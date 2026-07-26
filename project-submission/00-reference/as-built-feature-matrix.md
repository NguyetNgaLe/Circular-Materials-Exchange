# Ma trận chức năng as-built

> Trạng thái được đối chiếu với frontend route, API Gateway route/handler và database ngày 24/07/2026.

## Quy ước

- **Đã chạy:** có UI/API và thao tác dữ liệu thực.
- **Một phần:** có UI hoặc API nhưng còn stub, bypass hoặc thiếu validation.
- **Chưa có:** chỉ là yêu cầu/hướng phát triển.

| ID | Nhóm | Chức năng | Trạng thái | Bằng chứng/Ghi chú |
|---|---|---|---|---|
| F01 | Auth | Đăng ký | Đã chạy | `POST /api/auth/register`, ghi `auth_db.users` |
| F02 | Auth | Đăng nhập | Đã chạy | Auth Service kiểm tra bcrypt và phát JWT; middleware giữ tương thích token demo cũ |
| F03 | Auth | Xem người dùng hiện tại | Đã chạy | `GET /api/auth/me` |
| F04 | Auth | Cập nhật hồ sơ | Chưa expose | Có trong proto/service, không có REST route/frontend |
| F05 | Company | Tạo hồ sơ doanh nghiệp | Đã chạy | Gateway gọi Company Service; `GET /auth/me` ghép Company theo `owner_id` |
| F06 | Company | Xem/list doanh nghiệp | Đã chạy | `GET /api/companies`, `GET /api/companies/:id` |
| F07 | Company | Duyệt/từ chối | Đã chạy | Admin route, trạng thái `verified/rejected` |
| F08 | Company | Cập nhật doanh nghiệp | Chưa expose | Có trong proto/service, không có REST route |
| F09 | Marketplace | Xem danh mục | Đã chạy | Public `GET /api/categories` |
| F10 | Marketplace | Xem/tìm nguồn cung | Đã chạy | Public list/detail; lọc phía handler còn đơn giản |
| F11 | Marketplace | Upload ảnh | Đã chạy | Gateway truyền byte ảnh qua Material gRPC; Material Service PUT tới MinIO |
| F12 | Marketplace | Đăng nguồn cung | Đã chạy | Yêu cầu doanh nghiệp `verified` |
| F13 | Marketplace | Sửa/xóa nguồn cung | Một phần | Sửa còn stub; xóa đã kiểm tra Owner/Admin và gọi Material gRPC để xóa database |
| F14 | Demand | Xem nhu cầu | Đã chạy | Public `GET /api/demands` |
| F15 | Demand | Đăng nhu cầu | Một phần | `POST /api/demands` chưa INSERT database |
| F16 | Demand | Seller báo giá Demand | Chưa có | Chưa có API/luồng UI |
| F17 | Offer | Gửi Offer | Đã chạy ở mức prototype | Order Service ghi Offer + escrow holding; Notification qua gRPC/NATS idempotent |
| F18 | Offer | Xem Offer gửi/nhận | Đã chạy | `GET /api/offers?role=buyer|seller` |
| F19 | Offer | Chấp nhận/từ chối | Một phần | Enforce trạng thái `pending`, tạo Transaction và gắn escrow; ownership validation còn thiếu |
| F20 | Offer | Hủy Offer | Chưa expose | Service/proto có, Gateway không có route |
| F21 | Transaction | Xem danh sách/chi tiết/timeline | Đã chạy | `transactions` và `transaction_events` |
| F22 | Transaction | Seller xác nhận giao hàng | Đã chạy | UI chuyển `confirmed -> in_progress` |
| F23 | Transaction | Buyer hoàn tất | Đã chạy | UI chuyển `in_progress -> completed` |
| F24 | Transaction | Enforce state machine | Chưa có | Endpoint chấp nhận chuỗi trạng thái từ client |
| F25 | Review | Tạo/xem đánh giá | Đã chạy | `review_db.reviews` |
| F26 | Notification | Xem/đánh dấu thông báo | Đã chạy | List, unread count, mark one/all |
| F27 | Event | NATS notification consumer | Đã chạy | Order publish `cme.orders.*`; Notification subscribe queue group và chống trùng bằng `reference_id` |
| F28 | Finance | Dashboard phí/ví sàn | Đã chạy ở mức prototype | Gateway gọi Finance RPC thuộc Order Service |
| F29 | Escrow | Tạo/giải phóng escrow | Đã chạy ở mức prototype | Không có tiền thật hoặc payment gateway |
| F30 | Seller wallet | Xem ví/lịch sử | Đã chạy ở mức prototype | Dữ liệu ledger nội bộ |
| F31 | Withdrawal | Gửi và duyệt rút tiền | API có, UI Seller chưa hoàn chỉnh | Không chuyển tiền ngân hàng |
| F32 | Admin | Duyệt doanh nghiệp | Đã chạy | UI + API |
| F33 | Admin | Finance/Escrow | Đã chạy ở mức prototype | UI + API |
| F34 | Admin | Category/Listing/Report/Export | Một phần | Có route UI; REST CRUD/báo cáo chưa đầy đủ |
| F35 | Matching | Tự động ghép cặp cung-cầu | Chưa có | Hướng phát triển |
| F36 | Logistics | Điều phối/theo dõi vận chuyển | Chưa có | Thực hiện ngoài hệ thống |
| F37 | Payment | Thanh toán ngân hàng thật | Chưa có | Chỉ mô phỏng ledger/escrow |

Lưu ý kỹ thuật: escrow vẫn là ledger prototype, không khóa tiền thật. Luồng accept hiện gắn escrow holding vào Transaction nhưng chuỗi cập nhật nhiều bảng chưa nằm trong một SQL transaction duy nhất.

## Quy tắc dùng trong báo cáo

1. Chỉ dùng cụm “đã triển khai” cho dòng **Đã chạy**.
2. Với dòng **Một phần**, phải nêu rõ stub, demo, bypass hoặc thiếu validation.
3. Dòng **Chưa có** chỉ xuất hiện trong phạm vi ngoài MVP, hạn chế hoặc hướng phát triển.
4. Sơ đồ yêu cầu có thể thể hiện chức năng mục tiêu, nhưng caption phải phân biệt với hiện trạng triển khai.
