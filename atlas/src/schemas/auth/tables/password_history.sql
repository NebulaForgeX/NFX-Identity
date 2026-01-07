-- Password History table to prevent password reuse
-- Enterprise compliance: prevent using last N passwords
CREATE TABLE IF NOT EXISTS "auth"."password_history" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID NOT NULL, -- References directory.users.id (application-level consistency)
  "tenant_id" UUID NOT NULL,
  "password_hash" VARCHAR(255) NOT NULL,
  "hash_alg" VARCHAR(50), -- Algorithm used: 'bcrypt', 'argon2id', etc.
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_password_history_user_id" ON "auth"."password_history"("user_id", "created_at" DESC);
CREATE INDEX IF NOT EXISTS "idx_password_history_tenant_id" ON "auth"."password_history"("tenant_id");
CREATE INDEX IF NOT EXISTS "idx_password_history_created_at" ON "auth"."password_history"("created_at");

