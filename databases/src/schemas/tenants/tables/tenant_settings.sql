-- Tenant Settings table: Tenant-level policies and configurations
-- Enterprise policies: enforce MFA, password minimum length, session duration, allowed login methods, email domain restrictions, etc.
-- Not recommended to scatter in metadata
-- Note: id directly references tenants.id (one-to-one relationship)
CREATE TABLE IF NOT EXISTS "tenants"."tenant_settings" (
  "id" UUID PRIMARY KEY REFERENCES "tenants"."tenants"("id") ON DELETE CASCADE,
  "enforce_mfa" BOOLEAN NOT NULL DEFAULT false, -- Whether to enforce MFA
  "allowed_email_domains" TEXT[], -- Allowed email domains for this tenant
  "session_ttl_minutes" INTEGER, -- Session TTL in minutes
  "password_policy" JSONB DEFAULT '{}'::jsonb, -- Password policy: {"min_length": 8, "require_uppercase": true, "require_lowercase": true, "require_numbers": true, "require_symbols": false, "max_age_days": 90, ...}
  "login_policy" JSONB DEFAULT '{}'::jsonb, -- Login policy: {"allowed_methods": ["password", "sso"], "sso_provider": "...", "ip_restrictions": [...], ...}
  "mfa_policy" JSONB DEFAULT '{}'::jsonb, -- MFA policy: {"required_for_admin": true, "allowed_methods": ["totp", "sms"], ...}
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_by" UUID
);


