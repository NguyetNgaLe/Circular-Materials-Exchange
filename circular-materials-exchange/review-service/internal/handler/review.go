package handler

import (
	"context"
	"review-service/internal/repository"
	"review-service/internal/service"
	"review-service/pb"
	"time"
)

type ReviewHandler struct {
	pb.UnimplementedReviewServiceServer
	svc *service.ReviewService
}

func NewReviewHandler(svc *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{svc: svc}
}

func (h *ReviewHandler) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.Review, error) {
	review, err := h.svc.CreateReview(
		req.GetTransactionId(),
		req.GetReviewerId(),
		req.GetReviewerName(),
		req.GetRevieweeId(),
		req.GetRevieweeName(),
		req.GetRating(),
		req.GetComment(),
	)
	if err != nil {
		return nil, err
	}
	return reviewToProto(review), nil
}

func (h *ReviewHandler) GetReview(ctx context.Context, req *pb.GetReviewRequest) (*pb.Review, error) {
	review, err := h.svc.GetReview(req.GetId())
	if err != nil {
		return nil, err
	}
	return reviewToProto(review), nil
}

func (h *ReviewHandler) ListReviews(ctx context.Context, req *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error) {
	var reviews []repository.Review
	var total int64
	var err error
	if req.GetUserId() != "" && req.GetRevieweeId() == "" {
		reviews, total, err = h.svc.ListReviewsForUser(req.GetUserId(), int(req.GetPage()), int(req.GetPageSize()))
	} else {
		reviews, total, err = h.svc.ListReviews(req.GetRevieweeId(), int(req.GetPage()), int(req.GetPageSize()))
	}
	if err != nil {
		return nil, err
	}

	protoReviews := make([]*pb.Review, len(reviews))
	for i, r := range reviews {
		protoReviews[i] = reviewToProto(&r)
	}

	avg, _, _ := h.svc.GetUserRating(req.GetRevieweeId())

	return &pb.ListReviewsResponse{
		Reviews:       protoReviews,
		Total:         int32(total),
		AverageRating: avg,
	}, nil
}

func (h *ReviewHandler) GetUserRating(ctx context.Context, req *pb.GetUserRatingRequest) (*pb.UserRatingResponse, error) {
	avg, count, err := h.svc.GetUserRating(req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.UserRatingResponse{
		Average: avg,
		Count:   int32(count),
	}, nil
}

func reviewToProto(r *repository.Review) *pb.Review {
	return &pb.Review{
		Id:            r.ID,
		TransactionId: r.TransactionID,
		ReviewerId:    r.ReviewerID,
		ReviewerName:  r.ReviewerName,
		RevieweeId:    r.RevieweeID,
		RevieweeName:  r.RevieweeName,
		Rating:        r.Rating,
		Comment:       r.Comment,
		CreatedAt:     r.CreatedAt.Format(time.RFC3339),
	}
}
