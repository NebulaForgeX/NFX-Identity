-- Create "actions" table
CREATE TABLE "access"."actions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "key" character varying(255) NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "is_system" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "actions_key_key" UNIQUE ("key")
);
-- Create index "idx_actions_deleted_at" to table: "actions"
CREATE INDEX "idx_actions_deleted_at" ON "access"."actions" ("deleted_at");
-- Create index "idx_actions_is_system" to table: "actions"
CREATE INDEX "idx_actions_is_system" ON "access"."actions" ("is_system");
-- Create index "idx_actions_key" to table: "actions"
CREATE INDEX "idx_actions_key" ON "access"."actions" ("key");
-- Create "action_requirements" table
CREATE TABLE "access"."action_requirements" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "action_id" uuid NOT NULL,
  "permission_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "action_requirements_action_id_permission_id_key" UNIQUE ("action_id", "permission_id"),
  CONSTRAINT "action_requirements_action_id_fkey" FOREIGN KEY ("action_id") REFERENCES "access"."actions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "action_requirements_permission_id_fkey" FOREIGN KEY ("permission_id") REFERENCES "access"."permissions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_action_requirements_action_id" to table: "action_requirements"
CREATE INDEX "idx_action_requirements_action_id" ON "access"."action_requirements" ("action_id");
-- Create index "idx_action_requirements_action_permission" to table: "action_requirements"
CREATE INDEX "idx_action_requirements_action_permission" ON "access"."action_requirements" ("action_id", "permission_id");
-- Create index "idx_action_requirements_permission_id" to table: "action_requirements"
CREATE INDEX "idx_action_requirements_permission_id" ON "access"."action_requirements" ("permission_id");
