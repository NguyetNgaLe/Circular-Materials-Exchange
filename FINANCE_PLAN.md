# PLAN: Hệ Thống Dòng Tiền Thu Chi - Circular Materials Exchange

## 1. Tổng Quan

### Mô hình kinh doanh
- **Sàn giao dịch** đóng vai trò trung gian giữa Buyer và Seller
- **Thu phí giao dịch**: 2% trên mỗi giao dịch thành công
- **Mô hình Escrow**: Giữ tiền tạm thời, giải ngân khi giao dịch hoàn tất

### Luồng tiền (Shopee-style)
```
Buyer thanh toán → Escrow (giữ tiền)
        ↓
Seller giao hàng → Buyer xác nhận nhận hàng
        ↓
Escrow giải ngân:
  - 98% → Ví Seller
  - 2% → Platform (phí giao dịch)
```

---

## 2. Luồng Giao Dịch Chi Tiết

### 2.1 Buyer gửi đề nghị mua
```
1. Buyer chọn sản phẩm → Click "Gửi Đề Nghị Mua"
2. Nhập số lượng, giá đề xuất
3. Click "Thanh Toán & Gửi Đề Nghị"
4. Popup xác nhận thanh toán
5. Hệ thống:
   - Tạo Offer (status: pending)
   - Tạo Escrow (giữ tiền, status: holding)
   - Gửi thông báo cho Seller
6. Popup "Bill Succeed" hiển thị
```

### 2.2 Seller chấp nhận đề nghị
```
1. Seller xem đề nghị → Click "Chấp nhận"
2. Hệ thống:
   - Cập nhật Offer (status: accepted)
   - Tạo Transaction (status: confirmed)
   - Gửi thông báo cho Buyer
```

### 2.3 Seller giao hàng
```
1. Seller click "Xác Nhận Đã Giao Hàng"
2. Transaction status: confirmed → in_progress
3. Gửi thông báo cho Buyer
```

### 2.4 Buyer nhận hàng & Hoàn tất
```
1. Buyer click "Xác Nhận Đã Nhận Hàng & Hoàn Tất"
2. Transaction status: in_progress → completed
3. Hệ thống tự động:
   - Giải ngân Escrow
   - 98% → Ví Seller (seller_amount)
   - 2% → Platform wallet (fee_amount)
   - Tạo bản ghi platform_fees
   - Tạo lịch sử ví Seller
   - Gửi thông báo cho Seller
```

---

## 3. Database Design

### 3.1 Bảng `escrow_transactions` (Quỹ giữ tiền)
```sql
CREATE TABLE escrow_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID,                    -- NULL khi chưa có transaction
    buyer_id UUID NOT NULL,
    buyer_name VARCHAR(255),
    seller_id UUID NOT NULL,
    seller_name VARCHAR(255),
    amount DECIMAL(15,2) NOT NULL,          -- Tổng tiền
    fee_rate DECIMAL(5,4) DEFAULT 0.0200,   -- Tỷ lệ phí (2%)
    fee_amount DECIMAL(15,2) NOT NULL,      -- Tiền phí
    seller_amount DECIMAL(15,2) NOT NULL,   -- Tiền seller nhận
    status VARCHAR(20) DEFAULT 'holding',   -- holding, released
    hold_until TIMESTAMP NOT NULL,          -- Thời gian giữ tối đa
    released_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 3.2 Bảng `platform_wallet` (Ví sàn)
```sql
CREATE TABLE platform_wallet (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    balance DECIMAL(15,2) DEFAULT 0,
    total_income DECIMAL(15,2) DEFAULT 0,
    total_expense DECIMAL(15,2) DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 3.3 Bảng `platform_fees` (Phí giao dịch)
```sql
CREATE TABLE platform_fees (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID,
    seller_id UUID NOT NULL,
    buyer_id UUID NOT NULL,
    transaction_amount DECIMAL(15,2) NOT NULL,
    fee_rate DECIMAL(5,4) DEFAULT 0.0200,
    fee_amount DECIMAL(15,2) NOT NULL,
    fee_type VARCHAR(20) DEFAULT 'transaction',
    status VARCHAR(20) DEFAULT 'collected',
    collected_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 3.4 Bảng `seller_wallet` (Ví seller)
```sql
CREATE TABLE seller_wallet (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    seller_id UUID UNIQUE NOT NULL,
    seller_name VARCHAR(255),
    balance DECIMAL(15,2) DEFAULT 0,
    total_earned DECIMAL(15,2) DEFAULT 0,
    total_fees_paid DECIMAL(15,2) DEFAULT 0,
    total_withdrawn DECIMAL(15,2) DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 3.5 Bảng `seller_wallet_transactions` (Lịch sử ví seller)
```sql
CREATE TABLE seller_wallet_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    seller_id UUID NOT NULL,
    type VARCHAR(20) NOT NULL,              -- credit, debit
    amount DECIMAL(15,2) NOT NULL,
    balance_after DECIMAL(15,2),
    reference_type VARCHAR(50),
    reference_id UUID,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 3.6 Bảng `withdrawal_requests` (Yêu cầu rút tiền)
```sql
CREATE TABLE withdrawal_requests (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    seller_id UUID NOT NULL,
    seller_name VARCHAR(255),
    amount DECIMAL(15,2) NOT NULL,
    bank_name VARCHAR(100),
    bank_account VARCHAR(50),
    bank_owner VARCHAR(255),
    status VARCHAR(20) DEFAULT 'pending',   -- pending, completed, rejected
    admin_note TEXT,
    processed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

## 4. Business Logic

### 4.1 Khi Buyer gửi đề nghị mua (CreateOffer)
```
1. Kiểm tra buyer != seller
2. Kiểm tra DN buyer đã được duyệt
3. Tạo offer (status: pending)
4. Tạo escrow:
   - amount = quantity × proposed_price
   - fee_amount = amount × 2%
   - seller_amount = amount - fee_amount
   - status = holding
   - hold_until = now + 3 days
5. Tạo notification cho seller
```

### 4.2 Khi Seller chấp nhận đề nghị (AcceptOffer)
```
1. Cập nhật offer status = accepted
2. Tạo transaction (status: confirmed)
3. Tạo transaction_event
4. Gửi notification cho buyer
```

### 4.3 Khi Seller giao hàng (UpdateTransactionStatus → in_progress)
```
1. Cập nhật transaction status = in_progress
2. Tạo transaction_event
3. Gửi notification cho buyer
```

### 4.4 Khi Buyer xác nhận nhận hàng (UpdateTransactionStatus → completed)
```
1. Cập nhật transaction status = completed
2. Tạo transaction_event
3. Giải ngân escrow:
   - Cập nhật escrow status = released
   - Tạo platform_fees (fee_amount)
   - Cập nhật platform_wallet (balance += fee_amount)
   - Cập nhật seller_wallet (balance += seller_amount)
   - Tạo seller_wallet_transactions
4. Gửi notification cho seller
```

### 4.5 Các loại phí
| Loại phí | Tỷ lệ | Mô tả |
|----------|-------|-------|
| Phí giao dịch | 2% | Thu từ seller khi giao dịch hoàn thành |

### 4.6 Quy tắc tính phí
- Phí chỉ tính khi `transaction.status = 'completed'`
- Phí tính trên `quantity × agreed_price`
- Phí do **Seller** chịu (trừ vào tiền nhận)
- Tiền được giữ tạm trong escrow cho đến khi giao dịch hoàn tất

---

## 5. API Endpoints

### 5.1 Offer APIs
```
POST   /api/offers                  -- Tạo đề nghị (+ tạo escrow)
GET    /api/offers                  -- Danh sách đề nghị
POST   /api/offers/:id/accept       -- Chấp nhận (+ tạo transaction)
POST   /api/offers/:id/reject       -- Từ chối
```

### 5.2 Transaction APIs
```
GET    /api/transactions            -- Danh sách giao dịch
GET    /api/transactions/:id        -- Chi tiết giao dịch
POST   /api/transactions/:id/status -- Cập nhật trạng thái (+ giải ngân nếu completed)
```

### 5.3 Seller Wallet APIs
```
GET    /api/seller/wallet                -- Thông tin ví seller
GET    /api/seller/wallet/transactions   -- Lịch sử ví
POST   /api/seller/withdraw              -- Yêu cầu rút tiền
GET    /api/seller/withdrawals           -- Lịch sử rút tiền
```

### 5.4 Admin Finance APIs
```
GET    /api/admin/finance/overview       -- Tổng quan tài chính
GET    /api/admin/finance/fees           -- Danh sách phí đã thu
GET    /api/admin/finance/wallet         -- Thông tin ví sàn
```

### 5.5 Admin Escrow APIs
```
GET    /api/admin/escrow                 -- Danh sách escrow
POST   /api/admin/escrow/:id/release    -- Giải ngân thủ công
```

### 5.6 Admin Withdrawal APIs
```
GET    /api/admin/withdrawals            -- Danh sách yêu cầu rút tiền
POST   /api/admin/withdrawals/:id/approve -- Duyệt rút tiền
POST   /api/admin/withdrawals/:id/reject  -- Từ chối rút tiền
```

---

## 6. Frontend Pages

### 6.1 Buyer Flow
| Trang | Chức năng |
|-------|-----------|
| `/marketplace` | Duyệt sản phẩm |
| `/material/:id` | Xem chi tiết + "Gửi Đề Nghị Mua" |
| `/offers/new/:id` | Form gửi đề nghị + Popup thanh toán + Bill Succeed |
| `/offers/sent` | Xem đề nghị đã gửi |
| `/transactions/:id` | Xác nhận nhận hàng & Hoàn tất |

### 6.2 Seller Flow
| Trang | Chức năng |
|-------|-----------|
| `/offers/received` | Xem đề nghị đã nhận + Chấp nhận/Từ chối |
| `/transactions/:id` | Xác nhận giao hàng |
| `/wallet` | Xem ví + Lịch sử giao dịch |

### 6.3 Admin Flow
| Trang | Chức năng |
|-------|-----------|
| `/admin/finance` | Dashboard tài chính (doanh thu, phí, giao dich) |
| `/admin/escrow` | Quản lý escrow (tiền đang giữ, giải ngân) |

---

## 7. Trạng Thái Giao Dịch

```
pending (offer) → accepted → confirmed → in_progress → completed
                                                      → cancelled
```

| Trạng thái | Mô tả | Escrow |
|------------|-------|--------|
| pending | Offer mới tạo | holding |
| accepted | Seller chấp nhận | holding |
| confirmed | Transaction được tạo | holding |
| in_progress | Seller giao hàng | holding |
| completed | Buyer xác nhận nhận hàng | released → ví seller |

---

## 8. Ví Dụ Cụ Thể

**GreenPack mua Nhua PE cua EcoPoly: 30,000,000đ**

| Bước | Hành động | Kết quả |
|------|-----------|---------|
| 1 | Buyer gửi đề nghị | Escrow: 30,000,000đ (holding) |
| 2 | Seller chấp nhận | Transaction: confirmed |
| 3 | Seller giao hàng | Transaction: in_progress |
| 4 | Buyer xác nhận | Transaction: completed |
| 5 | Giải ngân | Seller: +29,400,000đ / Platform: +600,000đ |

---

## 9. Cấu Hình Phí

```go
const (
    DefaultFeeRate = 0.02  // 2%
    HoldDays       = 3     // Giữ tiền 3 ngày
)
```

---

## 10. Mở Rộng Tương Lai
- Hệ thống thanh toán online (VNPay, MoMo)
- Hóa đơn điện tử
- Báo cáo tài chính chi tiết
- Phí theo membership tier
- Phí đăng tin premium
