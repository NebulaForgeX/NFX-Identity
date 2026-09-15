CREATE TABLE IF NOT EXISTS "auth"."AuthorityProfileAvatars" (
  "id" UUID PRIMARY KEY,
  "profile_id" UUID NOT NULL,
  "image_id" UUID NOT NULL,
  "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_authority_profile_avatars_profile_id" FOREIGN KEY ("profile_id") REFERENCES "auth"."AuthorityProfiles" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_authority_profile_avatars_profile_id" ON "auth"."AuthorityProfileAvatars" ("profile_id");
CREATE INDEX IF NOT EXISTS "idx_authority_profile_avatars_image_id" ON "auth"."AuthorityProfileAvatars" ("image_id");
CREATE INDEX IF NOT EXISTS "idx_authority_profile_avatars_deleted_at" ON "auth"."AuthorityProfileAvatars" ("deleted_at");
CREATE UNIQUE INDEX IF NOT EXISTS "uq_authority_profile_avatars_one_active_per_profile" ON "auth"."AuthorityProfileAvatars" ("profile_id") WHERE "is_active" = TRUE AND "deleted_at" IS NULL;

COMMENT ON TABLE "auth"."AuthorityProfileAvatars" IS 'Avatar history per authority profile; at most one active non-deleted row per AuthorityProfiles row; image_id → asset.Images at app level.';
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."profile_id" IS 'FK auth.AuthorityProfiles.id; CASCADE delete.';
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."image_id" IS 'Reference to asset.Images.id (application-level, no FK).';
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."is_active" IS 'Exactly one non-deleted active row per profile_id; see partial unique index.';
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."created_at" IS 'When this avatar record was created.';
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."updated_at" IS 'Last change to this row.';
COMMENT ON COLUMN "auth"."AuthorityProfileAvatars"."deleted_at" IS 'Soft-delete; set when this history row is removed from use.';

-- Maintainer notes (after DDL):
-- profile_id = AuthorityProfiles.id; one active avatar per profile.
