package handler

import (
	"company-service/internal/repository"
	"company-service/internal/service"
	companyPb "company-service/pb"
	"context"
	"strings"
)

type CompanyHandler struct {
	companyPb.UnimplementedCompanyServiceServer
	svc *service.CompanyService
}

func NewCompanyHandler(svc *service.CompanyService) *CompanyHandler {
	return &CompanyHandler{svc: svc}
}

func (h *CompanyHandler) CreateCompany(ctx context.Context, req *companyPb.CreateCompanyRequest) (*companyPb.Company, error) {
	c, err := h.svc.CreateCompany(req.GetName(), req.GetTaxCode(), req.GetAddress(), req.GetCity(), req.GetDescription(), req.GetOwnerId())
	if err != nil {
		return nil, err
	}
	return companyToProto(c), nil
}

func (h *CompanyHandler) GetCompany(ctx context.Context, req *companyPb.GetCompanyRequest) (*companyPb.Company, error) {
	c, err := h.svc.GetCompany(req.GetId())
	if err != nil {
		return nil, err
	}
	return companyToProto(c), nil
}

func (h *CompanyHandler) GetCompanyByOwner(ctx context.Context, req *companyPb.GetCompanyByOwnerRequest) (*companyPb.Company, error) {
	c, err := h.svc.GetCompanyByOwner(req.GetOwnerId())
	if err != nil {
		return nil, err
	}
	return companyToProto(c), nil
}

func (h *CompanyHandler) ListCompanies(ctx context.Context, req *companyPb.ListCompaniesRequest) (*companyPb.ListCompaniesResponse, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	companies, total, err := h.svc.ListCompanies(req.GetStatus(), page, pageSize)
	if err != nil {
		return nil, err
	}

	var pbCompanies []*companyPb.Company
	for _, c := range companies {
		pbCompanies = append(pbCompanies, companyToProto(&c))
	}

	return &companyPb.ListCompaniesResponse{
		Companies: pbCompanies,
		Total:     int32(total),
	}, nil
}

func (h *CompanyHandler) UpdateCompany(ctx context.Context, req *companyPb.UpdateCompanyRequest) (*companyPb.Company, error) {
	c, err := h.svc.UpdateCompany(req.GetId(), req.GetName(), req.GetAddress(), req.GetCity(), req.GetDescription())
	if err != nil {
		return nil, err
	}
	return companyToProto(c), nil
}

func (h *CompanyHandler) ApproveCompany(ctx context.Context, req *companyPb.ApproveCompanyRequest) (*companyPb.Company, error) {
	c, err := h.svc.ApproveCompany(req.GetId())
	if err != nil {
		return nil, err
	}
	return companyToProto(c), nil
}

func (h *CompanyHandler) RejectCompany(ctx context.Context, req *companyPb.RejectCompanyRequest) (*companyPb.Company, error) {
	c, err := h.svc.RejectCompany(req.GetId(), req.GetReason())
	if err != nil {
		return nil, err
	}
	return companyToProto(c), nil
}

func companyToProto(c *repository.Company) *companyPb.Company {
	certs := []string{}
	if c.Certifications != "" {
		certs = strings.Split(c.Certifications, ",")
	}
	return &companyPb.Company{
		Id:             c.ID,
		Name:           c.Name,
		TaxCode:        c.TaxCode,
		Address:        c.Address,
		City:           c.City,
		Description:    c.Description,
		Status:         c.Status,
		RejectReason:   c.RejectReason,
		OwnerId:        c.OwnerID,
		Rating:         c.Rating,
		ReviewCount:    c.ReviewCount,
		MemberSince:    c.MemberSince.Format("2006-01-02"),
		Certifications: certs,
		ImageUrl:       c.ImageURL,
	}
}
