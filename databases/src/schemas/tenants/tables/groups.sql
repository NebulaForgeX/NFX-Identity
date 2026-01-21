-- Groups table: Organizational structure (minimum implementation)
-- Enterprise requirement: "Can we assign permissions by department? Can we grant permissions to a team?"
-- Supports: departments, teams, groups with optional tree structure
CREATE TYPE "tenants".group_type AS ENUM ('department', 'team', 'group', 'other');

CREATE TABLE IF NOT EXISTS "tenants"."groups" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "group_id" VARCHAR(255) NOT NULL UNIQUE, -- Public group identifier
  "tenant_id" UUID NOT NULL REFERENCES "tenants"."tenants"("id") ON DELETE CASCADE,
  "name" VARCHAR(255) NOT NULL,
  "type" "tenants".group_type NOT NULL DEFAULT 'group',
  "parent_group_id" UUID REFERENCES "tenants"."groups"("id") ON DELETE SET NULL, -- Optional: parent group for tree structure (self-reference)
  "description" TEXT,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "deleted_at" TIMESTAMP,
  "metadata" JSONB DEFAULT '{}'::jsonb
);

CREATE INDEX IF NOT EXISTS "idx_groups_group_id" ON "tenants"."groups"("group_id");
CREATE INDEX IF NOT EXISTS "idx_groups_tenant_id" ON "tenants"."groups"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_groups_parent_group_id" ON "tenants"."groups"("parent_group_id");
CREATE INDEX IF NOT EXISTS "idx_groups_type" ON "tenants"."groups"("type");
CREATE INDEX IF NOT EXISTS "idx_groups_deleted_at" ON "tenants"."groups"("deleted_at");

