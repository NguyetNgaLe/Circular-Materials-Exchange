# Chương 5 - Thực hiện và triển khai

## 5.1. Cấu trúc mã nguồn và thành phần

| Thư mục | Nội dung |
|---|---|
| `stitch-app/` | React SPA, route, page, store và API client |
| `circular-materials-exchange/api-gateway/` | REST route, middleware, HTTP handler và gRPC client |
| `circular-materials-exchange/auth-service/` | Auth gRPC Handler-Service-Repository |
| `circular-materials-exchange/company-service/` | Company gRPC Handler-Service-Repository |
| `circular-materials-exchange/material-service/` | Material gRPC Handler-Service-Repository |
| `circular-materials-exchange/order-service/` | Offer/Transaction/Finance/Escrow gRPC service, NATS publish |
| `circular-materials-exchange/review-service/` | Review gRPC service |
| `circular-materials-exchange/notification-service/` | Notification gRPC service và NATS consumer |
| `circular-materials-exchange/proto/` | Hợp đồng Protocol Buffers |
| `circular-materials-exchange/*/migrations/` | Migration theo service |
| `circular-materials-exchange/scripts/init-databases.sql` | Bootstrap 6 database và 18 bảng |
| `circular-materials-exchange/scripts/migrate-existing.sql` | Nâng cấp idempotent volume database hiện có |
| `circular-materials-exchange/docker-compose.yml` | PostgreSQL, NATS, MinIO, 6 service và Gateway |

## 5.2. Môi trường triển khai

### Frontend server

- Nginx phục vụ bản build Vite.
- Client-side route dùng SPA fallback.
- `/api/` được reverse proxy sang API Gateway.
- Frontend không nằm trong Docker Compose hiện tại.

### Backend Docker server

- API Gateway: HTTP port 8085, publish `8085:8085`, tham gia `cme-network`.
- 6 gRPC service: 50051-50056.
- PostgreSQL 15: host port 5433, 6 database logic.
- NATS: 4222 và monitoring 8222.
- MinIO: 9000 và console 9001.
- Persistent volume: `pg-data`, `minio-data`.

## 5.3. Luồng triển khai

1. Cấu hình `DB_PASSWORD`, `JWT_SECRET` và `MINIO_PASSWORD` bằng biến môi trường.
2. Chạy Docker Compose để dựng backend.
3. PostgreSQL thực thi `scripts/init-databases.sql` khi volume mới được tạo.
4. Build frontend bằng `npm run build`.
5. Đưa `dist/` lên frontend server và cấu hình Nginx.
6. Kiểm tra `/health`, `/ready`, trang `/`, upload ảnh, NATS consumer và các API được bảo vệ.

## 5.4. Kiến trúc as-built

Các sơ đồ D11-D12 và D24 thể hiện:

- REST chỉ được expose tại Gateway;
- Gateway gọi 6 service bằng gRPC và không truy cập PostgreSQL;
- Material Service upload ảnh lên MinIO;
- Order Service publish NATS và Notification Service subscribe;
- mỗi repository chỉ truy cập database logic của service mình.

## 5.5. Hạn chế triển khai

- Middleware còn giữ tương thích với demo token đã phát trước khi refactor.
- Chuỗi Offer -> Transaction -> Escrow chưa nằm trong một SQL transaction duy nhất.
- Finance/escrow/withdrawal là ledger nội bộ, không có ngân hàng.
- Chưa có HTTPS end-to-end, mTLS nội bộ, monitoring tập trung, backup job và CI/CD chính thức.
