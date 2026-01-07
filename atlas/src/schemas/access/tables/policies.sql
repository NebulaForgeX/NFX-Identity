-- Policies table: ABAC (Attribute-Based Access Control) conditional authorization
-- Answers: Under what conditions is access allowed/denied?
-- Not all scenarios can be covered by simple RBAC
-- Typical ABAC conditions:
--   - "Can only read users in the same tenant" (often implemented via tenant_id claim + query conditions, may not need policy)
--   - "Support can only read tickets for users they are responsible for"
--   - "Can only access export interface during working hours"
--   - "High-risk permissions require secondary verification/MFA" (usually linked with verification/risk)
CREATE TYPE "access".policy_effect AS ENUM ('ALLOW', 'DENY');
CREATE TYPE "access".resource_type AS ENUM ('user', 'tenant', 'app', 'asset', 'client', 'other');

CREATE TABLE IF NOT EXISTS "access"."policies" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" VARCHAR(255) NOT NULL, -- Policy name
  "description" TEXT, -- Policy description
  "effect" "access".policy_effect NOT NULL DEFAULT 'ALLOW', -- ALLOW or DENY
  "priority" INTEGER NOT NULL DEFAULT 100, -- Priority when conflicts occur (lower number = higher priority)
  "condition" JSONB DEFAULT '{}'::jsonb, -- Condition expression/rules: {"tenant_id": "...", "time_range": {"start": "09:00", "end": "18:00"}, "mfa_required": true, ...}
  "resource_type" "access".resource_type, -- Resource type this policy applies to
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID, -- Who created this policy
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_policies_name" ON "access"."policies"("name");
CREATE INDEX IF NOT EXISTS "idx_policies_effect" ON "access"."policies"("effect");
CREATE INDEX IF NOT EXISTS "idx_policies_priority" ON "access"."policies"("priority");
CREATE INDEX IF NOT EXISTS "idx_policies_resource_type" ON "access"."policies"("resource_type");
CREATE INDEX IF NOT EXISTS "idx_policies_deleted_at" ON "access"."policies"("deleted_at");

