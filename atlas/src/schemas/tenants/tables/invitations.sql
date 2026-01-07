-- Invitations table: Invitation mechanism (enterprise essential)
-- "Company A admin invites colleagues to join backend" relies on this
-- Enterprise invitation table must consider security and traceability
CREATE TYPE "tenants".invitation_status AS ENUM ('PENDING', 'ACCEPTED', 'EXPIRED', 'REVOKED');

CREATE TABLE IF NOT EXISTS "tenants"."invitations" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "invite_id" VARCHAR(255) NOT NULL UNIQUE, -- Invitation identifier
  "tenant_id" UUID NOT NULL REFERENCES "tenants"."tenants"("id") ON DELETE CASCADE,
  "email" VARCHAR(255) NOT NULL, -- Normalized email (lowercase + trimmed)
  "token_hash" VARCHAR(255) NOT NULL, -- Hashed invitation token (never store plaintext)
  "expires_at" TIMESTAMP NOT NULL,
  "status" "tenants".invitation_status NOT NULL DEFAULT 'PENDING',
  "invited_by" UUID NOT NULL, -- Who sent the invitation
  "invited_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "accepted_by_user_id" UUID, -- User ID when accepted (filled on acceptance)
  "accepted_at" TIMESTAMP,
  "revoked_by" UUID, -- Who revoked this invitation
  "revoked_at" TIMESTAMP,
  "revoke_reason" TEXT,
  "role_ids" UUID[], -- Pre-assigned role IDs (references access.roles.id, application-level consistency)
  "metadata" JSONB DEFAULT '{}'::jsonb -- Extended fields: {"app_id": "...", "message": "...", ...}
);

CREATE INDEX IF NOT EXISTS "idx_invitations_invite_id" ON "tenants"."invitations"("invite_id");
CREATE INDEX IF NOT EXISTS "idx_invitations_tenant_id" ON "tenants"."invitations"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_invitations_email" ON "tenants"."invitations"("email");
CREATE INDEX IF NOT EXISTS "idx_invitations_status" ON "tenants"."invitations"("status");
CREATE INDEX IF NOT EXISTS "idx_invitations_tenant_status" ON "tenants"."invitations"("tenant_id", "status");
CREATE INDEX IF NOT EXISTS "idx_invitations_expires_at" ON "tenants"."invitations"("expires_at");
CREATE INDEX IF NOT EXISTS "idx_invitations_role_ids" ON "tenants"."invitations" USING GIN("role_ids"); -- GIN index for array queries

