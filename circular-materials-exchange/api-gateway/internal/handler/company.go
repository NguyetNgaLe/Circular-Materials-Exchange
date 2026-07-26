package handler

import (
	companypb "api-gateway/internal/pb/company"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	client companypb.CompanyServiceClient
}

func NewCompanyHandler(client companypb.CompanyServiceClient) *CompanyHandler {
	return &CompanyHandler{client: client}
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
	ctx, cancel := rpcContext(c)
	defer cancel()
	company, err := h.client.CreateCompany(ctx, &companypb.CreateCompanyRequest{
		Name: req.Name, TaxCode: req.TaxCode, Address: req.Address,
		City: req.City, Description: req.Description, OwnerId: stringValue(userID),
	})
	if err != nil {
		writeRPCError(c, err, "Loi tao doanh nghiep")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": companyJSON(company)})
}

func (h *CompanyHandler) GetCompany(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	company, err := h.client.GetCompany(ctx, &companypb.GetCompanyRequest{Id: c.Param("id")})
	if err != nil {
		writeRPCError(c, err, "Not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": companyJSON(company)})
}

func (h *CompanyHandler) ListCompanies(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	role, _ := c.Get("role")
	userID, _ := c.Get("user_id")

	companies := []*companypb.Company{}
	total := 0
	if stringValue(role) == "admin" {
		page, pageSize := pagination(c)
		response, err := h.client.ListCompanies(ctx, &companypb.ListCompaniesRequest{
			Status: c.Query("status"), Page: page, PageSize: pageSize,
		})
		if err != nil {
			writeRPCError(c, err, "Loi lay danh sach doanh nghiep")
			return
		}
		companies = response.GetCompanies()
		total = int(response.GetTotal())
	} else {
		company, err := h.client.GetCompanyByOwner(ctx, &companypb.GetCompanyByOwnerRequest{OwnerId: stringValue(userID)})
		if err == nil && company != nil {
			companies = append(companies, company)
		}
		total = len(companies)
	}

	items := make([]gin.H, 0, len(companies))
	for _, company := range companies {
		items = append(items, companyJSON(company))
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"companies": items, "total": total}})
}

func (h *CompanyHandler) ApproveCompany(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	company, err := h.client.ApproveCompany(ctx, &companypb.ApproveCompanyRequest{Id: c.Param("id")})
	if err != nil {
		writeRPCError(c, err, "Loi duyet doanh nghiep")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": company.GetId(), "status": company.GetStatus()}})
}

func (h *CompanyHandler) RejectCompany(c *gin.Context) {
	var req struct {
		Reason string `json:"reason"`
	}
	_ = c.ShouldBindJSON(&req)
	ctx, cancel := rpcContext(c)
	defer cancel()
	company, err := h.client.RejectCompany(ctx, &companypb.RejectCompanyRequest{Id: c.Param("id"), Reason: req.Reason})
	if err != nil {
		writeRPCError(c, err, "Loi tu choi doanh nghiep")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"id": company.GetId(), "status": company.GetStatus(), "reject_reason": company.GetRejectReason(),
	}})
}

func companyJSON(company *companypb.Company) gin.H {
	if company == nil {
		return gin.H{}
	}
	return gin.H{
		"id":             company.GetId(),
		"name":           company.GetName(),
		"taxCode":        company.GetTaxCode(),
		"address":        company.GetAddress(),
		"city":           company.GetCity(),
		"description":    company.GetDescription(),
		"status":         company.GetStatus(),
		"rejectReason":   company.GetRejectReason(),
		"ownerId":        company.GetOwnerId(),
		"rating":         company.GetRating(),
		"reviewCount":    company.GetReviewCount(),
		"memberSince":    company.GetMemberSince(),
		"certifications": company.GetCertifications(),
		"imageUrl":       company.GetImageUrl(),
	}
}

func getCompanyByOwner(ctx context.Context, client companypb.CompanyServiceClient, ownerID string) (*companypb.Company, error) {
	return client.GetCompanyByOwner(ctx, &companypb.GetCompanyByOwnerRequest{OwnerId: ownerID})
}

func stringValue(value interface{}) string {
	if text, ok := value.(string); ok {
		return text
	}
	return ""
}
