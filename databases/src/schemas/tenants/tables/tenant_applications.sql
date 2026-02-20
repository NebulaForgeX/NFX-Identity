CREATE TYPE "tenants".tenant_application_status AS ENUM ('ACTIVE', 'DISABLED', 'SUSPENDED');

CREATE TABLE IF NOT EXISTS "tenants"."tenant_applications" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "tenant_id" UUID NOT NULL REFERENCES "tenants"."tenants"("id") ON DELETE CASCADE,
  "application_id" UUID NOT NULL,
  "status" "tenants".tenant_application_status NOT NULL DEFAULT 'ACTIVE',
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "settings" JSONB DEFAULT '{}'::jsonb,
  UNIQUE("tenant_id", "application_id")
);

CREATE INDEX IF NOT EXISTS "idx_tenant_applications_tenant_id" ON "tenants"."tenant_applications"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_tenant_applications_application_id" ON "tenants"."tenant_applications"("application_id");
CREATE INDEX IF NOT EXISTS "idx_tenant_applications_status" ON "tenants"."tenant_applications"("status");
