-- User with Role View
-- Purpose: Provides user information with role details in a single query
-- Use Case: API endpoints, admin dashboards, authorization checks
-- Returns: One row per user with role information

CREATE VIEW "auth"."user_with_role_view" AS
SELECT
  u."id" AS "user_id",
  u."username",
  u."email",
  u."phone",
  u."status",
  u."is_verified",
  u."last_login_at",
  u."role_id",
  r."name" AS "role_name",
  r."description" AS "role_description",
  r."permissions" AS "role_permissions",
  u."created_at" AS "user_created_at",
  u."updated_at" AS "user_updated_at"
FROM
  "auth"."users" u
  LEFT JOIN "auth"."roles" r ON u."role_id" = r."id"
WHERE
  u."deleted_at" IS NULL;

