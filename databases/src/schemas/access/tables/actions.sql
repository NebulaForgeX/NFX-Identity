-- Actions table: registry of callable actions (service:resource.verb)
-- CheckAccess Step A: request action_key must exist here and status = 'active'
-- key format: service:resource.verb, e.g. "directory:user.get", "directory:users.count", "tenants:tenant.create"
CREATE TABLE IF NOT EXISTS "access"."actions" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "key" VARCHAR(255) NOT NULL UNIQUE,
  "service" VARCHAR(255) NOT NULL, -- Service name, e.g. "directory", "tenants" (prefix of key)
  "status" VARCHAR(50) NOT NULL DEFAULT 'active', -- active | deprecated | disabled; only active can pass Step A
  "name" VARCHAR(255) NOT NULL,
  "description" TEXT,
  "is_system" BOOLEAN NOT NULL DEFAULT false,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_actions_key" ON "access"."actions"("key");
CREATE INDEX IF NOT EXISTS "idx_actions_service" ON "access"."actions"("service");
CREATE INDEX IF NOT EXISTS "idx_actions_status" ON "access"."actions"("status");
CREATE INDEX IF NOT EXISTS "idx_actions_is_system" ON "access"."actions"("is_system");
CREATE INDEX IF NOT EXISTS "idx_actions_deleted_at" ON "access"."actions"("deleted_at");
