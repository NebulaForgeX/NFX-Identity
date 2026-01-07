-- Credential Events table for credential change audit trail
-- Records password changes, MFA bindings, credential disables, etc. for security audit
CREATE TYPE "auth".actor_type AS ENUM ('user', 'admin', 'system');

CREATE TYPE "auth".credential_event_type AS ENUM (
  'password_changed',
  'password_reset',
  'mfa_enabled',
  'mfa_disabled',
  'mfa_factor_added',
  'mfa_factor_removed',
  'credential_disabled',
  'credential_enabled',
  'credential_expired',
  'passkey_added',
  'passkey_removed',
  'oauth_linked',
  'oauth_unlinked',
  'other'
);

CREATE TABLE IF NOT EXISTS "auth"."credential_events" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "event_id" VARCHAR(255) NOT NULL UNIQUE, -- Unique event identifier
  "tenant_id" UUID NOT NULL,
  "user_id" UUID NOT NULL, -- References directory.users.id (application-level consistency)
  "event_type" "auth".credential_event_type NOT NULL,
  "metadata" JSONB DEFAULT '{}'::jsonb, -- Additional event data: {"ip": "...", "device": "...", "factor_type": "totp", ...}
  "actor_id" UUID, -- Who performed the action (user_id or admin_id)
  "actor_type" "auth".actor_type, -- 'user', 'admin', 'system'
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_credential_events_event_id" ON "auth"."credential_events"("event_id");
CREATE INDEX IF NOT EXISTS "idx_credential_events_user_id" ON "auth"."credential_events"("user_id");
CREATE INDEX IF NOT EXISTS "idx_credential_events_tenant_id" ON "auth"."credential_events"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_credential_events_type" ON "auth"."credential_events"("event_type");
CREATE INDEX IF NOT EXISTS "idx_credential_events_created_at" ON "auth"."credential_events"("created_at");
CREATE INDEX IF NOT EXISTS "idx_credential_events_user_created" ON "auth"."credential_events"("user_id", "created_at");

