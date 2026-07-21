package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type Review struct {
	ID            string
	TransactionID string
	ReviewerID    string
	ReviewerName  string
	RevieweeID    string
	RevieweeName  string
	Rating        int32
	Comment       string
	CreatedAt     time.Time
}

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) Create(review *Review) error {
	_, err := r.db.Exec(`INSERT INTO reviews (id, transaction_id, reviewer_id, reviewer_name, reviewee_id, reviewee_name, rating, comment, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		review.ID, review.TransactionID, review.ReviewerID, review.ReviewerName, review.RevieweeID, review.RevieweeName, review.Rating, review.Comment, review.CreatedAt)
	return err
}

func (r *ReviewRepository) FindByID(id string) (*Review, error) {
	var review Review
	err := r.db.QueryRow(`SELECT id, transaction_id, reviewer_id, reviewer_name, reviewee_id, reviewee_name, rating, comment, created_at FROM reviews WHERE id=$1`, id).
		Scan(&review.ID, &review.TransactionID, &review.ReviewerID, &review.ReviewerName, &review.RevieweeID, &review.RevieweeName, &review.Rating, &review.Comment, &review.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("review not found: %w", err)
	}
	return &review, nil
}

func (r *ReviewRepository) ListByReviewee(revieweeID string, page, pageSize int) ([]Review, int64, error) {
	var total int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM reviews WHERE reviewee_id=$1`, revieweeID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	rows, err := r.db.Query(`SELECT id, transaction_id, reviewer_id, reviewer_name, reviewee_id, reviewee_name, rating, comment, created_at FROM reviews WHERE reviewee_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, revieweeID, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		if err := rows.Scan(&review.ID, &review.TransactionID, &review.ReviewerID, &review.ReviewerName, &review.RevieweeID, &review.RevieweeName, &review.Rating, &review.Comment, &review.CreatedAt); err != nil {
			return nil, 0, err
		}
		reviews = append(reviews, review)
	}
	return reviews, total, nil
}

func (r *ReviewRepository) GetAverageRating(userID string) (float64, int64, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM reviews WHERE reviewee_id=$1`, userID).Scan(&count)
	if err != nil {
		return 0, 0, err
	}
	if count == 0 {
		return 0, 0, nil
	}

	var avg float64
	err = r.db.QueryRow(`SELECT COALESCE(AVG(rating), 0) FROM reviews WHERE reviewee_id=$1`, userID).Scan(&avg)
	if err != nil {
		return 0, 0, err
	}
	return avg, count, nil
}
