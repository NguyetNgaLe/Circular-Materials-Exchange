# Luồng nghiệp vụ cốt lõi

## Luồng nguồn cung đến giao dịch

1. Seller đăng nguồn cung vật liệu dư thừa.
2. Hệ thống công bố nguồn cung trên marketplace.
3. Buyer tìm kiếm, lọc và xem chi tiết vật liệu.
4. Buyer gửi Offer gồm số lượng, giá đề xuất và ghi chú.
5. Seller xem Offer và chấp nhận hoặc từ chối.
6. Khi Seller chấp nhận, hệ thống tạo Transaction.
7. Hai bên thực hiện thanh toán/giao nhận ngoài hệ thống.
8. Buyer và Seller cập nhật trạng thái; hệ thống lưu TransactionEvent.
9. Khi hoàn tất, hai bên có thể đánh giá nhau.

## Luồng nhu cầu mua

1. Buyer đăng nhu cầu mua.
2. Hệ thống công khai nhu cầu để các doanh nghiệp khác tìm kiếm.
3. Seller có thể nhận biết cơ hội cung ứng.
4. Trong MVP hiện tại, giao dịch chính vẫn được khởi tạo bằng Offer gắn với SupplyListing.

Luồng Seller gửi báo giá trực tiếp cho Demand là phần mở rộng chưa có API hoàn chỉnh.

