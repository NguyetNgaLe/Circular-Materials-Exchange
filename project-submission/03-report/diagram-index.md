# Danh mục sơ đồ

## Chủ sở hữu theo 13 loại

| STT | Loại | Người phụ trách |
|---:|---|---:|
| 1 | Cơ cấu tổ chức | 1 - Lead |
| 2 | Activity/BPMN | 2 |
| 3 | Domain Class | 2 |
| 4 | Functional Decomposition/Package | 1 - Lead |
| 5 | Use Case tổng quan | 1 - Lead |
| 6 | Use Case theo nhóm | 3 |
| 7 | Logical Architecture | 3 |
| 8 | Deployment | 3 |
| 9 | ERD | 4 |
| 10 | Design Class | 4 |
| 11 | Sequence | 5 |
| 12 | State Machine | 2 |
| 13 | Docker/server | 5 |

Mỗi loại có đúng một người phụ trách. Một loại có thể gồm nhiều hình con nhưng không chia các hình con cho nhiều người.

## Danh mục hình

| Mã | Chương | Tên sơ đồ | File nguồn | Người |
|---|---:|---|---|---:|
| D01 | 2 | Cơ cấu tổ chức nghiệp vụ | `D01-organization-structure.mmd` | 1 |
| D02 | 2 | Activity đăng ký và duyệt doanh nghiệp | `D02-company-approval-activity.puml` | 2 |
| D03 | 2 | Activity nguồn cung, Offer, giao dịch và ledger | `D03-circular-transaction-activity.puml` | 2 |
| D04 | 2 | Domain Class gồm finance/escrow/wallet | `D04-domain-class.puml` | 2 |
| D05 | 3 | Functional Decomposition/Package | `D05-functional-decomposition.puml` | 1 |
| D06 | 3 | Use Case tổng quan | `D06-overview-use-case.puml` | 1 |
| D07 | 3 | Use Case tài khoản và doanh nghiệp | `D07-auth-company-use-case.puml` | 3 |
| D08 | 3 | Use Case marketplace, upload và Demand | `D08-marketplace-use-case.puml` | 3 |
| D09 | 3 | Use Case Offer, Transaction, wallet và Review | `D09-transaction-review-use-case.puml` | 3 |
| D10 | 3 | Use Case Notification, finance và Admin | `D10-admin-notification-use-case.puml` | 3 |
| D11 | 4 | Logical Architecture microservice as-built | `D11-logical-architecture.dsl` | 3 |
| D12 | 4 | Deployment Diagram hai server | `D12-deployment.puml` | 3 |
| D13 | 4 | ERD 6 database, 18 bảng | `D13-erd.dbml` | 4 |
| D14 | 4 | Design Class Auth/Company microservice as-built | `D14-auth-company-design-class.puml` | 4 |
| D15 | 4 | Design Class Material/Upload microservice as-built | `D15-material-design-class.puml` | 4 |
| D16 | 4 | Design Class Order/Finance/Escrow/Review/Notification | `D16-order-support-design-class.puml` | 4 |
| D17 | 4 | Sequence tạo và duyệt doanh nghiệp | `D17-company-approval-sequence.puml` | 5 |
| D18 | 4 | Sequence upload ảnh và đăng nguồn cung | `D18-create-listing-sequence.puml` | 5 |
| D19 | 4 | Sequence tìm kiếm và gửi Offer | `D19-send-offer-sequence.puml` | 5 |
| D20 | 4 | Sequence chấp nhận Offer và tạo Transaction | `D20-accept-offer-sequence.puml` | 5 |
| D21 | 4 | Sequence giao hàng, hoàn tất, release ledger và Review | `D21-complete-review-sequence.puml` | 5 |
| D22 | 4 | State Machine Offer | `D22-offer-state.mmd` | 2 |
| D23 | 4 | State Machine Transaction as-built | `D23-transaction-state.mmd` | 2 |
| D24 | 5 | Mô hình triển khai hai server và Docker | `D24-docker-server-model.puml` | 5 |

## Quy tắc caption

- D11 phải ghi rõ REST Gateway, gRPC service và database ownership.
- D13 phải có `image_url`, `images`, `reference_id` và ghi chú tương thích `specs`.
- D19-D21 phải gọi escrow/fee/wallet là ledger prototype.
- D22-D23 phải ghi backend chưa enforce state transition đầy đủ.
