\set ON_ERROR_STOP on

-- Idempotent migration for deployments that already have a PostgreSQL volume.
-- Fresh deployments are fully created by init-databases.sql.

\connect company_db

ALTER TABLE companies ADD COLUMN IF NOT EXISTS image_url TEXT DEFAULT '';

\connect material_db

ALTER TABLE categories ADD COLUMN IF NOT EXISTS image_url TEXT DEFAULT '';
ALTER TABLE supply_listings ADD COLUMN IF NOT EXISTS images TEXT DEFAULT '';
ALTER TABLE supply_listings ADD COLUMN IF NOT EXISTS image_url TEXT DEFAULT '';

\connect notif_db

ALTER TABLE notifications ADD COLUMN IF NOT EXISTS reference_id UUID;
CREATE UNIQUE INDEX IF NOT EXISTS idx_notifications_reference
    ON notifications(user_id, type, reference_id)
    WHERE reference_id IS NOT NULL;
