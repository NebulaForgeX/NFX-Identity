-- Scope Permissions table: Maps OAuth scopes to internal permissions
-- Answers: Which internal permissions does a scope include?
-- Example: "users:read" scope may include permissions: users.read, users.count, users.detail
CREATE TABLE IF NOT EXISTS "access"."scope_permissions" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "scope" VARCHAR(255) NOT NULL REFERENCES "access"."scopes"("scope") ON DELETE CASCADE,
  "permission_id" UUID NOT NULL REFERENCES "access"."permissions"("id") ON DELETE CASCADE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("scope", "permission_id")
);

CREATE INDEX IF NOT EXISTS "idx_scope_permissions_scope" ON "access"."scope_permissions"("scope");
CREATE INDEX IF NOT EXISTS "idx_scope_permissions_permission_id" ON "access"."scope_permissions"("permission_id");
CREATE INDEX IF NOT EXISTS "idx_scope_permissions_scope_permission" ON "access"."scope_permissions"("scope", "permission_id");

