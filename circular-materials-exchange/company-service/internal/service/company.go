package service

import (
	"company-service/internal/repository"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type CompanyService struct {
	repo *repository.CompanyRepository
}

func NewCompanyService(repo *repository.CompanyRepository) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) CreateCompany(name, taxCode, address, city, desc, ownerID string) (*repository.Company, error) {
	existing, _ := s.repo.FindByOwnerID(ownerID)
	if existing != nil {
		return nil, errors.New("company already exists for this owner")
	}

	c := &repository.Company{
		ID:          uuid.New().String(),
		Name:        name,
		TaxCode:     taxCode,
		Address:     address,
		City:        city,
		Description: desc,
		Status:      "pending",
		OwnerID:     ownerID,
		MemberSince: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CompanyService) GetCompany(id string) (*repository.Company, error) {
	return s.repo.FindByID(id)
}

func (s *CompanyService) GetCompanyByOwner(ownerID string) (*repository.Company, error) {
	return s.repo.FindByOwnerID(ownerID)
}

func (s *CompanyService) ListCompanies(status string, page, pageSize int) ([]repository.Company, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	return s.repo.List(status, page, pageSize)
}

func (s *CompanyService) UpdateCompany(id, name, address, city, desc string) (*repository.Company, error) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if name != "" {
		c.Name = name
	}
	if address != "" {
		c.Address = address
	}
	if city != "" {
		c.City = city
	}
	if desc != "" {
		c.Description = desc
	}
	if err := s.repo.Update(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CompanyService) ApproveCompany(id string) (*repository.Company, error) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("company not found")
		}
		return nil, err
	}
	c.Status = "verified"
	if err := s.repo.Update(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CompanyService) RejectCompany(id, reason string) (*repository.Company, error) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	c.Status = "rejected"
	c.RejectReason = reason
	if err := s.repo.Update(c); err != nil {
		return nil, err
	}
	return c, nil
}
