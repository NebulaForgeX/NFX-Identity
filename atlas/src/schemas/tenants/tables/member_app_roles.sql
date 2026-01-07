-- Member App Roles table: App-level member permissions
-- Enterprise common: In the same company, some employees can only manage a specific product/app
-- Refines permissions to app level
CREATE TABLE IF NOT EXISTS "tenants"."member_app_roles" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "member_id" UUID NOT NULL REFERENCES "tenants"."members"("member_id") ON DELETE CASCADE,
  "app_id" UUID NOT NULL, -- References clients.apps.id (application-level consistency)
  "role_id" UUID NOT NULL, -- References access.roles.id (application-level consistency)
  "assigned_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "assigned_by" UUID,
  "expires_at" TIMESTAMP, -- Optional: temporary app-level permissions
  "revoked_at" TIMESTAMP,
  "revoked_by" UUID,
  "revoke_reason" TEXT,
  UNIQUE("member_id", "app_id", "role_id")
);

CREATE INDEX IF NOT EXISTS "idx_member_app_roles_member_id" ON "tenants"."member_app_roles"("member_id");
CREATE INDEX IF NOT EXISTS "idx_member_app_roles_app_id" ON "tenants"."member_app_roles"("app_id");
CREATE INDEX IF NOT EXISTS "idx_member_app_roles_role_id" ON "tenants"."member_app_roles"("role_id");
CREATE INDEX IF NOT EXISTS "idx_member_app_roles_expires_at" ON "tenants"."member_app_roles"("expires_at") WHERE "expires_at" IS NOT NULL;

