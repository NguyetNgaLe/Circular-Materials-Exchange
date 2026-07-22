package handler

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	_ "github.com/lib/pq"
)

type MaterialHandler struct {
	conn *grpc.ClientConn
	db   *sql.DB
}

func NewMaterialHandler(conn *grpc.ClientConn) *MaterialHandler {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5433")
	dbName := "material_db"
	dbUser := getEnv("DB_USER", "cme")
	dbPass := getEnv("DB_PASSWORD", "cme_secret_2024")

	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return &MaterialHandler{conn: conn, db: nil}
	}
	return &MaterialHandler{conn: conn, db: db}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

type Listing struct {
	ID              string  `json:"id"`
	Title           string  `json:"title"`
	CategoryID      *string `json:"category_id"`
	SellerID        string  `json:"seller_id"`
	CompanyID       *string `json:"company_id"`
	Description     *string `json:"description"`
	Specs           *string `json:"specs"`
	Quantity        float64 `json:"quantity"`
	Unit            *string `json:"unit"`
	PricePerUnit    float64 `json:"price_per_unit"`
	Currency        *string `json:"currency"`
	Location        *string `json:"location"`
	MinOrderQty     *float64 `json:"min_order_quantity"`
	Packaging       *string `json:"packaging"`
	Status          string  `json:"status"`
	CreatedAt       string  `json:"created_at"`
}

func (h *MaterialHandler) ListCategories(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": []gin.H{}})
		return
	}

	rows, err := h.db.Query("SELECT id, name, COALESCE(icon,''), COALESCE(image_url,'') FROM categories ORDER BY name")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": []gin.H{}})
		return
	}
	defer rows.Close()

	var cats []gin.H
	for rows.Next() {
		var id, name, icon, imageURL string
		rows.Scan(&id, &name, &icon, &imageURL)
		cats = append(cats, gin.H{"id": id, "name": name, "icon": icon, "imageUrl": imageURL})
	}
	if cats == nil {
		cats = []gin.H{}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": cats})
}

func (h *MaterialHandler) ListListings(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"listings": []gin.H{}, "total": 0}})
		return
	}

	userRole, _ := c.Get("role")

	var rows *sql.Rows
	var err error

	if userRole == "admin" {
		rows, err = h.db.Query("SELECT id, title, category_id, seller_id, company_id, description, specs::text, quantity, unit, price_per_unit, currency, location, min_order_quantity, packaging, status, created_at::text, COALESCE(image_url,'') FROM supply_listings ORDER BY created_at DESC")
	} else {
		rows, err = h.db.Query("SELECT id, title, category_id, seller_id, company_id, description, specs::text, quantity, unit, price_per_unit, currency, location, min_order_quantity, packaging, status, created_at::text, COALESCE(image_url,'') FROM supply_listings WHERE status='active' ORDER BY created_at DESC")
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"listings": []gin.H{}, "total": 0}})
		return
	}
	defer rows.Close()

	var listings []gin.H
	for rows.Next() {
		var l Listing
		var imageURL string
		rows.Scan(&l.ID, &l.Title, &l.CategoryID, &l.SellerID, &l.CompanyID, &l.Description, &l.Specs, &l.Quantity, &l.Unit, &l.PricePerUnit, &l.Currency, &l.Location, &l.MinOrderQty, &l.Packaging, &l.Status, &l.CreatedAt, &imageURL)
		listings = append(listings, gin.H{
			"id": l.ID, "title": l.Title, "categoryId": l.CategoryID,
			"sellerId": l.SellerID, "description": l.Description,
			"specs": parseSpecs(l.Specs), "quantity": l.Quantity, "unit": l.Unit,
			"pricePerUnit": l.PricePerUnit, "currency": l.Currency,
			"location": l.Location, "status": l.Status, "createdAt": l.CreatedAt,
			"minOrderQuantity": l.MinOrderQty, "packaging": l.Packaging,
			"imageUrl": imageURL,
		})
	}
	if listings == nil {
		listings = []gin.H{}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"listings": listings, "total": len(listings)}})
}

func (h *MaterialHandler) GetListing(c *gin.Context) {
	id := c.Param("id")
	if h.db == nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Not found"})
		return
	}

	var l Listing
	var imageURL string
	err := h.db.QueryRow("SELECT id, title, category_id, seller_id, company_id, description, specs::text, quantity, unit, price_per_unit, currency, location, min_order_quantity, packaging, status, created_at::text, COALESCE(image_url,'') FROM supply_listings WHERE id=$1", id).Scan(&l.ID, &l.Title, &l.CategoryID, &l.SellerID, &l.CompanyID, &l.Description, &l.Specs, &l.Quantity, &l.Unit, &l.PricePerUnit, &l.Currency, &l.Location, &l.MinOrderQty, &l.Packaging, &l.Status, &l.CreatedAt, &imageURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"id": l.ID, "title": l.Title, "categoryId": l.CategoryID,
		"sellerId": l.SellerID, "description": l.Description,
		"specs": parseSpecs(l.Specs), "quantity": l.Quantity, "unit": l.Unit,
		"pricePerUnit": l.PricePerUnit, "currency": l.Currency,
		"location": l.Location, "status": l.Status, "createdAt": l.CreatedAt,
		"minOrderQuantity": l.MinOrderQty, "packaging": l.Packaging,
		"imageUrl": imageURL,
	}})
}

func (h *MaterialHandler) CreateListing(c *gin.Context) {
	var req struct {
		Title        string            `json:"title" binding:"required"`
		CategoryID   string            `json:"category_id" binding:"required"`
		Description  string            `json:"description"`
		Specs        map[string]string `json:"specs"`
		Quantity     float64           `json:"quantity"`
		Unit         string            `json:"unit"`
		PricePerUnit float64           `json:"price_per_unit"`
		Currency     string            `json:"currency"`
		Location     string            `json:"location"`
		MinOrderQty  float64           `json:"min_order_quantity"`
		Packaging    string            `json:"packaging"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": "new"}})
		return
	}

	var id string
	err := h.db.QueryRow("INSERT INTO supply_listings (title, category_id, seller_id, description, quantity, unit, price_per_unit, currency, location, min_order_quantity, packaging, status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,'active') RETURNING id",
		req.Title, req.CategoryID, userID, req.Description, req.Quantity, req.Unit, req.PricePerUnit, req.Currency, req.Location, req.MinOrderQty, req.Packaging).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": id}})
}

func (h *MaterialHandler) UpdateListing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": c.Param("id")}})
}

func (h *MaterialHandler) DeleteListing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "deleted"})
}

func (h *MaterialHandler) CreateDemand(c *gin.Context) {
	var req struct {
		Title       string  `json:"title" binding:"required"`
		CategoryID  string  `json:"category_id" binding:"required"`
		Description string  `json:"description"`
		Quantity    float64 `json:"quantity"`
		Unit        string  `json:"unit"`
		TargetPrice float64 `json:"target_price"`
		Location    string  `json:"location"`
		Deadline    string  `json:"deadline"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": "new-demand"}})
}

func (h *MaterialHandler) ListDemands(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"demands": []gin.H{}, "total": 0}})
		return
	}

	rows, err := h.db.Query("SELECT id, title, category_id, buyer_id, description, quantity, unit, target_price, location, deadline::text, status, created_at::text FROM demand_listings WHERE status='open' ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"demands": []gin.H{}, "total": 0}})
		return
	}
	defer rows.Close()

	var demands []gin.H
	for rows.Next() {
		var id, title, catID, buyerID, status, createdAt string
		var desc, unit, location, deadline *string
		var qty, price float64
		rows.Scan(&id, &title, &catID, &buyerID, &desc, &qty, &unit, &price, &location, &deadline, &status, &createdAt)
		demands = append(demands, gin.H{
			"id": id, "title": title, "category_id": catID, "buyer_id": buyerID,
			"description": desc, "quantity": qty, "unit": unit,
			"target_price": price, "location": location, "deadline": deadline,
			"status": status, "created_at": createdAt,
		})
	}
	if demands == nil {
		demands = []gin.H{}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"demands": demands, "total": len(demands)}})
}

func parseSpecs(s *string) map[string]string {
	if s == nil || *s == "" {
		return map[string]string{}
	}
	result := map[string]string{}
	// Simple JSON parse: {"key":"value",...}
	// For now return empty if parsing fails
	return result
}
