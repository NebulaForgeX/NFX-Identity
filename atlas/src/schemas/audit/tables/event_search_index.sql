-- Event Search Index table: Query-optimized redundant index table
-- For fast filtering in audit admin pages: by actor, action, target, result, time range
-- Solves slow filtering on audit.events jsonb metadata
-- Allows extracting hot fields and even lightweight inverted index (if not using ELK)
-- Note: This is a "query acceleration table", can be written asynchronously by pipeline
-- Allows slight delay (seconds/minutes) compared to events table (acceptable)
CREATE TABLE IF NOT EXISTS "audit"."event_search_index" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "event_id" VARCHAR(255) NOT NULL UNIQUE, -- References audit.events.event_id (application-level consistency)
  "tenant_id" UUID, -- Multi-tenant isolation
  "app_id" UUID, -- App isolation
  "actor_type" "audit".actor_type NOT NULL,
  "actor_id" UUID NOT NULL,
  "action" VARCHAR(255) NOT NULL,
  "target_type" VARCHAR(100),
  "target_id" UUID,
  "result" "audit".result_type NOT NULL,
  "occurred_at" TIMESTAMP NOT NULL,
  "ip" INET,
  "risk_level" "audit".risk_level NOT NULL DEFAULT 'low',
  "data_classification" "audit".data_classification NOT NULL DEFAULT 'internal',
  "tags" TEXT[], -- Optional: ["security", "admin", "export", "sensitive", ...]
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes optimized for search queries
CREATE INDEX IF NOT EXISTS "idx_event_search_index_event_id" ON "audit"."event_search_index"("event_id");
CREATE INDEX IF NOT EXISTS "idx_event_search_index_tenant_occurred" ON "audit"."event_search_index"("tenant_id", "occurred_at");
CREATE INDEX IF NOT EXISTS "idx_event_search_index_actor" ON "audit"."event_search_index"("actor_type", "actor_id", "occurred_at");
CREATE INDEX IF NOT EXISTS "idx_event_search_index_action" ON "audit"."event_search_index"("action", "occurred_at");
CREATE INDEX IF NOT EXISTS "idx_event_search_index_target" ON "audit"."event_search_index"("target_type", "target_id");
CREATE INDEX IF NOT EXISTS "idx_event_search_index_result" ON "audit"."event_search_index"("result", "occurred_at");
CREATE INDEX IF NOT EXISTS "idx_event_search_index_risk_level" ON "audit"."event_search_index"("risk_level", "occurred_at");
CREATE INDEX IF NOT EXISTS "idx_event_search_index_tags" ON "audit"."event_search_index" USING GIN("tags"); -- GIN index for array search

