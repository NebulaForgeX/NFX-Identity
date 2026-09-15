CREATE TABLE IF NOT EXISTS "auth"."Emails" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "is_primary" BOOLEAN NOT NULL DEFAULT FALSE,
  "verified_at" TIMESTAMP NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_emails_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "uq_emails_email_ci" ON "auth"."Emails" (LOWER("email")) WHERE "deleted_at" IS NULL;

CREATE INDEX IF NOT EXISTS "idx_emails_account_id" ON "auth"."Emails" ("account_id");
CREATE INDEX IF NOT EXISTS "idx_emails_deleted_at" ON "auth"."Emails" ("deleted_at");
CREATE UNIQUE INDEX IF NOT EXISTS "uq_emails_one_primary_per_account" ON "auth"."Emails" ("account_id") WHERE "is_primary" = TRUE AND "deleted_at" IS NULL;

COMMENT ON TABLE "auth"."Emails" IS 'Email bindings; soft-deleted rows kept for history. Partial unique (LOWER(email)) WHERE deleted_at IS NULL avoids collision when re-registering. Avoid hard-deleting auth.Accounts if you need email audit trail (CASCADE removes child rows).';
COMMENT ON COLUMN "auth"."Emails"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."Emails"."account_id" IS 'Owning auth.Accounts row; CASCADE delete.';
COMMENT ON COLUMN "auth"."Emails"."email" IS 'Email address; uniqueness enforced case-insensitively among non-deleted rows.';
COMMENT ON COLUMN "auth"."Emails"."verified_at" IS 'NULL = not verified yet.';
COMMENT ON COLUMN "auth"."Emails"."is_primary" IS 'At most one non-deleted primary email per account (partial unique index).';
COMMENT ON COLUMN "auth"."Emails"."created_at" IS 'Row insert time.';
COMMENT ON COLUMN "auth"."Emails"."updated_at" IS 'Last edit time.';
COMMENT ON COLUMN "auth"."Emails"."deleted_at" IS 'Soft-delete only: row remains for audit. Partial unique applies only where deleted_at IS NULL. Re-register adds another row.';

-- Maintainer notes (after DDL):
-- 登录邮箱（可多邮箱；一条标记为主邮箱）。
-- 同一 LOWER(email) 可存在多行：多行「已注销」+ 至多一行「当前有效」。
--   注销：本行写 deleted_at，不删行。再注册同邮箱：INSERT 新行（新 id、新 account_id、deleted_at NULL），旧行 deleted_at 不动 → 每次注销时间仍留在旧行上。
-- 部分唯一索引 (LOWER(email) WHERE deleted_at IS NULL) 只保证：同时只能有一条「未删除」的该邮箱；不限制历史行条数。
-- 若对 Accounts 物理 DELETE，CASCADE 会删掉子表 Emails 行 → 历史消失；用户侧应软删账号，勿物理删。
