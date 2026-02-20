-- One row per role per tenant. When a tenant is created, insert exactly 4 rows (owner, admin, editor, viewer).
-- tenant_id references tenants.tenants.id (application-level; access loads before tenants)
CREATE TABLE IF NOT EXISTS "access"."tenant_roles" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "tenant_id" UUID NOT NULL,
  "role_key" VARCHAR(32) NOT NULL, -- 'owner' | 'admin' | 'editor' | 'viewer'
  "name" VARCHAR(255), -- Display name
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("tenant_id", "role_key")
);

CREATE INDEX IF NOT EXISTS "idx_tenant_roles_tenant_id" ON "access"."tenant_roles"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_tenant_roles_role_key" ON "access"."tenant_roles"("role_key");
