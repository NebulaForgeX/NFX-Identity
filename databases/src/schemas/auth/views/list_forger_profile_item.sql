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
