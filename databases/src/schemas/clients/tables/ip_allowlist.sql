CREATE TYPE "clients".allowlist_status AS ENUM ('active', 'disabled', 'revoked');

CREATE TABLE IF NOT EXISTS "clients"."ip_allowlist" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "rule_id" VARCHAR(255) NOT NULL UNIQUE,
  "application_id" UUID NOT NULL REFERENCES "clients"."applications"("id") ON DELETE CASCADE,
  "cidr" CIDR NOT NULL,
  "description" TEXT,
  "status" "clients".allowlist_status NOT NULL DEFAULT 'active',
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_by" UUID,
  "revoked_at" TIMESTAMP,
  "revoked_by" UUID,
  "revoke_reason" TEXT
);

CREATE INDEX IF NOT EXISTS "idx_ip_allowlist_rule_id" ON "clients"."ip_allowlist"("rule_id");
CREATE INDEX IF NOT EXISTS "idx_ip_allowlist_application_id" ON "clients"."ip_allowlist"("application_id");
CREATE INDEX IF NOT EXISTS "idx_ip_allowlist_status" ON "clients"."ip_allowlist"("status");
CREATE INDEX IF NOT EXISTS "idx_ip_allowlist_app_status" ON "clients"."ip_allowlist"("application_id", "status") WHERE "status" = 'active';
