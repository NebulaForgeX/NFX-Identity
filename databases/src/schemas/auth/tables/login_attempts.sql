-- Login Attempts table for authentication events and risk control
-- Records every authentication attempt for locking, rate limiting, audit, and risk scoring
-- Note: No tenant_id because login happens before tenant selection (user can belong to multiple tenants)
CREATE TYPE "auth".failure_code AS ENUM (
  'bad_password',
  'user_not_found',
  'locked',
  'mfa_required',
  'mfa_failed',
  'account_disabled',
  'credential_expired',
  'rate_limited',
  'ip_blocked',
  'device_not_trusted',
  'other'
);

CREATE TABLE IF NOT EXISTS "auth"."login_attempts" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "identifier" VARCHAR(255) NOT NULL, -- Normalized email/phone/username
  "user_id" UUID, -- References directory.users.id (filled after successful login, application-level consistency)
  "ip" INET,
  "ua_hash" VARCHAR(64), -- User agent hash
  "device_fingerprint" VARCHAR(255), -- Device fingerprint for risk control
  "success" BOOLEAN NOT NULL DEFAULT false,
  "failure_code" "auth".failure_code,
  "mfa_required" BOOLEAN NOT NULL DEFAULT false,
  "mfa_verified" BOOLEAN NOT NULL DEFAULT false,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_login_attempts_identifier" ON "auth"."login_attempts"("identifier", "created_at");
CREATE INDEX IF NOT EXISTS "idx_login_attempts_ip" ON "auth"."login_attempts"("ip", "created_at");
CREATE INDEX IF NOT EXISTS "idx_login_attempts_user_id" ON "auth"."login_attempts"("user_id", "created_at");
CREATE INDEX IF NOT EXISTS "idx_login_attempts_success" ON "auth"."login_attempts"("success", "created_at");
CREATE INDEX IF NOT EXISTS "idx_login_attempts_created_at" ON "auth"."login_attempts"("created_at");
