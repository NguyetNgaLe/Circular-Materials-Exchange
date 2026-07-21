package pb

type Notification struct {
	Id        string `json:"id,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	Title     string `json:"title,omitempty"`
	Message   string `json:"message,omitempty"`
	Type      string `json:"type,omitempty"`
	Read      bool   `json:"read,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (x *Notification) Reset()         {}
func (x *Notification) String() string { return "" }
func (x *Notification) ProtoMessage()  {}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Notification) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Notification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Notification) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Notification) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

func (x *Notification) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateNotificationRequest struct {
	UserId  string `json:"user_id,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
}

func (x *CreateNotificationRequest) Reset()         {}
func (x *CreateNotificationRequest) String() string { return "" }
func (x *CreateNotificationRequest) ProtoMessage()  {}

func (x *CreateNotificationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateNotificationRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateNotificationRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateNotificationRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ListNotificationsRequest struct {
	UserId   string `json:"user_id,omitempty"`
	Page     int32  `json:"page,omitempty"`
	PageSize int32  `json:"page_size,omitempty"`
}

func (x *ListNotificationsRequest) Reset()         {}
func (x *ListNotificationsRequest) String() string { return "" }
func (x *ListNotificationsRequest) ProtoMessage()  {}

func (x *ListNotificationsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListNotificationsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListNotificationsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListNotificationsResponse struct {
	Notifications []*Notification `json:"notifications,omitempty"`
	Total         int32           `json:"total,omitempty"`
}

func (x *ListNotificationsResponse) Reset()         {}
func (x *ListNotificationsResponse) String() string { return "" }
func (x *ListNotificationsResponse) ProtoMessage()  {}

func (x *ListNotificationsResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

func (x *ListNotificationsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type MarkReadRequest struct {
	Id string `json:"id,omitempty"`
}

func (x *MarkReadRequest) Reset()         {}
func (x *MarkReadRequest) String() string { return "" }
func (x *MarkReadRequest) ProtoMessage()  {}

func (x *MarkReadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MarkAllReadRequest struct {
	UserId string `json:"user_id,omitempty"`
}

func (x *MarkAllReadRequest) Reset()         {}
func (x *MarkAllReadRequest) String() string { return "" }
func (x *MarkAllReadRequest) ProtoMessage()  {}

func (x *MarkAllReadRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUnreadCountRequest struct {
	UserId string `json:"user_id,omitempty"`
}

func (x *GetUnreadCountRequest) Reset()         {}
func (x *GetUnreadCountRequest) String() string { return "" }
func (x *GetUnreadCountRequest) ProtoMessage()  {}

func (x *GetUnreadCountRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UnreadCountResponse struct {
	Count int32 `json:"count,omitempty"`
}

func (x *UnreadCountResponse) Reset()         {}
func (x *UnreadCountResponse) String() string { return "" }
func (x *UnreadCountResponse) ProtoMessage()  {}

func (x *UnreadCountResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Empty struct{}

func (x *Empty) Reset()         {}
func (x *Empty) String() string { return "" }
func (x *Empty) ProtoMessage()  {}
