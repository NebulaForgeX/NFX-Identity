CREATE OR REPLACE VIEW "asset"."AudiosActiveView" AS
SELECT
  "id", "file_path", "file_name", "file_size", "mime_type",
  "duration_seconds", "uploader_id", "created_at", "updated_at"
FROM "asset"."Audios"
WHERE "deleted_at" IS NULL;

COMMENT ON VIEW "asset"."AudiosActiveView" IS 'Non-deleted audio rows only; omits deleted_at column.';
