-- Add new schema named "asset"
CREATE SCHEMA "asset";
-- Set comment to schema: "asset"
COMMENT ON SCHEMA "asset" IS 'Independent Asset service schema: Images, Files, Videos, Audios metadata. uploader_id is logical ref to Auth (no cross-schema FK). Stored in Stack MinIO.';
-- Add new schema named "auth"
CREATE SCHEMA "auth";
-- Set comment to schema: "auth"
COMMENT ON SCHEMA "auth" IS 'Login center: accounts, identities, forger/authority profiles. Avatars/backgrounds link to asset.Images.';
-- Create extension "pgcrypto"
CREATE EXTENSION "pgcrypto" WITH SCHEMA "public" VERSION "1.3";
-- Create extension "btree_gist"
CREATE EXTENSION "btree_gist" WITH SCHEMA "public" VERSION "1.8";
-- Create "Audios" table
CREATE TABLE "asset"."Audios" (
  "id" uuid NOT NULL,
  "file_path" character varying(500) NOT NULL,
  "file_name" character varying(255) NOT NULL,
  "file_size" bigint NOT NULL,
  "mime_type" character varying(100) NOT NULL,
  "duration_seconds" double precision NULL,
  "uploader_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_audios_deleted_at" to table: "Audios"
CREATE INDEX "idx_audios_deleted_at" ON "asset"."Audios" ("deleted_at");
-- Create index "idx_audios_file_path" to table: "Audios"
CREATE INDEX "idx_audios_file_path" ON "asset"."Audios" ("file_path");
-- Create index "idx_audios_mime_type" to table: "Audios"
CREATE INDEX "idx_audios_mime_type" ON "asset"."Audios" ("mime_type");
-- Create index "idx_audios_uploader_id" to table: "Audios"
CREATE INDEX "idx_audios_uploader_id" ON "asset"."Audios" ("uploader_id");
-- Set comment to table: "Audios"
COMMENT ON TABLE "asset"."Audios" IS 'Audio binary metadata; duration when known; uploader_id references auth.Accounts at app level.';
-- Set comment to column: "id" on table: "Audios"
COMMENT ON COLUMN "asset"."Audios"."id" IS 'Primary key UUID; typically matches stored object id.';
-- Set comment to column: "file_path" on table: "Audios"
COMMENT ON COLUMN "asset"."Audios"."file_path" IS 'Relative path from storage root or object key.';
-- Set comment to column: "file_name" on table: "Audios"
COMMENT ON COLUMN "asset"."Audios"."file_name" IS 'Original filename.';
-- Set comment to column: "file_size" on table: "Audios"
COMMENT ON COLUMN "asset"."Audios"."file_size" IS 'File size in bytes.';
-- Set comment to column: "mime_type" on table: "Audios"
COMMENT ON COLUMN "asset"."Audios"."mime_type" IS 'MIME type as uploaded or detected.';
-- Set comment to column: "duration_seconds" on table: "Audios"
COMMENT ON COLUMN "asset"."Audios"."duration_seconds" IS 'Media duration when known.';
-- Set comment to column: "uploader_id" on table: "Audios"
COMMENT ON COLUMN "asset"."Audios"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
-- Set comment to column: "created_at" on table: "Audios"
COMMENT ON COLUMN "asset"."Audios"."created_at" IS 'Insert time.';
-- Set comment to column: "updated_at" on table: "Audios"
COMMENT ON COLUMN "asset"."Audios"."updated_at" IS 'Last metadata update.';
-- Set comment to column: "deleted_at" on table: "Audios"
COMMENT ON COLUMN "asset"."Audios"."deleted_at" IS 'Soft-delete; NULL if object still active.';
-- Create "Files" table
CREATE TABLE "asset"."Files" (
  "id" uuid NOT NULL,
  "file_path" character varying(500) NOT NULL,
  "file_name" character varying(255) NOT NULL,
  "file_size" bigint NOT NULL,
  "mime_type" character varying(100) NOT NULL,
  "uploader_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_files_deleted_at" to table: "Files"
CREATE INDEX "idx_files_deleted_at" ON "asset"."Files" ("deleted_at");
-- Create index "idx_files_file_path" to table: "Files"
CREATE INDEX "idx_files_file_path" ON "asset"."Files" ("file_path");
-- Create index "idx_files_mime_type" to table: "Files"
CREATE INDEX "idx_files_mime_type" ON "asset"."Files" ("mime_type");
-- Create index "idx_files_uploader_id" to table: "Files"
CREATE INDEX "idx_files_uploader_id" ON "asset"."Files" ("uploader_id");
-- Set comment to table: "Files"
COMMENT ON TABLE "asset"."Files" IS 'Non-image binary metadata (documents, archives, etc.); same soft-delete shape as Images without pixel fields.';
-- Set comment to column: "id" on table: "Files"
COMMENT ON COLUMN "asset"."Files"."id" IS 'Primary key UUID; typically matches stored object id.';
-- Set comment to column: "file_path" on table: "Files"
COMMENT ON COLUMN "asset"."Files"."file_path" IS 'Relative path from storage root or object key.';
-- Set comment to column: "file_name" on table: "Files"
COMMENT ON COLUMN "asset"."Files"."file_name" IS 'Original filename.';
-- Set comment to column: "file_size" on table: "Files"
COMMENT ON COLUMN "asset"."Files"."file_size" IS 'File size in bytes.';
-- Set comment to column: "mime_type" on table: "Files"
COMMENT ON COLUMN "asset"."Files"."mime_type" IS 'MIME type as uploaded or detected.';
-- Set comment to column: "uploader_id" on table: "Files"
COMMENT ON COLUMN "asset"."Files"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
-- Set comment to column: "created_at" on table: "Files"
COMMENT ON COLUMN "asset"."Files"."created_at" IS 'Insert time.';
-- Set comment to column: "updated_at" on table: "Files"
COMMENT ON COLUMN "asset"."Files"."updated_at" IS 'Last metadata update.';
-- Set comment to column: "deleted_at" on table: "Files"
COMMENT ON COLUMN "asset"."Files"."deleted_at" IS 'Soft-delete; NULL if object still active.';
-- Create "Images" table
CREATE TABLE "asset"."Images" (
  "id" uuid NOT NULL,
  "file_path" character varying(500) NOT NULL,
  "file_name" character varying(255) NOT NULL,
  "file_size" bigint NOT NULL,
  "mime_type" character varying(100) NOT NULL,
  "width" integer NULL,
  "height" integer NULL,
  "alt_text" character varying(255) NULL,
  "uploader_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_images_deleted_at" to table: "Images"
CREATE INDEX "idx_images_deleted_at" ON "asset"."Images" ("deleted_at");
-- Create index "idx_images_file_path" to table: "Images"
CREATE INDEX "idx_images_file_path" ON "asset"."Images" ("file_path");
-- Create index "idx_images_mime_type" to table: "Images"
CREATE INDEX "idx_images_mime_type" ON "asset"."Images" ("mime_type");
-- Create index "idx_images_uploader_id" to table: "Images"
CREATE INDEX "idx_images_uploader_id" ON "asset"."Images" ("uploader_id");
-- Set comment to table: "Images"
COMMENT ON TABLE "asset"."Images" IS 'Raster image metadata; path is relative to storage root or object key; uploader_id references auth.Accounts at app level.';
-- Set comment to column: "id" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."id" IS 'Primary key UUID; typically matches stored object id.';
-- Set comment to column: "file_path" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."file_path" IS 'Relative path from storage root or object key.';
-- Set comment to column: "file_name" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."file_name" IS 'Original filename.';
-- Set comment to column: "file_size" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."file_size" IS 'File size in bytes.';
-- Set comment to column: "mime_type" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."mime_type" IS 'MIME type as uploaded or detected.';
-- Set comment to column: "width" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."width" IS 'Pixel width if known.';
-- Set comment to column: "height" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."height" IS 'Pixel height if known.';
-- Set comment to column: "alt_text" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."alt_text" IS 'Accessibility / caption text.';
-- Set comment to column: "uploader_id" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
-- Set comment to column: "created_at" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."created_at" IS 'Insert time.';
-- Set comment to column: "updated_at" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."updated_at" IS 'Last metadata update.';
-- Set comment to column: "deleted_at" on table: "Images"
COMMENT ON COLUMN "asset"."Images"."deleted_at" IS 'Soft-delete; NULL if object still active.';
-- Create "Videos" table
CREATE TABLE "asset"."Videos" (
  "id" uuid NOT NULL,
  "file_path" character varying(500) NOT NULL,
  "file_name" character varying(255) NOT NULL,
  "file_size" bigint NOT NULL,
  "mime_type" character varying(100) NOT NULL,
  "duration_seconds" double precision NULL,
  "width" integer NULL,
  "height" integer NULL,
  "uploader_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_videos_deleted_at" to table: "Videos"
CREATE INDEX "idx_videos_deleted_at" ON "asset"."Videos" ("deleted_at");
-- Create index "idx_videos_file_path" to table: "Videos"
CREATE INDEX "idx_videos_file_path" ON "asset"."Videos" ("file_path");
-- Create index "idx_videos_mime_type" to table: "Videos"
CREATE INDEX "idx_videos_mime_type" ON "asset"."Videos" ("mime_type");
-- Create index "idx_videos_uploader_id" to table: "Videos"
CREATE INDEX "idx_videos_uploader_id" ON "asset"."Videos" ("uploader_id");
-- Set comment to table: "Videos"
COMMENT ON TABLE "asset"."Videos" IS 'Video binary metadata; duration and optional display dimensions; uploader_id references auth.Accounts at app level.';
-- Set comment to column: "id" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."id" IS 'Primary key UUID; typically matches stored object id.';
-- Set comment to column: "file_path" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."file_path" IS 'Relative path from storage root or object key.';
-- Set comment to column: "file_name" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."file_name" IS 'Original filename.';
-- Set comment to column: "file_size" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."file_size" IS 'File size in bytes.';
-- Set comment to column: "mime_type" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."mime_type" IS 'MIME type as uploaded or detected.';
-- Set comment to column: "duration_seconds" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."duration_seconds" IS 'Media duration when known.';
-- Set comment to column: "width" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."width" IS 'Pixel width if known.';
-- Set comment to column: "height" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."height" IS 'Pixel height if known.';
-- Set comment to column: "uploader_id" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
-- Set comment to column: "created_at" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."created_at" IS 'Insert time.';
-- Set comment to column: "updated_at" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."updated_at" IS 'Last metadata update.';
-- Set comment to column: "deleted_at" on table: "Videos"
COMMENT ON COLUMN "asset"."Videos"."deleted_at" IS 'Soft-delete; NULL if object still active.';
-- Create enum type "identity_provider"
CREATE TYPE "auth"."identity_provider" AS ENUM ('password');
-- Create enum type "forger_role"
CREATE TYPE "auth"."forger_role" AS ENUM ('forger');
-- Create enum type "authority_role"
CREATE TYPE "auth"."authority_role" AS ENUM ('auditor', 'administrator', 'owner');
-- Create enum type "account_status"
CREATE TYPE "auth"."account_status" AS ENUM ('active', 'suspended', 'deleted');
-- Create enum type "signup_platform"
CREATE TYPE "auth"."signup_platform" AS ENUM ('nfxidentity', 'nfxnews', 'nfxstorages', 'nfxvault');
-- Create enum type "profile_language"
CREATE TYPE "auth"."profile_language" AS ENUM ('en', 'zh', 'fr');
-- Create enum type "profile_scope"
CREATE TYPE "auth"."profile_scope" AS ENUM ('community', 'authority');
-- Create "Accounts" table
CREATE TABLE "auth"."Accounts" (
  "id" uuid NOT NULL,
  "account_status" "auth"."account_status" NOT NULL DEFAULT 'active',
  "signup_platform" "auth"."signup_platform" NOT NULL DEFAULT 'nfxidentity',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_accounts_account_status" to table: "Accounts"
CREATE INDEX "idx_accounts_account_status" ON "auth"."Accounts" ("account_status");
-- Create index "idx_accounts_deleted_at" to table: "Accounts"
CREATE INDEX "idx_accounts_deleted_at" ON "auth"."Accounts" ("deleted_at");
-- Create index "idx_accounts_signup_platform" to table: "Accounts"
CREATE INDEX "idx_accounts_signup_platform" ON "auth"."Accounts" ("signup_platform");
-- Set comment to table: "Accounts"
COMMENT ON TABLE "auth"."Accounts" IS 'Core account; credentials in Identities, forger profile in ForgerProfiles, authority profile in AuthorityProfiles, emails in Emails.';
-- Set comment to column: "id" on table: "Accounts"
COMMENT ON COLUMN "auth"."Accounts"."id" IS 'Primary key UUID; typically matches application user id.';
-- Set comment to column: "account_status" on table: "Accounts"
COMMENT ON COLUMN "auth"."Accounts"."account_status" IS 'Lifecycle: auth.account_status enum.';
-- Set comment to column: "signup_platform" on table: "Accounts"
COMMENT ON COLUMN "auth"."Accounts"."signup_platform" IS 'Product that first created this account (auth.signup_platform); written only at signup.';
-- Set comment to column: "created_at" on table: "Accounts"
COMMENT ON COLUMN "auth"."Accounts"."created_at" IS 'Row insert time.';
-- Set comment to column: "updated_at" on table: "Accounts"
COMMENT ON COLUMN "auth"."Accounts"."updated_at" IS 'Last account row update.';
-- Set comment to column: "deleted_at" on table: "Accounts"
COMMENT ON COLUMN "auth"."Accounts"."deleted_at" IS 'Optional soft-delete timestamp when account_status = deleted.';
-- Add new schema named "system"
CREATE SCHEMA "system";
-- Set comment to schema: "system"
COMMENT ON SCHEMA "system" IS 'System-level state and administration schema';
-- Create "system_state" table
CREATE TABLE "system"."system_state" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "initialized" boolean NOT NULL DEFAULT false,
  "initialized_at" timestamp NULL,
  "initialization_version" character varying(50) NULL,
  "last_reset_at" timestamp NULL,
  "last_reset_by" uuid NULL,
  "reset_count" integer NOT NULL DEFAULT 0,
  "metadata" jsonb NULL DEFAULT '{}',
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);
-- Create index "idx_system_state_created_at" to table: "system_state"
CREATE INDEX "idx_system_state_created_at" ON "system"."system_state" ("created_at" DESC);
-- Create index "idx_system_state_initialized" to table: "system_state"
CREATE INDEX "idx_system_state_initialized" ON "system"."system_state" ("initialized");
-- Set comment to table: "system_state"
COMMENT ON TABLE "system"."system_state" IS 'System initialization state: tracks if system has been bootstrapped. Query logic: SELECT initialized FROM system_state ORDER BY created_at DESC LIMIT 1. If no record exists OR initialized = false, system is not initialized.';
-- Set comment to column: "id" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."id" IS 'UUID primary key (not fixed, allows multiple records for reset/re-initialization)';
-- Set comment to column: "initialized" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."initialized" IS 'Whether system has been initialized (checked on service startup). Always check latest record by created_at DESC.';
-- Set comment to column: "initialized_at" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."initialized_at" IS 'Timestamp when system was initialized via /bootstrap/initialize';
-- Set comment to column: "initialization_version" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."initialization_version" IS 'Version of initialization schema/data for migration tracking';
-- Set comment to column: "last_reset_at" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."last_reset_at" IS 'Timestamp when system was last reset';
-- Set comment to column: "last_reset_by" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."last_reset_by" IS 'User ID who reset the system. Even if database is cleared after reset, can be traced via log files for accountability';
-- Set comment to column: "created_at" on table: "system_state"
COMMENT ON COLUMN "system"."system_state"."created_at" IS 'Record creation time. Used to determine latest state (ORDER BY created_at DESC LIMIT 1)';
-- Create "AuthorityProfiles" table
CREATE TABLE "auth"."AuthorityProfiles" (
  "id" uuid NOT NULL,
  "account_id" uuid NOT NULL,
  "authority_roles" "auth"."authority_role"[] NOT NULL DEFAULT ARRAY['auditor'::auth.authority_role],
  "profile_language" "auth"."profile_language" NOT NULL DEFAULT 'zh',
  "preference" jsonb NULL,
  "display_name" character varying(150) NULL,
  "first_name" character varying(100) NULL,
  "last_name" character varying(100) NULL,
  "country" character varying(100) NULL,
  "city" character varying(100) NULL,
  "gender" character varying(100) NULL,
  "birthday" date NULL,
  "website" character varying(100) NULL,
  "timezone" character varying(100) NULL,
  "bio" text NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_authority_profiles_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "chk_authority_profiles_authority_roles_nonempty" CHECK (cardinality(authority_roles) >= 1)
);
-- Create index "idx_authority_profiles_account_id" to table: "AuthorityProfiles"
CREATE INDEX "idx_authority_profiles_account_id" ON "auth"."AuthorityProfiles" ("account_id");
-- Create index "idx_authority_profiles_authority_roles" to table: "AuthorityProfiles"
CREATE INDEX "idx_authority_profiles_authority_roles" ON "auth"."AuthorityProfiles" USING GIN ("authority_roles");
-- Create index "idx_authority_profiles_deleted_at" to table: "AuthorityProfiles"
CREATE INDEX "idx_authority_profiles_deleted_at" ON "auth"."AuthorityProfiles" ("deleted_at");
-- Create index "idx_authority_profiles_profile_language" to table: "AuthorityProfiles"
CREATE INDEX "idx_authority_profiles_profile_language" ON "auth"."AuthorityProfiles" ("profile_language");
-- Set comment to table: "AuthorityProfiles"
COMMENT ON TABLE "auth"."AuthorityProfiles" IS 'Staff/authority display profile; one account may own many profiles (1:N). Distinct from ForgerProfiles; carries authority_roles[]. Phone numbers live in auth.Phones.';
-- Set comment to column: "id" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."id" IS 'Primary key UUID (profile_id); referenced by content/social as profile_id.';
-- Set comment to column: "account_id" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."account_id" IS 'Owning account; FK to auth.Accounts; one account may have many profiles; CASCADE delete.';
-- Set comment to column: "authority_roles" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."authority_roles" IS 'auth.authority_role[]; capability = membership only (reviewer, moderator, …). Default {reviewer}.';
-- Set comment to column: "profile_language" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."profile_language" IS 'Preferred UI locale (auth.profile_language).';
-- Set comment to column: "preference" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."preference" IS 'User preferences JSON: theme, layout, etc.';
-- Set comment to column: "display_name" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."display_name" IS 'Public handle / nickname; optional if first+last used.';
-- Set comment to column: "first_name" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."first_name" IS 'Given name; optional.';
-- Set comment to column: "last_name" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."last_name" IS 'Family name; optional.';
-- Set comment to column: "country" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."country" IS 'Country string for display or filters.';
-- Set comment to column: "city" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."city" IS 'City string.';
-- Set comment to column: "gender" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."gender" IS 'Self-identified gender label; app-defined vocabulary.';
-- Set comment to column: "birthday" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."birthday" IS 'Birth date if collected.';
-- Set comment to column: "website" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."website" IS 'Personal or social URL.';
-- Set comment to column: "timezone" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."timezone" IS 'IANA timezone name if set.';
-- Set comment to column: "bio" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."bio" IS 'Free-text bio.';
-- Set comment to column: "created_at" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."created_at" IS 'Row insert time.';
-- Set comment to column: "updated_at" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."updated_at" IS 'Last profile edit.';
-- Set comment to column: "deleted_at" on table: "AuthorityProfiles"
COMMENT ON COLUMN "auth"."AuthorityProfiles"."deleted_at" IS 'Soft-delete profile snapshot.';
-- Create "AuthorityProfileAvatars" table
CREATE TABLE "auth"."AuthorityProfileAvatars" (
  "id" uuid NOT NULL,
  "profile_id" uuid NOT NULL,
  "image_id" uuid NOT NULL,
  "is_active" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_authority_profile_avatars_profile_id" FOREIGN KEY ("profile_id") REFERENCES "auth"."AuthorityProfiles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_authority_profile_avatars_deleted_at" to table: "AuthorityProfileAvatars"
CREATE INDEX "idx_authority_profile_avatars_deleted_at" ON "auth"."AuthorityProfileAvatars" ("deleted_at");
-- Create index "idx_authority_profile_avatars_image_id" to table: "AuthorityProfileAvatars"
CREATE INDEX "idx_authority_profile_avatars_image_id" ON "auth"."AuthorityProfileAvatars" ("image_id");
-- Create index "idx_authority_profile_avatars_profile_id" to table: "AuthorityProfileAvatars"
CREATE INDEX "idx_authority_profile_avatars_profile_id" ON "auth"."AuthorityProfileAvatars" ("profile_id");
-- Create index "uq_authority_profile_avatars_one_active_per_profile" to table: "AuthorityProfileAvatars"
CREATE UNIQUE INDEX "uq_authority_profile_avatars_one_active_per_profile" ON "auth"."AuthorityProfileAvatars" ("profile_id") WHERE ((is_active = true) AND (deleted_at IS NULL));
-- Set comment to table: "AuthorityProfileAvatars"
COMMENT ON TABLE "auth"."AuthorityProfileAvatars" IS 'Avatar history per authority profile; at most one active non-deleted row per AuthorityProfiles row; image_id → asset.Images at app level.';
-- Set comment to column: "id" on table: "AuthorityProfileAvatars"
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."id" IS 'Primary key UUID.';
-- Set comment to column: "profile_id" on table: "AuthorityProfileAvatars"
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."profile_id" IS 'FK auth.AuthorityProfiles.id; CASCADE delete.';
-- Set comment to column: "image_id" on table: "AuthorityProfileAvatars"
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."image_id" IS 'Reference to asset.Images.id (application-level, no FK).';
-- Set comment to column: "is_active" on table: "AuthorityProfileAvatars"
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."is_active" IS 'Exactly one non-deleted active row per profile_id; see partial unique index.';
-- Set comment to column: "created_at" on table: "AuthorityProfileAvatars"
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."created_at" IS 'When this avatar record was created.';
-- Set comment to column: "updated_at" on table: "AuthorityProfileAvatars"
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."updated_at" IS 'Last change to this row.';
-- Set comment to column: "deleted_at" on table: "AuthorityProfileAvatars"
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."deleted_at" IS 'Soft-delete; set when this history row is removed from use.';
-- Create "AuthorityProfileBackgrounds" table
CREATE TABLE "auth"."AuthorityProfileBackgrounds" (
  "id" uuid NOT NULL,
  "profile_id" uuid NOT NULL,
  "image_id" uuid NOT NULL,
  "sort_order" integer NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_authority_profile_backgrounds_profile_id" FOREIGN KEY ("profile_id") REFERENCES "auth"."AuthorityProfiles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_authority_profile_backgrounds_deleted_at" to table: "AuthorityProfileBackgrounds"
CREATE INDEX "idx_authority_profile_backgrounds_deleted_at" ON "auth"."AuthorityProfileBackgrounds" ("deleted_at");
-- Create index "idx_authority_profile_backgrounds_image_id" to table: "AuthorityProfileBackgrounds"
CREATE INDEX "idx_authority_profile_backgrounds_image_id" ON "auth"."AuthorityProfileBackgrounds" ("image_id");
-- Create index "idx_authority_profile_backgrounds_profile_id" to table: "AuthorityProfileBackgrounds"
CREATE INDEX "idx_authority_profile_backgrounds_profile_id" ON "auth"."AuthorityProfileBackgrounds" ("profile_id");
-- Create index "idx_authority_profile_backgrounds_profile_sort" to table: "AuthorityProfileBackgrounds"
CREATE INDEX "idx_authority_profile_backgrounds_profile_sort" ON "auth"."AuthorityProfileBackgrounds" ("profile_id", "sort_order");
-- Set comment to table: "AuthorityProfileBackgrounds"
COMMENT ON TABLE "auth"."AuthorityProfileBackgrounds" IS 'Background images per authority profile; order by sort_order; soft-delete via deleted_at; no cross-schema FK to asset.Images.';
-- Set comment to column: "id" on table: "AuthorityProfileBackgrounds"
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."id" IS 'Primary key UUID.';
-- Set comment to column: "profile_id" on table: "AuthorityProfileBackgrounds"
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."profile_id" IS 'FK auth.AuthorityProfiles.id; CASCADE delete.';
-- Set comment to column: "image_id" on table: "AuthorityProfileBackgrounds"
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."image_id" IS 'Reference to asset.Images.id (application-level, no FK).';
-- Set comment to column: "sort_order" on table: "AuthorityProfileBackgrounds"
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."sort_order" IS 'Display order within profile; lower = higher priority.';
-- Set comment to column: "created_at" on table: "AuthorityProfileBackgrounds"
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."created_at" IS 'When this background row was created.';
-- Set comment to column: "updated_at" on table: "AuthorityProfileBackgrounds"
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."updated_at" IS 'Last change to this row.';
-- Set comment to column: "deleted_at" on table: "AuthorityProfileBackgrounds"
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."deleted_at" IS 'Soft-delete row.';
-- Create "AuthorityProfileSettings" table
CREATE TABLE "auth"."AuthorityProfileSettings" (
  "id" uuid NOT NULL,
  "login_notification" boolean NOT NULL DEFAULT true,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_authority_profile_settings_id" FOREIGN KEY ("id") REFERENCES "auth"."AuthorityProfiles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_authority_profile_settings_deleted_at" to table: "AuthorityProfileSettings"
CREATE INDEX "idx_authority_profile_settings_deleted_at" ON "auth"."AuthorityProfileSettings" ("deleted_at");
-- Set comment to table: "AuthorityProfileSettings"
COMMENT ON TABLE "auth"."AuthorityProfileSettings" IS 'Per-authority-profile system settings; id equals AuthorityProfiles.id (1:1).';
-- Set comment to column: "id" on table: "AuthorityProfileSettings"
COMMENT ON COLUMN "auth"."AuthorityProfileSettings"."id" IS 'Primary key; same as auth.AuthorityProfiles.id.';
-- Set comment to column: "login_notification" on table: "AuthorityProfileSettings"
COMMENT ON COLUMN "auth"."AuthorityProfileSettings"."login_notification" IS 'When true, send email on successful login.';
-- Set comment to column: "created_at" on table: "AuthorityProfileSettings"
COMMENT ON COLUMN "auth"."AuthorityProfileSettings"."created_at" IS 'When this settings row was created.';
-- Set comment to column: "updated_at" on table: "AuthorityProfileSettings"
COMMENT ON COLUMN "auth"."AuthorityProfileSettings"."updated_at" IS 'Last change to this row.';
-- Set comment to column: "deleted_at" on table: "AuthorityProfileSettings"
COMMENT ON COLUMN "auth"."AuthorityProfileSettings"."deleted_at" IS 'Soft-delete timestamp.';
-- Create "Emails" table
CREATE TABLE "auth"."Emails" (
  "id" uuid NOT NULL,
  "account_id" uuid NOT NULL,
  "email" character varying(255) NOT NULL,
  "is_primary" boolean NOT NULL DEFAULT false,
  "verified_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_emails_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_emails_account_id" to table: "Emails"
CREATE INDEX "idx_emails_account_id" ON "auth"."Emails" ("account_id");
-- Create index "idx_emails_deleted_at" to table: "Emails"
CREATE INDEX "idx_emails_deleted_at" ON "auth"."Emails" ("deleted_at");
-- Create index "uq_emails_email_ci" to table: "Emails"
CREATE UNIQUE INDEX "uq_emails_email_ci" ON "auth"."Emails" ((lower((email)::text))) WHERE (deleted_at IS NULL);
-- Create index "uq_emails_one_primary_per_account" to table: "Emails"
CREATE UNIQUE INDEX "uq_emails_one_primary_per_account" ON "auth"."Emails" ("account_id") WHERE ((is_primary = true) AND (deleted_at IS NULL));
-- Set comment to table: "Emails"
COMMENT ON TABLE "auth"."Emails" IS 'Email bindings; soft-deleted rows kept for history. Partial unique (LOWER(email)) WHERE deleted_at IS NULL avoids collision when re-registering. Avoid hard-deleting auth.Accounts if you need email audit trail (CASCADE removes child rows).';
-- Set comment to column: "id" on table: "Emails"
COMMENT ON COLUMN "auth"."Emails"."id" IS 'Primary key UUID.';
-- Set comment to column: "account_id" on table: "Emails"
COMMENT ON COLUMN "auth"."Emails"."account_id" IS 'Owning auth.Accounts row; CASCADE delete.';
-- Set comment to column: "email" on table: "Emails"
COMMENT ON COLUMN "auth"."Emails"."email" IS 'Email address; uniqueness enforced case-insensitively among non-deleted rows.';
-- Set comment to column: "is_primary" on table: "Emails"
COMMENT ON COLUMN "auth"."Emails"."is_primary" IS 'At most one non-deleted primary email per account (partial unique index).';
-- Set comment to column: "verified_at" on table: "Emails"
COMMENT ON COLUMN "auth"."Emails"."verified_at" IS 'NULL = not verified yet.';
-- Set comment to column: "created_at" on table: "Emails"
COMMENT ON COLUMN "auth"."Emails"."created_at" IS 'Row insert time.';
-- Set comment to column: "updated_at" on table: "Emails"
COMMENT ON COLUMN "auth"."Emails"."updated_at" IS 'Last edit time.';
-- Set comment to column: "deleted_at" on table: "Emails"
COMMENT ON COLUMN "auth"."Emails"."deleted_at" IS 'Soft-delete only: row remains for audit. Partial unique applies only where deleted_at IS NULL. Re-register adds another row.';
-- Create "ForgerProfiles" table
CREATE TABLE "auth"."ForgerProfiles" (
  "id" uuid NOT NULL,
  "account_id" uuid NOT NULL,
  "forger_roles" "auth"."forger_role"[] NOT NULL DEFAULT ARRAY['forger'::auth.forger_role],
  "profile_language" "auth"."profile_language" NOT NULL DEFAULT 'en',
  "preference" jsonb NULL,
  "display_name" character varying(150) NULL,
  "first_name" character varying(100) NULL,
  "last_name" character varying(100) NULL,
  "country" character varying(100) NULL,
  "city" character varying(100) NULL,
  "gender" character varying(100) NULL,
  "birthday" date NULL,
  "website" character varying(100) NULL,
  "timezone" character varying(100) NULL,
  "bio" text NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_forger_profiles_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "chk_forger_profiles_forger_roles_nonempty" CHECK (cardinality(forger_roles) >= 1)
);
-- Create index "idx_forger_profiles_account_id" to table: "ForgerProfiles"
CREATE INDEX "idx_forger_profiles_account_id" ON "auth"."ForgerProfiles" ("account_id");
-- Create index "idx_forger_profiles_deleted_at" to table: "ForgerProfiles"
CREATE INDEX "idx_forger_profiles_deleted_at" ON "auth"."ForgerProfiles" ("deleted_at");
-- Create index "idx_forger_profiles_forger_roles" to table: "ForgerProfiles"
CREATE INDEX "idx_forger_profiles_forger_roles" ON "auth"."ForgerProfiles" USING GIN ("forger_roles");
-- Create index "idx_forger_profiles_profile_language" to table: "ForgerProfiles"
CREATE INDEX "idx_forger_profiles_profile_language" ON "auth"."ForgerProfiles" ("profile_language");
-- Set comment to table: "ForgerProfiles"
COMMENT ON TABLE "auth"."ForgerProfiles" IS 'End-user forger display profile; one account may own many profiles (1:N). Distinct from AuthorityProfiles. Phone numbers live in auth.Phones.';
-- Set comment to column: "id" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."id" IS 'Primary key UUID (profile_id); referenced by content/social as profile_id.';
-- Set comment to column: "account_id" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."account_id" IS 'Owning account; FK to auth.Accounts; one account may have many profiles; CASCADE delete.';
-- Set comment to column: "forger_roles" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."forger_roles" IS 'auth.forger_role[]; capability = membership. Default {forger}.';
-- Set comment to column: "profile_language" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."profile_language" IS 'Preferred UI locale (auth.profile_language).';
-- Set comment to column: "preference" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."preference" IS 'User preferences JSON: theme, layout, etc.';
-- Set comment to column: "display_name" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."display_name" IS 'Public handle / nickname; optional if first+last used.';
-- Set comment to column: "first_name" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."first_name" IS 'Given name; optional.';
-- Set comment to column: "last_name" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."last_name" IS 'Family name; optional.';
-- Set comment to column: "country" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."country" IS 'Country string for display or filters.';
-- Set comment to column: "city" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."city" IS 'City string.';
-- Set comment to column: "gender" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."gender" IS 'Self-identified gender label; app-defined vocabulary.';
-- Set comment to column: "birthday" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."birthday" IS 'Birth date if collected.';
-- Set comment to column: "website" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."website" IS 'Personal or social URL.';
-- Set comment to column: "timezone" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."timezone" IS 'IANA timezone name if set.';
-- Set comment to column: "bio" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."bio" IS 'Free-text bio.';
-- Set comment to column: "created_at" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."created_at" IS 'Row insert time.';
-- Set comment to column: "updated_at" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."updated_at" IS 'Last profile edit.';
-- Set comment to column: "deleted_at" on table: "ForgerProfiles"
COMMENT ON COLUMN "auth"."ForgerProfiles"."deleted_at" IS 'Soft-delete profile snapshot.';
-- Create "ForgerProfileAvatars" table
CREATE TABLE "auth"."ForgerProfileAvatars" (
  "id" uuid NOT NULL,
  "profile_id" uuid NOT NULL,
  "image_id" uuid NOT NULL,
  "is_active" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_forger_profile_avatars_profile_id" FOREIGN KEY ("profile_id") REFERENCES "auth"."ForgerProfiles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_forger_profile_avatars_deleted_at" to table: "ForgerProfileAvatars"
CREATE INDEX "idx_forger_profile_avatars_deleted_at" ON "auth"."ForgerProfileAvatars" ("deleted_at");
-- Create index "idx_forger_profile_avatars_image_id" to table: "ForgerProfileAvatars"
CREATE INDEX "idx_forger_profile_avatars_image_id" ON "auth"."ForgerProfileAvatars" ("image_id");
-- Create index "idx_forger_profile_avatars_profile_id" to table: "ForgerProfileAvatars"
CREATE INDEX "idx_forger_profile_avatars_profile_id" ON "auth"."ForgerProfileAvatars" ("profile_id");
-- Create index "uq_forger_profile_avatars_one_active_per_profile" to table: "ForgerProfileAvatars"
CREATE UNIQUE INDEX "uq_forger_profile_avatars_one_active_per_profile" ON "auth"."ForgerProfileAvatars" ("profile_id") WHERE ((is_active = true) AND (deleted_at IS NULL));
-- Set comment to table: "ForgerProfileAvatars"
COMMENT ON TABLE "auth"."ForgerProfileAvatars" IS 'Avatar history per forger profile; at most one active non-deleted row per ForgerProfiles row; image_id → asset.Images at app level.';
-- Set comment to column: "id" on table: "ForgerProfileAvatars"
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."id" IS 'Primary key UUID.';
-- Set comment to column: "profile_id" on table: "ForgerProfileAvatars"
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."profile_id" IS 'FK auth.ForgerProfiles.id; CASCADE delete.';
-- Set comment to column: "image_id" on table: "ForgerProfileAvatars"
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."image_id" IS 'Reference to asset.Images.id (application-level, no FK).';
-- Set comment to column: "is_active" on table: "ForgerProfileAvatars"
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."is_active" IS 'Exactly one non-deleted active row per profile_id; see partial unique index.';
-- Set comment to column: "created_at" on table: "ForgerProfileAvatars"
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."created_at" IS 'When this avatar record was created.';
-- Set comment to column: "updated_at" on table: "ForgerProfileAvatars"
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."updated_at" IS 'Last change to this row.';
-- Set comment to column: "deleted_at" on table: "ForgerProfileAvatars"
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."deleted_at" IS 'Soft-delete; set when this history row is removed from use.';
-- Create "ForgerProfileBackgrounds" table
CREATE TABLE "auth"."ForgerProfileBackgrounds" (
  "id" uuid NOT NULL,
  "profile_id" uuid NOT NULL,
  "image_id" uuid NOT NULL,
  "sort_order" integer NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_forger_profile_backgrounds_profile_id" FOREIGN KEY ("profile_id") REFERENCES "auth"."ForgerProfiles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_forger_profile_backgrounds_deleted_at" to table: "ForgerProfileBackgrounds"
CREATE INDEX "idx_forger_profile_backgrounds_deleted_at" ON "auth"."ForgerProfileBackgrounds" ("deleted_at");
-- Create index "idx_forger_profile_backgrounds_image_id" to table: "ForgerProfileBackgrounds"
CREATE INDEX "idx_forger_profile_backgrounds_image_id" ON "auth"."ForgerProfileBackgrounds" ("image_id");
-- Create index "idx_forger_profile_backgrounds_profile_id" to table: "ForgerProfileBackgrounds"
CREATE INDEX "idx_forger_profile_backgrounds_profile_id" ON "auth"."ForgerProfileBackgrounds" ("profile_id");
-- Create index "idx_forger_profile_backgrounds_profile_sort" to table: "ForgerProfileBackgrounds"
CREATE INDEX "idx_forger_profile_backgrounds_profile_sort" ON "auth"."ForgerProfileBackgrounds" ("profile_id", "sort_order");
-- Set comment to table: "ForgerProfileBackgrounds"
COMMENT ON TABLE "auth"."ForgerProfileBackgrounds" IS 'Background images per forger profile; order by sort_order; soft-delete via deleted_at; no cross-schema FK to asset.Images.';
-- Set comment to column: "id" on table: "ForgerProfileBackgrounds"
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."id" IS 'Primary key UUID.';
-- Set comment to column: "profile_id" on table: "ForgerProfileBackgrounds"
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."profile_id" IS 'FK auth.ForgerProfiles.id; CASCADE delete.';
-- Set comment to column: "image_id" on table: "ForgerProfileBackgrounds"
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."image_id" IS 'Reference to asset.Images.id (application-level, no FK).';
-- Set comment to column: "sort_order" on table: "ForgerProfileBackgrounds"
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."sort_order" IS 'Display order within profile; lower = higher priority.';
-- Set comment to column: "created_at" on table: "ForgerProfileBackgrounds"
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."created_at" IS 'When this background row was created.';
-- Set comment to column: "updated_at" on table: "ForgerProfileBackgrounds"
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."updated_at" IS 'Last change to this row.';
-- Set comment to column: "deleted_at" on table: "ForgerProfileBackgrounds"
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."deleted_at" IS 'Soft-delete row.';
-- Create "ForgerProfileSettings" table
CREATE TABLE "auth"."ForgerProfileSettings" (
  "id" uuid NOT NULL,
  "login_notification" boolean NOT NULL DEFAULT true,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_forger_profile_settings_id" FOREIGN KEY ("id") REFERENCES "auth"."ForgerProfiles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_forger_profile_settings_deleted_at" to table: "ForgerProfileSettings"
CREATE INDEX "idx_forger_profile_settings_deleted_at" ON "auth"."ForgerProfileSettings" ("deleted_at");
-- Set comment to table: "ForgerProfileSettings"
COMMENT ON TABLE "auth"."ForgerProfileSettings" IS 'Per-forger-profile system settings; id equals ForgerProfiles.id (1:1).';
-- Set comment to column: "id" on table: "ForgerProfileSettings"
COMMENT ON COLUMN "auth"."ForgerProfileSettings"."id" IS 'Primary key; same as auth.ForgerProfiles.id.';
-- Set comment to column: "login_notification" on table: "ForgerProfileSettings"
COMMENT ON COLUMN "auth"."ForgerProfileSettings"."login_notification" IS 'When true, send email on successful login.';
-- Set comment to column: "created_at" on table: "ForgerProfileSettings"
COMMENT ON COLUMN "auth"."ForgerProfileSettings"."created_at" IS 'When this settings row was created.';
-- Set comment to column: "updated_at" on table: "ForgerProfileSettings"
COMMENT ON COLUMN "auth"."ForgerProfileSettings"."updated_at" IS 'Last change to this row.';
-- Set comment to column: "deleted_at" on table: "ForgerProfileSettings"
COMMENT ON COLUMN "auth"."ForgerProfileSettings"."deleted_at" IS 'Soft-delete timestamp.';
-- Create "Identities" table
CREATE TABLE "auth"."Identities" (
  "id" uuid NOT NULL,
  "account_id" uuid NOT NULL,
  "identity_provider" "auth"."identity_provider" NOT NULL,
  "provider_subject" text NOT NULL,
  "password_hash" text NULL,
  "metadata" jsonb NULL,
  "last_login_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_identities_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_identities_account_id" to table: "Identities"
CREATE INDEX "idx_identities_account_id" ON "auth"."Identities" ("account_id");
-- Create index "idx_identities_deleted_at" to table: "Identities"
CREATE INDEX "idx_identities_deleted_at" ON "auth"."Identities" ("deleted_at");
-- Create index "uq_identities_identity_provider_subject_active" to table: "Identities"
CREATE UNIQUE INDEX "uq_identities_identity_provider_subject_active" ON "auth"."Identities" ("identity_provider", "provider_subject") WHERE (deleted_at IS NULL);
-- Set comment to table: "Identities"
COMMENT ON TABLE "auth"."Identities" IS 'Auth credentials per identity_provider; PASSWORD uses provider_subject = normalized login id (e.g. email).';
-- Set comment to column: "id" on table: "Identities"
COMMENT ON COLUMN "auth"."Identities"."id" IS 'Primary key UUID.';
-- Set comment to column: "account_id" on table: "Identities"
COMMENT ON COLUMN "auth"."Identities"."account_id" IS 'Owning auth.Accounts row; CASCADE delete.';
-- Set comment to column: "identity_provider" on table: "Identities"
COMMENT ON COLUMN "auth"."Identities"."identity_provider" IS 'auth.identity_provider enum (password / future OAuth).';
-- Set comment to column: "provider_subject" on table: "Identities"
COMMENT ON COLUMN "auth"."Identities"."provider_subject" IS 'Unique within identity_provider among non-deleted rows; email for password, sub for OAuth.';
-- Set comment to column: "password_hash" on table: "Identities"
COMMENT ON COLUMN "auth"."Identities"."password_hash" IS 'Set only when identity_provider = password; bcrypt/argon2 hash at application layer.';
-- Set comment to column: "metadata" on table: "Identities"
COMMENT ON COLUMN "auth"."Identities"."metadata" IS 'Optional provider-specific payload (e.g. OAuth raw claims subset).';
-- Set comment to column: "last_login_at" on table: "Identities"
COMMENT ON COLUMN "auth"."Identities"."last_login_at" IS 'Last successful login instant for this identity.';
-- Set comment to column: "created_at" on table: "Identities"
COMMENT ON COLUMN "auth"."Identities"."created_at" IS 'Row insert time.';
-- Set comment to column: "updated_at" on table: "Identities"
COMMENT ON COLUMN "auth"."Identities"."updated_at" IS 'Last credential/metadata update.';
-- Set comment to column: "deleted_at" on table: "Identities"
COMMENT ON COLUMN "auth"."Identities"."deleted_at" IS 'Soft-delete identity; frees provider_subject for re-link.';
-- Create "Phones" table
CREATE TABLE "auth"."Phones" (
  "id" uuid NOT NULL,
  "account_id" uuid NOT NULL,
  "phone" character varying(32) NOT NULL,
  "is_primary" boolean NOT NULL DEFAULT false,
  "verified_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_phones_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_phones_account_id" to table: "Phones"
CREATE INDEX "idx_phones_account_id" ON "auth"."Phones" ("account_id");
-- Create index "idx_phones_deleted_at" to table: "Phones"
CREATE INDEX "idx_phones_deleted_at" ON "auth"."Phones" ("deleted_at");
-- Create index "uq_phones_one_primary_per_account" to table: "Phones"
CREATE UNIQUE INDEX "uq_phones_one_primary_per_account" ON "auth"."Phones" ("account_id") WHERE ((is_primary = true) AND (deleted_at IS NULL));
-- Create index "uq_phones_phone_active" to table: "Phones"
CREATE UNIQUE INDEX "uq_phones_phone_active" ON "auth"."Phones" ("phone") WHERE (deleted_at IS NULL);
-- Set comment to table: "Phones"
COMMENT ON TABLE "auth"."Phones" IS 'Phone bindings; partial unique on phone WHERE deleted_at IS NULL mirrors Emails pattern.';
-- Set comment to column: "id" on table: "Phones"
COMMENT ON COLUMN "auth"."Phones"."id" IS 'Primary key UUID.';
-- Set comment to column: "account_id" on table: "Phones"
COMMENT ON COLUMN "auth"."Phones"."account_id" IS 'Owning auth.Accounts row; CASCADE delete.';
-- Set comment to column: "phone" on table: "Phones"
COMMENT ON COLUMN "auth"."Phones"."phone" IS 'Canonical phone string; unique among non-deleted rows.';
-- Set comment to column: "is_primary" on table: "Phones"
COMMENT ON COLUMN "auth"."Phones"."is_primary" IS 'At most one non-deleted primary phone per account.';
-- Set comment to column: "verified_at" on table: "Phones"
COMMENT ON COLUMN "auth"."Phones"."verified_at" IS 'NULL = not verified yet.';
-- Set comment to column: "created_at" on table: "Phones"
COMMENT ON COLUMN "auth"."Phones"."created_at" IS 'Row insert time.';
-- Set comment to column: "updated_at" on table: "Phones"
COMMENT ON COLUMN "auth"."Phones"."updated_at" IS 'Last edit time.';
-- Set comment to column: "deleted_at" on table: "Phones"
COMMENT ON COLUMN "auth"."Phones"."deleted_at" IS 'Soft-delete row; same canonical phone may be re-bound via INSERT new row.';
-- Create "RefreshTokens" table
CREATE TABLE "auth"."RefreshTokens" (
  "id" uuid NOT NULL,
  "account_id" uuid NOT NULL,
  "identity_id" uuid NULL,
  "profile_id" uuid NULL,
  "profile_scope" "auth"."profile_scope" NULL,
  "device_id" text NULL,
  "token_hash" text NOT NULL,
  "expires_at" timestamp NOT NULL,
  "revoked_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_refresh_tokens_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_refresh_tokens_identity_id" FOREIGN KEY ("identity_id") REFERENCES "auth"."Identities" ("id") ON UPDATE NO ACTION ON DELETE SET NULL
);
-- Create index "idx_refresh_tokens_account_device_active" to table: "RefreshTokens"
CREATE INDEX "idx_refresh_tokens_account_device_active" ON "auth"."RefreshTokens" ("account_id", "device_id") WHERE ((device_id IS NOT NULL) AND (revoked_at IS NULL) AND (deleted_at IS NULL));
-- Create index "idx_refresh_tokens_account_id" to table: "RefreshTokens"
CREATE INDEX "idx_refresh_tokens_account_id" ON "auth"."RefreshTokens" ("account_id");
-- Create index "idx_refresh_tokens_deleted_at" to table: "RefreshTokens"
CREATE INDEX "idx_refresh_tokens_deleted_at" ON "auth"."RefreshTokens" ("deleted_at");
-- Create index "idx_refresh_tokens_device_id" to table: "RefreshTokens"
CREATE INDEX "idx_refresh_tokens_device_id" ON "auth"."RefreshTokens" ("device_id");
-- Create index "idx_refresh_tokens_expires_at" to table: "RefreshTokens"
CREATE INDEX "idx_refresh_tokens_expires_at" ON "auth"."RefreshTokens" ("expires_at");
-- Create index "idx_refresh_tokens_identity_id" to table: "RefreshTokens"
CREATE INDEX "idx_refresh_tokens_identity_id" ON "auth"."RefreshTokens" ("identity_id");
-- Create index "idx_refresh_tokens_revoked_at" to table: "RefreshTokens"
CREATE INDEX "idx_refresh_tokens_revoked_at" ON "auth"."RefreshTokens" ("revoked_at");
-- Set comment to table: "RefreshTokens"
COMMENT ON TABLE "auth"."RefreshTokens" IS 'Opaque refresh tokens; store only hashes at rest.';
-- Set comment to column: "id" on table: "RefreshTokens"
COMMENT ON COLUMN "auth"."RefreshTokens"."id" IS 'Primary key UUID.';
-- Set comment to column: "account_id" on table: "RefreshTokens"
COMMENT ON COLUMN "auth"."RefreshTokens"."account_id" IS 'Owning account; CASCADE delete.';
-- Set comment to column: "identity_id" on table: "RefreshTokens"
COMMENT ON COLUMN "auth"."RefreshTokens"."identity_id" IS 'Which login issued the token; NULL if legacy / unspecified.';
-- Set comment to column: "device_id" on table: "RefreshTokens"
COMMENT ON COLUMN "auth"."RefreshTokens"."device_id" IS 'Client device/session identifier for same-device token revocation policy.';
-- Set comment to column: "token_hash" on table: "RefreshTokens"
COMMENT ON COLUMN "auth"."RefreshTokens"."token_hash" IS 'Hash of opaque refresh token at rest.';
-- Set comment to column: "expires_at" on table: "RefreshTokens"
COMMENT ON COLUMN "auth"."RefreshTokens"."expires_at" IS 'Absolute expiry instant.';
-- Set comment to column: "revoked_at" on table: "RefreshTokens"
COMMENT ON COLUMN "auth"."RefreshTokens"."revoked_at" IS 'When token was invalidated server-side.';
-- Set comment to column: "created_at" on table: "RefreshTokens"
COMMENT ON COLUMN "auth"."RefreshTokens"."created_at" IS 'Issued-at time.';
-- Set comment to column: "deleted_at" on table: "RefreshTokens"
COMMENT ON COLUMN "auth"."RefreshTokens"."deleted_at" IS 'Soft-delete token row (e.g. purge); prefer revoked_at for normal logout.';
-- Create "FullAccountInformationWithAuthorityProfileSettingsView" view
CREATE VIEW "auth"."FullAccountInformationWithAuthorityProfileSettingsView" (
  "id",
  "login_notification",
  "created_at",
  "updated_at"
) AS SELECT id,
    login_notification,
    created_at,
    updated_at
   FROM auth."AuthorityProfileSettings" s
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithAuthorityProfileSettingsView"
COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileSettingsView" IS 'Non-deleted profile settings; query by id (profile_id) for the single-authority-profile read model.';
-- Create "FullAccountInformationPhoneView" view
CREATE VIEW "auth"."FullAccountInformationPhoneView" (
  "id",
  "account_id",
  "phone",
  "is_primary",
  "verified_at",
  "created_at",
  "updated_at"
) AS SELECT id,
    account_id,
    phone,
    is_primary,
    verified_at,
    created_at,
    updated_at
   FROM auth."Phones" ph
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationPhoneView"
COMMENT ON VIEW "auth"."FullAccountInformationPhoneView" IS 'Non-deleted phones; query by account_id for account details.';
-- Create "SystemStateActiveView" view
CREATE VIEW "system"."SystemStateActiveView" (
  "id",
  "initialized",
  "initialized_at",
  "initialization_version",
  "last_reset_at",
  "last_reset_by",
  "reset_count",
  "metadata",
  "created_at",
  "updated_at"
) AS SELECT id,
    initialized,
    initialized_at,
    initialization_version,
    last_reset_at,
    last_reset_by,
    reset_count,
    metadata,
    created_at,
    updated_at
   FROM system.system_state;
-- Set comment to view: "SystemStateActiveView"
COMMENT ON VIEW "system"."SystemStateActiveView" IS 'All system_state rows for latest-by-created_at reads.';
-- Create "AudiosActiveView" view
CREATE VIEW "asset"."AudiosActiveView" (
  "id",
  "file_path",
  "file_name",
  "file_size",
  "mime_type",
  "duration_seconds",
  "uploader_id",
  "created_at",
  "updated_at"
) AS SELECT id,
    file_path,
    file_name,
    file_size,
    mime_type,
    duration_seconds,
    uploader_id,
    created_at,
    updated_at
   FROM asset."Audios"
  WHERE deleted_at IS NULL;
-- Set comment to view: "AudiosActiveView"
COMMENT ON VIEW "asset"."AudiosActiveView" IS 'Non-deleted audio rows only; omits deleted_at column.';
-- Create "FilesActiveView" view
CREATE VIEW "asset"."FilesActiveView" (
  "id",
  "file_path",
  "file_name",
  "file_size",
  "mime_type",
  "uploader_id",
  "created_at",
  "updated_at"
) AS SELECT id,
    file_path,
    file_name,
    file_size,
    mime_type,
    uploader_id,
    created_at,
    updated_at
   FROM asset."Files"
  WHERE deleted_at IS NULL;
-- Set comment to view: "FilesActiveView"
COMMENT ON VIEW "asset"."FilesActiveView" IS 'Non-deleted file rows only; omits deleted_at column.';
-- Create "ImagesActiveView" view
CREATE VIEW "asset"."ImagesActiveView" (
  "id",
  "file_path",
  "file_name",
  "file_size",
  "mime_type",
  "width",
  "height",
  "alt_text",
  "uploader_id",
  "created_at",
  "updated_at"
) AS SELECT id,
    file_path,
    file_name,
    file_size,
    mime_type,
    width,
    height,
    alt_text,
    uploader_id,
    created_at,
    updated_at
   FROM asset."Images"
  WHERE deleted_at IS NULL;
-- Set comment to view: "ImagesActiveView"
COMMENT ON VIEW "asset"."ImagesActiveView" IS 'Non-deleted image rows only; omits deleted_at column.';
-- Create "VideosActiveView" view
CREATE VIEW "asset"."VideosActiveView" (
  "id",
  "file_path",
  "file_name",
  "file_size",
  "mime_type",
  "duration_seconds",
  "width",
  "height",
  "uploader_id",
  "created_at",
  "updated_at"
) AS SELECT id,
    file_path,
    file_name,
    file_size,
    mime_type,
    duration_seconds,
    width,
    height,
    uploader_id,
    created_at,
    updated_at
   FROM asset."Videos"
  WHERE deleted_at IS NULL;
-- Set comment to view: "VideosActiveView"
COMMENT ON VIEW "asset"."VideosActiveView" IS 'Non-deleted video rows only; omits deleted_at column.';
-- Create "ListPhoneItems" view
CREATE VIEW "auth"."ListPhoneItems" (
  "id",
  "account_id",
  "phone",
  "is_primary",
  "verified_at",
  "created_at",
  "updated_at"
) AS SELECT id,
    account_id,
    phone,
    is_primary,
    verified_at,
    created_at,
    updated_at
   FROM auth."Phones" ph
  WHERE deleted_at IS NULL;
-- Set comment to view: "ListPhoneItems"
COMMENT ON VIEW "auth"."ListPhoneItems" IS 'One row per non-deleted phone. Filter by account_id to list account phones.';
-- Create "ListForgerProfileItems" view
CREATE VIEW "auth"."ListForgerProfileItems" (
  "profile_id",
  "account_id",
  "forger_roles",
  "display_name",
  "profile_language",
  "city",
  "country",
  "website",
  "timezone",
  "birthday",
  "created_at",
  "avatar_image_id",
  "background_image_id"
) AS SELECT p.id AS profile_id,
    p.account_id,
    p.forger_roles,
    p.display_name,
    p.profile_language,
    p.city,
    p.country,
    p.website,
    p.timezone,
    p.birthday,
    p.created_at,
    av.image_id AS avatar_image_id,
    bg.image_id AS background_image_id
   FROM auth."ForgerProfiles" p
     LEFT JOIN LATERAL ( SELECT a.image_id
           FROM auth."ForgerProfileAvatars" a
          WHERE a.profile_id = p.id AND a.is_active = true AND a.deleted_at IS NULL
          ORDER BY a.created_at DESC
         LIMIT 1) av ON true
     LEFT JOIN LATERAL ( SELECT b.image_id
           FROM auth."ForgerProfileBackgrounds" b
          WHERE b.profile_id = p.id AND b.deleted_at IS NULL
          ORDER BY b.sort_order, b.created_at
         LIMIT 1) bg ON true
  WHERE p.deleted_at IS NULL;
-- Set comment to view: "ListForgerProfileItems"
COMMENT ON VIEW "auth"."ListForgerProfileItems" IS 'One row per non-deleted forger profile with public list/search fields, active avatar, and primary background (lowest sort_order). Private fields (first_name, last_name) are not exposed; read those from account-scoped profile APIs only.';
-- Create "ListEmailItems" view
CREATE VIEW "auth"."ListEmailItems" (
  "id",
  "account_id",
  "email",
  "is_primary",
  "verified_at",
  "created_at",
  "updated_at"
) AS SELECT id,
    account_id,
    email,
    is_primary,
    verified_at,
    created_at,
    updated_at
   FROM auth."Emails" em
  WHERE deleted_at IS NULL;
-- Set comment to view: "ListEmailItems"
COMMENT ON VIEW "auth"."ListEmailItems" IS 'One row per non-deleted email. Filter by account_id to list account emails.';
-- Create "ListAuthorityProfileItems" view
CREATE VIEW "auth"."ListAuthorityProfileItems" (
  "profile_id",
  "account_id",
  "authority_roles",
  "display_name",
  "profile_language",
  "city",
  "country",
  "website",
  "timezone",
  "birthday",
  "created_at",
  "avatar_image_id",
  "background_image_id"
) AS SELECT p.id AS profile_id,
    p.account_id,
    p.authority_roles,
    p.display_name,
    p.profile_language,
    p.city,
    p.country,
    p.website,
    p.timezone,
    p.birthday,
    p.created_at,
    av.image_id AS avatar_image_id,
    bg.image_id AS background_image_id
   FROM auth."AuthorityProfiles" p
     LEFT JOIN LATERAL ( SELECT a.image_id
           FROM auth."AuthorityProfileAvatars" a
          WHERE a.profile_id = p.id AND a.is_active = true AND a.deleted_at IS NULL
          ORDER BY a.created_at DESC
         LIMIT 1) av ON true
     LEFT JOIN LATERAL ( SELECT b.image_id
           FROM auth."AuthorityProfileBackgrounds" b
          WHERE b.profile_id = p.id AND b.deleted_at IS NULL
          ORDER BY b.sort_order, b.created_at
         LIMIT 1) bg ON true
  WHERE p.deleted_at IS NULL;
-- Set comment to view: "ListAuthorityProfileItems"
COMMENT ON VIEW "auth"."ListAuthorityProfileItems" IS 'One row per non-deleted authority profile with public list/search fields, active avatar, and primary background (lowest sort_order). Mirrors ListCommunityProfileItems but carries authority_roles[].';
-- Create "FullAccountInformationWithForgerProfileView" view
CREATE VIEW "auth"."FullAccountInformationWithForgerProfileView" (
  "id",
  "account_id",
  "forger_roles",
  "profile_language",
  "preference",
  "display_name",
  "first_name",
  "last_name",
  "country",
  "city",
  "gender",
  "birthday",
  "website",
  "timezone",
  "bio",
  "created_at",
  "updated_at"
) AS SELECT id,
    account_id,
    forger_roles,
    profile_language,
    preference,
    display_name,
    first_name,
    last_name,
    country,
    city,
    gender,
    birthday,
    website,
    timezone,
    bio,
    created_at,
    updated_at
   FROM auth."ForgerProfiles" p
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithForgerProfileView"
COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileView" IS 'Non-deleted forger profiles; query by id (current profile from token) or account_id (fallback). Backs the single-forger-profile read model.';
-- Create "FullAccountInformationWithForgerProfileSettingsView" view
CREATE VIEW "auth"."FullAccountInformationWithForgerProfileSettingsView" (
  "id",
  "login_notification",
  "created_at",
  "updated_at"
) AS SELECT id,
    login_notification,
    created_at,
    updated_at
   FROM auth."ForgerProfileSettings" s
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithForgerProfileSettingsView"
COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileSettingsView" IS 'Non-deleted profile settings; query by id (profile_id) for the single-forger-profile read model.';
-- Create "FullAccountInformationWithForgerProfilePhoneView" view
CREATE VIEW "auth"."FullAccountInformationWithForgerProfilePhoneView" (
  "id",
  "account_id",
  "phone",
  "is_primary",
  "verified_at",
  "created_at",
  "updated_at"
) AS SELECT id,
    account_id,
    phone,
    is_primary,
    verified_at,
    created_at,
    updated_at
   FROM auth."Phones" ph
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithForgerProfilePhoneView"
COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfilePhoneView" IS 'Non-deleted phones; query by account_id for the single-forger-profile read model.';
-- Create "FullAccountInformationWithForgerProfileEmailView" view
CREATE VIEW "auth"."FullAccountInformationWithForgerProfileEmailView" (
  "id",
  "account_id",
  "email",
  "is_primary",
  "verified_at",
  "created_at",
  "updated_at"
) AS SELECT id,
    account_id,
    email,
    is_primary,
    verified_at,
    created_at,
    updated_at
   FROM auth."Emails" em
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithForgerProfileEmailView"
COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileEmailView" IS 'Non-deleted emails; query by account_id for the single-forger-profile read model.';
-- Create "FullAccountInformationWithForgerProfileBackgroundView" view
CREATE VIEW "auth"."FullAccountInformationWithForgerProfileBackgroundView" (
  "id",
  "profile_id",
  "image_id",
  "sort_order",
  "created_at",
  "updated_at"
) AS SELECT id,
    profile_id,
    image_id,
    sort_order,
    created_at,
    updated_at
   FROM auth."ForgerProfileBackgrounds" bg
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithForgerProfileBackgroundView"
COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileBackgroundView" IS 'Non-deleted profile background links; query by profile_id for the single-forger-profile read model.';
-- Create "FullAccountInformationAccountView" view
CREATE VIEW "auth"."FullAccountInformationAccountView" (
  "account_id",
  "account_status",
  "signup_platform",
  "created_at",
  "updated_at"
) AS SELECT id AS account_id,
    account_status,
    signup_platform,
    created_at,
    updated_at
   FROM auth."Accounts" a
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationAccountView"
COMMENT ON VIEW "auth"."FullAccountInformationAccountView" IS 'Active account base fields only; one row per account.';
-- Create "FullAccountInformationAvatarView" view
CREATE VIEW "auth"."FullAccountInformationAvatarView" (
  "id",
  "profile_id",
  "image_id",
  "is_active",
  "created_at",
  "updated_at"
) AS SELECT id,
    profile_id,
    image_id,
    is_active,
    created_at,
    updated_at
   FROM auth."ForgerProfileAvatars" av
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationAvatarView"
COMMENT ON VIEW "auth"."FullAccountInformationAvatarView" IS 'Non-deleted forger avatar links; query by profile_id.';
-- Create "FullAccountInformationBackgroundView" view
CREATE VIEW "auth"."FullAccountInformationBackgroundView" (
  "id",
  "profile_id",
  "image_id",
  "sort_order",
  "created_at",
  "updated_at"
) AS SELECT id,
    profile_id,
    image_id,
    sort_order,
    created_at,
    updated_at
   FROM auth."ForgerProfileBackgrounds" bg
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationBackgroundView"
COMMENT ON VIEW "auth"."FullAccountInformationBackgroundView" IS 'Non-deleted forger profile background links; query by profile_id.';
-- Create "FullAccountInformationEmailView" view
CREATE VIEW "auth"."FullAccountInformationEmailView" (
  "id",
  "account_id",
  "email",
  "is_primary",
  "verified_at",
  "created_at",
  "updated_at"
) AS SELECT id,
    account_id,
    email,
    is_primary,
    verified_at,
    created_at,
    updated_at
   FROM auth."Emails" em
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationEmailView"
COMMENT ON VIEW "auth"."FullAccountInformationEmailView" IS 'Non-deleted emails; query by account_id for account details.';
-- Create "FullAccountInformationForgerProfileSettingsView" view
CREATE VIEW "auth"."FullAccountInformationForgerProfileSettingsView" (
  "id",
  "login_notification",
  "created_at",
  "updated_at"
) AS SELECT id,
    login_notification,
    created_at,
    updated_at
   FROM auth."ForgerProfileSettings" s
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationForgerProfileSettingsView"
COMMENT ON VIEW "auth"."FullAccountInformationForgerProfileSettingsView" IS 'Non-deleted forger profile settings; query by id (profile_id).';
-- Create "FullAccountInformationForgerProfileView" view
CREATE VIEW "auth"."FullAccountInformationForgerProfileView" (
  "id",
  "account_id",
  "forger_roles",
  "profile_language",
  "preference",
  "display_name",
  "first_name",
  "last_name",
  "country",
  "city",
  "gender",
  "birthday",
  "website",
  "timezone",
  "bio",
  "created_at",
  "updated_at"
) AS SELECT id,
    account_id,
    forger_roles,
    profile_language,
    preference,
    display_name,
    first_name,
    last_name,
    country,
    city,
    gender,
    birthday,
    website,
    timezone,
    bio,
    created_at,
    updated_at
   FROM auth."ForgerProfiles" p
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationForgerProfileView"
COMMENT ON VIEW "auth"."FullAccountInformationForgerProfileView" IS 'Non-deleted forger profiles; one account may have many rows; query by account_id (list) or id (single profile).';
-- Create "FullAccountInformationWithForgerProfileAvatarView" view
CREATE VIEW "auth"."FullAccountInformationWithForgerProfileAvatarView" (
  "id",
  "profile_id",
  "image_id",
  "is_active",
  "created_at",
  "updated_at"
) AS SELECT id,
    profile_id,
    image_id,
    is_active,
    created_at,
    updated_at
   FROM auth."ForgerProfileAvatars" av
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithForgerProfileAvatarView"
COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileAvatarView" IS 'Non-deleted avatar links; query by profile_id for the single-forger-profile read model.';
-- Create "FullAccountInformationWithAuthorityProfileAccountView" view
CREATE VIEW "auth"."FullAccountInformationWithAuthorityProfileAccountView" (
  "account_id",
  "account_status",
  "signup_platform",
  "created_at",
  "updated_at"
) AS SELECT id AS account_id,
    account_status,
    signup_platform,
    created_at,
    updated_at
   FROM auth."Accounts" a
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithAuthorityProfileAccountView"
COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileAccountView" IS 'Active account base fields only; one row per account. Backs the single-authority-profile login-state read model.';
-- Create "FullAccountInformationWithAuthorityProfileAvatarView" view
CREATE VIEW "auth"."FullAccountInformationWithAuthorityProfileAvatarView" (
  "id",
  "profile_id",
  "image_id",
  "is_active",
  "created_at",
  "updated_at"
) AS SELECT id,
    profile_id,
    image_id,
    is_active,
    created_at,
    updated_at
   FROM auth."AuthorityProfileAvatars" av
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithAuthorityProfileAvatarView"
COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileAvatarView" IS 'Non-deleted avatar links; query by profile_id for the single-authority-profile read model.';
-- Create "FullAccountInformationWithAuthorityProfileBackgroundView" view
CREATE VIEW "auth"."FullAccountInformationWithAuthorityProfileBackgroundView" (
  "id",
  "profile_id",
  "image_id",
  "sort_order",
  "created_at",
  "updated_at"
) AS SELECT id,
    profile_id,
    image_id,
    sort_order,
    created_at,
    updated_at
   FROM auth."AuthorityProfileBackgrounds" bg
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithAuthorityProfileBackgroundView"
COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileBackgroundView" IS 'Non-deleted profile background links; query by profile_id for the single-authority-profile read model.';
-- Create "FullAccountInformationWithAuthorityProfileEmailView" view
CREATE VIEW "auth"."FullAccountInformationWithAuthorityProfileEmailView" (
  "id",
  "account_id",
  "email",
  "is_primary",
  "verified_at",
  "created_at",
  "updated_at"
) AS SELECT id,
    account_id,
    email,
    is_primary,
    verified_at,
    created_at,
    updated_at
   FROM auth."Emails" em
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithAuthorityProfileEmailView"
COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileEmailView" IS 'Non-deleted emails; query by account_id for the single-authority-profile read model.';
-- Create "FullAccountInformationWithAuthorityProfilePhoneView" view
CREATE VIEW "auth"."FullAccountInformationWithAuthorityProfilePhoneView" (
  "id",
  "account_id",
  "phone",
  "is_primary",
  "verified_at",
  "created_at",
  "updated_at"
) AS SELECT id,
    account_id,
    phone,
    is_primary,
    verified_at,
    created_at,
    updated_at
   FROM auth."Phones" ph
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithAuthorityProfilePhoneView"
COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfilePhoneView" IS 'Non-deleted phones; query by account_id for the single-authority-profile read model.';
-- Create "FullAccountInformationWithAuthorityProfileView" view
CREATE VIEW "auth"."FullAccountInformationWithAuthorityProfileView" (
  "id",
  "account_id",
  "authority_roles",
  "profile_language",
  "preference",
  "display_name",
  "first_name",
  "last_name",
  "country",
  "city",
  "gender",
  "birthday",
  "website",
  "timezone",
  "bio",
  "created_at",
  "updated_at"
) AS SELECT id,
    account_id,
    authority_roles,
    profile_language,
    preference,
    display_name,
    first_name,
    last_name,
    country,
    city,
    gender,
    birthday,
    website,
    timezone,
    bio,
    created_at,
    updated_at
   FROM auth."AuthorityProfiles" p
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithAuthorityProfileView"
COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileView" IS 'Non-deleted authority profiles; query by id (current profile from token) or account_id (fallback). Backs the single-authority-profile read model.';
-- Create "FullAccountInformationWithForgerProfileAccountView" view
CREATE VIEW "auth"."FullAccountInformationWithForgerProfileAccountView" (
  "account_id",
  "account_status",
  "signup_platform",
  "created_at",
  "updated_at"
) AS SELECT id AS account_id,
    account_status,
    signup_platform,
    created_at,
    updated_at
   FROM auth."Accounts" a
  WHERE deleted_at IS NULL;
-- Set comment to view: "FullAccountInformationWithForgerProfileAccountView"
COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileAccountView" IS 'Active account base fields only; one row per account. Backs the single-forger-profile login-state read model.';
