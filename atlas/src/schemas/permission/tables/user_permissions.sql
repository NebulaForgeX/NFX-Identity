-- User-Permission relationship table
-- Links users from auth.users to permissions in permission.permissions
-- Note: user_id references auth.users, so this is a cross-schema reference
CREATE TABLE IF NOT EXISTS "permission"."user_permissions" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID NOT NULL, -- References auth.users(id) - cross-schema reference
  "permission_id" UUID NOT NULL REFERENCES "permission"."permissions"("id") ON DELETE CASCADE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP,
  UNIQUE("user_id", "permission_id")
);

CREATE INDEX IF NOT EXISTS "idx_user_permissions_user_id" ON "permission"."user_permissions"("user_id");
CREATE INDEX IF NOT EXISTS "idx_user_permissions_permission_id" ON "permission"."user_permissions"("permission_id");
CREATE INDEX IF NOT EXISTS "idx_user_permissions_deleted_at" ON "permission"."user_permissions"("deleted_at");

-- Add foreign key constraint to auth.users (cross-schema)
-- Note: This requires the auth schema to exist first
-- ALTER TABLE "permission"."user_permissions" 
--   ADD CONSTRAINT "fk_user_permissions_user_id" 
--   FOREIGN KEY ("user_id") REFERENCES "auth"."users"("id") ON DELETE CASCADE;

