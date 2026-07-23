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
	if err := clients.Connect(); err != nil {
		log.Fatalf("Failed to initialize gRPC clients: %v", err)
	}
	defer clients.Close()

	// Create handlers
	authHandler := handler.NewAuthHandler(clients.Auth, clients.Company)
	companyHandler := handler.NewCompanyHandler(clients.Company)
	materialHandler := handler.NewMaterialHandler(clients.Material, clients.Company)
	orderHandler := handler.NewOrderHandler(clients.Order, clients.Company)
	reviewHandler := handler.NewReviewHandler(clients.Review, clients.Auth)
	notificationHandler := handler.NewNotificationHandler(clients.Notification)
	financeHandler := handler.NewFinanceHandler(clients.Order)
	escrowHandler := handler.NewEscrowHandler(clients.Order)
	uploadHandler := handler.NewUploadHandler(clients.Material)

	// Ket noi OrderHandler voi NotificationHandler de tao thong bao
	orderHandler.SetNotificationHandler(notificationHandler)

	// Setup Gin
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/ready", func(c *gin.Context) {
		if !clients.Ready() {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "not_ready"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})

	api := r.Group("/api")
	authMiddleware := middleware.JWTAuth(clients.Auth)

	// Auth routes (public)
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.GET("/me", authMiddleware, authHandler.GetMe)
	}

	// Public browse routes (no auth required)
	api.GET("/listings", materialHandler.ListListings)
	api.GET("/listings/:id", materialHandler.GetListing)
	api.GET("/categories", materialHandler.ListCategories)
	api.GET("/demands", materialHandler.ListDemands)

	// Upload (auth required)
	api.POST("/upload", authMiddleware, uploadHandler.UploadImage)

	// Protected routes
	protected := api.Group("")
	protected.Use(authMiddleware)
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

		// Seller Wallet
		protected.GET("/seller/wallet", escrowHandler.GetSellerWallet)
		protected.GET("/seller/wallet/transactions", escrowHandler.GetSellerWalletTransactions)
		protected.POST("/seller/withdraw", escrowHandler.CreateWithdrawal)
		protected.GET("/seller/withdrawals", escrowHandler.GetSellerWithdrawals)
	}

	// Admin routes
	admin := api.Group("/admin")
	admin.Use(authMiddleware, middleware.AdminOnly())
	{
		// Finance
		admin.GET("/finance/overview", financeHandler.GetOverview)
		admin.GET("/finance/fees", financeHandler.ListFees)
		admin.GET("/finance/wallet", financeHandler.GetWallet)
		admin.GET("/finance/wallet-transactions", financeHandler.ListWalletTransactions)
		admin.POST("/finance/collect-fee", financeHandler.CollectFee)

		// Escrow
		admin.POST("/escrow", escrowHandler.CreateEscrow)
		admin.GET("/escrow", escrowHandler.ListEscrows)
		admin.POST("/escrow/:id/release", escrowHandler.ReleaseEscrow)

		// Withdrawals
		admin.GET("/withdrawals", escrowHandler.ListWithdrawals)
		admin.POST("/withdrawals/:id/approve", escrowHandler.ApproveWithdrawal)
		admin.POST("/withdrawals/:id/reject", escrowHandler.RejectWithdrawal)
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
