package pb

type Offer struct {
	Id            string  `json:"id,omitempty"`
	Type          string  `json:"type,omitempty"`
	ListingId     string  `json:"listing_id,omitempty"`
	ListingTitle  string  `json:"listing_title,omitempty"`
	BuyerId       string  `json:"buyer_id,omitempty"`
	BuyerName     string  `json:"buyer_name,omitempty"`
	SellerId      string  `json:"seller_id,omitempty"`
	SellerName    string  `json:"seller_name,omitempty"`
	Quantity      float64 `json:"quantity,omitempty"`
	Unit          string  `json:"unit,omitempty"`
	ProposedPrice float64 `json:"proposed_price,omitempty"`
	Currency      string  `json:"currency,omitempty"`
	Message       string  `json:"message,omitempty"`
	Status        string  `json:"status,omitempty"`
	CreatedAt     string  `json:"created_at,omitempty"`
}

func (x *Offer) Reset()         {}
func (x *Offer) String() string { return "" }
func (x *Offer) ProtoMessage()  {}

func (x *Offer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Offer) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Offer) GetListingId() string {
	if x != nil {
		return x.ListingId
	}
	return ""
}

func (x *Offer) GetListingTitle() string {
	if x != nil {
		return x.ListingTitle
	}
	return ""
}

func (x *Offer) GetBuyerId() string {
	if x != nil {
		return x.BuyerId
	}
	return ""
}

func (x *Offer) GetBuyerName() string {
	if x != nil {
		return x.BuyerName
	}
	return ""
}

func (x *Offer) GetSellerId() string {
	if x != nil {
		return x.SellerId
	}
	return ""
}

func (x *Offer) GetSellerName() string {
	if x != nil {
		return x.SellerName
	}
	return ""
}

func (x *Offer) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Offer) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Offer) GetProposedPrice() float64 {
	if x != nil {
		return x.ProposedPrice
	}
	return 0
}

func (x *Offer) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Offer) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Offer) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Offer) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Transaction struct {
	Id             string              `json:"id,omitempty"`
	OfferId        string              `json:"offer_id,omitempty"`
	ListingTitle   string              `json:"listing_title,omitempty"`
	BuyerId        string              `json:"buyer_id,omitempty"`
	BuyerName      string              `json:"buyer_name,omitempty"`
	SellerId       string              `json:"seller_id,omitempty"`
	SellerName     string              `json:"seller_name,omitempty"`
	Quantity       float64             `json:"quantity,omitempty"`
	Unit           string              `json:"unit,omitempty"`
	AgreedPrice    float64             `json:"agreed_price,omitempty"`
	Currency       string              `json:"currency,omitempty"`
	PaymentStatus  string              `json:"payment_status,omitempty"`
	PaymentMethod  string              `json:"payment_method,omitempty"`
	SettlementNote string              `json:"settlement_note,omitempty"`
	Status         string              `json:"status,omitempty"`
	CreatedAt      string              `json:"created_at,omitempty"`
	Events         []*TransactionEvent `json:"events,omitempty"`
}

func (x *Transaction) Reset()         {}
func (x *Transaction) String() string { return "" }
func (x *Transaction) ProtoMessage()  {}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetOfferId() string {
	if x != nil {
		return x.OfferId
	}
	return ""
}

func (x *Transaction) GetListingTitle() string {
	if x != nil {
		return x.ListingTitle
	}
	return ""
}

func (x *Transaction) GetBuyerId() string {
	if x != nil {
		return x.BuyerId
	}
	return ""
}

func (x *Transaction) GetBuyerName() string {
	if x != nil {
		return x.BuyerName
	}
	return ""
}

func (x *Transaction) GetSellerId() string {
	if x != nil {
		return x.SellerId
	}
	return ""
}

func (x *Transaction) GetSellerName() string {
	if x != nil {
		return x.SellerName
	}
	return ""
}

func (x *Transaction) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Transaction) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Transaction) GetAgreedPrice() float64 {
	if x != nil {
		return x.AgreedPrice
	}
	return 0
}

func (x *Transaction) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Transaction) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

func (x *Transaction) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *Transaction) GetSettlementNote() string {
	if x != nil {
		return x.SettlementNote
	}
	return ""
}

func (x *Transaction) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Transaction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Transaction) GetEvents() []*TransactionEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type TransactionEvent struct {
	Id         string `json:"id,omitempty"`
	ActorId    string `json:"actor_id,omitempty"`
	ActorName  string `json:"actor_name,omitempty"`
	FromStatus string `json:"from_status,omitempty"`
	ToStatus   string `json:"to_status,omitempty"`
	Note       string `json:"note,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}

func (x *TransactionEvent) Reset()         {}
func (x *TransactionEvent) String() string { return "" }
func (x *TransactionEvent) ProtoMessage()  {}

func (x *TransactionEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TransactionEvent) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

func (x *TransactionEvent) GetActorName() string {
	if x != nil {
		return x.ActorName
	}
	return ""
}

func (x *TransactionEvent) GetFromStatus() string {
	if x != nil {
		return x.FromStatus
	}
	return ""
}

func (x *TransactionEvent) GetToStatus() string {
	if x != nil {
		return x.ToStatus
	}
	return ""
}

func (x *TransactionEvent) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *TransactionEvent) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateOfferRequest struct {
	Type          string  `json:"type,omitempty"`
	ListingId     string  `json:"listing_id,omitempty"`
	ListingTitle  string  `json:"listing_title,omitempty"`
	BuyerId       string  `json:"buyer_id,omitempty"`
	BuyerName     string  `json:"buyer_name,omitempty"`
	SellerId      string  `json:"seller_id,omitempty"`
	SellerName    string  `json:"seller_name,omitempty"`
	Quantity      float64 `json:"quantity,omitempty"`
	Unit          string  `json:"unit,omitempty"`
	ProposedPrice float64 `json:"proposed_price,omitempty"`
	Currency      string  `json:"currency,omitempty"`
	Message       string  `json:"message,omitempty"`
}

func (x *CreateOfferRequest) Reset()         {}
func (x *CreateOfferRequest) String() string { return "" }
func (x *CreateOfferRequest) ProtoMessage()  {}

func (x *CreateOfferRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateOfferRequest) GetListingId() string {
	if x != nil {
		return x.ListingId
	}
	return ""
}

func (x *CreateOfferRequest) GetListingTitle() string {
	if x != nil {
		return x.ListingTitle
	}
	return ""
}

func (x *CreateOfferRequest) GetBuyerId() string {
	if x != nil {
		return x.BuyerId
	}
	return ""
}

func (x *CreateOfferRequest) GetBuyerName() string {
	if x != nil {
		return x.BuyerName
	}
	return ""
}

func (x *CreateOfferRequest) GetSellerId() string {
	if x != nil {
		return x.SellerId
	}
	return ""
}

func (x *CreateOfferRequest) GetSellerName() string {
	if x != nil {
		return x.SellerName
	}
	return ""
}

func (x *CreateOfferRequest) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CreateOfferRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *CreateOfferRequest) GetProposedPrice() float64 {
	if x != nil {
		return x.ProposedPrice
	}
	return 0
}

func (x *CreateOfferRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CreateOfferRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetOfferRequest struct {
	Id string `json:"id,omitempty"`
}

func (x *GetOfferRequest) Reset()         {}
func (x *GetOfferRequest) String() string { return "" }
func (x *GetOfferRequest) ProtoMessage()  {}

func (x *GetOfferRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListOffersRequest struct {
	UserId   string `json:"user_id,omitempty"`
	Role     string `json:"role,omitempty"`
	Status   string `json:"status,omitempty"`
	Page     int32  `json:"page,omitempty"`
	PageSize int32  `json:"page_size,omitempty"`
}

func (x *ListOffersRequest) Reset()         {}
func (x *ListOffersRequest) String() string { return "" }
func (x *ListOffersRequest) ProtoMessage()  {}

func (x *ListOffersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListOffersRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ListOffersRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ListOffersRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListOffersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListOffersResponse struct {
	Offers []*Offer `json:"offers,omitempty"`
	Total  int32    `json:"total,omitempty"`
}

func (x *ListOffersResponse) Reset()         {}
func (x *ListOffersResponse) String() string { return "" }
func (x *ListOffersResponse) ProtoMessage()  {}

func (x *ListOffersResponse) GetOffers() []*Offer {
	if x != nil {
		return x.Offers
	}
	return nil
}

func (x *ListOffersResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AcceptOfferRequest struct {
	OfferId   string `json:"offer_id,omitempty"`
	ActorId   string `json:"actor_id,omitempty"`
	ActorName string `json:"actor_name,omitempty"`
}

func (x *AcceptOfferRequest) Reset()         {}
func (x *AcceptOfferRequest) String() string { return "" }
func (x *AcceptOfferRequest) ProtoMessage()  {}

func (x *AcceptOfferRequest) GetOfferId() string {
	if x != nil {
		return x.OfferId
	}
	return ""
}

func (x *AcceptOfferRequest) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

func (x *AcceptOfferRequest) GetActorName() string {
	if x != nil {
		return x.ActorName
	}
	return ""
}

type RejectOfferRequest struct {
	OfferId string `json:"offer_id,omitempty"`
	ActorId string `json:"actor_id,omitempty"`
}

func (x *RejectOfferRequest) Reset()         {}
func (x *RejectOfferRequest) String() string { return "" }
func (x *RejectOfferRequest) ProtoMessage()  {}

func (x *RejectOfferRequest) GetOfferId() string {
	if x != nil {
		return x.OfferId
	}
	return ""
}

func (x *RejectOfferRequest) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

type CancelOfferRequest struct {
	OfferId string `json:"offer_id,omitempty"`
	ActorId string `json:"actor_id,omitempty"`
}

func (x *CancelOfferRequest) Reset()         {}
func (x *CancelOfferRequest) String() string { return "" }
func (x *CancelOfferRequest) ProtoMessage()  {}

func (x *CancelOfferRequest) GetOfferId() string {
	if x != nil {
		return x.OfferId
	}
	return ""
}

func (x *CancelOfferRequest) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

type GetTransactionRequest struct {
	Id string `json:"id,omitempty"`
}

func (x *GetTransactionRequest) Reset()         {}
func (x *GetTransactionRequest) String() string { return "" }
func (x *GetTransactionRequest) ProtoMessage()  {}

func (x *GetTransactionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListTransactionsRequest struct {
	UserId   string `json:"user_id,omitempty"`
	Status   string `json:"status,omitempty"`
	Page     int32  `json:"page,omitempty"`
	PageSize int32  `json:"page_size,omitempty"`
}

func (x *ListTransactionsRequest) Reset()         {}
func (x *ListTransactionsRequest) String() string { return "" }
func (x *ListTransactionsRequest) ProtoMessage()  {}

func (x *ListTransactionsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListTransactionsRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ListTransactionsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListTransactionsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListTransactionsResponse struct {
	Transactions []*Transaction `json:"transactions,omitempty"`
	Total        int32          `json:"total,omitempty"`
}

func (x *ListTransactionsResponse) Reset()         {}
func (x *ListTransactionsResponse) String() string { return "" }
func (x *ListTransactionsResponse) ProtoMessage()  {}

func (x *ListTransactionsResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *ListTransactionsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateStatusRequest struct {
	TransactionId string `json:"transaction_id,omitempty"`
	NewStatus     string `json:"new_status,omitempty"`
	ActorId       string `json:"actor_id,omitempty"`
	ActorName     string `json:"actor_name,omitempty"`
	Note          string `json:"note,omitempty"`
}

func (x *UpdateStatusRequest) Reset()         {}
func (x *UpdateStatusRequest) String() string { return "" }
func (x *UpdateStatusRequest) ProtoMessage()  {}

func (x *UpdateStatusRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *UpdateStatusRequest) GetNewStatus() string {
	if x != nil {
		return x.NewStatus
	}
	return ""
}

func (x *UpdateStatusRequest) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

func (x *UpdateStatusRequest) GetActorName() string {
	if x != nil {
		return x.ActorName
	}
	return ""
}

func (x *UpdateStatusRequest) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}
