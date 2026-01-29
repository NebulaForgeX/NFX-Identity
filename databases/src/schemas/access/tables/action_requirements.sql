-- Action Requirements table: which permissions are required for an action
-- One action can require multiple permissions (AND semantics: user must have all listed permissions)
-- Used when loading policy into Casbin: action -> required permission set
CREATE TABLE IF NOT EXISTS "access"."action_requirements" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "action_id" UUID NOT NULL REFERENCES "access"."actions"("id") ON DELETE CASCADE,
  "permission_id" UUID NOT NULL REFERENCES "access"."permissions"("id") ON DELETE CASCADE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("action_id", "permission_id")
);

CREATE INDEX IF NOT EXISTS "idx_action_requirements_action_id" ON "access"."action_requirements"("action_id");
CREATE INDEX IF NOT EXISTS "idx_action_requirements_permission_id" ON "access"."action_requirements"("permission_id");
CREATE INDEX IF NOT EXISTS "idx_action_requirements_action_permission" ON "access"."action_requirements"("action_id", "permission_id");
