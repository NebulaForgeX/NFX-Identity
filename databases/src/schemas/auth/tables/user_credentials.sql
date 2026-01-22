-- User Credentials table for authentication materials
-- Stores how to verify a user (password, passkey, oauth links, etc.)
-- Note: id directly references directory.users.id (one-to-one relationship)
CREATE TYPE "auth".credential_type AS ENUM ('password', 'passkey', 'oauth_link', 'saml', 'ldap');
CREATE TYPE "auth".credential_status AS ENUM ('active', 'disabled', 'expired');

CREATE TABLE IF NOT EXISTS "auth"."user_credentials" (
  "id" UUID PRIMARY KEY, -- References directory.users.id (application-level consistency, no FK to avoid cross-schema dependency)
  "credential_type" "auth".credential_type NOT NULL DEFAULT 'password',
  "password_hash" VARCHAR(255), -- For password type
  "hash_alg" VARCHAR(50), -- e.g., 'bcrypt', 'argon2id', 'scrypt'
  "hash_params" JSONB DEFAULT '{}'::jsonb, -- Algorithm parameters: {"cost": 10, "salt": "..."}
  "password_updated_at" TIMESTAMP,
  "last_success_login_at" TIMESTAMP, -- Last successful login timestamp
  "status" "auth".credential_status NOT NULL DEFAULT 'active',
  "must_change_password" BOOLEAN NOT NULL DEFAULT false, -- Force password change on next login
  "version" INTEGER NOT NULL DEFAULT 1, -- Optimistic locking for concurrent updates
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_user_credentials_type" ON "auth"."user_credentials"("credential_type");
CREATE INDEX IF NOT EXISTS "idx_user_credentials_status" ON "auth"."user_credentials"("status");
CREATE INDEX IF NOT EXISTS "idx_user_credentials_deleted_at" ON "auth"."user_credentials"("deleted_at");

