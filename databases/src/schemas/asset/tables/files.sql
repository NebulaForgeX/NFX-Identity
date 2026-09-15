CREATE TABLE IF NOT EXISTS "asset"."Files" (
  "id" UUID PRIMARY KEY,
  "file_path" VARCHAR(500) NOT NULL,
  "file_name" VARCHAR(255) NOT NULL,
  "file_size" BIGINT NOT NULL,
  "mime_type" VARCHAR(100) NOT NULL,
  "uploader_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
CREATE INDEX IF NOT EXISTS "idx_files_file_path" ON "asset"."Files"("file_path");
CREATE INDEX IF NOT EXISTS "idx_files_uploader_id" ON "asset"."Files"("uploader_id");
CREATE INDEX IF NOT EXISTS "idx_files_mime_type" ON "asset"."Files"("mime_type");
CREATE INDEX IF NOT EXISTS "idx_files_deleted_at" ON "asset"."Files"("deleted_at");

COMMENT ON TABLE "asset"."Files" IS 'Non-image binary metadata (documents, archives, etc.); same soft-delete shape as Images without pixel fields.';
COMMENT ON COLUMN "asset"."Files"."id" IS 'Primary key UUID; typically matches stored object id.';
COMMENT ON COLUMN "asset"."Files"."file_path" IS 'Relative path from storage root or object key.';
COMMENT ON COLUMN "asset"."Files"."file_name" IS 'Original filename.';
COMMENT ON COLUMN "asset"."Files"."file_size" IS 'File size in bytes.';
COMMENT ON COLUMN "asset"."Files"."mime_type" IS 'MIME type as uploaded or detected.';
COMMENT ON COLUMN "asset"."Files"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
COMMENT ON COLUMN "asset"."Files"."created_at" IS 'Insert time.';
COMMENT ON COLUMN "asset"."Files"."updated_at" IS 'Last metadata update.';
COMMENT ON COLUMN "asset"."Files"."deleted_at" IS 'Soft-delete; NULL if object still active.';
