package handler

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"order-service/internal/repository"
	"order-service/internal/service"
	"order-service/pb"
)

type OrderHandler struct {
	pb.UnimplementedOrderServiceServer
	svc *service.OrderService
}

func NewOrderHandler(svc *service.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

func (h *OrderHandler) CreateOffer(ctx context.Context, req *pb.CreateOfferRequest) (*pb.Offer, error) {
	offer := &repository.Offer{
		Type:          req.GetType(),
		ListingID:     req.GetListingId(),
		ListingTitle:  req.GetListingTitle(),
		BuyerID:       req.GetBuyerId(),
		BuyerName:     req.GetBuyerName(),
		SellerID:      req.GetSellerId(),
		SellerName:    req.GetSellerName(),
		Quantity:      req.GetQuantity(),
		Unit:          req.GetUnit(),
		ProposedPrice: req.GetProposedPrice(),
		Currency:      req.GetCurrency(),
		Message:       req.GetMessage(),
	}

	created, err := h.svc.CreateOffer(offer)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create offer: %v", err)
	}

	return offerToProto(created), nil
}

func (h *OrderHandler) GetOffer(ctx context.Context, req *pb.GetOfferRequest) (*pb.Offer, error) {
	offer, err := h.svc.GetOffer(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "offer not found: %v", err)
	}
	return offerToProto(offer), nil
}

func (h *OrderHandler) ListOffers(ctx context.Context, req *pb.ListOffersRequest) (*pb.ListOffersResponse, error) {
	offers, total, err := h.svc.ListOffers(req.GetUserId(), req.GetRole(), req.GetStatus(), int(req.GetPage()), int(req.GetPageSize()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list offers: %v", err)
	}

	pbOffers := make([]*pb.Offer, len(offers))
	for i, o := range offers {
		pbOffers[i] = offerToProto(&o)
	}

	return &pb.ListOffersResponse{
		Offers: pbOffers,
		Total:  int32(total),
	}, nil
}

func (h *OrderHandler) AcceptOffer(ctx context.Context, req *pb.AcceptOfferRequest) (*pb.Transaction, error) {
	tx, err := h.svc.AcceptOffer(req.GetOfferId(), req.GetActorId(), req.GetActorName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to accept offer: %v", err)
	}
	return transactionToProto(tx, nil), nil
}

func (h *OrderHandler) RejectOffer(ctx context.Context, req *pb.RejectOfferRequest) (*pb.Offer, error) {
	offer, err := h.svc.RejectOffer(req.GetOfferId(), req.GetActorId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to reject offer: %v", err)
	}
	return offerToProto(offer), nil
}

func (h *OrderHandler) CancelOffer(ctx context.Context, req *pb.CancelOfferRequest) (*pb.Offer, error) {
	offer, err := h.svc.CancelOffer(req.GetOfferId(), req.GetActorId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to cancel offer: %v", err)
	}
	return offerToProto(offer), nil
}

func (h *OrderHandler) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.Transaction, error) {
	tx, events, err := h.svc.GetTransaction(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "transaction not found: %v", err)
	}
	return transactionToProto(tx, events), nil
}

func (h *OrderHandler) ListTransactions(ctx context.Context, req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	transactions, total, err := h.svc.ListTransactions(req.GetUserId(), req.GetStatus(), int(req.GetPage()), int(req.GetPageSize()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list transactions: %v", err)
	}

	pbTxns := make([]*pb.Transaction, len(transactions))
	for i, t := range transactions {
		pbTxns[i] = transactionToProto(&t, nil)
	}

	return &pb.ListTransactionsResponse{
		Transactions: pbTxns,
		Total:        int32(total),
	}, nil
}

func (h *OrderHandler) UpdateTransactionStatus(ctx context.Context, req *pb.UpdateStatusRequest) (*pb.Transaction, error) {
	tx, err := h.svc.UpdateTransactionStatus(req.GetTransactionId(), req.GetNewStatus(), req.GetActorId(), req.GetActorName(), req.GetNote())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update transaction status: %v", err)
	}
	return transactionToProto(tx, nil), nil
}

func offerToProto(o *repository.Offer) *pb.Offer {
	return &pb.Offer{
		Id:            o.ID,
		Type:          o.Type,
		ListingId:     o.ListingID,
		ListingTitle:  o.ListingTitle,
		BuyerId:       o.BuyerID,
		BuyerName:     o.BuyerName,
		SellerId:      o.SellerID,
		SellerName:    o.SellerName,
		Quantity:      o.Quantity,
		Unit:          o.Unit,
		ProposedPrice: o.ProposedPrice,
		Currency:      o.Currency,
		Message:       o.Message,
		Status:        o.Status,
		CreatedAt:     o.CreatedAt.Format(time.RFC3339),
	}
}

func transactionToProto(tx *repository.Transaction, events []repository.TransactionEvent) *pb.Transaction {
	pbTx := &pb.Transaction{
		Id:             tx.ID,
		OfferId:        tx.OfferID,
		ListingTitle:   tx.ListingTitle,
		BuyerId:        tx.BuyerID,
		BuyerName:      tx.BuyerName,
		SellerId:       tx.SellerID,
		SellerName:     tx.SellerName,
		Quantity:       tx.Quantity,
		Unit:           tx.Unit,
		AgreedPrice:    tx.AgreedPrice,
		Currency:       tx.Currency,
		PaymentStatus:  tx.PaymentStatus,
		PaymentMethod:  tx.PaymentMethod,
		SettlementNote: tx.SettlementNote,
		Status:         tx.Status,
		CreatedAt:      tx.CreatedAt.Format(time.RFC3339),
	}

	if events != nil {
		pbEvents := make([]*pb.TransactionEvent, len(events))
		for i, e := range events {
			pbEvents[i] = &pb.TransactionEvent{
				Id:         e.ID,
				ActorId:    e.ActorID,
				ActorName:  e.ActorName,
				FromStatus: e.FromStatus,
				ToStatus:   e.ToStatus,
				Note:       e.Note,
				CreatedAt:  e.CreatedAt.Format(time.RFC3339),
			}
		}
		pbTx.Events = pbEvents
	}

	return pbTx
}
