package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	_ "github.com/lib/pq"
)

type ReviewHandler struct {
	conn   *grpc.ClientConn
	db     *sql.DB
	authDB *sql.DB
}

func NewReviewHandler(conn *grpc.ClientConn) *ReviewHandler {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5433")
	dbUser := getEnv("DB_USER", "cme")
	dbPass := getEnv("DB_PASSWORD", "cme_secret_2024")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=review_db sslmode=disable", dbHost, dbPort, dbUser, dbPass)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		db = nil
	}

	authDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=auth_db sslmode=disable", dbHost, dbPort, dbUser, dbPass)
	authDB, err := sql.Open("postgres", authDSN)
	if err != nil {
		authDB = nil
	}

	return &ReviewHandler{conn: conn, db: db, authDB: authDB}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var req struct {
		TransactionID string `json:"transactionId" binding:"required"`
		RevieweeID    string `json:"revieweeId" binding:"required"`
		Rating        int32  `json:"rating" binding:"required"`
		Comment       string `json:"comment"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")
	userName, _ := c.Get("email")

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": "rv-demo"}})
		return
	}

	// Lay ten reviewee tu auth_db
	revieweeName := ""
	if h.authDB != nil {
		var rName sql.NullString
		h.authDB.QueryRow("SELECT name FROM users WHERE id=$1", req.RevieweeID).Scan(&rName)
		if rName.Valid {
			revieweeName = rName.String
		}
	}

	id := uuid.New().String()
	_, err := h.db.Exec(`INSERT INTO reviews (id, transaction_id, reviewer_id, reviewer_name, reviewee_id, reviewee_name, rating, comment) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		id, req.TransactionID, userID, userName, req.RevieweeID, revieweeName, req.Rating, req.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id": id, "transactionId": req.TransactionID,
			"reviewerId": userID, "reviewerName": userName,
			"revieweeId": req.RevieweeID, "revieweeName": revieweeName,
			"rating": req.Rating, "comment": req.Comment,
		},
	})
}

func (h *ReviewHandler) ListReviews(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"reviews": []gin.H{}, "total": 0}})
		return
	}

	userID, _ := c.Get("user_id")
	revieweeID := c.Query("reviewee_id")

	var rows *sql.Rows
	var err error

	if revieweeID != "" {
		rows, err = h.db.Query("SELECT id, transaction_id::text, reviewer_id::text, reviewer_name, reviewee_id::text, reviewee_name, rating, comment, created_at::text FROM reviews WHERE reviewee_id=$1 ORDER BY created_at DESC", revieweeID)
	} else {
		rows, err = h.db.Query("SELECT id, transaction_id::text, reviewer_id::text, reviewer_name, reviewee_id::text, reviewee_name, rating, comment, created_at::text FROM reviews WHERE reviewer_id=$1 OR reviewee_id=$1 ORDER BY created_at DESC", userID)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"reviews": []gin.H{}, "total": 0}})
		return
	}
	defer rows.Close()

	var reviews []gin.H
	var totalRating float64
	for rows.Next() {
		var id, txID, reviewerID, reviewerName, revieweeIDStr, revieweeName, createdAt string
		var comment sql.NullString
		var rating int
		rows.Scan(&id, &txID, &reviewerID, &reviewerName, &revieweeIDStr, &revieweeName, &rating, &comment, &createdAt)
		totalRating += float64(rating)
		reviews = append(reviews, gin.H{
			"id": id, "transactionId": txID,
			"reviewerId": reviewerID, "reviewerName": reviewerName,
			"revieweeId": revieweeIDStr, "revieweeName": revieweeName,
			"rating": rating, "comment": comment.String, "createdAt": createdAt,
		})
	}
	if reviews == nil {
		reviews = []gin.H{}
	}

	avgRating := 0.0
	if len(reviews) > 0 {
		avgRating = totalRating / float64(len(reviews))
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"reviews": reviews, "total": len(reviews), "averageRating": avgRating,
	}})
}
