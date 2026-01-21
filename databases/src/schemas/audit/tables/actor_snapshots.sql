-- Actor Snapshots table: Actor snapshots for forensics
-- When user/client is deleted, renamed, or tenant migrated, audit records remain readable and forensically valid
-- Avoids cross-service lookups for "who was this actor at that time" in every audit query
-- Enterprise requirement: Audit goal is "forensic evidence", evidence cannot depend on external data still existing
CREATE TABLE IF NOT EXISTS "audit"."actor_snapshots" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "actor_type" "audit".actor_type NOT NULL,
  "actor_id" UUID NOT NULL,
  "display_name" VARCHAR(255), -- Display name at that time
  "email" VARCHAR(255), -- Email at that time (may be masked)
  "client_name" VARCHAR(255), -- Client/app name (for service actors)
  "tenant_id" UUID, -- Tenant at that time (references tenants.tenants.id, application-level consistency)
  "snapshot_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- When this snapshot was taken
  "snapshot_data" JSONB DEFAULT '{}'::jsonb, -- Extended snapshot data (may be masked)
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE("actor_type", "actor_id", "snapshot_at")
);

CREATE INDEX IF NOT EXISTS "idx_actor_snapshots_actor" ON "audit"."actor_snapshots"("actor_type", "actor_id", "snapshot_at" DESC);
CREATE INDEX IF NOT EXISTS "idx_actor_snapshots_tenant_id" ON "audit"."actor_snapshots"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_actor_snapshots_snapshot_at" ON "audit"."actor_snapshots"("snapshot_at");

