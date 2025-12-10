-- Role table for user roles and permissions
CREATE TABLE IF NOT EXISTS "auth"."roles" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" VARCHAR(50) NOT NULL UNIQUE,
  "description" TEXT,
  "permissions" JSONB DEFAULT '[]'::jsonb,
  "is_system" BOOLEAN NOT NULL DEFAULT false,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_roles_name" ON "auth"."roles"("name");
CREATE INDEX IF NOT EXISTS "idx_roles_deleted_at" ON "auth"."roles"("deleted_at");

