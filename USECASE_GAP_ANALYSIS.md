# Circular Material Exchange Platform - Gap Analysis & MVP Solution

## 1. Nhận định tổng quan

File use case hiện tại đã đủ để làm project môn học, vì nó bao phủ được nghiệp vụ cốt lõi của một sàn giao dịch vật liệu:

- Đăng ký, đăng nhập, xác thực tài khoản.
- Tạo và xác minh doanh nghiệp.
- Đăng nguồn cung vật liệu.
- Đăng nhu cầu mua vật liệu.
- Tìm kiếm, lọc, xem chi tiết vật liệu.
- Gửi đề nghị mua / RFQ.
- Seller chấp nhận hoặc từ chối đề nghị.
- Hệ thống tạo giao dịch.
- Buyer và seller theo dõi, xác nhận hoàn tất.
- Đánh giá sau giao dịch.
- Admin quản lý user, doanh nghiệp, vật liệu, listing, giao dịch, báo cáo.

Tuy nhiên, nếu xem như sản phẩm thương mại thật thì bản hiện tại chưa đủ, vì các phần thanh toán, logistics, hợp đồng, định danh doanh nghiệp, kiểm định chất lượng vật liệu, tranh chấp, audit và tuân thủ pháp lý chưa đủ sâu. Với phạm vi môn học, điều này chấp nhận được. Nên trình bày rõ ràng rằng hệ thống là prototype học thuật/MVP, chưa xử lý thanh toán và logistics thật.

## 2. Hướng MVP phù hợp nhất

Nên chốt MVP theo luồng sau:

1. Guest xem trang chủ, marketplace công khai và chi tiết vật liệu.
2. User đăng ký, xác thực OTP demo, đăng nhập.
3. User tạo hồ sơ doanh nghiệp.
4. Admin duyệt doanh nghiệp thủ công.
5. Seller đăng nguồn cung vật liệu.
6. Buyer tìm kiếm vật liệu và gửi đề nghị mua.
7. Seller xem RFQ/de nghị mua và chấp nhận hoặc từ chối.
8. Khi seller chấp nhận, hệ thống tự động tạo giao dịch.
9. Thanh toán được bypass bằng trang "Xác nhận thỏa thuận giao dịch".
10. Buyer và seller cập nhật trạng thái, xác nhận hoàn tất.
11. Hệ thống cho phép hai bên đánh giá nhau.
12. Admin xem dashboard và quản lý dữ liệu chính.

Đây là luồng demo tốt nhất vì thể hiện đủ giá trị của sàn giao dịch mà không cần tích hợp bên thứ ba.

## 3. Giải pháp cho thanh toán

Không nên làm thanh toán online thật trong project môn học. Nên thay bằng cơ chế "demo/manual settlement".

### Đề xuất nghiệp vụ

Trang hiện tại `thanh_to_n_k_qu_giao_d_ch` không nên hiển thị như thanh toán online thật. Nên đổi thành:

- "Xác nhận thỏa thuận giao dịch"
- "Ghi nhận thanh toán ngoài hệ thống"
- "Hoàn tất bước xác nhận demo"

Nút "Xác nhận & Thanh toán" nên đổi thành:

- "Xác nhận giao dịch"
- hoặc "Ghi nhận đã thỏa thuận"

Sau khi bấm, hệ thống không gọi payment gateway. Hệ thống chỉ tạo bản ghi:

- `payment_status = bypassed_demo`
- `payment_method = manual_offline`
- `settlement_note = Thanh toán được thực hiện ngoài hệ thống trong phạm vi prototype`

### Trạng thái gợi ý

- `offer.pending`: buyer đã gửi đề nghị.
- `offer.accepted`: seller đã chấp nhận.
- `transaction.confirmed`: giao dịch được tạo.
- `transaction.in_progress`: hai bên đang thực hiện ngoài hệ thống.
- `transaction.completed`: buyer và seller đều xác nhận hoàn tất.
- `transaction.cancelled`: giao dịch bị hủy.
- `transaction.disputed`: có tranh chấp, chỉ nên để trong future hoặc admin xử lý đơn giản.

Không cần làm escrow, refund, QR banking, thẻ tín dụng, vì các phần đó đều cần đối tác thanh toán và pháp lý.

## 4. Các chức năng cần bên thứ ba và solution thay thế

| Chức năng | Cần bên thứ ba nếu làm thật | Giải pháp cho project môn học |
|---|---|---|
| Thanh toán online | Payment gateway, ngân hàng, escrow, đối soát | Bypass/manual offline, chỉ ghi nhận trạng thái demo |
| OTP SMS | Nhà cung cấp SMS | Dùng OTP giả lập `123456` hoặc gửi email nội bộ |
| Email xác thực | SMTP/email service | Có thể mock bằng màn OTP, hoặc log link ra console |
| Xác minh mã số thuế | API cơ quan nhà nước/nguồn dữ liệu doanh nghiệp | Admin duyệt thủ công dựa trên file upload demo |
| Logistics/vận chuyển | Đơn vị vận chuyển, bản đồ, tracking | Lưu địa điểm giao nhận, ghi chú vận chuyển, trạng thái tự cập nhật thủ công |
| Hợp đồng điện tử | E-signature provider, CA/chữ ký số | Checkbox đồng ý điều khoản + xuất phiếu giao dịch PDF/CSV |
| Kiểm định chất lượng vật liệu | Lab/chứng chỉ vật liệu | Upload chứng chỉ dạng file, hiển thị "chưa xác minh" hoặc "admin đã xem" |
| LCA/CO2/ESG | Bộ hệ số phát thải, chuyên gia môi trường | Tính ước lượng theo bảng hệ số cố định, ghi rõ "estimated" |
| Carbon credit | Đơn vị đăng ký tín chỉ carbon | Chuyển sang định hướng phát triển, không đưa vào MVP |
| Google/LinkedIn login | OAuth app và cấu hình security | Ẩn nút, disable, hoặc để "coming soon" |
| AI matching / recommendation | Model AI, dữ liệu lịch sử | Lọc theo rule đơn giản: vật liệu, tỉnh/thành, khối lượng, giá |

## 5. Màn hình đang có trong design

Folder `design/stitch_minimalist_web_interface` đã có các màn quan trọng:

- Trang chủ.
- Đăng nhập/đăng ký.
- OTP/xác thực danh tính.
- Bảo mật tài khoản.
- Đăng ký doanh nghiệp.
- Marketplace nguồn cung vật liệu.
- Chi tiết vật liệu.
- Đăng vật liệu mới.
- Sàn nhu cầu mua vật liệu.
- Đăng nhu cầu mua.
- Đăng nhu cầu thành công.
- Quản lý RFQ / đàm phán.
- Chi tiết đề nghị mua cho seller.
- Quản lý giao dịch và đánh giá.
- Đánh giá seller / buyer.
- Dashboard doanh nghiệp.
- Dashboard admin.
- Báo cáo tác động môi trường.
- Màn thanh toán / kết quả giao dịch.

Concept UI phù hợp với B2B: tối giản, nhiều dữ liệu, màu xanh lá cho hành động chính, xanh dương cho navigation/tin cậy, cam cho hành động giao dịch. Nên giữ hướng này.

## 6. Màn hình còn thiếu hoặc nên bổ sung

### Bắt buộc nên có cho MVP

1. **Hồ sơ cá nhân**
   - Phục vụ UC007.
   - Bao gồm họ tên, email, số điện thoại, chức vụ, đổi mật khẩu nhanh.

2. **Hồ sơ doanh nghiệp - xem/sửa**
   - Sau khi đăng ký doanh nghiệp, user cần có nơi xem trạng thái: draft, pending, verified, rejected.
   - Cần hiển thị lý do từ chối nếu admin không duyệt.

3. **Danh sách nguồn cung của tôi**
   - Phục vụ UC019, UC020, UC021.
   - Cần có nút sửa, ẩn/đóng, xem offer liên quan.

4. **Danh sách nhu cầu mua của tôi**
   - Phục vụ UC025, UC026.
   - Cần có nút sửa, đóng nhu cầu, xem seller đã gửi báo giá.

5. **Offer đã gửi**
   - Buyer xem các đề nghị mua đã gửi.
   - Cho phép hủy khi offer còn `pending`.

6. **Offer đã nhận**
   - Seller xem các RFQ/de nghị mua theo từng listing.
   - Hiển thị rõ nút chấp nhận, từ chối, xem chi tiết.

7. **Chi tiết giao dịch**
   - Hiển thị timeline: offer accepted -> transaction confirmed -> in progress -> buyer confirmed -> seller confirmed -> completed.
   - Có nút buyer/seller xác nhận hoàn tất.
   - Có nút yêu cầu hủy và lý do hủy.

8. **Danh sách thông báo**
   - Phục vụ UC052, UC053.
   - Có thể làm đơn giản bằng notification in-app, không cần push/email.

9. **Báo cáo vi phạm**
   - Form report cho listing, transaction, company hoặc user.
   - Cần màn admin xử lý report.

10. **Màn quên mật khẩu / đặt lại mật khẩu**
    - Trong use case có UC005 nhưng design hiện mới thấy auth/OTP.

### Nên có nếu kịp

1. **Quản lý thành viên doanh nghiệp**
   - Phục vụ UC012, UC013.
   - Nếu thời gian ít, có thể rút gọn: mỗi doanh nghiệp chỉ có 1 owner.

2. **Saved listings / saved search**
   - Use case có UC032.
   - Không bắt buộc cho demo giao dịch.

3. **Admin quản lý danh mục vật liệu**
   - Phục vụ UC015-UC017.
   - Rất nên có vì admin dashboard hiện có chưa thể hiện CRUD danh mục rõ ràng.

4. **Admin kiểm duyệt listing và demand**
   - Phục vụ UC023 và quản trị dữ liệu.
   - Có thể làm bằng status `pending_review`, `active`, `rejected`.

5. **Export CSV**
   - Để thay cho xuất PDF nếu thời gian ít.
   - CSV dễ làm, đủ để chấm điểm chức năng báo cáo.

## 7. Điểm chưa phù hợp trong design hiện tại

1. **Thương hiệu chưa thống nhất**
   - Đang có nhiều tên: Circular Materials, CircuTrade, Circulix, Circular Exchange.
   - Nên chốt 1 tên. Gợi ý: **Circular Materials Exchange** hoặc **CircuTrade**.

2. **Ngôn ngữ chưa thống nhất**
   - Có màn tiếng Việt, có màn tiếng Anh.
   - Nếu báo cáo môn học bằng tiếng Việt, nên Việt hóa các màn chính.

3. **Nav đang hơi marketing**
   - Các mục "Solutions", "Resources", "About", "Careers" phù hợp landing page hơn app nghiệp vụ.
   - App nav nên ưu tiên: Marketplace, Nhu cầu mua, Đăng vật liệu, Giao dịch, Dashboard, Thông báo, Tài khoản.

4. **LCA/ESG/logistics đang quá nặng**
   - Các tab "LCA & Sustainability", "Logistics & Handling", "Download Full LCA Report" dễ làm giảng viên nghĩ có tích hợp thật.
   - Nên đổi thành "Tác động môi trường ước tính", "Thông tin giao nhận", "Tài liệu đính kèm".

5. **Thanh toán đang tạo cảm giác có payment gateway**
   - Cần đổi thành manual/bypass như mục 3.

6. **Social login nên bỏ khỏi MVP**
   - Google/LinkedIn cần OAuth thật.
   - Nên ẩn hoặc để disabled.

7. **Bảo mật 2FA nên để optional**
   - Màn security center đẹp, nhưng 2FA thật không cần trong MVP.
   - Nên giữ đổi mật khẩu, lịch sử đăng nhập demo, OTP xác thực tài khoản.

8. **Admin dashboard chưa đủ CRUD**
   - Hiện có tổng quan và hàng chờ xác minh.
   - Cần thêm các list/detail quản trị: user, company, material category, listing, demand, transaction, report.

## 8. Gap lớn nhất trong use case

### 8.1. Luồng nhu cầu mua chưa khép kín

Use case có "Buyer đăng nhu cầu mua" và design có "Sàn nhu cầu mua vật liệu", nhưng use case chưa đặc tả rõ luồng seller phản hồi nhu cầu mua.

Nên thêm 2 use case:

- **UC058 - Seller gửi báo giá cho nhu cầu mua**
  - Seller xem nhu cầu mua.
  - Seller chọn "Gửi báo giá".
  - Seller nhập số lượng, giá, thời gian giao, ghi chú.
  - Hệ thống tạo offer loại `seller_to_buyer`.

- **UC059 - Buyer xử lý báo giá từ seller**
  - Buyer xem báo giá nhận được cho nhu cầu.
  - Buyer chấp nhận hoặc từ chối.
  - Nếu chấp nhận, hệ thống tạo transaction.

Nếu nhóm muốn đơn giản hơn, có thể ghi rõ trong scope: "Đăng nhu cầu mua chỉ dùng để công khai nhu cầu; giao dịch chính vẫn bắt đầu từ buyer gửi offer cho supply listing."

### 8.2. Chưa nói rõ chất lượng và rủi ro vật liệu

Nên thêm quy tắc:

- Vật liệu nguy hại chỉ được đánh dấu `hazardous_level`, không xử lý logistics/hợp đồng trong MVP.
- Seller phải khai báo chất lượng, tình trạng, ảnh/chứng từ nếu có.
- Hệ thống hiển thị disclaimer "Thông tin do doanh nghiệp đăng tải, nên cần xác minh trước khi giao dịch thật."

### 8.3. Chưa có audit lịch sử trạng thái

Nên thêm bảng `TransactionEvent` hoặc `StatusHistory`:

- transaction_id
- actor_id
- from_status
- to_status
- note
- created_at

Phần này giúp demo timeline giao dịch và làm use case chặt chẽ hơn.

## 9. Điều chỉnh thực thể dữ liệu

Nên bổ sung:

- `VerificationRequest`: lưu hồ sơ xác minh doanh nghiệp và trạng thái duyệt.
- `MaterialImage`: ảnh của supply listing.
- `ListingAttachment`: chứng chỉ, file kiểm định, hóa đơn, tài liệu vật liệu.
- `OfferHistory`: lịch sử đàm phán giá/khối lượng.
- `TransactionEvent`: timeline giao dịch.
- `PaymentRecord`: bản ghi thanh toán demo/manual.
- `ReportEvidence`: file bằng chứng khi báo cáo vi phạm.

Nên thêm trường:

- `SupplyListing.min_order_quantity`
- `SupplyListing.packaging`
- `SupplyListing.pickup_address`
- `SupplyListing.quality_certificate_status`
- `SupplyListing.hazardous_level`
- `DemandListing.deadline`
- `DemandListing.delivery_location`
- `PurchaseOffer.offer_type` với giá trị `buyer_to_seller` hoặc `seller_to_buyer`
- `Transaction.payment_status` với giá trị `not_required`, `manual_offline`, `bypassed_demo`

## 10. Screen map để hoàn thiện

### Guest

- Trang chủ
- Marketplace công khai
- Chi tiết nguồn cung
- Đăng ký
- Đăng nhập
- Quên mật khẩu
- Đặt lại mật khẩu
- OTP/xác thực tài khoản

### Business User

- Dashboard doanh nghiệp
- Hồ sơ cá nhân
- Hồ sơ doanh nghiệp
- Tạo/cập nhật doanh nghiệp
- Danh sách nguồn cung của tôi
- Tạo/sửa nguồn cung
- Danh sách nhu cầu mua của tôi
- Tạo/sửa nhu cầu mua
- Marketplace nguồn cung
- Sàn nhu cầu mua
- Chi tiết nguồn cung
- Chi tiết nhu cầu mua
- Offer đã gửi
- Offer đã nhận
- Chi tiết offer/RFQ
- Danh sách giao dịch
- Chi tiết giao dịch
- Xác nhận thỏa thuận giao dịch
- Xác nhận hoàn tất giao dịch
- Đánh giá đối tác
- Thông báo
- Báo cáo vi phạm
- Bảo mật tài khoản

### Admin

- Dashboard admin
- Quản lý user
- Quản lý doanh nghiệp
- Duyệt doanh nghiệp
- Quản lý danh mục vật liệu
- Quản lý nguồn cung
- Quản lý nhu cầu mua
- Quản lý offer/RFQ
- Quản lý giao dịch
- Quản lý báo cáo vi phạm
- Quản lý đánh giá
- Xuất báo cáo CSV/PDF

## 11. Ưu tiên triển khai

### Must-have để demo được

- Auth cơ bản.
- Hồ sơ doanh nghiệp.
- Admin duyệt doanh nghiệp.
- Danh mục vật liệu.
- Đăng nguồn cung.
- Marketplace + lọc/tìm kiếm.
- Chi tiết vật liệu.
- Gửi đề nghị mua.
- Seller chấp nhận/từ chối.
- Tạo giao dịch tự động.
- Chi tiết giao dịch + timeline.
- Buyer/seller xác nhận hoàn tất.
- Đánh giá sau giao dịch.
- Dashboard business và admin có số liệu giả lập/thực.

### Should-have nếu còn thời gian

- Đăng nhu cầu mua.
- Seller gửi báo giá cho nhu cầu mua.
- Thông báo in-app.
- Báo cáo vi phạm.
- Export CSV.
- Upload ảnh/tài liệu.

### Future / không nên làm trong bản nộp

- Thanh toán online.
- Escrow/ký quỹ.
- Logistics thật.
- Hợp đồng điện tử.
- Carbon credit.
- Digital Product Passport.
- ESG scoring nghiêm túc.
- AI/Graph ML.
- Social login.
- SMS OTP thật.

## 12. Kết luận để đưa vào báo cáo

Hệ thống nên được định vị là prototype học thuật của sàn giao dịch vật liệu tuần hoàn B2B. Bản MVP tập trung vào việc kết nối nguồn cung và nhu cầu, xử lý RFQ/de nghị mua, tạo giao dịch, theo dõi trạng thái và đánh giá uy tín. Các nghiệp vụ phụ thuộc bên thứ ba như thanh toán, logistics, hợp đồng điện tử và xác minh pháp lý được mô phỏng bằng cơ chế thủ công hoặc bypass demo, đồng thời đưa vào phần định hướng phát triển tương lai.

Hướng này giúp project vừa đủ sâu về nghiệp vụ, vừa tránh phạm vi quá lớn, và vẫn demo được một vòng đời giao dịch hoàn chỉnh.
