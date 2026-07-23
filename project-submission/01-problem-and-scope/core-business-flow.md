# Luồng nghiệp vụ cốt lõi

## Luồng doanh nghiệp

1. Business User đăng ký hoặc đăng nhập.
2. Người dùng tạo hồ sơ doanh nghiệp, hệ thống lưu trạng thái `pending`.
3. Admin duyệt thành `verified` hoặc từ chối thành `rejected`.
4. Chỉ doanh nghiệp `verified` mới được đăng nguồn cung và gửi Offer trong luồng demo.

## Luồng nguồn cung đến giao dịch

1. Seller upload ảnh và đăng nguồn cung vật liệu.
2. Hệ thống công bố nguồn cung trên marketplace.
3. Buyer tìm kiếm, lọc và xem chi tiết.
4. Buyer nhập số lượng, giá đề xuất, lời nhắn và gửi Offer.
5. Hệ thống ghi Offer `pending`, tạo escrow ledger `holding` và Notification cho Seller.
6. Seller chấp nhận hoặc từ chối Offer.
7. Khi chấp nhận, hệ thống tạo Transaction `confirmed`, TransactionEvent và liên kết escrow với giao dịch.
8. Seller xác nhận đã giao hàng, Transaction chuyển `in_progress`.
9. Buyer xác nhận đã nhận hàng, Transaction chuyển `completed`.
10. Hệ thống release escrow ledger, ghi phí 2%, cộng Seller wallet và tạo Notification.
11. Buyer có thể tạo Review cho Seller.

Escrow và ví chỉ là dữ liệu mô phỏng. Thanh toán và vận chuyển thật diễn ra ngoài hệ thống.

## Luồng nhu cầu mua

1. Hệ thống hiển thị các Demand đang có trong `material_db`.
2. Buyer có giao diện đăng nhu cầu mua.
3. Endpoint tạo Demand hiện mới trả ID demo, chưa ghi database.
4. Chưa có luồng Seller gửi báo giá trực tiếp cho Demand.

Do đó luồng giao dịch đầy đủ trong bản demo phải bắt đầu từ Buyer gửi Offer cho SupplyListing.

## Luồng quản trị tài chính prototype

1. Khi giao dịch hoàn tất, hệ thống ghi platform fee và số tiền Seller được nhận.
2. Seller xem tổng tiền nhận, phí và lịch sử wallet.
3. Seller có API tạo withdrawal request; Admin có API duyệt hoặc từ chối.
4. Admin xem finance overview, fee, platform wallet và escrow.
5. Không có kết nối ngân hàng; các thao tác chỉ thay đổi ledger trong `order_db`.

