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
