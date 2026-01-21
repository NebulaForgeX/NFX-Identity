-- IP Allowlist table for source constraints
-- Constrains requests using this client/API key must come from specified IPs/CIDRs
-- Reduces risk of credential leakage
-- Enterprise common: production only allows fixed egress IPs
CREATE TYPE "clients".allowlist_status AS ENUM ('active', 'disabled', 'revoked');

CREATE TABLE IF NOT EXISTS "clients"."ip_allowlist" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "rule_id" VARCHAR(255) NOT NULL UNIQUE, -- Unique rule identifier
  "app_id" UUID NOT NULL REFERENCES "clients"."apps"("id") ON DELETE CASCADE,
  "cidr" CIDR NOT NULL, -- IP range: e.g., 203.0.113.0/24, 2001:db8::/32
  "description" TEXT, -- Rule description: "Production server IPs", "Office network"
  "status" "clients".allowlist_status NOT NULL DEFAULT 'active',
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID, -- Who created this rule
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_by" UUID,
  "revoked_at" TIMESTAMP,
  "revoked_by" UUID,
  "revoke_reason" TEXT
);

CREATE INDEX IF NOT EXISTS "idx_ip_allowlist_rule_id" ON "clients"."ip_allowlist"("rule_id");
CREATE INDEX IF NOT EXISTS "idx_ip_allowlist_app_id" ON "clients"."ip_allowlist"("app_id");
CREATE INDEX IF NOT EXISTS "idx_ip_allowlist_status" ON "clients"."ip_allowlist"("status");
CREATE INDEX IF NOT EXISTS "idx_ip_allowlist_app_status" ON "clients"."ip_allowlist"("app_id", "status") WHERE "status" = 'active';
-- Note: GiST index on cidr is optional. For small-scale deployments (dozens to hundreds of rules),
-- the composite B-tree index (app_id, status) is sufficient. GiST index is only needed for:
-- 1. High-volume CIDR matching queries (IP <<= cidr)
-- 2. Large allowlist datasets (thousands+ rules per app)
-- See: databases/migrations/optional/add_gist_index_for_ip_allowlist.sql for optional migration

