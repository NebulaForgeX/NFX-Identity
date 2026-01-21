-- User Phones table for multiple phone numbers per user
-- One user can have multiple phone numbers (primary, secondary, etc.)
CREATE TABLE IF NOT EXISTS "directory"."user_phones" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID NOT NULL REFERENCES "directory"."users"("id") ON DELETE CASCADE,
  "phone" VARCHAR(20) NOT NULL,
  "country_code" VARCHAR(10), -- Country code (e.g., +1, +86)
  "is_primary" BOOLEAN NOT NULL DEFAULT false, -- Primary phone number
  "is_verified" BOOLEAN NOT NULL DEFAULT false, -- Phone verification status
  "verified_at" TIMESTAMP, -- When the phone was verified
  "verification_code" VARCHAR(10), -- SMS verification code
  "verification_expires_at" TIMESTAMP, -- Verification code expiration time
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP,
  UNIQUE("phone")
);

CREATE INDEX IF NOT EXISTS "idx_user_phones_user_id" ON "directory"."user_phones"("user_id");
CREATE INDEX IF NOT EXISTS "idx_user_phones_phone" ON "directory"."user_phones"("phone");
CREATE INDEX IF NOT EXISTS "idx_user_phones_is_primary" ON "directory"."user_phones"("user_id", "is_primary") WHERE "is_primary" = true;
CREATE INDEX IF NOT EXISTS "idx_user_phones_is_verified" ON "directory"."user_phones"("is_verified");
CREATE INDEX IF NOT EXISTS "idx_user_phones_deleted_at" ON "directory"."user_phones"("deleted_at");

