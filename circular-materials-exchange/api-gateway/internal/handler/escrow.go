package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	_ "github.com/lib/pq"
)

type EscrowHandler struct {
	conn *grpc.ClientConn
	db   *sql.DB
}

func NewEscrowHandler(conn *grpc.ClientConn) *EscrowHandler {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5433")
	dbName := "order_db"
	dbUser := getEnv("DB_USER", "cme")
	dbPass := getEnv("DB_PASSWORD", "cme_secret_2024")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return &EscrowHandler{conn: conn, db: nil}
	}
	return &EscrowHandler{conn: conn, db: db}
}

// Tao escrow khi transaction duoc tao
func (h *EscrowHandler) CreateEscrow(c *gin.Context) {
	var req struct {
		TransactionID string  `json:"transactionId" binding:"required"`
		BuyerID       string  `json:"buyerId" binding:"required"`
		BuyerName     string  `json:"buyerName"`
		SellerID      string  `json:"sellerId" binding:"required"`
		SellerName    string  `json:"sellerName"`
		Amount        float64 `json:"amount" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": "escrow-demo"}})
		return
	}

	feeRate := 0.02
	feeAmount := req.Amount * feeRate
	sellerAmount := req.Amount - feeAmount
	holdUntil := time.Now().AddDate(0, 0, 3) // 3 ngay

	id := uuid.New().String()
	_, err := h.db.Exec(`INSERT INTO escrow_transactions (id, transaction_id, buyer_id, buyer_name, seller_id, seller_name, amount, fee_rate, fee_amount, seller_amount, status, hold_until) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,'holding',$11)`,
		id, req.TransactionID, req.BuyerID, req.BuyerName, req.SellerID, req.SellerName, req.Amount, feeRate, feeAmount, sellerAmount, holdUntil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Cap nhat platform_wallet
	h.db.Exec("UPDATE platform_wallet SET balance = balance + $1, total_income = total_income + $1, updated_at = NOW()", feeAmount)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id": id, "amount": req.Amount, "feeAmount": feeAmount,
			"sellerAmount": sellerAmount, "status": "holding", "holdUntil": holdUntil,
		},
	})
}

// Giai ngan tu dong (cron job hoac goi thu cong)
func (h *EscrowHandler) ReleaseEscrow(c *gin.Context) {
	escrowID := c.Param("id")

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"status": "released"}})
		return
	}

	// Lay thong tin escrow
	var sellerID, sellerName string
	var sellerAmount, feeAmount float64
	err := h.db.QueryRow("SELECT seller_id, seller_name, seller_amount, fee_amount FROM escrow_transactions WHERE id=$1 AND status='holding'", escrowID).Scan(
		&sellerID, &sellerName, &sellerAmount, &feeAmount)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Escrow not found or already released"})
		return
	}

	// Cap nhat escrow status
	h.db.Exec("UPDATE escrow_transactions SET status='released', released_at=NOW() WHERE id=$1", escrowID)

	// Cap nhat seller_wallet
	h.db.Exec(`INSERT INTO seller_wallet (seller_id, seller_name, balance, total_earned, total_fees_paid) 
		VALUES ($1, $2, $3, $3, $4) 
		ON CONFLICT (seller_id) DO UPDATE SET 
		balance = seller_wallet.balance + $3, 
		total_earned = seller_wallet.total_earned + $3,
		total_fees_paid = seller_wallet.total_fees_paid + $4,
		updated_at = NOW()`,
		sellerID, sellerName, sellerAmount, feeAmount)

	// Lay so du moi
	var newBalance float64
	h.db.QueryRow("SELECT balance FROM seller_wallet WHERE seller_id=$1", sellerID).Scan(&newBalance)

	// Tao lich su vi
	h.db.Exec(`INSERT INTO seller_wallet_transactions (seller_id, type, amount, balance_after, reference_type, reference_id, description) VALUES ($1,'credit',$2,$3,'escrow_release',$4,$5)`,
		sellerID, sellerAmount, newBalance, escrowID, fmt.Sprintf("Nhan tien tu giao dich (tru phi %.0f%%)", feeAmount/sellerAmount*100))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"escrowId": escrowID, "status": "released",
			"sellerAmount": sellerAmount, "feeAmount": feeAmount,
			"sellerBalance": newBalance,
		},
	})
}

// Giai ngan tu dong cho cac escrow da qua thoi gian giu
func (h *EscrowHandler) AutoReleaseEscrows() {
	if h.db == nil {
		return
	}

	rows, err := h.db.Query("SELECT id FROM escrow_transactions WHERE status='holding' AND hold_until <= NOW()")
	if err != nil {
		return
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		rows.Scan(&id)
		ids = append(ids, id)
	}

	for _, id := range ids {
		// Lay thong tin escrow
		var sellerID, sellerName string
		var sellerAmount, feeAmount float64
		err := h.db.QueryRow("SELECT seller_id, seller_name, seller_amount, fee_amount FROM escrow_transactions WHERE id=$1", id).Scan(
			&sellerID, &sellerName, &sellerAmount, &feeAmount)
		if err != nil {
			continue
		}

		// Cap nhat escrow status
		h.db.Exec("UPDATE escrow_transactions SET status='released', released_at=NOW() WHERE id=$1", id)

		// Cap nhat seller_wallet
		h.db.Exec(`INSERT INTO seller_wallet (seller_id, seller_name, balance, total_earned, total_fees_paid) 
			VALUES ($1, $2, $3, $3, $4) 
			ON CONFLICT (seller_id) DO UPDATE SET 
			balance = seller_wallet.balance + $3, 
			total_earned = seller_wallet.total_earned + $3,
			total_fees_paid = seller_wallet.total_fees_paid + $4,
			updated_at = NOW()`,
			sellerID, sellerName, sellerAmount, feeAmount)
	}
}

// Xem vi seller
func (h *EscrowHandler) GetSellerWallet(c *gin.Context) {
	userID, _ := c.Get("user_id")

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
			"balance": 0, "totalEarned": 0, "totalFeesPaid": 0, "totalWithdrawn": 0,
		}})
		return
	}

	var balance, totalEarned, totalFeesPaid, totalWithdrawn float64
	err := h.db.QueryRow("SELECT COALESCE(balance,0), COALESCE(total_earned,0), COALESCE(total_fees_paid,0), COALESCE(total_withdrawn,0) FROM seller_wallet WHERE seller_id=$1", userID).Scan(
		&balance, &totalEarned, &totalFeesPaid, &totalWithdrawn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
			"balance": 0, "totalEarned": 0, "totalFeesPaid": 0, "totalWithdrawn": 0,
		}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"balance": balance, "totalEarned": totalEarned,
			"totalFeesPaid": totalFeesPaid, "totalWithdrawn": totalWithdrawn,
		},
	})
}

// Lich su giao dich vi seller
func (h *EscrowHandler) GetSellerWalletTransactions(c *gin.Context) {
	userID, _ := c.Get("user_id")

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": []gin.H{}}})
		return
	}

	rows, err := h.db.Query("SELECT id, type, amount, balance_after, reference_type, description, created_at::text FROM seller_wallet_transactions WHERE seller_id=$1 ORDER BY created_at DESC LIMIT 50", userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": []gin.H{}}})
		return
	}
	defer rows.Close()

	var txs []gin.H
	for rows.Next() {
		var id, txType, refType, desc, createdAt string
		var amount, balanceAfter float64
		rows.Scan(&id, &txType, &amount, &balanceAfter, &refType, &desc, &createdAt)
		txs = append(txs, gin.H{
			"id": id, "type": txType, "amount": amount, "balanceAfter": balanceAfter,
			"referenceType": refType, "description": desc, "createdAt": createdAt,
		})
	}
	if txs == nil {
		txs = []gin.H{}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": txs}})
}

// Yeu cau rut tien
func (h *EscrowHandler) CreateWithdrawal(c *gin.Context) {
	var req struct {
		Amount      float64 `json:"amount" binding:"required"`
		BankName    string  `json:"bankName" binding:"required"`
		BankAccount string  `json:"bankAccount" binding:"required"`
		BankOwner   string  `json:"bankOwner" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	userName, _ := c.Get("email")

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": "wd-demo"}})
		return
	}

	// Kiem tra so du
	var balance float64
	err := h.db.QueryRow("SELECT COALESCE(balance,0) FROM seller_wallet WHERE seller_id=$1", userID).Scan(&balance)
	if err != nil || balance < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "So du khong du"})
		return
	}

	// Tru so du tam thoi
	h.db.Exec("UPDATE seller_wallet SET balance = balance - $1, updated_at = NOW() WHERE seller_id=$1", req.Amount, userID)

	// Tao yeu cau rut tien
	id := uuid.New().String()
	h.db.Exec(`INSERT INTO withdrawal_requests (id, seller_id, seller_name, amount, bank_name, bank_account, bank_owner) VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		id, userID, userName, req.Amount, req.BankName, req.BankAccount, req.BankOwner)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id": id, "amount": req.Amount, "status": "pending",
			"bankName": req.BankName, "bankAccount": req.BankAccount,
		},
	})
}

// Lich su rut tien cua seller
func (h *EscrowHandler) GetSellerWithdrawals(c *gin.Context) {
	userID, _ := c.Get("user_id")

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"withdrawals": []gin.H{}}})
		return
	}

	rows, err := h.db.Query("SELECT id, amount, bank_name, bank_account, bank_owner, status, admin_note, processed_at::text, created_at::text FROM withdrawal_requests WHERE seller_id=$1 ORDER BY created_at DESC", userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"withdrawals": []gin.H{}}})
		return
	}
	defer rows.Close()

	var items []gin.H
	for rows.Next() {
		var id, bankName, bankAccount, bankOwner, status, createdAt string
		var adminNote, processedAt sql.NullString
		var amount float64
		rows.Scan(&id, &amount, &bankName, &bankAccount, &bankOwner, &status, &adminNote, &processedAt, &createdAt)
		items = append(items, gin.H{
			"id": id, "amount": amount, "bankName": bankName, "bankAccount": bankAccount,
			"bankOwner": bankOwner, "status": status, "adminNote": adminNote.String,
			"processedAt": processedAt.String, "createdAt": createdAt,
		})
	}
	if items == nil {
		items = []gin.H{}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"withdrawals": items}})
}

// Admin: Xem tat ca escrow
func (h *EscrowHandler) ListEscrows(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"escrows": []gin.H{}, "totalHolding": 0}})
		return
	}

	rows, err := h.db.Query("SELECT id, COALESCE(transaction_id::text,''), COALESCE(buyer_id::text,''), COALESCE(buyer_name,''), COALESCE(seller_id::text,''), COALESCE(seller_name,''), COALESCE(amount,0), COALESCE(fee_amount,0), COALESCE(seller_amount,0), COALESCE(status,''), COALESCE(hold_until::text,''), COALESCE(released_at::text,''), COALESCE(created_at::text,'') FROM escrow_transactions ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"escrows": []gin.H{}, "totalHolding": 0}})
		return
	}
	defer rows.Close()

	var totalHolding float64
	var escrows []gin.H
	for rows.Next() {
		var id, txID, buyerID, buyerName, sellerID, sellerName, status, holdUntil, releasedAt, createdAt string
		var amount, feeAmount, sellerAmount float64
		rows.Scan(&id, &txID, &buyerID, &buyerName, &sellerID, &sellerName, &amount, &feeAmount, &sellerAmount, &status, &holdUntil, &releasedAt, &createdAt)
		if status == "holding" {
			totalHolding += amount
		}
		escrows = append(escrows, gin.H{
			"id": id, "transactionId": txID, "buyerId": buyerID, "buyerName": buyerName,
			"sellerId": sellerID, "sellerName": sellerName, "amount": amount,
			"feeAmount": feeAmount, "sellerAmount": sellerAmount, "status": status,
			"holdUntil": holdUntil, "releasedAt": releasedAt, "createdAt": createdAt,
		})
	}
	if escrows == nil {
		escrows = []gin.H{}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"escrows": escrows, "totalHolding": totalHolding}})
}

// Admin: Xem tat ca yeu cau rut tien
func (h *EscrowHandler) ListWithdrawals(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"withdrawals": []gin.H{}}})
		return
	}

	rows, err := h.db.Query("SELECT id, seller_id::text, seller_name, amount, bank_name, bank_account, bank_owner, status, admin_note, processed_at::text, created_at::text FROM withdrawal_requests ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"withdrawals": []gin.H{}}})
		return
	}
	defer rows.Close()

	var items []gin.H
	for rows.Next() {
		var id, sellerID, sellerName, bankName, bankAccount, bankOwner, status, createdAt string
		var adminNote, processedAt sql.NullString
		var amount float64
		rows.Scan(&id, &sellerID, &sellerName, &amount, &bankName, &bankAccount, &bankOwner, &status, &adminNote, &processedAt, &createdAt)
		items = append(items, gin.H{
			"id": id, "sellerId": sellerID, "sellerName": sellerName, "amount": amount,
			"bankName": bankName, "bankAccount": bankAccount, "bankOwner": bankOwner,
			"status": status, "adminNote": adminNote.String, "processedAt": processedAt.String, "createdAt": createdAt,
		})
	}
	if items == nil {
		items = []gin.H{}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"withdrawals": items}})
}

// Admin: Duyet rut tien
func (h *EscrowHandler) ApproveWithdrawal(c *gin.Context) {
	id := c.Param("id")

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"status": "completed"}})
		return
	}

	var sellerID string
	var amount float64
	err := h.db.QueryRow("SELECT seller_id, amount FROM withdrawal_requests WHERE id=$1 AND status='pending'", id).Scan(&sellerID, &amount)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Not found or already processed"})
		return
	}

	h.db.Exec("UPDATE withdrawal_requests SET status='completed', processed_at=NOW() WHERE id=$1", id)
	h.db.Exec("UPDATE seller_wallet SET total_withdrawn = total_withdrawn + $1, updated_at = NOW() WHERE seller_id=$2", amount, sellerID)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": id, "status": "completed"}})
}

// Admin: Tu choi rut tien
func (h *EscrowHandler) RejectWithdrawal(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Reason string `json:"reason"`
	}
	c.ShouldBindJSON(&req)

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"status": "rejected"}})
		return
	}

	// Hoan tien lai cho seller
	var sellerID string
	var amount float64
	err := h.db.QueryRow("SELECT seller_id, amount FROM withdrawal_requests WHERE id=$1 AND status='pending'", id).Scan(&sellerID, &amount)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Not found or already processed"})
		return
	}

	h.db.Exec("UPDATE withdrawal_requests SET status='rejected', admin_note=$1, processed_at=NOW() WHERE id=$2", req.Reason, id)
	h.db.Exec("UPDATE seller_wallet SET balance = balance + $1, updated_at = NOW() WHERE seller_id=$2", amount, sellerID)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": id, "status": "rejected"}})
}
