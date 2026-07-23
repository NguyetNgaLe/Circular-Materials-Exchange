# Danh mục Use Case

> Danh mục này dùng cho mục 3.1.2 của báo cáo. Cột trạng thái phải được giữ khi chuyển nội dung sang Word/LaTeX để phân biệt yêu cầu mục tiêu với code as-built.

## Tài khoản và doanh nghiệp

| ID | Tên Use Case | Actor chính | Trạng thái | Sơ đồ liên quan |
|---|---|---|---|---|
| UC001 | Đăng ký tài khoản | Guest | Đã chạy | D06, D07 |
| UC002 | Đăng nhập | Guest | Đã chạy: JWT, tương thích demo token cũ | D06, D07 |
| UC003 | Xem thông tin người dùng hiện tại | Business User | Đã chạy | D07 |
| UC004 | Tạo hồ sơ doanh nghiệp | Business User | Đã chạy | D02, D07, D17 |
| UC005 | Xem hồ sơ/trạng thái doanh nghiệp | Business User | Đã chạy | D07 |
| UC006 | Duyệt doanh nghiệp | Admin | Đã chạy | D02, D07, D17 |
| UC007 | Từ chối doanh nghiệp | Admin | Đã chạy | D02, D07 |

## Marketplace cung và cầu

| ID | Tên Use Case | Actor chính | Trạng thái | Sơ đồ liên quan |
|---|---|---|---|---|
| UC008 | Xem danh mục vật liệu | Guest | Đã chạy | D08 |
| UC009 | Xem, tìm kiếm và lọc nguồn cung | Guest | Đã chạy, bộ lọc còn đơn giản | D06, D08, D19 |
| UC010 | Xem chi tiết nguồn cung | Guest | Đã chạy | D08, D19 |
| UC011 | Upload ảnh vật liệu | Seller | Đã chạy | D08, D18 |
| UC012 | Đăng nguồn cung | Seller | Đã chạy | D03, D08, D18 |
| UC013 | Cập nhật nguồn cung | Seller | Một phần: handler stub | D08 |
| UC014 | Xóa nguồn cung | Seller | Một phần: handler stub | D08 |
| UC015 | Xem danh sách nhu cầu | Guest | Đã chạy | D08 |
| UC016 | Đăng nhu cầu mua | Buyer | Một phần: chưa INSERT database | D03, D08 |

## Offer và giao dịch

| ID | Tên Use Case | Actor chính | Trạng thái | Sơ đồ liên quan |
|---|---|---|---|---|
| UC017 | Gửi Offer | Buyer | Đã chạy | D03, D09, D19 |
| UC018 | Xem Offer đã gửi | Buyer | Đã chạy | D09 |
| UC019 | Xem Offer đã nhận | Seller | Đã chạy | D09 |
| UC020 | Chấp nhận Offer | Seller | Đã chạy, validation còn thiếu | D03, D09, D20, D22 |
| UC021 | Từ chối Offer | Seller | Đã chạy, validation còn thiếu | D03, D09, D22 |
| UC022 | Xem danh sách giao dịch | Buyer/Seller | Đã chạy | D09 |
| UC023 | Xem chi tiết và timeline | Buyer/Seller | Đã chạy | D09, D21 |
| UC024 | Xác nhận đã giao hàng | Seller | Đã chạy | D09, D21, D23 |
| UC025 | Xác nhận nhận hàng và hoàn tất | Buyer | Đã chạy | D09, D21, D23 |

## Review và Notification

| ID | Tên Use Case | Actor chính | Trạng thái | Sơ đồ liên quan |
|---|---|---|---|---|
| UC026 | Tạo Review | Buyer/Seller | Đã chạy; UI chính hỗ trợ Buyer đánh giá Seller | D09, D21 |
| UC027 | Xem Review/điểm trung bình | Business User | Đã chạy | D09 |
| UC028 | Xem Notification và số chưa đọc | Business User | Đã chạy | D10 |
| UC029 | Đánh dấu Notification đã đọc | Business User | Đã chạy | D10 |

## Finance, escrow và quản trị

| ID | Tên Use Case | Actor chính | Trạng thái | Sơ đồ liên quan |
|---|---|---|---|---|
| UC030 | Xem Seller wallet và lịch sử | Seller | Đã chạy ở mức ledger prototype | D09, D10 |
| UC031 | Tạo withdrawal request | Seller | API có, UI chưa hoàn chỉnh | D10 |
| UC032 | Xem finance overview và phí | Admin | Đã chạy ở mức prototype | D10 |
| UC033 | Xem/giải phóng escrow | Admin | Đã chạy ở mức prototype | D10, D21 |
| UC034 | Duyệt/từ chối withdrawal | Admin | API có, UI chưa hoàn chỉnh | D10 |

## Yêu cầu tương lai, không đánh mã UC triển khai

- Tự động matching/recommendation cung-cầu.
- Seller gửi báo giá trực tiếp cho Demand.
- Thanh toán ngân hàng và escrow thật.
- Logistics/tracking.
- Hợp đồng điện tử, kiểm định và tranh chấp.
