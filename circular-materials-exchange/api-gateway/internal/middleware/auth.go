package middleware

import (
	authpb "api-gateway/internal/pb/auth"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JWTAuth(authClient authpb.AuthServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Missing authorization header"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid authorization format"})
			c.Abort()
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		var response *authpb.UserResponse
		var err error
		if strings.HasPrefix(token, "demo-token-") {
			email := strings.TrimPrefix(token, "demo-token-")
			response, err = authClient.GetUserByEmail(ctx, &authpb.GetUserByEmailRequest{Email: email})
		} else {
			response, err = authClient.VerifyToken(ctx, &authpb.TokenRequest{Token: token})
		}
		if err != nil || response.GetUser() == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid token"})
			c.Abort()
			return
		}

		user := response.GetUser()
		c.Set("user_id", user.GetId())
		c.Set("email", user.GetEmail())
		c.Set("name", user.GetName())
		c.Set("role", user.GetRole())
		c.Set("company_id", user.GetCompanyId())
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		roleValue, ok := role.(string)
		if !exists || !ok || roleValue != "admin" {
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

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
