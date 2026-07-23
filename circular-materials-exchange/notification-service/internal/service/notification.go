package service

import (
	"notification-service/internal/repository"
	"time"

	"github.com/google/uuid"
)

type NotificationService struct {
	repo *repository.NotificationRepository
}

func NewNotificationService(repo *repository.NotificationRepository) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) CreateNotification(userID, title, message, notifType, referenceID string) (*repository.Notification, error) {
	notification := &repository.Notification{
		ID:          uuid.New().String(),
		UserID:      userID,
		Title:       title,
		Message:     message,
		Type:        notifType,
		ReferenceID: referenceID,
		Read:        false,
		CreatedAt:   time.Now(),
	}

	if err := s.repo.Create(notification); err != nil {
		return nil, err
	}

	return notification, nil
}

func (s *NotificationService) ListNotifications(userID string, page, pageSize int) ([]repository.Notification, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return s.repo.ListByUser(userID, page, pageSize)
}

func (s *NotificationService) MarkRead(id string) (*repository.Notification, error) {
	return s.repo.MarkRead(id)
}

func (s *NotificationService) MarkAllRead(userID string) error {
	return s.repo.MarkAllRead(userID)
}

func (s *NotificationService) GetUnreadCount(userID string) (int64, error) {
	return s.repo.GetUnreadCount(userID)
}
