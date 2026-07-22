CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Bảng phí giao dịch
CREATE TABLE IF NOT EXISTS platform_fees (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID REFERENCES transactions(id),
    seller_id UUID NOT NULL,
    buyer_id UUID NOT NULL,
    transaction_amount DECIMAL(15,2) NOT NULL,
    fee_rate DECIMAL(5,4) DEFAULT 0.0200,
    fee_amount DECIMAL(15,2) NOT NULL,
    fee_type VARCHAR(20) DEFAULT 'transaction',
    status VARCHAR(20) DEFAULT 'pending',
    collected_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Ví sàn
CREATE TABLE IF NOT EXISTS platform_wallet (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    balance DECIMAL(15,2) DEFAULT 0,
    total_income DECIMAL(15,2) DEFAULT 0,
    total_expense DECIMAL(15,2) DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Lịch sử giao dịch ví
CREATE TABLE IF NOT EXISTS wallet_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    wallet_id UUID REFERENCES platform_wallet(id),
    type VARCHAR(20) NOT NULL,
    amount DECIMAL(15,2) NOT NULL,
    reference_type VARCHAR(50),
    reference_id UUID,
    description TEXT,
    balance_after DECIMAL(15,2),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Ví doanh nghiệp
CREATE TABLE IF NOT EXISTS company_wallet (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id UUID UNIQUE,
    balance DECIMAL(15,2) DEFAULT 0,
    total_earned DECIMAL(15,2) DEFAULT 0,
    total_spent DECIMAL(15,2) DEFAULT 0,
    total_fees_paid DECIMAL(15,2) DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_platform_fees_seller ON platform_fees(seller_id);
CREATE INDEX IF NOT EXISTS idx_platform_fees_status ON platform_fees(status);
CREATE INDEX IF NOT EXISTS idx_platform_fees_created ON platform_fees(created_at);
CREATE INDEX IF NOT EXISTS idx_wallet_transactions_wallet ON wallet_transactions(wallet_id);
CREATE INDEX IF NOT EXISTS idx_wallet_transactions_type ON wallet_transactions(type);

-- Seed ví sàn
INSERT INTO platform_wallet (id, balance, total_income, total_expense) 
VALUES ('00000000-0000-0000-0000-000000000001', 0, 0, 0)
ON CONFLICT (id) DO NOTHING;
