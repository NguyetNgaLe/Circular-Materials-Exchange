# Khung báo cáo theo guideline

## Chương 1. Đặt vấn đề

- Bài toán thực tế.
- Hiện trạng và nhu cầu.
- Mục tiêu và phạm vi MVP.
- Giải pháp tổng quan.

Nguồn: `01-problem-and-scope`.

## Chương 2. Phân tích nghiệp vụ

- 2.1 Cơ cấu tổ chức: D01.
- 2.2 Quy trình nghiệp vụ: D02, D03.
- 2.3 Các lớp lĩnh vực: D04.

## Chương 3. Phân tích yêu cầu

- 3.1.1 Phân rã chức năng: D05.
- 3.1.1 Use Case tổng quan: D06.
- Các nhóm chức năng: D07-D10.
- 3.1.2 Đặc tả các Use Case quan trọng.
- 3.2 Yêu cầu phi chức năng.

## Chương 4. Thiết kế

- Giải thích lựa chọn kiến trúc microservice và ranh giới 6 service.
- Làm rõ Database per Service ở mức logic dù môi trường demo dùng chung một PostgreSQL container.
- 4.1.1 Kiến trúc logic: D11.
- 4.1.2 Kiến trúc triển khai: D12.
- Thiết kế dữ liệu: D13.
- Thiết kế lớp: D14-D16.
- Thiết kế hành vi: D17-D23.

## Chương 5. Thực hiện và triển khai

- Cấu trúc mã nguồn.
- Thành phần triển khai.
- Docker/server: D24.
- Môi trường chạy và hướng dẫn demo.

## Chương 6. Thử nghiệm và đánh giá

- Kịch bản đăng ký và duyệt doanh nghiệp.
- Kịch bản đăng nguồn cung và gửi Offer.
- Kịch bản chấp nhận Offer, hoàn tất và đánh giá.
- Đánh giá kết quả và giới hạn MVP.

## Chương 7. Kết luận

- Kết quả đạt được.
- Giá trị kinh tế tuần hoàn.
- Hướng phát triển: matching, payment, logistics, verification và dispute.
