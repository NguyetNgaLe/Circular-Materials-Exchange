# Screen Map - Circular Materials Exchange

## Quy ước

- `reuse`: màn đã có trong bộ design nháp ban đầu.
- `new`: màn được bổ sung trong lần generate này.
- Các màn mới dùng shared style tại `_shared/cme-ui.css`, có `code.html` và `screen.png`.

## Must-have để demo

| Nghiệp vụ | Màn hình | Trạng thái | Folder |
|---|---|---:|---|
| Auth cơ bản | Đăng nhập / đăng ký | reuse | `ng_nh_p_ng_k_b_o_m_t`, `ng_nh_p_ng_k_circular_materials` |
| OTP demo | Xác thực danh tính OTP | reuse | `x_c_th_c_danh_t_nh_otp` |
| Hồ sơ cá nhân | Hồ sơ cá nhân | new | `ho_so_ca_nhan_vn` |
| Hồ sơ doanh nghiệp | Xem/sửa hồ sơ doanh nghiệp | new | `ho_so_doanh_nghiep_vn` |
| Tạo hồ sơ doanh nghiệp | Đăng ký doanh nghiệp | reuse | `ng_k_doanh_nghi_p_vn`, `ng_k_doanh_nghi_p_x_c_minh_b_o_m_t` |
| Admin duyệt doanh nghiệp | Chi tiết duyệt hồ sơ | new | `admin_duyet_doanh_nghiep_detail_vn` |
| Danh mục vật liệu | Admin quản lý danh mục vật liệu | new | `admin_quan_ly_danh_muc_vat_lieu_vn` |
| Đăng nguồn cung | Form đăng vật liệu mới | reuse | `ng_b_n_v_t_li_u_m_i_vn` |
| Nguồn cung của tôi | Danh sách nguồn cung seller | new | `danh_sach_nguon_cung_cua_toi_vn` |
| Marketplace + lọc/tìm kiếm | Sàn giao dịch vật liệu | reuse | `s_n_giao_d_ch_v_t_li_u_vn` |
| Chi tiết vật liệu | Chi tiết nguồn cung HDPE | reuse | `chi_ti_t_v_t_li_u_nh_a_t_i_ch_hdpe_vn` |
| Gửi đề nghị mua | Form gửi offer buyer-to-seller | new | `gui_de_nghi_mua_vn` |
| Offer đã gửi | Buyer quản lý offer đã gửi | new | `offer_da_gui_buyer_vn` |
| Seller chấp nhận/từ chối | Seller quản lý offer đã nhận | new | `offer_da_nhan_seller_vn` |
| Chi tiết đề nghị mua | Chi tiết offer/RFQ cho seller | reuse | `chi_ti_t_ngh_mua_seller`, `qu_n_l_b_o_gi_rfq_vn` |
| Tạo giao dịch tự động | Màn chấp nhận đề nghị thành công | reuse | `x_c_nh_n_giao_d_ch_th_nh_c_ng_seller` |
| Thanh toán bypass | Xác nhận thỏa thuận giao dịch | new | `xac_nhan_thoa_thuan_giao_dich_vn` |
| Chi tiết giao dịch + timeline | Giao dịch TX có TransactionEvent | new | `chi_tiet_giao_dich_timeline_vn` |
| Buyer/seller xác nhận hoàn tất | Giao dịch và review management | reuse + new | `qu_n_l_giao_d_ch_nh_gi`, `chi_tiet_giao_dich_timeline_vn` |
| Đánh giá sau giao dịch | Buyer/seller đánh giá nhau | reuse | `nh_gi_ng_i_mua_seller`, `nh_gi_nh_cung_c_p_buyer` |
| Dashboard business | Dashboard doanh nghiệp | reuse | `b_ng_i_u_khi_n_ng_i_d_ng_vn` |
| Dashboard admin | Dashboard admin | reuse | `trang_qu_n_tr_h_th_ng` |

## Should-have nếu còn thời gian

| Nghiệp vụ | Màn hình | Trạng thái | Folder |
|---|---|---:|---|
| Đăng nhu cầu mua | Form đăng nhu cầu | reuse | `ng_nhu_c_u_mua_v_t_li_u`, `ng_nhu_c_u_th_nh_c_ng_circular_materials` |
| Sàn nhu cầu mua | Demand marketplace | reuse | `s_n_nhu_c_u_mua_v_t_li_u_vn` |
| Nhu cầu mua của tôi | Buyer quản lý nhu cầu đã đăng | new | `danh_sach_nhu_cau_mua_cua_toi_vn` |
| Seller gửi báo giá cho nhu cầu mua | Offer seller-to-buyer | new | `seller_gui_bao_gia_nhu_cau_mua_vn` |
| Thông báo in-app | Danh sách thông báo | new | `danh_sach_thong_bao_vn` |
| Báo cáo vi phạm | Form report của business user | new | `bao_cao_vi_pham_vn` |
| Admin xử lý báo cáo vi phạm | Moderation report detail | new | `admin_quan_ly_bao_cao_vi_pham_vn` |
| Export CSV/PDF demo | Màn xuất báo cáo giao dịch | new | `admin_xuat_bao_cao_csv_vn` |
| Upload ảnh/tài liệu | Ảnh listing và chứng từ vật liệu | new | `upload_tai_lieu_vat_lieu_vn` |

## Màn nên đổi tên hoặc chỉnh nội dung khi triển khai thật

| Màn hiện có | Vấn đề | Gợi ý chỉnh |
|---|---|---|
| `thanh_to_n_k_qu_giao_d_ch` | Tạo cảm giác có thanh toán online thật | Dùng `xac_nhan_thoa_thuan_giao_dich_vn` thay thế trong MVP |
| `x_c_nh_n_thanh_to_n_th_nh_c_ng_buyer` | Có wording ký quỹ/thanh toán thành công | Đổi thành "Ghi nhận thỏa thuận thành công" nếu vẫn dùng |
| Các màn có `LCA`, `ESG`, `Logistics` | Dễ vượt scope môn học | Đổi thành "ước tính", "thông tin giao nhận", "tài liệu đính kèm" |
| Các màn có `CircuTrade`, `Circulix`, `Circular Exchange` | Brand chưa thống nhất | Chốt một brand: `Circular Materials Exchange` hoặc `CircuTrade` |

## Luồng demo khuyến nghị

1. `ng_nh_p_ng_k_b_o_m_t` → đăng nhập.
2. `ho_so_doanh_nghiep_vn` → xem doanh nghiệp đã verified.
3. `ng_b_n_v_t_li_u_m_i_vn` → seller đăng nguồn cung.
4. `s_n_giao_d_ch_v_t_li_u_vn` → buyer tìm kiếm marketplace.
5. `chi_ti_t_v_t_li_u_nh_a_t_i_ch_hdpe_vn` → buyer xem chi tiết.
6. `gui_de_nghi_mua_vn` → buyer gửi offer.
7. `offer_da_nhan_seller_vn` → seller chấp nhận offer.
8. `xac_nhan_thoa_thuan_giao_dich_vn` → bypass thanh toán.
9. `chi_tiet_giao_dich_timeline_vn` → theo dõi timeline.
10. `qu_n_l_giao_d_ch_nh_gi` hoặc `nh_gi_nh_cung_c_p_buyer` → đánh giá sau giao dịch.
11. `trang_qu_n_tr_h_th_ng` → admin dashboard.
