package pb

type Review struct {
	Id            string `json:"id,omitempty"`
	TransactionId string `json:"transaction_id,omitempty"`
	ReviewerId    string `json:"reviewer_id,omitempty"`
	ReviewerName  string `json:"reviewer_name,omitempty"`
	RevieweeId    string `json:"reviewee_id,omitempty"`
	RevieweeName  string `json:"reviewee_name,omitempty"`
	Rating        int32  `json:"rating,omitempty"`
	Comment       string `json:"comment,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
}

func (x *Review) Reset()         {}
func (x *Review) String() string { return "" }
func (x *Review) ProtoMessage()  {}

func (x *Review) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Review) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *Review) GetReviewerId() string {
	if x != nil {
		return x.ReviewerId
	}
	return ""
}

func (x *Review) GetReviewerName() string {
	if x != nil {
		return x.ReviewerName
	}
	return ""
}

func (x *Review) GetRevieweeId() string {
	if x != nil {
		return x.RevieweeId
	}
	return ""
}

func (x *Review) GetRevieweeName() string {
	if x != nil {
		return x.RevieweeName
	}
	return ""
}

func (x *Review) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Review) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Review) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateReviewRequest struct {
	TransactionId string `json:"transaction_id,omitempty"`
	ReviewerId    string `json:"reviewer_id,omitempty"`
	ReviewerName  string `json:"reviewer_name,omitempty"`
	RevieweeId    string `json:"reviewee_id,omitempty"`
	RevieweeName  string `json:"reviewee_name,omitempty"`
	Rating        int32  `json:"rating,omitempty"`
	Comment       string `json:"comment,omitempty"`
}

func (x *CreateReviewRequest) Reset()         {}
func (x *CreateReviewRequest) String() string { return "" }
func (x *CreateReviewRequest) ProtoMessage()  {}

func (x *CreateReviewRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *CreateReviewRequest) GetReviewerId() string {
	if x != nil {
		return x.ReviewerId
	}
	return ""
}

func (x *CreateReviewRequest) GetReviewerName() string {
	if x != nil {
		return x.ReviewerName
	}
	return ""
}

func (x *CreateReviewRequest) GetRevieweeId() string {
	if x != nil {
		return x.RevieweeId
	}
	return ""
}

func (x *CreateReviewRequest) GetRevieweeName() string {
	if x != nil {
		return x.RevieweeName
	}
	return ""
}

func (x *CreateReviewRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *CreateReviewRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type GetReviewRequest struct {
	Id string `json:"id,omitempty"`
}

func (x *GetReviewRequest) Reset()         {}
func (x *GetReviewRequest) String() string { return "" }
func (x *GetReviewRequest) ProtoMessage()  {}

func (x *GetReviewRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListReviewsRequest struct {
	RevieweeId string `json:"reviewee_id,omitempty"`
	Page       int32  `json:"page,omitempty"`
	PageSize   int32  `json:"page_size,omitempty"`
}

func (x *ListReviewsRequest) Reset()         {}
func (x *ListReviewsRequest) String() string { return "" }
func (x *ListReviewsRequest) ProtoMessage()  {}

func (x *ListReviewsRequest) GetRevieweeId() string {
	if x != nil {
		return x.RevieweeId
	}
	return ""
}

func (x *ListReviewsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListReviewsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListReviewsResponse struct {
	Reviews       []*Review `json:"reviews,omitempty"`
	Total         int32     `json:"total,omitempty"`
	AverageRating float64   `json:"average_rating,omitempty"`
}

func (x *ListReviewsResponse) Reset()         {}
func (x *ListReviewsResponse) String() string { return "" }
func (x *ListReviewsResponse) ProtoMessage()  {}

func (x *ListReviewsResponse) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

func (x *ListReviewsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListReviewsResponse) GetAverageRating() float64 {
	if x != nil {
		return x.AverageRating
	}
	return 0
}

type GetUserRatingRequest struct {
	UserId string `json:"user_id,omitempty"`
}

func (x *GetUserRatingRequest) Reset()         {}
func (x *GetUserRatingRequest) String() string { return "" }
func (x *GetUserRatingRequest) ProtoMessage()  {}

func (x *GetUserRatingRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserRatingResponse struct {
	Average float64 `json:"average,omitempty"`
	Count   int32   `json:"count,omitempty"`
}

func (x *UserRatingResponse) Reset()         {}
func (x *UserRatingResponse) String() string { return "" }
func (x *UserRatingResponse) ProtoMessage()  {}

func (x *UserRatingResponse) GetAverage() float64 {
	if x != nil {
		return x.Average
	}
	return 0
}

func (x *UserRatingResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}
