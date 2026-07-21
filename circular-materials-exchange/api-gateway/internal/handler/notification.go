package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type NotificationHandler struct {
	conn *grpc.ClientConn
}

func NewNotificationHandler(conn *grpc.ClientConn) *NotificationHandler {
	return &NotificationHandler{conn: conn}
}

func (h *NotificationHandler) ListNotifications(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"notifications": []gin.H{},
			"total":         0,
		},
	})
}

func (h *NotificationHandler) MarkRead(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id":   id,
			"read": true,
		},
	})
}

func (h *NotificationHandler) MarkAllRead(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "All notifications marked as read",
	})
}

func (h *NotificationHandler) GetUnreadCount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"count": 0,
		},
	})
}
