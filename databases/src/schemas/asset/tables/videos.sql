CREATE TABLE IF NOT EXISTS "asset"."Videos" (
  "id" UUID PRIMARY KEY,
  "file_path" VARCHAR(500) NOT NULL,
  "file_name" VARCHAR(255) NOT NULL,
  "file_size" BIGINT NOT NULL,
  "mime_type" VARCHAR(100) NOT NULL,
  "duration_seconds" DOUBLE PRECISION,
  "width" INTEGER,
  "height" INTEGER,
  "uploader_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
CREATE INDEX IF NOT EXISTS "idx_videos_file_path" ON "asset"."Videos"("file_path");
CREATE INDEX IF NOT EXISTS "idx_videos_uploader_id" ON "asset"."Videos"("uploader_id");
CREATE INDEX IF NOT EXISTS "idx_videos_mime_type" ON "asset"."Videos"("mime_type");
CREATE INDEX IF NOT EXISTS "idx_videos_deleted_at" ON "asset"."Videos"("deleted_at");

COMMENT ON TABLE "asset"."Videos" IS 'Video binary metadata; duration and optional display dimensions; uploader_id references auth.Accounts at app level.';
COMMENT ON COLUMN "asset"."Videos"."id" IS 'Primary key UUID; typically matches stored object id.';
COMMENT ON COLUMN "asset"."Videos"."file_path" IS 'Relative path from storage root or object key.';
COMMENT ON COLUMN "asset"."Videos"."file_name" IS 'Original filename.';
COMMENT ON COLUMN "asset"."Videos"."file_size" IS 'File size in bytes.';
COMMENT ON COLUMN "asset"."Videos"."mime_type" IS 'MIME type as uploaded or detected.';
COMMENT ON COLUMN "asset"."Videos"."duration_seconds" IS 'Media duration when known.';
COMMENT ON COLUMN "asset"."Videos"."width" IS 'Pixel width if known.';
COMMENT ON COLUMN "asset"."Videos"."height" IS 'Pixel height if known.';
COMMENT ON COLUMN "asset"."Videos"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
COMMENT ON COLUMN "asset"."Videos"."created_at" IS 'Insert time.';
COMMENT ON COLUMN "asset"."Videos"."updated_at" IS 'Last metadata update.';
COMMENT ON COLUMN "asset"."Videos"."deleted_at" IS 'Soft-delete; NULL if object still active.';
