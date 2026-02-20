CREATE TYPE "clients".credential_status AS ENUM ('active', 'expired', 'revoked', 'rotating');

CREATE TABLE IF NOT EXISTS "clients"."client_credentials" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "application_id" UUID NOT NULL REFERENCES "clients"."applications"("id") ON DELETE CASCADE,
  "client_id" VARCHAR(255) NOT NULL UNIQUE,
  "secret_hash" VARCHAR(255) NOT NULL,
  "hash_alg" VARCHAR(50) NOT NULL,
  "status" "clients".credential_status NOT NULL DEFAULT 'active',
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "rotated_at" TIMESTAMP,
  "expires_at" TIMESTAMP,
  "last_used_at" TIMESTAMP,
  "created_by" UUID,
  "revoked_at" TIMESTAMP,
  "revoked_by" UUID,
  "revoke_reason" TEXT
);

CREATE INDEX IF NOT EXISTS "idx_client_credentials_client_id" ON "clients"."client_credentials"("client_id");
CREATE INDEX IF NOT EXISTS "idx_client_credentials_application_id" ON "clients"."client_credentials"("application_id");
CREATE INDEX IF NOT EXISTS "idx_client_credentials_status" ON "clients"."client_credentials"("status");
CREATE INDEX IF NOT EXISTS "idx_client_credentials_app_status" ON "clients"."client_credentials"("application_id", "status") WHERE "status" = 'active';
CREATE INDEX IF NOT EXISTS "idx_client_credentials_expires_at" ON "clients"."client_credentials"("expires_at") WHERE "expires_at" IS NOT NULL;
CREATE INDEX IF NOT EXISTS "idx_client_credentials_last_used_at" ON "clients"."client_credentials"("last_used_at");
