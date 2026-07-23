CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Bang escrow (giu tien)
CREATE TABLE IF NOT EXISTS escrow_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID REFERENCES transactions(id),
    buyer_id UUID NOT NULL,
    buyer_name VARCHAR(255),
    seller_id UUID NOT NULL,
    seller_name VARCHAR(255),
    amount DECIMAL(15,2) NOT NULL,
    fee_rate DECIMAL(5,4) DEFAULT 0.0200,
    fee_amount DECIMAL(15,2) NOT NULL,
    seller_amount DECIMAL(15,2) NOT NULL,
    status VARCHAR(20) DEFAULT 'holding',
    hold_until TIMESTAMP NOT NULL,
    released_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Bang vi seller
CREATE TABLE IF NOT EXISTS seller_wallet (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    seller_id UUID UNIQUE NOT NULL,
    seller_name VARCHAR(255),
    balance DECIMAL(15,2) DEFAULT 0,
    total_earned DECIMAL(15,2) DEFAULT 0,
    total_fees_paid DECIMAL(15,2) DEFAULT 0,
    total_withdrawn DECIMAL(15,2) DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Bang lich su vi seller
CREATE TABLE IF NOT EXISTS seller_wallet_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    seller_id UUID NOT NULL,
    type VARCHAR(20) NOT NULL,
    amount DECIMAL(15,2) NOT NULL,
    balance_after DECIMAL(15,2),
    reference_type VARCHAR(50),
    reference_id UUID,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Bang yeu cau rut tien
CREATE TABLE IF NOT EXISTS withdrawal_requests (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    seller_id UUID NOT NULL,
    seller_name VARCHAR(255),
    amount DECIMAL(15,2) NOT NULL,
    bank_name VARCHAR(100),
    bank_account VARCHAR(50),
    bank_owner VARCHAR(255),
    status VARCHAR(20) DEFAULT 'pending',
    admin_note TEXT,
    processed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_escrow_status ON escrow_transactions(status);
CREATE INDEX IF NOT EXISTS idx_escrow_seller ON escrow_transactions(seller_id);
CREATE INDEX IF NOT EXISTS idx_escrow_hold_until ON escrow_transactions(hold_until);
CREATE INDEX IF NOT EXISTS idx_seller_wallet_seller ON seller_wallet(seller_id);
CREATE INDEX IF NOT EXISTS idx_seller_wallet_tx_seller ON seller_wallet_transactions(seller_id);
CREATE INDEX IF NOT EXISTS idx_withdrawals_seller ON withdrawal_requests(seller_id);
CREATE INDEX IF NOT EXISTS idx_withdrawals_status ON withdrawal_requests(status);
