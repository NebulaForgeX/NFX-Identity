-- Applications/Client main table
-- Represents a system/application/backend service that can integrate with the platform
-- Company A may have multiple systems, each should have one app
CREATE TYPE "clients".app_type AS ENUM ('server', 'service', 'internal', 'partner', 'third_party');
CREATE TYPE "clients".app_status AS ENUM ('active', 'disabled', 'suspended', 'pending');
CREATE TYPE "clients".environment AS ENUM ('production', 'staging', 'development', 'test');

CREATE TABLE IF NOT EXISTS "clients"."apps" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "app_id" VARCHAR(255) NOT NULL UNIQUE, -- Public application identifier
  "tenant_id" UUID NOT NULL, -- Which company this app belongs to (references tenants.tenants.id)
  "name" VARCHAR(255) NOT NULL,
  "description" TEXT,
  "type" "clients".app_type NOT NULL DEFAULT 'server',
  "status" "clients".app_status NOT NULL DEFAULT 'pending',
  "environment" "clients".environment NOT NULL DEFAULT 'development', -- Critical: separate prod/staging/dev credentials
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID, -- Admin user_id or member_id who created this app (for accountability)
  "updated_by" UUID, -- Admin user_id or member_id who last updated
  "metadata" JSONB DEFAULT '{}'::jsonb, -- Extended fields: {"contact_email": "...", "webhook_url": "...", ...}
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_apps_app_id" ON "clients"."apps"("app_id");
CREATE INDEX IF NOT EXISTS "idx_apps_tenant_id" ON "clients"."apps"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_apps_status" ON "clients"."apps"("status");
CREATE INDEX IF NOT EXISTS "idx_apps_environment" ON "clients"."apps"("environment");
CREATE INDEX IF NOT EXISTS "idx_apps_tenant_environment" ON "clients"."apps"("tenant_id", "environment");
CREATE INDEX IF NOT EXISTS "idx_apps_deleted_at" ON "clients"."apps"("deleted_at");

