package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Category struct {
	ID        string
	Name      string
	Icon      string
	ImageURL  string
	CreatedAt time.Time
}

type SupplyListing struct {
	ID               string
	Title            string
	CategoryID       string
	SellerID         string
	CompanyID        string
	Description      string
	Specs            string
	Quantity         float64
	Unit             string
	PricePerUnit     float64
	Currency         string
	Location         string
	MinOrderQuantity float64
	Packaging        string
	Status           string
	Images           string
	ImageURL         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type DemandListing struct {
	ID          string
	Title       string
	CategoryID  string
	BuyerID     string
	CompanyID   string
	Description string
	Quantity    float64
	Unit        string
	TargetPrice float64
	Location    string
	Deadline    string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MaterialRepository struct {
	db *sql.DB
}

func NewMaterialRepository(db *sql.DB) *MaterialRepository {
	return &MaterialRepository{db: db}
}

func (r *MaterialRepository) CreateCategory(cat *Category) error {
	_, err := r.db.Exec(
		`INSERT INTO categories (id, name, icon, image_url, created_at) VALUES ($1, $2, $3, $4, $5)`,
		cat.ID, cat.Name, cat.Icon, cat.ImageURL, cat.CreatedAt,
	)
	return err
}

func (r *MaterialRepository) ListCategories() ([]Category, error) {
	rows, err := r.db.Query(`SELECT id, name, COALESCE(icon,''), COALESCE(image_url,''), created_at FROM categories ORDER BY name ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Icon, &c.ImageURL, &c.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *MaterialRepository) CreateListing(listing *SupplyListing) error {
	_, err := r.db.Exec(
		`INSERT INTO supply_listings (id, title, category_id, seller_id, company_id, description, specs, quantity, unit, price_per_unit, currency, location, min_order_quantity, packaging, status, images, image_url, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`,
		listing.ID, listing.Title, listing.CategoryID, listing.SellerID, listing.CompanyID, listing.Description, listing.Specs, listing.Quantity, listing.Unit, listing.PricePerUnit, listing.Currency, listing.Location, listing.MinOrderQuantity, listing.Packaging, listing.Status, listing.Images, listing.ImageURL, listing.CreatedAt, listing.UpdatedAt,
	)
	return err
}

func (r *MaterialRepository) FindListingByID(id string) (*SupplyListing, error) {
	l := &SupplyListing{}
	err := r.db.QueryRow(
		`SELECT id, title, COALESCE(category_id::text,''), seller_id,
			COALESCE(company_id::text,''), COALESCE(description,''),
			COALESCE(specs::text,'{}'), COALESCE(quantity,0), COALESCE(unit,''),
			COALESCE(price_per_unit,0), COALESCE(currency,'VND'), COALESCE(location,''),
			COALESCE(min_order_quantity,0), COALESCE(packaging,''), COALESCE(status,'active'),
			COALESCE(images,''), COALESCE(image_url,''), created_at, updated_at
		 FROM supply_listings WHERE id = $1`, id,
	).Scan(&l.ID, &l.Title, &l.CategoryID, &l.SellerID, &l.CompanyID, &l.Description, &l.Specs, &l.Quantity, &l.Unit, &l.PricePerUnit, &l.Currency, &l.Location, &l.MinOrderQuantity, &l.Packaging, &l.Status, &l.Images, &l.ImageURL, &l.CreatedAt, &l.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (r *MaterialRepository) ListListings(categoryID, keyword, location string, page, pageSize int32) ([]SupplyListing, int64, error) {
	var total int64
	countQuery := `SELECT COUNT(*) FROM supply_listings WHERE 1=1`
	args := []interface{}{}
	argIdx := 1

	if categoryID != "" {
		countQuery += fmt.Sprintf(` AND category_id = $%d`, argIdx)
		args = append(args, categoryID)
		argIdx++
	}
	if keyword != "" {
		like := fmt.Sprintf("%%%s%%", keyword)
		countQuery += fmt.Sprintf(` AND (title ILIKE $%d OR description ILIKE $%d)`, argIdx, argIdx+1)
		args = append(args, like, like)
		argIdx += 2
	}
	if location != "" {
		countQuery += fmt.Sprintf(` AND location ILIKE $%d`, argIdx)
		args = append(args, fmt.Sprintf("%%%s%%", location))
		argIdx++
	}

	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, title, COALESCE(category_id::text,''), seller_id,
		COALESCE(company_id::text,''), COALESCE(description,''),
		COALESCE(specs::text,'{}'), COALESCE(quantity,0), COALESCE(unit,''),
		COALESCE(price_per_unit,0), COALESCE(currency,'VND'), COALESCE(location,''),
		COALESCE(min_order_quantity,0), COALESCE(packaging,''), COALESCE(status,'active'),
		COALESCE(images,''), COALESCE(image_url,''), created_at, updated_at
		FROM supply_listings WHERE 1=1`
	args2 := []interface{}{}
	argIdx2 := 1

	if categoryID != "" {
		query += fmt.Sprintf(` AND category_id = $%d`, argIdx2)
		args2 = append(args2, categoryID)
		argIdx2++
	}
	if keyword != "" {
		like := fmt.Sprintf("%%%s%%", keyword)
		query += fmt.Sprintf(` AND (title ILIKE $%d OR description ILIKE $%d)`, argIdx2, argIdx2+1)
		args2 = append(args2, like, like)
		argIdx2 += 2
	}
	if location != "" {
		query += fmt.Sprintf(` AND location ILIKE $%d`, argIdx2)
		args2 = append(args2, fmt.Sprintf("%%%s%%", location))
		argIdx2++
	}

	offset := (page - 1) * pageSize
	query += fmt.Sprintf(` ORDER BY created_at DESC LIMIT $%d OFFSET $%d`, argIdx2, argIdx2+1)
	args2 = append(args2, pageSize, offset)

	rows, err := r.db.Query(query, args2...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var listings []SupplyListing
	for rows.Next() {
		var l SupplyListing
		if err := rows.Scan(&l.ID, &l.Title, &l.CategoryID, &l.SellerID, &l.CompanyID, &l.Description, &l.Specs, &l.Quantity, &l.Unit, &l.PricePerUnit, &l.Currency, &l.Location, &l.MinOrderQuantity, &l.Packaging, &l.Status, &l.Images, &l.ImageURL, &l.CreatedAt, &l.UpdatedAt); err != nil {
			return nil, 0, err
		}
		listings = append(listings, l)
	}
	return listings, total, nil
}

func (r *MaterialRepository) UpdateListing(listing *SupplyListing) error {
	listing.UpdatedAt = time.Now()
	_, err := r.db.Exec(
		`UPDATE supply_listings SET
			title=$1, category_id=$2, description=$3, specs=$4, quantity=$5,
			unit=$6, price_per_unit=$7, currency=$8, location=$9,
			min_order_quantity=$10, packaging=$11, status=$12, images=$13,
			image_url=$14, updated_at=$15
		 WHERE id=$16`,
		listing.Title, listing.CategoryID, listing.Description, listing.Specs,
		listing.Quantity, listing.Unit, listing.PricePerUnit, listing.Currency,
		listing.Location, listing.MinOrderQuantity, listing.Packaging,
		listing.Status, listing.Images, listing.ImageURL, listing.UpdatedAt,
		listing.ID,
	)
	return err
}

func (r *MaterialRepository) DeleteListing(id string) error {
	_, err := r.db.Exec(`DELETE FROM supply_listings WHERE id = $1`, id)
	return err
}

func (r *MaterialRepository) CreateDemand(demand *DemandListing) error {
	_, err := r.db.Exec(
		`INSERT INTO demand_listings (id, title, category_id, buyer_id, company_id, description, quantity, unit, target_price, location, deadline, status, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`,
		demand.ID, demand.Title, demand.CategoryID, demand.BuyerID, demand.CompanyID, demand.Description, demand.Quantity, demand.Unit, demand.TargetPrice, demand.Location, demand.Deadline, demand.Status, demand.CreatedAt, demand.UpdatedAt,
	)
	return err
}

func (r *MaterialRepository) FindDemandByID(id string) (*DemandListing, error) {
	d := &DemandListing{}
	err := r.db.QueryRow(
		`SELECT id, title, COALESCE(category_id::text,''), buyer_id,
			COALESCE(company_id::text,''), COALESCE(description,''), COALESCE(quantity,0),
			COALESCE(unit,''), COALESCE(target_price,0), COALESCE(location,''),
			COALESCE(deadline::text,''), COALESCE(status,'open'), created_at, updated_at
		 FROM demand_listings WHERE id = $1`, id,
	).Scan(&d.ID, &d.Title, &d.CategoryID, &d.BuyerID, &d.CompanyID, &d.Description, &d.Quantity, &d.Unit, &d.TargetPrice, &d.Location, &d.Deadline, &d.Status, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *MaterialRepository) ListDemands(categoryID, keyword string, page, pageSize int32) ([]DemandListing, int64, error) {
	var total int64
	countQuery := `SELECT COUNT(*) FROM demand_listings WHERE 1=1`
	args := []interface{}{}
	argIdx := 1

	if categoryID != "" {
		countQuery += fmt.Sprintf(` AND category_id = $%d`, argIdx)
		args = append(args, categoryID)
		argIdx++
	}
	if keyword != "" {
		like := fmt.Sprintf("%%%s%%", keyword)
		countQuery += fmt.Sprintf(` AND (title ILIKE $%d OR description ILIKE $%d)`, argIdx, argIdx+1)
		args = append(args, like, like)
		argIdx += 2
	}

	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, title, COALESCE(category_id::text,''), buyer_id,
		COALESCE(company_id::text,''), COALESCE(description,''), COALESCE(quantity,0),
		COALESCE(unit,''), COALESCE(target_price,0), COALESCE(location,''),
		COALESCE(deadline::text,''), COALESCE(status,'open'), created_at, updated_at
		FROM demand_listings WHERE 1=1`
	args2 := []interface{}{}
	argIdx2 := 1

	if categoryID != "" {
		query += fmt.Sprintf(` AND category_id = $%d`, argIdx2)
		args2 = append(args2, categoryID)
		argIdx2++
	}
	if keyword != "" {
		like := fmt.Sprintf("%%%s%%", keyword)
		query += fmt.Sprintf(` AND (title ILIKE $%d OR description ILIKE $%d)`, argIdx2, argIdx2+1)
		args2 = append(args2, like, like)
		argIdx2 += 2
	}

	offset := (page - 1) * pageSize
	query += fmt.Sprintf(` ORDER BY created_at DESC LIMIT $%d OFFSET $%d`, argIdx2, argIdx2+1)
	args2 = append(args2, pageSize, offset)

	rows, err := r.db.Query(query, args2...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var demands []DemandListing
	for rows.Next() {
		var d DemandListing
		if err := rows.Scan(&d.ID, &d.Title, &d.CategoryID, &d.BuyerID, &d.CompanyID, &d.Description, &d.Quantity, &d.Unit, &d.TargetPrice, &d.Location, &d.Deadline, &d.Status, &d.CreatedAt, &d.UpdatedAt); err != nil {
			return nil, 0, err
		}
		demands = append(demands, d)
	}
	return demands, total, nil
}

func SpecsToJSON(specs map[string]string) string {
	if specs == nil {
		return "{}"
	}
	b, _ := json.Marshal(specs)
	return string(b)
}

func JSONToSpecs(s string) map[string]string {
	if s == "" {
		return nil
	}
	var specs map[string]string
	json.Unmarshal([]byte(s), &specs)
	return specs
}

func ImagesToString(images []string) string {
	if images == nil {
		return ""
	}
	return strings.Join(images, ",")
}

func StringToImages(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}
