package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MonthlyRevenue struct {
	Month   string
	Revenue float64
}

type FinanceOverview struct {
	TotalIncome       float64
	MonthIncome       float64
	TotalTransactions int64
	WalletBalance     float64
	MonthlyRevenue    []MonthlyRevenue
}

type PlatformFee struct {
	ID                string
	TransactionID     string
	SellerID          string
	BuyerID           string
	TransactionAmount float64
	FeeRate           float64
	FeeAmount         float64
	FeeType           string
	Status            string
	CollectedAt       string
	CreatedAt         string
	ListingTitle      string
}

type PlatformWallet struct {
	Balance      float64
	TotalIncome  float64
	TotalExpense float64
}

type PlatformWalletTransaction struct {
	ID            string
	Type          string
	Amount        float64
	BalanceAfter  float64
	ReferenceType string
	ReferenceID   string
	Description   string
	CreatedAt     string
}

type EscrowTransaction struct {
	ID            string
	TransactionID string
	BuyerID       string
	BuyerName     string
	SellerID      string
	SellerName    string
	Amount        float64
	FeeRate       float64
	FeeAmount     float64
	SellerAmount  float64
	Status        string
	HoldUntil     string
	ReleasedAt    string
	CreatedAt     string
}

type SellerWallet struct {
	SellerID       string
	SellerName     string
	Balance        float64
	TotalEarned    float64
	TotalFees      float64
	TotalWithdrawn float64
}

type SellerWalletTransaction struct {
	ID            string
	Type          string
	Amount        float64
	BalanceAfter  float64
	ReferenceType string
	ReferenceID   string
	Description   string
	CreatedAt     string
}

type WithdrawalRequest struct {
	ID          string
	SellerID    string
	SellerName  string
	Amount      float64
	BankName    string
	BankAccount string
	BankOwner   string
	Status      string
	AdminNote   string
	ProcessedAt string
	CreatedAt   string
}

func (r *OrderRepository) GetFinanceOverview() (*FinanceOverview, error) {
	result := &FinanceOverview{}
	if err := r.db.QueryRow(`SELECT COALESCE(SUM(fee_amount),0) FROM platform_fees WHERE status='collected'`).Scan(&result.TotalIncome); err != nil {
		return nil, err
	}
	if err := r.db.QueryRow(`SELECT COALESCE(SUM(fee_amount),0) FROM platform_fees WHERE status='collected' AND created_at >= date_trunc('month',CURRENT_DATE)`).Scan(&result.MonthIncome); err != nil {
		return nil, err
	}
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM platform_fees`).Scan(&result.TotalTransactions); err != nil {
		return nil, err
	}
	if err := r.db.QueryRow(`SELECT COALESCE(balance,0) FROM platform_wallet LIMIT 1`).Scan(&result.WalletBalance); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	rows, err := r.db.Query(`SELECT to_char(date_trunc('month',created_at),'YYYY-MM'), COALESCE(SUM(fee_amount),0)
		FROM platform_fees WHERE created_at >= CURRENT_DATE - INTERVAL '6 months'
		GROUP BY date_trunc('month',created_at) ORDER BY date_trunc('month',created_at)`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item MonthlyRevenue
		if err := rows.Scan(&item.Month, &item.Revenue); err != nil {
			return nil, err
		}
		result.MonthlyRevenue = append(result.MonthlyRevenue, item)
	}
	return result, rows.Err()
}

func (r *OrderRepository) ListPlatformFees(page, pageSize int) ([]PlatformFee, int64, error) {
	var total int64
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM platform_fees`).Scan(&total); err != nil {
		return nil, 0, err
	}
	rows, err := r.db.Query(`SELECT f.id::text, COALESCE(f.transaction_id::text,''), f.seller_id::text, f.buyer_id::text,
		f.transaction_amount, f.fee_rate, f.fee_amount, f.fee_type, f.status, COALESCE(f.collected_at::text,''),
		f.created_at::text, COALESCE(t.listing_title,'') FROM platform_fees f
		LEFT JOIN transactions t ON t.id=f.transaction_id
		ORDER BY f.created_at DESC LIMIT $1 OFFSET $2`, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var items []PlatformFee
	for rows.Next() {
		var item PlatformFee
		if err := rows.Scan(&item.ID, &item.TransactionID, &item.SellerID, &item.BuyerID,
			&item.TransactionAmount, &item.FeeRate, &item.FeeAmount, &item.FeeType,
			&item.Status, &item.CollectedAt, &item.CreatedAt, &item.ListingTitle); err != nil {
			return nil, 0, err
		}
		items = append(items, item)
	}
	return items, total, rows.Err()
}

func (r *OrderRepository) GetPlatformWallet() (*PlatformWallet, error) {
	item := &PlatformWallet{}
	err := r.db.QueryRow(`SELECT COALESCE(balance,0),COALESCE(total_income,0),COALESCE(total_expense,0) FROM platform_wallet LIMIT 1`).
		Scan(&item.Balance, &item.TotalIncome, &item.TotalExpense)
	return item, err
}

func (r *OrderRepository) ListPlatformWalletTransactions(page, pageSize int) ([]PlatformWalletTransaction, int64, error) {
	var total int64
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM wallet_transactions`).Scan(&total); err != nil {
		return nil, 0, err
	}
	rows, err := r.db.Query(`SELECT id::text,type,amount,COALESCE(balance_after,0),COALESCE(reference_type,''),
		COALESCE(reference_id::text,''),COALESCE(description,''),created_at::text
		FROM wallet_transactions ORDER BY created_at DESC LIMIT $1 OFFSET $2`, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var items []PlatformWalletTransaction
	for rows.Next() {
		var item PlatformWalletTransaction
		if err := rows.Scan(&item.ID, &item.Type, &item.Amount, &item.BalanceAfter,
			&item.ReferenceType, &item.ReferenceID, &item.Description, &item.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, item)
	}
	return items, total, rows.Err()
}

func (r *OrderRepository) CollectFee(transactionID, sellerID, buyerID string, amount, feeRate float64) (string, float64, float64, error) {
	if feeRate <= 0 {
		feeRate = 0.02
	}
	feeID := uuid.New().String()
	feeAmount := amount * feeRate
	tx, err := r.db.Begin()
	if err != nil {
		return "", 0, 0, err
	}
	defer tx.Rollback()
	if _, err = tx.Exec(`INSERT INTO platform_fees (id,transaction_id,seller_id,buyer_id,transaction_amount,fee_rate,fee_amount,fee_type,status,collected_at)
		VALUES ($1,NULLIF($2,'')::uuid,$3,$4,$5,$6,$7,'transaction','collected',NOW())`,
		feeID, transactionID, sellerID, buyerID, amount, feeRate, feeAmount); err != nil {
		return "", 0, 0, err
	}
	var walletID string
	var balance float64
	if err = tx.QueryRow(`UPDATE platform_wallet SET balance=balance+$1,total_income=total_income+$1,updated_at=NOW()
		RETURNING id::text,balance`, feeAmount).Scan(&walletID, &balance); err != nil {
		return "", 0, 0, err
	}
	if _, err = tx.Exec(`INSERT INTO wallet_transactions (id,wallet_id,type,amount,reference_type,reference_id,description,balance_after)
		VALUES ($1,$2,'credit',$3,'transaction_fee',NULLIF($4,'')::uuid,'Phí giao dịch',$5)`,
		uuid.New().String(), walletID, feeAmount, transactionID, balance); err != nil {
		return "", 0, 0, err
	}
	if err = tx.Commit(); err != nil {
		return "", 0, 0, err
	}
	return feeID, feeAmount, balance, nil
}

func (r *OrderRepository) CreateEscrow(transactionID, buyerID, buyerName, sellerID, sellerName string, amount float64, holdDays int, creditPlatform bool) (*EscrowTransaction, error) {
	if holdDays <= 0 {
		holdDays = 7
	}
	feeRate := 0.02
	feeAmount := amount * feeRate
	item := &EscrowTransaction{
		ID: uuid.New().String(), TransactionID: transactionID, BuyerID: buyerID, BuyerName: buyerName,
		SellerID: sellerID, SellerName: sellerName, Amount: amount, FeeRate: feeRate,
		FeeAmount: feeAmount, SellerAmount: amount - feeAmount, Status: "holding",
		HoldUntil: time.Now().Add(time.Duration(holdDays) * 24 * time.Hour).Format(time.RFC3339),
	}
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	if _, err = tx.Exec(`INSERT INTO escrow_transactions
		(id,transaction_id,buyer_id,buyer_name,seller_id,seller_name,amount,fee_rate,fee_amount,seller_amount,status,hold_until)
		VALUES ($1,NULLIF($2,'')::uuid,$3,$4,$5,$6,$7,$8,$9,$10,'holding',$11)`,
		item.ID, transactionID, buyerID, buyerName, sellerID, sellerName, amount,
		feeRate, feeAmount, item.SellerAmount, item.HoldUntil); err != nil {
		return nil, err
	}
	if creditPlatform {
		if _, err = tx.Exec(`UPDATE platform_wallet SET balance=balance+$1,total_income=total_income+$1,updated_at=NOW()`, feeAmount); err != nil {
			return nil, err
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return r.FindEscrowByID(item.ID)
}

func (r *OrderRepository) AttachPendingEscrow(transactionID, buyerID, sellerID string, amount float64) error {
	result, err := r.db.Exec(`UPDATE escrow_transactions SET transaction_id=$1
		WHERE id=(SELECT id FROM escrow_transactions WHERE transaction_id IS NULL AND buyer_id=$2 AND seller_id=$3
			AND amount=$4 AND status='holding' ORDER BY created_at DESC LIMIT 1)`,
		transactionID, buyerID, sellerID, amount)
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		_, err = r.CreateEscrow(transactionID, buyerID, "", sellerID, "", amount, 7, false)
	}
	return err
}

func (r *OrderRepository) FindEscrowByID(id string) (*EscrowTransaction, error) {
	item := &EscrowTransaction{}
	err := r.db.QueryRow(`SELECT id::text,COALESCE(transaction_id::text,''),buyer_id::text,COALESCE(buyer_name,''),
		seller_id::text,COALESCE(seller_name,''),amount,fee_rate,fee_amount,seller_amount,status,
		hold_until::text,COALESCE(released_at::text,''),created_at::text FROM escrow_transactions WHERE id=$1`, id).
		Scan(&item.ID, &item.TransactionID, &item.BuyerID, &item.BuyerName, &item.SellerID,
			&item.SellerName, &item.Amount, &item.FeeRate, &item.FeeAmount, &item.SellerAmount,
			&item.Status, &item.HoldUntil, &item.ReleasedAt, &item.CreatedAt)
	return item, err
}

func (r *OrderRepository) ListEscrows(page, pageSize int) ([]EscrowTransaction, int64, error) {
	var total int64
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM escrow_transactions`).Scan(&total); err != nil {
		return nil, 0, err
	}
	rows, err := r.db.Query(`SELECT id::text,COALESCE(transaction_id::text,''),buyer_id::text,COALESCE(buyer_name,''),
		seller_id::text,COALESCE(seller_name,''),amount,fee_rate,fee_amount,seller_amount,status,
		hold_until::text,COALESCE(released_at::text,''),created_at::text
		FROM escrow_transactions ORDER BY created_at DESC LIMIT $1 OFFSET $2`, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var items []EscrowTransaction
	for rows.Next() {
		var item EscrowTransaction
		if err := rows.Scan(&item.ID, &item.TransactionID, &item.BuyerID, &item.BuyerName,
			&item.SellerID, &item.SellerName, &item.Amount, &item.FeeRate, &item.FeeAmount,
			&item.SellerAmount, &item.Status, &item.HoldUntil, &item.ReleasedAt, &item.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, item)
	}
	return items, total, rows.Err()
}

func (r *OrderRepository) ReleaseEscrow(id string, collectPlatformFee bool) (*EscrowTransaction, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var item EscrowTransaction
	if err = tx.QueryRow(`SELECT id::text,COALESCE(transaction_id::text,''),buyer_id::text,COALESCE(buyer_name,''),
		seller_id::text,COALESCE(seller_name,''),amount,fee_rate,fee_amount,seller_amount,status,
		hold_until::text,COALESCE(released_at::text,''),created_at::text
		FROM escrow_transactions WHERE id=$1 FOR UPDATE`, id).
		Scan(&item.ID, &item.TransactionID, &item.BuyerID, &item.BuyerName, &item.SellerID,
			&item.SellerName, &item.Amount, &item.FeeRate, &item.FeeAmount, &item.SellerAmount,
			&item.Status, &item.HoldUntil, &item.ReleasedAt, &item.CreatedAt); err != nil {
		return nil, err
	}
	if item.Status == "released" {
		return &item, nil
	}
	if item.Status != "holding" {
		return nil, fmt.Errorf("escrow is not holding")
	}
	if _, err = tx.Exec(`UPDATE escrow_transactions SET status='released',released_at=NOW() WHERE id=$1`, id); err != nil {
		return nil, err
	}
	if _, err = tx.Exec(`INSERT INTO seller_wallet (seller_id,seller_name,balance,total_earned,total_fees_paid)
		VALUES ($1,$2,$3,$3,$4) ON CONFLICT (seller_id) DO UPDATE SET
		balance=seller_wallet.balance+$3,total_earned=seller_wallet.total_earned+$3,
		total_fees_paid=seller_wallet.total_fees_paid+$4,updated_at=NOW()`,
		item.SellerID, item.SellerName, item.SellerAmount, item.FeeAmount); err != nil {
		return nil, err
	}
	var balance float64
	if err = tx.QueryRow(`SELECT balance FROM seller_wallet WHERE seller_id=$1`, item.SellerID).Scan(&balance); err != nil {
		return nil, err
	}
	if _, err = tx.Exec(`INSERT INTO seller_wallet_transactions
		(id,seller_id,type,amount,balance_after,reference_type,reference_id,description)
		VALUES ($1,$2,'credit',$3,$4,'escrow_release',$5,'Giải phóng escrow')`,
		uuid.New().String(), item.SellerID, item.SellerAmount, balance, item.ID); err != nil {
		return nil, err
	}
	if collectPlatformFee {
		if _, err = tx.Exec(`INSERT INTO platform_fees
			(id,transaction_id,seller_id,buyer_id,transaction_amount,fee_rate,fee_amount,fee_type,status,collected_at)
			VALUES ($1,NULLIF($2,'')::uuid,$3,$4,$5,$6,$7,'transaction','collected',NOW())
			ON CONFLICT DO NOTHING`, uuid.New().String(), item.TransactionID, item.SellerID,
			item.BuyerID, item.Amount, item.FeeRate, item.FeeAmount); err != nil {
			return nil, err
		}
		if _, err = tx.Exec(`UPDATE platform_wallet SET balance=balance+$1,total_income=total_income+$1,updated_at=NOW()`, item.FeeAmount); err != nil {
			return nil, err
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return r.FindEscrowByID(id)
}

func (r *OrderRepository) ReleaseEscrowByTransaction(transactionID string) (*EscrowTransaction, error) {
	var id string
	err := r.db.QueryRow(`SELECT id::text FROM escrow_transactions WHERE transaction_id=$1 AND status='holding' ORDER BY created_at DESC LIMIT 1`, transactionID).Scan(&id)
	if err != nil {
		return nil, err
	}
	return r.ReleaseEscrow(id, true)
}

func (r *OrderRepository) GetSellerWallet(sellerID string) (*SellerWallet, error) {
	item := &SellerWallet{SellerID: sellerID}
	err := r.db.QueryRow(`SELECT seller_id::text,COALESCE(seller_name,''),COALESCE(balance,0),
		COALESCE(total_earned,0),COALESCE(total_fees_paid,0),COALESCE(total_withdrawn,0)
		FROM seller_wallet WHERE seller_id=$1`, sellerID).
		Scan(&item.SellerID, &item.SellerName, &item.Balance, &item.TotalEarned, &item.TotalFees, &item.TotalWithdrawn)
	if err == sql.ErrNoRows {
		return item, nil
	}
	return item, err
}

func (r *OrderRepository) ListSellerWalletTransactions(sellerID string, page, pageSize int) ([]SellerWalletTransaction, int64, error) {
	var total int64
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM seller_wallet_transactions WHERE seller_id=$1`, sellerID).Scan(&total); err != nil {
		return nil, 0, err
	}
	rows, err := r.db.Query(`SELECT id::text,type,amount,COALESCE(balance_after,0),COALESCE(reference_type,''),
		COALESCE(reference_id::text,''),COALESCE(description,''),created_at::text
		FROM seller_wallet_transactions WHERE seller_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`,
		sellerID, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var items []SellerWalletTransaction
	for rows.Next() {
		var item SellerWalletTransaction
		if err := rows.Scan(&item.ID, &item.Type, &item.Amount, &item.BalanceAfter,
			&item.ReferenceType, &item.ReferenceID, &item.Description, &item.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, item)
	}
	return items, total, rows.Err()
}

func (r *OrderRepository) CreateWithdrawal(sellerID, sellerName string, amount float64, bankName, bankAccount, bankOwner string) (*WithdrawalRequest, error) {
	if amount <= 0 {
		return nil, fmt.Errorf("amount must be positive")
	}
	item := &WithdrawalRequest{
		ID: uuid.New().String(), SellerID: sellerID, SellerName: sellerName, Amount: amount,
		BankName: bankName, BankAccount: bankAccount, BankOwner: bankOwner, Status: "pending",
	}
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var balance float64
	if err = tx.QueryRow(`SELECT balance FROM seller_wallet WHERE seller_id=$1 FOR UPDATE`, sellerID).Scan(&balance); err != nil {
		return nil, err
	}
	if balance < amount {
		return nil, fmt.Errorf("insufficient balance")
	}
	if _, err = tx.Exec(`UPDATE seller_wallet SET balance=balance-$1,updated_at=NOW() WHERE seller_id=$2`, amount, sellerID); err != nil {
		return nil, err
	}
	if _, err = tx.Exec(`INSERT INTO withdrawal_requests
		(id,seller_id,seller_name,amount,bank_name,bank_account,bank_owner,status)
		VALUES ($1,$2,$3,$4,$5,$6,$7,'pending')`,
		item.ID, sellerID, sellerName, amount, bankName, bankAccount, bankOwner); err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return r.FindWithdrawalByID(item.ID)
}

func (r *OrderRepository) FindWithdrawalByID(id string) (*WithdrawalRequest, error) {
	item := &WithdrawalRequest{}
	err := r.db.QueryRow(`SELECT id::text,seller_id::text,COALESCE(seller_name,''),amount,
		COALESCE(bank_name,''),COALESCE(bank_account,''),COALESCE(bank_owner,''),status,
		COALESCE(admin_note,''),COALESCE(processed_at::text,''),created_at::text
		FROM withdrawal_requests WHERE id=$1`, id).
		Scan(&item.ID, &item.SellerID, &item.SellerName, &item.Amount, &item.BankName,
			&item.BankAccount, &item.BankOwner, &item.Status, &item.AdminNote,
			&item.ProcessedAt, &item.CreatedAt)
	return item, err
}

func (r *OrderRepository) ListWithdrawals(sellerID string, page, pageSize int) ([]WithdrawalRequest, int64, error) {
	where := ""
	args := []interface{}{}
	if sellerID != "" {
		where = " WHERE seller_id=$1"
		args = append(args, sellerID)
	}
	var total int64
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM withdrawal_requests`+where, args...).Scan(&total); err != nil {
		return nil, 0, err
	}
	limitArg := len(args) + 1
	query := fmt.Sprintf(`SELECT id::text,seller_id::text,COALESCE(seller_name,''),amount,
		COALESCE(bank_name,''),COALESCE(bank_account,''),COALESCE(bank_owner,''),status,
		COALESCE(admin_note,''),COALESCE(processed_at::text,''),created_at::text
		FROM withdrawal_requests%s ORDER BY created_at DESC LIMIT $%d OFFSET $%d`, where, limitArg, limitArg+1)
	args = append(args, pageSize, (page-1)*pageSize)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var items []WithdrawalRequest
	for rows.Next() {
		var item WithdrawalRequest
		if err := rows.Scan(&item.ID, &item.SellerID, &item.SellerName, &item.Amount,
			&item.BankName, &item.BankAccount, &item.BankOwner, &item.Status,
			&item.AdminNote, &item.ProcessedAt, &item.CreatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, item)
	}
	return items, total, rows.Err()
}

func (r *OrderRepository) ProcessWithdrawal(id, status, reason string) (*WithdrawalRequest, error) {
	if status != "completed" && status != "rejected" {
		return nil, fmt.Errorf("invalid withdrawal status")
	}
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var sellerID string
	var amount float64
	var currentStatus string
	if err = tx.QueryRow(`SELECT seller_id::text,amount,status FROM withdrawal_requests WHERE id=$1 FOR UPDATE`, id).
		Scan(&sellerID, &amount, &currentStatus); err != nil {
		return nil, err
	}
	if currentStatus != "pending" {
		return nil, fmt.Errorf("withdrawal is not pending")
	}
	if _, err = tx.Exec(`UPDATE withdrawal_requests SET status=$1,admin_note=$2,processed_at=NOW() WHERE id=$3`, status, reason, id); err != nil {
		return nil, err
	}
	if status == "completed" {
		_, err = tx.Exec(`UPDATE seller_wallet SET total_withdrawn=total_withdrawn+$1,updated_at=NOW() WHERE seller_id=$2`, amount, sellerID)
	} else {
		_, err = tx.Exec(`UPDATE seller_wallet SET balance=balance+$1,updated_at=NOW() WHERE seller_id=$2`, amount, sellerID)
	}
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return r.FindWithdrawalByID(id)
}
