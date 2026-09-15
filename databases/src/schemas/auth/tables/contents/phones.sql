CREATE TABLE IF NOT EXISTS "auth"."Phones" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "phone" VARCHAR(32) NOT NULL,
  "is_primary" BOOLEAN NOT NULL DEFAULT FALSE,
  "verified_at" TIMESTAMP NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_phones_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "uq_phones_phone_active" ON "auth"."Phones" ("phone") WHERE "deleted_at" IS NULL;

CREATE INDEX IF NOT EXISTS "idx_phones_account_id" ON "auth"."Phones" ("account_id");
CREATE INDEX IF NOT EXISTS "idx_phones_deleted_at" ON "auth"."Phones" ("deleted_at");
CREATE UNIQUE INDEX IF NOT EXISTS "uq_phones_one_primary_per_account" ON "auth"."Phones" ("account_id") WHERE "is_primary" = TRUE AND "deleted_at" IS NULL;

COMMENT ON COLUMN "auth"."Phones"."verified_at" IS 'NULL = not verified yet.';
COMMENT ON COLUMN "auth"."Phones"."is_primary" IS 'At most one non-deleted primary phone per account.';
COMMENT ON COLUMN "auth"."Phones"."deleted_at" IS 'Soft-delete row; same canonical phone may be re-bound via INSERT new row.';
COMMENT ON TABLE "auth"."Phones" IS 'Phone bindings; partial unique on phone WHERE deleted_at IS NULL mirrors Emails pattern.';
COMMENT ON COLUMN "auth"."Phones"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."Phones"."account_id" IS 'Owning auth.Accounts row; CASCADE delete.';
COMMENT ON COLUMN "auth"."Phones"."phone" IS 'Canonical phone string; unique among non-deleted rows.';
COMMENT ON COLUMN "auth"."Phones"."created_at" IS 'Row insert time.';
COMMENT ON COLUMN "auth"."Phones"."updated_at" IS 'Last edit time.';

-- Maintainer notes (after DDL):
-- 手机号（可多号码；一条标记为主号码）。语义对齐 Emails：软删历史 + 再绑定同号用新行。
-- 建议在应用层存 E.164 等规范格式，避免同一号码多种写法绕过唯一。
