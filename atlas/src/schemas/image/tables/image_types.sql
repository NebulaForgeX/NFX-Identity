-- Image Types Table
-- Purpose: Defines image usage categories and specifications (avatar, background, product_cover, etc.)
-- This table allows different business domains to have different image requirements
-- Example types: avatar (1:1), profile_background (16:9), product_cover, product_gallery, post_image, banner, badge_icon

CREATE TABLE "image"."image_types" (
  -- Primary key
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  
  -- Type identifier (e.g., 'avatar', 'background', 'product_cover')
  -- Must be unique across all types
  "key" TEXT NOT NULL UNIQUE,
  
  -- Human-readable description of the image type
  "description" TEXT,
  
  -- Maximum allowed dimensions (for validation during upload)
  "max_width" INT,
  "max_height" INT,
  
  -- Aspect ratio hint (e.g., '1:1', '16:9', '4:3') for frontend auto-cropping
  "aspect_ratio" TEXT,
  
  -- System types (is_system = TRUE) are protected from deletion
  -- Custom types can be created by users/applications
  "is_system" BOOLEAN DEFAULT FALSE,
  
  -- Timestamps
  "created_at" TIMESTAMP DEFAULT NOW(),
  "updated_at" TIMESTAMP DEFAULT NOW()
);

-- Indexes for common queries
CREATE INDEX "idx_image_types_key" ON "image"."image_types"("key");
CREATE INDEX "idx_image_types_is_system" ON "image"."image_types"("is_system");

-- Table and column comments
COMMENT ON TABLE "image"."image_types" IS 'Image type definitions for different use cases and business domains';
COMMENT ON COLUMN "image"."image_types"."key" IS 'Type key identifier: avatar, background, product_cover, product_gallery, post_image, banner, badge_icon, etc.';
COMMENT ON COLUMN "image"."image_types"."description" IS 'Human-readable description of what this image type is used for';
COMMENT ON COLUMN "image"."image_types"."max_width" IS 'Maximum allowed width in pixels (NULL means no limit)';
COMMENT ON COLUMN "image"."image_types"."max_height" IS 'Maximum allowed height in pixels (NULL means no limit)';
COMMENT ON COLUMN "image"."image_types"."aspect_ratio" IS 'Preferred aspect ratio hint: 1:1, 16:9, 4:3, etc. Used for frontend auto-cropping';
COMMENT ON COLUMN "image"."image_types"."is_system" IS 'TRUE for system-built types (cannot be deleted), FALSE for custom types';

