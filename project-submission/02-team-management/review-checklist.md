# Checklist nghiệm thu sơ đồ

## Nội dung chung

- [ ] Mã Dxx và tên hình khớp `03-report/diagram-index.md`.
- [ ] Thuật ngữ thống nhất: Business User, Buyer, Seller, Admin, SupplyListing, DemandListing, Offer, Transaction.
- [ ] Không ghi hệ thống đã tự động matching cung-cầu.
- [ ] Không mô tả thanh toán, logistics hoặc xác minh pháp lý như tích hợp thật.
- [ ] Buyer và Seller là vai trò của Business User.
- [ ] Có file source, SVG và PNG.
- [ ] Chữ đọc rõ khi hình được đặt vừa trang A4.

## Use Case và Activity

- [ ] Activity tập trung quy trình nghiệp vụ, không lẫn lời gọi kỹ thuật gRPC/SQL.
- [ ] Actor nằm ngoài biên hệ thống.
- [ ] Use case dùng động từ và thể hiện mục tiêu người dùng.
- [ ] Không thêm luồng Seller báo giá Demand như chức năng đã triển khai.

## Class và ERD

- [ ] Domain Class không chứa Handler/Service/Repository.
- [ ] Design Class có đủ Handler, Service, Repository và Entity.
- [ ] ERD thể hiện 6 database logic.
- [ ] Chỉ dùng FK vật lý trong cùng database; quan hệ qua service được ghi là logical reference.

## Architecture và Deployment

- [ ] Ghi rõ backend là kiến trúc microservice, không gọi là monolith hoặc kiến trúc nhiều lớp duy nhất.
- [ ] Có React SPA, API Gateway và 6 gRPC service.
- [ ] Thể hiện ranh giới sở hữu database của từng service; không vẽ service truy cập database của service khác.
- [ ] Docker hiện có một PostgreSQL container, NATS và MinIO.
- [ ] API Gateway chạy `network_mode: host` trong Compose hiện tại.
- [ ] Frontend không được vẽ như Docker service đang có.
- [ ] NATS chỉ được vẽ là Order Service publish; subscriber của Notification là thiết kế tương lai nếu xuất hiện.

## State và Sequence

- [ ] Offer chỉ chuyển từ `pending` sang `accepted`, `rejected` hoặc `cancelled`.
- [ ] Transaction state ghi rõ là quy tắc mục tiêu vì code chưa enforce transition.
- [ ] Sequence phân biệt REST từ frontend đến Gateway và gRPC từ Gateway đến service.
- [ ] Khi hoàn tất, Review được gửi qua API riêng; không khẳng định tự động tạo Review.
