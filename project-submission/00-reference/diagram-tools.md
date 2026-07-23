# Công cụ và quy tắc xuất 13 loại sơ đồ

| STT | Loại sơ đồ | Chương | Công cụ chính | File nguồn hiện có | Hình |
|---:|---|---:|---|---|---|
| 1 | Cơ cấu tổ chức | 2 | Mermaid Flowchart | `.mmd` | D01 |
| 2 | Activity Diagram/BPMN | 2 | PlantUML Activity | `.puml` | D02-D03 |
| 3 | Domain Class Diagram | 2 | PlantUML Class | `.puml` | D04 |
| 4 | Functional Decomposition/Package | 3 | PlantUML Package | `.puml` | D05 |
| 5 | Use Case tổng quan | 3 | PlantUML Use Case | `.puml` | D06 |
| 6 | Use Case theo nhóm | 3 | PlantUML Use Case | `.puml` | D07-D10 |
| 7 | Logical Architecture | 4 | Structurizr DSL/C4 | `.dsl` | D11 |
| 8 | Deployment Diagram | 4 | PlantUML Deployment | `.puml` | D12 |
| 9 | ERD | 4 | dbdiagram.io/DBML | `.dbml` | D13 |
| 10 | Design Class Diagram | 4 | PlantUML Class | `.puml` | D14-D16 |
| 11 | Sequence Diagram | 4 | PlantUML Sequence | `.puml` | D17-D21 |
| 12 | State Machine Diagram | 4 | Mermaid `stateDiagram-v2` | `.mmd` | D22-D23 |
| 13 | Mô hình Docker/server | 5 | PlantUML Deployment; có thể dàn trang lại bằng diagrams.net | `.puml` | D24 |

Sequence Diagram dùng PlantUML thay cho Mermaid vì luồng có nhiều `alt`, database và service. D24 giữ PlantUML làm nguồn có thể version-control; nếu cần hình trình bày đẹp hơn, Người 5 có thể dựng lại bằng diagrams.net nhưng nội dung không được khác D24.

## Link công cụ

- Mermaid Live Editor: <https://mermaid.live/>
- PlantUML Server: <https://www.plantuml.com/plantuml/>
- Structurizr: <https://structurizr.com/>
- dbdiagram.io: <https://dbdiagram.io/>
- diagrams.net: <https://app.diagrams.net/>
- bpmn.io khi giảng viên yêu cầu BPMN chuẩn: <https://bpmn.io/>

## Chuẩn trình bày

- Tên actor, use case, trạng thái và mã UC phải thống nhất giữa hình và báo cáo.
- Mỗi hình có tiêu đề tiếng Việt và mã D01-D24.
- Ưu tiên nền trắng, màu xanh lá/xanh dương đồng bộ với giao diện.
- Không chụp màn hình editor; phải dùng chức năng Export.
- Xuất SVG để chèn báo cáo và PNG độ rộng tối thiểu 2000 px để dự phòng.
- Giữ file nguồn; không sửa trực tiếp file ảnh xuất.
- Caption phải phân biệt chức năng đang chạy, chức năng một phần và kiến trúc mục tiêu.

## Kiểm tra trước khi export

1. Đối chiếu `system-baseline.md` và `as-built-feature-matrix.md`.
2. Kiểm tra owner trong `task-assignment.md`; một loại chỉ có một owner.
3. Render source, sửa lỗi cú pháp và kiểm tra chữ không bị chồng/cắt.
4. Export SVG/PNG vào `99-final-export/` theo đúng mã Dxx.
5. Chuyển trạng thái sang `REVIEW`; chỉ Lead được chốt `DONE`.
