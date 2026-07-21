CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS offers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    type VARCHAR(20) NOT NULL,
    listing_id UUID NOT NULL,
    listing_title VARCHAR(255),
    buyer_id UUID NOT NULL,
    buyer_name VARCHAR(255),
    seller_id UUID NOT NULL,
    seller_name VARCHAR(255),
    quantity DECIMAL(10,2),
    unit VARCHAR(20),
    proposed_price DECIMAL(15,2),
    currency VARCHAR(10) DEFAULT 'VND',
    message TEXT,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    offer_id UUID REFERENCES offers(id),
    listing_title VARCHAR(255),
    buyer_id UUID NOT NULL,
    buyer_name VARCHAR(255),
    seller_id UUID NOT NULL,
    seller_name VARCHAR(255),
    quantity DECIMAL(10,2),
    unit VARCHAR(20),
    agreed_price DECIMAL(15,2),
    currency VARCHAR(10) DEFAULT 'VND',
    payment_status VARCHAR(20) DEFAULT 'bypassed_demo',
    payment_method VARCHAR(20) DEFAULT 'manual_offline',
    settlement_note TEXT DEFAULT 'Thanh toan duoc thuc hien ngoai he thong trong pham vi prototype',
    status VARCHAR(20) DEFAULT 'confirmed',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transaction_events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID REFERENCES transactions(id),
    actor_id UUID NOT NULL,
    actor_name VARCHAR(255),
    from_status VARCHAR(50),
    to_status VARCHAR(50),
    note TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_offers_buyer ON offers(buyer_id);
CREATE INDEX IF NOT EXISTS idx_offers_seller ON offers(seller_id);
CREATE INDEX IF NOT EXISTS idx_offers_status ON offers(status);
CREATE INDEX IF NOT EXISTS idx_transactions_buyer ON transactions(buyer_id);
CREATE INDEX IF NOT EXISTS idx_transactions_seller ON transactions(seller_id);
CREATE INDEX IF NOT EXISTS idx_transactions_status ON transactions(status);
CREATE INDEX IF NOT EXISTS idx_tx_events_transaction ON transaction_events(transaction_id);
