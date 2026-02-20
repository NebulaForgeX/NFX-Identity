CREATE TYPE "clients".rate_limit_type AS ENUM ('requests_per_second', 'requests_per_minute', 'requests_per_hour', 'requests_per_day');

CREATE TABLE IF NOT EXISTS "clients"."rate_limits" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "application_id" UUID NOT NULL REFERENCES "clients"."applications"("id") ON DELETE CASCADE,
  "limit_type" "clients".rate_limit_type NOT NULL DEFAULT 'requests_per_minute',
  "limit_value" INTEGER NOT NULL,
  "window_seconds" INTEGER NOT NULL,
  "description" TEXT,
  "status" VARCHAR(50) NOT NULL DEFAULT 'active',
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_by" UUID,
  UNIQUE("application_id", "limit_type")
);

CREATE INDEX IF NOT EXISTS "idx_rate_limits_application_id" ON "clients"."rate_limits"("application_id");
CREATE INDEX IF NOT EXISTS "idx_rate_limits_status" ON "clients"."rate_limits"("status");
