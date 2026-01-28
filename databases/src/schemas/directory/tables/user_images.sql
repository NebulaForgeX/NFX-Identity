-- User Images table for managing user images (backgrounds, gallery, etc.)
-- One user can have multiple images with ordering
-- display_order = 0 means the current/active image (for types like background)
-- Note: image_id references image.images.id (no foreign key to maintain service independence)
-- Image type is stored in image.images.type_id, no need to duplicate here
CREATE TABLE IF NOT EXISTS "directory"."user_images" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID NOT NULL REFERENCES "directory"."users"("id") ON DELETE CASCADE,
  "image_id" UUID NOT NULL, -- References image.images.id (application-level consistency)
  "display_order" INTEGER NOT NULL DEFAULT 0, -- Display order (0 = current/active, higher = later)
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

-- Indexes for common queries
CREATE INDEX IF NOT EXISTS "idx_user_images_user_id" ON "directory"."user_images"("user_id");
CREATE INDEX IF NOT EXISTS "idx_user_images_image_id" ON "directory"."user_images"("image_id");
CREATE INDEX IF NOT EXISTS "idx_user_images_display_order" ON "directory"."user_images"("user_id", "display_order") WHERE "deleted_at" IS NULL;
CREATE INDEX IF NOT EXISTS "idx_user_images_deleted_at" ON "directory"."user_images"("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_user_images_created_at" ON "directory"."user_images"("created_at");

-- Composite index for common query: get all images for a user, ordered
CREATE INDEX IF NOT EXISTS "idx_user_images_user_order" 
  ON "directory"."user_images"("user_id", "display_order") 
  WHERE "deleted_at" IS NULL;

-- Unique constraint: prevent duplicate image for the same user
CREATE UNIQUE INDEX IF NOT EXISTS "idx_user_images_user_image_unique" 
  ON "directory"."user_images"("user_id", "image_id") 
  WHERE "deleted_at" IS NULL;

-- Table and column comments
COMMENT ON TABLE "directory"."user_images" IS 'User images table - stores all user images with display order. Image type is stored in image.images.type_id';
COMMENT ON COLUMN "directory"."user_images"."user_id" IS 'Reference to directory.users.id';
COMMENT ON COLUMN "directory"."user_images"."image_id" IS 'Reference to image.images.id (application-level consistency, no foreign key). Image type can be queried via image.images.type_id';
COMMENT ON COLUMN "directory"."user_images"."display_order" IS 'Display order: 0 = current/active image, higher numbers = later in sequence';
COMMENT ON COLUMN "directory"."user_images"."deleted_at" IS 'Soft delete timestamp (NULL = active, NOT NULL = deleted)';
