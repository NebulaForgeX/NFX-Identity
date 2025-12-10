-- Add new schema named "auth"
CREATE SCHEMA "auth";
-- Set comment to schema: "auth"
COMMENT ON SCHEMA "auth" IS 'Authentication schema';
-- Add new schema named "image"
CREATE SCHEMA "image";
-- Set comment to schema: "image"
COMMENT ON SCHEMA "image" IS 'Image schema';
-- Create enum type "user_status"
CREATE TYPE "auth"."user_status" AS ENUM ('pending', 'active', 'deactive');
-- Create "roles" table
CREATE TABLE "auth"."roles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(50) NOT NULL,
  "description" text NULL,
  "permissions" jsonb NULL DEFAULT '[]',
  "is_system" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "roles_name_key" UNIQUE ("name")
);
-- Create index "idx_roles_deleted_at" to table: "roles"
CREATE INDEX "idx_roles_deleted_at" ON "auth"."roles" ("deleted_at");
-- Create index "idx_roles_name" to table: "roles"
CREATE INDEX "idx_roles_name" ON "auth"."roles" ("name");
-- Create "users" table
CREATE TABLE "auth"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "username" character varying(50) NOT NULL,
  "email" character varying(255) NOT NULL,
  "phone" character varying(20) NOT NULL,
  "password_hash" character varying(255) NOT NULL,
  "role_id" uuid NULL,
  "status" "auth"."user_status" NOT NULL DEFAULT 'pending',
  "is_verified" boolean NOT NULL DEFAULT false,
  "last_login_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "users_email_key" UNIQUE ("email"),
  CONSTRAINT "users_phone_key" UNIQUE ("phone"),
  CONSTRAINT "users_username_key" UNIQUE ("username"),
  CONSTRAINT "users_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "auth"."roles" ("id") ON UPDATE NO ACTION ON DELETE SET NULL
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "auth"."users" ("deleted_at");
-- Create index "idx_users_email" to table: "users"
CREATE INDEX "idx_users_email" ON "auth"."users" ("email");
-- Create index "idx_users_phone" to table: "users"
CREATE INDEX "idx_users_phone" ON "auth"."users" ("phone");
-- Create index "idx_users_role_id" to table: "users"
CREATE INDEX "idx_users_role_id" ON "auth"."users" ("role_id");
-- Create index "idx_users_status" to table: "users"
CREATE INDEX "idx_users_status" ON "auth"."users" ("status");
-- Create index "idx_users_username" to table: "users"
CREATE INDEX "idx_users_username" ON "auth"."users" ("username");
-- Create "profiles" table
CREATE TABLE "auth"."profiles" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "first_name" character varying(100) NULL,
  "last_name" character varying(100) NULL,
  "nickname" character varying(50) NULL,
  "display_name" character varying(100) NULL,
  "avatar_id" uuid NULL,
  "background_id" uuid NULL,
  "background_ids" uuid[] NULL,
  "bio" text NULL,
  "phone" character varying(20) NULL,
  "birthday" date NULL,
  "age" integer NULL,
  "gender" character varying(20) NULL,
  "location" character varying(255) NULL,
  "website" character varying(255) NULL,
  "github" character varying(255) NULL,
  "social_links" jsonb NULL DEFAULT '{}',
  "preferences" jsonb NULL DEFAULT '{}',
  "skills" jsonb NULL DEFAULT '{}',
  "privacy_settings" jsonb NULL DEFAULT '{}',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "profiles_nickname_key" UNIQUE ("nickname"),
  CONSTRAINT "profiles_user_id_key" UNIQUE ("user_id"),
  CONSTRAINT "profiles_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "auth"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_profiles_avatar_id" to table: "profiles"
CREATE INDEX "idx_profiles_avatar_id" ON "auth"."profiles" ("avatar_id");
-- Create index "idx_profiles_background_id" to table: "profiles"
CREATE INDEX "idx_profiles_background_id" ON "auth"."profiles" ("background_id");
-- Create index "idx_profiles_birthday" to table: "profiles"
CREATE INDEX "idx_profiles_birthday" ON "auth"."profiles" ("birthday");
-- Create index "idx_profiles_deleted_at" to table: "profiles"
CREATE INDEX "idx_profiles_deleted_at" ON "auth"."profiles" ("deleted_at");
-- Create index "idx_profiles_display_name" to table: "profiles"
CREATE INDEX "idx_profiles_display_name" ON "auth"."profiles" ("display_name");
-- Create index "idx_profiles_first_and_last_name" to table: "profiles"
CREATE INDEX "idx_profiles_first_and_last_name" ON "auth"."profiles" ("first_name", "last_name");
-- Create index "idx_profiles_gender" to table: "profiles"
CREATE INDEX "idx_profiles_gender" ON "auth"."profiles" ("gender");
-- Create index "idx_profiles_github" to table: "profiles"
CREATE INDEX "idx_profiles_github" ON "auth"."profiles" ("github");
-- Create index "idx_profiles_location" to table: "profiles"
CREATE INDEX "idx_profiles_location" ON "auth"."profiles" ("location");
-- Create index "idx_profiles_nickname" to table: "profiles"
CREATE INDEX "idx_profiles_nickname" ON "auth"."profiles" ("nickname");
-- Create index "idx_profiles_phone" to table: "profiles"
CREATE INDEX "idx_profiles_phone" ON "auth"."profiles" ("phone");
-- Create index "idx_profiles_user_id" to table: "profiles"
CREATE INDEX "idx_profiles_user_id" ON "auth"."profiles" ("user_id");
-- Create index "idx_profiles_website" to table: "profiles"
CREATE INDEX "idx_profiles_website" ON "auth"."profiles" ("website");
-- Create "educations" table
CREATE TABLE "auth"."educations" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "profile_id" uuid NOT NULL,
  "school" character varying(255) NOT NULL,
  "degree" character varying(100) NULL,
  "major" character varying(255) NULL,
  "field_of_study" character varying(255) NULL,
  "start_date" date NULL,
  "end_date" date NULL,
  "is_current" boolean NOT NULL DEFAULT false,
  "description" text NULL,
  "grade" character varying(50) NULL,
  "activities" text NULL,
  "achievements" text NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "educations_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "auth"."profiles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_educations_degree" to table: "educations"
CREATE INDEX "idx_educations_degree" ON "auth"."educations" ("degree");
-- Create index "idx_educations_deleted_at" to table: "educations"
CREATE INDEX "idx_educations_deleted_at" ON "auth"."educations" ("deleted_at");
-- Create index "idx_educations_profile_id" to table: "educations"
CREATE INDEX "idx_educations_profile_id" ON "auth"."educations" ("profile_id");
-- Create index "idx_educations_school" to table: "educations"
CREATE INDEX "idx_educations_school" ON "auth"."educations" ("school");
-- Create "image_types" table
CREATE TABLE "image"."image_types" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "key" text NOT NULL,
  "description" text NULL,
  "max_width" integer NULL,
  "max_height" integer NULL,
  "aspect_ratio" text NULL,
  "is_system" boolean NULL DEFAULT false,
  "created_at" timestamp NULL DEFAULT now(),
  "updated_at" timestamp NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "image_types_key_key" UNIQUE ("key")
);
-- Create index "idx_image_types_is_system" to table: "image_types"
CREATE INDEX "idx_image_types_is_system" ON "image"."image_types" ("is_system");
-- Create index "idx_image_types_key" to table: "image_types"
CREATE INDEX "idx_image_types_key" ON "image"."image_types" ("key");
-- Set comment to table: "image_types"
COMMENT ON TABLE "image"."image_types" IS 'Image type definitions for different use cases and business domains';
-- Set comment to column: "key" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."key" IS 'Type key identifier: avatar, background, product_cover, product_gallery, post_image, banner, badge_icon, etc.';
-- Set comment to column: "description" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."description" IS 'Human-readable description of what this image type is used for';
-- Set comment to column: "max_width" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."max_width" IS 'Maximum allowed width in pixels (NULL means no limit)';
-- Set comment to column: "max_height" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."max_height" IS 'Maximum allowed height in pixels (NULL means no limit)';
-- Set comment to column: "aspect_ratio" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."aspect_ratio" IS 'Preferred aspect ratio hint: 1:1, 16:9, 4:3, etc. Used for frontend auto-cropping';
-- Set comment to column: "is_system" on table: "image_types"
COMMENT ON COLUMN "image"."image_types"."is_system" IS 'TRUE for system-built types (cannot be deleted), FALSE for custom types';
-- Create "images" table
CREATE TABLE "image"."images" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "type_id" uuid NULL,
  "user_id" uuid NULL,
  "source_domain" text NULL,
  "filename" text NOT NULL,
  "original_filename" text NOT NULL,
  "mime_type" text NOT NULL,
  "size" bigint NOT NULL,
  "width" integer NULL,
  "height" integer NULL,
  "storage_path" text NOT NULL,
  "url" text NULL,
  "is_public" boolean NOT NULL DEFAULT false,
  "metadata" jsonb NULL DEFAULT '{}',
  "created_at" timestamp NULL DEFAULT now(),
  "updated_at" timestamp NULL DEFAULT now(),
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "images_type_id_fkey" FOREIGN KEY ("type_id") REFERENCES "image"."image_types" ("id") ON UPDATE NO ACTION ON DELETE SET NULL
);
-- Create index "idx_images_created_at" to table: "images"
CREATE INDEX "idx_images_created_at" ON "image"."images" ("created_at");
-- Create index "idx_images_deleted_at" to table: "images"
CREATE INDEX "idx_images_deleted_at" ON "image"."images" ("deleted_at");
-- Create index "idx_images_filename" to table: "images"
CREATE INDEX "idx_images_filename" ON "image"."images" ("filename");
-- Create index "idx_images_is_public" to table: "images"
CREATE INDEX "idx_images_is_public" ON "image"."images" ("is_public");
-- Create index "idx_images_mime_type" to table: "images"
CREATE INDEX "idx_images_mime_type" ON "image"."images" ("mime_type");
-- Create index "idx_images_source_domain" to table: "images"
CREATE INDEX "idx_images_source_domain" ON "image"."images" ("source_domain");
-- Create index "idx_images_type_id" to table: "images"
CREATE INDEX "idx_images_type_id" ON "image"."images" ("type_id");
-- Create index "idx_images_type_public" to table: "images"
CREATE INDEX "idx_images_type_public" ON "image"."images" ("type_id", "is_public") WHERE (deleted_at IS NULL);
-- Create index "idx_images_user_id" to table: "images"
CREATE INDEX "idx_images_user_id" ON "image"."images" ("user_id");
-- Create index "idx_images_user_public" to table: "images"
CREATE INDEX "idx_images_user_public" ON "image"."images" ("user_id", "is_public") WHERE (deleted_at IS NULL);
-- Set comment to table: "images"
COMMENT ON TABLE "image"."images" IS 'Main table storing original image metadata for all uploaded images';
-- Set comment to column: "type_id" on table: "images"
COMMENT ON COLUMN "image"."images"."type_id" IS 'Reference to image_types table defining the image category/usage';
-- Set comment to column: "user_id" on table: "images"
COMMENT ON COLUMN "image"."images"."user_id" IS 'UUID of user who uploaded the image (from auth service, no FK constraint)';
-- Set comment to column: "source_domain" on table: "images"
COMMENT ON COLUMN "image"."images"."source_domain" IS 'Source service identifier: auth, product, post, cms, etc.';
-- Set comment to column: "filename" on table: "images"
COMMENT ON COLUMN "image"."images"."filename" IS 'Current filename after processing (may differ from original_filename)';
-- Set comment to column: "original_filename" on table: "images"
COMMENT ON COLUMN "image"."images"."original_filename" IS 'Original filename as uploaded by user';
-- Set comment to column: "mime_type" on table: "images"
COMMENT ON COLUMN "image"."images"."mime_type" IS 'MIME type: image/jpeg, image/png, image/webp, image/gif, etc.';
-- Set comment to column: "size" on table: "images"
COMMENT ON COLUMN "image"."images"."size" IS 'File size in bytes';
-- Set comment to column: "width" on table: "images"
COMMENT ON COLUMN "image"."images"."width" IS 'Image width in pixels';
-- Set comment to column: "height" on table: "images"
COMMENT ON COLUMN "image"."images"."height" IS 'Image height in pixels';
-- Set comment to column: "storage_path" on table: "images"
COMMENT ON COLUMN "image"."images"."storage_path" IS 'Storage backend path (filesystem path or S3 object key)';
-- Set comment to column: "url" on table: "images"
COMMENT ON COLUMN "image"."images"."url" IS 'Public accessible URL (CDN URL or direct storage URL)';
-- Set comment to column: "is_public" on table: "images"
COMMENT ON COLUMN "image"."images"."is_public" IS 'TRUE = publicly accessible, FALSE = requires authentication';
-- Set comment to column: "metadata" on table: "images"
COMMENT ON COLUMN "image"."images"."metadata" IS 'Extended JSON metadata: EXIF, color profiles, AI labels, etc.';
-- Set comment to column: "deleted_at" on table: "images"
COMMENT ON COLUMN "image"."images"."deleted_at" IS 'Soft delete timestamp (NULL = active, NOT NULL = deleted)';
-- Create "image_tags" table
CREATE TABLE "image"."image_tags" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "image_id" uuid NOT NULL,
  "tag" text NOT NULL,
  "confidence" double precision NULL,
  "created_at" timestamp NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "image_tags_image_id_tag_key" UNIQUE ("image_id", "tag"),
  CONSTRAINT "image_tags_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "image"."images" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_image_tags_confidence" to table: "image_tags"
CREATE INDEX "idx_image_tags_confidence" ON "image"."image_tags" ("confidence") WHERE (confidence IS NOT NULL);
-- Create index "idx_image_tags_image_id" to table: "image_tags"
CREATE INDEX "idx_image_tags_image_id" ON "image"."image_tags" ("image_id");
-- Create index "idx_image_tags_tag" to table: "image_tags"
CREATE INDEX "idx_image_tags_tag" ON "image"."image_tags" ("tag");
-- Create index "idx_image_tags_tag_confidence" to table: "image_tags"
CREATE INDEX "idx_image_tags_tag_confidence" ON "image"."image_tags" ("tag", "confidence" DESC);
-- Set comment to table: "image_tags"
COMMENT ON TABLE "image"."image_tags" IS 'Image tags for content-based search: AI labels and user-defined tags';
-- Set comment to column: "image_id" on table: "image_tags"
COMMENT ON COLUMN "image"."image_tags"."image_id" IS 'Reference to image in images table';
-- Set comment to column: "tag" on table: "image_tags"
COMMENT ON COLUMN "image"."image_tags"."tag" IS 'Tag text (normalized, e.g., lowercase). Examples: cat, shoes, selfie, nature';
-- Set comment to column: "confidence" on table: "image_tags"
COMMENT ON COLUMN "image"."image_tags"."confidence" IS 'Confidence score 0.0-1.0 for AI-generated tags, NULL for user tags';
-- Set comment to column: "created_at" on table: "image_tags"
COMMENT ON COLUMN "image"."image_tags"."created_at" IS 'When this tag was added to the image';
-- Create "image_variants" table
CREATE TABLE "image"."image_variants" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "image_id" uuid NOT NULL,
  "variant_key" text NOT NULL,
  "width" integer NULL,
  "height" integer NULL,
  "size" bigint NULL,
  "mime_type" text NULL,
  "storage_path" text NOT NULL,
  "url" text NULL,
  "created_at" timestamp NULL DEFAULT now(),
  "updated_at" timestamp NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "image_variants_image_id_variant_key_key" UNIQUE ("image_id", "variant_key"),
  CONSTRAINT "image_variants_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "image"."images" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_image_variants_image_id" to table: "image_variants"
CREATE INDEX "idx_image_variants_image_id" ON "image"."image_variants" ("image_id");
-- Create index "idx_image_variants_image_key" to table: "image_variants"
CREATE INDEX "idx_image_variants_image_key" ON "image"."image_variants" ("image_id", "variant_key");
-- Create index "idx_image_variants_url" to table: "image_variants"
CREATE INDEX "idx_image_variants_url" ON "image"."image_variants" ("url") WHERE (url IS NOT NULL);
-- Create index "idx_image_variants_variant_key" to table: "image_variants"
CREATE INDEX "idx_image_variants_variant_key" ON "image"."image_variants" ("variant_key");
-- Set comment to table: "image_variants"
COMMENT ON TABLE "image"."image_variants" IS 'Stores derived versions of images: thumbnails, resized versions, format conversions';
-- Set comment to column: "image_id" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."image_id" IS 'Reference to parent image in images table';
-- Set comment to column: "variant_key" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."variant_key" IS 'Variant type: thumbnail (150x150), small (300px), medium (720px), large (1080px), webp (converted), etc.';
-- Set comment to column: "width" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."width" IS 'Variant width in pixels';
-- Set comment to column: "height" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."height" IS 'Variant height in pixels';
-- Set comment to column: "size" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."size" IS 'Variant file size in bytes';
-- Set comment to column: "mime_type" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."mime_type" IS 'Variant MIME type (may differ from original, e.g., webp conversion)';
-- Set comment to column: "storage_path" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."storage_path" IS 'Storage backend path for this variant';
-- Set comment to column: "url" on table: "image_variants"
COMMENT ON COLUMN "image"."image_variants"."url" IS 'Public accessible URL for this variant';
-- Create "occupations" table
CREATE TABLE "auth"."occupations" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "profile_id" uuid NOT NULL,
  "company" character varying(255) NOT NULL,
  "position" character varying(255) NOT NULL,
  "department" character varying(255) NULL,
  "industry" character varying(100) NULL,
  "location" character varying(255) NULL,
  "employment_type" character varying(50) NULL,
  "start_date" date NULL,
  "end_date" date NULL,
  "is_current" boolean NOT NULL DEFAULT false,
  "description" text NULL,
  "responsibilities" text NULL,
  "achievements" text NULL,
  "skills_used" text[] NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "occupations_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "auth"."profiles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_occupations_company" to table: "occupations"
CREATE INDEX "idx_occupations_company" ON "auth"."occupations" ("company");
-- Create index "idx_occupations_deleted_at" to table: "occupations"
CREATE INDEX "idx_occupations_deleted_at" ON "auth"."occupations" ("deleted_at");
-- Create index "idx_occupations_industry" to table: "occupations"
CREATE INDEX "idx_occupations_industry" ON "auth"."occupations" ("industry");
-- Create index "idx_occupations_is_current" to table: "occupations"
CREATE INDEX "idx_occupations_is_current" ON "auth"."occupations" ("is_current");
-- Create index "idx_occupations_position" to table: "occupations"
CREATE INDEX "idx_occupations_position" ON "auth"."occupations" ("position");
-- Create index "idx_occupations_profile_id" to table: "occupations"
CREATE INDEX "idx_occupations_profile_id" ON "auth"."occupations" ("profile_id");
-- Create "badges" table
CREATE TABLE "auth"."badges" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying(50) NOT NULL,
  "description" text NULL,
  "icon_url" character varying(255) NULL,
  "color" character varying(20) NULL,
  "category" character varying(50) NULL,
  "is_system" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "badges_name_key" UNIQUE ("name")
);
-- Create index "idx_badges_category" to table: "badges"
CREATE INDEX "idx_badges_category" ON "auth"."badges" ("category");
-- Create index "idx_badges_deleted_at" to table: "badges"
CREATE INDEX "idx_badges_deleted_at" ON "auth"."badges" ("deleted_at");
-- Create index "idx_badges_name" to table: "badges"
CREATE INDEX "idx_badges_name" ON "auth"."badges" ("name");
-- Create "profile_badges" table
CREATE TABLE "auth"."profile_badges" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "profile_id" uuid NOT NULL,
  "badge_id" uuid NOT NULL,
  "description" text NULL DEFAULT '',
  "level" integer NULL DEFAULT 1,
  "earned_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "profile_badges_profile_id_badge_id_key" UNIQUE ("profile_id", "badge_id"),
  CONSTRAINT "profile_badges_badge_id_fkey" FOREIGN KEY ("badge_id") REFERENCES "auth"."badges" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "profile_badges_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "auth"."profiles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_profile_badges_badge_id" to table: "profile_badges"
CREATE INDEX "idx_profile_badges_badge_id" ON "auth"."profile_badges" ("badge_id");
-- Create index "idx_profile_badges_earned_at" to table: "profile_badges"
CREATE INDEX "idx_profile_badges_earned_at" ON "auth"."profile_badges" ("earned_at");
-- Create index "idx_profile_badges_profile_id" to table: "profile_badges"
CREATE INDEX "idx_profile_badges_profile_id" ON "auth"."profile_badges" ("profile_id");
