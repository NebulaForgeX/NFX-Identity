-- User Avatars table for storing user avatar
-- One user has one avatar (one-to-one relationship)
-- Note: image_id references image.images.id (no foreign key to maintain service independence)
CREATE TABLE IF NOT EXISTS "directory"."user_avatars" (
  "user_id" UUID PRIMARY KEY REFERENCES "directory"."users"("id") ON DELETE CASCADE,
  "image_id" UUID NOT NULL, -- References image.images.id (application-level consistency)
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for common queries
CREATE INDEX IF NOT EXISTS "idx_user_avatars_image_id" ON "directory"."user_avatars"("image_id");

-- Table and column comments
COMMENT ON TABLE "directory"."user_avatars" IS 'User avatar table - stores current avatar for each user (one-to-one relationship)';
COMMENT ON COLUMN "directory"."user_avatars"."user_id" IS 'Reference to directory.users.id (primary key, one-to-one)';
COMMENT ON COLUMN "directory"."user_avatars"."image_id" IS 'Reference to image.images.id (application-level consistency, no foreign key)';
