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

-- Seed categories
INSERT INTO categories (id, name, icon) VALUES
('00000000-0000-0000-0000-000000000001', 'Nhựa', 'recycling'),
('00000000-0000-0000-0000-000000000002', 'Kim Loại', 'hardware'),
('00000000-0000-0000-0000-000000000003', 'Giấy & Bìa Cứng', 'description'),
('00000000-0000-0000-0000-000000000004', 'Gỗ', 'forest'),
('00000000-0000-0000-0000-000000000005', 'Dệt May', 'checkroom'),
('00000000-0000-0000-0000-000000000006', 'Thủy Tinh', 'local_drink')
ON CONFLICT (id) DO NOTHING;
