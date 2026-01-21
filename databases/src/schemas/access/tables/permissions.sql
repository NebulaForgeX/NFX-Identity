-- Permissions table: defines system capability points
-- Answers: What actions can be performed in the system?
-- Permission points are stable product contracts: system upgrades and new features are essentially adding/deprecating permission points
CREATE TABLE IF NOT EXISTS "access"."permissions" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "key" VARCHAR(255) NOT NULL UNIQUE, -- Stable key: "users.read", "users.count", "users.export", "assets.read", "tenants.members.manage", "clients.credentials.rotate"
  "name" VARCHAR(255) NOT NULL, -- Display name
  "description" TEXT, -- Permission description
  "is_system" BOOLEAN NOT NULL DEFAULT false, -- System built-in, cannot be deleted
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_permissions_key" ON "access"."permissions"("key");
CREATE INDEX IF NOT EXISTS "idx_permissions_is_system" ON "access"."permissions"("is_system");
CREATE INDEX IF NOT EXISTS "idx_permissions_deleted_at" ON "access"."permissions"("deleted_at");

