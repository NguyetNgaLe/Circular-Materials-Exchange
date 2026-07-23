package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	_ "github.com/lib/pq"
)

type FinanceHandler struct {
	conn *grpc.ClientConn
	db   *sql.DB
}

func NewFinanceHandler(conn *grpc.ClientConn) *FinanceHandler {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5433")
	dbName := "order_db"
	dbUser := getEnv("DB_USER", "cme")
	dbPass := getEnv("DB_PASSWORD", "cme_secret_2024")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return &FinanceHandler{conn: conn, db: nil}
	}
	return &FinanceHandler{conn: conn, db: db}
}

// GET /api/admin/finance/overview
func (h *FinanceHandler) GetOverview(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
			"totalIncome": 0, "monthIncome": 0, "totalTransactions": 0, "walletBalance": 0,
		}})
		return
	}

	// Tổng doanh thu
	var totalIncome float64
	h.db.QueryRow("SELECT COALESCE(SUM(fee_amount), 0) FROM platform_fees WHERE status='collected'").Scan(&totalIncome)

	// Doanh thu tháng này
	var monthIncome float64
	h.db.QueryRow("SELECT COALESCE(SUM(fee_amount), 0) FROM platform_fees WHERE status='collected' AND created_at >= date_trunc('month', CURRENT_DATE)").Scan(&monthIncome)

	// Số giao dịch
	var totalTx int
	h.db.QueryRow("SELECT COUNT(*) FROM platform_fees").Scan(&totalTx)

	// Số dư ví
	var walletBalance float64
	h.db.QueryRow("SELECT COALESCE(balance, 0) FROM platform_wallet LIMIT 1").Scan(&walletBalance)

	// Doanh thu theo tháng (6 tháng gần nhất)
	rows, _ := h.db.Query(`
		SELECT to_char(created_at, 'YYYY-MM') as month, SUM(fee_amount) as total 
		FROM platform_fees WHERE status='collected' 
		GROUP BY to_char(created_at, 'YYYY-MM') 
		ORDER BY month DESC LIMIT 6
	`)
	var monthlyData []gin.H
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var month string
			var total float64
			rows.Scan(&month, &total)
			monthlyData = append(monthlyData, gin.H{"month": month, "total": total})
		}
	}
	if monthlyData == nil {
		monthlyData = []gin.H{}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"totalIncome":       totalIncome,
			"monthIncome":       monthIncome,
			"totalTransactions": totalTx,
			"walletBalance":     walletBalance,
			"monthlyData":       monthlyData,
		},
	})
}

// GET /api/admin/finance/fees
func (h *FinanceHandler) ListFees(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"fees": []gin.H{}, "total": 0}})
		return
	}

	rows, err := h.db.Query(`
		SELECT f.id, f.transaction_id, f.seller_id, f.buyer_id, 
		       f.transaction_amount, f.fee_rate, f.fee_amount, 
		       f.fee_type, f.status, f.collected_at::text, f.created_at::text,
		       t.listing_title
		FROM platform_fees f
		LEFT JOIN transactions t ON f.transaction_id = t.id
		ORDER BY f.created_at DESC
		LIMIT 100
	`)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"fees": []gin.H{}, "total": 0}})
		return
	}
	defer rows.Close()

	var fees []gin.H
	for rows.Next() {
		var id, txID, sellerID, buyerID, feeType, status, createdAt string
		var collectedAt, listingTitle sql.NullString
		var txAmount, feeRate, feeAmount float64
		rows.Scan(&id, &txID, &sellerID, &buyerID, &txAmount, &feeRate, &feeAmount, &feeType, &status, &collectedAt, &createdAt, &listingTitle)
		fees = append(fees, gin.H{
			"id": id, "transactionId": txID, "sellerId": sellerID, "buyerId": buyerID,
			"transactionAmount": txAmount, "feeRate": feeRate, "feeAmount": feeAmount,
			"feeType": feeType, "status": status, "collectedAt": collectedAt.String,
			"createdAt": createdAt, "listingTitle": listingTitle.String,
		})
	}
	if fees == nil {
		fees = []gin.H{}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"fees": fees, "total": len(fees)}})
}

// GET /api/admin/finance/wallet
func (h *FinanceHandler) GetWallet(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
			"balance": 0, "totalIncome": 0, "totalExpense": 0,
		}})
		return
	}

	var balance, totalIncome, totalExpense float64
	err := h.db.QueryRow("SELECT COALESCE(balance,0), COALESCE(total_income,0), COALESCE(total_expense,0) FROM platform_wallet LIMIT 1").Scan(&balance, &totalIncome, &totalExpense)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
			"balance": 0, "totalIncome": 0, "totalExpense": 0,
		}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"balance":     balance,
			"totalIncome": totalIncome,
			"totalExpense": totalExpense,
		},
	})
}

// GET /api/admin/finance/wallet-transactions
func (h *FinanceHandler) ListWalletTransactions(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": []gin.H{}, "total": 0}})
		return
	}

	rows, err := h.db.Query(`
		SELECT id, type, amount, reference_type, reference_id::text, 
		       description, balance_after, created_at::text
		FROM wallet_transactions
		ORDER BY created_at DESC
		LIMIT 50
	`)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": []gin.H{}, "total": 0}})
		return
	}
	defer rows.Close()

	var txs []gin.H
	for rows.Next() {
		var id, txType, refType, refID, desc, createdAt string
		var amount, balanceAfter float64
		rows.Scan(&id, &txType, &amount, &refType, &refID, &desc, &balanceAfter, &createdAt)
		txs = append(txs, gin.H{
			"id": id, "type": txType, "amount": amount, "referenceType": refType,
			"referenceId": refID, "description": desc, "balanceAfter": balanceAfter, "createdAt": createdAt,
		})
	}
	if txs == nil {
		txs = []gin.H{}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": txs, "total": len(txs)}})
}

// POST /api/admin/finance/collect-fee (tính phí cho transaction)
func (h *FinanceHandler) CollectFee(c *gin.Context) {
	var req struct {
		TransactionID string  `json:"transactionId" binding:"required"`
		Amount        float64 `json:"amount" binding:"required"`
		SellerID      string  `json:"sellerId" binding:"required"`
		BuyerID       string  `json:"buyerId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"feeId": "demo-fee"}})
		return
	}

	feeRate := 0.02
	feeAmount := req.Amount * feeRate

	// Tạo bản ghi phí
	var feeID string
	err := h.db.QueryRow(`
		INSERT INTO platform_fees (transaction_id, seller_id, buyer_id, transaction_amount, fee_rate, fee_amount, status, collected_at)
		VALUES ($1, $2, $3, $4, $5, $6, 'collected', NOW())
		RETURNING id
	`, req.TransactionID, req.SellerID, req.BuyerID, req.Amount, feeRate, feeAmount).Scan(&feeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Cập nhật ví sàn
	h.db.Exec("UPDATE platform_wallet SET balance = balance + $1, total_income = total_income + $1, updated_at = NOW()", feeAmount)

	// Tạo lịch sử ví
	var balance float64
	h.db.QueryRow("SELECT balance FROM platform_wallet LIMIT 1").Scan(&balance)
	h.db.Exec(`
		INSERT INTO wallet_transactions (wallet_id, type, amount, reference_type, reference_id, description, balance_after)
		VALUES ('00000000-0000-0000-0000-000000000001', 'income', $1, 'platform_fee', $2, $3, $4)
	`, feeAmount, feeID, fmt.Sprintf("Phí giao dịch %s", req.TransactionID), balance)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"feeId":     feeID,
			"feeAmount": feeAmount,
			"feeRate":   feeRate,
			"balance":   balance,
		},
	})
}
