# PLAN: Hệ Thống Dòng Tiền Thu Chi - Circular Materials Exchange

## 1. Tổng Quan

### Mô hình kinh doanh
- **Sàn giao dịch** đóng vai trò trung gian giữa Buyer và Seller
- **Thu phí giao dịch**: Tỷ lệ % trên mỗi giao dịch thành công
- **Phí đăng tin**: (Tùy chọn) Phí cho listing premium/featured

### Luồng tiền
```
Buyer → [Thanh toán] → Seller
                 ↓
         [Phí giao dịch] → Sàn (Platform)
```

---

## 2. Sửa Lỗi: Không Được Tự Mua Hàng Của Chính Mình

### Vấn đề
Hiện tại `CreateOffer` không kiểm tra `buyer_id != seller_id`

### Giải pháp
Thêm validation trong `order.go` - `CreateOffer` function

```go
// Kiểm tra không tự mua hàng của mình
if userID == req.SellerID {
    c.JSON(http.StatusBadRequest, gin.H{
        "success": false, 
        "message": "Khong the tu mua hang cua chinh minh",
    })
    return
}
```

---

## 3. Database Design - Bảng Mới

### 3.1 Bảng `platform_fees` (Phí giao dịch)
```sql
CREATE TABLE platform_fees (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID REFERENCES transactions(id),
    seller_id UUID NOT NULL,
    buyer_id UUID NOT NULL,
    transaction_amount DECIMAL(15,2) NOT NULL,  -- Tổng tiền giao dịch
    fee_rate DECIMAL(5,4) DEFAULT 0.0200,        -- Tỷ lệ phí (2%)
    fee_amount DECIMAL(15,2) NOT NULL,           -- Số tiền phí
    fee_type VARCHAR(20) DEFAULT 'transaction',  -- transaction, listing_premium
    status VARCHAR(20) DEFAULT 'pending',        -- pending, collected, refunded
    collected_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_platform_fees_seller ON platform_fees(seller_id);
CREATE INDEX idx_platform_fees_status ON platform_fees(status);
CREATE INDEX idx_platform_fees_created ON platform_fees(created_at);
```

### 3.2 Bảng `platform_wallet` (Ví sàn)
```sql
CREATE TABLE platform_wallet (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    balance DECIMAL(15,2) DEFAULT 0,             -- Số dư hiện tại
    total_income DECIMAL(15,2) DEFAULT 0,        -- Tổng thu
    total_expense DECIMAL(15,2) DEFAULT 0,       -- Tổng chi
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### 3.3 Bảng `wallet_transactions` (Lịch sử giao dịch ví)
```sql
CREATE TABLE wallet_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    wallet_id UUID REFERENCES platform_wallet(id),
    type VARCHAR(20) NOT NULL,                   -- income, expense, refund
    amount DECIMAL(15,2) NOT NULL,
    reference_type VARCHAR(50),                  -- platform_fee, withdrawal, etc
    reference_id UUID,                           -- ID của fee/transaction liên quan
    description TEXT,
    balance_after DECIMAL(15,2),                 -- Số dư sau giao dịch
    created_at TIMESTAMP DEFAULT NOW()
);
```

### 3.4 Bảng `company_wallet` (Ví doanh nghiệp) - Tùy chọn
```sql
CREATE TABLE company_wallet (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id UUID REFERENCES companies(id),
    balance DECIMAL(15,2) DEFAULT 0,
    total_earned DECIMAL(15,2) DEFAULT 0,        -- Tổng tiền bán
    total_spent DECIMAL(15,2) DEFAULT 0,         -- Tổng tiền mua
    total_fees_paid DECIMAL(15,2) DEFAULT 0,     -- Tổng phí đã trả
    updated_at TIMESTAMP DEFAULT NOW()
);
```

---

## 4. Business Logic

### 4.1 Khi giao dịch hoàn thành (Transaction status = 'completed')
```
1. Tính phí: fee_amount = transaction_amount × fee_rate
2. Tạo bản ghi platform_fees
3. Cập nhật platform_wallet:
   - balance += fee_amount
   - total_income += fee_amount
4. Tạo wallet_transactions (type: 'income')
5. (Tùy chọn) Cập nhật company_wallet cho seller:
   - total_fees_paid += fee_amount
```

### 4.2 Các loại phí
| Loại phí | Tỷ lệ | Mô tả |
|----------|-------|-------|
| Phí giao dịch | 2% | Thu từ seller khi giao dịch hoàn thành |
| Phí đăng tin premium | 50,000đ/tin | Tin đăng được ưu tiên hiển thị |
| Phí thành viên | - | (Tương lai) Phí duy trì tài khoản |

### 4.3 Quy tắc tính phí
- Phí chỉ tính khi `transaction.status = 'completed'`
- Phí tính trên `agreed_price × quantity`
- Phí do **Seller** chịu (trừ vào tiền nhận)

---

## 5. API Endpoints

### 5.1 Finance APIs (Admin only)
```
GET    /api/admin/finance/overview       -- Tổng quan tài chính
GET    /api/admin/finance/fees           -- Danh sách phí đã thu
GET    /api/admin/finance/wallet         -- Thông tin ví sàn
GET    /api/admin/finance/transactions   -- Lịch sử giao dịch ví
POST   /api/admin/finance/withdraw       -- Rút tiền (mock)
GET    /api/admin/finance/reports        -- Báo cáo theo tháng/quý
```

### 5.2 Validation APIs
```
POST   /api/offers                       -- Thêm kiểm tra buyer != seller
```

---

## 6. Frontend Pages

### 6.1 Admin Finance Dashboard (`/admin/finance`)
- **Stats cards**: Tổng doanh thu, Phí tháng này, Số giao dịch, Số dư ví
- **Biểu đồ**: Doanh thu theo tháng (line chart)
- **Bảng**: Giao dịch phí gần đây
- **Nút**: Xuất báo cáo CSV

### 6.2 Admin Finance Detail (`/admin/finance/fees`)
- Danh sách tất cả phí đã thu
- Filter theo ngày, trạng thái
- Chi tiết từng phí

---

## 7. Implementation Steps

### Phase 1: Sửa lỗi (Ưu tiên cao)
- [ ] Thêm validation `buyer_id != seller_id` trong CreateOffer
- [ ] Thêm validation trong Frontend (ẩn nút mua nếu là owner)

### Phase 2: Database
- [ ] Tạo migration cho 4 bảng mới
- [ ] Seed data mẫu

### Phase 3: Backend
- [ ] Tính phí tự động khi transaction completed
- [ ] API tổng quan tài chính
- [ ] API danh sách phí
- [ ] API báo cáo

### Phase 4: Frontend
- [ ] Admin Finance Dashboard
- [ ] Biểu đồ doanh thu
- [ ] Bảng chi tiết phí

---

## 8. Cấu hình phí

```go
// Trong code
const (
    DefaultFeeRate     = 0.02  // 2%
    PremiumListingFee  = 50000 // 50,000 VND
    MinTransactionFee  = 1000  // 1,000 VND (phí tối thiểu)
)
```

---

## 9. Mở rộng tương lai
- Ví doanh nghiệp (nạp/rút)
- Hệ thống thanh toán online (VNPay, MoMo)
- Hóa đơn điện tử
- Báo cáo tài chính chi tiết
- Phí theo membership tier
