# Mục tiêu project

## Mục tiêu tổng quát

Xây dựng prototype nền tảng trao đổi vật liệu tuần hoàn B2B, giúp doanh nghiệp công bố nguồn vật liệu dư, tìm kiếm cơ hội tái sử dụng và số hóa vòng đời đề nghị-giao dịch.

## Mục tiêu nghiệp vụ

1. Đăng và quản lý nguồn cung vật liệu dư thừa.
2. Công bố nhu cầu mua vật liệu.
3. Tìm kiếm, lọc và xem chi tiết nguồn cung/nhu cầu.
4. Gửi và xử lý Offer giữa Buyer và Seller.
5. Tạo Transaction khi Seller chấp nhận Offer.
6. Theo dõi trạng thái và TransactionEvent.
7. Ghi nhận escrow, phí và ví ở mức prototype.
8. Xác nhận hoàn tất, tạo Notification và Review.
9. Hỗ trợ Admin duyệt doanh nghiệp, theo dõi finance/escrow và dữ liệu chính.

## Mục tiêu kỹ thuật

- Xây dựng React SPA và API Gateway làm điểm vào REST.
- Tổ chức backend theo 6 miền microservice với proto/gRPC và database logic riêng.
- Đóng gói backend bằng Docker Compose.
- Sử dụng PostgreSQL, MinIO và NATS trong môi trường demo.
- Triển khai frontend và backend trên hai máy chủ độc lập.

## Giới hạn mục tiêu

Mục tiêu tự động matching cung-cầu là hướng phát triển, chưa có trong code hiện tại. Thanh toán, escrow, ví và rút tiền chỉ là ledger mô phỏng; không tích hợp ngân hàng. Logistics và giao nhận được thực hiện ngoài hệ thống.
