-- Refresh Tokens table for long-lived sessions with rotation support
-- Supports token rotation and revocation for enterprise security
CREATE TYPE "auth".revoke_reason AS ENUM (
  'user_logout',
  'admin_revoke',
  'password_changed',
  'rotation',
  'account_locked',
  'device_changed',
  'suspicious_activity',
  'other'
);

CREATE TABLE IF NOT EXISTS "auth"."refresh_tokens" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "token_id" VARCHAR(255) NOT NULL UNIQUE, -- Random token identifier (not guessable)
  "user_id" UUID NOT NULL, -- References directory.users.id (application-level consistency)
  "tenant_id" UUID NOT NULL, -- Multi-tenant isolation
  "app_id" UUID, -- Which application/client this token belongs to (references clients.apps.id)
  "client_id" VARCHAR(255), -- OAuth client identifier
  "session_id" UUID, -- Link to sessions table (optional)
  "issued_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "expires_at" TIMESTAMP NOT NULL,
  "revoked_at" TIMESTAMP,
  "revoke_reason" "auth".revoke_reason,
  "rotated_from" UUID REFERENCES "auth"."refresh_tokens"("id"), -- Previous token ID (for rotation chain)
  "device_id" VARCHAR(255), -- Device identifier for device management
  "ip" INET, -- IP address when issued
  "ua_hash" VARCHAR(64), -- User agent hash for risk control
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_token_id" ON "auth"."refresh_tokens"("token_id");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_user_id" ON "auth"."refresh_tokens"("user_id");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_tenant_id" ON "auth"."refresh_tokens"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_app_id" ON "auth"."refresh_tokens"("app_id");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_session_id" ON "auth"."refresh_tokens"("session_id");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_expires_at" ON "auth"."refresh_tokens"("expires_at");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_revoked_at" ON "auth"."refresh_tokens"("revoked_at") WHERE "revoked_at" IS NULL;
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_user_tenant" ON "auth"."refresh_tokens"("user_id", "tenant_id");

