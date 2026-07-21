package handler

import (
	"context"
	"notification-service/internal/repository"
	"notification-service/internal/service"
	"notification-service/pb"
	"time"
)

type NotificationHandler struct {
	pb.UnimplementedNotificationServiceServer
	svc *service.NotificationService
}

func NewNotificationHandler(svc *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{svc: svc}
}

func (h *NotificationHandler) CreateNotification(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.Notification, error) {
	notification, err := h.svc.CreateNotification(
		req.GetUserId(),
		req.GetTitle(),
		req.GetMessage(),
		req.GetType(),
	)
	if err != nil {
		return nil, err
	}
	return notificationToProto(notification), nil
}

func (h *NotificationHandler) ListNotifications(ctx context.Context, req *pb.ListNotificationsRequest) (*pb.ListNotificationsResponse, error) {
	notifications, total, err := h.svc.ListNotifications(
		req.GetUserId(),
		int(req.GetPage()),
		int(req.GetPageSize()),
	)
	if err != nil {
		return nil, err
	}

	protoNotifications := make([]*pb.Notification, len(notifications))
	for i, n := range notifications {
		protoNotifications[i] = notificationToProto(&n)
	}

	return &pb.ListNotificationsResponse{
		Notifications: protoNotifications,
		Total:         int32(total),
	}, nil
}

func (h *NotificationHandler) MarkRead(ctx context.Context, req *pb.MarkReadRequest) (*pb.Notification, error) {
	notification, err := h.svc.MarkRead(req.GetId())
	if err != nil {
		return nil, err
	}
	return notificationToProto(notification), nil
}

func (h *NotificationHandler) MarkAllRead(ctx context.Context, req *pb.MarkAllReadRequest) (*pb.Empty, error) {
	err := h.svc.MarkAllRead(req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h *NotificationHandler) GetUnreadCount(ctx context.Context, req *pb.GetUnreadCountRequest) (*pb.UnreadCountResponse, error) {
	count, err := h.svc.GetUnreadCount(req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.UnreadCountResponse{
		Count: int32(count),
	}, nil
}

func notificationToProto(n *repository.Notification) *pb.Notification {
	return &pb.Notification{
		Id:        n.ID,
		UserId:    n.UserID,
		Title:     n.Title,
		Message:   n.Message,
		Type:      n.Type,
		Read:      n.Read,
		CreatedAt: n.CreatedAt.Format(time.RFC3339),
	}
}
