# Circular Materials Exchange - Full UI Prototype

Đây là bản UI prototype thống nhất cho toàn bộ use case của project môn học. Bản này dùng mock data trong `app.js`, không cần backend, không cần payment gateway, không cần dịch vụ ngoài.

## Cách mở

Mở trực tiếp file:

`design/circular_materials_full_ui/index.html`

Hoặc nếu muốn chạy qua local server:

```powershell
cd "G:\A Cao Học\PTTKHT\design\circular_materials_full_ui"
python -m http.server 5173
```

Sau đó mở `http://localhost:5173`.

## Nội dung đã có

- Chuyển ngôn ngữ UI cố định: `Tiếng Việt` / `English`.
- Dữ liệu động giữ nguyên ngôn ngữ gốc trong mock data.
- Role switcher: Guest, Buyer, Seller, Admin.
- Dashboard business/admin.
- Auth: đăng ký, đăng nhập, OTP demo, quên/đổi mật khẩu.
- Hồ sơ cá nhân, hồ sơ doanh nghiệp, thành viên và phân quyền.
- Danh mục vật liệu và CRUD mock cho admin.
- Marketplace nguồn cung, chi tiết listing, saved listing mock.
- Nguồn cung của seller, form đăng/sửa/đóng listing.
- Nhu cầu mua, form đăng nhu cầu, seller gửi báo giá cho demand.
- Offer/RFQ: buyer gửi offer, seller accept/reject, buyer hủy offer.
- Transaction: tạo tự động, timeline, thanh toán manual/bypass, xác nhận hoàn tất, hủy giao dịch.
- Review/reputation: buyer/seller đánh giá nhau.
- Report violation: user gửi report, admin xử lý report.
- Notification center: in-app notification, mark-read.
- Report/export: preview CSV/PDF demo.
- Use case matrix: coverage UC001-UC059.

## Luồng demo khuyến nghị

1. `Dashboard` để giới thiệu số liệu tổng quan.
2. `Auth & bảo mật` để nói rõ OTP và session là demo.
3. `Hồ sơ doanh nghiệp` để giải thích xác minh doanh nghiệp.
4. `Marketplace` để buyer tìm nguồn cung.
5. `Offer / RFQ` để buyer gửi offer và seller xử lý.
6. `Giao dịch` để xem transaction timeline và payment bypass.
7. `Đánh giá` để hoàn tất uy tín sau giao dịch.
8. `Admin` để duyệt doanh nghiệp, kiểm duyệt listing/report.
9. `Use case matrix` để chứng minh bao phủ toàn bộ use case.

## Ghi chú scope

Thanh toán online, SMS OTP thật, logistics thật, hợp đồng điện tử, carbon credit và AI matching không nằm trong MVP. Prototype mô phỏng các phần này bằng trạng thái mock như `manual_offline`, `bypassed_demo`, `pending_verification` và `admin_review`.

## Ghi chú i18n

Language switch chỉ dịch phần cố định của giao diện như menu, tiêu đề, nhãn form, nút, table header và mô tả màn hình. Các phần sinh từ mock data như tên doanh nghiệp, tên vật liệu, tiêu đề listing, nội dung report, transaction event và message offer được giữ nguyên theo dữ liệu gốc.
