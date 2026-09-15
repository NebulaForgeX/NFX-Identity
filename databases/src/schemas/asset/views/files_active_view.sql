CREATE OR REPLACE VIEW "asset"."FilesActiveView" AS
SELECT
  "id", "file_path", "file_name", "file_size", "mime_type",
  "uploader_id", "created_at", "updated_at"
FROM "asset"."Files"
WHERE "deleted_at" IS NULL;

COMMENT ON VIEW "asset"."FilesActiveView" IS 'Non-deleted file rows only; omits deleted_at column.';
