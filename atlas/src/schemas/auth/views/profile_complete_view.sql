-- Active: 1765190245476@@127.0.0.1@10105@lyuauth_dev@auth
-- Profile Complete View
-- Purpose: Provides complete profile information with occupations and educations
-- Use Case: User profile pages, resume displays, comprehensive user information
-- Returns: One row per profile with aggregated occupations and educations

CREATE VIEW "auth"."profile_complete_view" AS
SELECT
  u."id" AS "user_id",
  u."username",
  u."email",
  u."phone" AS "user_phone",
  u."status" AS "user_status",
  u."is_verified",
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
  ARRAY_AGG(
    jsonb_build_object(
      'id', o."id",
      'company', o."company",
      'position', o."position",
      'department', o."department",
      'industry', o."industry",
      'location', o."location",
      'employment_type', o."employment_type",
      'start_date', o."start_date",
      'end_date', o."end_date",
      'is_current', o."is_current",
      'description', o."description"
    )
    ORDER BY o."is_current" DESC, o."start_date" DESC
  ) FILTER (WHERE o."id" IS NOT NULL) AS "occupations",
  ARRAY_AGG(
    jsonb_build_object(
      'id', e."id",
      'school', e."school",
      'degree', e."degree",
      'major', e."major",
      'field_of_study', e."field_of_study",
      'start_date', e."start_date",
      'end_date', e."end_date",
      'is_current', e."is_current",
      'grade', e."grade"
    )
    ORDER BY e."is_current" DESC, e."start_date" DESC
  ) FILTER (WHERE e."id" IS NOT NULL) AS "educations",
  p."created_at" AS "profile_created_at",
  p."updated_at" AS "profile_updated_at"
FROM
  "auth"."users" u
  LEFT JOIN "auth"."profiles" p ON p."user_id" = u."id"
  LEFT JOIN "auth"."profile_occupations" o ON o."profile_id" = p."id" AND o."deleted_at" IS NULL
  LEFT JOIN "auth"."profile_educations" e ON e."profile_id" = p."id" AND e."deleted_at" IS NULL
WHERE
  u."deleted_at" IS NULL
GROUP BY
  u."id", u."username", u."email", u."phone", u."status", u."is_verified",
  p."id", p."first_name", p."last_name", p."nickname", p."display_name",
  p."avatar_id", p."background_id", p."background_ids", p."bio", p."phone",
  p."birthday", p."age", p."gender", p."location", p."website", p."github",
  p."social_links", p."preferences", p."skills", p."privacy_settings",
  p."created_at", p."updated_at";

