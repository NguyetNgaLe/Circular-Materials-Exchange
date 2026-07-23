package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type Company struct {
	ID             string
	Name           string
	TaxCode        string
	Address        string
	City           string
	Description    string
	Status         string
	RejectReason   string
	OwnerID        string
	Rating         float64
	ReviewCount    int32
	MemberSince    time.Time
	Certifications string
	ImageURL       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) Create(c *Company) error {
	_, err := r.db.Exec(
		`INSERT INTO companies (id, name, tax_code, address, city, description, status, reject_reason, owner_id, rating, review_count, member_since, certifications, image_url, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`,
		c.ID, c.Name, c.TaxCode, c.Address, c.City, c.Description, c.Status, c.RejectReason, c.OwnerID, c.Rating, c.ReviewCount, c.MemberSince, c.Certifications, c.ImageURL, c.CreatedAt, c.UpdatedAt,
	)
	return err
}

func (r *CompanyRepository) FindByID(id string) (*Company, error) {
	c := &Company{}
	err := r.db.QueryRow(
		`SELECT id, name, COALESCE(tax_code,''), COALESCE(address,''), COALESCE(city,''),
			COALESCE(description,''), COALESCE(status,'pending'), COALESCE(reject_reason,''),
			COALESCE(owner_id::text,''), COALESCE(rating,0), COALESCE(review_count,0),
			COALESCE(member_since,CURRENT_DATE), COALESCE(certifications,''),
			COALESCE(image_url,''), created_at, updated_at
		 FROM companies WHERE id = $1`, id,
	).Scan(&c.ID, &c.Name, &c.TaxCode, &c.Address, &c.City, &c.Description, &c.Status, &c.RejectReason, &c.OwnerID, &c.Rating, &c.ReviewCount, &c.MemberSince, &c.Certifications, &c.ImageURL, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CompanyRepository) FindByOwnerID(ownerID string) (*Company, error) {
	c := &Company{}
	err := r.db.QueryRow(
		`SELECT id, name, COALESCE(tax_code,''), COALESCE(address,''), COALESCE(city,''),
			COALESCE(description,''), COALESCE(status,'pending'), COALESCE(reject_reason,''),
			COALESCE(owner_id::text,''), COALESCE(rating,0), COALESCE(review_count,0),
			COALESCE(member_since,CURRENT_DATE), COALESCE(certifications,''),
			COALESCE(image_url,''), created_at, updated_at
		 FROM companies WHERE owner_id = $1`, ownerID,
	).Scan(&c.ID, &c.Name, &c.TaxCode, &c.Address, &c.City, &c.Description, &c.Status, &c.RejectReason, &c.OwnerID, &c.Rating, &c.ReviewCount, &c.MemberSince, &c.Certifications, &c.ImageURL, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CompanyRepository) List(status string, page, pageSize int) ([]Company, int64, error) {
	var total int64
	countQuery := `SELECT COUNT(*) FROM companies`
	args := []interface{}{}
	if status != "" {
		countQuery += ` WHERE status = $1`
		args = append(args, status)
	}
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, name, COALESCE(tax_code,''), COALESCE(address,''), COALESCE(city,''),
		COALESCE(description,''), COALESCE(status,'pending'), COALESCE(reject_reason,''),
		COALESCE(owner_id::text,''), COALESCE(rating,0), COALESCE(review_count,0),
		COALESCE(member_since,CURRENT_DATE), COALESCE(certifications,''),
		COALESCE(image_url,''), created_at, updated_at FROM companies`
	if status != "" {
		query += ` WHERE status = $1`
	}
	offset := (page - 1) * pageSize
	query += ` ORDER BY created_at DESC LIMIT $` + itoa(len(args)+1) + ` OFFSET $` + itoa(len(args)+2)
	args = append(args, pageSize, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var companies []Company
	for rows.Next() {
		var c Company
		if err := rows.Scan(&c.ID, &c.Name, &c.TaxCode, &c.Address, &c.City, &c.Description, &c.Status, &c.RejectReason, &c.OwnerID, &c.Rating, &c.ReviewCount, &c.MemberSince, &c.Certifications, &c.ImageURL, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, 0, err
		}
		companies = append(companies, c)
	}
	return companies, total, nil
}

func (r *CompanyRepository) Update(c *Company) error {
	c.UpdatedAt = time.Now()
	_, err := r.db.Exec(
		`UPDATE companies SET name=$1, tax_code=$2, address=$3, city=$4, description=$5, status=$6, reject_reason=$7, rating=$8, review_count=$9, member_since=$10, certifications=$11, image_url=$12, updated_at=$13 WHERE id=$14`,
		c.Name, c.TaxCode, c.Address, c.City, c.Description, c.Status, c.RejectReason, c.Rating, c.ReviewCount, c.MemberSince, c.Certifications, c.ImageURL, c.UpdatedAt, c.ID,
	)
	return err
}

func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
