package pb

type Company struct {
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	TaxCode        string   `json:"tax_code"`
	Address        string   `json:"address"`
	City           string   `json:"city"`
	Description    string   `json:"description"`
	Status         string   `json:"status"`
	RejectReason   string   `json:"reject_reason"`
	OwnerId        string   `json:"owner_id"`
	Rating         float64  `json:"rating"`
	ReviewCount    int32    `json:"review_count"`
	MemberSince    string   `json:"member_since"`
	Certifications []string `json:"certifications"`
}

func (x *Company) GetId() string             { return x.Id }
func (x *Company) GetName() string           { return x.Name }
func (x *Company) GetTaxCode() string        { return x.TaxCode }
func (x *Company) GetAddress() string        { return x.Address }
func (x *Company) GetCity() string           { return x.City }
func (x *Company) GetDescription() string    { return x.Description }
func (x *Company) GetStatus() string         { return x.Status }
func (x *Company) GetRejectReason() string   { return x.RejectReason }
func (x *Company) GetOwnerId() string        { return x.OwnerId }
func (x *Company) GetRating() float64        { return x.Rating }
func (x *Company) GetReviewCount() int32     { return x.ReviewCount }
func (x *Company) GetMemberSince() string    { return x.MemberSince }
func (x *Company) GetCertifications() []string { return x.Certifications }

type CreateCompanyRequest struct {
	Name        string `json:"name"`
	TaxCode     string `json:"tax_code"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Description string `json:"description"`
	OwnerId     string `json:"owner_id"`
}

func (x *CreateCompanyRequest) GetName() string        { return x.Name }
func (x *CreateCompanyRequest) GetTaxCode() string     { return x.TaxCode }
func (x *CreateCompanyRequest) GetAddress() string     { return x.Address }
func (x *CreateCompanyRequest) GetCity() string        { return x.City }
func (x *CreateCompanyRequest) GetDescription() string { return x.Description }
func (x *CreateCompanyRequest) GetOwnerId() string     { return x.OwnerId }

type GetCompanyRequest struct {
	Id string `json:"id"`
}

func (x *GetCompanyRequest) GetId() string { return x.Id }

type GetCompanyByOwnerRequest struct {
	OwnerId string `json:"owner_id"`
}

func (x *GetCompanyByOwnerRequest) GetOwnerId() string { return x.OwnerId }

type ListCompaniesRequest struct {
	Status   string `json:"status"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"page_size"`
}

func (x *ListCompaniesRequest) GetStatus() string   { return x.Status }
func (x *ListCompaniesRequest) GetPage() int32      { return x.Page }
func (x *ListCompaniesRequest) GetPageSize() int32  { return x.PageSize }

type ListCompaniesResponse struct {
	Companies []*Company `json:"companies"`
	Total     int32      `json:"total"`
}

func (x *ListCompaniesResponse) GetCompanies() []*Company { return x.Companies }
func (x *ListCompaniesResponse) GetTotal() int32          { return x.Total }

type UpdateCompanyRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Description string `json:"description"`
}

func (x *UpdateCompanyRequest) GetId() string          { return x.Id }
func (x *UpdateCompanyRequest) GetName() string        { return x.Name }
func (x *UpdateCompanyRequest) GetAddress() string     { return x.Address }
func (x *UpdateCompanyRequest) GetCity() string        { return x.City }
func (x *UpdateCompanyRequest) GetDescription() string { return x.Description }

type ApproveCompanyRequest struct {
	Id string `json:"id"`
}

func (x *ApproveCompanyRequest) GetId() string { return x.Id }

type RejectCompanyRequest struct {
	Id     string `json:"id"`
	Reason string `json:"reason"`
}

func (x *RejectCompanyRequest) GetId() string     { return x.Id }
func (x *RejectCompanyRequest) GetReason() string { return x.Reason }
