-- Member Groups table: Relationship between members and groups
-- Represents which groups/departments/teams a member belongs to
CREATE TABLE IF NOT EXISTS "tenants"."member_groups" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "member_id" UUID NOT NULL REFERENCES "tenants"."members"("member_id") ON DELETE CASCADE,
  "group_id" UUID NOT NULL REFERENCES "tenants"."groups"("id") ON DELETE CASCADE,
  "assigned_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "assigned_by" UUID,
  "revoked_at" TIMESTAMP,
  "revoked_by" UUID,
  UNIQUE("member_id", "group_id")
);

CREATE INDEX IF NOT EXISTS "idx_member_groups_member_id" ON "tenants"."member_groups"("member_id");
CREATE INDEX IF NOT EXISTS "idx_member_groups_group_id" ON "tenants"."member_groups"("group_id");

