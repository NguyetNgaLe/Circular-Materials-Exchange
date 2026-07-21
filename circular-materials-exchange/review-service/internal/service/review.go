package service

import (
	"review-service/internal/repository"
	"time"

	"github.com/google/uuid"
)

type ReviewService struct {
	repo *repository.ReviewRepository
}

func NewReviewService(repo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) CreateReview(transactionID, reviewerID, reviewerName, revieweeID, revieweeName string, rating int32, comment string) (*repository.Review, error) {
	review := &repository.Review{
		ID:            uuid.New().String(),
		TransactionID: transactionID,
		ReviewerID:    reviewerID,
		ReviewerName:  reviewerName,
		RevieweeID:    revieweeID,
		RevieweeName:  revieweeName,
		Rating:        rating,
		Comment:       comment,
		CreatedAt:     time.Now(),
	}

	if err := s.repo.Create(review); err != nil {
		return nil, err
	}

	return review, nil
}

func (s *ReviewService) GetReview(id string) (*repository.Review, error) {
	return s.repo.FindByID(id)
}

func (s *ReviewService) ListReviews(revieweeID string, page, pageSize int) ([]repository.Review, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return s.repo.ListByReviewee(revieweeID, page, pageSize)
}

func (s *ReviewService) GetUserRating(userID string) (float64, int64, error) {
	return s.repo.GetAverageRating(userID)
}
