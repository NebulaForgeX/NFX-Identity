-- Tenants table: Main entity for tenants (companies/organizations)
-- Each Company A/B/C in the system is a tenant
-- Enterprise responsibilities:
--   - Isolation boundary: users, applications, assets, audits are scoped to a tenant
--   - Lifecycle: create, activate, suspend, close, delete (usually soft delete)
--   - Billing/contract (future): plan, quota, expiration (can be reserved for future)
CREATE TYPE "tenants".tenant_status AS ENUM ('ACTIVE', 'SUSPENDED', 'CLOSED', 'PENDING');

CREATE TABLE IF NOT EXISTS "tenants"."tenants" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "tenant_id" VARCHAR(255) NOT NULL UNIQUE, -- Public tenant identifier
  "name" VARCHAR(255) NOT NULL, -- Tenant name
  "display_name" VARCHAR(255), -- Display name
  "owner_id" UUID NOT NULL, -- Creator (references directory.users.id); gets tenant role owner
  "status" "tenants".tenant_status NOT NULL DEFAULT 'PENDING',
  "primary_domain" VARCHAR(255), -- Optional: companyA.com, used for email domain restriction/SSO
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP,
  "metadata" JSONB DEFAULT '{}'::jsonb -- Flexible extension: {"plan": "...", "quota": {...}, "contract_expires_at": "...", ...}
);

CREATE INDEX IF NOT EXISTS "idx_tenants_tenant_id" ON "tenants"."tenants"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_tenants_name" ON "tenants"."tenants"("name");
CREATE INDEX IF NOT EXISTS "idx_tenants_status" ON "tenants"."tenants"("status");
CREATE INDEX IF NOT EXISTS "idx_tenants_primary_domain" ON "tenants"."tenants"("primary_domain") WHERE "primary_domain" IS NOT NULL;
CREATE INDEX IF NOT EXISTS "idx_tenants_deleted_at" ON "tenants"."tenants"("deleted_at");

