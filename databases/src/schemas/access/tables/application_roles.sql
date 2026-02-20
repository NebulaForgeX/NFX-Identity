CREATE TABLE IF NOT EXISTS "access"."application_roles" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "application_id" UUID NOT NULL,
  "role_key" VARCHAR(32) NOT NULL,
  "name" VARCHAR(255),
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("application_id", "role_key")
);

CREATE INDEX IF NOT EXISTS "idx_application_roles_application_id" ON "access"."application_roles"("application_id");
CREATE INDEX IF NOT EXISTS "idx_application_roles_role_key" ON "access"."application_roles"("role_key");
