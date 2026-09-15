CREATE TABLE IF NOT EXISTS "auth"."AuthorityProfiles" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "authority_roles" "auth".authority_role[] NOT NULL DEFAULT ARRAY['auditor']::"auth".authority_role[],
  "profile_language" "auth".profile_language NOT NULL DEFAULT 'zh',
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
  CONSTRAINT "fk_authority_profiles_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON DELETE CASCADE,
  CONSTRAINT "chk_authority_profiles_authority_roles_nonempty" CHECK (cardinality("authority_roles") >= 1)
);

CREATE INDEX IF NOT EXISTS "idx_authority_profiles_account_id" ON "auth"."AuthorityProfiles" ("account_id");
CREATE INDEX IF NOT EXISTS "idx_authority_profiles_deleted_at" ON "auth"."AuthorityProfiles" ("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_authority_profiles_profile_language" ON "auth"."AuthorityProfiles" ("profile_language");
CREATE INDEX IF NOT EXISTS "idx_authority_profiles_authority_roles" ON "auth"."AuthorityProfiles" USING GIN ("authority_roles");

COMMENT ON TABLE "auth"."AuthorityProfiles" IS 'Staff/authority display profile; one account may own many profiles (1:N). Distinct from ForgerProfiles; carries authority_roles[]. Phone numbers live in auth.Phones.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."id" IS 'Primary key UUID (profile_id); referenced by content/social as profile_id.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."account_id" IS 'Owning account; FK to auth.Accounts; one account may have many profiles; CASCADE delete.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."authority_roles" IS 'auth.authority_role[]; capability = membership only (reviewer, moderator, …). Default {reviewer}.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."display_name" IS 'Public handle / nickname; optional if first+last used.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."deleted_at" IS 'Soft-delete profile snapshot.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."profile_language" IS 'Preferred UI locale (auth.profile_language).';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."preference" IS 'User preferences JSON: theme, layout, etc.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."first_name" IS 'Given name; optional.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."last_name" IS 'Family name; optional.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."country" IS 'Country string for display or filters.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."city" IS 'City string.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."gender" IS 'Self-identified gender label; app-defined vocabulary.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."birthday" IS 'Birth date if collected.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."website" IS 'Personal or social URL.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."timezone" IS 'IANA timezone name if set.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."bio" IS 'Free-text bio.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."created_at" IS 'Row insert time.';
COMMENT ON COLUMN "auth"."AuthorityProfiles"."updated_at" IS 'Last profile edit.';

-- Maintainer notes (after DDL):
-- File authority_profiles.sql → table AuthorityProfiles; one account : many profiles via account_id FK.
-- One role = one purpose; check with HasRole(need, roles) / SQL @> or = ANY. No hierarchical bundling.
