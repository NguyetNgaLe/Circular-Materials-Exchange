# Chương 6 - Thử nghiệm và đánh giá

## 6.1. Kiểm tra kỹ thuật đã thực hiện

| Mã | Kiểm tra | Kết quả gần nhất |
|---|---|---|
| T00 | `go test ./...` cho API Gateway và 6 service | Đạt; hiện chưa có test case Go riêng |
| T01 | `npm run build` frontend | Đạt |
| T02 | PostgreSQL fresh bootstrap | Đạt: 6 database, 18 bảng |
| T03 | Chạy lại init script lần hai | Đạt |
| T04 | Backend `/health` | HTTP 200 |
| T05 | Backend `/ready` và 15 nhóm API đọc qua gRPC | HTTP 200 |
| T06 | Frontend `/` và reverse proxy `/api/listings` | HTTP 200 |
| T07 | Login thật qua Auth gRPC | Đạt, trả JWT và role đúng |
| T08 | Upload Gateway -> Material gRPC -> MinIO | Đạt; ảnh test tải được và đã dọn |
| T09 | NATS Order event -> Notification consumer | Đạt; tạo đúng một notification idempotent và đã dọn |

## 6.2. Kịch bản chức năng cần chạy khi nghiệm thu

### Kịch bản 1 - Đăng ký và duyệt doanh nghiệp

- Tiền điều kiện: email mới; có tài khoản Admin.
- Bước: đăng ký -> tạo Company -> Admin xem danh sách -> duyệt.
- Mong đợi: User được tạo; Company `pending -> verified`; `GET /auth/me` ghép đúng Company theo owner.

### Kịch bản 2 - Chặn doanh nghiệp chưa được duyệt

- Bước: dùng Business User chưa verified để đăng nguồn cung/gửi Offer.
- Mong đợi: API trả 403; UI hướng người dùng tới hồ sơ doanh nghiệp.

### Kịch bản 3 - Upload ảnh và đăng nguồn cung

- Bước: chọn ảnh hợp lệ dưới 5 MB -> upload -> tạo listing.
- Mong đợi: MinIO nhận file; SupplyListing `active`; marketplace hiển thị ảnh và thông tin.
- Kiểm tra lỗi: file quá 5 MB hoặc sai đuôi phải bị từ chối.

### Kịch bản 4 - Tìm nguồn cung và gửi Offer

- Bước: Buyer tìm listing -> nhập số lượng/giá -> xác nhận gửi.
- Mong đợi: Offer `pending`; escrow ledger `holding`; Seller nhận Notification.

### Kịch bản 5 - Seller chấp nhận Offer

- Bước: Seller mở Offer nhận được -> chấp nhận.
- Mong đợi: Offer `accepted`; Transaction `confirmed`; có TransactionEvent; Buyer nhận Notification.
- Kiểm tra: escrow holding được gắn `transaction_id`, không tạo thêm escrow trùng cho lần accept.

### Kịch bản 6 - Giao hàng và hoàn tất

- Bước: Seller xác nhận giao hàng -> Buyer xác nhận nhận hàng.
- Mong đợi: `confirmed -> in_progress -> completed`; timeline có đủ sự kiện.

### Kịch bản 7 - Release escrow và cập nhật ví

- Bước: hoàn tất Transaction hoặc Admin release escrow.
- Mong đợi: escrow `released`; ghi fee 2%; Seller wallet tăng; có wallet transaction.
- Chú ý: chỉ kiểm tra ledger, không khẳng định tiền thật được chuyển.

### Kịch bản 8 - Đánh giá

- Bước: Buyer mở Transaction completed -> tạo rating/comment.
- Mong đợi: Review được lưu và hiển thị trong danh sách/chi tiết Seller.

### Kịch bản 9 - Notification

- Bước: xem danh sách -> đánh dấu một -> đánh dấu tất cả.
- Mong đợi: unread count giảm đúng.

### Kịch bản 10 - Phân quyền Admin

- Bước: gọi `/api/admin/*` bằng Business User và Admin.
- Mong đợi: Business User nhận 403; Admin truy cập được.

## 6.3. Kịch bản phải ghi là chưa đạt

| Chức năng | Kết quả dự kiến hiện tại |
|---|---|
| Tạo Demand | Trả ID demo, không có bản ghi mới |
| Sửa/xóa Listing | Handler trả success nhưng chưa thay đổi database |
| Hủy Offer | Không có REST route |
| Ownership/state machine đầy đủ | Mới enforce Offer pending; chưa kiểm tra đầy đủ actor sở hữu |
| Atomicity Offer/Transaction/Escrow | Chưa có một SQL transaction bao trọn chuỗi cập nhật |
| Matching tự động | Không có |
| Thanh toán/logistics thật | Không có |

## 6.4. Đánh giá

Prototype đã thể hiện được luồng chính từ đăng ký doanh nghiệp đến đăng nguồn cung, Offer, Transaction, timeline, escrow ledger, wallet và Review. Điểm mạnh là có giao diện hoàn chỉnh, dữ liệu thật trong PostgreSQL và deployment hai server.

Gateway đã gọi service qua gRPC, Auth phát JWT, bootstrap/migration đã đồng nhất và NATS consumer đã hoạt động. Hạn chế chính còn lại là một số stub, nhánh tương thích demo token, state/ownership validation chưa chặt, atomicity và thiếu test tích hợp tự động trong CI. Đây là cơ sở cho phần hướng phát triển ở Chương 7.
