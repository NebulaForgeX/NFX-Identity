-- Role Permissions table: Many-to-Many relationship between roles and permissions
-- Answers: What can tenant.admin actually do?
-- Enterprise needs "version evolution": when a permission is added to admin role, it must be traceable
CREATE TABLE IF NOT EXISTS "access"."role_permissions" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "role_id" UUID NOT NULL REFERENCES "access"."roles"("id") ON DELETE CASCADE,
  "permission_id" UUID NOT NULL REFERENCES "access"."permissions"("id") ON DELETE CASCADE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID, -- Who added this permission to the role (for audit)
  UNIQUE("role_id", "permission_id")
);

CREATE INDEX IF NOT EXISTS "idx_role_permissions_role_id" ON "access"."role_permissions"("role_id");
CREATE INDEX IF NOT EXISTS "idx_role_permissions_permission_id" ON "access"."role_permissions"("permission_id");
CREATE INDEX IF NOT EXISTS "idx_role_permissions_role_permission" ON "access"."role_permissions"("role_id", "permission_id");

