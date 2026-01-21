-- Audit Events table: Core audit event fact table
-- Stores every audit event record (cannot be lost)
-- This is the source data for all audit queries, compliance exports, and post-incident forensics
-- Must answer: who did what, to what, when, where, result, risk level
CREATE TYPE "audit".actor_type AS ENUM ('user', 'service', 'system', 'admin');
CREATE TYPE "audit".result_type AS ENUM ('success', 'failure', 'deny', 'error');
CREATE TYPE "audit".risk_level AS ENUM ('low', 'medium', 'high', 'critical');
CREATE TYPE "audit".data_classification AS ENUM ('public', 'internal', 'confidential', 'restricted');

CREATE TABLE IF NOT EXISTS "audit"."events" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "event_id" VARCHAR(255) NOT NULL UNIQUE, -- Event identifier (ULID recommended for sort-friendly)
  "occurred_at" TIMESTAMP NOT NULL, -- Event occurrence time (not just write time)
  "received_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Time written to audit database (for delay investigation)
  "tenant_id" UUID, -- Multi-tenant isolation (references tenants.tenants.id, application-level consistency)
  "app_id" UUID, -- App isolation (references clients.apps.id, application-level consistency)
  "actor_type" "audit".actor_type NOT NULL, -- user|service|system|admin
  "actor_id" UUID NOT NULL, -- user_id or client_id (references directory.users.id or clients.apps.id, application-level consistency)
  "actor_tenant_member_id" UUID, -- Optional: backend member identity (if distinguishing member)
  "action" VARCHAR(255) NOT NULL, -- Action enum string: "user.login", "users.list", "grants.update", "clients.secret.rotate", etc.
  "target_type" VARCHAR(100), -- Target resource type: "user", "tenant", "client", "role", "asset", "token", "export_job", etc.
  "target_id" UUID, -- Target resource ID (nullable, e.g., for "list query")
  "result" "audit".result_type NOT NULL, -- success|failure|deny|error
  "failure_reason_code" VARCHAR(100), -- Failure reason: "INVALID_PASSWORD", "INSUFFICIENT_SCOPE", "RATE_LIMITED", etc.
  "http_method" VARCHAR(10), -- HTTP method: "GET", "POST", "PUT", "DELETE", etc.
  "http_path" VARCHAR(500), -- HTTP path: "/api/v1/users", "/api/v1/users/count", etc.
  "http_status" INTEGER, -- HTTP status code: 200, 401, 403, 500, etc.
  "request_id" VARCHAR(255), -- Request ID for request tracing (very important)
  "trace_id" VARCHAR(255), -- Trace ID for distributed tracing
  "ip" INET, -- IP address
  "user_agent" TEXT, -- User agent string
  "geo_country" VARCHAR(10), -- Optional: country code from IP geolocation
  "risk_level" "audit".risk_level NOT NULL DEFAULT 'low', -- low|medium|high|critical
  "data_classification" "audit".data_classification NOT NULL DEFAULT 'internal', -- public|internal|confidential|restricted
  "prev_hash" VARCHAR(64), -- Previous event hash for tamper-proof chain
  "event_hash" VARCHAR(64), -- Current event hash: SHA256(prev_hash + canonical_json(event_fields))
  "metadata" JSONB DEFAULT '{}'::jsonb, -- Extended fields (don't put all key fields in metadata)
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Core indexes for common queries
CREATE INDEX IF NOT EXISTS "idx_events_event_id" ON "audit"."events"("event_id");
CREATE INDEX IF NOT EXISTS "idx_events_occurred_at" ON "audit"."events"("occurred_at");
CREATE INDEX IF NOT EXISTS "idx_events_tenant_id" ON "audit"."events"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_events_app_id" ON "audit"."events"("app_id");
CREATE INDEX IF NOT EXISTS "idx_events_actor" ON "audit"."events"("actor_type", "actor_id");
CREATE INDEX IF NOT EXISTS "idx_events_action" ON "audit"."events"("action");
CREATE INDEX IF NOT EXISTS "idx_events_target" ON "audit"."events"("target_type", "target_id");
CREATE INDEX IF NOT EXISTS "idx_events_result" ON "audit"."events"("result");
CREATE INDEX IF NOT EXISTS "idx_events_risk_level" ON "audit"."events"("risk_level");
CREATE INDEX IF NOT EXISTS "idx_events_data_classification" ON "audit"."events"("data_classification");

-- Composite indexes for common query patterns
CREATE INDEX IF NOT EXISTS "idx_events_tenant_occurred" ON "audit"."events"("tenant_id", "occurred_at");
CREATE INDEX IF NOT EXISTS "idx_events_actor_occurred" ON "audit"."events"("actor_type", "actor_id", "occurred_at");
CREATE INDEX IF NOT EXISTS "idx_events_action_occurred" ON "audit"."events"("action", "occurred_at");
CREATE INDEX IF NOT EXISTS "idx_events_tenant_action" ON "audit"."events"("tenant_id", "action", "occurred_at");
CREATE INDEX IF NOT EXISTS "idx_events_request_id" ON "audit"."events"("request_id");
CREATE INDEX IF NOT EXISTS "idx_events_trace_id" ON "audit"."events"("trace_id");

-- Partitioning hint: This table should be partitioned by occurred_at (monthly or weekly)
-- Example: PARTITION BY RANGE (occurred_at)
-- Partition management: detach old partitions, archive for compliance retention

