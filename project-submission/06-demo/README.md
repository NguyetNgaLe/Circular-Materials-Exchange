# Demo

## Luồng demo chính

1. Kiểm tra backend `/health` và mở marketplace.
2. Business User đăng nhập.
3. Tạo doanh nghiệp; Admin duyệt `pending -> verified`.
4. Seller upload ảnh lên MinIO và đăng nguồn cung nhựa PP.
5. Buyer tìm nguồn cung, xem Review và gửi Offer.
6. Kiểm tra Seller nhận Notification và escrow ledger `holding`.
7. Seller chấp nhận; hệ thống tạo Transaction `confirmed`.
8. Seller xác nhận đã giao hàng -> `in_progress`.
9. Buyer xác nhận đã nhận hàng -> `completed`.
10. Kiểm tra escrow `released`, phí 2%, Seller wallet và timeline.
11. Buyer tạo Review cho Seller.
12. Admin xem finance/escrow.

## Luồng dự phòng

- Doanh nghiệp chưa verified bị chặn đăng nguồn cung/gửi Offer.
- File upload quá 5 MB hoặc sai định dạng bị từ chối.
- Business User truy cập `/api/admin/*` nhận 403.

## Không demo như chức năng hoàn thiện

- Tạo Demand mới, sửa/xóa Listing.
- Matching tự động.
- Thanh toán/chuyển khoản ngân hàng thật.
- Logistics/tracking.
- JWT thật.

Khi trình bày escrow, dùng cụm “ledger mô phỏng” hoặc “ghi nhận prototype”.

