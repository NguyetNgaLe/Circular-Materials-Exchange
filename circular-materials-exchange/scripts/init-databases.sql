-- Bootstrap every service database for a fresh Docker volume.
-- This script is executed by the official PostgreSQL image through
-- /docker-entrypoint-initdb.d/init.sql.

-- POSTGRES_DB may have already created auth_db, so create databases
-- conditionally to keep the bootstrap safe to run more than once.
SELECT 'CREATE DATABASE auth_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'auth_db')\gexec

SELECT 'CREATE DATABASE company_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'company_db')\gexec

SELECT 'CREATE DATABASE material_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'material_db')\gexec

SELECT 'CREATE DATABASE order_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'order_db')\gexec

SELECT 'CREATE DATABASE review_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'review_db')\gexec

SELECT 'CREATE DATABASE notif_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'notif_db')\gexec

-- ============================================================================
-- Auth service
-- ============================================================================
\connect auth_db

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'business',
    avatar VARCHAR(500) DEFAULT '',
    company_id UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);

-- Demo accounts use the same bcrypt password as the deployed demo environment.
INSERT INTO users (id, name, email, phone, password_hash, role) VALUES
('a0000000-0000-0000-0000-000000000001', 'Admin', 'admin@cme.vn', '0900000000', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'admin'),
('a0000000-0000-0000-0000-000000000002', 'Nguyễn Văn An', 'an@ecopoly.vn', '0901234567', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'business'),
('a0000000-0000-0000-0000-000000000003', 'Trần Thị Bình', 'binh@greenpack.vn', '0912345678', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'business')
ON CONFLICT (email) DO NOTHING;

-- ============================================================================
-- Company service
-- ============================================================================
\connect company_db

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS companies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    tax_code VARCHAR(50) UNIQUE,
    address TEXT,
    city VARCHAR(100),
    description TEXT,
    status VARCHAR(20) DEFAULT 'pending',
    reject_reason TEXT,
    owner_id UUID NOT NULL,
    rating DECIMAL(3,2) DEFAULT 0,
    review_count INT DEFAULT 0,
    member_since DATE DEFAULT CURRENT_DATE,
    certifications TEXT DEFAULT '',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_companies_owner ON companies(owner_id);
CREATE INDEX IF NOT EXISTS idx_companies_status ON companies(status);

-- ============================================================================
-- Material service
-- ============================================================================
\connect material_db

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    icon VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS supply_listings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    category_id UUID REFERENCES categories(id),
    seller_id UUID NOT NULL,
    company_id UUID,
    description TEXT,
    specs JSONB DEFAULT '{}',
    quantity DECIMAL(10,2),
    unit VARCHAR(20),
    price_per_unit DECIMAL(15,2),
    currency VARCHAR(10) DEFAULT 'VND',
    location VARCHAR(255),
    min_order_quantity DECIMAL(10,2),
    packaging VARCHAR(100),
    status VARCHAR(20) DEFAULT 'active',
    images TEXT DEFAULT '',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS demand_listings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    category_id UUID REFERENCES categories(id),
    buyer_id UUID NOT NULL,
    company_id UUID,
    description TEXT,
    quantity DECIMAL(10,2),
    unit VARCHAR(20),
    target_price DECIMAL(15,2),
    location VARCHAR(255),
    deadline DATE,
    status VARCHAR(20) DEFAULT 'open',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_supply_category ON supply_listings(category_id);
CREATE INDEX IF NOT EXISTS idx_supply_seller ON supply_listings(seller_id);
CREATE INDEX IF NOT EXISTS idx_supply_status ON supply_listings(status);
CREATE INDEX IF NOT EXISTS idx_demand_category ON demand_listings(category_id);
CREATE INDEX IF NOT EXISTS idx_demand_buyer ON demand_listings(buyer_id);

INSERT INTO categories (id, name, icon) VALUES
('00000000-0000-0000-0000-000000000001', 'Nhựa', 'recycling'),
('00000000-0000-0000-0000-000000000002', 'Kim Loại', 'hardware'),
('00000000-0000-0000-0000-000000000003', 'Giấy & Bìa Cứng', 'description'),
('00000000-0000-0000-0000-000000000004', 'Gỗ', 'forest'),
('00000000-0000-0000-0000-000000000005', 'Dệt May', 'checkroom'),
('00000000-0000-0000-0000-000000000006', 'Thủy Tinh', 'local_drink')
ON CONFLICT (id) DO NOTHING;

-- ============================================================================
-- Order, finance, and escrow services
-- ============================================================================
\connect order_db

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

CREATE TABLE IF NOT EXISTS platform_wallet (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    balance DECIMAL(15,2) DEFAULT 0,
    total_income DECIMAL(15,2) DEFAULT 0,
    total_expense DECIMAL(15,2) DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW()
);

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

CREATE TABLE IF NOT EXISTS company_wallet (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    company_id UUID UNIQUE,
    balance DECIMAL(15,2) DEFAULT 0,
    total_earned DECIMAL(15,2) DEFAULT 0,
    total_spent DECIMAL(15,2) DEFAULT 0,
    total_fees_paid DECIMAL(15,2) DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_platform_fees_seller ON platform_fees(seller_id);
CREATE INDEX IF NOT EXISTS idx_platform_fees_status ON platform_fees(status);
CREATE INDEX IF NOT EXISTS idx_platform_fees_created ON platform_fees(created_at);
CREATE INDEX IF NOT EXISTS idx_wallet_transactions_wallet ON wallet_transactions(wallet_id);
CREATE INDEX IF NOT EXISTS idx_wallet_transactions_type ON wallet_transactions(type);

INSERT INTO platform_wallet (id, balance, total_income, total_expense)
VALUES ('00000000-0000-0000-0000-000000000001', 0, 0, 0)
ON CONFLICT (id) DO NOTHING;

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

CREATE INDEX IF NOT EXISTS idx_escrow_status ON escrow_transactions(status);
CREATE INDEX IF NOT EXISTS idx_escrow_seller ON escrow_transactions(seller_id);
CREATE INDEX IF NOT EXISTS idx_escrow_hold_until ON escrow_transactions(hold_until);
CREATE INDEX IF NOT EXISTS idx_seller_wallet_seller ON seller_wallet(seller_id);
CREATE INDEX IF NOT EXISTS idx_seller_wallet_tx_seller ON seller_wallet_transactions(seller_id);
CREATE INDEX IF NOT EXISTS idx_withdrawals_seller ON withdrawal_requests(seller_id);
CREATE INDEX IF NOT EXISTS idx_withdrawals_status ON withdrawal_requests(status);

-- ============================================================================
-- Review service
-- ============================================================================
\connect review_db

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS reviews (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction_id UUID NOT NULL,
    reviewer_id UUID NOT NULL,
    reviewer_name VARCHAR(255),
    reviewee_id UUID NOT NULL,
    reviewee_name VARCHAR(255),
    rating INT CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_reviews_reviewee ON reviews(reviewee_id);
CREATE INDEX IF NOT EXISTS idx_reviews_transaction ON reviews(transaction_id);

-- ============================================================================
-- Notification service
-- ============================================================================
\connect notif_db

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    title VARCHAR(255),
    message TEXT,
    type VARCHAR(20),
    read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_notifications_user ON notifications(user_id);
CREATE INDEX IF NOT EXISTS idx_notifications_read ON notifications(user_id, read);
