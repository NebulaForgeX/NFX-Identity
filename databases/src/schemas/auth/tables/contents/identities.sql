CREATE TABLE IF NOT EXISTS "auth"."Identities" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "identity_provider" "auth".identity_provider NOT NULL,
  "provider_subject" TEXT NOT NULL,
  "password_hash" TEXT NULL,
  "metadata" JSONB NULL,
  "last_login_at" TIMESTAMP NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_identities_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_identities_account_id" ON "auth"."Identities" ("account_id");
CREATE INDEX IF NOT EXISTS "idx_identities_deleted_at" ON "auth"."Identities" ("deleted_at");
CREATE UNIQUE INDEX IF NOT EXISTS "uq_identities_identity_provider_subject_active" ON "auth"."Identities" ("identity_provider", "provider_subject") WHERE "deleted_at" IS NULL;

COMMENT ON TABLE "auth"."Identities" IS 'Auth credentials per identity_provider; PASSWORD uses provider_subject = normalized login id (e.g. email).';
COMMENT ON COLUMN "auth"."Identities"."identity_provider" IS 'auth.identity_provider enum (password / future OAuth).';
COMMENT ON COLUMN "auth"."Identities"."provider_subject" IS 'Unique within identity_provider among non-deleted rows; email for password, sub for OAuth.';
COMMENT ON COLUMN "auth"."Identities"."password_hash" IS 'Set only when identity_provider = password; bcrypt/argon2 hash at application layer.';
COMMENT ON COLUMN "auth"."Identities"."metadata" IS 'Optional provider-specific payload (e.g. OAuth raw claims subset).';
COMMENT ON COLUMN "auth"."Identities"."deleted_at" IS 'Soft-delete identity; frees provider_subject for re-link.';
COMMENT ON COLUMN "auth"."Identities"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."Identities"."account_id" IS 'Owning auth.Accounts row; CASCADE delete.';
COMMENT ON COLUMN "auth"."Identities"."last_login_at" IS 'Last successful login instant for this identity.';
COMMENT ON COLUMN "auth"."Identities"."created_at" IS 'Row insert time.';
COMMENT ON COLUMN "auth"."Identities"."updated_at" IS 'Last credential/metadata update.';

-- Maintainer notes (after DDL):
-- 登录身份：密码 / 未来 OAuth；凭证不与 Accounts 混放在一行。
