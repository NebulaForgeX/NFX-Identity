-- Image Tags Table
-- Purpose: Stores tags for images (AI-generated labels and user-defined tags)
-- Enables content-based search and discovery (e.g., find all images tagged with 'cat', 'shoes', 'selfie')
-- Supports both manual user tagging and automated AI labeling

CREATE TABLE "image"."image_tags" (
  -- Primary key
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  
  -- Reference to image (CASCADE delete: if image is deleted, all tags are deleted)
  "image_id" UUID NOT NULL REFERENCES "image"."images"("id") ON DELETE CASCADE,
  
  -- Tag text (normalized to lowercase recommended)
  -- Examples: 'cat', 'shoes', 'selfie', 'nature', 'food', 'portrait'
  "tag" TEXT NOT NULL,
  
  -- Confidence score for AI-generated tags (0.0 to 1.0)
  -- NULL for user-defined tags (which are assumed to be 100% accurate)
  -- Used for ranking search results and filtering low-confidence AI tags
  "confidence" FLOAT,
  
  -- Timestamp when tag was created
  "created_at" TIMESTAMP DEFAULT NOW(),
  
  -- Unique constraint: one tag per image (prevents duplicate tags)
  UNIQUE("image_id", "tag")
);

-- Indexes for common queries
CREATE INDEX "idx_image_tags_image_id" ON "image"."image_tags"("image_id");
CREATE INDEX "idx_image_tags_tag" ON "image"."image_tags"("tag");
CREATE INDEX "idx_image_tags_confidence" ON "image"."image_tags"("confidence") WHERE "confidence" IS NOT NULL;

-- Composite index for tag search queries
CREATE INDEX "idx_image_tags_tag_confidence" ON "image"."image_tags"("tag", "confidence" DESC);

-- Table and column comments
COMMENT ON TABLE "image"."image_tags" IS 'Image tags for content-based search: AI labels and user-defined tags';
COMMENT ON COLUMN "image"."image_tags"."image_id" IS 'Reference to image in images table';
COMMENT ON COLUMN "image"."image_tags"."tag" IS 'Tag text (normalized, e.g., lowercase). Examples: cat, shoes, selfie, nature';
COMMENT ON COLUMN "image"."image_tags"."confidence" IS 'Confidence score 0.0-1.0 for AI-generated tags, NULL for user tags';
COMMENT ON COLUMN "image"."image_tags"."created_at" IS 'When this tag was added to the image';

