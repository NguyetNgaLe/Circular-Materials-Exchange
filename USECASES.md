# Circular Materials Exchange - Use Cases

## Tổng quan hệ thống

**Sàn giao dịch vật liệu tuần hoàn B2B** - Nền tảng cho phép các doanh nghiệp mua bán vật liệu tái chế (nhựa, kim loại, giấy, gỗ, dệt may, thủy tinh).

---

## 1. Use Case: Đăng ký tài khoản

| Mục | Mô tả |
|-----|-------|
| **Actor** | Người dùng mới |
| **Precondition** | Chưa có tài khoản |
| **Flow** | 1. Truy cập trang `/login` → Chọn tab "Đăng ký" <br> 2. Nhập tên doanh nghiệp, email, mật khẩu <br> 3. Click "Tạo tài khoản mới" <br> 4. Hệ thống tạo tài khoản, chuyển đến Dashboard |
| **Postcondition** | Tài khoản được tạo trong `auth_db`, trạng thái: business |
| **API** | `POST /api/auth/register` |

---

## 2. Use Case: Đăng nhập / Đăng xuất

| Mục | Mô tả |
|-----|-------|
| **Actor** | Người dùng đã đăng ký |
| **Precondition** | Có tài khoản |
| **Flow** | 1. Truy cập `/login` <br> 2. Nhập email, mật khẩu <br> 3. Click "Đăng nhập" <br> 4. Hệ thống xác thực, chuyển đến Dashboard (business) hoặc Admin (admin) |
| **Đăng xuất** | Click "Đăng xuất" → Xóa token, chuyển về trang login |
| **API** | `POST /api/auth/login` |

---

## 3. Use Case: Đăng ký doanh nghiệp

| Mục | Mô tả |
|-----|-------|
| **Actor** | Người dùng business |
| **Precondition** | Đã đăng nhập, chưa có hồ sơ doanh nghiệp |
| **Flow** | 1. Vào "Doanh nghiệp" → Click "Tạo hồ sơ doanh nghiệp" <br> 2. Nhập tên, mã số thuế, địa chỉ, thành phố, mô tả <br> 3. Click "Tạo hồ sơ" <br> 4. Hệ thống tạo hồ sơ, trạng thái: pending |
| **Postcondition** | Company được tạo trong `company_db`, `user.company_id` được cập nhật |
| **API** | `POST /api/companies` |

---

## 4. Use Case: Admin duyệt doanh nghiệp

| Mục | Mô tả |
|-----|-------|
| **Actor** | Admin |
| **Precondition** | Có doanh nghiệp chờ duyệt |
| **Flow** | 1. Vào Admin → "Duyệt doanh nghiệp" <br> 2. Xem danh sách doanh nghiệp pending <br> 3. Click "Duyệt" hoặc "Từ chối" (nhập lý do) |
| **Postcondition** | Company status: verified hoặc rejected |
| **API** | `POST /api/companies/:id/approve` <br> `POST /api/companies/:id/reject` |

---

## 5. Use Case: Đăng vật liệu (Nguồn cung)

| Mục | Mô tả |
|-----|-------|
| **Actor** | Doanh nghiệp đã được duyệt |
| **Precondition** | Company status = verified |
| **Flow** | 1. Vào "Nguồn cung của tôi" → "Đăng vật liệu mới" <br> 2. Nhập tên, danh mục, mô tả, số lượng, giá, địa điểm <br> 3. Upload hình ảnh (từ máy) <br> 4. Click "Đăng Vật Liệu" |
| **Postcondition** | Listing được tạo trong `material_db`, status: active |
| **API** | `POST /api/listings` <br> `POST /api/upload` |
| **Validation** | Nếu DN chưa duyệt → Hiển thị "Doanh nghiệp chưa được duyệt" |

---

## 6. Use Case: Duyệt Marketplace

| Mục | Mô tả |
|-----|-------|
| **Actor** | Tất cả người dùng |
| **Flow** | 1. Truy cập `/marketplace` <br> 2. Xem danh sách vật liệu (có ảnh, giá, số lượng) <br> 3. Lọc theo danh mục, địa điểm, tìm kiếm <br> 4. Sắp xếp theo giá, ngày, số lượng |
| **API** | `GET /api/listings` <br> `GET /api/categories` |

---

## 7. Use Case: Xem chi tiết vật liệu

| Mục | Mô tả |
|-----|-------|
| **Actor** | Tất cả người dùng |
| **Flow** | 1. Click vào vật liệu trên Marketplace <br> 2. Xem: mô tả, thông số, ảnh, giá, số lượng <br> 3. Xem hồ sơ nhà cung cấp <br> 4. Xem đánh giá từ người mua (tab "Đánh giá") |
| **API** | `GET /api/listings/:id` <br> `GET /api/reviews?reviewee_id=...` |

---

## 8. Use Case: Gửi đề nghị mua

| Mục | Mô tả |
|-----|-------|
| **Actor** | Buyer (doanh nghiệp đã duyệt) |
| **Precondition** | DN đã duyệt, không phải sản phẩm của mình |
| **Flow** | 1. Click "Gửi Đề Nghị Mua" trên trang chi tiết <br> 2. Nhập số lượng, giá đề xuất, lời nhắn <br> 3. Click "Thanh Toán & Gửi Đề Nghị" <br> 4. Popup xác nhận thanh toán → Click "Thanh Toán" <br> 5. Popup "Bill Succeed" hiển thị <br> 6. Click "Đã hiểu" → Chuyển đến Đề nghị đã gửi |
| **Postcondition** | Offer tạo (status: pending), Escrow tạo (giữ tiền) |
| **API** | `POST /api/offers` |
| **Validation** | - Không tự mua hàng của mình <br> - DN phải được duyệt |

---

## 9. Use Case: Xem đề nghị đã gửi (Buyer)

| Mục | Mô tả |
|-----|-------|
| **Actor** | Buyer |
| **Flow** | 1. Vào "Đề nghị đã gửi" <br> 2. Xem danh sách: vật liệu, người bán, số lượng, giá, trạng thái |
| **API** | `GET /api/offers?role=buyer` |

---

## 10. Use Case: Xem đề nghị đã nhận (Seller)

| Mục | Mô tả |
|-----|-------|
| **Actor** | Seller |
| **Flow** | 1. Vào "Đề nghị đã nhận" <br> 2. Xem danh sách: vật liệu, người mua, số lượng, giá <br> 3. Click "Chấp nhận" hoặc "Từ chối" |
| **API** | `GET /api/offers?role=seller` <br> `POST /api/offers/:id/accept` <br> `POST /api/offers/:id/reject` |

---

## 11. Use Case: Chấp nhận đề nghị (Seller)

| Mục | Mô tả |
|-----|-------|
| **Actor** | Seller |
| **Flow** | 1. Click "Chấp nhận" trên đề nghị <br> 2. Hệ thống tự động: <br>   - Tạo Transaction (status: confirmed) <br>   - Tạo Escrow (giữ tiền, tính phí 2%) <br>   - Gửi thông báo cho Buyer |
| **Postcondition** | Transaction + Escrow được tạo |
| **API** | `POST /api/offers/:id/accept` |

---

## 12. Use Case: Xác nhận giao hàng (Seller)

| Mục | Mô tả |
|-----|-------|
| **Actor** | Seller |
| **Precondition** | Transaction status = confirmed |
| **Flow** | 1. Vào chi tiết giao dịch <br> 2. Click "Xác Nhận Đã Giao Hàng" <br> 3. Status chuyển: confirmed → in_progress |
| **API** | `POST /api/transactions/:id/status` |

---

## 13. Use Case: Xác nhận nhận hàng & Hoàn tất (Buyer)

| Mục | Mô tả |
|-----|-------|
| **Actor** | Buyer |
| **Precondition** | Transaction status = in_progress |
| **Flow** | 1. Vào chi tiết giao dịch <br> 2. Click "Xác Nhận Đã Nhận Hàng & Hoàn Tất" <br> 3. Hệ thống tự động: <br>   - Status: in_progress → completed <br>   - Giải ngân escrow: 98% → Seller, 2% → Platform <br>   - Cập nhật ví seller <br>   - Cập nhật platform wallet <br>   - Gửi thông báo cho Seller |
| **API** | `POST /api/transactions/:id/status` |

---

## 14. Use Case: Đánh giá người bán (Buyer)

| Mục | Mô tả |
|-----|-------|
| **Actor** | Buyer (sau khi giao dịch hoàn tất) |
| **Precondition** | Transaction status = completed |
| **Flow** | 1. Vào chi tiết giao dịch → Click "Đánh Giá Người Bán" <br> 2. Chọn số sao (1-5) <br> 3. Nhập nhận xét <br> 4. Click "Gửi Đánh Giá" |
| **API** | `POST /api/reviews` |
| **Note** | Chỉ Buyer mới được đánh giá Seller |

---

## 15. Use Case: Xem ví doanh nghiệp (Seller)

| Mục | Mô tả |
|-----|-------|
| **Actor** | Seller |
| **Flow** | 1. Vào "Lịch sử GD" (sidebar) <br> 2. Xem: tổng tiền đã nhận, tổng phí đã trả <br> 3. Xem lịch sử giao dịch |
| **API** | `GET /api/seller/wallet` <br> `GET /api/seller/wallet/transactions` |

---

## 16. Use Case: Admin - Quản lý Escrow

| Mục | Mô tả |
|-----|-------|
| **Actor** | Admin |
| **Flow** | 1. Vào Admin → "Escrow" <br> 2. Xem: tiền đang giữ, số giao dịch holding/released <br> 3. Xem danh sách escrow chi tiết <br> 4. Click "Giải ngân" (nếu cần thủ công) |
| **API** | `GET /api/admin/escrow` <br> `POST /api/admin/escrow/:id/release` |

---

## 17. Use Case: Admin - Quản lý Tài chính

| Mục | Mô tả |
|-----|-------|
| **Actor** | Admin |
| **Flow** | 1. Vào Admin → "Tài chính" <br> 2. Xem: tổng doanh thu, doanh thu tháng, số giao dịch <br> 3. Xem doanh thu theo tháng <br> 4. Xem lịch sử phí giao dịch |
| **API** | `GET /api/admin/finance/overview` <br> `GET /api/admin/finance/fees` |

---

## 18. Use Case: Admin - Quản lý danh mục

| Mục | Mô tả |
|-----|-------|
| **Actor** | Admin |
| **Flow** | 1. Vào Admin → "Danh mục vật liệu" <br> 2. Xem, thêm, sửa, xóa danh mục |
| **API** | `GET /api/categories` |

---

## 19. Use Case: Admin - Quản lý nguồn cung

| Mục | Mô tả |
|-----|-------|
| **Actor** | Admin |
| **Flow** | 1. Vào Admin → "Quản lý nguồn cung" <br> 2. Xem tất cả listings trên hệ thống |
| **API** | `GET /api/listings` (admin xem tất cả) |

---

## 20. Use Case: Thông báo

| Mục | Mô tả |
|-----|-------|
| **Actor** | Tất cả người dùng |
| **Trigger** | - Có đề nghị mua mới <br> - Đề nghị được chấp nhận <br> - Seller giao hàng <br> - Giao dịch hoàn tất |
| **Flow** | 1. Click biểu tượng chuông trên Navbar <br> 2. Xem danh sách thông báo <br> 3. Click "Đánh dấu đã đọc" |
| **API** | `GET /api/notifications` <br> `PUT /api/notifications/:id/read` <br> `PUT /api/notifications/read-all` |

---

## 21. Use Case: Upload hình ảnh

| Mục | Mô tả |
|-----|-------|
| **Actor** | Doanh nghiệp đã duyệt |
| **Flow** | 1. Khi đăng vật liệu mới <br> 2. Click "Chọn ảnh" → Chọn file từ máy <br> 3. Xem preview <br> 4. Hệ thống upload lên MinIO, trả về URL |
| **API** | `POST /api/upload` |
| **Validation** | - Định dạng: jpg, png, gif, webp <br> - Kích thước: tối đa 5MB |

---

## 22. Use Case: Xuất CSV

| Mục | Mô tả |
|-----|-------|
| **Actor** | Admin |
| **Flow** | 1. Vào Admin → "Xuất CSV" <br> 2. Xem danh sách giao dịch <br> 3. Click "Xuất CSV" để tải file |
| **API** | `GET /api/transactions` |

---

## Bảng tổng hợp API

| Nhóm | Endpoint | Method | Mô tả |
|------|----------|--------|-------|
| **Auth** | `/api/auth/register` | POST | Đăng ký |
| | `/api/auth/login` | POST | Đăng nhập |
| | `/api/auth/me` | GET | Lấy thông tin user |
| **Companies** | `/api/companies` | POST | Tạo doanh nghiệp |
| | `/api/companies` | GET | Danh sách DN |
| | `/api/companies/:id` | GET | Chi tiết DN |
| | `/api/companies/:id/approve` | POST | Duyệt DN |
| | `/api/companies/:id/reject` | POST | Từ chối DN |
| **Listings** | `/api/listings` | GET | Danh sách vật liệu |
| | `/api/listings/:id` | GET | Chi tiết vật liệu |
| | `/api/listings` | POST | Đăng vật liệu |
| | `/api/listings/:id` | PUT | Sửa vật liệu |
| | `/api/listings/:id` | DELETE | Xóa vật liệu |
| **Categories** | `/api/categories` | GET | Danh mục |
| **Demands** | `/api/demands` | GET | Nhu cầu mua |
| | `/api/demands` | POST | Đăng nhu cầu |
| **Offers** | `/api/offers` | POST | Tạo đề nghị |
| | `/api/offers` | GET | Danh sách đề nghị |
| | `/api/offers/:id/accept` | POST | Chấp nhận |
| | `/api/offers/:id/reject` | POST | Từ chối |
| **Transactions** | `/api/transactions` | GET | Danh sách GD |
| | `/api/transactions/:id` | GET | Chi tiết GD |
| | `/api/transactions/:id/status` | POST | Cập nhật trạng thái |
| **Reviews** | `/api/reviews` | POST | Tạo đánh giá |
| | `/api/reviews` | GET | Danh sách đánh giá |
| **Notifications** | `/api/notifications` | GET | Thông báo |
| | `/api/notifications/:id/read` | PUT | Đánh dấu đã đọc |
| | `/api/notifications/read-all` | PUT | Đọc tất cả |
| | `/api/notifications/unread-count` | GET | Số chưa đọc |
| **Upload** | `/api/upload` | POST | Upload hình ảnh |
| **Seller** | `/api/seller/wallet` | GET | Ví seller |
| | `/api/seller/wallet/transactions` | GET | Lịch sử ví |
| | `/api/seller/withdraw` | POST | Yêu cầu rút tiền |
| | `/api/seller/withdrawals` | GET | Lịch sử rút tiền |
| **Admin Finance** | `/api/admin/finance/overview` | GET | Tổng quan tài chính |
| | `/api/admin/finance/fees` | GET | Danh sách phí |
| | `/api/admin/finance/wallet` | GET | Ví platform |
| **Admin Escrow** | `/api/admin/escrow` | GET | Danh sách escrow |
| | `/api/admin/escrow/:id/release` | POST | Giải ngân |
| **Admin Withdrawals** | `/api/admin/withdrawals` | GET | Yêu cầu rút tiền |
| | `/api/admin/withdrawals/:id/approve` | POST | Duyệt rút tiền |
| | `/api/admin/withdrawals/:id/reject` | POST | Từ chối rút tiền |

---

## Luồng tiền (Escrow Flow)

```
Buyer thanh toán → Escrow (giữ tiền)
        ↓
Seller giao hàng (status: in_progress)
        ↓
Buyer xác nhận nhận hàng (status: completed)
        ↓
Escrow giải ngân:
  - 98% → Ví Seller
  - 2% → Platform (phí giao dịch)
```

---

## Trạng thái giao dịch

```
pending (offer) → accepted → confirmed → in_progress → completed
                                                      → cancelled
```

---

## Vai trò (Roles)

| Role | Quyền |
|------|-------|
| **business** | Đăng vật liệu, gửi/nhận đề nghị, giao dịch, đánh giá |
| **admin** | Duyệt DN, quản lý danh mục, escrow, tài chính, thống kê |
