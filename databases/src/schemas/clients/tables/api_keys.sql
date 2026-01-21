-- API Keys table for simplified credentials
-- Used by scripts/internal systems/low-complexity integrators
-- Higher risk than OAuth: direct access, requires stronger constraints (IP/permissions/expiration)
CREATE TYPE "clients".api_key_status AS ENUM ('active', 'revoked', 'expired');

CREATE TABLE IF NOT EXISTS "clients"."api_keys" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "key_id" VARCHAR(255) NOT NULL UNIQUE, -- Public key identifier (prefix)
  "app_id" UUID NOT NULL REFERENCES "clients"."apps"("id") ON DELETE CASCADE,
  "key_hash" VARCHAR(255) NOT NULL, -- Hashed API key (never store plaintext)
  "hash_alg" VARCHAR(50) NOT NULL, -- Algorithm used
  "name" VARCHAR(255) NOT NULL, -- User-friendly name: "prod-batch-job-key", "staging-webhook-key"
  "status" "clients".api_key_status NOT NULL DEFAULT 'active',
  "expires_at" TIMESTAMP, -- Expiration time (strongly recommended)
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "revoked_at" TIMESTAMP,
  "revoked_by" UUID, -- Who revoked this key
  "revoke_reason" TEXT,
  "last_used_at" TIMESTAMP, -- Last time this key was used
  "created_by" UUID, -- Who created this key
  "metadata" JSONB DEFAULT '{}'::jsonb -- Extended fields: {"purpose": "...", "contact": "...", ...}
);

CREATE INDEX IF NOT EXISTS "idx_api_keys_key_id" ON "clients"."api_keys"("key_id");
CREATE INDEX IF NOT EXISTS "idx_api_keys_app_id" ON "clients"."api_keys"("app_id");
CREATE INDEX IF NOT EXISTS "idx_api_keys_status" ON "clients"."api_keys"("status");
CREATE INDEX IF NOT EXISTS "idx_api_keys_app_status" ON "clients"."api_keys"("app_id", "status") WHERE "status" = 'active';
CREATE INDEX IF NOT EXISTS "idx_api_keys_expires_at" ON "clients"."api_keys"("expires_at") WHERE "expires_at" IS NOT NULL;
CREATE INDEX IF NOT EXISTS "idx_api_keys_last_used_at" ON "clients"."api_keys"("last_used_at");

