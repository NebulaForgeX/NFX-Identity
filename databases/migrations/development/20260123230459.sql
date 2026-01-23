-- Modify "account_lockouts" table
ALTER TABLE "auth"."account_lockouts" DROP COLUMN "tenant_id";
-- Modify "login_attempts" table
ALTER TABLE "auth"."login_attempts" DROP COLUMN "tenant_id";
-- Create index "idx_login_attempts_identifier" to table: "login_attempts"
CREATE INDEX "idx_login_attempts_identifier" ON "auth"."login_attempts" ("identifier", "created_at");
-- Create index "idx_login_attempts_ip" to table: "login_attempts"
CREATE INDEX "idx_login_attempts_ip" ON "auth"."login_attempts" ("ip", "created_at");
-- Modify "mfa_factors" table
ALTER TABLE "auth"."mfa_factors" DROP COLUMN "tenant_id";
-- Modify "password_history" table
ALTER TABLE "auth"."password_history" DROP COLUMN "tenant_id";
-- Modify "password_resets" table
ALTER TABLE "auth"."password_resets" DROP COLUMN "tenant_id";
-- Modify "refresh_tokens" table
ALTER TABLE "auth"."refresh_tokens" DROP COLUMN "tenant_id";
-- Modify "sessions" table
ALTER TABLE "auth"."sessions" DROP COLUMN "tenant_id";
-- Create index "idx_sessions_user_active" to table: "sessions"
CREATE INDEX "idx_sessions_user_active" ON "auth"."sessions" ("user_id", "revoked_at") WHERE (revoked_at IS NULL);
-- Modify "trusted_devices" table
ALTER TABLE "auth"."trusted_devices" DROP COLUMN "tenant_id", ADD CONSTRAINT "trusted_devices_user_id_device_id_key" UNIQUE ("user_id", "device_id");
