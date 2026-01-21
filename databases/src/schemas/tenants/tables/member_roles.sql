-- Member Roles table: Role assignment within tenant
-- Represents "what permission level does a Company A employee have in Company A"
-- Enterprise recommendation: Use member_id instead of user_id because:
--   - Users have different identities/statuses in different tenants; member_id is cleaner
--   - Audit: "member" is closer to business semantics than "user"
--   - Future support for "temporary roles/expiring roles/delegation" is easier
CREATE TABLE IF NOT EXISTS "tenants"."member_roles" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "tenant_id" UUID NOT NULL REFERENCES "tenants"."tenants"("id") ON DELETE CASCADE,
  "member_id" UUID NOT NULL REFERENCES "tenants"."members"("member_id") ON DELETE CASCADE,
  "role_id" UUID NOT NULL, -- References access.roles.id (application-level consistency)
  "assigned_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "assigned_by" UUID, -- Who assigned this role
  "expires_at" TIMESTAMP, -- Enterprise common: temporary permissions
  "scope" VARCHAR(100), -- Optional: role only effective in certain app, scope of effect
  "revoked_at" TIMESTAMP,
  "revoked_by" UUID,
  "revoke_reason" TEXT,
  UNIQUE("tenant_id", "member_id", "role_id") -- Avoid duplicate role assignment
);

CREATE INDEX IF NOT EXISTS "idx_member_roles_tenant_id" ON "tenants"."member_roles"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_member_roles_member_id" ON "tenants"."member_roles"("member_id");
CREATE INDEX IF NOT EXISTS "idx_member_roles_role_id" ON "tenants"."member_roles"("role_id");
CREATE INDEX IF NOT EXISTS "idx_member_roles_expires_at" ON "tenants"."member_roles"("expires_at") WHERE "expires_at" IS NOT NULL;
CREATE INDEX IF NOT EXISTS "idx_member_roles_revoked_at" ON "tenants"."member_roles"("revoked_at") WHERE "revoked_at" IS NULL;

