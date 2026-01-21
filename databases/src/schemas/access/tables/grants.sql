-- Grants table: Core authorization table - who is granted what in what scope
-- Answers: What roles or permissions do users/services actually have?
-- Enterprise must be able to express:
--   - subject: user or service_client
--   - scope: in tenant, in app, or even on specific resources (resource-level)
--   - grant: grant role or directly grant permission (both must be supported, enterprises often need exception grants)
CREATE TYPE "access".subject_type AS ENUM ('USER', 'CLIENT');
CREATE TYPE "access".grant_type AS ENUM ('ROLE', 'PERMISSION');
CREATE TYPE "access".grant_effect AS ENUM ('ALLOW', 'DENY');

CREATE TABLE IF NOT EXISTS "access"."grants" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "subject_type" "access".subject_type NOT NULL, -- USER or CLIENT
  "subject_id" UUID NOT NULL, -- user_id or client_id (references directory.users.id or clients.apps.id, application-level consistency)
  "grant_type" "access".grant_type NOT NULL, -- ROLE or PERMISSION
  "grant_ref_id" UUID NOT NULL, -- role_id or permission_id (references access.roles.id or access.permissions.id, application-level consistency)
  "tenant_id" UUID, -- NULL means global grant
  "app_id" UUID, -- NULL means not app-scoped
  "resource_type" VARCHAR(100), -- Resource type: "user", "tenant", "app", "asset", etc.
  "resource_id" UUID, -- Specific resource ID (for resource-level authorization)
  "effect" "access".grant_effect NOT NULL DEFAULT 'ALLOW', -- ALLOW or DENY (otherwise DENY can only be achieved via policy)
  "expires_at" TIMESTAMP, -- Temporary grants are common
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID, -- Who granted this (admin user_id)
  "revoked_at" TIMESTAMP,
  "revoked_by" UUID, -- Who revoked this
  "revoke_reason" TEXT
);

CREATE INDEX IF NOT EXISTS "idx_grants_subject" ON "access"."grants"("subject_type", "subject_id");
CREATE INDEX IF NOT EXISTS "idx_grants_grant_type" ON "access"."grants"("grant_type", "grant_ref_id");
CREATE INDEX IF NOT EXISTS "idx_grants_tenant_id" ON "access"."grants"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_grants_app_id" ON "access"."grants"("app_id");
CREATE INDEX IF NOT EXISTS "idx_grants_resource" ON "access"."grants"("resource_type", "resource_id");
CREATE INDEX IF NOT EXISTS "idx_grants_effect" ON "access"."grants"("effect");
CREATE INDEX IF NOT EXISTS "idx_grants_expires_at" ON "access"."grants"("expires_at") WHERE "expires_at" IS NOT NULL;
CREATE INDEX IF NOT EXISTS "idx_grants_revoked_at" ON "access"."grants"("revoked_at") WHERE "revoked_at" IS NULL;
CREATE INDEX IF NOT EXISTS "idx_grants_subject_tenant" ON "access"."grants"("subject_type", "subject_id", "tenant_id");

