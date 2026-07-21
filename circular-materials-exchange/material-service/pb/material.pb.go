package pb

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func (x *Category) GetId() string   { return x.Id }
func (x *Category) GetName() string { return x.Name }
func (x *Category) GetIcon() string { return x.Icon }

type SupplyListing struct {
	Id               string            `json:"id"`
	Title            string            `json:"title"`
	CategoryId       string            `json:"category_id"`
	SellerId         string            `json:"seller_id"`
	CompanyId        string            `json:"company_id"`
	Description      string            `json:"description"`
	Specs            map[string]string `json:"specs"`
	Quantity         float64           `json:"quantity"`
	Unit             string            `json:"unit"`
	PricePerUnit     float64           `json:"price_per_unit"`
	Currency         string            `json:"currency"`
	Location         string            `json:"location"`
	MinOrderQuantity float64           `json:"min_order_quantity"`
	Packaging        string            `json:"packaging"`
	Status           string            `json:"status"`
	Images           []string          `json:"images"`
	CreatedAt        string            `json:"created_at"`
}

func (x *SupplyListing) GetId() string               { return x.Id }
func (x *SupplyListing) GetTitle() string             { return x.Title }
func (x *SupplyListing) GetCategoryId() string        { return x.CategoryId }
func (x *SupplyListing) GetSellerId() string          { return x.SellerId }
func (x *SupplyListing) GetCompanyId() string         { return x.CompanyId }
func (x *SupplyListing) GetDescription() string       { return x.Description }
func (x *SupplyListing) GetSpecs() map[string]string  { return x.Specs }
func (x *SupplyListing) GetQuantity() float64         { return x.Quantity }
func (x *SupplyListing) GetUnit() string              { return x.Unit }
func (x *SupplyListing) GetPricePerUnit() float64     { return x.PricePerUnit }
func (x *SupplyListing) GetCurrency() string          { return x.Currency }
func (x *SupplyListing) GetLocation() string          { return x.Location }
func (x *SupplyListing) GetMinOrderQuantity() float64 { return x.MinOrderQuantity }
func (x *SupplyListing) GetPackaging() string         { return x.Packaging }
func (x *SupplyListing) GetStatus() string            { return x.Status }
func (x *SupplyListing) GetImages() []string          { return x.Images }
func (x *SupplyListing) GetCreatedAt() string         { return x.CreatedAt }

type DemandListing struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	CategoryId  string  `json:"category_id"`
	BuyerId     string  `json:"buyer_id"`
	CompanyId   string  `json:"company_id"`
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	Unit        string  `json:"unit"`
	TargetPrice float64 `json:"target_price"`
	Location    string  `json:"location"`
	Deadline    string  `json:"deadline"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
}

func (x *DemandListing) GetId() string          { return x.Id }
func (x *DemandListing) GetTitle() string        { return x.Title }
func (x *DemandListing) GetCategoryId() string   { return x.CategoryId }
func (x *DemandListing) GetBuyerId() string      { return x.BuyerId }
func (x *DemandListing) GetCompanyId() string    { return x.CompanyId }
func (x *DemandListing) GetDescription() string  { return x.Description }
func (x *DemandListing) GetQuantity() float64    { return x.Quantity }
func (x *DemandListing) GetUnit() string         { return x.Unit }
func (x *DemandListing) GetTargetPrice() float64 { return x.TargetPrice }
func (x *DemandListing) GetLocation() string     { return x.Location }
func (x *DemandListing) GetDeadline() string     { return x.Deadline }
func (x *DemandListing) GetStatus() string       { return x.Status }
func (x *DemandListing) GetCreatedAt() string    { return x.CreatedAt }

type CreateListingRequest struct {
	Title            string            `json:"title"`
	CategoryId       string            `json:"category_id"`
	SellerId         string            `json:"seller_id"`
	CompanyId        string            `json:"company_id"`
	Description      string            `json:"description"`
	Specs            map[string]string `json:"specs"`
	Quantity         float64           `json:"quantity"`
	Unit             string            `json:"unit"`
	PricePerUnit     float64           `json:"price_per_unit"`
	Currency         string            `json:"currency"`
	Location         string            `json:"location"`
	MinOrderQuantity float64           `json:"min_order_quantity"`
	Packaging        string            `json:"packaging"`
}

func (x *CreateListingRequest) GetTitle() string             { return x.Title }
func (x *CreateListingRequest) GetCategoryId() string        { return x.CategoryId }
func (x *CreateListingRequest) GetSellerId() string          { return x.SellerId }
func (x *CreateListingRequest) GetCompanyId() string         { return x.CompanyId }
func (x *CreateListingRequest) GetDescription() string       { return x.Description }
func (x *CreateListingRequest) GetSpecs() map[string]string  { return x.Specs }
func (x *CreateListingRequest) GetQuantity() float64         { return x.Quantity }
func (x *CreateListingRequest) GetUnit() string              { return x.Unit }
func (x *CreateListingRequest) GetPricePerUnit() float64     { return x.PricePerUnit }
func (x *CreateListingRequest) GetCurrency() string          { return x.Currency }
func (x *CreateListingRequest) GetLocation() string          { return x.Location }
func (x *CreateListingRequest) GetMinOrderQuantity() float64 { return x.MinOrderQuantity }
func (x *CreateListingRequest) GetPackaging() string         { return x.Packaging }

type GetListingRequest struct {
	Id string `json:"id"`
}

func (x *GetListingRequest) GetId() string { return x.Id }

type ListListingsRequest struct {
	CategoryId string `json:"category_id"`
	Keyword    string `json:"keyword"`
	Location   string `json:"location"`
	Page       int32  `json:"page"`
	PageSize   int32  `json:"page_size"`
}

func (x *ListListingsRequest) GetCategoryId() string { return x.CategoryId }
func (x *ListListingsRequest) GetKeyword() string    { return x.Keyword }
func (x *ListListingsRequest) GetLocation() string   { return x.Location }
func (x *ListListingsRequest) GetPage() int32        { return x.Page }
func (x *ListListingsRequest) GetPageSize() int32    { return x.PageSize }

type ListListingsResponse struct {
	Listings []*SupplyListing `json:"listings"`
	Total    int32            `json:"total"`
}

func (x *ListListingsResponse) GetListings() []*SupplyListing { return x.Listings }
func (x *ListListingsResponse) GetTotal() int32               { return x.Total }

type UpdateListingRequest struct {
	Id           string  `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Quantity     float64 `json:"quantity"`
	PricePerUnit float64 `json:"price_per_unit"`
	Status       string  `json:"status"`
}

func (x *UpdateListingRequest) GetId() string           { return x.Id }
func (x *UpdateListingRequest) GetTitle() string        { return x.Title }
func (x *UpdateListingRequest) GetDescription() string  { return x.Description }
func (x *UpdateListingRequest) GetQuantity() float64    { return x.Quantity }
func (x *UpdateListingRequest) GetPricePerUnit() float64 { return x.PricePerUnit }
func (x *UpdateListingRequest) GetStatus() string       { return x.Status }

type DeleteListingRequest struct {
	Id string `json:"id"`
}

func (x *DeleteListingRequest) GetId() string { return x.Id }

type CreateDemandRequest struct {
	Title       string  `json:"title"`
	CategoryId  string  `json:"category_id"`
	BuyerId     string  `json:"buyer_id"`
	CompanyId   string  `json:"company_id"`
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	Unit        string  `json:"unit"`
	TargetPrice float64 `json:"target_price"`
	Location    string  `json:"location"`
	Deadline    string  `json:"deadline"`
}

func (x *CreateDemandRequest) GetTitle() string        { return x.Title }
func (x *CreateDemandRequest) GetCategoryId() string   { return x.CategoryId }
func (x *CreateDemandRequest) GetBuyerId() string      { return x.BuyerId }
func (x *CreateDemandRequest) GetCompanyId() string    { return x.CompanyId }
func (x *CreateDemandRequest) GetDescription() string  { return x.Description }
func (x *CreateDemandRequest) GetQuantity() float64    { return x.Quantity }
func (x *CreateDemandRequest) GetUnit() string         { return x.Unit }
func (x *CreateDemandRequest) GetTargetPrice() float64 { return x.TargetPrice }
func (x *CreateDemandRequest) GetLocation() string     { return x.Location }
func (x *CreateDemandRequest) GetDeadline() string     { return x.Deadline }

type GetDemandRequest struct {
	Id string `json:"id"`
}

func (x *GetDemandRequest) GetId() string { return x.Id }

type ListDemandsRequest struct {
	CategoryId string `json:"category_id"`
	Keyword    string `json:"keyword"`
	Page       int32  `json:"page"`
	PageSize   int32  `json:"page_size"`
}

func (x *ListDemandsRequest) GetCategoryId() string { return x.CategoryId }
func (x *ListDemandsRequest) GetKeyword() string    { return x.Keyword }
func (x *ListDemandsRequest) GetPage() int32        { return x.Page }
func (x *ListDemandsRequest) GetPageSize() int32    { return x.PageSize }

type ListDemandsResponse struct {
	Demands []*DemandListing `json:"demands"`
	Total   int32            `json:"total"`
}

func (x *ListDemandsResponse) GetDemands() []*DemandListing { return x.Demands }
func (x *ListDemandsResponse) GetTotal() int32              { return x.Total }

type ListCategoriesRequest struct{}

func (x *ListCategoriesRequest) Getunused() bool { return false }

type ListCategoriesResponse struct {
	Categories []*Category `json:"categories"`
}

func (x *ListCategoriesResponse) GetCategories() []*Category { return x.Categories }

type CreateCategoryRequest struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func (x *CreateCategoryRequest) GetName() string { return x.Name }
func (x *CreateCategoryRequest) GetIcon() string { return x.Icon }

type Empty struct{}

func (x *Empty) Getunused() bool { return false }
