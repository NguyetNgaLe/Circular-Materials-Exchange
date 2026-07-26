package handler

import (
	companypb "api-gateway/internal/pb/company"
	materialpb "api-gateway/internal/pb/material"
	"net/http"
	"strings"

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

func (h *MaterialHandler) ListMyListings(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()

	userID, _ := c.Get("user_id")
	response, err := h.material.ListListings(ctx, &materialpb.ListListingsRequest{
		Page: 1, PageSize: 1000,
	})
	if err != nil {
		writeRPCError(c, err, "Loi lay danh sach nguon cung")
		return
	}

	items := make([]gin.H, 0)
	for _, listing := range response.GetListings() {
		if listing.GetSellerId() == stringValue(userID) {
			items = append(items, listingJSON(listing, ""))
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    gin.H{"listings": items, "total": len(items)},
	})
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

func (h *MaterialHandler) UpdateListing(c *gin.Context) {
	var req struct {
		Title        *string           `json:"title"`
		CategoryID   *string           `json:"category_id"`
		Description  *string           `json:"description"`
		Specs        map[string]string `json:"specs"`
		Quantity     *float64          `json:"quantity"`
		Unit         *string           `json:"unit"`
		PricePerUnit *float64          `json:"price_per_unit"`
		Currency     *string           `json:"currency"`
		Location     *string           `json:"location"`
		MinOrderQty  *float64          `json:"min_order_quantity"`
		Packaging    *string           `json:"packaging"`
		Status       *string           `json:"status"`
		ImageURL     *string           `json:"image_url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	ctx, cancel := rpcContext(c)
	defer cancel()
	current, err := h.material.GetListing(ctx, &materialpb.GetListingRequest{Id: c.Param("id")})
	if err != nil {
		writeRPCError(c, err, "Khong tim thay nguon cung")
		return
	}

	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	if stringValue(role) != "admin" && current.GetSellerId() != stringValue(userID) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Ban khong co quyen chinh sua nguon cung nay",
		})
		return
	}

	title := current.GetTitle()
	categoryID := current.GetCategoryId()
	description := current.GetDescription()
	specs := current.GetSpecs()
	quantity := current.GetQuantity()
	unit := current.GetUnit()
	pricePerUnit := current.GetPricePerUnit()
	currency := current.GetCurrency()
	location := current.GetLocation()
	minOrderQuantity := current.GetMinOrderQuantity()
	packaging := current.GetPackaging()
	status := current.GetStatus()
	images := append([]string(nil), current.GetImages()...)

	if req.Title != nil {
		title = strings.TrimSpace(*req.Title)
		if title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Ten vat lieu khong duoc de trong"})
			return
		}
	}
	if req.CategoryID != nil {
		categoryID = strings.TrimSpace(*req.CategoryID)
		if categoryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Danh muc khong duoc de trong"})
			return
		}
	}
	if req.Description != nil {
		description = *req.Description
	}
	if req.Specs != nil {
		specs = req.Specs
	}
	if req.Quantity != nil {
		if *req.Quantity <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "So luong phai lon hon 0"})
			return
		}
		quantity = *req.Quantity
	}
	if req.Unit != nil {
		unit = strings.TrimSpace(*req.Unit)
		if unit == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Don vi khong duoc de trong"})
			return
		}
	}
	if req.PricePerUnit != nil {
		if *req.PricePerUnit < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Gia khong duoc am"})
			return
		}
		pricePerUnit = *req.PricePerUnit
	}
	if req.Currency != nil {
		currency = strings.TrimSpace(*req.Currency)
	}
	if req.Location != nil {
		location = strings.TrimSpace(*req.Location)
		if location == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Dia diem khong duoc de trong"})
			return
		}
	}
	if req.MinOrderQty != nil {
		if *req.MinOrderQty < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Don hang toi thieu khong duoc am"})
			return
		}
		minOrderQuantity = *req.MinOrderQty
	}
	if req.Packaging != nil {
		packaging = *req.Packaging
	}
	if req.Status != nil {
		status = strings.TrimSpace(*req.Status)
		if status != "active" && status != "hidden" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Trang thai chi co the la active hoac hidden"})
			return
		}
	}
	if req.ImageURL != nil {
		images = nil
		if imageURL := strings.TrimSpace(*req.ImageURL); imageURL != "" {
			images = []string{imageURL}
		}
	}

	updated, err := h.material.UpdateListing(ctx, &materialpb.UpdateListingRequest{
		Id: current.GetId(), Title: title, CategoryId: categoryID,
		Description: description, Specs: specs, Quantity: quantity, Unit: unit,
		PricePerUnit: pricePerUnit, Currency: currency, Location: location,
		MinOrderQuantity: minOrderQuantity, Packaging: packaging, Status: status,
		Images: images,
	})
	if err != nil {
		writeRPCError(c, err, "Loi cap nhat nguon cung")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": listingJSON(updated, "")})
}

func (h *MaterialHandler) DeleteListing(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()

	listing, err := h.material.GetListing(ctx, &materialpb.GetListingRequest{Id: c.Param("id")})
	if err != nil {
		writeRPCError(c, err, "Khong tim thay nguon cung")
		return
	}

	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	if stringValue(role) != "admin" && listing.GetSellerId() != stringValue(userID) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Ban khong co quyen xoa nguon cung nay",
		})
		return
	}

	if _, err := h.material.DeleteListing(ctx, &materialpb.DeleteListingRequest{Id: listing.GetId()}); err != nil {
		writeRPCError(c, err, "Loi xoa nguon cung")
		return
	}

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
