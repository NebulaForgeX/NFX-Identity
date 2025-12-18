-- Create "authorization_codes" table
CREATE TABLE "permission"."authorization_codes" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "code" character varying(255) NOT NULL,
  "max_uses" integer NOT NULL DEFAULT 1,
  "used_count" integer NOT NULL DEFAULT 0,
  "created_by" uuid NULL,
  "expires_at" timestamp NULL,
  "is_active" boolean NOT NULL DEFAULT true,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "authorization_codes_code_key" UNIQUE ("code")
);
-- Create index "idx_authorization_codes_code" to table: "authorization_codes"
CREATE INDEX "idx_authorization_codes_code" ON "permission"."authorization_codes" ("code");
-- Create index "idx_authorization_codes_created_by" to table: "authorization_codes"
CREATE INDEX "idx_authorization_codes_created_by" ON "permission"."authorization_codes" ("created_by");
-- Create index "idx_authorization_codes_deleted_at" to table: "authorization_codes"
CREATE INDEX "idx_authorization_codes_deleted_at" ON "permission"."authorization_codes" ("deleted_at");
-- Create index "idx_authorization_codes_expires_at" to table: "authorization_codes"
CREATE INDEX "idx_authorization_codes_expires_at" ON "permission"."authorization_codes" ("expires_at");
-- Create index "idx_authorization_codes_is_active" to table: "authorization_codes"
CREATE INDEX "idx_authorization_codes_is_active" ON "permission"."authorization_codes" ("is_active");
