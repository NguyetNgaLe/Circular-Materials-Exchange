# Checklist nghiệm thu sơ đồ và báo cáo

## Nội dung chung

- [ ] Mã Dxx và tên hình khớp `03-report/diagram-index.md`.
- [ ] Thuật ngữ thống nhất với `00-reference/system-baseline.md`.
- [ ] UC dùng mã UC001-UC034 từ `03-report/use-case-catalog.md`.
- [ ] Phân biệt `Đã chạy`, `Một phần`, `Chưa có`.
- [ ] Không ghi hệ thống đã tự động matching.
- [ ] Không mô tả bank payment, escrow, withdrawal hoặc logistics như tích hợp thật.
- [ ] Buyer và Seller là vai trò của Business User.
- [ ] Có file source, SVG và PNG; chữ đọc rõ khi đặt trên trang A4.

## Use Case và Activity

- [ ] Activity chỉ mô tả nghiệp vụ, không lẫn chi tiết SQL/gRPC.
- [ ] Actor nằm ngoài biên hệ thống.
- [ ] D08 ghi Update/Hide/Delete Listing đã triển khai ownership; Create Demand vẫn là stub.
- [ ] Không thêm Seller báo giá Demand như chức năng đã có.
- [ ] D09-D10 có finance/escrow/wallet đúng mức prototype.

## Class và ERD

- [ ] Domain Class D04 không chứa Handler/Service/Repository.
- [ ] D04 có các lớp escrow, fee, wallet và withdrawal.
- [ ] D13 có đúng 6 database logic và 18 bảng.
- [ ] Chỉ vẽ FK vật lý trong cùng database; quan hệ xuyên database là logical reference.
- [ ] D13 có `image_url`, `images`, `reference_id`; ghi `specs` tương thích dữ liệu demo.
- [ ] Design Class D14-D16 thể hiện Gateway client, gRPC Handler-Service-Repository và DB ownership.

## Architecture và Deployment

- [ ] Có React SPA, Nginx, API Gateway, 6 gRPC service, PostgreSQL, NATS và MinIO.
- [ ] Frontend và backend nằm trên hai server khác nhau.
- [ ] API Gateway tham gia `cme-network` và publish port 8085.
- [ ] D11-D12 thể hiện Gateway gọi đủ 6 service qua gRPC, không có đường SQL từ Gateway.
- [ ] Không vẽ frontend như Docker service.
- [ ] Order Service publish NATS và Notification Service subscribe queue group.
- [ ] Không đưa các container không thuộc CME trên server vào sơ đồ.

## State và Sequence

- [ ] D22 dùng quy tắc `pending -> accepted/rejected`; ghi Order Service đã enforce `pending` nhưng ownership còn thiếu.
- [ ] D23 dùng luồng UI `confirmed -> in_progress -> completed`.
- [ ] D17-D21 thể hiện REST từ SPA đến Gateway, gRPC tới service và SQL ở repository.
- [ ] D18 có bước upload MinIO và kiểm tra Company verified.
- [ ] D19 có Offer, escrow ledger và Notification.
- [ ] D20 thể hiện Order Service gắn escrow holding vào Transaction.
- [ ] D21 có fee 2%, Seller wallet và Review chủ động.

## Báo cáo

- [ ] Chương 1 dùng bản `.tex` đã cập nhật.
- [ ] Chương 3 có phân rã, Use Case, đặc tả UC và NFR.
- [ ] Chương 4 phân biệt logical architecture và as-built.
- [ ] Chương 5 mô tả hai server và bootstrap 18 bảng.
- [ ] Chương 6 có cả kịch bản đạt và kịch bản chưa đạt.
- [ ] Chương 7 nêu technical debt, không chỉ nêu ưu điểm.
