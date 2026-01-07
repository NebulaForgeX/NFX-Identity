-- Token Usage Logs table for audit and analytics
-- Records each token usage for security audit, analytics, and abuse detection
CREATE TABLE IF NOT EXISTS "clients"."token_usage_logs" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "token_id" VARCHAR(255) NOT NULL REFERENCES "clients"."service_tokens"("token_id") ON DELETE CASCADE, -- References service_tokens.token_id
  "app_id" UUID NOT NULL REFERENCES "clients"."apps"("id") ON DELETE CASCADE,
  "client_id" VARCHAR(255) REFERENCES "clients"."client_credentials"("client_id") ON DELETE SET NULL,
  "endpoint" VARCHAR(255), -- API endpoint accessed: "/api/v1/users", "/api/v1/users/count"
  "method" VARCHAR(10), -- HTTP method: "GET", "POST", etc.
  "ip" INET,
  "ua_hash" VARCHAR(64), -- User agent hash
  "status_code" INTEGER, -- HTTP status code: 200, 401, 403, etc.
  "response_time_ms" INTEGER, -- Response time in milliseconds
  "request_size" INTEGER, -- Request size in bytes
  "response_size" INTEGER, -- Response size in bytes
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_token_usage_logs_token_id" ON "clients"."token_usage_logs"("token_id", "created_at");
CREATE INDEX IF NOT EXISTS "idx_token_usage_logs_app_id" ON "clients"."token_usage_logs"("app_id", "created_at");
CREATE INDEX IF NOT EXISTS "idx_token_usage_logs_client_id" ON "clients"."token_usage_logs"("client_id", "created_at");
CREATE INDEX IF NOT EXISTS "idx_token_usage_logs_created_at" ON "clients"."token_usage_logs"("created_at");
CREATE INDEX IF NOT EXISTS "idx_token_usage_logs_endpoint" ON "clients"."token_usage_logs"("endpoint", "created_at");

