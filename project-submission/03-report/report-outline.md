# Khung báo cáo theo guideline

## Chương 1. Đặt vấn đề

1. Bối cảnh kinh tế tuần hoàn và bài toán thực tế.
2. Hiện trạng, nhu cầu và hệ quả.
3. Mục tiêu, phạm vi MVP.
4. Yêu cầu chức năng sơ bộ.
5. Giải pháp nghiệp vụ và kỹ thuật.
6. Giới hạn: matching, payment, logistics.

Nguồn soạn: `01-problem-and-scope/*` và `latex/chuong-01-dat-van-de.tex`.

## Chương 2. Phân tích nghiệp vụ (Optional)

Nguồn soạn chính: `business-analysis.md`.

### 2.1. Cơ cấu tổ chức

- Doanh nghiệp cung, doanh nghiệp mua và đơn vị vận hành.
- Sơ đồ D01.

### 2.2. Các quy trình nghiệp vụ

- Đăng ký và duyệt doanh nghiệp: D02.
- Nguồn cung -> Offer -> giao dịch -> escrow ledger -> Review: D03.

### 2.3. Các lớp lĩnh vực

- User, Company, Listing, Offer, Transaction, finance/escrow/wallet, Review, Notification.
- Domain Class D04.

## Chương 3. Phân tích yêu cầu

### 3.1. Yêu cầu chức năng

- 3.1.1 Phân rã chức năng: D05.
- Use Case tổng quan: D06.
- Use Case theo nhóm: D07-D10.
- 3.1.2 Danh mục UC001-UC034: `use-case-catalog.md`.
- Đặc tả UC trọng yếu: `use-case-specifications.md`.
- Mỗi UC phải ghi trạng thái Đã chạy/Một phần/Chưa có.

### 3.2. Yêu cầu phi chức năng

- Dùng `non-functional-requirements.md`.
- Tách “yêu cầu mong muốn” khỏi “hiện trạng đạt được”.

## Chương 4. Thiết kế

Nguồn soạn chính: `architecture-design.md`.

### 4.1. Kiến trúc

- 4.1.1 Kiến trúc logic/as-built microservice: D11.
- 4.1.2 Kiến trúc triển khai/as-built: D12.
- Giải thích ranh giới REST Gateway, gRPC service và database ownership.

### 4.2. Các kỹ thuật thiết kế

- React component/page/store/API client.
- REST Gateway, middleware và gRPC contract.
- Handler-Service-Repository trong 6 service.
- Repository/PostgreSQL và Database per Service ở mức logic.
- NATS event và MinIO object storage.

### 4.3. Thiết kế ca sử dụng

- Dữ liệu/ERD 18 bảng: D13.
- Design Class: D14-D16.
- Sequence cho UC004/UC006/UC012/UC017/UC020/UC024/UC025/UC026: D17-D21.
- State Machine Offer/Transaction: D22-D23.

## Chương 5. Thực hiện và triển khai

- 5.1 Cấu trúc source và thành phần.
- 5.2 Frontend Nginx server và backend Docker server.
- Bootstrap 6 database/18 bảng.
- Mô hình Docker/server: D24.
- Nguồn soạn: `implementation-deployment.md`.

## Chương 6. Thử nghiệm và đánh giá

- 6.1 Kiểm tra build, health, database bootstrap và 10 kịch bản chức năng.
- 6.2 Đánh giá kết quả, stub và technical debt.
- Nguồn soạn: `test-scenarios.md`.

## Chương 7. Kết luận

1. Kết quả: marketplace, Offer, Transaction, timeline, ledger prototype và deployment.
2. Giá trị kinh tế tuần hoàn.
3. Hạn chế: nhánh tương thích token demo, Create Demand còn stub, ownership/state/atomicity ở các luồng khác.
4. Hướng phát triển: matching, payment/logistics, security, testing và observability.

Nguồn soạn: `conclusion-notes.md`.

## Kiểm tra độ phủ

Dùng `guideline-coverage.md` để kiểm tra từng mục trong guideline trước khi ghép bản báo cáo cuối.
