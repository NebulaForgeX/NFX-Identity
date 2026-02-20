CREATE TABLE IF NOT EXISTS "access"."application_role_assignments" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID NOT NULL,
  "application_id" UUID NOT NULL,
  "application_role_id" UUID NOT NULL,
  "assigned_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "assigned_by" UUID,
  UNIQUE("user_id", "application_id")
);

CREATE INDEX IF NOT EXISTS "idx_application_role_assignments_user_id" ON "access"."application_role_assignments"("user_id");
CREATE INDEX IF NOT EXISTS "idx_application_role_assignments_application_id" ON "access"."application_role_assignments"("application_id");
CREATE INDEX IF NOT EXISTS "idx_application_role_assignments_application_role_id" ON "access"."application_role_assignments"("application_role_id");
