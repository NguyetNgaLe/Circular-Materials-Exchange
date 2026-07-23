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

type OrderHandler struct {
	conn         *grpc.ClientConn
	db           *sql.DB
	companyDB    *sql.DB
	notifHandler *NotificationHandler
}

func NewOrderHandler(conn *grpc.ClientConn) *OrderHandler {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5433")
	dbUser := getEnv("DB_USER", "cme")
	dbPass := getEnv("DB_PASSWORD", "cme_secret_2024")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=order_db sslmode=disable", dbHost, dbPort, dbUser, dbPass)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		db = nil
	}

	companyDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=company_db sslmode=disable", dbHost, dbPort, dbUser, dbPass)
	companyDB, err := sql.Open("postgres", companyDSN)
	if err != nil {
		companyDB = nil
	}

	return &OrderHandler{conn: conn, db: db, companyDB: companyDB}
}

func (h *OrderHandler) SetNotificationHandler(nh *NotificationHandler) {
	h.notifHandler = nh
}

func (h *OrderHandler) CreateOffer(c *gin.Context) {
	var req struct {
		Type          string  `json:"type" binding:"required"`
		ListingID     string  `json:"listingId" binding:"required"`
		ListingTitle  string  `json:"listingTitle"`
		BuyerID       string  `json:"buyerId"`
		BuyerName     string  `json:"buyerName"`
		SellerID      string  `json:"sellerId" binding:"required"`
		SellerName    string  `json:"sellerName"`
		Quantity      float64 `json:"quantity" binding:"required"`
		Unit          string  `json:"unit"`
		ProposedPrice float64 `json:"proposedPrice" binding:"required"`
		Currency      string  `json:"currency"`
		Message       string  `json:"message"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")

	// Kiem tra doanh nghiep buyer da duoc duyet chua
	if h.companyDB != nil {
		var companyStatus string
		err := h.companyDB.QueryRow("SELECT status FROM companies WHERE owner_id=$1 LIMIT 1", userID).Scan(&companyStatus)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Ban can co ho do doanh nghiep de gui de nghi mua"})
			return
		}
		if companyStatus != "verified" {
			c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Doanh nghiep chua duoc admin duyet"})
			return
		}
	}

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": "of-demo"}})
		return
	}

	id := uuid.New().String()
	_, err := h.db.Exec(`INSERT INTO offers (id, type, listing_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, proposed_price, currency, message, status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,'pending')`,
		id, req.Type, req.ListingID, req.ListingTitle, userID, req.BuyerName, req.SellerID, req.SellerName, req.Quantity, req.Unit, req.ProposedPrice, req.Currency, req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Tao escrow (giu tien) khi buyer thanh toan
	totalAmount := req.Quantity * req.ProposedPrice
	feeRate := 0.02
	feeAmount := totalAmount * feeRate
	sellerAmount := totalAmount - feeAmount
	escrowID := uuid.New().String()
	holdUntil := time.Now().AddDate(0, 0, 3) // 3 ngay

	h.db.Exec(`INSERT INTO escrow_transactions (id, transaction_id, buyer_id, buyer_name, seller_id, seller_name, amount, fee_rate, fee_amount, seller_amount, status, hold_until) VALUES ($1,NULL,$2,$3,$4,$5,$6,$7,$8,$9,'holding',$10)`,
		escrowID, userID, req.BuyerName, req.SellerID, req.SellerName, totalAmount, feeRate, feeAmount, sellerAmount, holdUntil)

	// Tao thong bao cho seller
	if h.notifHandler != nil {
		h.notifHandler.CreateNotification(req.SellerID,
			"De nghi mua moi",
			fmt.Sprintf("%s muon mua %s %.0f %s voi gia %.0f VND/%s", req.BuyerName, req.ListingTitle, req.Quantity, req.Unit, req.ProposedPrice, req.Unit),
			"offer", id)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id": id, "type": req.Type, "listingId": req.ListingID, "listingTitle": req.ListingTitle,
			"buyerId": userID, "buyerName": req.BuyerName, "sellerId": req.SellerID, "sellerName": req.SellerName,
			"quantity": req.Quantity, "unit": req.Unit, "proposedPrice": req.ProposedPrice,
			"currency": req.Currency, "message": req.Message, "status": "pending",
		},
	})
}

func (h *OrderHandler) ListOffers(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"offers": []gin.H{}, "total": 0}})
		return
	}

	userID, _ := c.Get("user_id")
	userRole, _ := c.Get("role")
	role := c.Query("role")

	var rows *sql.Rows
	var err error

	if userRole == "admin" {
		rows, err = h.db.Query("SELECT id, type, listing_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, proposed_price, currency, message, status, created_at::text FROM offers ORDER BY created_at DESC")
	} else if role == "buyer" {
		rows, err = h.db.Query("SELECT id, type, listing_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, proposed_price, currency, message, status, created_at::text FROM offers WHERE buyer_id=$1 ORDER BY created_at DESC", userID)
	} else if role == "seller" {
		rows, err = h.db.Query("SELECT id, type, listing_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, proposed_price, currency, message, status, created_at::text FROM offers WHERE seller_id=$1 ORDER BY created_at DESC", userID)
	} else {
		rows, err = h.db.Query("SELECT id, type, listing_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, proposed_price, currency, message, status, created_at::text FROM offers WHERE buyer_id=$1 OR seller_id=$1 ORDER BY created_at DESC", userID)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"offers": []gin.H{}, "total": 0}})
		return
	}
	defer rows.Close()

	var offers []gin.H
	for rows.Next() {
		var id, typ, listingID, listingTitle, buyerID, buyerName, sellerID, sellerName, unit, currency, message, status, createdAt string
		var qty, price float64
		rows.Scan(&id, &typ, &listingID, &listingTitle, &buyerID, &buyerName, &sellerID, &sellerName, &qty, &unit, &price, &currency, &message, &status, &createdAt)
		offers = append(offers, gin.H{
			"id": id, "type": typ, "listingId": listingID, "listingTitle": listingTitle,
			"buyerId": buyerID, "buyerName": buyerName, "sellerId": sellerID, "sellerName": sellerName,
			"quantity": qty, "unit": unit, "proposedPrice": price, "currency": currency,
			"message": message, "status": status, "createdAt": createdAt,
		})
	}
	if offers == nil {
		offers = []gin.H{}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"offers": offers, "total": len(offers)}})
}

func (h *OrderHandler) AcceptOffer(c *gin.Context) {
	offerID := c.Param("id")

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": offerID, "status": "accepted"}})
		return
	}

	// Update offer status
	_, err := h.db.Exec("UPDATE offers SET status='accepted', updated_at=NOW() WHERE id=$1", offerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Get offer details
	var offer struct {
		ListingTitle, BuyerID, BuyerName, SellerID, SellerName, Unit string
		Quantity, ProposedPrice                                       float64
	}
	err = h.db.QueryRow("SELECT listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, proposed_price FROM offers WHERE id=$1", offerID).Scan(
		&offer.ListingTitle, &offer.BuyerID, &offer.BuyerName, &offer.SellerID, &offer.SellerName, &offer.Quantity, &offer.Unit, &offer.ProposedPrice)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": offerID, "status": "accepted"}})
		return
	}

	// Create transaction
	txID := uuid.New().String()
	_, err = h.db.Exec(`INSERT INTO transactions (id, offer_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, agreed_price, currency, payment_status, payment_method, settlement_note, status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,'VND','bypassed_demo','manual_offline','Thanh toan duoc thuc hien ngoai he thong trong pham vi prototype','confirmed')`,
		txID, offerID, offer.ListingTitle, offer.BuyerID, offer.BuyerName, offer.SellerID, offer.SellerName, offer.Quantity, offer.Unit, offer.ProposedPrice)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": offerID, "status": "accepted"}})
		return
	}

	// Create event
	evID := uuid.New().String()
	h.db.Exec(`INSERT INTO transaction_events (id, transaction_id, actor_id, actor_name, from_status, to_status, note) VALUES ($1,$2,$3,$4,'offer.accepted','transaction.confirmed','Giao dich duoc tao tu dong khi seller chap nhan offer')`,
		evID, txID, offer.SellerID, offer.SellerName)

	// Create escrow (tu dong giu tien)
	totalAmount := offer.Quantity * offer.ProposedPrice
	feeRate := 0.02
	feeAmount := totalAmount * feeRate
	sellerAmount := totalAmount - feeAmount
	escrowID := uuid.New().String()
	holdUntil := time.Now().AddDate(0, 0, 3) // 3 ngay

	h.db.Exec(`INSERT INTO escrow_transactions (id, transaction_id, buyer_id, buyer_name, seller_id, seller_name, amount, fee_rate, fee_amount, seller_amount, status, hold_until) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,'holding',$11)`,
		escrowID, txID, offer.BuyerID, offer.BuyerName, offer.SellerID, offer.SellerName, totalAmount, feeRate, feeAmount, sellerAmount, holdUntil)

	// Tao thong bao cho buyer
	if h.notifHandler != nil {
		h.notifHandler.CreateNotification(offer.BuyerID,
			"De nghi da duoc chap nhan",
			fmt.Sprintf("%s da chap nhan de nghi mua %s. Giao dich da duoc tao.", offer.SellerName, offer.ListingTitle),
			"offer_accepted", offerID)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id": offerID, "status": "accepted",
			"transaction": gin.H{"id": txID, "status": "confirmed"},
			"escrow": gin.H{"id": escrowID, "amount": totalAmount, "feeAmount": feeAmount, "sellerAmount": sellerAmount, "holdUntil": holdUntil},
		},
	})
}

func (h *OrderHandler) RejectOffer(c *gin.Context) {
	offerID := c.Param("id")
	if h.db != nil {
		h.db.Exec("UPDATE offers SET status='rejected', updated_at=NOW() WHERE id=$1", offerID)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": offerID, "status": "rejected"}})
}

func (h *OrderHandler) ListTransactions(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": []gin.H{}, "total": 0}})
		return
	}

	userID, _ := c.Get("user_id")
	userRole, _ := c.Get("role")

	var rows *sql.Rows
	var err error

	if userRole == "admin" {
		rows, err = h.db.Query("SELECT id, offer_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, agreed_price, currency, payment_status, payment_method, settlement_note, status, created_at::text FROM transactions ORDER BY created_at DESC")
	} else {
		rows, err = h.db.Query("SELECT id, offer_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, agreed_price, currency, payment_status, payment_method, settlement_note, status, created_at::text FROM transactions WHERE buyer_id=$1 OR seller_id=$1 ORDER BY created_at DESC", userID)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": []gin.H{}, "total": 0}})
		return
	}
	defer rows.Close()

	var txs []gin.H
	for rows.Next() {
		var id, offerID, listingTitle, buyerID, buyerName, sellerID, sellerName, unit, currency, payStatus, payMethod, note, status, createdAt string
		var qty, price float64
		rows.Scan(&id, &offerID, &listingTitle, &buyerID, &buyerName, &sellerID, &sellerName, &qty, &unit, &price, &currency, &payStatus, &payMethod, &note, &status, &createdAt)

		// Get events
		evRows, _ := h.db.Query("SELECT id, actor_id, actor_name, from_status, to_status, note, created_at::text FROM transaction_events WHERE transaction_id=$1 ORDER BY created_at", id)
		var events []gin.H
		if evRows != nil {
			for evRows.Next() {
				var evID, actorID, actorName, fromSt, toSt, evNote, evTime string
				evRows.Scan(&evID, &actorID, &actorName, &fromSt, &toSt, &evNote, &evTime)
				events = append(events, gin.H{"id": evID, "actorId": actorID, "actorName": actorName, "fromStatus": fromSt, "toStatus": toSt, "note": evNote, "createdAt": evTime})
			}
			evRows.Close()
		}
		if events == nil {
			events = []gin.H{}
		}

		txs = append(txs, gin.H{
			"id": id, "offerId": offerID, "listingTitle": listingTitle,
			"buyerId": buyerID, "buyerName": buyerName, "sellerId": sellerID, "sellerName": sellerName,
			"quantity": qty, "unit": unit, "agreedPrice": price, "currency": currency,
			"paymentStatus": payStatus, "paymentMethod": payMethod, "settlementNote": note,
			"status": status, "createdAt": createdAt, "events": events,
		})
	}
	if txs == nil {
		txs = []gin.H{}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": txs, "total": len(txs)}})
}

func (h *OrderHandler) GetTransaction(c *gin.Context) {
	id := c.Param("id")
	if h.db == nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Not found"})
		return
	}

	var tx struct {
		ID, OfferID, ListingTitle, BuyerID, BuyerName, SellerID, SellerName, Unit, Currency, PayStatus, PayMethod, Note, Status, CreatedAt string
		Qty, Price                                                                                                                           float64
	}
	err := h.db.QueryRow("SELECT id, offer_id, listing_title, buyer_id, buyer_name, seller_id, seller_name, quantity, unit, agreed_price, currency, payment_status, payment_method, settlement_note, status, created_at::text FROM transactions WHERE id=$1", id).Scan(
		&tx.ID, &tx.OfferID, &tx.ListingTitle, &tx.BuyerID, &tx.BuyerName, &tx.SellerID, &tx.SellerName, &tx.Qty, &tx.Unit, &tx.Price, &tx.Currency, &tx.PayStatus, &tx.PayMethod, &tx.Note, &tx.Status, &tx.CreatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Not found"})
		return
	}

	// Get events
	evRows, _ := h.db.Query("SELECT id, actor_id, actor_name, from_status, to_status, note, created_at::text FROM transaction_events WHERE transaction_id=$1 ORDER BY created_at", id)
	var events []gin.H
	if evRows != nil {
		for evRows.Next() {
			var evID, actorID, actorName, fromSt, toSt, evNote, evTime string
			evRows.Scan(&evID, &actorID, &actorName, &fromSt, &toSt, &evNote, &evTime)
			events = append(events, gin.H{"id": evID, "actorId": actorID, "actorName": actorName, "fromStatus": fromSt, "toStatus": toSt, "note": evNote, "createdAt": evTime})
		}
		evRows.Close()
	}
	if events == nil {
		events = []gin.H{}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"id": tx.ID, "offerId": tx.OfferID, "listingTitle": tx.ListingTitle,
		"buyerId": tx.BuyerID, "buyerName": tx.BuyerName, "sellerId": tx.SellerID, "sellerName": tx.SellerName,
		"quantity": tx.Qty, "unit": tx.Unit, "agreedPrice": tx.Price, "currency": tx.Currency,
		"paymentStatus": tx.PayStatus, "paymentMethod": tx.PayMethod, "settlementNote": tx.Note,
		"status": tx.Status, "createdAt": tx.CreatedAt, "events": events,
	}})
}

func (h *OrderHandler) UpdateTransactionStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Status string `json:"status" binding:"required"`
		Note   string `json:"note"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": id, "status": req.Status}})
		return
	}

	// Get current status
	var oldStatus string
	h.db.QueryRow("SELECT status FROM transactions WHERE id=$1", id).Scan(&oldStatus)

	// Update
	_, err := h.db.Exec("UPDATE transactions SET status=$1, updated_at=NOW() WHERE id=$2", req.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	// Create event
	userID, _ := c.Get("user_id")
	evID := uuid.New().String()
	h.db.Exec(`INSERT INTO transaction_events (id, transaction_id, actor_id, actor_name, from_status, to_status, note) VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		evID, id, userID, userID, "transaction."+oldStatus, "transaction."+req.Status, req.Note)

	// Neu transaction completed -> tu dong giai ngan escrow
	if req.Status == "completed" {
		var escrowID, sellerID, sellerName, buyerID, buyerName string
		var sellerAmount, feeAmount, totalAmount float64
		err := h.db.QueryRow("SELECT id, seller_id, seller_name, buyer_id, buyer_name, amount, seller_amount, fee_amount FROM escrow_transactions WHERE transaction_id=$1 AND status='holding'", id).Scan(
			&escrowID, &sellerID, &sellerName, &buyerID, &buyerName, &totalAmount, &sellerAmount, &feeAmount)
		if err == nil {
			// Giai ngan escrow
			h.db.Exec("UPDATE escrow_transactions SET status='released', released_at=NOW() WHERE id=$1", escrowID)

			// Them vao platform_fees (de FinanceHandler doc duoc)
			feeID := uuid.New().String()
			h.db.Exec(`INSERT INTO platform_fees (id, transaction_id, seller_id, buyer_id, transaction_amount, fee_rate, fee_amount, fee_type, status, collected_at) VALUES ($1,$2,$3,$4,$5,0.02,$6,'transaction','collected',NOW())`,
				feeID, id, sellerID, buyerID, totalAmount, feeAmount)

			// Cap nhat seller_wallet
			h.db.Exec(`INSERT INTO seller_wallet (seller_id, seller_name, balance, total_earned, total_fees_paid) 
				VALUES ($1, $2, $3, $3, $4) 
				ON CONFLICT (seller_id) DO UPDATE SET 
				balance = seller_wallet.balance + $3, 
				total_earned = seller_wallet.total_earned + $3,
				total_fees_paid = seller_wallet.total_fees_paid + $4,
				updated_at = NOW()`,
				sellerID, sellerName, sellerAmount, feeAmount)

			// Cap nhat platform_wallet (san nhan phi)
			h.db.Exec("UPDATE platform_wallet SET balance = balance + $1, total_income = total_income + $1, updated_at = NOW()", feeAmount)

			// Lay so du moi
			var newBalance float64
			h.db.QueryRow("SELECT balance FROM seller_wallet WHERE seller_id=$1", sellerID).Scan(&newBalance)

			// Tao lich su vi seller
			walletTxID := uuid.New().String()
			h.db.Exec(`INSERT INTO seller_wallet_transactions (id, seller_id, type, amount, balance_after, reference_type, reference_id, description) VALUES ($1,$2,'credit',$3,$4,'escrow_release',$5,$6)`,
				walletTxID, sellerID, sellerAmount, newBalance, escrowID, fmt.Sprintf("Nhan tien tu giao dich %s (tru phi %.0f%%)", id[:8], feeAmount/sellerAmount*100))
		}
	}

	// Tao thong bao khi trang thai giao dich thay doi
	if h.notifHandler != nil {
		var buyerID, sellerID, listingTitle string
		h.db.QueryRow("SELECT buyer_id, seller_id, listing_title FROM transactions WHERE id=$1", id).Scan(&buyerID, &sellerID, &listingTitle)

		if req.Status == "in_progress" {
			h.notifHandler.CreateNotification(buyerID,
				"Seller da giao hang",
				fmt.Sprintf("San pham %s da duoc giao. Vui long xac nhan nhan hang.", listingTitle),
				"transaction", id)
		} else if req.Status == "completed" {
			h.notifHandler.CreateNotification(sellerID,
				"Giao dich hoan tat",
				fmt.Sprintf("Giao dich %s da hoan tat. Tien da chuyen vao vi cua ban.", listingTitle),
				"transaction", id)
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": id, "status": req.Status}})
}
