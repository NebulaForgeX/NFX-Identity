CREATE TABLE IF NOT EXISTS "asset"."Audios" (
  "id" UUID PRIMARY KEY,
  "file_path" VARCHAR(500) NOT NULL,
  "file_name" VARCHAR(255) NOT NULL,
  "file_size" BIGINT NOT NULL,
  "mime_type" VARCHAR(100) NOT NULL,
  "duration_seconds" DOUBLE PRECISION,
  "uploader_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
CREATE INDEX IF NOT EXISTS "idx_audios_file_path" ON "asset"."Audios"("file_path");
CREATE INDEX IF NOT EXISTS "idx_audios_uploader_id" ON "asset"."Audios"("uploader_id");
CREATE INDEX IF NOT EXISTS "idx_audios_mime_type" ON "asset"."Audios"("mime_type");
CREATE INDEX IF NOT EXISTS "idx_audios_deleted_at" ON "asset"."Audios"("deleted_at");

COMMENT ON TABLE "asset"."Audios" IS 'Audio binary metadata; duration when known; uploader_id references auth.Accounts at app level.';
COMMENT ON COLUMN "asset"."Audios"."id" IS 'Primary key UUID; typically matches stored object id.';
COMMENT ON COLUMN "asset"."Audios"."file_path" IS 'Relative path from storage root or object key.';
COMMENT ON COLUMN "asset"."Audios"."file_name" IS 'Original filename.';
COMMENT ON COLUMN "asset"."Audios"."file_size" IS 'File size in bytes.';
COMMENT ON COLUMN "asset"."Audios"."mime_type" IS 'MIME type as uploaded or detected.';
COMMENT ON COLUMN "asset"."Audios"."duration_seconds" IS 'Media duration when known.';
COMMENT ON COLUMN "asset"."Audios"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
COMMENT ON COLUMN "asset"."Audios"."created_at" IS 'Insert time.';
COMMENT ON COLUMN "asset"."Audios"."updated_at" IS 'Last metadata update.';
COMMENT ON COLUMN "asset"."Audios"."deleted_at" IS 'Soft-delete; NULL if object still active.';
