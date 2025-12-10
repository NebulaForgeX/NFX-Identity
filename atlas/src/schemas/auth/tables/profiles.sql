-- User profile table for additional user information
-- Note: avatar_id, background_id, and background_ids are UUIDs (not foreign keys) to avoid cross-service dependencies
-- These UUIDs can be used to query the image service via API
-- background_id: current/active background image
-- background_ids: collection of background images (history/favorites)
CREATE TABLE IF NOT EXISTS "auth"."profiles" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID NOT NULL UNIQUE REFERENCES "auth"."users"("id") ON DELETE CASCADE,
  "first_name" VARCHAR(100),
  "last_name" VARCHAR(100),
  "nickname" VARCHAR(50) UNIQUE, -- 昵称（唯一）
  "display_name" VARCHAR(100), -- 显示名称（不唯一）
  "avatar_id" UUID,
  "background_id" UUID,
  "background_ids" UUID[],
  "bio" TEXT,
  "phone" VARCHAR(20),
  "birthday" DATE,
  "age" INTEGER, -- 年龄（可以从 birthday 计算，也可以单独存储）
  "gender" VARCHAR(20),
  "location" VARCHAR(255), -- 位置信息（从结构化数据拼接：country, province, city, timezone）
  "website" VARCHAR(255),
  "github" VARCHAR(255), -- GitHub 用户名或 URL
  "social_links" JSONB DEFAULT '{}'::jsonb, -- 结构化社交链接: {"twitter": "https://twitter.com/username", "linkedin": "https://linkedin.com/in/username", "instagram": "https://instagram.com/username", "youtube": "https://youtube.com/username"} (不包含 github/website)
  "preferences" JSONB DEFAULT '{}'::jsonb, -- 偏好设置: {"theme": "dark", "language": "en", ...}
  "skills" JSONB DEFAULT '{}'::jsonb, -- 技能: {"golang": 10, "python": 8, ...}
  "privacy_settings" JSONB DEFAULT '{}'::jsonb, -- 隐私设置
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_profiles_user_id" ON "auth"."profiles"("user_id");
CREATE INDEX IF NOT EXISTS "idx_profiles_nickname" ON "auth"."profiles"("nickname");
CREATE INDEX IF NOT EXISTS "idx_profiles_github" ON "auth"."profiles"("github");
CREATE INDEX IF NOT EXISTS "idx_profiles_avatar_id" ON "auth"."profiles"("avatar_id");
CREATE INDEX IF NOT EXISTS "idx_profiles_background_id" ON "auth"."profiles"("background_id");
CREATE INDEX IF NOT EXISTS "idx_profiles_deleted_at" ON "auth"."profiles"("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_profiles_display_name" ON "auth"."profiles"("display_name");
CREATE INDEX IF NOT EXISTS "idx_profiles_first_and_last_name" ON "auth"."profiles"("first_name", "last_name");
CREATE INDEX IF NOT EXISTS "idx_profiles_phone" ON "auth"."profiles"("phone");
CREATE INDEX IF NOT EXISTS "idx_profiles_birthday" ON "auth"."profiles"("birthday");
CREATE INDEX IF NOT EXISTS "idx_profiles_gender" ON "auth"."profiles"("gender");
CREATE INDEX IF NOT EXISTS "idx_profiles_location" ON "auth"."profiles"("location");
CREATE INDEX IF NOT EXISTS "idx_profiles_website" ON "auth"."profiles"("website");


