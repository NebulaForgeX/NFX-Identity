CREATE TABLE IF NOT EXISTS "auth"."AuthorityProfileSettings" (
  "id" UUID PRIMARY KEY,
  "login_notification" BOOLEAN NOT NULL DEFAULT TRUE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_authority_profile_settings_id" FOREIGN KEY ("id") REFERENCES "auth"."AuthorityProfiles" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_authority_profile_settings_deleted_at" ON "auth"."AuthorityProfileSettings" ("deleted_at");

COMMENT ON TABLE "auth"."AuthorityProfileSettings" IS 'Per-authority-profile system settings; id equals AuthorityProfiles.id (1:1).';
COMMENT ON COLUMN "auth"."AuthorityProfileSettings"."id" IS 'Primary key; same as auth.AuthorityProfiles.id.';
COMMENT ON COLUMN "auth"."AuthorityProfileSettings"."login_notification" IS 'When true, send email on successful login.';
COMMENT ON COLUMN "auth"."AuthorityProfileSettings"."created_at" IS 'When this settings row was created.';
COMMENT ON COLUMN "auth"."AuthorityProfileSettings"."updated_at" IS 'Last change to this row.';
COMMENT ON COLUMN "auth"."AuthorityProfileSettings"."deleted_at" IS 'Soft-delete timestamp.';

-- Maintainer notes (after DDL):
-- id = AuthorityProfiles.id; one settings row per profile.
