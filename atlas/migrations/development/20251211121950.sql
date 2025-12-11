-- Add new schema named "permission"
CREATE SCHEMA "permission";
-- Create "profile_educations" table
CREATE TABLE "auth"."profile_educations" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "profile_id" uuid NOT NULL,
  "school" character varying(255) NOT NULL,
  "degree" character varying(100) NULL,
  "major" character varying(255) NULL,
  "field_of_study" character varying(255) NULL,
  "start_date" date NULL,
  "end_date" date NULL,
  "is_current" boolean NOT NULL DEFAULT false,
  "description" text NULL,
  "grade" character varying(50) NULL,
  "activities" text NULL,
  "achievements" text NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "profile_educations_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "auth"."profiles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_profile_educations_degree" to table: "profile_educations"
CREATE INDEX "idx_profile_educations_degree" ON "auth"."profile_educations" ("degree");
-- Create index "idx_profile_educations_deleted_at" to table: "profile_educations"
CREATE INDEX "idx_profile_educations_deleted_at" ON "auth"."profile_educations" ("deleted_at");
-- Create index "idx_profile_educations_profile_id" to table: "profile_educations"
CREATE INDEX "idx_profile_educations_profile_id" ON "auth"."profile_educations" ("profile_id");
-- Create index "idx_profile_educations_school" to table: "profile_educations"
CREATE INDEX "idx_profile_educations_school" ON "auth"."profile_educations" ("school");
-- Create "profile_occupations" table
CREATE TABLE "auth"."profile_occupations" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "profile_id" uuid NOT NULL,
  "company" character varying(255) NOT NULL,
  "position" character varying(255) NOT NULL,
  "department" character varying(255) NULL,
  "industry" character varying(100) NULL,
  "location" character varying(255) NULL,
  "employment_type" character varying(50) NULL,
  "start_date" date NULL,
  "end_date" date NULL,
  "is_current" boolean NOT NULL DEFAULT false,
  "description" text NULL,
  "responsibilities" text NULL,
  "achievements" text NULL,
  "skills_used" text[] NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "profile_occupations_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "auth"."profiles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_profile_occupations_company" to table: "profile_occupations"
CREATE INDEX "idx_profile_occupations_company" ON "auth"."profile_occupations" ("company");
-- Create index "idx_profile_occupations_deleted_at" to table: "profile_occupations"
CREATE INDEX "idx_profile_occupations_deleted_at" ON "auth"."profile_occupations" ("deleted_at");
-- Create index "idx_profile_occupations_industry" to table: "profile_occupations"
CREATE INDEX "idx_profile_occupations_industry" ON "auth"."profile_occupations" ("industry");
-- Create index "idx_profile_occupations_is_current" to table: "profile_occupations"
CREATE INDEX "idx_profile_occupations_is_current" ON "auth"."profile_occupations" ("is_current");
-- Create index "idx_profile_occupations_position" to table: "profile_occupations"
CREATE INDEX "idx_profile_occupations_position" ON "auth"."profile_occupations" ("position");
-- Create index "idx_profile_occupations_profile_id" to table: "profile_occupations"
CREATE INDEX "idx_profile_occupations_profile_id" ON "auth"."profile_occupations" ("profile_id");
-- Create "permissions" table
CREATE TABLE "permission"."permissions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "tag" character varying(100) NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" text NULL,
  "category" character varying(50) NULL,
  "is_system" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "permissions_tag_key" UNIQUE ("tag")
);
-- Create index "idx_permissions_category" to table: "permissions"
CREATE INDEX "idx_permissions_category" ON "permission"."permissions" ("category");
-- Create index "idx_permissions_deleted_at" to table: "permissions"
CREATE INDEX "idx_permissions_deleted_at" ON "permission"."permissions" ("deleted_at");
-- Create index "idx_permissions_tag" to table: "permissions"
CREATE INDEX "idx_permissions_tag" ON "permission"."permissions" ("tag");
-- Create "user_permissions" table
CREATE TABLE "permission"."user_permissions" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "permission_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_permissions_user_id_permission_id_key" UNIQUE ("user_id", "permission_id"),
  CONSTRAINT "user_permissions_permission_id_fkey" FOREIGN KEY ("permission_id") REFERENCES "permission"."permissions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_permissions_deleted_at" to table: "user_permissions"
CREATE INDEX "idx_user_permissions_deleted_at" ON "permission"."user_permissions" ("deleted_at");
-- Create index "idx_user_permissions_permission_id" to table: "user_permissions"
CREATE INDEX "idx_user_permissions_permission_id" ON "permission"."user_permissions" ("permission_id");
-- Create index "idx_user_permissions_user_id" to table: "user_permissions"
CREATE INDEX "idx_user_permissions_user_id" ON "permission"."user_permissions" ("user_id");
-- Modify "users" table
ALTER TABLE "auth"."users" DROP COLUMN "role_id";
-- Create "user_roles" table
CREATE TABLE "auth"."user_roles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "role_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_roles_user_id_role_id_key" UNIQUE ("user_id", "role_id"),
  CONSTRAINT "user_roles_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "auth"."roles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "user_roles_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "auth"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_roles_role_id" to table: "user_roles"
CREATE INDEX "idx_user_roles_role_id" ON "auth"."user_roles" ("role_id");
-- Create index "idx_user_roles_user_id" to table: "user_roles"
CREATE INDEX "idx_user_roles_user_id" ON "auth"."user_roles" ("user_id");
-- Drop "educations" table
DROP TABLE "auth"."educations";
-- Drop "occupations" table
DROP TABLE "auth"."occupations";
