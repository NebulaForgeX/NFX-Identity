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

COMMENT ON TABLE "auth"."Accounts" IS 'Core account; credentials in Identities, forger profile in ForgerProfiles, authority profile in AuthorityProfiles, emails in Emails.';
COMMENT ON COLUMN "auth"."Accounts"."id" IS 'Primary key UUID; typically matches application user id.';
COMMENT ON COLUMN "auth"."Accounts"."account_status" IS 'Lifecycle: auth.account_status enum.';
COMMENT ON COLUMN "auth"."Accounts"."signup_platform" IS 'Product that first created this account (auth.signup_platform); written only at signup.';
COMMENT ON COLUMN "auth"."Accounts"."created_at" IS 'Row insert time.';
COMMENT ON COLUMN "auth"."Accounts"."updated_at" IS 'Last account row update.';
COMMENT ON COLUMN "auth"."Accounts"."deleted_at" IS 'Optional soft-delete timestamp when account_status = deleted.';

-- Maintainer notes (after DDL):
-- 账号核心表：不含邮箱/密码（见 Emails / Identities）；角色下沉到各 profile（ForgerProfiles.forger_roles[] / AuthorityProfiles.authority_roles[]）。
-- signup_platform：仅注册创建时写入；DEFAULT nfxidentity。
