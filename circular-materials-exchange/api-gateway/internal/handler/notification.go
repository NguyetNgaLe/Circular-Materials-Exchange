package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	_ "github.com/lib/pq"
)

type NotificationHandler struct {
	conn   *grpc.ClientConn
	db     *sql.DB
	authDB *sql.DB
}

func NewNotificationHandler(conn *grpc.ClientConn) *NotificationHandler {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5433")
	dbUser := getEnv("DB_USER", "cme")
	dbPass := getEnv("DB_PASSWORD", "cme_secret_2024")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=notif_db sslmode=disable", dbHost, dbPort, dbUser, dbPass)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		db = nil
	}

	authDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=auth_db sslmode=disable", dbHost, dbPort, dbUser, dbPass)
	authDB, err := sql.Open("postgres", authDSN)
	if err != nil {
		authDB = nil
	}

	return &NotificationHandler{conn: conn, db: db, authDB: authDB}
}

// Tao thong bao moi
func (h *NotificationHandler) CreateNotification(userID, title, message, notifType, referenceID string) {
	if h.db == nil {
		return
	}
	id := uuid.New().String()
	h.db.Exec(`INSERT INTO notifications (id, user_id, title, message, type, read) VALUES ($1,$2,$3,$4,$5,false)`,
		id, userID, title, message, notifType)
}

func (h *NotificationHandler) ListNotifications(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"notifications": []gin.H{}, "total": 0}})
		return
	}

	userID, _ := c.Get("user_id")
	rows, err := h.db.Query("SELECT id, title, message, type, read, created_at::text FROM notifications WHERE user_id=$1 ORDER BY created_at DESC LIMIT 50", userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"notifications": []gin.H{}, "total": 0}})
		return
	}
	defer rows.Close()

	var notifications []gin.H
	for rows.Next() {
		var id, title, message, notifType, createdAt string
		var isRead bool
		rows.Scan(&id, &title, &message, &notifType, &isRead, &createdAt)
		notifications = append(notifications, gin.H{
			"id": id, "title": title, "message": message, "type": notifType,
			"read": isRead, "createdAt": createdAt,
		})
	}
	if notifications == nil {
		notifications = []gin.H{}
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"notifications": notifications, "total": len(notifications)}})
}

func (h *NotificationHandler) MarkRead(c *gin.Context) {
	id := c.Param("id")
	if h.db != nil {
		h.db.Exec("UPDATE notifications SET read=true WHERE id=$1", id)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": id, "read": true}})
}

func (h *NotificationHandler) MarkAllRead(c *gin.Context) {
	userID, _ := c.Get("user_id")
	if h.db != nil {
		h.db.Exec("UPDATE notifications SET read=true WHERE user_id=$1", userID)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Da doc tat ca thong bao"})
}

func (h *NotificationHandler) GetUnreadCount(c *gin.Context) {
	userID, _ := c.Get("user_id")
	count := 0
	if h.db != nil {
		h.db.QueryRow("SELECT COUNT(*) FROM notifications WHERE user_id=$1 AND read=false", userID).Scan(&count)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"count": count}})
}
