-- Drop index "idx_action_requirements_action_permission" from table: "action_requirements"
DROP INDEX "access"."idx_action_requirements_action_permission";
-- Modify "action_requirements" table
ALTER TABLE "access"."action_requirements" ADD COLUMN "group_id" integer NOT NULL DEFAULT 1;
-- Create index "idx_action_requirements_action_group" to table: "action_requirements"
CREATE INDEX "idx_action_requirements_action_group" ON "access"."action_requirements" ("action_id", "group_id");
-- Modify "actions" table
ALTER TABLE "access"."actions" ADD COLUMN "service" character varying(255) NOT NULL, ADD COLUMN "status" character varying(50) NOT NULL DEFAULT 'active';
-- Create index "idx_actions_service" to table: "actions"
CREATE INDEX "idx_actions_service" ON "access"."actions" ("service");
-- Create index "idx_actions_status" to table: "actions"
CREATE INDEX "idx_actions_status" ON "access"."actions" ("status");
