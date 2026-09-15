CREATE TABLE IF NOT EXISTS "auth"."ForgerProfileSettings" (
  "id" UUID PRIMARY KEY,
  "login_notification" BOOLEAN NOT NULL DEFAULT TRUE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_forger_profile_settings_id" FOREIGN KEY ("id") REFERENCES "auth"."ForgerProfiles" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_forger_profile_settings_deleted_at" ON "auth"."ForgerProfileSettings" ("deleted_at");

COMMENT ON TABLE "auth"."ForgerProfileSettings" IS 'Per-forger-profile system settings; id equals ForgerProfiles.id (1:1).';
COMMENT ON COLUMN "auth"."ForgerProfileSettings"."id" IS 'Primary key; same as auth.ForgerProfiles.id.';
COMMENT ON COLUMN "auth"."ForgerProfileSettings"."login_notification" IS 'When true, send email on successful login.';
COMMENT ON COLUMN "auth"."ForgerProfileSettings"."created_at" IS 'When this settings row was created.';
COMMENT ON COLUMN "auth"."ForgerProfileSettings"."updated_at" IS 'Last change to this row.';
COMMENT ON COLUMN "auth"."ForgerProfileSettings"."deleted_at" IS 'Soft-delete timestamp.';

-- Maintainer notes (after DDL):
-- id = ForgerProfiles.id; one settings row per profile.
