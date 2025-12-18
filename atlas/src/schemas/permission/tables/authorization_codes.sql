-- Authorization code table for admin panel registration
-- Used by companies to allow their employees to register and login to Identity-Admin platform
-- Senior administrators can create authorization entries, and employees can register using these codes
CREATE TABLE IF NOT EXISTS "permission"."authorization_codes" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "code" VARCHAR(255) NOT NULL UNIQUE, -- Unique authorization code
  "max_uses" INTEGER NOT NULL DEFAULT 1, -- Maximum number of times this code can be used
  "used_count" INTEGER NOT NULL DEFAULT 0, -- Current number of times this code has been used
  "created_by" UUID, -- User ID who created this authorization code (references auth.users)
  "expires_at" TIMESTAMP, -- Optional expiration date
  "is_active" BOOLEAN NOT NULL DEFAULT true, -- Whether this code is currently active
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_authorization_codes_code" ON "permission"."authorization_codes"("code");
CREATE INDEX IF NOT EXISTS "idx_authorization_codes_created_by" ON "permission"."authorization_codes"("created_by");
CREATE INDEX IF NOT EXISTS "idx_authorization_codes_is_active" ON "permission"."authorization_codes"("is_active");
CREATE INDEX IF NOT EXISTS "idx_authorization_codes_expires_at" ON "permission"."authorization_codes"("expires_at");
CREATE INDEX IF NOT EXISTS "idx_authorization_codes_deleted_at" ON "permission"."authorization_codes"("deleted_at");

-- Add foreign key constraint to auth.users (cross-schema reference)
-- Note: This requires the auth schema to exist first
-- ALTER TABLE "permission"."authorization_codes" 
--   ADD CONSTRAINT "fk_authorization_codes_created_by" 
--   FOREIGN KEY ("created_by") REFERENCES "auth"."users"("id") ON DELETE SET NULL;
