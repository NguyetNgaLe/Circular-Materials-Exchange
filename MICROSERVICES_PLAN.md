# Circular Materials Exchange - Microservices Architecture Plan

> **Trạng thái 24/07/2026:** Plan đã được triển khai và deploy. API Gateway chạy
> trong `cme-network` tại port 8085, gọi 6 service hoàn toàn qua gRPC và không
> truy cập PostgreSQL trực tiếp. Order Service publish `cme.orders.*`;
> Notification Service consume event với `reference_id` idempotent. Một
> PostgreSQL container chứa 6 database logic/18 bảng để phù hợp tài nguyên demo.

## Thông tin project

- **Tên dự án:** Circular Materials Exchange (Sàn giao dịch vật liệu tuần hoàn B2B)
- **Loại project:** Project giữa kỳ môn Phân tích và Thiết kế Hệ thống (PTTKHT)
- **Trạng thái:** MVP / Prototype học thuật
- **Lưu ý:** Các chức năng cần tích hợp bên thứ ba (thanh toán, logistics, SMS, email) được **bypass/manual demo** vì phạm vi môn học.

---

## 1. Tổng quan kiến trúc

```
┌─────────────────────────────────────────────────────────────────┐
│                    React Frontend (port 5173)                    │
│                    (Giữ nguyên code hiện tại)                    │
└──────────────────────────┬──────────────────────────────────────┘
                           │ HTTPS
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│                    API GATEWAY (:8085)                           │
│           Go + Gin │ JWT Auth │ Rate Limit │ CORS               │
└──────┬───────┬───────┬───────┬───────┬───────┬──────────────────┘
       │       │       │       │       │       │
       ▼       ▼       ▼       ▼       ▼       ▼
    ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐
    │Auth │ │Comp-│ │Mate-│ │Order│ │Revi-│ │Notif│
    │Svc  │ │any  │ │rial │ │Svc  │ │ew   │ │Svc  │
    │:50051│ │:50052│ │:50053│ │:50054│ │:50055│ │:50056│
    └─────┘ └─────┘ └─────┘ └─────┘ └─────┘ └─────┘
       │       │       │       │       │       │
       ▼       ▼       ▼       ▼       ▼       ▼
    ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐
    │auth │ │comp-│ │mater│ │order│ │review│ │notif│
    │_db  │ │any_db│ │ial_db│ │_db  │ │_db  │ │_db  │
    └─────┘ └─────┘ └─────┘ └─────┘ └─────┘ └─────┘
                     (PostgreSQL)

                    ┌─────────────┐
                    │NATS JetStream│
                    │  (Async)     │
                    └─────────────┘
```

---

## 2. Danh sách Microservices (6 services)

| # | Service | Port (gRPC) | Database | Responsibility |
|---|---------|-------------|----------|----------------|
| 1 | **auth-service** | :50051 | auth_db | Đăng ký, đăng nhập, JWT, OTP, quản lý user |
| 2 | **company-service** | :50052 | company_db | CRUD doanh nghiệp, xác minh, admin duyệt |
| 3 | **material-service** | :50053 | material_db | Danh mục, nguồn cung, nhu cầu mua, tìm kiếm |
| 4 | **order-service** | :50054 | order_db | Offer, transaction, timeline, xác nhận hoàn tất |
| 5 | **review-service** | :50055 | review_db | Đánh giá buyer/seller, rating |
| 6 | **notification-service** | :50056 | notif_db | Thông báo in-app |

---

## 3. Cấu trúc thư mục

```
circular-materials-exchange/
├── api-gateway/                    # API Gateway
│   ├── cmd/
│   │   └── server/main.go
│   ├── internal/
│   │   ├── handler/               # HTTP handlers
│   │   │   ├── auth.go
│   │   │   ├── company.go
│   │   │   ├── material.go
│   │   │   ├── order.go
│   │   │   ├── review.go
│   │   │   └── notification.go
│   │   ├── middleware/             # JWT, CORS, Rate Limit
│   │   │   ├── auth.go
│   │   │   ├── cors.go
│   │   │   └── ratelimit.go
│   │   └── proxy/                 # gRPC client connections
│   │       └── clients.go
│   ├── go.mod
│   └── go.sum
│
├── auth-service/                   # Authentication Service
│   ├── cmd/
│   │   └── server/main.go
│   ├── internal/
│   │   ├── service/               # Business logic
│   │   │   └── auth.go
│   │   ├── repository/            # Database access
│   │   │   └── user.go
│   │   └── handler/               # gRPC handlers
│   │       └── auth.go
│   ├── migrations/
│   │   └── 001_create_users.sql
│   ├── go.mod
│   └── go.sum
│
├── company-service/                # Company Service
│   ├── cmd/server/main.go
│   ├── internal/
│   │   ├── service/
│   │   ├── repository/
│   │   └── handler/
│   ├── migrations/
│   │   └── 001_create_companies.sql
│   └── go.mod
│
├── material-service/               # Material Service
│   ├── cmd/server/main.go
│   ├── internal/
│   │   ├── service/
│   │   ├── repository/
│   │   └── handler/
│   ├── migrations/
│   │   └── 001_create_materials.sql
│   └── go.mod
│
├── order-service/                  # Order Service
│   ├── cmd/server/main.go
│   ├── internal/
│   │   ├── service/
│   │   ├── repository/
│   │   └── handler/
│   ├── migrations/
│   │   └── 001_create_orders.sql
│   └── go.mod
│
├── review-service/                 # Review Service
│   ├── cmd/server/main.go
│   ├── internal/
│   │   ├── service/
│   │   ├── repository/
│   │   └── handler/
│   ├── migrations/
│   │   └── 001_create_reviews.sql
│   └── go.mod
│
├── notification-service/           # Notification Service
│   ├── cmd/server/main.go
│   ├── internal/
│   │   ├── service/
│   │   ├── repository/
│   │   └── handler/
│   ├── migrations/
│   │   └── 001_create_notifications.sql
│   └── go.mod
│
├── proto/                          # gRPC Proto definitions
│   ├── auth.proto
│   ├── company.proto
│   ├── material.proto
│   ├── order.proto
│   ├── review.proto
│   └── notification.proto
│
├── ui/                             # Frontend React (giữ nguyên)
│   └── ...
│
├── docker-compose.yml              # Docker orchestration
├── Makefile                        # Build & run commands
└── README.md
```

---

## 4. Proto Definitions (gRPC)

### auth.proto

```protobuf
syntax = "proto3";
package auth;
option go_package = "./pb";

service AuthService {
  rpc Register(RegisterRequest) returns (AuthResponse);
  rpc Login(LoginRequest) returns (AuthResponse);
  rpc VerifyToken(TokenRequest) returns (UserResponse);
  rpc GetUser(GetUserRequest) returns (UserResponse);
  rpc UpdateProfile(UpdateProfileRequest) returns (UserResponse);
}

message User {
  string id = 1;
  string name = 2;
  string email = 3;
  string phone = 4;
  string role = 5;
  string avatar = 6;
  string company_id = 7;
}

message RegisterRequest {
  string name = 1;
  string email = 2;
  string password = 3;
  string phone = 4;
}

message LoginRequest {
  string email = 1;
  string password = 2;
}

message AuthResponse {
  string token = 1;
  User user = 2;
}

message TokenRequest {
  string token = 1;
}

message GetUserRequest {
  string id = 1;
}

message UpdateProfileRequest {
  string id = 1;
  string name = 2;
  string phone = 3;
  string avatar = 4;
}

message UserResponse {
  User user = 1;
}
```

### company.proto

```protobuf
syntax = "proto3";
package company;
option go_package = "./pb";

service CompanyService {
  rpc CreateCompany(CreateCompanyRequest) returns (Company);
  rpc GetCompany(GetCompanyRequest) returns (Company);
  rpc ListCompanies(ListCompaniesRequest) returns (ListCompaniesResponse);
  rpc UpdateCompany(UpdateCompanyRequest) returns (Company);
  rpc ApproveCompany(ApproveCompanyRequest) returns (Company);
  rpc RejectCompany(RejectCompanyRequest) returns (Company);
}

message Company {
  string id = 1;
  string name = 2;
  string tax_code = 3;
  string address = 4;
  string city = 5;
  string description = 6;
  string status = 7;       // draft, pending, verified, rejected
  string reject_reason = 8;
  string owner_id = 9;
  double rating = 10;
  int32 review_count = 11;
  string member_since = 12;
  repeated string certifications = 13;
}

message CreateCompanyRequest {
  string name = 1;
  string tax_code = 2;
  string address = 3;
  string city = 4;
  string description = 5;
  string owner_id = 6;
}

message GetCompanyRequest {
  string id = 1;
}

message ListCompaniesRequest {
  string status = 1;  // filter by status
  int32 page = 2;
  int32 page_size = 3;
}

message ListCompaniesResponse {
  repeated Company companies = 1;
  int32 total = 2;
}

message UpdateCompanyRequest {
  string id = 1;
  string name = 2;
  string address = 3;
  string city = 4;
  string description = 5;
}

message ApproveCompanyRequest {
  string id = 1;
}

message RejectCompanyRequest {
  string id = 1;
  string reason = 2;
}
```

### material.proto

```protobuf
syntax = "proto3";
package material;
option go_package = "./pb";

service MaterialService {
  // Categories
  rpc ListCategories(ListCategoriesRequest) returns (ListCategoriesResponse);
  rpc CreateCategory(CreateCategoryRequest) returns (Category);

  // Supply Listings
  rpc CreateListing(CreateListingRequest) returns (SupplyListing);
  rpc GetListing(GetListingRequest) returns (SupplyListing);
  rpc ListListings(ListListingsRequest) returns (ListListingsResponse);
  rpc UpdateListing(UpdateListingRequest) returns (SupplyListing);
  rpc DeleteListing(DeleteListingRequest) returns (Empty);

  // Demand Listings
  rpc CreateDemand(CreateDemandRequest) returns (DemandListing);
  rpc GetDemand(GetDemandRequest) returns (DemandListing);
  rpc ListDemands(ListDemandsRequest) returns (ListDemandsResponse);
}

message Category {
  string id = 1;
  string name = 2;
  string icon = 3;
}

message SupplyListing {
  string id = 1;
  string title = 2;
  string category_id = 3;
  string seller_id = 4;
  string company_id = 5;
  string description = 6;
  map<string, string> specs = 7;
  double quantity = 8;
  string unit = 9;
  double price_per_unit = 10;
  string currency = 11;
  string location = 12;
  double min_order_quantity = 13;
  string packaging = 14;
  string status = 15;    // active, pending_review, sold, hidden
  repeated string images = 16;
  string created_at = 17;
}

message DemandListing {
  string id = 1;
  string title = 2;
  string category_id = 3;
  string buyer_id = 4;
  string company_id = 5;
  string description = 6;
  double quantity = 7;
  string unit = 8;
  double target_price = 9;
  string location = 10;
  string deadline = 11;
  string status = 12;    // open, closed, matched
  string created_at = 13;
}

message CreateListingRequest {
  string title = 1;
  string category_id = 2;
  string seller_id = 3;
  string company_id = 4;
  string description = 5;
  map<string, string> specs = 6;
  double quantity = 7;
  string unit = 8;
  double price_per_unit = 9;
  string currency = 10;
  string location = 11;
  double min_order_quantity = 12;
  string packaging = 13;
}

message GetListingRequest {
  string id = 1;
}

message ListListingsRequest {
  string category_id = 1;
  string keyword = 2;
  string location = 3;
  int32 page = 4;
  int32 page_size = 5;
}

message ListListingsResponse {
  repeated SupplyListing listings = 1;
  int32 total = 2;
}

message UpdateListingRequest {
  string id = 1;
  string title = 2;
  string description = 3;
  double quantity = 4;
  double price_per_unit = 5;
  string status = 6;
}

message DeleteListingRequest {
  string id = 1;
}

message CreateDemandRequest {
  string title = 1;
  string category_id = 2;
  string buyer_id = 3;
  string company_id = 4;
  string description = 5;
  double quantity = 6;
  string unit = 7;
  double target_price = 8;
  string location = 9;
  string deadline = 10;
}

message GetDemandRequest {
  string id = 1;
}

message ListDemandsRequest {
  string category_id = 1;
  string keyword = 2;
  int32 page = 3;
  int32 page_size = 4;
}

message ListDemandsResponse {
  repeated DemandListing demands = 1;
  int32 total = 2;
}

message ListCategoriesRequest {}
message ListCategoriesResponse {
  repeated Category categories = 1;
}

message CreateCategoryRequest {
  string name = 1;
  string icon = 2;
}

message Empty {}
```

### order.proto

```protobuf
syntax = "proto3";
package order;
option go_package = "./pb";

service OrderService {
  rpc CreateOffer(CreateOfferRequest) returns (Offer);
  rpc GetOffer(GetOfferRequest) returns (Offer);
  rpc ListOffers(ListOffersRequest) returns (ListOffersResponse);
  rpc AcceptOffer(AcceptOfferRequest) returns (Transaction);
  rpc RejectOffer(RejectOfferRequest) returns (Offer);
  rpc CancelOffer(CancelOfferRequest) returns (Offer);

  rpc GetTransaction(GetTransactionRequest) returns (Transaction);
  rpc ListTransactions(ListTransactionsRequest) returns (ListTransactionsResponse);
  rpc UpdateTransactionStatus(UpdateStatusRequest) returns (Transaction);
}

message Offer {
  string id = 1;
  string type = 2;           // buyer_to_seller / seller_to_buyer
  string listing_id = 3;
  string listing_title = 4;
  string buyer_id = 5;
  string buyer_name = 6;
  string seller_id = 7;
  string seller_name = 8;
  double quantity = 9;
  string unit = 10;
  double proposed_price = 11;
  string currency = 12;
  string message = 13;
  string status = 14;        // pending, accepted, rejected, cancelled, expired
  string created_at = 15;
}

message Transaction {
  string id = 1;
  string offer_id = 2;
  string listing_title = 3;
  string buyer_id = 4;
  string buyer_name = 5;
  string seller_id = 6;
  string seller_name = 7;
  double quantity = 8;
  string unit = 9;
  double agreed_price = 10;
  string currency = 11;
  string payment_status = 12;   // bypassed_demo (demo mode)
  string payment_method = 13;   // manual_offline (demo mode)
  string settlement_note = 14;
  string status = 15;           // confirmed, in_progress, buyer_confirmed, seller_confirmed, completed, cancelled, disputed
  string created_at = 16;
  repeated TransactionEvent events = 17;
}

message TransactionEvent {
  string id = 1;
  string actor_id = 2;
  string actor_name = 3;
  string from_status = 4;
  string to_status = 5;
  string note = 6;
  string created_at = 7;
}

message CreateOfferRequest {
  string type = 1;
  string listing_id = 2;
  string listing_title = 3;
  string buyer_id = 4;
  string buyer_name = 5;
  string seller_id = 6;
  string seller_name = 7;
  double quantity = 8;
  string unit = 9;
  double proposed_price = 10;
  string currency = 11;
  string message = 12;
}

message GetOfferRequest {
  string id = 1;
}

message ListOffersRequest {
  string user_id = 1;
  string role = 2;      // buyer / seller
  string status = 3;
  int32 page = 4;
  int32 page_size = 5;
}

message ListOffersResponse {
  repeated Offer offers = 1;
  int32 total = 2;
}

message AcceptOfferRequest {
  string offer_id = 1;
  string actor_id = 2;
  string actor_name = 3;
}

message RejectOfferRequest {
  string offer_id = 1;
  string actor_id = 2;
}

message CancelOfferRequest {
  string offer_id = 1;
  string actor_id = 2;
}

message GetTransactionRequest {
  string id = 1;
}

message ListTransactionsRequest {
  string user_id = 1;
  string status = 2;
  int32 page = 3;
  int32 page_size = 4;
}

message ListTransactionsResponse {
  repeated Transaction transactions = 1;
  int32 total = 2;
}

message UpdateStatusRequest {
  string transaction_id = 1;
  string new_status = 2;
  string actor_id = 3;
  string actor_name = 4;
  string note = 5;
}
```

### review.proto

```protobuf
syntax = "proto3";
package review;
option go_package = "./pb";

service ReviewService {
  rpc CreateReview(CreateReviewRequest) returns (Review);
  rpc GetReview(GetReviewRequest) returns (Review);
  rpc ListReviews(ListReviewsRequest) returns (ListReviewsResponse);
  rpc GetUserRating(GetUserRatingRequest) returns (UserRatingResponse);
}

message Review {
  string id = 1;
  string transaction_id = 2;
  string reviewer_id = 3;
  string reviewer_name = 4;
  string reviewee_id = 5;
  string reviewee_name = 6;
  int32 rating = 7;
  string comment = 8;
  string created_at = 9;
}

message CreateReviewRequest {
  string transaction_id = 1;
  string reviewer_id = 2;
  string reviewer_name = 3;
  string reviewee_id = 4;
  string reviewee_name = 5;
  int32 rating = 6;
  string comment = 7;
}

message GetReviewRequest {
  string id = 1;
}

message ListReviewsRequest {
  string reviewee_id = 1;
  int32 page = 2;
  int32 page_size = 3;
}

message ListReviewsResponse {
  repeated Review reviews = 1;
  int32 total = 2;
  double average_rating = 3;
}

message GetUserRatingRequest {
  string user_id = 1;
}

message UserRatingResponse {
  double average = 1;
  int32 count = 2;
}
```

### notification.proto

```protobuf
syntax = "proto3";
package notification;
option go_package = "./pb";

service NotificationService {
  rpc CreateNotification(CreateNotificationRequest) returns (Notification);
  rpc ListNotifications(ListNotificationsRequest) returns (ListNotificationsResponse);
  rpc MarkRead(MarkReadRequest) returns (Notification);
  rpc MarkAllRead(MarkAllReadRequest) returns (Empty);
  rpc GetUnreadCount(GetUnreadCountRequest) returns (UnreadCountResponse);
}

message Notification {
  string id = 1;
  string user_id = 2;
  string title = 3;
  string message = 4;
  string type = 5;    // offer, transaction, system, review
  bool read = 6;
  string created_at = 7;
}

message CreateNotificationRequest {
  string user_id = 1;
  string title = 2;
  string message = 3;
  string type = 4;
}

message ListNotificationsRequest {
  string user_id = 1;
  int32 page = 2;
  int32 page_size = 3;
}

message ListNotificationsResponse {
  repeated Notification notifications = 1;
  int32 total = 2;
}

message MarkReadRequest {
  string id = 1;
}

message MarkAllReadRequest {
  string user_id = 1;
}

message GetUnreadCountRequest {
  string user_id = 1;
}

message UnreadCountResponse {
  int32 count = 1;
}

message Empty {}
```

---

## 5. Communication Patterns

### 5.1 REST API (Client → Gateway → Services)

```
React App → API Gateway (:8080) → gRPC → Backend Services
```

**Response Envelope:**
```json
{
  "success": true,
  "message": "OK",
  "code": "SUCCESS",
  "request_id": "uuid",
  "data": { ... }
}
```

### 5.2 gRPC (Service ↔ Service)

```
API Gateway ──gRPC──▶ auth-service (:50051)
API Gateway ──gRPC──▶ company-service (:50052)
API Gateway ──gRPC──▶ material-service (:50053)
API Gateway ──gRPC──▶ order-service (:50054)
API Gateway ──gRPC──▶ review-service (:50055)
API Gateway ──gRPC──▶ notification-service (:50056)

order-service ──gRPC──▶ material-service (verify listing)
order-service ──gRPC──▶ company-service (verify company)
```

### 5.3 NATS JetStream (Async Events)

```
┌──────────────────────────────────────────────────────────┐
│                    NATS JetStream                         │
└──────────────────────────────────────────────────────────┘
       ▲                    ▲                    ▲
       │                    │                    │
  order-service       order-service       review-service
  (Publisher)          (Publisher)          (Publisher)
       │                    │                    │
       ▼                    ▼                    ▼
  notification-svc    review-service     notification-svc
  (Consumer)           (Consumer)          (Consumer)
```

**NATS Subjects:**

| Subject | Source | Consumer | Purpose |
|---------|--------|----------|---------|
| `cme.orders.offer.created` | order-service | notification-service | Thông báo có offer mới |
| `cme.orders.offer.accepted` | order-service | notification-service | Thông báo offer được chấp nhận |
| `cme.orders.transaction.created` | order-service | notification-service | Thông báo giao dịch mới |
| `cme.orders.transaction.completed` | order-service | review-service, notification-service | Giao dịch hoàn tất |
| `cme.companies.approved` | company-service | notification-service | Doanh nghiệp được duyệt |

---

## 6. Database Schema

### auth_db

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'business',  -- business, admin
    avatar VARCHAR(500),
    company_id UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
```

### company_db

```sql
CREATE TABLE companies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    tax_code VARCHAR(50) UNIQUE,
    address TEXT,
    city VARCHAR(100),
    description TEXT,
    status VARCHAR(20) DEFAULT 'pending',  -- draft, pending, verified, rejected
    reject_reason TEXT,
    owner_id UUID NOT NULL,
    rating DECIMAL(3,2) DEFAULT 0,
    review_count INT DEFAULT 0,
    member_since DATE DEFAULT CURRENT_DATE,
    certifications TEXT[],  -- array of certification names
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_companies_owner ON companies(owner_id);
CREATE INDEX idx_companies_status ON companies(status);
```

### material_db

```sql
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    icon VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE supply_listings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    category_id UUID REFERENCES categories(id),
    seller_id UUID NOT NULL,
    company_id UUID NOT NULL,
    description TEXT,
    specs JSONB DEFAULT '{}',
    quantity DECIMAL(10,2),
    unit VARCHAR(20),
    price_per_unit DECIMAL(15,2),
    currency VARCHAR(10) DEFAULT 'VND',
    location VARCHAR(255),
    min_order_quantity DECIMAL(10,2),
    packaging VARCHAR(100),
    status VARCHAR(20) DEFAULT 'active',  -- active, pending_review, sold, hidden
    images TEXT[],
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE demand_listings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    category_id UUID REFERENCES categories(id),
    buyer_id UUID NOT NULL,
    company_id UUID NOT NULL,
    description TEXT,
    quantity DECIMAL(10,2),
    unit VARCHAR(20),
    target_price DECIMAL(15,2),
    location VARCHAR(255),
    deadline DATE,
    status VARCHAR(20) DEFAULT 'open',  -- open, closed, matched
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_supply_category ON supply_listings(category_id);
CREATE INDEX idx_supply_seller ON supply_listings(seller_id);
CREATE INDEX idx_supply_status ON supply_listings(status);
CREATE INDEX idx_demand_category ON demand_listings(category_id);
CREATE INDEX idx_demand_buyer ON demand_listings(buyer_id);
```

### order_db

```sql
CREATE TABLE offers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type VARCHAR(20) NOT NULL,  -- buyer_to_seller, seller_to_buyer
    listing_id UUID NOT NULL,
    listing_title VARCHAR(255),
    buyer_id UUID NOT NULL,
    buyer_name VARCHAR(255),
    seller_id UUID NOT NULL,
    seller_name VARCHAR(255),
    quantity DECIMAL(10,2),
    unit VARCHAR(20),
    proposed_price DECIMAL(15,2),
    currency VARCHAR(10) DEFAULT 'VND',
    message TEXT,
    status VARCHAR(20) DEFAULT 'pending',  -- pending, accepted, rejected, cancelled, expired
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    offer_id UUID REFERENCES offers(id),
    listing_title VARCHAR(255),
    buyer_id UUID NOT NULL,
    buyer_name VARCHAR(255),
    seller_id UUID NOT NULL,
    seller_name VARCHAR(255),
    quantity DECIMAL(10,2),
    unit VARCHAR(20),
    agreed_price DECIMAL(15,2),
    currency VARCHAR(10) DEFAULT 'VND',
    payment_status VARCHAR(20) DEFAULT 'bypassed_demo',  -- bypassed_demo (MVP mode)
    payment_method VARCHAR(20) DEFAULT 'manual_offline',  -- manual_offline (MVP mode)
    settlement_note TEXT DEFAULT 'Thanh toan duoc thuc hien ngoai he thong trong pham vi prototype',
    status VARCHAR(20) DEFAULT 'confirmed',  -- confirmed, in_progress, buyer_confirmed, seller_confirmed, completed, cancelled, disputed
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE transaction_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transaction_id UUID REFERENCES transactions(id),
    actor_id UUID NOT NULL,
    actor_name VARCHAR(255),
    from_status VARCHAR(50),
    to_status VARCHAR(50),
    note TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_offers_buyer ON offers(buyer_id);
CREATE INDEX idx_offers_seller ON offers(seller_id);
CREATE INDEX idx_offers_status ON offers(status);
CREATE INDEX idx_transactions_buyer ON transactions(buyer_id);
CREATE INDEX idx_transactions_seller ON transactions(seller_id);
CREATE INDEX idx_transactions_status ON transactions(status);
CREATE INDEX idx_tx_events_transaction ON transaction_events(transaction_id);
```

### review_db

```sql
CREATE TABLE reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transaction_id UUID NOT NULL,
    reviewer_id UUID NOT NULL,
    reviewer_name VARCHAR(255),
    reviewee_id UUID NOT NULL,
    reviewee_name VARCHAR(255),
    rating INT CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_reviews_reviewee ON reviews(reviewee_id);
CREATE INDEX idx_reviews_transaction ON reviews(transaction_id);
```

### notif_db

```sql
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    title VARCHAR(255),
    message TEXT,
    type VARCHAR(20),  -- offer, transaction, system, review
    read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_notifications_user ON notifications(user_id);
CREATE INDEX idx_notifications_read ON notifications(user_id, read);
```

---

## 7. Docker Compose

```yaml
version: '3.8'

services:
  # ========== Databases ==========
  auth-db:
    image: postgres:15
    container_name: cme-auth-db
    environment:
      POSTGRES_DB: auth_db
      POSTGRES_USER: cme
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    ports:
      - "5432:5432"
    volumes:
      - auth-db-data:/var/lib/postgresql/data
    networks:
      - cme-network

  company-db:
    image: postgres:15
    container_name: cme-company-db
    environment:
      POSTGRES_DB: company_db
      POSTGRES_USER: cme
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    ports:
      - "5433:5432"
    volumes:
      - company-db-data:/var/lib/postgresql/data
    networks:
      - cme-network

  material-db:
    image: postgres:15
    container_name: cme-material-db
    environment:
      POSTGRES_DB: material_db
      POSTGRES_USER: cme
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    ports:
      - "5434:5432"
    volumes:
      - material-db-data:/var/lib/postgresql/data
    networks:
      - cme-network

  order-db:
    image: postgres:15
    container_name: cme-order-db
    environment:
      POSTGRES_DB: order_db
      POSTGRES_USER: cme
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    ports:
      - "5435:5432"
    volumes:
      - order-db-data:/var/lib/postgresql/data
    networks:
      - cme-network

  review-db:
    image: postgres:15
    container_name: cme-review-db
    environment:
      POSTGRES_DB: review_db
      POSTGRES_USER: cme
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    ports:
      - "5436:5432"
    volumes:
      - review-db-data:/var/lib/postgresql/data
    networks:
      - cme-network

  notif-db:
    image: postgres:15
    container_name: cme-notif-db
    environment:
      POSTGRES_DB: notif_db
      POSTGRES_USER: cme
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    ports:
      - "5437:5432"
    volumes:
      - notif-db-data:/var/lib/postgresql/data
    networks:
      - cme-network

  # ========== NATS ==========
  nats:
    image: nats:latest
    container_name: cme-nats
    command: "-js"
    ports:
      - "4222:4222"
      - "8222:8222"
    networks:
      - cme-network

  # ========== Services ==========
  api-gateway:
    build:
      context: ./api-gateway
      dockerfile: Dockerfile
    container_name: cme-api-gateway
    ports:
      - "8080:8080"
    environment:
      - AUTH_SERVICE_ADDR=auth-service:50051
      - COMPANY_SERVICE_ADDR=company-service:50052
      - MATERIAL_SERVICE_ADDR=material-service:50053
      - ORDER_SERVICE_ADDR=order-service:50054
      - REVIEW_SERVICE_ADDR=review-service:50055
      - NOTIFICATION_SERVICE_ADDR=notification-service:50056
      - JWT_SECRET=${JWT_SECRET}
    depends_on:
      - auth-service
      - company-service
      - material-service
      - order-service
      - review-service
      - notification-service
    networks:
      - cme-network

  auth-service:
    build:
      context: ./auth-service
      dockerfile: Dockerfile
    container_name: cme-auth-service
    environment:
      - DB_HOST=auth-db
      - DB_PORT=5432
      - DB_NAME=auth_db
      - DB_USER=cme
      - DB_PASSWORD=${DB_PASSWORD}
      - JWT_SECRET=${JWT_SECRET}
    depends_on:
      - auth-db
    networks:
      - cme-network

  company-service:
    build:
      context: ./company-service
      dockerfile: Dockerfile
    container_name: cme-company-service
    environment:
      - DB_HOST=company-db
      - DB_PORT=5432
      - DB_NAME=company_db
      - DB_USER=cme
      - DB_PASSWORD=${DB_PASSWORD}
      - NATS_URL=nats://nats:4222
    depends_on:
      - company-db
      - nats
    networks:
      - cme-network

  material-service:
    build:
      context: ./material-service
      dockerfile: Dockerfile
    container_name: cme-material-service
    environment:
      - DB_HOST=material-db
      - DB_PORT=5432
      - DB_NAME=material_db
      - DB_USER=cme
      - DB_PASSWORD=${DB_PASSWORD}
    depends_on:
      - material-db
    networks:
      - cme-network

  order-service:
    build:
      context: ./order-service
      dockerfile: Dockerfile
    container_name: cme-order-service
    environment:
      - DB_HOST=order-db
      - DB_PORT=5432
      - DB_NAME=order_db
      - DB_USER=cme
      - DB_PASSWORD=${DB_PASSWORD}
      - NATS_URL=nats://nats:4222
      - MATERIAL_SERVICE_ADDR=material-service:50053
    depends_on:
      - order-db
      - nats
      - material-service
    networks:
      - cme-network

  review-service:
    build:
      context: ./review-service
      dockerfile: Dockerfile
    container_name: cme-review-service
    environment:
      - DB_HOST=review-db
      - DB_PORT=5432
      - DB_NAME=review_db
      - DB_USER=cme
      - DB_PASSWORD=${DB_PASSWORD}
      - NATS_URL=nats://nats:4222
    depends_on:
      - review-db
      - nats
    networks:
      - cme-network

  notification-service:
    build:
      context: ./notification-service
      dockerfile: Dockerfile
    container_name: cme-notification-service
    environment:
      - DB_HOST=notif-db
      - DB_PORT=5432
      - DB_NAME=notif_db
      - DB_USER=cme
      - DB_PASSWORD=${DB_PASSWORD}
      - NATS_URL=nats://nats:4222
    depends_on:
      - notif-db
      - nats
    networks:
      - cme-network

  # ========== Frontend ==========
  frontend:
    build:
      context: ./ui
      dockerfile: Dockerfile
    container_name: cme-frontend
    ports:
      - "3000:80"
    depends_on:
      - api-gateway
    networks:
      - cme-network

volumes:
  auth-db-data:
  company-db-data:
  material-db-data:
  order-db-data:
  review-db-data:
  notif-db-data:

networks:
  cme-network:
    driver: bridge
```

---

## 8. Makefile

```makefile
.PHONY: proto build up down logs clean

# Generate gRPC code from proto files
proto:
	@echo "Generating gRPC code..."
	@for dir in auth company material order review notification; do \
		protoc --go_out=. --go-grpc_out=. proto/$$dir.proto; \
	done

# Build all services
build:
	@echo "Building all services..."
	docker-compose build

# Start all services
up:
	@echo "Starting all services..."
	docker-compose up -d

# Stop all services
down:
	@echo "Stopping all services..."
	docker-compose down

# View logs
logs:
	docker-compose logs -f

# View logs for specific service
logs-%:
	docker-compose logs -f $*

# Clean everything
clean:
	docker-compose down -v --rmi all

# Restart specific service
restart-%:
	docker-compose restart $*

# Run database migrations
migrate:
	@echo "Running migrations..."
	@for dir in auth company material order review notification; do \
		docker-compose exec $$dir-service /app/migrate; \
	done

# Development: start only databases and infrastructure
dev-infra:
	docker-compose up -d auth-db company-db material-db order-db review-db notif-db nats

# Development: run service locally
dev-%:
	cd $*-service && go run cmd/server/main.go
```

---

## 9. API Endpoints (API Gateway)

### Auth Routes
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/auth/register` | Đăng ký tài khoản |
| POST | `/api/auth/login` | Đăng nhập |
| GET | `/api/auth/me` | Lấy thông tin user hiện tại |
| PUT | `/api/auth/profile` | Cập nhật profile |

### Company Routes
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/companies` | Tạo doanh nghiệp |
| GET | `/api/companies/:id` | Lấy thông tin doanh nghiệp |
| GET | `/api/companies` | Danh sách doanh nghiệp |
| PUT | `/api/companies/:id` | Cập nhật doanh nghiệp |
| POST | `/api/companies/:id/approve` | Admin duyệt doanh nghiệp |
| POST | `/api/companies/:id/reject` | Admin từ chối doanh nghiệp |

### Material Routes
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/categories` | Danh sách danh mục |
| POST | `/api/categories` | Tạo danh mục (admin) |
| POST | `/api/listings` | Đăng nguồn cung |
| GET | `/api/listings` | Danh sách nguồn cung (marketplace) |
| GET | `/api/listings/:id` | Chi tiết nguồn cung |
| PUT | `/api/listings/:id` | Cập nhật nguồn cung |
| DELETE | `/api/listings/:id` | Xóa nguồn cung |
| POST | `/api/demands` | Đăng nhu cầu mua |
| GET | `/api/demands` | Danh sách nhu cầu mua |
| GET | `/api/demands/:id` | Chi tiết nhu cầu mua |

### Order Routes
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/offers` | Tạo đề nghị mua |
| GET | `/api/offers` | Danh sách đề nghị |
| GET | `/api/offers/:id` | Chi tiết đề nghị |
| POST | `/api/offers/:id/accept` | Chấp nhận đề nghị |
| POST | `/api/offers/:id/reject` | Từ chối đề nghị |
| POST | `/api/offers/:id/cancel` | Hủy đề nghị |
| GET | `/api/transactions` | Danh sách giao dịch |
| GET | `/api/transactions/:id` | Chi tiết giao dịch |
| POST | `/api/transactions/:id/status` | Cập nhật trạng thái giao dịch |

### Review Routes
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/reviews` | Tạo đánh giá |
| GET | `/api/reviews` | Danh sách đánh giá |
| GET | `/api/reviews/user/:id` | Đánh giá của user |

### Notification Routes
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/notifications` | Danh sách thông báo |
| PUT | `/api/notifications/:id/read` | Đánh dấu đã đọc |
| PUT | `/api/notifications/read-all` | Đọc tất cả |
| GET | `/api/notifications/unread-count` | Số thông báo chưa đọc |

---

## 10. Luồng nghiệp vụ chính

### 10.1 Buyer gửi Offer

```
React → POST /api/offers
  → API Gateway (validate JWT)
  → order-service.CreateOffer (gRPC)
  → material-service.GetListing (gRPC) - verify listing exists
  → Insert offer vào order_db
  → NATS publish "cme.orders.offer.created"
  → notification-service consume → tạo notification cho seller
  → Response 201
```

### 10.2 Seller chấp nhận Offer → Tạo Transaction

```
React → POST /api/offers/:id/accept
  → API Gateway
  → order-service.AcceptOffer (gRPC)
  → Update offer status = 'accepted'
  → Create transaction + transaction_event
  → NATS publish "cme.orders.transaction.created"
  → notification-service → notification cho buyer
  → Response 200
```

### 10.3 Buyer/Seller xác nhận hoàn tất giao dịch

```
React → POST /api/transactions/:id/status
  → API Gateway
  → order-service.UpdateTransactionStatus (gRPC)
  → Update transaction status
  → Insert transaction_event
  → If status == 'completed':
    → NATS publish "cme.orders.transaction.completed"
    → review-service → enable review for this transaction
    → notification-service → notify both parties
  → Response 200
```

### 10.4 Admin duyệt doanh nghiệp

```
React → POST /api/companies/:id/approve
  → API Gateway (require admin role)
  → company-service.ApproveCompany (gRPC)
  → Update company status = 'verified'
  → NATS publish "cme.companies.approved"
  → notification-service → notify company owner
  → Response 200
```

---

## 11. Các chức năng Bypass (Demo Mode)

Vì đây là project giữa kỳ, các chức năng sau được bypass:

| Chức năng | Cách bypass | Ghi chú |
|-----------|-------------|---------|
| **Thanh toán** | `payment_status = 'bypassed_demo'`, `payment_method = 'manual_offline'` | Không gọi payment gateway |
| **OTP SMS** | Dùng OTP giả lập `123456` hoặc skip | Không cần nhà cung cấp SMS |
| **Email xác thực** | Bỏ qua, auto-verify | Không cần SMTP |
| **Xác minh mã số thuế** | Admin duyệt thủ công | Không cần API cơ quan nhà nước |
| **Logistics** | Lưu địa chỉ, ghi chú | Không tích hợp đơn vị vận chuyển |
| **Hợp đồng điện tử** | Checkbox đồng ý điều khoản | Không cần e-signature |
| **Push notification** | Chỉ dùng in-app notification | Không cần FCM/APNs |
| **OAuth social login** | Ẩn nút Google/LinkedIn | Không cần OAuth app |

---

## 12. Ước lượng Effort

| Task | Thời gian | Chi tiết |
|------|-----------|----------|
| Setup project structure | 0.5 ngày | Tạo thư mục, go.mod, Dockerfile |
| Proto definitions | 0.5 ngày | Viết 6 file .proto, generate code |
| auth-service | 1 ngày | CRUD user, JWT, login/register |
| company-service | 0.5 ngày | CRUD company, approve/reject |
| material-service | 1 ngày | Categories, listings, demands, search |
| order-service | 1.5 ngày | Offers, transactions, timeline, NATS events |
| review-service | 0.5 ngày | CRUD reviews, rating calculation |
| notification-service | 0.5 ngày | CRUD notifications, NATS consumer |
| API Gateway | 1 ngày | HTTP handlers, middleware, gRPC clients |
| Docker compose | 0.5 ngày | Setup all containers, networking |
| Frontend integration | 1 ngày | Update React app gọi API thật |
| Testing & debugging | 1 ngày | Test toàn bộ flow |
| **Tổng** | **~9-10 ngày** | |

---

## 13. Thứ tự triển khai

1. **Phase 1: Foundation** (1 ngày)
   - Tạo project structure
   - Viết proto definitions
   - Generate gRPC code
   - Setup Docker Compose

2. **Phase 2: Core Services** (3 ngày)
   - auth-service
   - company-service
   - material-service

3. **Phase 3: Business Logic** (2 ngày)
   - order-service (quan trọng nhất)
   - NATS integration

4. **Phase 4: Supporting Services** (1 ngày)
   - review-service
   - notification-service

5. **Phase 5: Gateway & Integration** (2 ngày)
   - API Gateway
   - Frontend integration

6. **Phase 6: Testing & Polish** (1 ngày)
   - End-to-end testing
   - Bug fixes
   - Documentation

---

## 14. Tech Stack

| Component | Technology | Version |
|-----------|------------|---------|
| Backend | Go | 1.21+ |
| HTTP Framework | Gin | latest |
| gRPC | google.golang.org/grpc | latest |
| Proto | protobuf | v3 |
| Database | PostgreSQL | 15 |
| Message Broker | NATS JetStream | latest |
| Frontend | React + TypeScript | 18.2 |
| Build Tool | Vite | 5.x |
| Container | Docker + Docker Compose | latest |

---

## 15. Lưu ý quan trọng

1. **Đây là project giữa kỳ** - mục tiêu là demo luồng nghiệp vụ hoàn chỉnh, không cần production-ready.

2. **Thanh toán là bypass** - không tích hợp payment gateway thật. Giao dịch chỉ là trạng thái demo.

3. **Mỗi service có database riêng** - tuân theo nguyên tắc Database per Service trong microservices.

4. **gRPC cho internal communication** - tất cả communication giữa services qua gRPC, chỉ API Gateway expose REST API.

5. **NATS cho async events** - các event quan trọng (offer created, transaction completed) publish qua NATS để notification-service và review-service consume.

6. **Frontend giữ nguyên** - chỉ cần cập nhật base URL API trong React app để gọi API Gateway thay vì mock data.

---

*Tài liệu này được tạo cho project giữa kỳ môn Phân tích và Thiết kế Hệ thống.*
