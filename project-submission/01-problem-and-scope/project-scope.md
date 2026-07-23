# Phạm vi project

## Đã triển khai và dùng được trong demo

- Đăng ký, đăng nhập và khôi phục phiên ở frontend.
- Tạo, xem, duyệt hoặc từ chối hồ sơ doanh nghiệp.
- Xem danh mục, nguồn cung, nhu cầu và chi tiết vật liệu.
- Upload ảnh vật liệu lên MinIO.
- Đăng nguồn cung sau khi doanh nghiệp được duyệt.
- Gửi, xem, chấp nhận hoặc từ chối Offer.
- Tạo Transaction, TransactionEvent và xem timeline.
- Seller xác nhận giao hàng; Buyer xác nhận nhận hàng và hoàn tất.
- Tạo/xem Review.
- Xem và đánh dấu Notification.
- Finance, escrow, Seller wallet và withdrawal API ở mức prototype.
- Dashboard và các màn Admin cơ bản.

## Có trong hệ thống nhưng chưa hoàn chỉnh

- Login kiểm tra mật khẩu thật nhưng phát demo token thay vì JWT thật.
- Tìm kiếm/lọc còn đơn giản.
- Tạo Demand chưa ghi database.
- Sửa/xóa SupplyListing đang là stub.
- State transition chưa được backend kiểm tra chặt.
- Một số màn Admin category/listing/report/export mới ở mức UI/mock.
- Luồng REST chính đi qua Gateway và 6 gRPC service; Gateway không truy cập database trực tiếp.
- Schema ảnh trên server chưa được chuẩn hóa đầy đủ bằng migration.

## Ngoài phạm vi MVP

- Thuật toán tự động matching/recommendation.
- Thanh toán, escrow hoặc chuyển khoản ngân hàng thật.
- Tích hợp logistics và tracking.
- Hợp đồng điện tử, chữ ký số.
- Xác minh pháp lý tự động.
- Kiểm định chất lượng vật liệu bởi bên thứ ba.
- Xử lý tranh chấp hoàn chỉnh.
- Seller gửi báo giá trực tiếp cho Demand.

Các thuật ngữ “thanh toán”, “escrow”, “ví” và “rút tiền” trong giao diện chỉ mô tả dữ liệu mô phỏng phục vụ project, không khẳng định có dòng tiền thật.
