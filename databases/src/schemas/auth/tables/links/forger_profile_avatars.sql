CREATE TABLE IF NOT EXISTS "auth"."ForgerProfileAvatars" (
  "id" UUID PRIMARY KEY,
  "profile_id" UUID NOT NULL,
  "image_id" UUID NOT NULL,
  "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_forger_profile_avatars_profile_id" FOREIGN KEY ("profile_id") REFERENCES "auth"."ForgerProfiles" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_forger_profile_avatars_profile_id" ON "auth"."ForgerProfileAvatars" ("profile_id");
CREATE INDEX IF NOT EXISTS "idx_forger_profile_avatars_image_id" ON "auth"."ForgerProfileAvatars" ("image_id");
CREATE INDEX IF NOT EXISTS "idx_forger_profile_avatars_deleted_at" ON "auth"."ForgerProfileAvatars" ("deleted_at");
CREATE UNIQUE INDEX IF NOT EXISTS "uq_forger_profile_avatars_one_active_per_profile" ON "auth"."ForgerProfileAvatars" ("profile_id") WHERE "is_active" = TRUE AND "deleted_at" IS NULL;

COMMENT ON TABLE "auth"."ForgerProfileAvatars" IS 'Avatar history per forger profile; at most one active non-deleted row per ForgerProfiles row; image_id → asset.Images at app level.';
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."profile_id" IS 'FK auth.ForgerProfiles.id; CASCADE delete.';
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."image_id" IS 'Reference to asset.Images.id (application-level, no FK).';
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."is_active" IS 'Exactly one non-deleted active row per profile_id; see partial unique index.';
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."created_at" IS 'When this avatar record was created.';
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."updated_at" IS 'Last change to this row.';
COMMENT ON COLUMN "auth"."ForgerProfileAvatars"."deleted_at" IS 'Soft-delete; set when this history row is removed from use.';

-- Maintainer notes (after DDL):
-- profile_id = ForgerProfiles.id; one active avatar per profile.
