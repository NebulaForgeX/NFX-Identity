-- Service Tokens table for M2M access tokens
-- Stores issued service tokens (if opaque) or token metadata (if JWT)
-- Supports token revocation, tracking, and audit
CREATE TYPE "clients".token_type AS ENUM ('opaque', 'jwt');
CREATE TYPE "clients".token_status AS ENUM ('active', 'revoked', 'expired');

CREATE TABLE IF NOT EXISTS "clients"."service_tokens" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "token_id" VARCHAR(255) NOT NULL UNIQUE, -- Token identifier (for opaque) or JWT jti claim
  "app_id" UUID NOT NULL REFERENCES "clients"."apps"("id") ON DELETE CASCADE,
  "client_id" VARCHAR(255) REFERENCES "clients"."client_credentials"("client_id") ON DELETE SET NULL, -- OAuth client_id that issued this token
  "token_type" "clients".token_type NOT NULL DEFAULT 'jwt',
  "token_hash" VARCHAR(255), -- Hashed token (for opaque tokens, if storing)
  "scopes" TEXT[], -- Granted scopes array
  "issued_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "expires_at" TIMESTAMP NOT NULL,
  "status" "clients".token_status NOT NULL DEFAULT 'active',
  "revoked_at" TIMESTAMP,
  "revoked_by" UUID, -- Who revoked this token
  "revoke_reason" TEXT,
  "last_used_at" TIMESTAMP, -- Last time this token was used
  "ip" INET, -- IP address when token was issued
  "ua_hash" VARCHAR(64), -- User agent hash
  "metadata" JSONB DEFAULT '{}'::jsonb -- Extended fields: {"tenant_id": "...", "request_id": "...", ...}
);

CREATE INDEX IF NOT EXISTS "idx_service_tokens_token_id" ON "clients"."service_tokens"("token_id");
CREATE INDEX IF NOT EXISTS "idx_service_tokens_app_id" ON "clients"."service_tokens"("app_id");
CREATE INDEX IF NOT EXISTS "idx_service_tokens_client_id" ON "clients"."service_tokens"("client_id");
CREATE INDEX IF NOT EXISTS "idx_service_tokens_status" ON "clients"."service_tokens"("status");
CREATE INDEX IF NOT EXISTS "idx_service_tokens_expires_at" ON "clients"."service_tokens"("expires_at");
CREATE INDEX IF NOT EXISTS "idx_service_tokens_app_status" ON "clients"."service_tokens"("app_id", "status") WHERE "status" = 'active';
CREATE INDEX IF NOT EXISTS "idx_service_tokens_last_used_at" ON "clients"."service_tokens"("last_used_at");

