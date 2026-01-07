-- Invitation Roles table: Maps invitations to multiple roles
-- One invitation can grant multiple roles
-- Otherwise putting role_id in invitations column will quickly become insufficient
CREATE TABLE IF NOT EXISTS "tenants"."invitation_roles" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "invite_id" VARCHAR(255) NOT NULL, -- References tenants.invitations.invite_id (application-level consistency)
  "role_id" UUID NOT NULL, -- References access.roles.id (application-level consistency)
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("invite_id", "role_id")
);

CREATE INDEX IF NOT EXISTS "idx_invitation_roles_invite_id" ON "tenants"."invitation_roles"("invite_id");
CREATE INDEX IF NOT EXISTS "idx_invitation_roles_role_id" ON "tenants"."invitation_roles"("role_id");

