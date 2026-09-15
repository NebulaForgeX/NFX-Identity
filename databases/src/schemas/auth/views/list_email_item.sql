CREATE OR REPLACE VIEW "auth"."ListEmailItems" AS
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

COMMENT ON VIEW "auth"."ListEmailItems" IS
'One row per non-deleted email. Filter by account_id to list account emails.';
