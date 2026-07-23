# Chương 4 - Kiến trúc và kỹ thuật thiết kế

## 4.1. Kiến trúc

### 4.1.1. Kiến trúc logic

Kiến trúc hiện tại gồm React SPA, REST API Gateway và sáu gRPC microservice: Auth, Company, Material, Order, Review và Notification. Mỗi service sở hữu một database logic riêng. NATS hỗ trợ giao tiếp sự kiện và MinIO lưu ảnh.

D11 mô tả đường đi as-built: Gateway chuyển REST/JSON thành gRPC, từng service xử lý nghiệp vụ và chỉ repository của service truy cập database thuộc miền đó. Gateway không còn kết nối SQL.

### 4.1.2. Kiến trúc triển khai

Hệ thống chạy trên hai server:

- Server frontend: Nginx phục vụ bản build React/Vite và reverse proxy `/api/`.
- Server backend: Docker Compose chạy Gateway, sáu service, PostgreSQL, NATS và MinIO.

PostgreSQL là một container vật lý nhưng chứa sáu database logic và 18 bảng. D12 mô tả node/component; D24 trình bày mô hình Docker/server dùng trong demo.

## 4.2. Các kỹ thuật thiết kế

### Lớp giao diện

Frontend tổ chức theo page/component, route và API client. Trạng thái đăng nhập được giữ ở store phía client. Nginx dùng SPA fallback để các client-side route vẫn mở trực tiếp được.

### Web service

Gateway dùng REST/JSON cho frontend, bearer-token middleware cho route bảo vệ và generated gRPC client cho giao tiếp nội bộ. Các service được tổ chức theo Handler-Service-Repository. Auth Service xác minh JWT cho middleware; các handler có thể orchestration nhiều service nhưng không sở hữu dữ liệu của service khác.

### Lớp trường cửu

Repository dùng PostgreSQL. Database per Service được áp dụng ở mức logic; các tham chiếu xuyên database lưu bằng UUID nhưng không có foreign key vật lý. MinIO giữ file ảnh, còn database chỉ giữ URL.

### Giao tiếp sự kiện và tài chính

Order Service publish event NATS cho Offer/Transaction. Notification Service consume theo queue group và dùng `reference_id` idempotent; Gateway vẫn có đường tạo Notification gRPC đồng bộ để giữ phản hồi UI ổn định. Escrow, fee, wallet và withdrawal là ledger prototype, chưa tích hợp payment gateway/ngân hàng.

## 4.3. Thiết kế ca sử dụng

| Nhóm thiết kế | Sơ đồ | Ca sử dụng tiêu biểu |
|---|---|---|
| Dữ liệu | D13 | Toàn bộ miền dữ liệu, 6 database/18 bảng |
| Design Class | D14-D16 | Auth/Company, Material/Upload, Order và các service hỗ trợ |
| Sequence | D17 | UC004 tạo Company, UC006 duyệt Company |
| Sequence | D18 | UC012 upload ảnh và đăng nguồn cung |
| Sequence | D19 | UC017 gửi Offer |
| Sequence | D20 | UC020 chấp nhận Offer |
| Sequence | D21 | UC024/UC025 hoàn tất, UC026 đánh giá |
| State Machine | D22-D23 | Offer và Transaction |

Các Sequence Diagram dùng đường đi as-built REST -> gRPC -> Service -> Repository. Design Class thể hiện ranh giới Gateway, service và repository sở hữu dữ liệu.
