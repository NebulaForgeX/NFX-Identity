CREATE TABLE IF NOT EXISTS "auth"."ForgerProfileBackgrounds" (
  "id" UUID PRIMARY KEY,
  "profile_id" UUID NOT NULL,
  "image_id" UUID NOT NULL,
  "sort_order" INTEGER NOT NULL DEFAULT 0,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_forger_profile_backgrounds_profile_id" FOREIGN KEY ("profile_id") REFERENCES "auth"."ForgerProfiles" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_forger_profile_backgrounds_profile_id" ON "auth"."ForgerProfileBackgrounds" ("profile_id");
CREATE INDEX IF NOT EXISTS "idx_forger_profile_backgrounds_profile_sort" ON "auth"."ForgerProfileBackgrounds" ("profile_id", "sort_order");
CREATE INDEX IF NOT EXISTS "idx_forger_profile_backgrounds_image_id" ON "auth"."ForgerProfileBackgrounds" ("image_id");
CREATE INDEX IF NOT EXISTS "idx_forger_profile_backgrounds_deleted_at" ON "auth"."ForgerProfileBackgrounds" ("deleted_at");

COMMENT ON TABLE "auth"."ForgerProfileBackgrounds" IS 'Background images per forger profile; order by sort_order; soft-delete via deleted_at; no cross-schema FK to asset.Images.';
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."profile_id" IS 'FK auth.ForgerProfiles.id; CASCADE delete.';
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."image_id" IS 'Reference to asset.Images.id (application-level, no FK).';
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."sort_order" IS 'Display order within profile; lower = higher priority.';
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."created_at" IS 'When this background row was created.';
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."updated_at" IS 'Last change to this row.';
COMMENT ON COLUMN "auth"."ForgerProfileBackgrounds"."deleted_at" IS 'Soft-delete row.';

-- Maintainer notes (after DDL):
-- profile_id = ForgerProfiles.id; sort_order + business rules for "current" image.
