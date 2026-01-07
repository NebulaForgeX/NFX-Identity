-- MFA Factors table for multi-factor authentication credentials
-- Stores MFA binding information and enabled status
CREATE TYPE "auth".mfa_type AS ENUM ('totp', 'sms', 'email', 'webauthn', 'backup_code');

CREATE TABLE IF NOT EXISTS "auth"."mfa_factors" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "factor_id" VARCHAR(255) NOT NULL UNIQUE, -- Unique factor identifier
  "tenant_id" UUID NOT NULL,
  "user_id" UUID NOT NULL, -- References directory.users.id (application-level consistency)
  "type" "auth".mfa_type NOT NULL,
  "secret_encrypted" TEXT, -- Encrypted TOTP secret or public key (for webauthn)
  "phone" VARCHAR(20), -- For SMS factor
  "email" VARCHAR(255), -- For email factor
  "name" VARCHAR(255), -- User-friendly name: "My iPhone", "Work Phone", etc.
  "enabled" BOOLEAN NOT NULL DEFAULT false,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "last_used_at" TIMESTAMP,
  "recovery_codes_hash" TEXT, -- Hashed recovery codes (or use separate table)
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_mfa_factors_factor_id" ON "auth"."mfa_factors"("factor_id");
CREATE INDEX IF NOT EXISTS "idx_mfa_factors_user_id" ON "auth"."mfa_factors"("user_id");
CREATE INDEX IF NOT EXISTS "idx_mfa_factors_tenant_id" ON "auth"."mfa_factors"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_mfa_factors_type" ON "auth"."mfa_factors"("type");
CREATE INDEX IF NOT EXISTS "idx_mfa_factors_enabled" ON "auth"."mfa_factors"("user_id", "enabled") WHERE "enabled" = true;
CREATE INDEX IF NOT EXISTS "idx_mfa_factors_deleted_at" ON "auth"."mfa_factors"("deleted_at");

