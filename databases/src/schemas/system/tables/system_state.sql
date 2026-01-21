-- System State table: Records system initialization and bootstrap state
-- Used to check if system has been initialized on service startup
-- Logic: Check if no record exists OR latest record has initialized = false -> system is not initialized
-- On initialization: Create new record with initialized = true
-- On reset: Delete all records or set initialized = false (allows re-initialization)
CREATE TABLE IF NOT EXISTS "system"."system_state" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "initialized" BOOLEAN NOT NULL DEFAULT false, -- Whether system has been initialized
  "initialized_at" TIMESTAMP, -- When system was initialized
  "initialization_version" VARCHAR(50), -- Version of initialization schema/data
  "last_reset_at" TIMESTAMP, -- Last time system was reset
  "last_reset_by" UUID, -- User ID who reset the system (references directory.users.id, application-level consistency). Even if database is cleared after reset, can be traced via log files for accountability
  "reset_count" INTEGER NOT NULL DEFAULT 0, -- Number of times system has been reset
  "metadata" JSONB DEFAULT '{}'::jsonb, -- Extended fields: {"bootstrap_token": "...", "services_initialized": [...], ...}
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_system_state_initialized" ON "system"."system_state"("initialized");
CREATE INDEX IF NOT EXISTS "idx_system_state_created_at" ON "system"."system_state"("created_at" DESC); -- For querying latest record

COMMENT ON TABLE "system"."system_state" IS 'System initialization state: tracks if system has been bootstrapped. Query logic: SELECT initialized FROM system_state ORDER BY created_at DESC LIMIT 1. If no record exists OR initialized = false, system is not initialized.';
COMMENT ON COLUMN "system"."system_state"."id" IS 'UUID primary key (not fixed, allows multiple records for reset/re-initialization)';
COMMENT ON COLUMN "system"."system_state"."initialized" IS 'Whether system has been initialized (checked on service startup). Always check latest record by created_at DESC.';
COMMENT ON COLUMN "system"."system_state"."initialized_at" IS 'Timestamp when system was initialized via /bootstrap/initialize';
COMMENT ON COLUMN "system"."system_state"."initialization_version" IS 'Version of initialization schema/data for migration tracking';
COMMENT ON COLUMN "system"."system_state"."last_reset_at" IS 'Timestamp when system was last reset';
COMMENT ON COLUMN "system"."system_state"."last_reset_by" IS 'User ID who reset the system. Even if database is cleared after reset, can be traced via log files for accountability';
COMMENT ON COLUMN "system"."system_state"."created_at" IS 'Record creation time. Used to determine latest state (ORDER BY created_at DESC LIMIT 1)';
