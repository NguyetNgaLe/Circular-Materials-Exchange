package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	_ "github.com/lib/pq"
)

type CompanyHandler struct {
	conn *grpc.ClientConn
	db   *sql.DB
}

func NewCompanyHandler(conn *grpc.ClientConn) *CompanyHandler {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5433")
	dbName := "company_db"
	dbUser := getEnv("DB_USER", "cme")
	dbPass := getEnv("DB_PASSWORD", "cme_secret_2024")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return &CompanyHandler{conn: conn, db: nil}
	}
	return &CompanyHandler{conn: conn, db: db}
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		TaxCode     string `json:"tax_code"`
		Address     string `json:"address"`
		City        string `json:"city"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")

	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": "comp-new", "name": req.Name, "status": "pending"}})
		return
	}

	var id string
	err := h.db.QueryRow(`INSERT INTO companies (name, tax_code, address, city, description, owner_id, status) VALUES ($1,$2,$3,$4,$5,$6,'pending') RETURNING id`,
		req.Name, req.TaxCode, req.Address, req.City, req.Description, userID).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id": id, "name": req.Name, "tax_code": req.TaxCode,
			"address": req.Address, "city": req.City, "description": req.Description,
			"status": "pending", "owner_id": userID,
		},
	})
}

func (h *CompanyHandler) GetCompany(c *gin.Context) {
	id := c.Param("id")
	if h.db == nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Not found"})
		return
	}

	var name, status, ownerID string
	var taxCode, address, city, description, rejectReason, memberSince, certifications, imageURL sql.NullString
	var rating sql.NullFloat64
	var reviewCount sql.NullInt64

	err := h.db.QueryRow(`SELECT name, tax_code, address, city, description, status, reject_reason, owner_id, rating, review_count, member_since::text, certifications, COALESCE(image_url,'') FROM companies WHERE id=$1`, id).Scan(
		&name, &taxCode, &address, &city, &description, &status, &rejectReason, &ownerID, &rating, &reviewCount, &memberSince, &certifications, &imageURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"id": id, "name": name, "taxCode": taxCode.String, "address": address.String,
		"city": city.String, "description": description.String, "status": status,
		"rejectReason": rejectReason.String, "ownerId": ownerID,
		"rating": rating.Float64, "reviewCount": reviewCount.Int64,
		"memberSince": memberSince.String, "certifications": certifications.String,
		"imageUrl": imageURL.String,
	}})
}

func (h *CompanyHandler) ListCompanies(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"companies": []gin.H{}, "total": 0}})
		return
	}

	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	var rows *sql.Rows
	var err error

	if role == "admin" {
		rows, err = h.db.Query(`SELECT id, name, COALESCE(tax_code,''), COALESCE(address,''), COALESCE(city,''), COALESCE(description,''), status, COALESCE(reject_reason,''), owner_id, COALESCE(rating,0), COALESCE(review_count,0), member_since::text, COALESCE(certifications,''), COALESCE(image_url,'') FROM companies ORDER BY created_at DESC`)
	} else {
		rows, err = h.db.Query(`SELECT id, name, COALESCE(tax_code,''), COALESCE(address,''), COALESCE(city,''), COALESCE(description,''), status, COALESCE(reject_reason,''), owner_id, COALESCE(rating,0), COALESCE(review_count,0), member_since::text, COALESCE(certifications,''), COALESCE(image_url,'') FROM companies WHERE owner_id=$1 ORDER BY created_at DESC`, userID)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"companies": []gin.H{}, "total": 0}})
		return
	}
	defer rows.Close()

	var companies []gin.H
	for rows.Next() {
		var id, name, taxCode, address, city, description, status, rejectReason, ownerID, memberSince, certifications, imageURL string
		var rating float64
		var reviewCount int
		rows.Scan(&id, &name, &taxCode, &address, &city, &description, &status, &rejectReason, &ownerID, &rating, &reviewCount, &memberSince, &certifications, &imageURL)
		companies = append(companies, gin.H{
			"id": id, "name": name, "taxCode": taxCode, "address": address,
			"city": city, "description": description, "status": status,
			"rejectReason": rejectReason, "ownerId": ownerID,
			"rating": rating, "reviewCount": reviewCount,
			"memberSince": memberSince, "certifications": certifications,
			"imageUrl": imageURL,
		})
	}
	if companies == nil {
		companies = []gin.H{}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"companies": companies, "total": len(companies)}})
}

func (h *CompanyHandler) ApproveCompany(c *gin.Context) {
	id := c.Param("id")
	if h.db != nil {
		h.db.Exec("UPDATE companies SET status='verified', updated_at=NOW() WHERE id=$1", id)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": id, "status": "verified"}})
}

func (h *CompanyHandler) RejectCompany(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Reason string `json:"reason"`
	}
	c.ShouldBindJSON(&req)

	if h.db != nil {
		h.db.Exec("UPDATE companies SET status='rejected', reject_reason=$1, updated_at=NOW() WHERE id=$2", req.Reason, id)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": id, "status": "rejected", "reject_reason": req.Reason}})
}
