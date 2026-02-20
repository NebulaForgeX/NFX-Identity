CREATE TYPE "clients".app_type AS ENUM ('server', 'service', 'internal', 'partner', 'third_party');
CREATE TYPE "clients".app_status AS ENUM ('active', 'disabled', 'suspended', 'pending');
CREATE TYPE "clients".environment AS ENUM ('production', 'staging', 'development', 'test');

CREATE TABLE IF NOT EXISTS "clients"."applications" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "application_id" VARCHAR(255) NOT NULL UNIQUE,
  "tenant_id" UUID NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "description" TEXT,
  "type" "clients".app_type NOT NULL DEFAULT 'server',
  "status" "clients".app_status NOT NULL DEFAULT 'pending',
  "environment" "clients".environment NOT NULL DEFAULT 'development',
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "updated_by" UUID,
  "metadata" JSONB DEFAULT '{}'::jsonb,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_applications_application_id" ON "clients"."applications"("application_id");
CREATE INDEX IF NOT EXISTS "idx_applications_tenant_id" ON "clients"."applications"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_applications_status" ON "clients"."applications"("status");
CREATE INDEX IF NOT EXISTS "idx_applications_environment" ON "clients"."applications"("environment");
CREATE INDEX IF NOT EXISTS "idx_applications_tenant_environment" ON "clients"."applications"("tenant_id", "environment");
CREATE INDEX IF NOT EXISTS "idx_applications_deleted_at" ON "clients"."applications"("deleted_at");
