-- incremental views for query layer
CREATE OR REPLACE VIEW "auth"."ListPhoneItems" AS
SELECT
  ph."id" AS "id",
  ph."account_id" AS "account_id",
  ph."phone" AS "phone",
  ph."is_primary" AS "is_primary",
  ph."verified_at" AS "verified_at",
  ph."created_at" AS "created_at",
  ph."updated_at" AS "updated_at"
FROM "auth"."Phones" ph
WHERE ph."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."ListPhoneItems" IS
'One row per non-deleted phone. Filter by account_id to list account phones.';

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
