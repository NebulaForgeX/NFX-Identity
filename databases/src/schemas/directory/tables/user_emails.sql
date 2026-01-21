-- User Emails table for multiple email addresses per user
-- One user can have multiple email addresses (primary, secondary, etc.)
-- Note: Foreign key to users table is added in users.sql after users table is created
CREATE TABLE IF NOT EXISTS "directory"."user_emails" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID NOT NULL REFERENCES "directory"."users"("id") ON DELETE CASCADE,
  "email" VARCHAR(255) NOT NULL,
  "is_primary" BOOLEAN NOT NULL DEFAULT false, -- Primary email address
  "is_verified" BOOLEAN NOT NULL DEFAULT false, -- Email verification status
  "verified_at" TIMESTAMP, -- When the email was verified
  "verification_token" VARCHAR(255), -- Token for email verification
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP,
  UNIQUE("email")
);

CREATE INDEX IF NOT EXISTS "idx_user_emails_user_id" ON "directory"."user_emails"("user_id");
CREATE INDEX IF NOT EXISTS "idx_user_emails_email" ON "directory"."user_emails"("email");
CREATE INDEX IF NOT EXISTS "idx_user_emails_is_primary" ON "directory"."user_emails"("user_id", "is_primary") WHERE "is_primary" = true;
CREATE INDEX IF NOT EXISTS "idx_user_emails_is_verified" ON "directory"."user_emails"("is_verified");
CREATE INDEX IF NOT EXISTS "idx_user_emails_deleted_at" ON "directory"."user_emails"("deleted_at");

