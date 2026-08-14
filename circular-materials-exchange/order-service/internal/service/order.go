package service

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"order-service/internal/repository"
)

type NATSConn interface {
	Publish(subject string, data []byte) error
}

type OrderService struct {
	repo *repository.OrderRepository
	nc   NATSConn
}

func NewOrderService(repo *repository.OrderRepository, nc NATSConn) *OrderService {
	return &OrderService{repo: repo, nc: nc}
}

func (s *OrderService) CreateOffer(req *repository.Offer) (*repository.Offer, error) {
	req.ID = uuid.New().String()
	req.Status = "pending"
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()
	if err := s.repo.CreateOffer(req); err != nil {
		return nil, fmt.Errorf("failed to create offer: %w", err)
	}
	amount := req.Quantity * req.ProposedPrice
	if _, err := s.repo.CreateEscrow("", req.BuyerID, req.BuyerName, req.SellerID, req.SellerName, amount, 3, false); err != nil {
		return nil, fmt.Errorf("failed to create offer escrow: %w", err)
	}
	s.publishNATS("cme.orders.offer.created", map[string]interface{}{
		"offer_id":       req.ID,
		"buyer_id":       req.BuyerID,
		"buyer_name":     req.BuyerName,
		"seller_id":      req.SellerID,
		"seller_name":    req.SellerName,
		"listing_title":  req.ListingTitle,
		"quantity":       req.Quantity,
		"unit":           req.Unit,
		"proposed_price": req.ProposedPrice,
	})
	return req, nil
}

func (s *OrderService) GetOffer(id string) (*repository.Offer, error) {
	return s.repo.FindOfferByID(id)
}

func (s *OrderService) ListOffers(userID, role, status string, page, pageSize int) ([]repository.Offer, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return s.repo.ListOffers(userID, role, status, page, pageSize)
}

func (s *OrderService) AcceptOffer(offerID, actorID, actorName string) (*repository.Transaction, error) {
	offer, err := s.repo.FindOfferByID(offerID)
	if err != nil {
		return nil, fmt.Errorf("offer not found: %w", err)
	}
	if offer.Status != "pending" {
		return nil, fmt.Errorf("offer is not in pending status, current status: %s", offer.Status)
	}

	offer.Status = "accepted"
	offer.UpdatedAt = time.Now()
	if err := s.repo.UpdateOffer(offer); err != nil {
		return nil, fmt.Errorf("failed to update offer: %w", err)
	}

	tx := &repository.Transaction{
		ID:             uuid.New().String(),
		OfferID:        offer.ID,
		ListingTitle:   offer.ListingTitle,
		BuyerID:        offer.BuyerID,
		BuyerName:      offer.BuyerName,
		SellerID:       offer.SellerID,
		SellerName:     offer.SellerName,
		Quantity:       offer.Quantity,
		Unit:           offer.Unit,
		AgreedPrice:    offer.ProposedPrice,
		Currency:       offer.Currency,
		PaymentStatus:  "bypassed_demo",
		PaymentMethod:  "manual_offline",
		SettlementNote: "Thanh toán được thực hiện ngoài hệ thống trong phạm vi prototype",
		Status:         "confirmed",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := s.repo.CreateTransaction(tx); err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}
	if err := s.repo.AttachPendingEscrow(tx.ID, offer.BuyerID, offer.SellerID, offer.Quantity*offer.ProposedPrice); err != nil {
		return nil, fmt.Errorf("failed to attach escrow: %w", err)
	}

	event := &repository.TransactionEvent{
		ID:            uuid.New().String(),
		TransactionID: tx.ID,
		ActorID:       actorID,
		ActorName:     actorName,
		FromStatus:    "offer.accepted",
		ToStatus:      "transaction.confirmed",
		Note:          "Giao dịch được tạo tự động khi người bán chấp nhận đề nghị",
		CreatedAt:     time.Now(),
	}
	if err := s.repo.CreateTransactionEvent(event); err != nil {
		log.Printf("failed to create transaction event: %v", err)
	}

	s.publishNATS("cme.orders.transaction.created", map[string]interface{}{
		"transaction_id": tx.ID,
		"offer_id":       offer.ID,
		"buyer_id":       offer.BuyerID,
		"buyer_name":     offer.BuyerName,
		"seller_id":      offer.SellerID,
		"seller_name":    offer.SellerName,
		"listing_title":  offer.ListingTitle,
		"agreed_price":   offer.ProposedPrice,
		"currency":       offer.Currency,
		"quantity":       offer.Quantity,
		"unit":           offer.Unit,
	})

	return tx, nil
}

func (s *OrderService) RejectOffer(offerID, actorID string) (*repository.Offer, error) {
	offer, err := s.repo.FindOfferByID(offerID)
	if err != nil {
		return nil, fmt.Errorf("offer not found: %w", err)
	}
	if offer.Status != "pending" {
		return nil, fmt.Errorf("offer is not in pending status, current status: %s", offer.Status)
	}

	offer.Status = "rejected"
	offer.UpdatedAt = time.Now()
	if err := s.repo.UpdateOffer(offer); err != nil {
		return nil, fmt.Errorf("failed to update offer: %w", err)
	}
	return offer, nil
}

func (s *OrderService) CancelOffer(offerID, actorID string) (*repository.Offer, error) {
	offer, err := s.repo.FindOfferByID(offerID)
	if err != nil {
		return nil, fmt.Errorf("offer not found: %w", err)
	}
	if offer.Status != "pending" {
		return nil, fmt.Errorf("offer is not in pending status, current status: %s", offer.Status)
	}

	offer.Status = "cancelled"
	offer.UpdatedAt = time.Now()
	if err := s.repo.UpdateOffer(offer); err != nil {
		return nil, fmt.Errorf("failed to update offer: %w", err)
	}
	return offer, nil
}

func (s *OrderService) GetTransaction(id string) (*repository.Transaction, []repository.TransactionEvent, error) {
	tx, err := s.repo.FindTransactionByID(id)
	if err != nil {
		return nil, nil, fmt.Errorf("transaction not found: %w", err)
	}
	events, err := s.repo.FindEventsByTransactionID(id)
	if err != nil {
		return tx, nil, nil
	}
	return tx, events, nil
}

func (s *OrderService) ListTransactions(userID, status string, page, pageSize int) ([]repository.Transaction, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return s.repo.ListTransactions(userID, status, page, pageSize)
}

func (s *OrderService) UpdateTransactionStatus(transactionID, newStatus, actorID, actorName, note string) (*repository.Transaction, error) {
	tx, err := s.repo.FindTransactionByID(transactionID)
	if err != nil {
		return nil, fmt.Errorf("transaction not found: %w", err)
	}

	oldStatus := tx.Status
	tx.Status = newStatus
	tx.UpdatedAt = time.Now()
	if err := s.repo.UpdateTransaction(tx); err != nil {
		return nil, fmt.Errorf("failed to update transaction: %w", err)
	}

	event := &repository.TransactionEvent{
		ID:            uuid.New().String(),
		TransactionID: tx.ID,
		ActorID:       actorID,
		ActorName:     actorName,
		FromStatus:    "transaction." + oldStatus,
		ToStatus:      "transaction." + newStatus,
		Note:          note,
		CreatedAt:     time.Now(),
	}
	if err := s.repo.CreateTransactionEvent(event); err != nil {
		log.Printf("failed to create transaction event: %v", err)
	}

	if newStatus == "completed" {
		if _, err := s.repo.ReleaseEscrowByTransaction(tx.ID); err != nil {
			log.Printf("failed to release transaction escrow: %v", err)
		}
		s.publishNATS("cme.orders.transaction.completed", map[string]interface{}{
			"transaction_id": tx.ID,
			"buyer_id":       tx.BuyerID,
			"seller_id":      tx.SellerID,
			"listing_title":  tx.ListingTitle,
			"agreed_price":   tx.AgreedPrice,
			"currency":       tx.Currency,
		})
	} else if newStatus == "in_progress" {
		s.publishNATS("cme.orders.transaction.in_progress", map[string]interface{}{
			"transaction_id": tx.ID,
			"buyer_id":       tx.BuyerID,
			"buyer_name":     tx.BuyerName,
			"seller_id":      tx.SellerID,
			"seller_name":    tx.SellerName,
			"listing_title":  tx.ListingTitle,
		})
	}

	return tx, nil
}

func (s *OrderService) publishNATS(subject string, data interface{}) {
	if s.nc == nil {
		return
	}
	payload, err := json.Marshal(data)
	if err != nil {
		log.Printf("failed to marshal NATS message: %v", err)
		return
	}
	if err := s.nc.Publish(subject, payload); err != nil {
		log.Printf("failed to publish NATS message to %s: %v", subject, err)
	}
}
