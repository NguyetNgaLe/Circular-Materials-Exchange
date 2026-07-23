# Actor và quyền hạn

| Actor | Chức năng chính | Điều kiện/giới hạn hiện tại |
|---|---|---|
| Guest | Xem nguồn cung, nhu cầu, chi tiết; đăng ký, đăng nhập | Không được tạo dữ liệu nghiệp vụ |
| Business User | Xem dashboard, tạo/xem hồ sơ doanh nghiệp, xem thông báo | Login hiện dùng demo bearer token |
| Buyer | Tìm nguồn cung, gửi Offer, xem giao dịch, xác nhận nhận hàng, đánh giá Seller | Phải có doanh nghiệp `verified` khi gửi Offer |
| Seller | Upload ảnh, đăng nguồn cung, xử lý Offer, xác nhận giao hàng, xem Seller wallet | Phải có doanh nghiệp `verified` khi đăng nguồn cung |
| Admin | Duyệt/từ chối doanh nghiệp, xem finance/escrow, giải ngân, duyệt/từ chối withdrawal | Một số màn quản trị khác còn ở mức UI/mock |

`Buyer` và `Seller` là hai vai trò nghiệp vụ của `Business User`. Một tài khoản có thể vừa mua vừa bán trong các giao dịch khác nhau.

Các hệ thống bên ngoài như ngân hàng, đơn vị logistics, cơ quan xác minh pháp lý chưa được tích hợp và không được vẽ như actor đang hoạt động trong MVP.

