# Yêu cầu phi chức năng

> Các mục dưới đây là yêu cầu/tiêu chí đánh giá cho báo cáo, không phải khẳng định tất cả đã đạt.

| Mã | Nhóm | Yêu cầu | Hiện trạng/tiêu chí kiểm chứng |
|---|---|---|---|
| NFR01 | Bảo mật | Mật khẩu phải được băm, không lưu plaintext | Đã dùng bcrypt |
| NFR02 | Bảo mật | API ghi dữ liệu phải yêu cầu token; Admin route phải kiểm tra role | Có middleware; Auth Service phát JWT, vẫn tương thích demo token cũ |
| NFR03 | Bảo mật | Secret chỉ lấy từ biến môi trường | Đạt ở Compose/Gateway; chưa có secret manager |
| NFR04 | Bảo mật | Không commit password, `.env`, dump DB hoặc token thật | Bắt buộc kiểm tra trước khi nộp |
| NFR05 | Phân quyền | Buyer/Seller chỉ thao tác dữ liệu thuộc quyền | UI có kiểm tra một phần; backend cần tăng ownership validation |
| NFR06 | Hiệu năng | Trang danh sách chính phản hồi trong thời gian chấp nhận được ở dữ liệu demo | Đo bằng browser/API; chưa có benchmark chính thức |
| NFR07 | Khả dụng | Có health/readiness endpoint và container restart policy | `/health`, `/ready`; Docker `restart: on-failure` |
| NFR08 | Tin cậy dữ liệu | Giao dịch và thay đổi trạng thái phải có timeline | Có `transaction_events` |
| NFR09 | Nhất quán | Schema mới phải được dựng hoàn toàn bằng migration/init script | Bootstrap đã được test tạo đủ 6 database/18 bảng; có migration idempotent |
| NFR10 | Khả năng mở rộng | Miền nghiệp vụ được tách theo service/proto/database logic | Đã có 6 service; Gateway gọi bằng gRPC và không chứa SQL |
| NFR11 | Bảo trì | Code backend tách Handler-Service-Repository | Đã áp dụng trong 6 service; proto là nguồn contract chung |
| NFR12 | Triển khai | Backend có thể dựng bằng Docker Compose; frontend build bằng Vite | Đã build/deploy được trên hai máy chủ |
| NFR13 | Tương thích | Giao diện hoạt động trên trình duyệt hiện đại và responsive cơ bản | Cần kiểm thử Chrome/Edge, desktop/mobile |
| NFR14 | Dễ dùng | Trạng thái và hành động chính phải hiển thị rõ bằng tiếng Việt | Đã có UI; còn một số chuỗi pha tiếng Anh/không dấu |
| NFR15 | Quan sát | Có log lỗi và khả năng xem trạng thái container | Gin/Docker log có; chưa có metrics/tracing tập trung |
| NFR16 | Sao lưu | Dữ liệu PostgreSQL và MinIO cần volume, có phương án backup | Có `pg-data`, `minio-data`; chưa có job backup tự động |
| NFR17 | Quyền riêng tư | Không đưa thông tin ngân hàng thật vào dữ liệu demo/báo cáo | Withdrawal chỉ dùng dữ liệu demo |
| NFR18 | Trung thực chức năng | Không mô tả payment, escrow, logistics hoặc matching như tích hợp thật | Bắt buộc trong báo cáo và demo |

## Ưu tiên cải thiện trước khi coi là production

1. Loại bỏ nhánh tương thích demo token sau khi hết phiên demo cũ; bổ sung secret manager.
2. Enforce ownership và state transition tại backend.
3. Thêm transaction SQL cho các thao tác Offer-Transaction-Escrow-Wallet.
4. Bổ sung test contract/integration tự động.
5. Bổ sung mTLS, backup định kỳ, monitoring, tracing và audit log.
