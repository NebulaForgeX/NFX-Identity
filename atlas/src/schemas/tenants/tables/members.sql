-- Members table: Tenant member relationship table
-- Represents "who belongs to Company A" (employees/admins/developers)
-- Enterprise note: Members are not just "user_id", but also have member status and source
CREATE TYPE "tenants".member_status AS ENUM ('INVITED', 'ACTIVE', 'SUSPENDED', 'REMOVED');
CREATE TYPE "tenants".member_source AS ENUM ('MANUAL', 'INVITE', 'SCIM', 'SSO', 'HR_SYNC', 'IMPORT');

CREATE TABLE IF NOT EXISTS "tenants"."members" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "member_id" UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(), -- Member entity ID (avoid composite primary key for better references)
  "tenant_id" UUID NOT NULL REFERENCES "tenants"."tenants"("id") ON DELETE CASCADE,
  "user_id" UUID NOT NULL, -- References directory.users.id (application-level consistency)
  "status" "tenants".member_status NOT NULL DEFAULT 'INVITED',
  "source" "tenants".member_source NOT NULL DEFAULT 'MANUAL',
  "joined_at" TIMESTAMP, -- When member joined
  "left_at" TIMESTAMP, -- When member left/removed
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID, -- Who added/invited this member
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "external_ref" VARCHAR(255), -- External reference for SCIM/HR system integration
  "metadata" JSONB DEFAULT '{}'::jsonb, -- Extended fields
  UNIQUE("tenant_id", "user_id") -- One user can only have one member record per tenant
);

CREATE INDEX IF NOT EXISTS "idx_members_member_id" ON "tenants"."members"("member_id");
CREATE INDEX IF NOT EXISTS "idx_members_tenant_id" ON "tenants"."members"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_members_user_id" ON "tenants"."members"("user_id");
CREATE INDEX IF NOT EXISTS "idx_members_status" ON "tenants"."members"("status");
CREATE INDEX IF NOT EXISTS "idx_members_tenant_status" ON "tenants"."members"("tenant_id", "status");
CREATE INDEX IF NOT EXISTS "idx_members_source" ON "tenants"."members"("source");

