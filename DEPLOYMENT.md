# Circular Materials Exchange - Deployment Guide

## Kiến trúc tổng thể

```
┌─────────────────────────────────────────────────────────────────┐
│                    Browser (cùng mạng LAN)                      │
└──────────────────────────┬──────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│         Frontend - Server FRONTEND_IP:13005                     │
│         React + Vite + serve (SPA mode)                         │
│         Files: /tmp/stitch-app/                                 │
└──────────────────────────┬──────────────────────────────────────┘
                           │ API calls (http://BACKEND_IP:8085)
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│         Backend - Server BACKEND_IP                            │
│                                                                 │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │              API Gateway (:8085)                          │  │
│  │         Go + Gin │ JWT Auth │ CORS │ gRPC Proxy          │  │
│  └──────┬───────┬───────┬───────┬───────┬───────┬──────────┘  │
│         │       │       │       │       │       │              │
│         ▼       ▼       ▼       ▼       ▼       ▼              │
│      ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐        │
│      │Auth │ │Comp-│ │Mate-│ │Order│ │Revi-│ │Notif│        │
│      │:5051│ │:5052│ │:5053│ │:5054│ │:5055│ │:5056│        │
│      └─────┘ └─────┘ └─────┘ └─────┘ └─────┘ └─────┘        │
│         │       │       │       │       │       │              │
│         ▼       ▼       ▼       ▼       ▼       ▼              │
│      ┌─────────────────────────────────────────────────┐      │
│      │   PostgreSQL (:5433) - 6 databases               │      │
│      │   auth_db, company_db, material_db,              │      │
│      │   order_db, review_db, notif_db                  │      │
│      └─────────────────────────────────────────────────┘      │
│      ┌─────────────────────────────────────────────────┐      │
│      │   NATS JetStream (:4222) - Async messaging       │      │
│      └─────────────────────────────────────────────────┘      │
└─────────────────────────────────────────────────────────────────┘
```

---

## Server thông tin

| Server | IP | SSH | Vai trò |
|--------|-----|-----|---------|
| Backend | <SERVER_BACKEND_IP> | <USER> | API Gateway + 6 Microservices + PostgreSQL + NATS |
| Frontend | <SERVER_FRONTEND_IP> | <USER> | React SPA (serve) |

---

## Backend Deployment (Server BACKEND_IP)

### Cấu trúc thư mục

```
/home/ubuntu/circular-materials-exchange/
├── api-gateway/           # Go + Gin HTTP Gateway
│   ├── cmd/server/main.go
│   ├── internal/
│   │   ├── handler/       # HTTP handlers (auth, company, material, order, review, notification)
│   │   ├── middleware/     # JWT auth, CORS
│   │   └── proxy/         # gRPC client connections
│   ├── Dockerfile
│   └── go.mod
├── auth-service/          # Authentication (gRPC :50051)
├── company-service/       # Company CRUD (gRPC :50052)
├── material-service/      # Materials & Listings (gRPC :50053)
├── order-service/         # Offers & Transactions (gRPC :50054)
├── review-service/        # Reviews & Ratings (gRPC :50055)
├── notification-service/  # Notifications (gRPC :50056)
├── scripts/
│   └── init-databases.sql # Tạo 6 databases
├── docker-compose.yml
└── proto/                 # gRPC proto definitions
```

### Docker Compose

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15
    container_name: cme-postgres
    environment:
      POSTGRES_USER: cme
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: auth_db
    ports:
      - "5433:5432"
    volumes:
      - ./scripts/init-databases.sql:/docker-entrypoint-initdb.d/init.sql
      - pg-data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U cme"]
      interval: 5s
      timeout: 5s
      retries: 5

  nats:
    image: nats:latest
    container_name: cme-nats
    command: "-js"
    ports:
      - "4222:4222"
      - "8222:8222"

  # 6 microservices chạy qua Docker
  auth-service:
    build: ./auth-service
    container_name: cme-auth-service
    ports: ["50051:50051"]
    environment:
      DB_HOST: postgres
      DB_PORT: "5432"
      DB_NAME: auth_db
      DB_USER: cme
      DB_PASSWORD: ${DB_PASSWORD}
      JWT_SECRET: cme_jwt_secret_2024
    depends_on:
      postgres: { condition: service_healthy }

  # ... (company, material, order, review, notification services tương tự)

volumes:
  pg-data:
```

### API Gateway (chạy trực tiếp trên host, không qua Docker)

API Gateway chạy trực tiếp trên host vì cần kết nối đến các services qua Docker network.

```bash
# Build
cd /home/ubuntu/circular-materials-exchange/api-gateway
go build -buildvcs=false -o /tmp/api-gateway ./cmd/server

# Run
HTTP_PORT=8085 nohup /tmp/api-gateway > /tmp/api-gw.log 2>&1 &
```

### Database Schema

PostgreSQL container chạy 6 databases riêng biệt:

| Database | Tables | Purpose |
|----------|--------|---------|
| auth_db | users | Đăng ký, đăng nhập, JWT |
| company_db | companies | Quản lý doanh nghiệp |
| material_db | categories, supply_listings, demand_listings | Vật liệu, nhu cầu mua |
| order_db | offers, transactions, transaction_events | Đề nghị mua, giao dịch |
| review_db | reviews | Đánh giá |
| notif_db | notifications | Thông báo |

### Firewall Rules

```bash
# Mở các ports trên server 32
sudo iptables -I INPUT -p tcp --dport 8080 -j ACCEPT   # Service cũ
sudo iptables -I INPUT -p tcp --dport 8085 -j ACCEPT   # API Gateway
sudo iptables -I INPUT -p tcp --dport 50051 -j ACCEPT  # auth-service
sudo iptables -I INPUT -p tcp --dport 50052 -j ACCEPT  # company-service
sudo iptables -I INPUT -p tcp --dport 50053 -j ACCEPT  # material-service
sudo iptables -I INPUT -p tcp --dport 50054 -j ACCEPT  # order-service
sudo iptables -I INPUT -p tcp --dport 50055 -j ACCEPT  # review-service
sudo iptables -I INPUT -p tcp --dport 50056 -j ACCEPT  # notification-service
sudo iptables -I INPUT -p tcp --dport 5433 -j ACCEPT   # PostgreSQL
sudo iptables -I INPUT -p tcp --dport 4222 -j ACCEPT   # NATS
```

### Kiểm tra trạng thái

```bash
# Docker containers
docker ps | grep cme

# API Gateway health
curl http://localhost:8085/health

# Test API
curl -X POST http://localhost:8085/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@cme.vn","password":"admin123"}'
```

---

## Frontend Deployment (Server FRONTEND_IP)

### Cấu trúc

```
/tmp/stitch-app/
├── index.html
├── assets/
│   ├── index-*.js    # React bundle
│   └── index-*.css   # Styles
└── ...
```

### Tech Stack

- React 18.2 + TypeScript
- Vite 5.x (build tool)
- React Router 6.14 (client-side routing)
- Lucide React (icons)

### API Configuration

File `stitch-app/src/services/api.ts`:

```typescript
const API_BASE_URL = 'http://${SERVER_BACKEND_IP}:8085/api';

class ApiService {
  private token: string | null = null;

  setToken(token: string) {
    this.token = token;
    localStorage.setItem('token', token);
  }

  getToken(): string | null {
    if (!this.token) {
      this.token = localStorage.getItem('token');
    }
    return this.token;
  }

  private async request(endpoint: string, options: RequestInit = {}) {
    const token = this.getToken();
    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
      ...((options.headers as Record<string, string>) || {}),
    };
    if (token) {
      headers['Authorization'] = `Bearer ${token}`;
    }
    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
      ...options,
      headers,
    });
    const data = await response.json();
    if (!response.ok) {
      throw new Error(data.message || 'Request failed');
    }
    return data;
  }

  // Auth
  async login(email: string, password: string) {
    return this.request('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    });
  }

  // ... (các methods khác)
}

export const api = new ApiService();
```

### State Management

File `stitch-app/src/store/index.ts`:

- Sử dụng custom store pattern (không dùng Redux)
- Tất cả methods là async, gọi API backend
- Lưu token vào localStorage
- Lưu user info vào localStorage

```typescript
export function useStore() {
  return {
    get currentUser() { return _currentUser },

    async login(email: string, password: string): Promise<boolean> {
      const res = await api.login(email, password);
      if (res.success) {
        api.setToken(res.data.token);
        _currentUser = res.data.user;
        localStorage.setItem('user', JSON.stringify(_currentUser));
        return true;
      }
      return false;
    },

    async getListings(params?: any) {
      const res = await api.getListings(params);
      return res.success ? res.data.listings || [] : [];
    },

    // ... (các methods khác)
  }
}
```

### Build & Deploy

```bash
# Local build
cd stitch-app
npm install
npm run build

# Upload to server 33
# Upload dist/ folder to /tmp/stitch-app/ trên server 33

# Start serve trên server 33
pkill -f serve
nohup serve -s /tmp/stitch-app -l 13005 > /tmp/serve.log 2>&1 &
```

### Serve Mode

Sử dụng `serve` với flag `-s` (SPA mode) để trả về `index.html` cho mọi routes:

```bash
serve -s /tmp/stitch-app -l 13005
```

Flag `-s` rất quan trọng vì React Router dùng client-side routing. Không có flag này, truy cập `/login` hoặc `/marketplace` sẽ trả về 404.

---

## API Endpoints

### Auth (Public)
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/auth/login | Đăng nhập |
| POST | /api/auth/register | Đăng ký |

### Protected (cần JWT token)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/auth/me | Thông tin user hiện tại |
| GET | /api/categories | Danh mục vật liệu |
| GET | /api/listings | Danh sách nguồn cung |
| POST | /api/listings | Đăng nguồn cung mới |
| GET | /api/listings/:id | Chi tiết nguồn cung |
| GET | /api/demands | Danh sách nhu cầu mua |
| POST | /api/demands | Đăng nhu cầu mua mới |
| GET | /api/offers | Danh sách đề nghị |
| POST | /api/offers | Tạo đề nghị mua |
| POST | /api/offers/:id/accept | Chấp nhận đề nghị |
| POST | /api/offers/:id/reject | Từ chối đề nghị |
| GET | /api/transactions | Danh sách giao dịch |
| GET | /api/transactions/:id | Chi tiết giao dịch |
| POST | /api/transactions/:id/status | Cập nhật trạng thái |
| GET | /api/companies | Danh sách doanh nghiệp |
| POST | /api/companies | Tạo doanh nghiệp |
| POST | /api/companies/:id/approve | Admin duyệt doanh nghiệp |
| GET | /api/reviews | Đánh giá |
| POST | /api/reviews | Tạo đánh giá |
| GET | /api/notifications | Thông báo |
| PUT | /api/notifications/:id/read | Đánh dấu đã đọc |

---

## Seed Data

### Users (password: admin123)
| Email | Name | Role |
|-------|------|------|
| admin@cme.vn | Admin | admin |
| an@ecopoly.vn | Nguyen Van An | business |
| binh@greenpack.vn | Tran Thi Binh | business |

### Material Categories
| ID | Name | Icon |
|----|------|------|
| cat1 | Nhựa | recycling |
| cat2 | Kim Loại | hardware |
| cat3 | Giấy & Bìa Cứng | description |
| cat4 | Gỗ | forest |
| cat5 | Dệt May | checkroom |
| cat6 | Thủy Tinh | local_drink |

---

## Troubleshooting

### Frontend trắng xóa
- Kiểm tra F12 → Console có lỗi JS không
- Đảm bảo serve chạy với flag `-s` (SPA mode)
- Hard refresh: Ctrl+Shift+R

### API 401 Unauthorized
- Kiểm tra token trong localStorage
- Đảm bảo login thành công trước khi gọi API khác

### API Gateway không chạy
```bash
# Kiểm tra process
ps aux | grep api-gateway

# Kill và restart
pkill -f /tmp/api-gateway
HTTP_PORT=8085 nohup /tmp/api-gateway > /tmp/api-gw.log 2>&1 &
```

### Docker containers không chạy
```bash
# Kiểm tra trạng thái
docker ps -a | grep cme

# Restart tất cả
cd /home/ubuntu/circular-materials-exchange
docker-compose down
docker-compose up -d
```

### Port bị chiếm
```bash
# Kiểm tra ai đang dùng port
ss -tlnp | grep <port>

# Kill process
fuser -k <port>/tcp
```

---

## Lưu ý quan trọng

1. **Thanh toán là bypass** - `payment_status: 'bypassed_demo'`, không tích hợp payment gateway
2. **OTP giả lập** - Không có SMS OTP thật
3. **Admin duyệt thủ công** - Không có auto-verify
4. **Mỗi service có database riêng** - Tuân theo nguyên tắc Database per Service
5. **gRPC cho internal communication** - Chỉ API Gateway expose REST API
6. **NATS cho async events** - order-service publish events, notification-service consume

---

*Tài liệu cập nhật lần cuối: 09/07/2026*
