package handler

import (
	companypb "api-gateway/internal/pb/company"
	materialpb "api-gateway/internal/pb/material"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MaterialHandler struct {
	material materialpb.MaterialServiceClient
	company  companypb.CompanyServiceClient
}

func NewMaterialHandler(material materialpb.MaterialServiceClient, company companypb.CompanyServiceClient) *MaterialHandler {
	return &MaterialHandler{material: material, company: company}
}

func (h *MaterialHandler) ListCategories(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.material.ListCategories(ctx, &materialpb.ListCategoriesRequest{})
	if err != nil {
		writeRPCError(c, err, "Loi lay danh muc")
		return
	}
	items := make([]gin.H, 0, len(response.GetCategories()))
	for _, category := range response.GetCategories() {
		items = append(items, gin.H{
			"id": category.GetId(), "name": category.GetName(),
			"icon": category.GetIcon(), "imageUrl": category.GetImageUrl(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": items})
}

func (h *MaterialHandler) ListListings(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	page, pageSize := pagination(c)
	response, err := h.material.ListListings(ctx, &materialpb.ListListingsRequest{
		CategoryId: c.Query("category_id"), Keyword: c.Query("keyword"),
		Location: c.Query("location"), Page: page, PageSize: pageSize,
	})
	if err != nil {
		writeRPCError(c, err, "Loi lay danh sach nguon cung")
		return
	}
	role, _ := c.Get("role")
	items := make([]gin.H, 0, len(response.GetListings()))
	for _, listing := range response.GetListings() {
		if stringValue(role) != "admin" && listing.GetStatus() != "active" {
			continue
		}
		items = append(items, listingJSON(listing, ""))
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"listings": items, "total": response.GetTotal()}})
}

func (h *MaterialHandler) GetListing(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	listing, err := h.material.GetListing(ctx, &materialpb.GetListingRequest{Id: c.Param("id")})
	if err != nil {
		writeRPCError(c, err, "Not found")
		return
	}
	sellerName := ""
	if company, companyErr := getCompanyByOwner(ctx, h.company, listing.GetSellerId()); companyErr == nil {
		sellerName = company.GetName()
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": listingJSON(listing, sellerName)})
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
		ImageURL     string            `json:"image_url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")
	ctx, cancel := rpcContext(c)
	defer cancel()
	company, err := getCompanyByOwner(ctx, h.company, stringValue(userID))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Ban can co ho so doanh nghiep de dang nguon cung"})
		return
	}
	if company.GetStatus() != "verified" {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Doanh nghiep chua duoc admin duyet"})
		return
	}
	images := []string{}
	if req.ImageURL != "" {
		images = append(images, req.ImageURL)
	}
	listing, err := h.material.CreateListing(ctx, &materialpb.CreateListingRequest{
		Title: req.Title, CategoryId: req.CategoryID, SellerId: stringValue(userID),
		CompanyId: company.GetId(), Description: req.Description, Specs: req.Specs,
		Quantity: req.Quantity, Unit: req.Unit, PricePerUnit: req.PricePerUnit,
		Currency: req.Currency, Location: req.Location, MinOrderQuantity: req.MinOrderQty,
		Packaging: req.Packaging, Images: images,
	})
	if err != nil {
		writeRPCError(c, err, "Loi tao nguon cung")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": listing.GetId()}})
}

// These two endpoints intentionally preserve the current public behavior.
func (h *MaterialHandler) UpdateListing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": c.Param("id")}})
}

func (h *MaterialHandler) DeleteListing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "deleted"})
}

// CreateDemand remains a compatibility stub until the product behavior is approved separately.
func (h *MaterialHandler) CreateDemand(c *gin.Context) {
	var req struct {
		Title      string `json:"title" binding:"required"`
		CategoryID string `json:"category_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": "new-demand"}})
}

func (h *MaterialHandler) ListDemands(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	page, pageSize := pagination(c)
	response, err := h.material.ListDemands(ctx, &materialpb.ListDemandsRequest{
		CategoryId: c.Query("category_id"), Keyword: c.Query("keyword"), Page: page, PageSize: pageSize,
	})
	if err != nil {
		writeRPCError(c, err, "Loi lay danh sach nhu cau")
		return
	}
	items := make([]gin.H, 0, len(response.GetDemands()))
	for _, demand := range response.GetDemands() {
		if demand.GetStatus() != "open" {
			continue
		}
		items = append(items, gin.H{
			"id": demand.GetId(), "title": demand.GetTitle(), "category_id": demand.GetCategoryId(),
			"buyer_id": demand.GetBuyerId(), "description": demand.GetDescription(),
			"quantity": demand.GetQuantity(), "unit": demand.GetUnit(),
			"target_price": demand.GetTargetPrice(), "location": demand.GetLocation(),
			"deadline": demand.GetDeadline(), "status": demand.GetStatus(), "created_at": demand.GetCreatedAt(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"demands": items, "total": response.GetTotal()}})
}

func listingJSON(listing *materialpb.SupplyListing, sellerName string) gin.H {
	imageURL := ""
	if len(listing.GetImages()) > 0 {
		imageURL = listing.GetImages()[0]
	}
	data := gin.H{
		"id": listing.GetId(), "title": listing.GetTitle(), "categoryId": listing.GetCategoryId(),
		"sellerId": listing.GetSellerId(), "companyId": listing.GetCompanyId(),
		"description": listing.GetDescription(), "specs": listing.GetSpecs(),
		"quantity": listing.GetQuantity(), "unit": listing.GetUnit(),
		"pricePerUnit": listing.GetPricePerUnit(), "currency": listing.GetCurrency(),
		"location": listing.GetLocation(), "status": listing.GetStatus(),
		"createdAt": listing.GetCreatedAt(), "minOrderQuantity": listing.GetMinOrderQuantity(),
		"packaging": listing.GetPackaging(), "imageUrl": imageURL,
	}
	if sellerName != "" {
		data["sellerName"] = sellerName
	}
	return data
}
