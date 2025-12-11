-- User-Role relationship table (many-to-many)
-- This allows a user to have multiple roles
CREATE TABLE IF NOT EXISTS "auth"."user_roles" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID NOT NULL REFERENCES "auth"."users"("id") ON DELETE CASCADE,
  "role_id" UUID NOT NULL REFERENCES "auth"."roles"("id") ON DELETE CASCADE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("user_id", "role_id")
);

CREATE INDEX IF NOT EXISTS "idx_user_roles_user_id" ON "auth"."user_roles"("user_id");
CREATE INDEX IF NOT EXISTS "idx_user_roles_role_id" ON "auth"."user_roles"("role_id");

