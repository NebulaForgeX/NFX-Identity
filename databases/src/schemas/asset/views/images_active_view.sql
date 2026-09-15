CREATE OR REPLACE VIEW "asset"."ImagesActiveView" AS
SELECT
  "id", "file_path", "file_name", "file_size", "mime_type",
  "width", "height", "alt_text", "uploader_id", "created_at", "updated_at"
FROM "asset"."Images"
WHERE "deleted_at" IS NULL;

COMMENT ON VIEW "asset"."ImagesActiveView" IS 'Non-deleted image rows only; omits deleted_at column.';
