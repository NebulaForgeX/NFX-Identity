CREATE OR REPLACE VIEW "asset"."VideosActiveView" AS
SELECT
  "id", "file_path", "file_name", "file_size", "mime_type",
  "duration_seconds", "width", "height", "uploader_id", "created_at", "updated_at"
FROM "asset"."Videos"
WHERE "deleted_at" IS NULL;

COMMENT ON VIEW "asset"."VideosActiveView" IS 'Non-deleted video rows only; omits deleted_at column.';
