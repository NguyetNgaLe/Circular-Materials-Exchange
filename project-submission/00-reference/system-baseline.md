# System Baseline

## 1. Bài toán

Circular Materials Exchange là nền tảng B2B giúp doanh nghiệp công bố vật liệu dư thừa, tìm kiếm nguồn vật liệu tuần hoàn, trao đổi đề nghị mua và theo dõi giao dịch. Các ví dụ thực tế gồm nhựa PP, pallet gỗ, bã cà phê, vải vụn, giấy và kim loại.

## 2. Actor

| Actor | Vai trò |
|---|---|
| Guest | Xem marketplace, nhu cầu và chi tiết nguồn cung; đăng ký, đăng nhập |
| Business User | Quản lý hồ sơ cá nhân và doanh nghiệp |
| Buyer | Vai trò của Business User khi tìm vật liệu, đăng nhu cầu, gửi Offer và xác nhận giao dịch |
| Seller | Vai trò của Business User khi đăng nguồn cung, xử lý Offer và xác nhận giao dịch |
| Admin | Duyệt doanh nghiệp và quản trị dữ liệu hệ thống |

Buyer và Seller không phải hai loại tài khoản cố định. Một Business User có thể đảm nhiệm cả hai vai trò tùy giao dịch.

## 3. Thành phần đang có trong code

Backend sử dụng **kiến trúc microservice**. API Gateway là điểm vào REST duy nhất cho frontend; nghiệp vụ được tách thành 6 service triển khai độc lập và Gateway gọi các service bằng gRPC. Mỗi service sở hữu mô hình dữ liệu/database logic của mình theo nguyên tắc Database per Service. Việc 6 database đang chạy trong cùng một PostgreSQL container chỉ là lựa chọn triển khai cho môi trường project, không phải dùng chung schema hay gộp nghiệp vụ thành monolith.

| Thành phần | Công nghệ | Trách nhiệm |
|---|---|---|
| React SPA | React, TypeScript, Vite | Giao diện Guest, Business User và Admin |
| API Gateway | Go, Gin | REST API, JWT middleware, chuyển tiếp gRPC |
| Auth Service | Go, gRPC, auth_db | Đăng ký, đăng nhập, hồ sơ người dùng |
| Company Service | Go, gRPC, company_db | Hồ sơ và phê duyệt doanh nghiệp |
| Material Service | Go, gRPC, material_db | Danh mục, nguồn cung, nhu cầu mua |
| Order Service | Go, gRPC, order_db | Offer, Transaction, TransactionEvent |
| Review Service | Go, gRPC, review_db | Đánh giá và điểm trung bình |
| Notification Service | Go, gRPC, notif_db | Danh sách và trạng thái đã đọc của thông báo |
| PostgreSQL | PostgreSQL 15 | Một container, chứa 6 database logic |
| NATS | NATS JetStream image | Order Service publish sự kiện giao dịch |
| MinIO | MinIO | Object storage trong Docker Compose |

### Ranh giới microservice

| Microservice | Database sở hữu | Không được truy cập trực tiếp |
|---|---|---|
| Auth Service | `auth_db` | Database của 5 service còn lại |
| Company Service | `company_db` | `auth_db`, `material_db`, `order_db`, `review_db`, `notif_db` |
| Material Service | `material_db` | Database của các service khác |
| Order Service | `order_db` | Database của các service khác |
| Review Service | `review_db` | Database của các service khác |
| Notification Service | `notif_db` | Database của các service khác |

Các quan hệ xuyên service được lưu bằng ID logic hoặc trao đổi qua API/gRPC/event, không tạo join hay foreign key xuyên database.

## 4. Thực thể nghiệp vụ

`User`, `Company`, `Category`, `SupplyListing`, `DemandListing`, `Offer`, `Transaction`, `TransactionEvent`, `Review`, `Notification`.

Các ID tham chiếu giữa hai database khác nhau là quan hệ logic, không phải khóa ngoại vật lý. Ví dụ `seller_id`, `buyer_id`, `company_id` và `transaction_id` trong Review Service.

## 5. Trạng thái

### Company

`pending -> verified` hoặc `pending -> rejected`.

### Offer

`pending -> accepted`, `pending -> rejected` hoặc `pending -> cancelled`.

### Transaction

Thiết kế mục tiêu: `confirmed -> in_progress -> buyer_confirmed/seller_confirmed -> completed`, kèm nhánh `cancelled` hoặc `disputed`.

Code hiện cho phép client gửi chuỗi trạng thái bất kỳ và chưa kiểm tra bảng chuyển trạng thái. Sơ đồ trạng thái thể hiện quy tắc nghiệp vụ mục tiêu; báo cáo cần ghi đây là validation cần hoàn thiện.

## 6. Luồng chính đã triển khai

1. Đăng ký/đăng nhập và nhận JWT.
2. Tạo doanh nghiệp với trạng thái `pending`.
3. Admin duyệt hoặc từ chối doanh nghiệp.
4. Seller đăng nguồn cung.
5. Buyer tìm kiếm nguồn cung và gửi Offer.
6. Seller chấp nhận hoặc từ chối Offer.
7. Khi chấp nhận, Order Service tạo Transaction và TransactionEvent.
8. Hai bên cập nhật trạng thái và xem timeline.
9. Khi hoàn tất, người dùng gửi Review.
10. Người dùng xem và đánh dấu Notification đã đọc.

## 7. Giới hạn phải trình bày trung thực

- Chưa có thuật toán tự động ghép cặp cung-cầu. MVP chỉ hỗ trợ công bố, tìm kiếm, lọc và gửi đề nghị.
- Luồng Demand mới có tạo, xem và liệt kê; chưa có API seller gửi báo giá trực tiếp cho Demand.
- Order Service publish `cme.orders.transaction.created` và `cme.orders.transaction.completed`, nhưng Notification Service chưa có NATS subscriber trong code.
- Thanh toán và logistics được thực hiện ngoài hệ thống. Không mô tả như tích hợp thanh toán thật.
- Docker Compose hiện dùng một PostgreSQL container cho 6 database, không phải 6 PostgreSQL container.
- Frontend chưa được khai báo thành service trong Docker Compose.
- `ARCHITECTURE.md` ở thư mục gốc thuộc dự án VTrue/VPoint khác và phải loại khỏi bài nộp.

## 8. Thứ tự ưu tiên nguồn thông tin

1. Code, migrations, proto, route và `docker-compose.yml` hiện tại.
2. `USECASE_GAP_ANALYSIS.md` cho phạm vi MVP.
3. `MICROSERVICES_PLAN.md` chỉ dùng làm tài liệu định hướng; chi tiết khác code phải bỏ.
4. Không sử dụng `ARCHITECTURE.md` cho Circular Materials Exchange.
