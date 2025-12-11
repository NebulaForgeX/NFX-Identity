-- User with Roles View
-- Purpose: Provides user information with role details (multiple roles supported)
-- Use Case: API endpoints, admin dashboards, authorization checks
-- Returns: One row per user with aggregated roles information

CREATE VIEW "auth"."user_with_role_view" AS
SELECT
  u."id" AS "user_id",
  u."username",
  u."email",
  u."phone",
  u."status",
  u."is_verified",
  u."last_login_at",
  ARRAY_AGG(
    jsonb_build_object(
      'role_id', r."id",
      'role_name', r."name",
      'role_description', r."description",
      'role_permissions', r."permissions"
    )
  ) FILTER (WHERE r."id" IS NOT NULL) AS "roles",
  u."created_at" AS "user_created_at",
  u."updated_at" AS "user_updated_at"
FROM
  "auth"."users" u
  LEFT JOIN "auth"."user_roles" ur ON ur."user_id" = u."id"
  LEFT JOIN "auth"."roles" r ON r."id" = ur."role_id"
WHERE
  u."deleted_at" IS NULL
GROUP BY
  u."id", u."username", u."email", u."phone", u."status", u."is_verified", u."last_login_at", u."created_at", u."updated_at";

