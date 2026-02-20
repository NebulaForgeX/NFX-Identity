CREATE TABLE IF NOT EXISTS "access"."super_admins" (
  "user_id" UUID PRIMARY KEY,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_super_admins_user_id" ON "access"."super_admins"("user_id");
