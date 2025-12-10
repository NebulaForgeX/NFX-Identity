-- User Badges View
-- Purpose: Provides user information with aggregated badges array
-- Use Case: User profile pages, achievement displays, badge-based filtering
-- Returns: One row per user with all badges aggregated into arrays

CREATE VIEW "auth"."user_badges_view" AS
SELECT
  u."id" AS "user_id",
  u."username",
  u."email",
  p."id" AS "profile_id",
  p."display_name",
  ARRAY_AGG(
    jsonb_build_object(
      'badge_id', b."id",
      'badge_name', b."name",
      'badge_description', b."description",
      'icon_url', b."icon_url",
      'color', b."color",
      'category', b."category",
      'level', pb."level",
      'earned_at', pb."earned_at"
    )
    ORDER BY pb."earned_at" DESC
  ) FILTER (WHERE pb."id" IS NOT NULL) AS "badges"
FROM
  "auth"."users" u
  LEFT JOIN "auth"."profiles" p ON p."user_id" = u."id"
  LEFT JOIN "auth"."profile_badges" pb ON pb."profile_id" = p."id"
  LEFT JOIN "auth"."badges" b ON b."id" = pb."badge_id"
WHERE
  u."deleted_at" IS NULL
GROUP BY
  u."id", u."username", u."email", p."id", p."display_name";

