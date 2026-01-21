-- Rate Limits table for client-level rate limiting
-- Controls request rate per app/client to prevent abuse
CREATE TYPE "clients".rate_limit_type AS ENUM ('requests_per_second', 'requests_per_minute', 'requests_per_hour', 'requests_per_day');

CREATE TABLE IF NOT EXISTS "clients"."rate_limits" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "app_id" UUID NOT NULL REFERENCES "clients"."apps"("id") ON DELETE CASCADE,
  "limit_type" "clients".rate_limit_type NOT NULL DEFAULT 'requests_per_minute',
  "limit_value" INTEGER NOT NULL, -- Number of requests allowed
  "window_seconds" INTEGER NOT NULL, -- Time window in seconds
  "description" TEXT, -- Rule description
  "status" VARCHAR(50) NOT NULL DEFAULT 'active', -- 'active', 'disabled'
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_by" UUID,
  UNIQUE("app_id", "limit_type")
);

CREATE INDEX IF NOT EXISTS "idx_rate_limits_app_id" ON "clients"."rate_limits"("app_id");
CREATE INDEX IF NOT EXISTS "idx_rate_limits_status" ON "clients"."rate_limits"("status");

