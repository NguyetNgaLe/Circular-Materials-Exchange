# Đặc tả các Use Case trọng yếu

## UC004 - Tạo hồ sơ doanh nghiệp

| Thuộc tính | Nội dung |
|---|---|
| Actor chính | Business User |
| Tiền điều kiện | Người dùng đã đăng nhập; chưa có hồ sơ doanh nghiệp phù hợp |
| Kích hoạt | Người dùng gửi biểu mẫu hồ sơ doanh nghiệp |
| Luồng chính | 1. Nhập tên, mã số thuế, địa chỉ, thành phố và mô tả. 2. Frontend gửi `POST /api/companies`. 3. Gateway lấy `user_id` từ middleware. 4. Gateway gọi Company Service qua gRPC. 5. Service tạo Company `pending`. 6. Trả hồ sơ vừa tạo. |
| Luồng thay thế | Dữ liệu thiếu: trả 400. Mã số thuế trùng/lỗi DB: trả lỗi và không tạo hồ sơ. |
| Hậu điều kiện | Có Company `pending` gắn logic với User |
| Trạng thái code | Đã chạy |

## UC006 - Duyệt doanh nghiệp

| Thuộc tính | Nội dung |
|---|---|
| Actor chính | Admin |
| Tiền điều kiện | Admin đã đăng nhập; Company tồn tại |
| Kích hoạt | Admin chọn “Duyệt” |
| Luồng chính | 1. Frontend gửi `POST /api/companies/{id}/approve`. 2. Middleware kiểm tra role Admin qua Auth Service. 3. Gateway gọi Company Service cập nhật `verified`. 4. UI tải lại danh sách. |
| Ngoại lệ | Không phải Admin: 403. Company không tồn tại: code hiện chưa trả lỗi chặt ở mọi nhánh. |
| Hậu điều kiện | Company có trạng thái `verified` và có thể đăng nguồn cung/gửi Offer |
| Trạng thái code | Đã chạy; validation tồn tại bản ghi cần hoàn thiện |

## UC012 - Đăng nguồn cung

| Thuộc tính | Nội dung |
|---|---|
| Actor chính | Seller |
| Tiền điều kiện | Đã đăng nhập; Company của Seller có trạng thái `verified` |
| Kích hoạt | Seller hoàn tất biểu mẫu đăng vật liệu |
| Luồng chính | 1. Chọn và upload ảnh qua `POST /api/upload`. 2. Gateway kiểm tra loại file và giới hạn 5 MB. 3. Gateway gửi byte ảnh qua Material gRPC; Material Service PUT vào MinIO. 4. Seller nhập thông tin vật liệu. 5. Frontend gửi `POST /api/listings`. 6. Gateway kiểm tra Company qua gRPC. 7. Material Service ghi SupplyListing `active`. |
| Luồng thay thế | Company chưa duyệt: 403. File sai định dạng/quá lớn: 400. |
| Hậu điều kiện | Nguồn cung xuất hiện trên marketplace |
| Trạng thái code | Đã chạy; bootstrap và migration đã có `images`/`image_url` |

## UC017 - Gửi Offer

| Thuộc tính | Nội dung |
|---|---|
| Actor chính | Buyer |
| Tiền điều kiện | Buyer đăng nhập; Company `verified`; SupplyListing tồn tại trên UI |
| Kích hoạt | Buyer nhập số lượng, giá và xác nhận gửi |
| Luồng chính | 1. Frontend gửi `POST /api/offers`. 2. Gateway kiểm tra trạng thái Company qua gRPC. 3. Order Service ghi Offer `pending`. 4. Tính tổng tiền và phí 2%. 5. Ghi escrow ledger `holding`. 6. Publish NATS và tạo Notification gRPC idempotent cho Seller. 7. Trả Offer. |
| Ngoại lệ | Company chưa duyệt: 403. Dữ liệu thiếu: 400. Lỗi DB: 500. |
| Hậu điều kiện | Seller nhìn thấy Offer; có escrow ledger prototype |
| Trạng thái code | Đã chạy; chưa xác minh Listing ở backend và không có thanh toán thật |

## UC020 - Chấp nhận Offer

| Thuộc tính | Nội dung |
|---|---|
| Actor chính | Seller |
| Tiền điều kiện | Seller đăng nhập; Offer tồn tại và theo quy tắc phải ở `pending` |
| Kích hoạt | Seller chọn “Chấp nhận” |
| Luồng chính | 1. Frontend gửi `POST /api/offers/{id}/accept`. 2. Gateway gọi Order Service. 3. Service kiểm tra Offer `pending` và cập nhật `accepted`. 4. Tạo Transaction `confirmed`. 5. Tạo TransactionEvent. 6. Gắn escrow holding với Transaction. 7. Publish NATS và tạo Notification idempotent cho Buyer. |
| Ngoại lệ | Offer không tồn tại/lỗi DB: trả lỗi hoặc kết quả fallback tùy nhánh hiện tại. |
| Hậu điều kiện | Transaction mới được tạo và hai bên có thể theo dõi |
| Trạng thái code | Đã chạy; Order Service enforce `pending`, quyền Seller sở hữu Offer còn thiếu |

## UC024 - Seller xác nhận đã giao hàng

| Thuộc tính | Nội dung |
|---|---|
| Actor chính | Seller |
| Tiền điều kiện | Transaction `confirmed`; user là Seller trên giao dịch |
| Kích hoạt | Seller chọn “Xác nhận đã giao hàng” |
| Luồng chính | 1. Frontend gửi trạng thái `in_progress`. 2. Gateway gọi Order Service. 3. Service cập nhật Transaction và tạo TransactionEvent. 4. Publish event Notification. 5. Trả Transaction mới. |
| Hậu điều kiện | Transaction ở `in_progress`; Buyer có thể xác nhận nhận hàng |
| Trạng thái code | UI kiểm tra vai trò/trạng thái; backend chưa tự enforce đầy đủ |

## UC025 - Buyer xác nhận nhận hàng và hoàn tất

| Thuộc tính | Nội dung |
|---|---|
| Actor chính | Buyer |
| Tiền điều kiện | Transaction `in_progress`; user là Buyer |
| Kích hoạt | Buyer chọn “Đã nhận hàng và hoàn tất” |
| Luồng chính | 1. Frontend gửi trạng thái `completed`. 2. Gateway gọi Order Service cập nhật Transaction và Event. 3. Service tìm escrow `holding`. 4. Chuyển escrow sang `released`. 5. Ghi platform fee 2%. 6. Cộng Seller wallet. 7. Ghi wallet transaction. 8. Publish event Notification. |
| Ngoại lệ | Không tìm thấy escrow: giao dịch vẫn có thể được cập nhật nhưng ledger không được release. |
| Hậu điều kiện | Transaction `completed`; dữ liệu ledger được cập nhật; Buyer có thể Review |
| Trạng thái code | Đã chạy ở mức prototype, không chuyển tiền thật |

## UC026 - Tạo Review

| Thuộc tính | Nội dung |
|---|---|
| Actor chính | Buyer/Seller |
| Tiền điều kiện | Người dùng đăng nhập; về nghiệp vụ Transaction phải hoàn tất |
| Kích hoạt | Người dùng gửi rating và comment |
| Luồng chính | 1. Frontend gửi `POST /api/reviews`. 2. Gateway lấy reviewer từ token. 3. Tra tên reviewee qua Auth gRPC. 4. Gọi Review Service ghi Review. 5. Trả Review. |
| Ngoại lệ | Rating/dữ liệu thiếu: 400 hoặc lỗi DB. |
| Hậu điều kiện | Review xuất hiện trong danh sách và được dùng tính average |
| Trạng thái code | Đã chạy; backend chưa kiểm tra Transaction completed và chưa chống đánh giá trùng |

## UC033 - Xem và giải phóng escrow

| Thuộc tính | Nội dung |
|---|---|
| Actor chính | Admin |
| Tiền điều kiện | Admin đăng nhập; escrow `holding` tồn tại |
| Kích hoạt | Admin chọn “Giải ngân” |
| Luồng chính | 1. Xem danh sách qua `GET /api/admin/escrow`. 2. Gửi `POST /api/admin/escrow/{id}/release`. 3. Gateway gọi Order Service cập nhật escrow `released`. 4. Service cộng Seller wallet và ghi lịch sử trong transaction. |
| Ngoại lệ | Escrow không tồn tại/đã release: 404. |
| Hậu điều kiện | Ledger cho biết escrow đã release |
| Trạng thái code | Đã chạy ở mức mô phỏng; không có giao dịch ngân hàng |
