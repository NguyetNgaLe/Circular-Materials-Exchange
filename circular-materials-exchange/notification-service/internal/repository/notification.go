package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type Notification struct {
	ID          string
	UserID      string
	Title       string
	Message     string
	Type        string
	ReferenceID string
	Read        bool
	CreatedAt   time.Time
}

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) Create(notification *Notification) error {
	_, err := r.db.Exec(`INSERT INTO notifications
		(id, user_id, title, message, type, reference_id, read, created_at)
		VALUES ($1,$2,$3,$4,$5,NULLIF($6,'')::uuid,$7,$8)
		ON CONFLICT DO NOTHING`,
		notification.ID, notification.UserID, notification.Title, notification.Message,
		notification.Type, notification.ReferenceID, notification.Read, notification.CreatedAt)
	return err
}

func (r *NotificationRepository) FindByID(id string) (*Notification, error) {
	var n Notification
	err := r.db.QueryRow(`SELECT id, user_id, title, message, type,
		COALESCE(reference_id::text,''), read, created_at
		FROM notifications WHERE id=$1`, id).
		Scan(&n.ID, &n.UserID, &n.Title, &n.Message, &n.Type, &n.ReferenceID, &n.Read, &n.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("notification not found: %w", err)
	}
	return &n, nil
}

func (r *NotificationRepository) ListByUser(userID string, page, pageSize int) ([]Notification, int64, error) {
	var total int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM notifications WHERE user_id=$1`, userID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	rows, err := r.db.Query(`SELECT id, user_id, title, message, type,
		COALESCE(reference_id::text,''), read, created_at
		FROM notifications WHERE user_id=$1
		ORDER BY created_at DESC LIMIT $2 OFFSET $3`, userID, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var notifications []Notification
	for rows.Next() {
		var n Notification
		if err := rows.Scan(&n.ID, &n.UserID, &n.Title, &n.Message, &n.Type, &n.ReferenceID, &n.Read, &n.CreatedAt); err != nil {
			return nil, 0, err
		}
		notifications = append(notifications, n)
	}
	return notifications, total, nil
}

func (r *NotificationRepository) MarkRead(id string) (*Notification, error) {
	_, err := r.db.Exec(`UPDATE notifications SET read=true WHERE id=$1`, id)
	if err != nil {
		return nil, err
	}
	return r.FindByID(id)
}

func (r *NotificationRepository) MarkAllRead(userID string) error {
	_, err := r.db.Exec(`UPDATE notifications SET read=true WHERE user_id=$1 AND read=false`, userID)
	return err
}

func (r *NotificationRepository) GetUnreadCount(userID string) (int64, error) {
	var count int64
	err := r.db.QueryRow(`SELECT COUNT(*) FROM notifications WHERE user_id=$1 AND read=false`, userID).Scan(&count)
	return count, err
}
