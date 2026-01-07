-- Roles table: named wrappers for permission collections
-- Answers: What positions/identities exist?
-- Roles are for human readability and operational convenience
-- Enterprise customers often need "custom roles" (or modify based on templates)
CREATE TYPE "access".scope_type AS ENUM ('TENANT', 'APP', 'GLOBAL');

CREATE TABLE IF NOT EXISTS "access"."roles" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "key" VARCHAR(255) NOT NULL UNIQUE, -- Unique role key: "tenant.owner", "tenant.admin", "tenant.viewer", "app.operator", "app.support", "service.reader", "service.writer"
  "name" VARCHAR(255) NOT NULL, -- Display name
  "description" TEXT, -- Role description
  "scope_type" "access".scope_type NOT NULL DEFAULT 'TENANT', -- Where this role is effective: TENANT / APP / GLOBAL
  "is_system" BOOLEAN NOT NULL DEFAULT false, -- System built-in role
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_roles_key" ON "access"."roles"("key");
CREATE INDEX IF NOT EXISTS "idx_roles_scope_type" ON "access"."roles"("scope_type");
CREATE INDEX IF NOT EXISTS "idx_roles_is_system" ON "access"."roles"("is_system");
CREATE INDEX IF NOT EXISTS "idx_roles_deleted_at" ON "access"."roles"("deleted_at");

