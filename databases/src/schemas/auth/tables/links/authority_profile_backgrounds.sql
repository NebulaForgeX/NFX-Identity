CREATE TABLE IF NOT EXISTS "auth"."AuthorityProfileBackgrounds" (
  "id" UUID PRIMARY KEY,
  "profile_id" UUID NOT NULL,
  "image_id" UUID NOT NULL,
  "sort_order" INTEGER NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_authority_profile_backgrounds_profile_id" FOREIGN KEY ("profile_id") REFERENCES "auth"."AuthorityProfiles" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_authority_profile_backgrounds_profile_id" ON "auth"."AuthorityProfileBackgrounds" ("profile_id");
CREATE INDEX IF NOT EXISTS "idx_authority_profile_backgrounds_profile_sort" ON "auth"."AuthorityProfileBackgrounds" ("profile_id", "sort_order");
CREATE INDEX IF NOT EXISTS "idx_authority_profile_backgrounds_image_id" ON "auth"."AuthorityProfileBackgrounds" ("image_id");
CREATE INDEX IF NOT EXISTS "idx_authority_profile_backgrounds_deleted_at" ON "auth"."AuthorityProfileBackgrounds" ("deleted_at");

COMMENT ON TABLE "auth"."AuthorityProfileBackgrounds" IS 'Background images per authority profile; order by sort_order; soft-delete via deleted_at; no cross-schema FK to asset.Images.';
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."profile_id" IS 'FK auth.AuthorityProfiles.id; CASCADE delete.';
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."image_id" IS 'Reference to asset.Images.id (application-level, no FK).';
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."sort_order" IS 'Display order within profile; lower = higher priority.';
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."created_at" IS 'When this background row was created.';
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."updated_at" IS 'Last change to this row.';
COMMENT ON COLUMN "auth"."AuthorityProfileBackgrounds"."deleted_at" IS 'Soft-delete row.';

-- Maintainer notes (after DDL):
-- profile_id = AuthorityProfiles.id; sort_order + business rules for "current" image.
