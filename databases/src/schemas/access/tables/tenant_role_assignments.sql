-- Who has which role in which tenant. One row per (user, tenant); role lives only here (same schema as tenant_roles).
-- user_id = directory.users.id, tenant_id = tenants.tenants.id (application-level; no cross-schema FK)
CREATE TABLE IF NOT EXISTS "access"."tenant_role_assignments" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID NOT NULL,
  "tenant_id" UUID NOT NULL,
  "tenant_role_id" UUID NOT NULL, -- References access.tenant_roles.id (must belong to this tenant_id)
  "assigned_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "assigned_by" UUID,
  UNIQUE("user_id", "tenant_id")
);

CREATE INDEX IF NOT EXISTS "idx_tenant_role_assignments_user_id" ON "access"."tenant_role_assignments"("user_id");
CREATE INDEX IF NOT EXISTS "idx_tenant_role_assignments_tenant_id" ON "access"."tenant_role_assignments"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_tenant_role_assignments_tenant_role_id" ON "access"."tenant_role_assignments"("tenant_role_id");
