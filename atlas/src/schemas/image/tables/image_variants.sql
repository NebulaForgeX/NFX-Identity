-- Image Variants Table
-- Purpose: Stores multiple versions/derivatives of an original image
-- Supports: thumbnails, resized versions, format conversions (webp/jpg/png), cropped versions
-- This enables efficient delivery of appropriately sized images for different use cases

CREATE TABLE "image"."image_variants" (
  -- Primary key
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  
  -- Reference to parent image (CASCADE delete: if image is deleted, all variants are deleted)
  "image_id" UUID NOT NULL REFERENCES "image"."images"("id") ON DELETE CASCADE,
  
  -- Variant type identifier (e.g., 'thumbnail', 'small', 'medium', 'large', 'webp', 'original')
  -- Combined with image_id forms unique constraint
  "variant_key" TEXT NOT NULL,
  
  -- Variant dimensions in pixels (may differ from original)
  "width" INT,
  "height" INT,
  
  -- Variant file size in bytes
  "size" BIGINT,
  
  -- Variant MIME type (may differ from original, e.g., webp conversion)
  "mime_type" TEXT,
  
  -- Storage path for this variant (filesystem path or S3 object key)
  "storage_path" TEXT NOT NULL,
  
  -- Public URL for accessing this variant (CDN URL or direct storage URL)
  "url" TEXT,
  
  -- Timestamps
  "created_at" TIMESTAMP DEFAULT NOW(),
  "updated_at" TIMESTAMP DEFAULT NOW(),
  
  -- Unique constraint: one variant of each type per image
  UNIQUE("image_id", "variant_key")
);

-- Indexes for common queries
CREATE INDEX "idx_image_variants_image_id" ON "image"."image_variants"("image_id");
CREATE INDEX "idx_image_variants_variant_key" ON "image"."image_variants"("variant_key");
CREATE INDEX "idx_image_variants_url" ON "image"."image_variants"("url") WHERE "url" IS NOT NULL;

-- Composite index for common query pattern (get all variants for an image)
CREATE INDEX "idx_image_variants_image_key" ON "image"."image_variants"("image_id", "variant_key");

-- Table and column comments
COMMENT ON TABLE "image"."image_variants" IS 'Stores derived versions of images: thumbnails, resized versions, format conversions';
COMMENT ON COLUMN "image"."image_variants"."image_id" IS 'Reference to parent image in images table';
COMMENT ON COLUMN "image"."image_variants"."variant_key" IS 'Variant type: thumbnail (150x150), small (300px), medium (720px), large (1080px), webp (converted), etc.';
COMMENT ON COLUMN "image"."image_variants"."width" IS 'Variant width in pixels';
COMMENT ON COLUMN "image"."image_variants"."height" IS 'Variant height in pixels';
COMMENT ON COLUMN "image"."image_variants"."size" IS 'Variant file size in bytes';
COMMENT ON COLUMN "image"."image_variants"."mime_type" IS 'Variant MIME type (may differ from original, e.g., webp conversion)';
COMMENT ON COLUMN "image"."image_variants"."storage_path" IS 'Storage backend path for this variant';
COMMENT ON COLUMN "image"."image_variants"."url" IS 'Public accessible URL for this variant';

