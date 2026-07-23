package handler

import (
	notificationpb "api-gateway/internal/pb/notification"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	client notificationpb.NotificationServiceClient
}

func NewNotificationHandler(client notificationpb.NotificationServiceClient) *NotificationHandler {
	return &NotificationHandler{client: client}
}

func (h *NotificationHandler) CreateNotification(userID, title, message, notificationType, referenceID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := h.client.CreateNotification(ctx, &notificationpb.CreateNotificationRequest{
		UserId: userID, Title: title, Message: message, Type: notificationType, ReferenceId: referenceID,
	})
	if err != nil {
		log.Printf("create notification failed: %v", err)
	}
}

func (h *NotificationHandler) ListNotifications(c *gin.Context) {
	userID, _ := c.Get("user_id")
	page, pageSize := pagination(c)
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.client.ListNotifications(ctx, &notificationpb.ListNotificationsRequest{
		UserId: stringValue(userID), Page: page, PageSize: pageSize,
	})
	if err != nil {
		writeRPCError(c, err, "Loi lay thong bao")
		return
	}
	items := make([]gin.H, 0, len(response.GetNotifications()))
	for _, notification := range response.GetNotifications() {
		items = append(items, gin.H{
			"id": notification.GetId(), "title": notification.GetTitle(),
			"message": notification.GetMessage(), "type": notification.GetType(),
			"read": notification.GetRead(), "createdAt": notification.GetCreatedAt(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"notifications": items, "total": response.GetTotal()}})
}

func (h *NotificationHandler) MarkRead(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	notification, err := h.client.MarkRead(ctx, &notificationpb.MarkReadRequest{Id: c.Param("id")})
	if err != nil {
		writeRPCError(c, err, "Loi cap nhat thong bao")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": notification.GetId(), "read": notification.GetRead()}})
}

func (h *NotificationHandler) MarkAllRead(c *gin.Context) {
	userID, _ := c.Get("user_id")
	ctx, cancel := rpcContext(c)
	defer cancel()
	if _, err := h.client.MarkAllRead(ctx, &notificationpb.MarkAllReadRequest{UserId: stringValue(userID)}); err != nil {
		writeRPCError(c, err, "Loi cap nhat thong bao")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Da doc tat ca thong bao"})
}

func (h *NotificationHandler) GetUnreadCount(c *gin.Context) {
	userID, _ := c.Get("user_id")
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.client.GetUnreadCount(ctx, &notificationpb.GetUnreadCountRequest{UserId: stringValue(userID)})
	if err != nil {
		writeRPCError(c, err, "Loi dem thong bao")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"count": response.GetCount()}})
}
