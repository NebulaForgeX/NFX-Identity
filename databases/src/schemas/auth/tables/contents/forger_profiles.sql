CREATE TABLE IF NOT EXISTS "auth"."ForgerProfiles" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "forger_roles" "auth".forger_role[] NOT NULL DEFAULT ARRAY['forger']::"auth".forger_role[],
  "profile_language" "auth".profile_language NOT NULL DEFAULT 'en',
  "preference" JSONB NULL,
  "display_name" VARCHAR(150) NULL,
  "first_name" VARCHAR(100) NULL,
  "last_name" VARCHAR(100) NULL,
  "country" VARCHAR(100) NULL,
  "city" VARCHAR(100) NULL,
  "gender" VARCHAR(100) NULL,
  "birthday" DATE NULL,
  "website" VARCHAR(100) NULL,
  "timezone" VARCHAR(100) NULL,
  "bio" TEXT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_forger_profiles_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON DELETE CASCADE,
  CONSTRAINT "chk_forger_profiles_forger_roles_nonempty" CHECK (cardinality("forger_roles") >= 1)
);

CREATE INDEX IF NOT EXISTS "idx_forger_profiles_account_id" ON "auth"."ForgerProfiles" ("account_id");
CREATE INDEX IF NOT EXISTS "idx_forger_profiles_deleted_at" ON "auth"."ForgerProfiles" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_forger_profiles_profile_language" ON "auth"."ForgerProfiles" ("profile_language");
CREATE INDEX IF NOT EXISTS "idx_forger_profiles_forger_roles" ON "auth"."ForgerProfiles" USING GIN ("forger_roles");

COMMENT ON TABLE "auth"."ForgerProfiles" IS 'End-user forger display profile; one account may own many profiles (1:N). Distinct from AuthorityProfiles. Phone numbers live in auth.Phones.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."id" IS 'Primary key UUID (profile_id); referenced by content/social as profile_id.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."account_id" IS 'Owning account; FK to auth.Accounts; one account may have many profiles; CASCADE delete.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."forger_roles" IS 'auth.forger_role[]; capability = membership. Default {forger}.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."display_name" IS 'Public handle / nickname; optional if first+last used.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."deleted_at" IS 'Soft-delete profile snapshot.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."profile_language" IS 'Preferred UI locale (auth.profile_language).';
COMMENT ON COLUMN "auth"."ForgerProfiles"."preference" IS 'User preferences JSON: theme, layout, etc.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."first_name" IS 'Given name; optional.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."last_name" IS 'Family name; optional.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."country" IS 'Country string for display or filters.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."city" IS 'City string.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."gender" IS 'Self-identified gender label; app-defined vocabulary.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."birthday" IS 'Birth date if collected.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."website" IS 'Personal or social URL.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."timezone" IS 'IANA timezone name if set.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."bio" IS 'Free-text bio.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."created_at" IS 'Row insert time.';
COMMENT ON COLUMN "auth"."ForgerProfiles"."updated_at" IS 'Last profile edit.';

-- Maintainer notes (after DDL):
-- File forger_profiles.sql → table ForgerProfiles; one account : many profiles via account_id FK.
-- Roles are an array: check membership with HasRole(need, roles) / SQL @> or = ANY.
