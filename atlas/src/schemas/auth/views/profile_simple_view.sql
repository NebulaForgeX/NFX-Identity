-- Profile Simple View
-- Purpose: Provides basic user and profile information without aggregations
-- Use Case: Quick profile lookups, API endpoints that don't need full details
-- Returns: One row per user with basic profile information

CREATE VIEW "auth"."profile_simple_view" AS
SELECT
  u."id" AS "user_id",
  u."username",
  u."email",
  u."phone" AS "user_phone",
  u."status" AS "user_status",
  u."is_verified",
  u."last_login_at",
  u."created_at" AS "user_created_at",
  u."updated_at" AS "user_updated_at",
  -- profile fields
  p."id" AS "profile_id",
  p."first_name",
  p."last_name",
  p."nickname",
  p."display_name",
  p."avatar_id",
  p."background_id",
  p."background_ids",
  p."bio",
  p."phone" AS "profile_phone",
  p."birthday",
  p."age",
  p."gender",
  p."location",
  p."website",
  p."github",
  p."social_links",
  p."preferences",
  p."skills",
  p."privacy_settings",
  p."created_at" AS "profile_created_at",
  p."updated_at" AS "profile_updated_at"
FROM
  "auth"."users" u
  LEFT JOIN "auth"."profiles" p ON p."user_id" = u."id"
WHERE
  u."deleted_at" IS NULL;

