-- Permission table for internal admin panel permissions
-- These permissions are used by the Identity-Admin frontend panel
CREATE TABLE IF NOT EXISTS "permission"."permissions" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "tag" VARCHAR(100) NOT NULL UNIQUE, -- e.g., "addUser", "deleteUser", "viewUsers"
  "name" VARCHAR(255) NOT NULL, -- Human-readable name, e.g., "Add User"
  "description" TEXT,
  "category" VARCHAR(50), -- e.g., "user", "role", "permission", "system"
  "is_system" BOOLEAN NOT NULL DEFAULT false, -- System permissions cannot be deleted
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_permissions_tag" ON "permission"."permissions"("tag");
CREATE INDEX IF NOT EXISTS "idx_permissions_category" ON "permission"."permissions"("category");
CREATE INDEX IF NOT EXISTS "idx_permissions_deleted_at" ON "permission"."permissions"("deleted_at");

