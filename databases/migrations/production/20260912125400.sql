-- Create NFX-Identity auth+asset+system schemas from databases/src

-- imported from extensions/pgcrypto.sql
CREATE EXTENSION IF NOT EXISTS "pgcrypto"
WITH
  SCHEMA "public" VERSION "1.3";


-- imported from extensions/btree_gist.sql
-- Active: 1768427124487@@192.168.1.64@10105@postgres
CREATE EXTENSION IF NOT EXISTS "btree_gist"
WITH
  SCHEMA "public";

-- imported from schemas/auth/main.sql

-- imported from schemas/auth/schema.sql
CREATE SCHEMA IF NOT EXISTS "auth";
COMMENT ON SCHEMA "auth" IS 'Login center: accounts, identities, forger/authority profiles. Avatars/backgrounds link to asset.Images.';

-- imported from schemas/auth/enums/identity_provider.sql
CREATE TYPE "auth".identity_provider AS ENUM (
  'password',
  'github'
);
COMMENT ON TYPE "auth".identity_provider IS 'password=email/phone credential; github=OAuth.';

-- imported from schemas/auth/enums/forger_role.sql
CREATE TYPE "auth".forger_role AS ENUM (
  'forger'
);
COMMENT ON TYPE "auth".forger_role IS 'Forger profile roles; membership only, not hierarchical.';

-- imported from schemas/auth/enums/authority_role.sql
CREATE TYPE "auth".authority_role AS ENUM (
  'auditor',
  'administrator',
  'owner'
);
COMMENT ON TYPE "auth".authority_role IS 'Authority profile roles; membership only, not hierarchical.';

-- imported from schemas/auth/enums/account_status.sql
CREATE TYPE "auth".account_status AS ENUM (
  'active',
  'suspended',
  'deleted'
);
COMMENT ON TYPE "auth".account_status IS 'Account lifecycle.';

-- imported from schemas/auth/enums/signup_platform.sql
CREATE TYPE "auth".signup_platform AS ENUM (
  'nfxidentity',
  'nfxnews',
  'nfxstorages',
  'nfxvault'
);
COMMENT ON TYPE "auth".signup_platform IS 'Product that first created the account; written only at signup.';

-- imported from schemas/auth/enums/profile_language.sql
CREATE TYPE "auth".profile_language AS ENUM (
  'en',
  'zh',
  'fr'
);
COMMENT ON TYPE "auth".profile_language IS 'Preferred UI locale.';

-- imported from schemas/auth/enums/profile_scope.sql
CREATE TYPE "auth".profile_scope AS ENUM (
  'forger',
  'authority'
);
COMMENT ON TYPE "auth".profile_scope IS 'JWT profile_scope / login kind.';

-- imported from schemas/auth/tables/contents/accounts.sql
CREATE TABLE IF NOT EXISTS "auth"."Accounts" (
  "id" UUID PRIMARY KEY,
  "account_status" "auth".account_status NOT NULL DEFAULT 'active',
  "signup_platform" "auth".signup_platform NOT NULL DEFAULT 'nfxidentity',
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL
);

CREATE INDEX IF NOT EXISTS "idx_accounts_account_status" ON "auth"."Accounts" ("account_status");
CREATE INDEX IF NOT EXISTS "idx_accounts_signup_platform" ON "auth"."Accounts" ("signup_platform");
CREATE INDEX IF NOT EXISTS "idx_accounts_deleted_at" ON "auth"."Accounts" ("deleted_at");

COMMENT ON TABLE "auth"."Accounts" IS 'Core account; credentials in Identities, community profile in CommunityProfiles, authority profile in AuthorityProfiles, emails in Emails.';
COMMENT ON COLUMN "auth"."Accounts"."id" IS 'Primary key UUID; typically matches application user id.';
COMMENT ON COLUMN "auth"."Accounts"."account_status" IS 'Lifecycle: auth.account_status enum.';
COMMENT ON COLUMN "auth"."Accounts"."signup_platform" IS 'Product that first created this account (auth.signup_platform); written only at signup.';
COMMENT ON COLUMN "auth"."Accounts"."created_at" IS 'Row insert time.';
COMMENT ON COLUMN "auth"."Accounts"."updated_at" IS 'Last account row update.';
COMMENT ON COLUMN "auth"."Accounts"."deleted_at" IS 'Optional soft-delete timestamp when account_status = deleted.';

-- Maintainer notes (after DDL):
-- 账号核心表：不含邮箱/密码（见 Emails / Identities）；角色下沉到各 profile（CommunityProfiles.community_roles[] / AuthorityProfiles.authority_roles[]）。
-- signup_platform：仅注册创建时写入；DEFAULT nfxidentity。

-- imported from schemas/auth/tables/contents/emails.sql
CREATE TABLE IF NOT EXISTS "auth"."Emails" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "is_primary" BOOLEAN NOT NULL DEFAULT FALSE,
  "verified_at" TIMESTAMP NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_emails_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "uq_emails_email_ci" ON "auth"."Emails" (LOWER("email")) WHERE "deleted_at" IS NULL;

CREATE INDEX IF NOT EXISTS "idx_emails_account_id" ON "auth"."Emails" ("account_id");
CREATE INDEX IF NOT EXISTS "idx_emails_deleted_at" ON "auth"."Emails" ("deleted_at");
CREATE UNIQUE INDEX IF NOT EXISTS "uq_emails_one_primary_per_account" ON "auth"."Emails" ("account_id") WHERE "is_primary" = TRUE AND "deleted_at" IS NULL;

COMMENT ON TABLE "auth"."Emails" IS 'Email bindings; soft-deleted rows kept for history. Partial unique (LOWER(email)) WHERE deleted_at IS NULL avoids collision when re-registering. Avoid hard-deleting auth.Accounts if you need email audit trail (CASCADE removes child rows).';
COMMENT ON COLUMN "auth"."Emails"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."Emails"."account_id" IS 'Owning auth.Accounts row; CASCADE delete.';
COMMENT ON COLUMN "auth"."Emails"."email" IS 'Email address; uniqueness enforced case-insensitively among non-deleted rows.';
COMMENT ON COLUMN "auth"."Emails"."verified_at" IS 'NULL = not verified yet.';
COMMENT ON COLUMN "auth"."Emails"."is_primary" IS 'At most one non-deleted primary email per account (partial unique index).';
COMMENT ON COLUMN "auth"."Emails"."created_at" IS 'Row insert time.';
COMMENT ON COLUMN "auth"."Emails"."updated_at" IS 'Last edit time.';
COMMENT ON COLUMN "auth"."Emails"."deleted_at" IS 'Soft-delete only: row remains for audit. Partial unique applies only where deleted_at IS NULL. Re-register adds another row.';

-- Maintainer notes (after DDL):
-- 登录邮箱（可多邮箱；一条标记为主邮箱）。
-- 同一 LOWER(email) 可存在多行：多行「已注销」+ 至多一行「当前有效」。
--   注销：本行写 deleted_at，不删行。再注册同邮箱：INSERT 新行（新 id、新 account_id、deleted_at NULL），旧行 deleted_at 不动 → 每次注销时间仍留在旧行上。
-- 部分唯一索引 (LOWER(email) WHERE deleted_at IS NULL) 只保证：同时只能有一条「未删除」的该邮箱；不限制历史行条数。
-- 若对 Accounts 物理 DELETE，CASCADE 会删掉子表 Emails 行 → 历史消失；用户侧应软删账号，勿物理删。

-- imported from schemas/auth/tables/contents/phones.sql
CREATE TABLE IF NOT EXISTS "auth"."Phones" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "phone" VARCHAR(32) NOT NULL,
  "is_primary" BOOLEAN NOT NULL DEFAULT FALSE,
  "verified_at" TIMESTAMP NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_phones_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "uq_phones_phone_active" ON "auth"."Phones" ("phone") WHERE "deleted_at" IS NULL;

CREATE INDEX IF NOT EXISTS "idx_phones_account_id" ON "auth"."Phones" ("account_id");
CREATE INDEX IF NOT EXISTS "idx_phones_deleted_at" ON "auth"."Phones" ("deleted_at");
CREATE UNIQUE INDEX IF NOT EXISTS "uq_phones_one_primary_per_account" ON "auth"."Phones" ("account_id") WHERE "is_primary" = TRUE AND "deleted_at" IS NULL;

COMMENT ON COLUMN "auth"."Phones"."verified_at" IS 'NULL = not verified yet.';
COMMENT ON COLUMN "auth"."Phones"."is_primary" IS 'At most one non-deleted primary phone per account.';
COMMENT ON COLUMN "auth"."Phones"."deleted_at" IS 'Soft-delete row; same canonical phone may be re-bound via INSERT new row.';
COMMENT ON TABLE "auth"."Phones" IS 'Phone bindings; partial unique on phone WHERE deleted_at IS NULL mirrors Emails pattern.';
COMMENT ON COLUMN "auth"."Phones"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."Phones"."account_id" IS 'Owning auth.Accounts row; CASCADE delete.';
COMMENT ON COLUMN "auth"."Phones"."phone" IS 'Canonical phone string; unique among non-deleted rows.';
COMMENT ON COLUMN "auth"."Phones"."created_at" IS 'Row insert time.';
COMMENT ON COLUMN "auth"."Phones"."updated_at" IS 'Last edit time.';

-- Maintainer notes (after DDL):
-- 手机号（可多号码；一条标记为主号码）。语义对齐 Emails：软删历史 + 再绑定同号用新行。
-- 建议在应用层存 E.164 等规范格式，避免同一号码多种写法绕过唯一。

-- imported from schemas/auth/tables/contents/forger_profiles.sql
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

-- imported from schemas/auth/tables/contents/authority_profiles.sql
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

COMMENT ON TABLE "auth"."AuthorityProfiles" IS 'Staff/authority display profile; one account may own many profiles (1:N). Mirrors CommunityProfiles but carries authority_roles[]. Phone numbers live in auth.Phones.';
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

-- imported from schemas/auth/tables/contents/identities.sql
CREATE TABLE IF NOT EXISTS "auth"."Identities" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "identity_provider" "auth".identity_provider NOT NULL,
  "provider_subject" TEXT NOT NULL,
  "password_hash" TEXT NULL,
  "metadata" JSONB NULL,
  "last_login_at" TIMESTAMP NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_identities_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_identities_account_id" ON "auth"."Identities" ("account_id");
CREATE INDEX IF NOT EXISTS "idx_identities_deleted_at" ON "auth"."Identities" ("deleted_at");
CREATE UNIQUE INDEX IF NOT EXISTS "uq_identities_identity_provider_subject_active" ON "auth"."Identities" ("identity_provider", "provider_subject") WHERE "deleted_at" IS NULL;

COMMENT ON TABLE "auth"."Identities" IS 'Auth credentials per identity_provider; PASSWORD uses provider_subject = normalized login id (e.g. email).';
COMMENT ON COLUMN "auth"."Identities"."identity_provider" IS 'auth.identity_provider enum (password / future OAuth).';
COMMENT ON COLUMN "auth"."Identities"."provider_subject" IS 'Unique within identity_provider among non-deleted rows; email for password, sub for OAuth.';
COMMENT ON COLUMN "auth"."Identities"."password_hash" IS 'Set only when identity_provider = password; bcrypt/argon2 hash at application layer.';
COMMENT ON COLUMN "auth"."Identities"."metadata" IS 'Optional provider-specific payload (e.g. OAuth raw claims subset).';
COMMENT ON COLUMN "auth"."Identities"."deleted_at" IS 'Soft-delete identity; frees provider_subject for re-link.';
COMMENT ON COLUMN "auth"."Identities"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."Identities"."account_id" IS 'Owning auth.Accounts row; CASCADE delete.';
COMMENT ON COLUMN "auth"."Identities"."last_login_at" IS 'Last successful login instant for this identity.';
COMMENT ON COLUMN "auth"."Identities"."created_at" IS 'Row insert time.';
COMMENT ON COLUMN "auth"."Identities"."updated_at" IS 'Last credential/metadata update.';

-- Maintainer notes (after DDL):
-- 登录身份：密码 / 未来 OAuth；凭证不与 Accounts 混放在一行。

-- imported from schemas/auth/tables/contents/refresh_tokens.sql
CREATE TABLE IF NOT EXISTS "auth"."RefreshTokens" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "identity_id" UUID NULL,
  "profile_id" UUID NULL,
  "profile_scope" "auth".profile_scope NULL,
  "device_id" TEXT NULL,
  "token_hash" TEXT NOT NULL,
  "expires_at" TIMESTAMP NOT NULL,
  "revoked_at" TIMESTAMP NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_refresh_tokens_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_refresh_tokens_identity_id" FOREIGN KEY ("identity_id") REFERENCES "auth"."Identities" ("id") ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_account_id" ON "auth"."RefreshTokens" ("account_id");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_identity_id" ON "auth"."RefreshTokens" ("identity_id");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_device_id" ON "auth"."RefreshTokens" ("device_id");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_account_device_active"
  ON "auth"."RefreshTokens" ("account_id", "device_id")
  WHERE "device_id" IS NOT NULL AND "revoked_at" IS NULL AND "deleted_at" IS NULL;
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_expires_at" ON "auth"."RefreshTokens" ("expires_at");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_revoked_at" ON "auth"."RefreshTokens" ("revoked_at");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_deleted_at" ON "auth"."RefreshTokens" ("deleted_at");

COMMENT ON TABLE "auth"."RefreshTokens" IS 'Opaque refresh tokens; store only hashes at rest.';
COMMENT ON COLUMN "auth"."RefreshTokens"."identity_id" IS 'Which login issued the token; NULL if legacy / unspecified.';
COMMENT ON COLUMN "auth"."RefreshTokens"."device_id" IS 'Client device/session identifier for same-device token revocation policy.';
COMMENT ON COLUMN "auth"."RefreshTokens"."deleted_at" IS 'Soft-delete token row (e.g. purge); prefer revoked_at for normal logout.';
COMMENT ON COLUMN "auth"."RefreshTokens"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."RefreshTokens"."account_id" IS 'Owning account; CASCADE delete.';
COMMENT ON COLUMN "auth"."RefreshTokens"."token_hash" IS 'Hash of opaque refresh token at rest.';
COMMENT ON COLUMN "auth"."RefreshTokens"."expires_at" IS 'Absolute expiry instant.';
COMMENT ON COLUMN "auth"."RefreshTokens"."revoked_at" IS 'When token was invalidated server-side.';
COMMENT ON COLUMN "auth"."RefreshTokens"."created_at" IS 'Issued-at time.';

-- Maintainer notes (after DDL):
-- 刷新会话令牌（存哈希，不存明文）。
-- device_id 用于“同设备 revoke，其他设备保留”。

-- imported from schemas/auth/tables/links/forger_profile_avatars.sql
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

-- imported from schemas/auth/tables/links/forger_profile_backgrounds.sql
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

-- imported from schemas/auth/tables/links/forger_profile_settings.sql
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

-- imported from schemas/auth/tables/links/authority_profile_avatars.sql
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

-- imported from schemas/auth/tables/links/authority_profile_backgrounds.sql
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

-- imported from schemas/auth/tables/links/authority_profile_settings.sql
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

-- imported from schemas/auth/views/full_account_information.sql
CREATE OR REPLACE VIEW "auth"."FullAccountInformationAccountView" AS
SELECT
  a."id" AS "account_id",
  a."account_status" AS "account_status",
  a."signup_platform" AS "signup_platform",
  a."created_at" AS "created_at",
  a."updated_at" AS "updated_at"
FROM "auth"."Accounts" a
WHERE a."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationAccountView" IS 'Active account base fields only; one row per account.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationEmailView" AS
SELECT
  em."id" AS "id",
  em."account_id" AS "account_id",
  em."email" AS "email",
  em."is_primary" AS "is_primary",
  em."verified_at" AS "verified_at",
  em."created_at" AS "created_at",
  em."updated_at" AS "updated_at"
FROM "auth"."Emails" em
WHERE em."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationEmailView" IS 'Non-deleted emails; query by account_id for account details.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationPhoneView" AS
SELECT
  ph."id" AS "id",
  ph."account_id" AS "account_id",
  ph."phone" AS "phone",
  ph."is_primary" AS "is_primary",
  ph."verified_at" AS "verified_at",
  ph."created_at" AS "created_at",
  ph."updated_at" AS "updated_at"
FROM "auth"."Phones" ph
WHERE ph."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationPhoneView" IS 'Non-deleted phones; query by account_id for account details.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationForgerProfileView" AS
SELECT
  p."id" AS "id",
  p."account_id" AS "account_id",
  p."forger_roles" AS "forger_roles",
  p."profile_language" AS "profile_language",
  p."preference" AS "preference",
  p."display_name" AS "display_name",
  p."first_name" AS "first_name",
  p."last_name" AS "last_name",
  p."country" AS "country",
  p."city" AS "city",
  p."gender" AS "gender",
  p."birthday" AS "birthday",
  p."website" AS "website",
  p."timezone" AS "timezone",
  p."bio" AS "bio",
  p."created_at" AS "created_at",
  p."updated_at" AS "updated_at"
FROM "auth"."ForgerProfiles" p
WHERE p."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationForgerProfileView" IS 'Non-deleted forger profiles; one account may have many rows; query by account_id (list) or id (single profile).';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationAvatarView" AS
SELECT
  av."id" AS "id",
  av."profile_id" AS "profile_id",
  av."image_id" AS "image_id",
  av."is_active" AS "is_active",
  av."created_at" AS "created_at",
  av."updated_at" AS "updated_at"
FROM "auth"."ForgerProfileAvatars" av
WHERE av."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationAvatarView" IS 'Non-deleted forger avatar links; query by profile_id.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationBackgroundView" AS
SELECT
  bg."id" AS "id",
  bg."profile_id" AS "profile_id",
  bg."image_id" AS "image_id",
  bg."sort_order" AS "sort_order",
  bg."created_at" AS "created_at",
  bg."updated_at" AS "updated_at"
FROM "auth"."ForgerProfileBackgrounds" bg
WHERE bg."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationBackgroundView" IS 'Non-deleted forger profile background links; query by profile_id.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationForgerProfileSettingsView" AS
SELECT
  s."id" AS "id",
  s."login_notification" AS "login_notification",
  s."created_at" AS "created_at",
  s."updated_at" AS "updated_at"
FROM "auth"."ForgerProfileSettings" s
WHERE s."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationForgerProfileSettingsView" IS 'Non-deleted forger profile settings; query by id (profile_id).';

-- Maintainer notes:
-- FullAccountInformation 读模型（账号聚合，含账号下全部 forger profile）：FullAccountInformationVO
--   { Account, Emails[], Phones[], ForgerProfiles []ForgerProfileVO }。
-- 拆成多个 view，便于 Go service 按 account_id / id 精准查询并组装：
--   Account + Email + Phone + ForgerProfile(多 profile) + Avatar/Background/Settings(按 profile_id)。

-- imported from schemas/auth/views/full_account_information_with_forger_profile.sql
CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithForgerProfileAccountView" AS
SELECT
  a."id" AS "account_id",
  a."account_status" AS "account_status",
  a."signup_platform" AS "signup_platform",
  a."created_at" AS "created_at",
  a."updated_at" AS "updated_at"
FROM "auth"."Accounts" a
WHERE a."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileAccountView" IS 'Active account base fields only; one row per account. Backs the single-forger-profile login-state read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithForgerProfileEmailView" AS
SELECT
  em."id" AS "id",
  em."account_id" AS "account_id",
  em."email" AS "email",
  em."is_primary" AS "is_primary",
  em."verified_at" AS "verified_at",
  em."created_at" AS "created_at",
  em."updated_at" AS "updated_at"
FROM "auth"."Emails" em
WHERE em."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileEmailView" IS 'Non-deleted emails; query by account_id for the single-forger-profile read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithForgerProfilePhoneView" AS
SELECT
  ph."id" AS "id",
  ph."account_id" AS "account_id",
  ph."phone" AS "phone",
  ph."is_primary" AS "is_primary",
  ph."verified_at" AS "verified_at",
  ph."created_at" AS "created_at",
  ph."updated_at" AS "updated_at"
FROM "auth"."Phones" ph
WHERE ph."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfilePhoneView" IS 'Non-deleted phones; query by account_id for the single-forger-profile read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithForgerProfileView" AS
SELECT
  p."id" AS "id",
  p."account_id" AS "account_id",
  p."forger_roles" AS "forger_roles",
  p."profile_language" AS "profile_language",
  p."preference" AS "preference",
  p."display_name" AS "display_name",
  p."first_name" AS "first_name",
  p."last_name" AS "last_name",
  p."country" AS "country",
  p."city" AS "city",
  p."gender" AS "gender",
  p."birthday" AS "birthday",
  p."website" AS "website",
  p."timezone" AS "timezone",
  p."bio" AS "bio",
  p."created_at" AS "created_at",
  p."updated_at" AS "updated_at"
FROM "auth"."ForgerProfiles" p
WHERE p."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileView" IS 'Non-deleted forger profiles; query by id (current profile from token) or account_id (fallback). Backs the single-forger-profile read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithForgerProfileAvatarView" AS
SELECT
  av."id" AS "id",
  av."profile_id" AS "profile_id",
  av."image_id" AS "image_id",
  av."is_active" AS "is_active",
  av."created_at" AS "created_at",
  av."updated_at" AS "updated_at"
FROM "auth"."ForgerProfileAvatars" av
WHERE av."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileAvatarView" IS 'Non-deleted avatar links; query by profile_id for the single-forger-profile read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithForgerProfileBackgroundView" AS
SELECT
  bg."id" AS "id",
  bg."profile_id" AS "profile_id",
  bg."image_id" AS "image_id",
  bg."sort_order" AS "sort_order",
  bg."created_at" AS "created_at",
  bg."updated_at" AS "updated_at"
FROM "auth"."ForgerProfileBackgrounds" bg
WHERE bg."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileBackgroundView" IS 'Non-deleted profile background links; query by profile_id for the single-forger-profile read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithForgerProfileSettingsView" AS
SELECT
  s."id" AS "id",
  s."login_notification" AS "login_notification",
  s."created_at" AS "created_at",
  s."updated_at" AS "updated_at"
FROM "auth"."ForgerProfileSettings" s
WHERE s."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithForgerProfileSettingsView" IS 'Non-deleted profile settings; query by id (profile_id) for the single-forger-profile read model.';

-- Maintainer notes:
-- FullAccountInformationWithForgerProfile 读模型（当前登录态，单 forger profile）：FullAccountInformationWithForgerProfileVO
--   { AccountID, Account, Emails[], Phones[], ForgerProfile *ForgerProfileVO }（由 token 的 profile_id 决定单个 profile）。
-- 与 full_account_information.sql 结构对称、独立命名：
--   Account + Email + Phone(按 account_id) + ForgerProfile(按 token profile_id 取单个) + Avatar/Background/Settings(按该 profile_id)。

-- imported from schemas/auth/views/full_account_information_with_authority_profile.sql
CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithAuthorityProfileAccountView" AS
SELECT
  a."id" AS "account_id",
  a."account_status" AS "account_status",
  a."signup_platform" AS "signup_platform",
  a."created_at" AS "created_at",
  a."updated_at" AS "updated_at"
FROM "auth"."Accounts" a
WHERE a."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileAccountView" IS 'Active account base fields only; one row per account. Backs the single-authority-profile login-state read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithAuthorityProfileEmailView" AS
SELECT
  em."id" AS "id",
  em."account_id" AS "account_id",
  em."email" AS "email",
  em."is_primary" AS "is_primary",
  em."verified_at" AS "verified_at",
  em."created_at" AS "created_at",
  em."updated_at" AS "updated_at"
FROM "auth"."Emails" em
WHERE em."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileEmailView" IS 'Non-deleted emails; query by account_id for the single-authority-profile read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithAuthorityProfilePhoneView" AS
SELECT
  ph."id" AS "id",
  ph."account_id" AS "account_id",
  ph."phone" AS "phone",
  ph."is_primary" AS "is_primary",
  ph."verified_at" AS "verified_at",
  ph."created_at" AS "created_at",
  ph."updated_at" AS "updated_at"
FROM "auth"."Phones" ph
WHERE ph."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfilePhoneView" IS 'Non-deleted phones; query by account_id for the single-authority-profile read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithAuthorityProfileView" AS
SELECT
  p."id" AS "id",
  p."account_id" AS "account_id",
  p."authority_roles" AS "authority_roles",
  p."profile_language" AS "profile_language",
  p."preference" AS "preference",
  p."display_name" AS "display_name",
  p."first_name" AS "first_name",
  p."last_name" AS "last_name",
  p."country" AS "country",
  p."city" AS "city",
  p."gender" AS "gender",
  p."birthday" AS "birthday",
  p."website" AS "website",
  p."timezone" AS "timezone",
  p."bio" AS "bio",
  p."created_at" AS "created_at",
  p."updated_at" AS "updated_at"
FROM "auth"."AuthorityProfiles" p
WHERE p."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileView" IS 'Non-deleted authority profiles; query by id (current profile from token) or account_id (fallback). Backs the single-authority-profile read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithAuthorityProfileAvatarView" AS
SELECT
  av."id" AS "id",
  av."profile_id" AS "profile_id",
  av."image_id" AS "image_id",
  av."is_active" AS "is_active",
  av."created_at" AS "created_at",
  av."updated_at" AS "updated_at"
FROM "auth"."AuthorityProfileAvatars" av
WHERE av."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileAvatarView" IS 'Non-deleted avatar links; query by profile_id for the single-authority-profile read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithAuthorityProfileBackgroundView" AS
SELECT
  bg."id" AS "id",
  bg."profile_id" AS "profile_id",
  bg."image_id" AS "image_id",
  bg."sort_order" AS "sort_order",
  bg."created_at" AS "created_at",
  bg."updated_at" AS "updated_at"
FROM "auth"."AuthorityProfileBackgrounds" bg
WHERE bg."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileBackgroundView" IS 'Non-deleted profile background links; query by profile_id for the single-authority-profile read model.';

CREATE OR REPLACE VIEW "auth"."FullAccountInformationWithAuthorityProfileSettingsView" AS
SELECT
  s."id" AS "id",
  s."login_notification" AS "login_notification",
  s."created_at" AS "created_at",
  s."updated_at" AS "updated_at"
FROM "auth"."AuthorityProfileSettings" s
WHERE s."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."FullAccountInformationWithAuthorityProfileSettingsView" IS 'Non-deleted profile settings; query by id (profile_id) for the single-authority-profile read model.';

-- Maintainer notes:
-- FullAccountInformationWithAuthorityProfile 读模型（当前登录态，单 authority profile）：FullAccountInformationWithAuthorityProfileVO
--   { AccountID, Account, Emails[], Phones[], AuthorityProfile *AuthorityProfileVO }（由 token 的 profile_id 决定单个 profile）。
-- 镜像 full_account_information_with_community_profile.sql，但 profile 子视图取 AuthorityProfiles 并带 authority_roles[]。

-- imported from schemas/auth/views/list_forger_profile_item.sql
CREATE OR REPLACE VIEW "auth"."ListForgerProfileItems" AS
SELECT
  p."id" AS "profile_id",
  p."account_id" AS "account_id",
  p."forger_roles" AS "forger_roles",
  p."display_name" AS "display_name",
  p."profile_language" AS "profile_language",
  p."city" AS "city",
  p."country" AS "country",
  p."website" AS "website",
  p."timezone" AS "timezone",
  p."birthday" AS "birthday",
  p."created_at" AS "created_at",
  av."image_id" AS "avatar_image_id",
  bg."image_id" AS "background_image_id"
FROM "auth"."ForgerProfiles" p
LEFT JOIN LATERAL (
  SELECT a."image_id"
  FROM "auth"."ForgerProfileAvatars" a
  WHERE
    a."profile_id" = p."id"
    AND a."is_active" = TRUE
    AND a."deleted_at" IS NULL
  ORDER BY a."created_at" DESC
  LIMIT 1
) av ON TRUE
LEFT JOIN LATERAL (
  SELECT b."image_id"
  FROM "auth"."ForgerProfileBackgrounds" b
  WHERE
    b."profile_id" = p."id"
    AND b."deleted_at" IS NULL
  ORDER BY b."sort_order" ASC, b."created_at" ASC
  LIMIT 1
) bg ON TRUE
WHERE p."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."ListForgerProfileItems" IS
'One row per non-deleted forger profile with public list/search fields, active avatar, and primary background (lowest sort_order). Private fields (first_name, last_name) are not exposed; read those from account-scoped profile APIs only.';

-- imported from schemas/auth/views/list_authority_profile_item.sql
CREATE OR REPLACE VIEW "auth"."ListAuthorityProfileItems" AS
SELECT
  p."id" AS "profile_id",
  p."account_id" AS "account_id",
  p."authority_roles" AS "authority_roles",
  p."display_name" AS "display_name",
  p."profile_language" AS "profile_language",
  p."city" AS "city",
  p."country" AS "country",
  p."website" AS "website",
  p."timezone" AS "timezone",
  p."birthday" AS "birthday",
  p."created_at" AS "created_at",
  av."image_id" AS "avatar_image_id",
  bg."image_id" AS "background_image_id"
FROM "auth"."AuthorityProfiles" p
LEFT JOIN LATERAL (
  SELECT a."image_id"
  FROM "auth"."AuthorityProfileAvatars" a
  WHERE
    a."profile_id" = p."id"
    AND a."is_active" = TRUE
    AND a."deleted_at" IS NULL
  ORDER BY a."created_at" DESC
  LIMIT 1
) av ON TRUE
LEFT JOIN LATERAL (
  SELECT b."image_id"
  FROM "auth"."AuthorityProfileBackgrounds" b
  WHERE
    b."profile_id" = p."id"
    AND b."deleted_at" IS NULL
  ORDER BY b."sort_order" ASC, b."created_at" ASC
  LIMIT 1
) bg ON TRUE
WHERE p."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."ListAuthorityProfileItems" IS
'One row per non-deleted authority profile with public list/search fields, active avatar, and primary background (lowest sort_order). Mirrors ListCommunityProfileItems but carries authority_roles[].';

-- imported from schemas/auth/views/list_email_item.sql
CREATE OR REPLACE VIEW "auth"."ListEmailItems" AS
SELECT
  em."id" AS "id",
  em."account_id" AS "account_id",
  em."email" AS "email",
  em."is_primary" AS "is_primary",
  em."verified_at" AS "verified_at",
  em."created_at" AS "created_at",
  em."updated_at" AS "updated_at"
FROM "auth"."Emails" em
WHERE em."deleted_at" IS NULL;

COMMENT ON VIEW "auth"."ListEmailItems" IS
'One row per non-deleted email. Filter by account_id to list account emails.';

-- imported from schemas/asset/main.sql

-- imported from schemas/asset/schema.sql
CREATE SCHEMA IF NOT EXISTS "asset";
COMMENT ON SCHEMA "asset" IS 'Independent Asset service schema: Images, Files, Videos, Audios metadata. uploader_id is logical ref to Auth (no cross-schema FK). Stored in Stack MinIO.';

-- imported from schemas/asset/tables/images.sql
CREATE TABLE IF NOT EXISTS "asset"."Images" (
  "id" UUID PRIMARY KEY,
  "file_path" VARCHAR(500) NOT NULL,
  "file_name" VARCHAR(255) NOT NULL,
  "file_size" BIGINT NOT NULL,
  "mime_type" VARCHAR(100) NOT NULL,
  "width" INTEGER,
  "height" INTEGER,
  "alt_text" VARCHAR(255),
  "uploader_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
CREATE INDEX IF NOT EXISTS "idx_images_file_path" ON "asset"."Images"("file_path");
CREATE INDEX IF NOT EXISTS "idx_images_uploader_id" ON "asset"."Images"("uploader_id");
CREATE INDEX IF NOT EXISTS "idx_images_mime_type" ON "asset"."Images"("mime_type");
CREATE INDEX IF NOT EXISTS "idx_images_deleted_at" ON "asset"."Images"("deleted_at");

COMMENT ON TABLE "asset"."Images" IS 'Raster image metadata; path is relative to storage root or object key; uploader_id references auth.Accounts at app level.';
COMMENT ON COLUMN "asset"."Images"."id" IS 'Primary key UUID; typically matches stored object id.';
COMMENT ON COLUMN "asset"."Images"."file_path" IS 'Relative path from storage root or object key.';
COMMENT ON COLUMN "asset"."Images"."file_name" IS 'Original filename.';
COMMENT ON COLUMN "asset"."Images"."file_size" IS 'File size in bytes.';
COMMENT ON COLUMN "asset"."Images"."mime_type" IS 'MIME type as uploaded or detected.';
COMMENT ON COLUMN "asset"."Images"."width" IS 'Pixel width if known.';
COMMENT ON COLUMN "asset"."Images"."height" IS 'Pixel height if known.';
COMMENT ON COLUMN "asset"."Images"."alt_text" IS 'Accessibility / caption text.';
COMMENT ON COLUMN "asset"."Images"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
COMMENT ON COLUMN "asset"."Images"."created_at" IS 'Insert time.';
COMMENT ON COLUMN "asset"."Images"."updated_at" IS 'Last metadata update.';
COMMENT ON COLUMN "asset"."Images"."deleted_at" IS 'Soft-delete; NULL if object still active.';

-- imported from schemas/asset/tables/files.sql
CREATE TABLE IF NOT EXISTS "asset"."Files" (
  "id" UUID PRIMARY KEY,
  "file_path" VARCHAR(500) NOT NULL,
  "file_name" VARCHAR(255) NOT NULL,
  "file_size" BIGINT NOT NULL,
  "mime_type" VARCHAR(100) NOT NULL,
  "uploader_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
CREATE INDEX IF NOT EXISTS "idx_files_file_path" ON "asset"."Files"("file_path");
CREATE INDEX IF NOT EXISTS "idx_files_uploader_id" ON "asset"."Files"("uploader_id");
CREATE INDEX IF NOT EXISTS "idx_files_mime_type" ON "asset"."Files"("mime_type");
CREATE INDEX IF NOT EXISTS "idx_files_deleted_at" ON "asset"."Files"("deleted_at");

COMMENT ON TABLE "asset"."Files" IS 'Non-image binary metadata (documents, archives, etc.); same soft-delete shape as Images without pixel fields.';
COMMENT ON COLUMN "asset"."Files"."id" IS 'Primary key UUID; typically matches stored object id.';
COMMENT ON COLUMN "asset"."Files"."file_path" IS 'Relative path from storage root or object key.';
COMMENT ON COLUMN "asset"."Files"."file_name" IS 'Original filename.';
COMMENT ON COLUMN "asset"."Files"."file_size" IS 'File size in bytes.';
COMMENT ON COLUMN "asset"."Files"."mime_type" IS 'MIME type as uploaded or detected.';
COMMENT ON COLUMN "asset"."Files"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
COMMENT ON COLUMN "asset"."Files"."created_at" IS 'Insert time.';
COMMENT ON COLUMN "asset"."Files"."updated_at" IS 'Last metadata update.';
COMMENT ON COLUMN "asset"."Files"."deleted_at" IS 'Soft-delete; NULL if object still active.';

-- imported from schemas/asset/tables/videos.sql
CREATE TABLE IF NOT EXISTS "asset"."Videos" (
  "id" UUID PRIMARY KEY,
  "file_path" VARCHAR(500) NOT NULL,
  "file_name" VARCHAR(255) NOT NULL,
  "file_size" BIGINT NOT NULL,
  "mime_type" VARCHAR(100) NOT NULL,
  "duration_seconds" DOUBLE PRECISION,
  "width" INTEGER,
  "height" INTEGER,
  "uploader_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
CREATE INDEX IF NOT EXISTS "idx_videos_file_path" ON "asset"."Videos"("file_path");
CREATE INDEX IF NOT EXISTS "idx_videos_uploader_id" ON "asset"."Videos"("uploader_id");
CREATE INDEX IF NOT EXISTS "idx_videos_mime_type" ON "asset"."Videos"("mime_type");
CREATE INDEX IF NOT EXISTS "idx_videos_deleted_at" ON "asset"."Videos"("deleted_at");

COMMENT ON TABLE "asset"."Videos" IS 'Video binary metadata; duration and optional display dimensions; uploader_id references auth.Accounts at app level.';
COMMENT ON COLUMN "asset"."Videos"."id" IS 'Primary key UUID; typically matches stored object id.';
COMMENT ON COLUMN "asset"."Videos"."file_path" IS 'Relative path from storage root or object key.';
COMMENT ON COLUMN "asset"."Videos"."file_name" IS 'Original filename.';
COMMENT ON COLUMN "asset"."Videos"."file_size" IS 'File size in bytes.';
COMMENT ON COLUMN "asset"."Videos"."mime_type" IS 'MIME type as uploaded or detected.';
COMMENT ON COLUMN "asset"."Videos"."duration_seconds" IS 'Media duration when known.';
COMMENT ON COLUMN "asset"."Videos"."width" IS 'Pixel width if known.';
COMMENT ON COLUMN "asset"."Videos"."height" IS 'Pixel height if known.';
COMMENT ON COLUMN "asset"."Videos"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
COMMENT ON COLUMN "asset"."Videos"."created_at" IS 'Insert time.';
COMMENT ON COLUMN "asset"."Videos"."updated_at" IS 'Last metadata update.';
COMMENT ON COLUMN "asset"."Videos"."deleted_at" IS 'Soft-delete; NULL if object still active.';

-- imported from schemas/asset/tables/audios.sql
CREATE TABLE IF NOT EXISTS "asset"."Audios" (
  "id" UUID PRIMARY KEY,
  "file_path" VARCHAR(500) NOT NULL,
  "file_name" VARCHAR(255) NOT NULL,
  "file_size" BIGINT NOT NULL,
  "mime_type" VARCHAR(100) NOT NULL,
  "duration_seconds" DOUBLE PRECISION,
  "uploader_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);
CREATE INDEX IF NOT EXISTS "idx_audios_file_path" ON "asset"."Audios"("file_path");
CREATE INDEX IF NOT EXISTS "idx_audios_uploader_id" ON "asset"."Audios"("uploader_id");
CREATE INDEX IF NOT EXISTS "idx_audios_mime_type" ON "asset"."Audios"("mime_type");
CREATE INDEX IF NOT EXISTS "idx_audios_deleted_at" ON "asset"."Audios"("deleted_at");

COMMENT ON TABLE "asset"."Audios" IS 'Audio binary metadata; duration when known; uploader_id references auth.Accounts at app level.';
COMMENT ON COLUMN "asset"."Audios"."id" IS 'Primary key UUID; typically matches stored object id.';
COMMENT ON COLUMN "asset"."Audios"."file_path" IS 'Relative path from storage root or object key.';
COMMENT ON COLUMN "asset"."Audios"."file_name" IS 'Original filename.';
COMMENT ON COLUMN "asset"."Audios"."file_size" IS 'File size in bytes.';
COMMENT ON COLUMN "asset"."Audios"."mime_type" IS 'MIME type as uploaded or detected.';
COMMENT ON COLUMN "asset"."Audios"."duration_seconds" IS 'Media duration when known.';
COMMENT ON COLUMN "asset"."Audios"."uploader_id" IS 'Reference to auth.Accounts.id (application-level, no FK).';
COMMENT ON COLUMN "asset"."Audios"."created_at" IS 'Insert time.';
COMMENT ON COLUMN "asset"."Audios"."updated_at" IS 'Last metadata update.';
COMMENT ON COLUMN "asset"."Audios"."deleted_at" IS 'Soft-delete; NULL if object still active.';

-- imported from schemas/asset/views/images_active_view.sql
CREATE OR REPLACE VIEW "asset"."ImagesActiveView" AS
SELECT
  "id", "file_path", "file_name", "file_size", "mime_type",
  "width", "height", "alt_text", "uploader_id", "created_at", "updated_at"
FROM "asset"."Images"
WHERE "deleted_at" IS NULL;

COMMENT ON VIEW "asset"."ImagesActiveView" IS 'Non-deleted image rows only; omits deleted_at column.';

-- imported from schemas/asset/views/files_active_view.sql
CREATE OR REPLACE VIEW "asset"."FilesActiveView" AS
SELECT
  "id", "file_path", "file_name", "file_size", "mime_type",
  "uploader_id", "created_at", "updated_at"
FROM "asset"."Files"
WHERE "deleted_at" IS NULL;

COMMENT ON VIEW "asset"."FilesActiveView" IS 'Non-deleted file rows only; omits deleted_at column.';

-- imported from schemas/asset/views/videos_active_view.sql
CREATE OR REPLACE VIEW "asset"."VideosActiveView" AS
SELECT
  "id", "file_path", "file_name", "file_size", "mime_type",
  "duration_seconds", "width", "height", "uploader_id", "created_at", "updated_at"
FROM "asset"."Videos"
WHERE "deleted_at" IS NULL;

COMMENT ON VIEW "asset"."VideosActiveView" IS 'Non-deleted video rows only; omits deleted_at column.';

-- imported from schemas/asset/views/audios_active_view.sql
CREATE OR REPLACE VIEW "asset"."AudiosActiveView" AS
SELECT
  "id", "file_path", "file_name", "file_size", "mime_type",
  "duration_seconds", "uploader_id", "created_at", "updated_at"
FROM "asset"."Audios"
WHERE "deleted_at" IS NULL;

COMMENT ON VIEW "asset"."AudiosActiveView" IS 'Non-deleted audio rows only; omits deleted_at column.';

-- imported from schemas/system/main.sql

-- imported from schemas/system/schema.sql
CREATE SCHEMA IF NOT EXISTS "system";
COMMENT ON SCHEMA "system" IS 'System-level state and administration schema';

-- imported from schemas/system/tables/system_state.sql
-- System State table: Records system initialization and bootstrap state
-- Used to check if system has been initialized on service startup
-- Logic: Check if no record exists OR latest record has initialized = false -> system is not initialized
-- On initialization: Create new record with initialized = true
-- On reset: Delete all records or set initialized = false (allows re-initialization)
CREATE TABLE IF NOT EXISTS "system"."system_state" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "initialized" BOOLEAN NOT NULL DEFAULT false, -- Whether system has been initialized
  "initialized_at" TIMESTAMP, -- When system was initialized
  "initialization_version" VARCHAR(50), -- Version of initialization schema/data
  "last_reset_at" TIMESTAMP, -- Last time system was reset
  "last_reset_by" UUID, -- Account ID who reset the system (auth.Accounts.id, application-level).
  "reset_count" INTEGER NOT NULL DEFAULT 0, -- Number of times system has been reset
  "metadata" JSONB DEFAULT '{}'::jsonb, -- Extended fields: {"bootstrap_token": "...", "services_initialized": [...], ...}
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_system_state_initialized" ON "system"."system_state"("initialized");
CREATE INDEX IF NOT EXISTS "idx_system_state_created_at" ON "system"."system_state"("created_at" DESC); -- For querying latest record

COMMENT ON TABLE "system"."system_state" IS 'System initialization state: tracks if system has been bootstrapped. Query logic: SELECT initialized FROM system_state ORDER BY created_at DESC LIMIT 1. If no record exists OR initialized = false, system is not initialized.';
COMMENT ON COLUMN "system"."system_state"."id" IS 'UUID primary key (not fixed, allows multiple records for reset/re-initialization)';
COMMENT ON COLUMN "system"."system_state"."initialized" IS 'Whether system has been initialized (checked on service startup). Always check latest record by created_at DESC.';
COMMENT ON COLUMN "system"."system_state"."initialized_at" IS 'Timestamp when system was initialized via /bootstrap/initialize';
COMMENT ON COLUMN "system"."system_state"."initialization_version" IS 'Version of initialization schema/data for migration tracking';
COMMENT ON COLUMN "system"."system_state"."last_reset_at" IS 'Timestamp when system was last reset';
COMMENT ON COLUMN "system"."system_state"."last_reset_by" IS 'User ID who reset the system. Even if database is cleared after reset, can be traced via log files for accountability';
COMMENT ON COLUMN "system"."system_state"."created_at" IS 'Record creation time. Used to determine latest state (ORDER BY created_at DESC LIMIT 1)';
