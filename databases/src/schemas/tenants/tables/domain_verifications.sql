-- Domain Verifications table: Verify company domain ownership
-- Prevents others from impersonating companyA.com
CREATE TYPE "tenants".verification_method AS ENUM ('DNS', 'TXT', 'HTML', 'FILE');
CREATE TYPE "tenants".verification_status AS ENUM ('PENDING', 'VERIFIED', 'FAILED', 'EXPIRED');

CREATE TABLE IF NOT EXISTS "tenants"."domain_verifications" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "tenant_id" UUID NOT NULL REFERENCES "tenants"."tenants"("id") ON DELETE CASCADE,
  "domain" VARCHAR(255) NOT NULL, -- Domain to verify: companyA.com
  "verification_method" "tenants".verification_method NOT NULL DEFAULT 'DNS',
  "verification_token" VARCHAR(255), -- Verification token/record value
  "status" "tenants".verification_status NOT NULL DEFAULT 'PENDING',
  "verified_at" TIMESTAMP,
  "expires_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "metadata" JSONB DEFAULT '{}'::jsonb, -- Extended fields: {"dns_record": "...", "txt_record": "...", ...}
  UNIQUE("tenant_id", "domain")
);

CREATE INDEX IF NOT EXISTS "idx_domain_verifications_tenant_id" ON "tenants"."domain_verifications"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_domain_verifications_domain" ON "tenants"."domain_verifications"("domain");
CREATE INDEX IF NOT EXISTS "idx_domain_verifications_status" ON "tenants"."domain_verifications"("status");

