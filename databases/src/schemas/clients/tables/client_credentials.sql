-- OAuth Client Credentials table
-- Stores client_id and secret hash, supports multi-version rotation
-- Only stores hash, never plaintext (plaintext only returned once during creation)
CREATE TYPE "clients".credential_status AS ENUM ('active', 'expired', 'revoked', 'rotating');

CREATE TABLE IF NOT EXISTS "clients"."client_credentials" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "app_id" UUID NOT NULL REFERENCES "clients"."apps"("id") ON DELETE CASCADE,
  "client_id" VARCHAR(255) NOT NULL UNIQUE, -- Public OAuth client identifier
  "secret_hash" VARCHAR(255) NOT NULL, -- Hashed secret (Argon2/bcrypt/SCrypt)
  "hash_alg" VARCHAR(50) NOT NULL, -- Algorithm used: 'argon2id', 'bcrypt', 'scrypt'
  "status" "clients".credential_status NOT NULL DEFAULT 'active',
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "rotated_at" TIMESTAMP, -- When this credential was rotated (replaced by new one)
  "expires_at" TIMESTAMP, -- Optional expiration time (enterprise common)
  "last_used_at" TIMESTAMP, -- Last time this credential was used (very useful)
  "created_by" UUID, -- Who issued/rotated this credential
  "revoked_at" TIMESTAMP,
  "revoked_by" UUID, -- Who revoked this credential
  "revoke_reason" TEXT
);

CREATE INDEX IF NOT EXISTS "idx_client_credentials_client_id" ON "clients"."client_credentials"("client_id");
CREATE INDEX IF NOT EXISTS "idx_client_credentials_app_id" ON "clients"."client_credentials"("app_id");
CREATE INDEX IF NOT EXISTS "idx_client_credentials_status" ON "clients"."client_credentials"("status");
CREATE INDEX IF NOT EXISTS "idx_client_credentials_app_status" ON "clients"."client_credentials"("app_id", "status") WHERE "status" = 'active';
CREATE INDEX IF NOT EXISTS "idx_client_credentials_expires_at" ON "clients"."client_credentials"("expires_at") WHERE "expires_at" IS NOT NULL;
CREATE INDEX IF NOT EXISTS "idx_client_credentials_last_used_at" ON "clients"."client_credentials"("last_used_at");

