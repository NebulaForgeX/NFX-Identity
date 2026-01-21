-- Client Scopes table for allowed permission scopes (Allow-list)
-- Determines which scopes this app can request
-- In OAuth: final effective scope = requested scope ∩ allowed scope
CREATE TABLE IF NOT EXISTS "clients"."client_scopes" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "app_id" UUID NOT NULL REFERENCES "clients"."apps"("id") ON DELETE CASCADE,
  "scope" VARCHAR(255) NOT NULL, -- Scope identifier: "users.read", "users.count", "users.write", etc.
  "granted_by" UUID, -- Who approved this scope (admin user_id)
  "granted_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "expires_at" TIMESTAMP, -- Optional: temporary authorization
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "revoked_at" TIMESTAMP,
  "revoked_by" UUID,
  "revoke_reason" TEXT,
  UNIQUE("app_id", "scope")
);

CREATE INDEX IF NOT EXISTS "idx_client_scopes_app_id" ON "clients"."client_scopes"("app_id");
CREATE INDEX IF NOT EXISTS "idx_client_scopes_scope" ON "clients"."client_scopes"("scope");
CREATE INDEX IF NOT EXISTS "idx_client_scopes_app_scope" ON "clients"."client_scopes"("app_id", "scope");
CREATE INDEX IF NOT EXISTS "idx_client_scopes_expires_at" ON "clients"."client_scopes"("expires_at") WHERE "expires_at" IS NOT NULL;
CREATE INDEX IF NOT EXISTS "idx_client_scopes_revoked_at" ON "clients"."client_scopes"("revoked_at") WHERE "revoked_at" IS NULL;

