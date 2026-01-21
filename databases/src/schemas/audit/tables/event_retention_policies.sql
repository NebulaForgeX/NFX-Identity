-- Event Retention Policies table: Retention policy management
-- Defines retention periods for different event types/classifications
-- Supports compliance requirements: 180 days / 1 year / 7 years retention
CREATE TYPE "audit".retention_action AS ENUM ('archive', 'delete', 'export');

CREATE TABLE IF NOT EXISTS "audit"."event_retention_policies" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "policy_name" VARCHAR(255) NOT NULL UNIQUE,
  "tenant_id" UUID, -- NULL means global policy
  "action_pattern" VARCHAR(255), -- Action pattern: "auth.*", "directory.users.export", "*" for all
  "data_classification" "audit".data_classification, -- Data classification filter
  "risk_level" "audit".risk_level, -- Risk level filter
  "retention_days" INTEGER NOT NULL, -- Retention period in days
  "retention_action" "audit".retention_action NOT NULL DEFAULT 'archive', -- archive|delete|export
  "archive_location" TEXT, -- Archive storage location (S3 path, etc.)
  "status" VARCHAR(50) NOT NULL DEFAULT 'active', -- 'active', 'disabled'
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID
);

CREATE INDEX IF NOT EXISTS "idx_event_retention_policies_policy_name" ON "audit"."event_retention_policies"("policy_name");
CREATE INDEX IF NOT EXISTS "idx_event_retention_policies_tenant_id" ON "audit"."event_retention_policies"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_event_retention_policies_status" ON "audit"."event_retention_policies"("status");

