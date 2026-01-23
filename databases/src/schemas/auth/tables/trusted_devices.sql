-- Trusted Devices table for device trust management
-- "Remember this device for 30 days, skip MFA" and risk control whitelist
-- Note: No tenant_id because trusted devices are user-level, not tenant-level (user can belong to multiple tenants)
CREATE TABLE IF NOT EXISTS "auth"."trusted_devices" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "device_id" VARCHAR(255) NOT NULL,
  "user_id" UUID NOT NULL, -- References directory.users.id (application-level consistency)
  "device_fingerprint_hash" VARCHAR(255) NOT NULL, -- Hashed device fingerprint
  "device_name" VARCHAR(255), -- User-friendly device name
  "trusted_until" TIMESTAMP NOT NULL, -- Trust expiration time
  "last_used_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "ip" INET, -- IP when device was trusted
  "ua_hash" VARCHAR(64), -- User agent hash
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("user_id", "device_id")
);

CREATE INDEX IF NOT EXISTS "idx_trusted_devices_device_id" ON "auth"."trusted_devices"("device_id");
CREATE INDEX IF NOT EXISTS "idx_trusted_devices_user_id" ON "auth"."trusted_devices"("user_id");
CREATE INDEX IF NOT EXISTS "idx_trusted_devices_trusted_until" ON "auth"."trusted_devices"("trusted_until");
