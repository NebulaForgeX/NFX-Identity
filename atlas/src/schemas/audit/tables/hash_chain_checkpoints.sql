-- Hash Chain Checkpoints table: Tamper-evident checkpoints
-- Prevents admins/DBAs from directly modifying audit.events content without leaving traces
-- Important for compliance/security incident investigations
-- Note: This is not absolute tamper-proof (DB admin can still modify chain), but significantly increases tampering cost
-- True stronger protection: WORM storage/external audit warehouse, but lightweight version can be done first
CREATE TABLE IF NOT EXISTS "audit"."hash_chain_checkpoints" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "checkpoint_id" VARCHAR(255) NOT NULL UNIQUE, -- Checkpoint identifier
  "tenant_id" UUID, -- Optional: tenant-level checkpoint (NULL means global)
  "partition_date" DATE NOT NULL, -- Partition date (for daily/weekly/monthly checkpoints)
  "checkpoint_hash" VARCHAR(64) NOT NULL, -- Checkpoint hash: SHA256 of all events in this partition/period
  "prev_checkpoint_hash" VARCHAR(64), -- Previous checkpoint hash (chain link)
  "event_count" INTEGER NOT NULL, -- Number of events in this checkpoint
  "first_event_id" VARCHAR(255), -- First event ID in this checkpoint
  "last_event_id" VARCHAR(255), -- Last event ID in this checkpoint
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" VARCHAR(255) -- System/service that created this checkpoint
);

CREATE INDEX IF NOT EXISTS "idx_hash_chain_checkpoints_checkpoint_id" ON "audit"."hash_chain_checkpoints"("checkpoint_id");
CREATE INDEX IF NOT EXISTS "idx_hash_chain_checkpoints_tenant_id" ON "audit"."hash_chain_checkpoints"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_hash_chain_checkpoints_partition_date" ON "audit"."hash_chain_checkpoints"("partition_date" DESC);
CREATE INDEX IF NOT EXISTS "idx_hash_chain_checkpoints_prev_hash" ON "audit"."hash_chain_checkpoints"("prev_checkpoint_hash");

