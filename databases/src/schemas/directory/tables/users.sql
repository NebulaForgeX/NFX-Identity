-- User table for user directory management
-- Note: Authentication credentials are stored in auth.user_credentials
CREATE TYPE "directory".user_status AS ENUM ('pending', 'active', 'deactive');

CREATE TABLE IF NOT EXISTS "directory"."users" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "username" VARCHAR(50) NOT NULL UNIQUE,
  "status" "directory".user_status NOT NULL DEFAULT 'pending',
  "is_verified" BOOLEAN NOT NULL DEFAULT false,
  "last_login_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
CREATE INDEX IF NOT EXISTS "idx_users_username" ON "directory"."users"("username");
CREATE INDEX IF NOT EXISTS "idx_users_status" ON "directory"."users"("status");
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "directory"."users"("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_users_last_login_at" ON "directory"."users"("last_login_at");

