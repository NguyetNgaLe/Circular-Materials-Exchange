# Phân công 13 loại sơ đồ - Nhóm 5 người

## Nguyên tắc phân công

- Làm đủ 13 loại sơ đồ trong danh sách yêu cầu.
- Mỗi loại sơ đồ chỉ có **một người chịu trách nhiệm trọn vẹn**.
- Không chia Use Case theo nhóm, Design Class hoặc Sequence Diagram cho nhiều người.
- Một loại có thể cần nhiều hình con để tránh hình quá lớn, nhưng tất cả hình con vẫn do cùng một người thực hiện.
- Lead chịu trách nhiệm review và ghép bài, không sửa nội dung của người khác khi chưa trao đổi.

## Bảng phân công chính thức

| STT | Loại sơ đồ | Chương | Người phụ trách | File/hình thuộc gói |
|---:|---|---:|---:|---|
| 1 | Sơ đồ cơ cấu tổ chức | 2 | **Người 1 - Lead** | D01 |
| 2 | Activity Diagram/BPMN | 2 | **Người 2** | D02-D03 |
| 3 | Domain Class Diagram | 2 | **Người 2** | D04 |
| 4 | Functional Decomposition/Package | 3 | **Người 1 - Lead** | D05 |
| 5 | Use Case Diagram tổng quan | 3 | **Người 1 - Lead** | D06 |
| 6 | Use Case Diagram theo nhóm | 3 | **Người 3** | D07-D10 |
| 7 | Logical Architecture Diagram | 4 | **Người 3** | D11 |
| 8 | Deployment Diagram | 4 | **Người 3** | D12 |
| 9 | ERD | 4 | **Người 4** | D13 |
| 10 | Design Class Diagram | 4 | **Người 4** | D14-D16 |
| 11 | Sequence Diagram | 4 | **Người 5** | D17-D21 |
| 12 | State Machine Diagram | 4 | **Người 2** | D22-D23 |
| 13 | Mô hình triển khai Docker/server | 5 | **Người 5** | D24 |

## Người 1 - Lead

Phụ trách trọn vẹn các loại 1, 4 và 5:

- Cơ cấu tổ chức nghiệp vụ.
- Phân rã chức năng/Package.
- Use Case tổng quan.

Trách nhiệm Lead: chốt thuật ngữ, kiểm tra D01-D24, quản lý phiên bản, ghép hình vào báo cáo và kiểm tra bản xuất cuối.

## Người 2 - Phân tích nghiệp vụ

Phụ trách trọn vẹn các loại 2, 3 và 12:

- Activity Diagram/BPMN: gồm quy trình duyệt doanh nghiệp và quy trình giao dịch.
- Domain Class Diagram, gồm finance/escrow/wallet.
- State Machine: gồm vòng đời Offer và Transaction.

Không đưa Handler, Service hoặc Repository vào Domain Class.

## Người 3 - Yêu cầu và kiến trúc

Phụ trách trọn vẹn các loại 6, 7 và 8:

- Toàn bộ Use Case theo nhóm: tài khoản/doanh nghiệp, marketplace, giao dịch/đánh giá, quản trị/thông báo.
- Logical Architecture, thể hiện REST Gateway -> gRPC service -> database sở hữu.
- Deployment Diagram hai server.

Buyer và Seller phải được thể hiện là vai trò chuyên biệt của Business User.

## Người 4 - Dữ liệu và thiết kế lớp

Phụ trách trọn vẹn các loại 9 và 10:

- ERD tổng hợp 6 database logic, 18 bảng và các cột tương thích dữ liệu demo.
- Toàn bộ Design Class: Auth/Company, Material/Upload và Order/Finance/Escrow/Review/Notification.

Phải phân biệt khóa ngoại vật lý trong cùng database với tham chiếu ID logic giữa các microservice.

## Người 5 - Hành vi và triển khai

Phụ trách trọn vẹn các loại 11 và 13:

- Toàn bộ Sequence Diagram as-built: duyệt doanh nghiệp, upload/đăng nguồn cung, gửi Offer, chấp nhận Offer, hoàn tất/release ledger/đánh giá qua gRPC.
- Mô hình triển khai frontend server và backend Docker server.

Docker hiện có một PostgreSQL container chứa 6 database; frontend chạy bằng Nginx trên server riêng và chưa nằm trong Docker Compose.

## Định nghĩa hoàn thành

Một loại sơ đồ chỉ được đánh dấu `DONE` khi tất cả hình thuộc loại đó có đủ:

- File nguồn chỉnh sửa được.
- File SVG.
- File PNG độ rộng tối thiểu 2000 px.
- Tiêu đề và chú thích dùng được trong báo cáo.
- Không mâu thuẫn với `00-reference/system-baseline.md`.
- Được chính người phụ trách hoàn thiện và Lead review.
