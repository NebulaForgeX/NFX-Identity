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
