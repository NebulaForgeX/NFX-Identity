-- Tenant Apps table: Relationship between tenant and applications (client/app)
-- Determines: Which apps can Company A manage/use, and which tenant these app configurations belong to
-- Enterprise common usage:
--   - One tenant has multiple apps (Web, Mobile, API, Backend, etc.)
--   - App belongs to tenant, but app permissions/scopes, callbacks, security policies must be strongly bound to tenant
--   - Future "permission segmentation by app" also needs this table
CREATE TYPE "tenants".tenant_app_status AS ENUM ('ACTIVE', 'DISABLED', 'SUSPENDED');

CREATE TABLE IF NOT EXISTS "tenants"."tenant_apps" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "tenant_id" UUID NOT NULL REFERENCES "tenants"."tenants"("id") ON DELETE CASCADE,
  "app_id" UUID NOT NULL, -- References clients.apps.id (application-level consistency)
  "status" "tenants".tenant_app_status NOT NULL DEFAULT 'ACTIVE',
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID, -- Who created this relationship
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "settings" JSONB DEFAULT '{}'::jsonb, -- Optional: settings for this app in this tenant, e.g., login policy
  UNIQUE("tenant_id", "app_id")
);

CREATE INDEX IF NOT EXISTS "idx_tenant_apps_tenant_id" ON "tenants"."tenant_apps"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_tenant_apps_app_id" ON "tenants"."tenant_apps"("app_id");
CREATE INDEX IF NOT EXISTS "idx_tenant_apps_status" ON "tenants"."tenant_apps"("status");

