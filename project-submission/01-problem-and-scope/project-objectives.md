# Mục tiêu project

Xây dựng prototype nền tảng trao đổi vật liệu tuần hoàn giữa các doanh nghiệp, cho phép:

1. Đăng và quản lý nguồn cung vật liệu dư thừa.
2. Đăng và công bố nhu cầu mua vật liệu.
3. Tìm kiếm, lọc và khám phá nguồn cung hoặc nhu cầu phù hợp.
4. Gửi và xử lý Offer/đề nghị giao dịch.
5. Tạo giao dịch khi hai bên đạt thỏa thuận.
6. Theo dõi trạng thái và lịch sử giao dịch.
7. Xác nhận hoàn tất và đánh giá mức độ uy tín của đối tác.
8. Hỗ trợ Admin kiểm duyệt doanh nghiệp và quản trị dữ liệu chính.

Mục tiêu tự động ghép cặp cung-cầu được định hướng cho phiên bản tương lai. MVP hiện tại hỗ trợ kết nối thông qua công bố tập trung, tìm kiếm, lọc và gửi đề nghị.

Về kỹ thuật, backend được thiết kế theo kiến trúc microservice gồm API Gateway và 6 service nghiệp vụ độc lập, giao tiếp đồng bộ bằng gRPC và hỗ trợ sự kiện bất đồng bộ qua NATS.
