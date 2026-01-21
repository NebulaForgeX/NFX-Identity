-- Scopes table: OAuth/OIDC semantic layer permissions
-- Answers: What can this token do? (for client_credentials / user token)
-- Scopes are external API contracts, permissions are internal capability points
-- Example: "users:read" scope may map to multiple internal permissions: users.read + users.count
-- Or you can make scope and permission 1:1 (but long-term expansion will be painful)
CREATE TABLE IF NOT EXISTS "access"."scopes" (
  "scope" VARCHAR(255) PRIMARY KEY, -- Scope identifier: "users:read", "users:write", "assets:read", etc.
  "description" TEXT, -- Scope description
  "is_system" BOOLEAN NOT NULL DEFAULT false, -- System built-in scope
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_scopes_is_system" ON "access"."scopes"("is_system");
CREATE INDEX IF NOT EXISTS "idx_scopes_deleted_at" ON "access"."scopes"("deleted_at");

