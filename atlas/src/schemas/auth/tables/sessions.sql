-- Sessions table for device/session management
-- Supports: kick offline, view active devices, limit concurrent sessions, abnormal device alerts
CREATE TYPE "auth".session_revoke_reason AS ENUM (
  'user_logout',
  'admin_revoke',
  'password_changed',
  'device_changed',
  'account_locked',
  'suspicious_activity',
  'session_expired',
  'other'
);

CREATE TABLE IF NOT EXISTS "auth"."sessions" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "session_id" VARCHAR(255) NOT NULL UNIQUE, -- Unique session identifier
  "tenant_id" UUID NOT NULL,
  "user_id" UUID NOT NULL, -- References directory.users.id (application-level consistency)
  "app_id" UUID, -- Which application this session belongs to
  "client_id" VARCHAR(255), -- OAuth client identifier
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "last_seen_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "expires_at" TIMESTAMP NOT NULL,
  "ip" INET,
  "ua_hash" VARCHAR(64), -- User agent hash
  "device_id" VARCHAR(255), -- Device identifier
  "device_fingerprint" VARCHAR(255), -- Device fingerprint for risk control
  "device_name" VARCHAR(255), -- User-friendly device name
  "revoked_at" TIMESTAMP,
  "revoke_reason" "auth".session_revoke_reason,
  "revoked_by" VARCHAR(255), -- 'user' or admin user_id/username
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_sessions_session_id" ON "auth"."sessions"("session_id");
CREATE INDEX IF NOT EXISTS "idx_sessions_user_id" ON "auth"."sessions"("user_id");
CREATE INDEX IF NOT EXISTS "idx_sessions_tenant_id" ON "auth"."sessions"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_sessions_app_id" ON "auth"."sessions"("app_id");
CREATE INDEX IF NOT EXISTS "idx_sessions_expires_at" ON "auth"."sessions"("expires_at");
CREATE INDEX IF NOT EXISTS "idx_sessions_revoked_at" ON "auth"."sessions"("revoked_at") WHERE "revoked_at" IS NULL;
CREATE INDEX IF NOT EXISTS "idx_sessions_user_tenant_active" ON "auth"."sessions"("user_id", "tenant_id", "revoked_at") WHERE "revoked_at" IS NULL;
CREATE INDEX IF NOT EXISTS "idx_sessions_last_seen_at" ON "auth"."sessions"("last_seen_at");

