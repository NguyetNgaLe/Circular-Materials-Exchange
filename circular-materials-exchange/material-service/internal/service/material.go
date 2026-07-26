package service

import (
	"database/sql"
	"errors"
	"material-service/internal/repository"
	"time"

	"github.com/google/uuid"
)

type MaterialService struct {
	repo *repository.MaterialRepository
}

func NewMaterialService(repo *repository.MaterialRepository) *MaterialService {
	return &MaterialService{repo: repo}
}

func (s *MaterialService) CreateCategory(name, icon string) (*repository.Category, error) {
	cat := &repository.Category{
		ID:        uuid.New().String(),
		Name:      name,
		Icon:      icon,
		CreatedAt: time.Now(),
	}
	if err := s.repo.CreateCategory(cat); err != nil {
		return nil, err
	}
	return cat, nil
}

func (s *MaterialService) ListCategories() ([]repository.Category, error) {
	return s.repo.ListCategories()
}

func (s *MaterialService) CreateListing(title, categoryID, sellerID, companyID, description string,
	specs map[string]string, quantity float64, unit string, pricePerUnit float64,
	currency, location string, minOrderQuantity float64, packaging string, images []string) (*repository.SupplyListing, error) {

	listing := &repository.SupplyListing{
		ID:               uuid.New().String(),
		Title:            title,
		CategoryID:       categoryID,
		SellerID:         sellerID,
		CompanyID:        companyID,
		Description:      description,
		Specs:            repository.SpecsToJSON(specs),
		Quantity:         quantity,
		Unit:             unit,
		PricePerUnit:     pricePerUnit,
		Currency:         currency,
		Location:         location,
		MinOrderQuantity: minOrderQuantity,
		Packaging:        packaging,
		Status:           "active",
		Images:           repository.ImagesToString(images),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	if len(images) > 0 {
		listing.ImageURL = images[0]
	}

	if err := s.repo.CreateListing(listing); err != nil {
		return nil, err
	}
	return listing, nil
}

func (s *MaterialService) GetListing(id string) (*repository.SupplyListing, error) {
	listing, err := s.repo.FindListingByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("listing not found")
		}
		return nil, err
	}
	return listing, nil
}

func (s *MaterialService) ListListings(categoryID, keyword, location string, page, pageSize int32) ([]repository.SupplyListing, int64, error) {
	return s.repo.ListListings(categoryID, keyword, location, page, pageSize)
}

func (s *MaterialService) UpdateListing(
	id, title, categoryID, description string,
	specs map[string]string,
	quantity float64,
	unit string,
	pricePerUnit float64,
	currency, location string,
	minOrderQuantity float64,
	packaging, status string,
	images []string,
) (*repository.SupplyListing, error) {
	listing, err := s.repo.FindListingByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("listing not found")
		}
		return nil, err
	}

	listing.Title = title
	listing.CategoryID = categoryID
	listing.Description = description
	listing.Specs = repository.SpecsToJSON(specs)
	listing.Quantity = quantity
	listing.Unit = unit
	listing.PricePerUnit = pricePerUnit
	listing.Currency = currency
	listing.Location = location
	listing.MinOrderQuantity = minOrderQuantity
	listing.Packaging = packaging
	listing.Status = status
	listing.Images = repository.ImagesToString(images)
	listing.ImageURL = ""
	if len(images) > 0 {
		listing.ImageURL = images[0]
	}

	if err := s.repo.UpdateListing(listing); err != nil {
		return nil, err
	}
	return listing, nil
}

func (s *MaterialService) DeleteListing(id string) error {
	_, err := s.repo.FindListingByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("listing not found")
		}
		return err
	}
	return s.repo.DeleteListing(id)
}

func (s *MaterialService) CreateDemand(title, categoryID, buyerID, companyID, description string,
	quantity float64, unit string, targetPrice float64, location, deadline string) (*repository.DemandListing, error) {

	demand := &repository.DemandListing{
		ID:          uuid.New().String(),
		Title:       title,
		CategoryID:  categoryID,
		BuyerID:     buyerID,
		CompanyID:   companyID,
		Description: description,
		Quantity:    quantity,
		Unit:        unit,
		TargetPrice: targetPrice,
		Location:    location,
		Deadline:    deadline,
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateDemand(demand); err != nil {
		return nil, err
	}
	return demand, nil
}

func (s *MaterialService) GetDemand(id string) (*repository.DemandListing, error) {
	demand, err := s.repo.FindDemandByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("demand not found")
		}
		return nil, err
	}
	return demand, nil
}

func (s *MaterialService) ListDemands(categoryID, keyword string, page, pageSize int32) ([]repository.DemandListing, int64, error) {
	return s.repo.ListDemands(categoryID, keyword, page, pageSize)
}
