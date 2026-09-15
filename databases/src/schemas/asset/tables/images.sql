CREATE TABLE IF NOT EXISTS "asset"."Images" (
  "id" UUID PRIMARY KEY,
  "file_path" VARCHAR(500) NOT NULL,
  "file_name" VARCHAR(255) NOT NULL,
  "file_size" BIGINT NOT NULL,
  "mime_type" VARCHAR(100) NOT NULL,
  "width" INTEGER,
  "height" INTEGER,
  "alt_text" VARCHAR(255),
  "uploader_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
CREATE INDEX IF NOT EXISTS "idx_images_file_path" ON "asset"."Images"("file_path");
CREATE INDEX IF NOT EXISTS "idx_images_uploader_id" ON "asset"."Images"("uploader_id");
CREATE INDEX IF NOT EXISTS "idx_images_mime_type" ON "asset"."Images"("mime_type");
CREATE INDEX IF NOT EXISTS "idx_images_deleted_at" ON "asset"."Images"("deleted_at");

COMMENT ON TABLE "asset"."Images" IS 'Raster image metadata; path is relative to storage root or object key; uploader_id references auth.Accounts at app level.';
COMMENT ON COLUMN "asset"."Images"."id" IS 'Primary key UUID; typically matches stored object id.';
COMMENT ON COLUMN "asset"."Images"."file_path" IS 'Relative path from storage root or object key.';
COMMENT ON COLUMN "asset"."Images"."file_name" IS 'Original filename.';
COMMENT ON COLUMN "asset"."Images"."file_size" IS 'File size in bytes.';
COMMENT ON COLUMN "asset"."Images"."mime_type" IS 'MIME type as uploaded or detected.';
COMMENT ON COLUMN "asset"."Images"."width" IS 'Pixel width if known.';
COMMENT ON COLUMN "asset"."Images"."height" IS 'Pixel height if known.';
COMMENT ON COLUMN "asset"."Images"."alt_text" IS 'Accessibility / caption text.';
COMMENT ON COLUMN "asset"."Images"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
COMMENT ON COLUMN "asset"."Images"."created_at" IS 'Insert time.';
COMMENT ON COLUMN "asset"."Images"."updated_at" IS 'Last metadata update.';
COMMENT ON COLUMN "asset"."Images"."deleted_at" IS 'Soft-delete; NULL if object still active.';
