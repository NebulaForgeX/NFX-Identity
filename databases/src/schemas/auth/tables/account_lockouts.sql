-- Account Lockouts table for user lock/unlock state management
-- Provides explicit lock state instead of aggregating from login_attempts
-- Note: No tenant_id because account lockout is user-level, not tenant-level (user can belong to multiple tenants)
CREATE TYPE "auth".lock_reason AS ENUM (
  'too_many_attempts',
  'admin_lock',
  'risk_detected',
  'suspicious_activity',
  'compliance',
  'other'
);

CREATE TABLE IF NOT EXISTS "auth"."account_lockouts" (
  "user_id" UUID PRIMARY KEY, -- References directory.users.id (application-level consistency)
  "locked_until" TIMESTAMP, -- NULL means permanently locked (until manually unlocked)
  "lock_reason" "auth".lock_reason NOT NULL,
  "locked_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "locked_by" VARCHAR(255), -- 'system' or admin user_id/username
  "actor_id" UUID, -- Admin user who locked (if manual lock)
  "unlocked_at" TIMESTAMP,
  "unlocked_by" VARCHAR(255), -- Admin user who unlocked
  "unlock_actor_id" UUID, -- Admin user who unlocked (if manual unlock)
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_account_lockouts_locked_until" ON "auth"."account_lockouts"("locked_until") WHERE "locked_until" IS NOT NULL;
CREATE INDEX IF NOT EXISTS "idx_account_lockouts_locked_at" ON "auth"."account_lockouts"("locked_at");
