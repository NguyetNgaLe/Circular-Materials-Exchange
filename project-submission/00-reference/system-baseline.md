# System Baseline

> Baseline được đối chiếu ngày 24/07/2026 sau đợt refactor microservice, kiểm tra route/API, schema/migration và deployment trên hai máy chủ demo. Khi code thay đổi, phải cập nhật file này trước khi sửa báo cáo hoặc sơ đồ.

## 1. Bài toán

Circular Materials Exchange là nền tảng B2B giúp doanh nghiệp công bố vật liệu dư thừa, tìm kiếm nguồn vật liệu tuần hoàn, trao đổi đề nghị mua và theo dõi giao dịch. Ví dụ nghiệp vụ gồm nhựa PP, pallet gỗ, bã cà phê, vải vụn, giấy, kim loại và thủy tinh.

Mục tiêu dài hạn là tự động kết nối cung-cầu. Phiên bản hiện tại **chưa có thuật toán matching/recommendation**; việc kết nối được thực hiện bằng marketplace, tìm kiếm/lọc và Offer.

## 2. Actor

| Actor | Vai trò trong code hiện tại |
|---|---|
| Guest | Xem nguồn cung, nhu cầu, chi tiết vật liệu; đăng ký và đăng nhập |
| Business User | Quản lý phiên, hồ sơ doanh nghiệp và sử dụng dashboard |
| Buyer | Tìm nguồn cung, gửi Offer, theo dõi và hoàn tất giao dịch, đánh giá Seller |
| Seller | Đăng nguồn cung, xử lý Offer, xác nhận giao hàng, xem ví và đánh giá Buyer |
| Admin | Duyệt doanh nghiệp; xem dashboard, tài chính, escrow và thực hiện giải ngân |

Buyer và Seller là vai trò nghiệp vụ của Business User theo từng giao dịch, không phải hai loại tài khoản cố định.

## 3. Thành phần triển khai

| Thành phần | Công nghệ | Trạng thái và trách nhiệm hiện tại |
|---|---|---|
| React SPA | React, TypeScript, Vite | Giao diện Guest, Business User và Admin |
| Frontend Web Server | Nginx | Phục vụ bản build SPA trên máy frontend, chuyển tiếp `/api/` |
| API Gateway | Go, Gin | Điểm vào REST, bearer-token middleware, chuyển đổi REST/JSON sang gRPC; không truy cập SQL |
| Auth Service | Go, gRPC, `auth_db` | Có đầy đủ Handler-Service-Repository và đang chạy |
| Company Service | Go, gRPC, `company_db` | Có đầy đủ Handler-Service-Repository và đang chạy |
| Material Service | Go, gRPC, `material_db` | Có đầy đủ Handler-Service-Repository và đang chạy |
| Order Service | Go, gRPC, `order_db` | Sở hữu Offer/Transaction/finance/escrow/wallet và publish NATS |
| Review Service | Go, gRPC, `review_db` | Có đầy đủ Handler-Service-Repository và đang chạy |
| Notification Service | Go, gRPC, NATS, `notif_db` | Sở hữu Notification, nhận lệnh gRPC và consume event Order |
| PostgreSQL 15 | Một container, 6 database | Lưu 18 bảng nghiệp vụ |
| NATS | NATS JetStream | Order Service publish event; Notification Service subscribe theo queue group |
| MinIO | Object storage | Material Service nhận byte ảnh qua gRPC và upload bằng HTTP PUT |

Frontend và backend được triển khai trên hai máy chủ riêng. Backend dùng Docker Compose; frontend là bản build Vite do Nginx phục vụ và chưa phải service trong Compose.

## 4. Kiến trúc logic và kiến trúc as-built

### 4.1. Luồng đồng bộ

`React SPA -> REST API Gateway -> gRPC microservices -> database do service sở hữu`

Mỗi microservice có database logic riêng theo nguyên tắc Database per Service. Sáu database cùng nằm trong một PostgreSQL container chỉ là quyết định triển khai cho project.

### 4.2. Kiến trúc as-built của luồng web

API Gateway dùng client gRPC sinh từ các contract trong `proto/` để gọi đúng service sở hữu dữ liệu:

| Nhóm REST tại Gateway | Service đích |
|---|---|
| Auth và bearer-token middleware | Auth Service |
| Company và kiểm tra hồ sơ doanh nghiệp | Company Service |
| Category, Listing, Demand và Upload | Material Service |
| Offer, Transaction, Finance, Escrow, Wallet, Withdrawal | Order Service |
| Review | Review Service, tra tên qua Auth Service |
| Notification | Notification Service |

Gateway không chứa `database/sql`, không dùng driver PostgreSQL và chạy cùng `cme-network`. Mỗi service chỉ truy cập database logic thuộc miền của mình.

### 4.3. Luồng bất đồng bộ

Order Service publish các subject `cme.orders.*` qua NATS. Notification Service subscribe các event tạo Offer, tạo Transaction, giao hàng và hoàn tất; `reference_id` cùng unique index bảo đảm đường event và đường gRPC đồng bộ không tạo thông báo trùng.

## 5. Dữ liệu hiện tại

Hệ thống có 6 database logic và 18 bảng:

| Database | Bảng |
|---|---|
| `auth_db` | `users` |
| `company_db` | `companies` |
| `material_db` | `categories`, `supply_listings`, `demand_listings` |
| `order_db` | `offers`, `transactions`, `transaction_events`, `platform_fees`, `platform_wallet`, `wallet_transactions`, `company_wallet`, `escrow_transactions`, `seller_wallet`, `seller_wallet_transactions`, `withdrawal_requests` |
| `review_db` | `reviews` |
| `notif_db` | `notifications` |

Các tham chiếu qua database như `buyer_id`, `seller_id`, `company_id`, `listing_id` và `reviewee_id` là ID logic, không phải khóa ngoại vật lý.

### Quản lý schema

- `scripts/init-databases.sql` đã được kiểm tra trên PostgreSQL volume mới và tạo đủ 6 database/18 bảng.
- `image_url`, `images` và `notifications.reference_id` đã nằm trong bootstrap/migration.
- `scripts/migrate-existing.sql` nâng cấp idempotent database volume đang chạy.
- Repository Material đọc được cả `specs` dạng `text` của dữ liệu demo và JSON text sinh từ service.

## 6. Trạng thái nghiệp vụ

### Company

`pending -> verified` hoặc `pending -> rejected`.

### Offer

Quy tắc mục tiêu: `pending -> accepted` hoặc `pending -> rejected`.

Order Service đã chặn accept/reject khi Offer không còn `pending`; kiểm tra actor sở hữu Offer chưa đầy đủ. Route hủy Offer chưa được expose.

### Transaction

Luồng UI đang dùng:

`confirmed -> in_progress -> completed`

- Seller chuyển `confirmed -> in_progress` khi xác nhận đã giao hàng.
- Buyer chuyển `in_progress -> completed` khi xác nhận đã nhận hàng.

Endpoint cập nhật trạng thái hiện nhận chuỗi trạng thái từ client và chưa enforce state machine. Các trạng thái `buyer_confirmed`, `seller_confirmed`, `cancelled`, `disputed` thuộc thiết kế mục tiêu hoặc service layer, chưa phải luồng UI chính.

### Escrow và Withdrawal

- Escrow: `holding -> released`.
- Withdrawal: `pending -> completed` hoặc `pending -> rejected`.

Escrow, ví và giải ngân chỉ là **sổ ghi nhận trong prototype**. Hệ thống không kết nối ngân hàng, không nạp tiền thật và không chuyển tiền thật.

## 7. Luồng chính đang chạy

1. Người dùng đăng ký hoặc đăng nhập.
2. Business User tạo hồ sơ doanh nghiệp ở trạng thái `pending`.
3. Admin duyệt hoặc từ chối doanh nghiệp.
4. Doanh nghiệp `verified` tải ảnh và đăng nguồn cung.
5. Buyer tìm kiếm nguồn cung, xem chi tiết và gửi Offer.
6. Khi tạo Offer, Order Service ghi escrow prototype, publish event và Gateway tạo Notification đồng bộ có idempotency.
7. Seller chấp nhận Offer; Order Service tạo Transaction, TransactionEvent và gắn escrow đang giữ với giao dịch.
8. Seller xác nhận giao hàng (`in_progress`).
9. Buyer xác nhận nhận hàng (`completed`).
10. Order Service giải phóng escrow, ghi phí 2%, cộng ví Seller và publish Notification event.
11. Buyer có thể đánh giá Seller.
12. Admin xem dữ liệu finance/escrow và có thể giải ngân thủ công.

## 8. Giới hạn và chức năng chưa hoàn chỉnh

- Chưa có thuật toán tự động matching/recommendation.
- `POST /api/demands` đang trả ID demo, chưa ghi `demand_listings`.
- `PUT /api/listings/:id` đã kiểm tra Seller sở hữu (hoặc Admin), hỗ trợ sửa đầy đủ
  thông tin và ẩn/hiện nguồn cung qua Material Service gRPC. `GET /api/my/listings`
  trả cả nguồn cung đang ẩn của Seller. `DELETE /api/listings/:id` kiểm tra cùng quyền
  sở hữu rồi xóa dữ liệu trong `material_db`.
- Chưa có route cập nhật hồ sơ người dùng, cập nhật doanh nghiệp hoặc hủy Offer.
- Chưa có luồng Seller gửi báo giá trực tiếp cho Demand.
- Một số trang Admin (category, listing, report, export) chủ yếu là UI/mock và chưa có REST CRUD tương ứng.
- Login/Register phát JWT HS256 thật; middleware vẫn nhận `demo-token-*` cũ để không làm gián đoạn phiên demo đã lưu trong trình duyệt.
- Notification có cả đường gRPC đồng bộ và NATS consumer; unique reference ngăn ghi trùng.
- Chuỗi cập nhật Offer -> Transaction -> Escrow chưa được gom thành một SQL transaction duy nhất.
- Thanh toán, escrow, rút tiền và giải ngân không kết nối ngân hàng.
- Logistics, tracking, hợp đồng điện tử, kiểm định và giải quyết tranh chấp chưa được tích hợp.
- Chưa có TLS/mTLS cho gRPC nội bộ, rate-limit phân tán, tracing tập trung hoặc secret manager.

## 9. Thứ tự ưu tiên nguồn thông tin

1. Route, HTTP handler, gRPC contract/client, frontend API call, migrations và `docker-compose.yml` hiện tại.
2. Schema đang chạy dùng để kiểm tra tương thích; schema chuẩn được quản lý bằng bootstrap và migration.
3. `00-reference/as-built-feature-matrix.md`.
4. Proto và service layer để mô tả kiến trúc mục tiêu.
5. Tài liệu kế hoạch cũ chỉ dùng tham khảo; chi tiết khác code phải ghi là mục tiêu hoặc loại bỏ.
