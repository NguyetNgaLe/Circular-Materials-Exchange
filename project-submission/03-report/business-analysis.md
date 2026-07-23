# Chương 2 - Phân tích nghiệp vụ

## 2.1. Cơ cấu tổ chức

Hệ thống kết nối ba nhóm tham gia chính:

- **Doanh nghiệp cung:** công bố vật liệu dư thừa và đóng vai Seller trong giao dịch.
- **Doanh nghiệp cầu:** tìm vật liệu, gửi đề nghị mua và đóng vai Buyer.
- **Đơn vị vận hành nền tảng:** Admin duyệt hồ sơ doanh nghiệp, giám sát giao dịch và theo dõi sổ ghi nhận tài chính.

Một Business User có thể là Buyer ở giao dịch này và Seller ở giao dịch khác. Vì vậy Buyer/Seller là vai trò theo ngữ cảnh, không phải hai loại tài khoản cố định. Cơ cấu được trình bày tại D01.

## 2.2. Các quy trình nghiệp vụ

### Quy trình đăng ký và duyệt doanh nghiệp

1. Người dùng đăng ký tài khoản.
2. Người dùng tạo hồ sơ doanh nghiệp; hồ sơ có trạng thái `pending`.
3. Admin kiểm tra thông tin và chuyển trạng thái sang `verified` hoặc `rejected`.
4. Chỉ doanh nghiệp `verified` mới được đăng nguồn cung hoặc gửi Offer.

Quy trình được trình bày tại D02. Code hiện chưa có API cập nhật Company nên trường hợp bị từ chối cần được chỉnh trực tiếp hoặc bổ sung chức năng ở phiên bản sau.

### Quy trình trao đổi vật liệu

1. Seller tải ảnh và đăng Supply Listing.
2. Buyer tìm kiếm/lọc listing rồi gửi Offer.
3. Seller chấp nhận hoặc từ chối Offer.
4. Khi chấp nhận, hệ thống tạo Transaction ở trạng thái `confirmed`.
5. Seller xác nhận giao hàng để chuyển sang `in_progress`.
6. Buyer xác nhận nhận hàng để chuyển sang `completed`.
7. Hệ thống ghi release escrow, phí nền tảng 2% và cộng Seller wallet ở mức ledger nội bộ.
8. Buyer có thể tạo Review; hai bên theo dõi Notification và timeline.

Quy trình được trình bày tại D03. Logistics, thanh toán thật, hợp đồng và giải quyết tranh chấp diễn ra ngoài hệ thống.

### Quy trình công bố nhu cầu

Guest có thể xem Demand Listing. API tạo Demand đã tồn tại nhưng hiện trả ID demo và chưa ghi database; luồng Seller gửi báo giá cho Demand chưa được triển khai. Nội dung này phải được đánh dấu “một phần” trong báo cáo.

## 2.3. Các lớp lĩnh vực

| Nhóm | Lớp lĩnh vực | Ý nghĩa |
|---|---|---|
| Tài khoản | User, Company | Danh tính, vai trò và trạng thái xác minh doanh nghiệp |
| Marketplace | Category, SupplyListing, DemandListing | Phân loại, nguồn cung và nhu cầu vật liệu |
| Giao dịch | Offer, Transaction, TransactionEvent | Đề nghị mua, thỏa thuận và lịch sử trạng thái |
| Tài chính | EscrowTransaction, PlatformFee, PlatformWallet, SellerWallet, WalletTransaction, WithdrawalRequest | Sổ ghi nhận giữ tiền, phí, số dư và giải ngân prototype |
| Tương tác | Review, Notification | Đánh giá đối tác và thông báo |

Các quan hệ nghiệp vụ được trình bày tại D04. ID liên database chỉ là tham chiếu logic; không có khóa ngoại vật lý xuyên microservice.
