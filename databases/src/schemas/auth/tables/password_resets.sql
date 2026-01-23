-- Password Resets table for secure password recovery flow
-- Supports audit, abuse prevention, and secure token management
-- Note: No tenant_id because password reset is user-level, not tenant-level (user can belong to multiple tenants)
CREATE TYPE "auth".reset_delivery AS ENUM ('email', 'sms');
CREATE TYPE "auth".reset_status AS ENUM ('issued', 'used', 'expired', 'revoked');

CREATE TABLE IF NOT EXISTS "auth"."password_resets" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "reset_id" VARCHAR(255) NOT NULL UNIQUE, -- Random reset identifier
  "user_id" UUID NOT NULL, -- References directory.users.id (application-level consistency)
  "delivery" "auth".reset_delivery NOT NULL,
  "code_hash" VARCHAR(255) NOT NULL, -- Hashed reset code/token (never store plaintext)
  "expires_at" TIMESTAMP NOT NULL,
  "used_at" TIMESTAMP,
  "requested_ip" INET,
  "ua_hash" VARCHAR(64), -- User agent hash
  "attempt_count" INTEGER NOT NULL DEFAULT 0, -- Number of failed verification attempts
  "status" "auth".reset_status NOT NULL DEFAULT 'issued',
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_password_resets_reset_id" ON "auth"."password_resets"("reset_id");
CREATE INDEX IF NOT EXISTS "idx_password_resets_user_id" ON "auth"."password_resets"("user_id");
CREATE INDEX IF NOT EXISTS "idx_password_resets_status" ON "auth"."password_resets"("status");
CREATE INDEX IF NOT EXISTS "idx_password_resets_expires_at" ON "auth"."password_resets"("expires_at");
CREATE INDEX IF NOT EXISTS "idx_password_resets_created_at" ON "auth"."password_resets"("created_at");
