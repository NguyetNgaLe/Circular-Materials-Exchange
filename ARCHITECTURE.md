# VTrue / VPoint — Microservices Architecture

## Tổng quan

VTrue là nền tảng loyalty & gamification cho ứng dụng fintech Việt Nam. Kiến trúc microservices với API Gateway làm entry point, gRPC cho communication nội bộ, NATS cho async events.

```
┌─────────────────────────────────────────────────────────────────┐
│                        Mobile App                               │
│                   (Kotlin Multiplatform)                        │
└──────────────────────────┬──────────────────────────────────────┘
                           │ HTTPS / WSS
                           ▼
┌─────────────────────────────────────────────────────────────────┐
│                      API GATEWAY (:8900)                        │
│  Go 1.25 │ JWT Auth │ Rate Limit │ CORS │ WS Proxy             │
└──────┬───────┬───────┬───────┬───────┬───────┬──────────────────┘
       │       │       │       │       │       │
       ▼       ▼       ▼       ▼       ▼       ▼
    ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐ ┌─────┐ ┌─────────┐
    │Auth │ │Event│ │Loyal│ │Vouch│ │Notif│ │Streaming│
    │Core │ │Svc  │ │Core │ │Svc  │ │Core │ │ Core    │
    └─────┘ └─────┘ └─────┘ └─────┘ └─────┘ └─────────┘
```

---

## Danh sách Services (10 services)

| # | Service | Language | HTTP Port | gRPC Port | Database | Purpose |
|---|---------|----------|-----------|-----------|----------|---------|
| 1 | **api-gateway** | Go 1.25 | :8900→:8080 | — | Redis | HTTP gateway, JWT auth, rate limit, WS proxy |
| 2 | **auth-core** | Go 1.26 | — | :8101 | PostgreSQL | User auth, OTP, JWT, PIN, devices |
| 3 | **event-service** | Go 1.25 | :8200 | :8201 | PostgreSQL | Gamification: live, missions, checkin, wheel, leaderboard, gift box |
| 4 | **loyalty-api** (vpoint-core) | Go | — | :8081 | — (proxy) | BFF cho loyalty-core gRPC |
| 5 | **loyalty-core** (vpoint-core) | C++ | — | :50052 | PostgreSQL | Financial engine: wallet, ledger, point lots, double-entry accounting |
| 6 | **voucher-service** (vmarket-core) | Go | — | :9400 | PostgreSQL | Voucher marketplace, brands, purchases |
| 7 | **vpoint-campaign** (vmarket-core) | Go | — | :8951 | PostgreSQL | Campaign engine, earn rules |
| 8 | **notification-core** | Go 1.25 | — | :8501 | PostgreSQL | Notifications (in-app, push, email) |
| 9 | **payment-core-vnpay** | Go 1.26 | — | :50052 | PostgreSQL | VNPay topup integration |
| 10 | **streaming-core** | Go 1.25 | :8300 | — | PostgreSQL + Redis | Real-time WS: comments, likes, viewer tracking |
| 11 | **chat-core/bridge** | Haskell | — | :8800 | — | E2E encrypted chat WebSocket bridge |
| 12 | **brand-daemon** | Python | — | — | — | Brand data sync |

---

## Communication Patterns

### 1. REST API (Client → Gateway)

```
Mobile App ──HTTPS──▶ API Gateway (:8900) ──gRPC──▶ Backend Services
```

**Response Envelope:**
```json
{
  "success": true,
  "message": "OK",
  "code": "SUCCESS",
  "request_id": "uuid",
  "data": { ... }
}
```

### 2. gRPC (Service ↔ Service)

Internal communication giữa services sử dụng gRPC:

```
API Gateway ──gRPC──▶ auth-core (:8101)
API Gateway ──gRPC──▶ event-service (:8201)
loyalty-api ──gRPC──▶ loyalty-core (:50052)
API Gateway ──gRPC──▶ voucher-service (:9400)
API Gateway ──gRPC──▶ notification-core (:8501)
```

**Proto definitions:** `proto/gamification.proto` (event-service)

### 3. NATS JetStream (Async Events)

```
┌──────────────────────────────────────────────────────────┐
│                    NATS JetStream                         │
└──────────────────────────────────────────────────────────┘
       ▲                    ▲                    ▲
       │                    │                    │
  event-service      loyalty-core        voucher-service
  (Publisher)         (Publisher)           (Publisher)
       │                    │                    │
       ▼                    ▼                    ▼
  reward-worker      event-service       event-service
  (Consumer)          (Consumer)           (Consumer)
```

**NATS Subjects:**

| Subject | Source | Consumer | Purpose |
|---------|--------|----------|---------|
| `vpoint.events.rewards.claims` | event-service | reward-worker | Async reward claims (all types) |
| `vpoint.events.points.earned` | loyalty-core | event-service | Points earned events |
| `vpoint.events.voucher.purchased` | voucher-service | event-service | Voucher purchase events |
| `payment.topup.success` | payment-core | event-service | Topup success events |

**Pattern:** Transactional outbox → Relay worker → JetStream publish → Durable subscription with manual ack

### 4. WebSocket (Real-time)

```
Mobile App ──WSS──▶ API Gateway ──WS Proxy──▶ streaming-core (:8300)
Mobile App ──WSS──▶ API Gateway ──WS Proxy──▶ chat-bridge (:8800)
```

**Streaming Protocol (JSON):**
```
Client → Server:
  {"type":"comment","data":{"client_message_id":"uuid","content":"Hello!"}}
  {"type":"like","data":{}}
  {"type":"ping","data":{}}

Server → Client:
  {"type":"connected","data":{"user_id":"...","live_id":"..."}}
  {"type":"comment_new","data":{"id","user_id","display_name","content","created_at"}}
  {"type":"like_update","data":{"like_count":343,"delta":28}}
  {"type":"viewer_update","data":{"viewer_count":1288}}
```

---

## Middleware Stack (API Gateway)

```
Request → RequestID → RealIP → Recovery → Compress → Timeout(skip WS)
        → Security → CORS → Logger → RateLimit(IP) → JWT Auth → RateLimit(User)
        → Handler → Response
```

---

## Service Details

### api-gateway
- Entry point cho tất cả client requests
- JWT authentication & token validation
- IP-based & user-based rate limiting
- WebSocket proxy đến backend services
- CORS, compression, request timeout

### auth-core
- User registration, login (OTP, password)
- JWT token issuance & refresh
- PIN management
- Device management & tracking

### event-service
- **Gamification hub**: live sessions, missions, daily checkin, wheel spin, leaderboard, gift box
- **Reward system**: Points claim → NATS → reward-worker → loyalty-api Earn
- **Live streaming backend**: live_sessions, live_engagement, live_comments tables

### loyalty-core (C++)
- **Financial engine**: double-entry accounting
- Wallet management, ledger, point lots
- High-performance C++ for transaction processing

### streaming-core
- **WebSocket hub**: per-live-ID rooms, fan-out broadcast
- **Comment buffer**: in-memory → batch INSERT (1s interval, 200 batch max)
- **Like counter**: Redis INCR → broadcast → sync DB
- **Viewer tracker**: Redis ZSET with timestamps
- **Spam filter**: rate limit, duplicate detection

---

## Database Architecture

Each service owns its database (Database per Service pattern):

```
┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│  auth_db    │  │  event_db   │  │ loyalty_db  │
│ (PostgreSQL)│  │ (PostgreSQL)│  │ (PostgreSQL)│
└─────────────┘  └─────────────┘  └─────────────┘

┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│ voucher_db  │  │ notif_db    │  │ streaming_db│
│ (PostgreSQL)│  │ (PostgreSQL)│  │ (PostgreSQL)│
└─────────────┘  └─────────────┘  │ + Redis     │
                                   └─────────────┘
```

---

## Docker Networks

```
┌─────────────────────────────────────────────────────────┐
│  Network: vpoint                                         │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐   │
│  │api-gateway│ │auth-core │ │event-vpoint│ │loyalty-api│  │
│  └──────────┘ └──────────┘ └──────────┘ └──────────┘   │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐                │
│  │voucher-svc│ │notif-core│ │payment-core│               │
│  └──────────┘ └──────────┘ └──────────┘                │
└─────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────┐
│  Network: deploy_streaming-net                           │
│  ┌──────────────┐ ┌─────────────┐ ┌─────────────┐      │
│  │streaming-core │ │streaming-db │ │streaming-redis│     │
│  └──────────────┘ └─────────────┘ └─────────────┘      │
│  ┌─────────────┐                                        │
│  │streaming-nats│                                        │
│  └─────────────┘                                        │
└─────────────────────────────────────────────────────────┘

api-gateway kết nối CẢ HAI mạng để proxy đến streaming-core
```

---

## Reward Flow (End-to-End)

```
User ClaimReward
    │
    ▼
MarkRewardClaimed (FOR UPDATE lock)
    │
    ├─── if points ──▶ gamification balance + asyncSvc.CreateClaim()
    │                        │
    │                        ▼
    │                   NATS publish (vpoint.events.rewards.claims)
    │                        │
    │                        ▼
    │                   RewardWorker → loyalty-api.Earn()
    │                        │
    │                        ▼
    │                   Wallet credited
    │
    └─── if voucher ──▶ buildVoucherReward() [PLACEHOLDER]
```

---

## Build & Run

| Service | Build | Run |
|---------|-------|-----|
| api-gateway | `go build -o bin/api-gateway ./cmd/server` | `./bin/api-gateway` |
| event-service | `go build -o bin/event-vpoint ./cmd/server` | `./bin/event-vpoint` |
| streaming-core | `go build -o bin/streaming-core ./cmd/server` | `./bin/streaming-core` |
| kotlin_app | `./gradlew :androidApp:assembleDebug` | Install APK |

---

## Lint & Type Check

| Service | Command |
|---------|---------|
| All Go services | `go vet ./...` |
| kotlin_app | `./gradlew check` |

---

## Deployment

- **Server**: `171.244.52.69` (SSH port: 10225, user: ngaltn)
- **Base URL**: `http://171.244.52.69:8900`
- **CI/CD**: GitLab CI/CD for api-gateway (test → build → deploy)
- **Video Storage**: MinIO at `171.244.52.69:9000` (bucket: vpoint-live)
