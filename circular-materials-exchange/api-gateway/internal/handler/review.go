package handler

import (
	authpb "api-gateway/internal/pb/auth"
	reviewpb "api-gateway/internal/pb/review"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	review reviewpb.ReviewServiceClient
	auth   authpb.AuthServiceClient
}

func NewReviewHandler(review reviewpb.ReviewServiceClient, auth authpb.AuthServiceClient) *ReviewHandler {
	return &ReviewHandler{review: review, auth: auth}
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
	reviewerName, _ := c.Get("name")
	if stringValue(reviewerName) == "" {
		reviewerName, _ = c.Get("email")
	}
	ctx, cancel := rpcContext(c)
	defer cancel()
	revieweeName := ""
	if response, err := h.auth.GetUser(ctx, &authpb.GetUserRequest{Id: req.RevieweeID}); err == nil {
		revieweeName = response.GetUser().GetName()
	}
	review, err := h.review.CreateReview(ctx, &reviewpb.CreateReviewRequest{
		TransactionId: req.TransactionID, ReviewerId: stringValue(userID),
		ReviewerName: stringValue(reviewerName), RevieweeId: req.RevieweeID,
		RevieweeName: revieweeName, Rating: req.Rating, Comment: req.Comment,
	})
	if err != nil {
		writeRPCError(c, err, "Loi tao danh gia")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": reviewJSON(review)})
}

func (h *ReviewHandler) ListReviews(c *gin.Context) {
	userID, _ := c.Get("user_id")
	page, pageSize := pagination(c)
	revieweeID := c.Query("reviewee_id")
	request := &reviewpb.ListReviewsRequest{RevieweeId: revieweeID, Page: page, PageSize: pageSize}
	if revieweeID == "" {
		request.UserId = stringValue(userID)
	}
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.review.ListReviews(ctx, request)
	if err != nil {
		writeRPCError(c, err, "Loi lay danh gia")
		return
	}
	items := make([]gin.H, 0, len(response.GetReviews()))
	for _, review := range response.GetReviews() {
		items = append(items, reviewJSON(review))
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"reviews": items, "total": response.GetTotal(), "averageRating": response.GetAverageRating(),
	}})
}

func reviewJSON(review *reviewpb.Review) gin.H {
	return gin.H{
		"id": review.GetId(), "transactionId": review.GetTransactionId(),
		"reviewerId": review.GetReviewerId(), "reviewerName": review.GetReviewerName(),
		"revieweeId": review.GetRevieweeId(), "revieweeName": review.GetRevieweeName(),
		"rating": review.GetRating(), "comment": review.GetComment(), "createdAt": review.GetCreatedAt(),
	}
}
