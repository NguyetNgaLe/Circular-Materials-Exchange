# Circular Materials Exchange - Project Submission

Thư mục này chứa tài liệu làm bài và đóng gói bài nộp môn Phân tích thiết kế hệ thống. Baseline hiện tại được đồng bộ với code ngày 24/07/2026.

## Bắt đầu từ đâu

1. Đọc `00-reference/system-baseline.md`.
2. Xem trạng thái chức năng tại `00-reference/as-built-feature-matrix.md`.
3. Xem phần được giao trong `02-team-management/task-assignment.md`.
4. Đối chiếu vị trí hình tại `03-report/diagram-index.md`.
5. Chỉnh source trong `04-diagrams`, export SVG và PNG.
6. Lead nghiệm thu bằng `02-team-management/review-checklist.md`.

## Tài liệu báo cáo đã chuẩn bị

| Nội dung | File |
|---|---|
| Chương 1 LaTeX | `03-report/latex/chuong-01-dat-van-de.tex` |
| Khung 7 chương | `03-report/report-outline.md` |
| Ma trận đáp ứng guideline | `03-report/guideline-coverage.md` |
| Chương 2 phân tích nghiệp vụ | `03-report/business-analysis.md` |
| UC001-UC034 | `03-report/use-case-catalog.md` |
| Đặc tả UC trọng yếu | `03-report/use-case-specifications.md` |
| Yêu cầu phi chức năng | `03-report/non-functional-requirements.md` |
| Chương 4 kiến trúc và thiết kế | `03-report/architecture-design.md` |
| Chương 5 triển khai | `03-report/implementation-deployment.md` |
| Chương 6 kiểm thử | `03-report/test-scenarios.md` |
| Chương 7 kết luận | `03-report/conclusion-notes.md` |

## Mô hình phân công

Bài có 13 loại sơ đồ và 5 thành viên tính cả Lead. Mỗi loại giao trọn cho đúng một người. Một loại có thể gồm nhiều hình con nhưng không chia hình con cho nhiều người.

## Quy ước trạng thái

- `TODO`: chưa hoàn thiện/export.
- `DOING`: đang làm.
- `REVIEW`: đủ source và hình, chờ Lead.
- `DONE`: đạt checklist.

## Nguồn sự thật

Ưu tiên route/handler/frontend call/proto/migration/Compose hiện tại. Luồng as-built là REST từ SPA tới Gateway, gRPC từ Gateway tới service sở hữu dữ liệu và SQL chỉ nằm trong repository của service. Không dùng tài liệu VTrue/VPoint hoặc tài liệu kế hoạch cũ nếu mâu thuẫn code.
