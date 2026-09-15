CREATE TABLE IF NOT EXISTS "auth"."RefreshTokens" (
  "id" UUID PRIMARY KEY,
  "account_id" UUID NOT NULL,
  "identity_id" UUID NULL,
  "profile_id" UUID NULL,
  "profile_scope" "auth".profile_scope NULL,
  "device_id" TEXT NULL,
  "token_hash" TEXT NOT NULL,
  "expires_at" TIMESTAMP NOT NULL,
  "revoked_at" TIMESTAMP NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP NULL,
  CONSTRAINT "fk_refresh_tokens_account_id" FOREIGN KEY ("account_id") REFERENCES "auth"."Accounts" ("id") ON DELETE CASCADE,
  CONSTRAINT "fk_refresh_tokens_identity_id" FOREIGN KEY ("identity_id") REFERENCES "auth"."Identities" ("id") ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_account_id" ON "auth"."RefreshTokens" ("account_id");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_identity_id" ON "auth"."RefreshTokens" ("identity_id");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_device_id" ON "auth"."RefreshTokens" ("device_id");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_account_device_active"
  ON "auth"."RefreshTokens" ("account_id", "device_id")
  WHERE "device_id" IS NOT NULL AND "revoked_at" IS NULL AND "deleted_at" IS NULL;
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_expires_at" ON "auth"."RefreshTokens" ("expires_at");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_revoked_at" ON "auth"."RefreshTokens" ("revoked_at");
CREATE INDEX IF NOT EXISTS "idx_refresh_tokens_deleted_at" ON "auth"."RefreshTokens" ("deleted_at");

COMMENT ON TABLE "auth"."RefreshTokens" IS 'Opaque refresh tokens; store only hashes at rest.';
COMMENT ON COLUMN "auth"."RefreshTokens"."identity_id" IS 'Which login issued the token; NULL if legacy / unspecified.';
COMMENT ON COLUMN "auth"."RefreshTokens"."device_id" IS 'Client device/session identifier for same-device token revocation policy.';
COMMENT ON COLUMN "auth"."RefreshTokens"."deleted_at" IS 'Soft-delete token row (e.g. purge); prefer revoked_at for normal logout.';
COMMENT ON COLUMN "auth"."RefreshTokens"."id" IS 'Primary key UUID.';
COMMENT ON COLUMN "auth"."RefreshTokens"."account_id" IS 'Owning account; CASCADE delete.';
COMMENT ON COLUMN "auth"."RefreshTokens"."token_hash" IS 'Hash of opaque refresh token at rest.';
COMMENT ON COLUMN "auth"."RefreshTokens"."expires_at" IS 'Absolute expiry instant.';
COMMENT ON COLUMN "auth"."RefreshTokens"."revoked_at" IS 'When token was invalidated server-side.';
COMMENT ON COLUMN "auth"."RefreshTokens"."created_at" IS 'Issued-at time.';

-- Maintainer notes (after DDL):
-- 刷新会话令牌（存哈希，不存明文）。
-- device_id 用于“同设备 revoke，其他设备保留”。
