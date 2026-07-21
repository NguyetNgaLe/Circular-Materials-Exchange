package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type Offer struct {
	ID            string
	Type          string
	ListingID     string
	ListingTitle  string
	BuyerID       string
	BuyerName     string
	SellerID      string
	SellerName    string
	Quantity      float64
	Unit          string
	ProposedPrice float64
	Currency      string
	Message       string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Transaction struct {
	ID             string
	OfferID        string
	ListingTitle   string
	BuyerID        string
	BuyerName      string
	SellerID       string
	SellerName     string
	Quantity       float64
	Unit           string
	AgreedPrice    float64
	Currency       string
	PaymentStatus  string
	PaymentMethod  string
	SettlementNote string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TransactionEvent struct {
	ID            string
	TransactionID string
	ActorID       string
	ActorName     string
	FromStatus    string
	ToStatus      string
	Note          string
	CreatedAt     time.Time
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOffer(offer *Offer) error {
	_, err := r.db.Exec(`INSERT INTO offers (id, type, listing_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, proposed_price, currency, message, status, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)`,
		offer.ID, offer.Type, offer.ListingID, offer.ListingTitle, offer.BuyerID, offer.BuyerName, offer.SellerID, offer.SellerName, offer.Quantity, offer.Unit, offer.ProposedPrice, offer.Currency, offer.Message, offer.Status, offer.CreatedAt, offer.UpdatedAt)
	return err
}

func (r *OrderRepository) FindOfferByID(id string) (*Offer, error) {
	var o Offer
	err := r.db.QueryRow(`SELECT id, type, listing_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, proposed_price, currency, message, status, created_at, updated_at FROM offers WHERE id=$1`, id).
		Scan(&o.ID, &o.Type, &o.ListingID, &o.ListingTitle, &o.BuyerID, &o.BuyerName, &o.SellerID, &o.SellerName, &o.Quantity, &o.Unit, &o.ProposedPrice, &o.Currency, &o.Message, &o.Status, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("offer not found: %w", err)
	}
	return &o, nil
}

func (r *OrderRepository) ListOffers(userID, role, status string, page, pageSize int) ([]Offer, int64, error) {
	where := "1=1"
	args := []interface{}{}
	argIdx := 1

	if userID != "" {
		if role == "buyer" {
			where += fmt.Sprintf(" AND buyer_id=$%d", argIdx)
			args = append(args, userID)
			argIdx++
		} else if role == "seller" {
			where += fmt.Sprintf(" AND seller_id=$%d", argIdx)
			args = append(args, userID)
			argIdx++
		} else {
			where += fmt.Sprintf(" AND (buyer_id=$%d OR seller_id=$%d)", argIdx, argIdx+1)
			args = append(args, userID, userID)
			argIdx += 2
		}
	}
	if status != "" {
		where += fmt.Sprintf(" AND status=$%d", argIdx)
		args = append(args, status)
		argIdx++
	}

	var total int64
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM offers WHERE %s", where)
	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	args = append(args, pageSize, offset)
	query := fmt.Sprintf("SELECT id, type, listing_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, proposed_price, currency, message, status, created_at, updated_at FROM offers WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", where, argIdx, argIdx+1)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var offers []Offer
	for rows.Next() {
		var o Offer
		if err := rows.Scan(&o.ID, &o.Type, &o.ListingID, &o.ListingTitle, &o.BuyerID, &o.BuyerName, &o.SellerID, &o.SellerName, &o.Quantity, &o.Unit, &o.ProposedPrice, &o.Currency, &o.Message, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, 0, err
		}
		offers = append(offers, o)
	}
	return offers, total, nil
}

func (r *OrderRepository) UpdateOffer(offer *Offer) error {
	_, err := r.db.Exec(`UPDATE offers SET type=$2, listing_id=$3, listing_title=$4, buyer_id=$5, buyer_name=$6, seller_id=$7, seller_name=$8, quantity=$9, unit=$10, proposed_price=$11, currency=$12, message=$13, status=$14, updated_at=$15 WHERE id=$1`,
		offer.ID, offer.Type, offer.ListingID, offer.ListingTitle, offer.BuyerID, offer.BuyerName, offer.SellerID, offer.SellerName, offer.Quantity, offer.Unit, offer.ProposedPrice, offer.Currency, offer.Message, offer.Status, offer.UpdatedAt)
	return err
}

func (r *OrderRepository) CreateTransaction(tx *Transaction) error {
	_, err := r.db.Exec(`INSERT INTO transactions (id, offer_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, agreed_price, currency, payment_status, payment_method, settlement_note, status, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`,
		tx.ID, tx.OfferID, tx.ListingTitle, tx.BuyerID, tx.BuyerName, tx.SellerID, tx.SellerName, tx.Quantity, tx.Unit, tx.AgreedPrice, tx.Currency, tx.PaymentStatus, tx.PaymentMethod, tx.SettlementNote, tx.Status, tx.CreatedAt, tx.UpdatedAt)
	return err
}

func (r *OrderRepository) FindTransactionByID(id string) (*Transaction, error) {
	var tx Transaction
	err := r.db.QueryRow(`SELECT id, offer_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, agreed_price, currency, payment_status, payment_method, settlement_note, status, created_at, updated_at FROM transactions WHERE id=$1`, id).
		Scan(&tx.ID, &tx.OfferID, &tx.ListingTitle, &tx.BuyerID, &tx.BuyerName, &tx.SellerID, &tx.SellerName, &tx.Quantity, &tx.Unit, &tx.AgreedPrice, &tx.Currency, &tx.PaymentStatus, &tx.PaymentMethod, &tx.SettlementNote, &tx.Status, &tx.CreatedAt, &tx.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("transaction not found: %w", err)
	}
	return &tx, nil
}

func (r *OrderRepository) ListTransactions(userID, status string, page, pageSize int) ([]Transaction, int64, error) {
	where := "1=1"
	args := []interface{}{}
	argIdx := 1

	if userID != "" {
		where += fmt.Sprintf(" AND (buyer_id=$%d OR seller_id=$%d)", argIdx, argIdx+1)
		args = append(args, userID, userID)
		argIdx += 2
	}
	if status != "" {
		where += fmt.Sprintf(" AND status=$%d", argIdx)
		args = append(args, status)
		argIdx++
	}

	var total int64
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM transactions WHERE %s", where)
	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	args = append(args, pageSize, offset)
	query := fmt.Sprintf("SELECT id, offer_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, agreed_price, currency, payment_status, payment_method, settlement_note, status, created_at, updated_at FROM transactions WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", where, argIdx, argIdx+1)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var tx Transaction
		if err := rows.Scan(&tx.ID, &tx.OfferID, &tx.ListingTitle, &tx.BuyerID, &tx.BuyerName, &tx.SellerID, &tx.SellerName, &tx.Quantity, &tx.Unit, &tx.AgreedPrice, &tx.Currency, &tx.PaymentStatus, &tx.PaymentMethod, &tx.SettlementNote, &tx.Status, &tx.CreatedAt, &tx.UpdatedAt); err != nil {
			return nil, 0, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, total, nil
}

func (r *OrderRepository) CreateTransactionEvent(event *TransactionEvent) error {
	_, err := r.db.Exec(`INSERT INTO transaction_events (id, transaction_id, actor_id, actor_name, from_status, to_status, note, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		event.ID, event.TransactionID, event.ActorID, event.ActorName, event.FromStatus, event.ToStatus, event.Note, event.CreatedAt)
	return err
}

func (r *OrderRepository) UpdateTransaction(tx *Transaction) error {
	_, err := r.db.Exec(`UPDATE transactions SET offer_id=$2, listing_title=$3, buyer_id=$4, buyer_name=$5, seller_id=$6, seller_name=$7, quantity=$8, unit=$9, agreed_price=$10, currency=$11, payment_status=$12, payment_method=$13, settlement_note=$14, status=$15, updated_at=$16 WHERE id=$1`,
		tx.ID, tx.OfferID, tx.ListingTitle, tx.BuyerID, tx.BuyerName, tx.SellerID, tx.SellerName, tx.Quantity, tx.Unit, tx.AgreedPrice, tx.Currency, tx.PaymentStatus, tx.PaymentMethod, tx.SettlementNote, tx.Status, tx.UpdatedAt)
	return err
}

func (r *OrderRepository) FindEventsByTransactionID(transactionID string) ([]TransactionEvent, error) {
	rows, err := r.db.Query(`SELECT id, transaction_id, actor_id, actor_name, from_status, to_status, note, created_at FROM transaction_events WHERE transaction_id=$1 ORDER BY created_at ASC`, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []TransactionEvent
	for rows.Next() {
		var e TransactionEvent
		if err := rows.Scan(&e.ID, &e.TransactionID, &e.ActorID, &e.ActorName, &e.FromStatus, &e.ToStatus, &e.Note, &e.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
