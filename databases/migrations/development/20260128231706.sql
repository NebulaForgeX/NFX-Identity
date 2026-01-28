-- Modify "user_profiles" table
ALTER TABLE "directory"."user_profiles" DROP COLUMN "avatar_id", DROP COLUMN "background_id", DROP COLUMN "background_ids";
-- Drop index "idx_image_variants_image_key" from table: "image_variants"
DROP INDEX "image"."idx_image_variants_image_key";
-- Create "user_avatars" table
CREATE TABLE "directory"."user_avatars" (
  "user_id" uuid NOT NULL,
  "image_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("user_id"),
  CONSTRAINT "user_avatars_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_avatars_image_id" to table: "user_avatars"
CREATE INDEX "idx_user_avatars_image_id" ON "directory"."user_avatars" ("image_id");
-- Set comment to table: "user_avatars"
COMMENT ON TABLE "directory"."user_avatars" IS 'User avatar table - stores current avatar for each user (one-to-one relationship)';
-- Set comment to column: "user_id" on table: "user_avatars"
COMMENT ON COLUMN "directory"."user_avatars"."user_id" IS 'Reference to directory.users.id (primary key, one-to-one)';
-- Set comment to column: "image_id" on table: "user_avatars"
COMMENT ON COLUMN "directory"."user_avatars"."image_id" IS 'Reference to image.images.id (application-level consistency, no foreign key)';
-- Create "user_images" table
CREATE TABLE "directory"."user_images" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "image_id" uuid NOT NULL,
  "display_order" integer NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "user_images_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "directory"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_user_images_created_at" to table: "user_images"
CREATE INDEX "idx_user_images_created_at" ON "directory"."user_images" ("created_at");
-- Create index "idx_user_images_deleted_at" to table: "user_images"
CREATE INDEX "idx_user_images_deleted_at" ON "directory"."user_images" ("deleted_at");
-- Create index "idx_user_images_display_order" to table: "user_images"
CREATE INDEX "idx_user_images_display_order" ON "directory"."user_images" ("user_id", "display_order") WHERE (deleted_at IS NULL);
-- Create index "idx_user_images_image_id" to table: "user_images"
CREATE INDEX "idx_user_images_image_id" ON "directory"."user_images" ("image_id");
-- Create index "idx_user_images_user_id" to table: "user_images"
CREATE INDEX "idx_user_images_user_id" ON "directory"."user_images" ("user_id");
-- Create index "idx_user_images_user_image_unique" to table: "user_images"
CREATE UNIQUE INDEX "idx_user_images_user_image_unique" ON "directory"."user_images" ("user_id", "image_id") WHERE (deleted_at IS NULL);
-- Create index "idx_user_images_user_order" to table: "user_images"
CREATE INDEX "idx_user_images_user_order" ON "directory"."user_images" ("user_id", "display_order") WHERE (deleted_at IS NULL);
-- Set comment to table: "user_images"
COMMENT ON TABLE "directory"."user_images" IS 'User images table - stores all user images with display order. Image type is stored in image.images.type_id';
-- Set comment to column: "user_id" on table: "user_images"
COMMENT ON COLUMN "directory"."user_images"."user_id" IS 'Reference to directory.users.id';
-- Set comment to column: "image_id" on table: "user_images"
COMMENT ON COLUMN "directory"."user_images"."image_id" IS 'Reference to image.images.id (application-level consistency, no foreign key). Image type can be queried via image.images.type_id';
-- Set comment to column: "display_order" on table: "user_images"
COMMENT ON COLUMN "directory"."user_images"."display_order" IS 'Display order: 0 = current/active image, higher numbers = later in sequence';
-- Set comment to column: "deleted_at" on table: "user_images"
COMMENT ON COLUMN "directory"."user_images"."deleted_at" IS 'Soft delete timestamp (NULL = active, NOT NULL = deleted)';
