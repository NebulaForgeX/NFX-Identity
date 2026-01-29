-- Actions table: registry of callable actions (service + function)
-- Used by Casbin / Enforce: "can subject do this action in this tenant?"
-- action key = service name + function name, e.g. "directory.GetUser", "tenants.CreateTenant"
CREATE TABLE IF NOT EXISTS "access"."actions" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "key" VARCHAR(255) NOT NULL UNIQUE, -- Business key: "directory.GetUser", "tenants.CreateTenant"
  "name" VARCHAR(255) NOT NULL, -- Display name
  "description" TEXT, -- Action description
  "is_system" BOOLEAN NOT NULL DEFAULT false, -- System-built, cannot be deleted
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_actions_key" ON "access"."actions"("key");
CREATE INDEX IF NOT EXISTS "idx_actions_is_system" ON "access"."actions"("is_system");
CREATE INDEX IF NOT EXISTS "idx_actions_deleted_at" ON "access"."actions"("deleted_at");
