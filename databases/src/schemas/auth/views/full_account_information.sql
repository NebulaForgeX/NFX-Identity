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
