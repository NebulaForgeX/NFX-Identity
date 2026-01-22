-- User Preferences table for user settings and preferences
-- Stores user-specific preferences and settings
-- Note: id directly references users.id (one-to-one relationship)
CREATE TABLE IF NOT EXISTS "directory"."user_preferences" (
  "id" UUID PRIMARY KEY REFERENCES "directory"."users"("id") ON DELETE CASCADE,
  "theme" VARCHAR(50) DEFAULT 'light', -- Theme preference: light, dark, auto
  "language" VARCHAR(10) DEFAULT 'en', -- Language preference: en, zh, etc.
  "timezone" VARCHAR(50) DEFAULT 'Etc/UTC', -- Timezone preference
  "notifications" JSONB DEFAULT '{}'::jsonb, -- Notification preferences: {"email": true, "sms": false, "push": true, ...}
  "privacy" JSONB DEFAULT '{}'::jsonb, -- Privacy settings: {"profile_visibility": "public", "email_visibility": "private", ...}
  "display" JSONB DEFAULT '{}'::jsonb, -- Display preferences: {"date_format": "YYYY-MM-DD", "time_format": "24h", ...}
  "other" JSONB DEFAULT '{}'::jsonb, -- Other custom preferences
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_user_preferences_deleted_at" ON "directory"."user_preferences"("deleted_at");

