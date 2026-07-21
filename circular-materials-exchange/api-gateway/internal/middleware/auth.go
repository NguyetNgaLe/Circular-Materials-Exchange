package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
)

var authDB *sql.DB

func init() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5433")
	dbUser := getEnv("DB_USER", "cme")
	dbPass := getEnv("DB_PASSWORD", "")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=auth_db sslmode=disable", dbHost, dbPort, dbUser, dbPass)
	authDB, _ = sql.Open("postgres", dsn)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

type UserInfo struct {
	ID    string
	Email string
	Role  string
}

func JWTAuth() gin.HandlerFunc {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "cme_jwt_secret_2024"
	}

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Missing authorization header"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Handle demo tokens
		if strings.HasPrefix(tokenStr, "demo-token-") {
			email := strings.TrimPrefix(tokenStr, "demo-token-")

			// Query user from DB
			var uid, name, role string
			var phone, companyID sql.NullString
			if authDB != nil {
				err := authDB.QueryRow("SELECT id, name, role, phone, company_id FROM users WHERE email=$1", email).Scan(&uid, &name, &role, &phone, &companyID)
				if err != nil {
					// User not in DB, use fallback
					role = "business"
					if email == "admin@cme.vn" {
						role = "admin"
					}
					userIDs := map[string]string{
						"admin@cme.vn":      "a0000000-0000-0000-0000-000000000001",
						"an@ecopoly.vn":     "a0000000-0000-0000-0000-000000000002",
						"binh@greenpack.vn": "a0000000-0000-0000-0000-000000000003",
					}
					uid = userIDs[email]
					if uid == "" {
						uid = "a0000000-0000-0000-0000-000000000099"
					}
				}
			} else {
				role = "business"
				if email == "admin@cme.vn" {
					role = "admin"
				}
				uid = "a0000000-0000-0000-0000-000000000099"
			}

			c.Set("user_id", uid)
			c.Set("email", email)
			c.Set("role", role)
			c.Next()
			return
		}

		// Real JWT token
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid claims"})
			c.Abort()
			return
		}

		userID, _ := claims["user_id"].(string)
		email, _ := claims["email"].(string)
		role, _ := claims["role"].(string)

		c.Set("user_id", userID)
		c.Set("email", email)
		c.Set("role", role)
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role.(string) != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
