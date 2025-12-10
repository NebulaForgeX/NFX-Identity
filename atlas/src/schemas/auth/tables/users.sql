-- User table for authentication and user management
CREATE TYPE "auth".user_status AS ENUM ('pending', 'active', 'deactive');

CREATE TABLE IF NOT EXISTS "auth"."users" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "username" VARCHAR(50) NOT NULL UNIQUE,
  "email" VARCHAR(255) NOT NULL UNIQUE,
  "phone" VARCHAR(20) NOT NULL UNIQUE,
  "password_hash" VARCHAR(255) NOT NULL,
  "role_id" UUID REFERENCES "auth"."roles"("id") ON DELETE SET NULL,
  "status" "auth".user_status NOT NULL DEFAULT 'pending',
  "is_verified" BOOLEAN NOT NULL DEFAULT false,
  "last_login_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_users_email" ON "auth"."users"("email");
CREATE INDEX IF NOT EXISTS "idx_users_phone" ON "auth"."users"("phone");
CREATE INDEX IF NOT EXISTS "idx_users_username" ON "auth"."users"("username");
CREATE INDEX IF NOT EXISTS "idx_users_role_id" ON "auth"."users"("role_id");
CREATE INDEX IF NOT EXISTS "idx_users_status" ON "auth"."users"("status");
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "auth"."users"("deleted_at");

