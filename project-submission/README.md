# Circular Materials Exchange - Project Submission

Thư mục này là khu vực làm bài và đóng gói bài nộp môn Phân tích thiết kế hệ thống. Mã nguồn ứng dụng vẫn nằm ở các thư mục gốc; không sửa hoặc sao chép tùy tiện mã nguồn vào đây.

## Cách sử dụng

1. Đọc `00-reference/system-baseline.md` trước khi vẽ.
2. Xem phần được giao trong `02-team-management/task-assignment.md`.
3. Sửa file nguồn tương ứng trong `04-diagrams`.
4. Xuất đồng thời SVG và PNG, giữ nguyên mã `D01` đến `D24`.
5. Cập nhật `02-team-management/diagram-progress.md`.
6. Lead kiểm tra theo `02-team-management/review-checklist.md` trước khi đưa vào báo cáo.

## Mô hình phân công

Bài có 13 loại sơ đồ và 5 thành viên. Mỗi loại chỉ giao cho một người phụ trách trọn vẹn. Một số loại có nhiều hình con, nhưng không được chia các hình con đó cho nhiều người.

## Quy ước trạng thái

- `TODO`: chưa làm.
- `DOING`: đang làm.
- `REVIEW`: đã xuất hình, chờ Lead kiểm tra.
- `DONE`: đã đạt checklist và có đủ source, SVG, PNG.

## Nguồn sự thật

Sơ đồ phải ưu tiên code hiện tại trong `circular-materials-exchange`, `stitch-app`, migrations, proto và Docker Compose. File `ARCHITECTURE.md` ở gốc mô tả một hệ thống VTrue/VPoint khác nên không được dùng cho bài nộp này.
