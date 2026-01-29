-- Action Requirements table: which permissions are required for an action (CheckAccess Step B)
-- Same group_id = AND (user must have all listed permissions in that group); different group_id = OR (satisfy at least one group)
CREATE TABLE IF NOT EXISTS "access"."action_requirements" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "action_id" UUID NOT NULL REFERENCES "access"."actions"("id") ON DELETE CASCADE,
  "permission_id" UUID NOT NULL REFERENCES "access"."permissions"("id") ON DELETE CASCADE,
  "group_id" INTEGER NOT NULL DEFAULT 1, -- Same group_id = AND; different group_id = OR
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("action_id", "permission_id")
);

CREATE INDEX IF NOT EXISTS "idx_action_requirements_action_id" ON "access"."action_requirements"("action_id");
CREATE INDEX IF NOT EXISTS "idx_action_requirements_permission_id" ON "access"."action_requirements"("permission_id");
CREATE INDEX IF NOT EXISTS "idx_action_requirements_action_group" ON "access"."action_requirements"("action_id", "group_id");
