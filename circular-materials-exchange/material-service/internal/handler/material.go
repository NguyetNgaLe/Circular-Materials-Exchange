package handler

import (
	"context"
	"material-service/internal/repository"
	"material-service/internal/service"
	"material-service/pb"
)

type MaterialHandler struct {
	pb.UnimplementedMaterialServiceServer
	svc *service.MaterialService
}

func NewMaterialHandler(svc *service.MaterialService) *MaterialHandler {
	return &MaterialHandler{svc: svc}
}

func (h *MaterialHandler) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	cat, err := h.svc.CreateCategory(req.GetName(), req.GetIcon())
	if err != nil {
		return nil, err
	}
	return categoryToProto(cat), nil
}

func (h *MaterialHandler) ListCategories(ctx context.Context, req *pb.ListCategoriesRequest) (*pb.ListCategoriesResponse, error) {
	categories, err := h.svc.ListCategories()
	if err != nil {
		return nil, err
	}
	protoCats := make([]*pb.Category, len(categories))
	for i, c := range categories {
		protoCats[i] = categoryToProto(&c)
	}
	return &pb.ListCategoriesResponse{Categories: protoCats}, nil
}

func (h *MaterialHandler) CreateListing(ctx context.Context, req *pb.CreateListingRequest) (*pb.SupplyListing, error) {
	listing, err := h.svc.CreateListing(
		req.GetTitle(), req.GetCategoryId(), req.GetSellerId(), req.GetCompanyId(),
		req.GetDescription(), req.GetSpecs(), req.GetQuantity(), req.GetUnit(),
		req.GetPricePerUnit(), req.GetCurrency(), req.GetLocation(),
		req.GetMinOrderQuantity(), req.GetPackaging(), nil,
	)
	if err != nil {
		return nil, err
	}
	return supplyListingToProto(listing), nil
}

func (h *MaterialHandler) GetListing(ctx context.Context, req *pb.GetListingRequest) (*pb.SupplyListing, error) {
	listing, err := h.svc.GetListing(req.GetId())
	if err != nil {
		return nil, err
	}
	return supplyListingToProto(listing), nil
}

func (h *MaterialHandler) ListListings(ctx context.Context, req *pb.ListListingsRequest) (*pb.ListListingsResponse, error) {
	listings, total, err := h.svc.ListListings(
		req.GetCategoryId(), req.GetKeyword(), req.GetLocation(),
		req.GetPage(), req.GetPageSize(),
	)
	if err != nil {
		return nil, err
	}
	protoListings := make([]*pb.SupplyListing, len(listings))
	for i, l := range listings {
		protoListings[i] = supplyListingToProto(&l)
	}
	return &pb.ListListingsResponse{
		Listings: protoListings,
		Total:    int32(total),
	}, nil
}

func (h *MaterialHandler) UpdateListing(ctx context.Context, req *pb.UpdateListingRequest) (*pb.SupplyListing, error) {
	listing, err := h.svc.UpdateListing(
		req.GetId(), req.GetTitle(), req.GetDescription(),
		req.GetQuantity(), req.GetPricePerUnit(), req.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	return supplyListingToProto(listing), nil
}

func (h *MaterialHandler) DeleteListing(ctx context.Context, req *pb.DeleteListingRequest) (*pb.Empty, error) {
	if err := h.svc.DeleteListing(req.GetId()); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h *MaterialHandler) CreateDemand(ctx context.Context, req *pb.CreateDemandRequest) (*pb.DemandListing, error) {
	demand, err := h.svc.CreateDemand(
		req.GetTitle(), req.GetCategoryId(), req.GetBuyerId(), req.GetCompanyId(),
		req.GetDescription(), req.GetQuantity(), req.GetUnit(),
		req.GetTargetPrice(), req.GetLocation(), req.GetDeadline(),
	)
	if err != nil {
		return nil, err
	}
	return demandListingToProto(demand), nil
}

func (h *MaterialHandler) GetDemand(ctx context.Context, req *pb.GetDemandRequest) (*pb.DemandListing, error) {
	demand, err := h.svc.GetDemand(req.GetId())
	if err != nil {
		return nil, err
	}
	return demandListingToProto(demand), nil
}

func (h *MaterialHandler) ListDemands(ctx context.Context, req *pb.ListDemandsRequest) (*pb.ListDemandsResponse, error) {
	demands, total, err := h.svc.ListDemands(
		req.GetCategoryId(), req.GetKeyword(),
		req.GetPage(), req.GetPageSize(),
	)
	if err != nil {
		return nil, err
	}
	protoDemands := make([]*pb.DemandListing, len(demands))
	for i, d := range demands {
		protoDemands[i] = demandListingToProto(&d)
	}
	return &pb.ListDemandsResponse{
		Demands: protoDemands,
		Total:   int32(total),
	}, nil
}

func categoryToProto(c *repository.Category) *pb.Category {
	return &pb.Category{
		Id:   c.ID,
		Name: c.Name,
		Icon: c.Icon,
	}
}

func supplyListingToProto(l *repository.SupplyListing) *pb.SupplyListing {
	return &pb.SupplyListing{
		Id:               l.ID,
		Title:            l.Title,
		CategoryId:       l.CategoryID,
		SellerId:         l.SellerID,
		CompanyId:        l.CompanyID,
		Description:      l.Description,
		Specs:            repository.JSONToSpecs(l.Specs),
		Quantity:         l.Quantity,
		Unit:             l.Unit,
		PricePerUnit:     l.PricePerUnit,
		Currency:         l.Currency,
		Location:         l.Location,
		MinOrderQuantity: l.MinOrderQuantity,
		Packaging:        l.Packaging,
		Status:           l.Status,
		Images:           repository.StringToImages(l.Images),
		CreatedAt:        l.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func demandListingToProto(d *repository.DemandListing) *pb.DemandListing {
	return &pb.DemandListing{
		Id:          d.ID,
		Title:       d.Title,
		CategoryId:  d.CategoryID,
		BuyerId:     d.BuyerID,
		CompanyId:   d.CompanyID,
		Description: d.Description,
		Quantity:    d.Quantity,
		Unit:        d.Unit,
		TargetPrice: d.TargetPrice,
		Location:    d.Location,
		Deadline:    d.Deadline,
		Status:      d.Status,
		CreatedAt:   d.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
