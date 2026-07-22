package main

import (
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"
	"api-gateway/internal/proxy"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	httpPort := getEnv("HTTP_PORT", "8080")

	// Connect to gRPC services
	clients := proxy.NewGRPCClients()
	clients.Connect()
	defer clients.Close()

	// Create handlers
	authHandler := handler.NewAuthHandler(clients.AuthConn)
	companyHandler := handler.NewCompanyHandler(clients.CompanyConn)
	materialHandler := handler.NewMaterialHandler(clients.MaterialConn)
	orderHandler := handler.NewOrderHandler(clients.OrderConn)
	reviewHandler := handler.NewReviewHandler(clients.ReviewConn)
	notificationHandler := handler.NewNotificationHandler(clients.NotificationConn)

	// Setup Gin
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api")

	// Auth routes (public)
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.GET("/me", middleware.JWTAuth(), authHandler.GetMe)
	}

	// Public browse routes (no auth required)
	api.GET("/listings", materialHandler.ListListings)
	api.GET("/listings/:id", materialHandler.GetListing)
	api.GET("/categories", materialHandler.ListCategories)
	api.GET("/demands", materialHandler.ListDemands)

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.JWTAuth())
	{
		// Company
		protected.POST("/companies", companyHandler.CreateCompany)
		protected.GET("/companies", companyHandler.ListCompanies)
		protected.GET("/companies/:id", companyHandler.GetCompany)
		protected.POST("/companies/:id/approve", middleware.AdminOnly(), companyHandler.ApproveCompany)
		protected.POST("/companies/:id/reject", middleware.AdminOnly(), companyHandler.RejectCompany)

		// Material listings (write ops)
		protected.POST("/listings", materialHandler.CreateListing)
		protected.PUT("/listings/:id", materialHandler.UpdateListing)
		protected.DELETE("/listings/:id", materialHandler.DeleteListing)

		// Demands (write ops)
		protected.POST("/demands", materialHandler.CreateDemand)

		// Offers
		protected.POST("/offers", orderHandler.CreateOffer)
		protected.GET("/offers", orderHandler.ListOffers)
		protected.POST("/offers/:id/accept", orderHandler.AcceptOffer)
		protected.POST("/offers/:id/reject", orderHandler.RejectOffer)

		// Transactions
		protected.GET("/transactions", orderHandler.ListTransactions)
		protected.GET("/transactions/:id", orderHandler.GetTransaction)
		protected.POST("/transactions/:id/status", orderHandler.UpdateTransactionStatus)

		// Reviews
		protected.POST("/reviews", reviewHandler.CreateReview)
		protected.GET("/reviews", reviewHandler.ListReviews)

		// Notifications
		protected.GET("/notifications", notificationHandler.ListNotifications)
		protected.PUT("/notifications/:id/read", notificationHandler.MarkRead)
		protected.PUT("/notifications/read-all", notificationHandler.MarkAllRead)
		protected.GET("/notifications/unread-count", notificationHandler.GetUnreadCount)
	}

	log.Printf("API Gateway running on :%s", httpPort)
	if err := r.Run(":" + httpPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
