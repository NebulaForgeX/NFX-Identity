CREATE OR REPLACE VIEW "system"."SystemStateActiveView" AS
SELECT
  "id",
  "initialized",
  "initialized_at",
  "initialization_version",
  "last_reset_at",
  "last_reset_by",
  "reset_count",
  "metadata",
  "created_at",
  "updated_at"
FROM "system"."system_state";

COMMENT ON VIEW "system"."SystemStateActiveView" IS 'All system_state rows for latest-by-created_at reads.';
