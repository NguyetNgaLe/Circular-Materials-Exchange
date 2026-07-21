package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	conn *grpc.ClientConn
	db   *sql.DB
}

func NewAuthHandler(conn *grpc.ClientConn) *AuthHandler {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5433")
	dbName := "auth_db"
	dbUser := getEnv("DB_USER", "cme")
	dbPass := getEnv("DB_PASSWORD", "")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return &AuthHandler{conn: conn, db: nil}
	}
	return &AuthHandler{conn: conn, db: db}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	if h.db == nil {
		// Fallback demo mode
		role := "business"
		if req.Email == "admin@cme.vn" {
			role = "admin"
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"token": "demo-token-" + req.Email,
				"user": gin.H{
					"id":    "u-" + req.Email,
					"name":  getNameByEmail(req.Email),
					"email": req.Email,
					"role":  role,
				},
			},
		})
		return
	}

	// Query user from DB
	var id, name, email, passwordHash, role string
	var phone, companyID sql.NullString
	err := h.db.QueryRow("SELECT id, name, email, password_hash, role, phone, company_id FROM users WHERE email=$1", req.Email).Scan(
		&id, &name, &email, &passwordHash, &role, &phone, &companyID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Email hoac mat khau khong dung"})
		return
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Email hoac mat khau khong dung"})
		return
	}

	// Generate token
	token := "demo-token-" + email

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"token": token,
			"user": gin.H{
				"id":        id,
				"name":      name,
				"email":     email,
				"role":      role,
				"phone":     phone.String,
				"companyId": companyID.String,
			},
		},
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Phone    string `json:"phone"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	if h.db == nil {
		// Fallback demo mode
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"token": "demo-token-" + req.Email,
				"user": gin.H{
					"id":    "u-" + req.Email,
					"name":  req.Name,
					"email": req.Email,
					"phone": req.Phone,
					"role":  "business",
				},
			},
		})
		return
	}

	// Check if email already exists
	var exists int
	h.db.QueryRow("SELECT COUNT(*) FROM users WHERE email=$1", req.Email).Scan(&exists)
	if exists > 0 {
		c.JSON(http.StatusConflict, gin.H{"success": false, "message": "Email da ton tai"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Loi ma hoa mat khau"})
		return
	}

	// Insert user into DB
	id := uuid.New().String()
	_, err = h.db.Exec("INSERT INTO users (id, name, email, phone, password_hash, role) VALUES ($1,$2,$3,$4,$5,'business')",
		id, req.Name, req.Email, req.Phone, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Loi tao tai khoan: " + err.Error()})
		return
	}

	// Generate token
	token := "demo-token-" + req.Email

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"token": token,
			"user": gin.H{
				"id":    id,
				"name":  req.Name,
				"email": req.Email,
				"phone": req.Phone,
				"role":  "business",
			},
		},
	})
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")
	role, _ := c.Get("role")

	if h.db != nil {
		var id, name, emailAddr, userRole string
		var phone, companyID sql.NullString
		err := h.db.QueryRow("SELECT id, name, email, role, phone, company_id FROM users WHERE id=$1", userID).Scan(
			&id, &name, &emailAddr, &userRole, &phone, &companyID)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data": gin.H{
					"id":        id,
					"name":      name,
					"email":     emailAddr,
					"role":      userRole,
					"phone":     phone.String,
					"companyId": companyID.String,
				},
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id":    userID,
			"email": email,
			"role":  role,
		},
	})
}

func getNameByEmail(email string) string {
	names := map[string]string{
		"admin@cme.vn":      "Admin",
		"an@ecopoly.vn":     "Nguyen Van An",
		"binh@greenpack.vn": "Tran Thi Binh",
	}
	if name, ok := names[email]; ok {
		return name
	}
	return "User"
}
