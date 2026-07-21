package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ReviewHandler struct {
	conn *grpc.ClientConn
}

func NewReviewHandler(conn *grpc.ClientConn) *ReviewHandler {
	return &ReviewHandler{conn: conn}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var req struct {
		TransactionID string `json:"transaction_id" binding:"required"`
		RevieweeID    string `json:"reviewee_id" binding:"required"`
		RevieweeName  string `json:"reviewee_name"`
		Rating        int32  `json:"rating" binding:"required"`
		Comment       string `json:"comment"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id":             "rv-new",
			"transaction_id": req.TransactionID,
			"reviewer_id":    userID,
			"reviewee_id":    req.RevieweeID,
			"reviewee_name":  req.RevieweeName,
			"rating":         req.Rating,
			"comment":        req.Comment,
		},
	})
}

func (h *ReviewHandler) ListReviews(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"reviews":        []gin.H{},
			"total":          0,
			"average_rating": 0,
		},
	})
}
