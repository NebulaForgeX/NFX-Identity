-- User profile table for additional user information
-- Note: id directly references users.id (one-to-one relationship)
-- Avatar and background images are managed via user_avatars and user_images tables
CREATE TABLE IF NOT EXISTS "directory"."user_profiles" (
  "id" UUID PRIMARY KEY REFERENCES "directory"."users"("id") ON DELETE CASCADE,
  "role" VARCHAR(100), -- User role for external companies (e.g., "admin", "user", "manager")
  "first_name" VARCHAR(100),
  "last_name" VARCHAR(100),
  "nickname" VARCHAR(50) UNIQUE, -- 昵称（唯一）
  "display_name" VARCHAR(100), -- 显示名称（不唯一）
  "bio" TEXT,
  "birthday" DATE,
  "age" INTEGER, -- 年龄（可以从 birthday 计算，也可以单独存储）
  "gender" VARCHAR(20),
  "location" VARCHAR(255), -- 位置信息（从结构化数据拼接：country, province, city, timezone）
  "website" VARCHAR(255),
  "github" VARCHAR(255), -- GitHub 用户名或 URL
  "social_links" JSONB DEFAULT '{}'::jsonb, -- 结构化社交链接: {"twitter": "https://twitter.com/username", "linkedin": "https://linkedin.com/in/username", "instagram": "https://instagram.com/username", "youtube": "https://youtube.com/username"} (不包含 github/website)
  "skills" JSONB DEFAULT '{}'::jsonb, -- 技能: {"golang": 10, "python": 8, ...}
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_user_profiles_nickname" ON "directory"."user_profiles"("nickname");
CREATE INDEX IF NOT EXISTS "idx_user_profiles_github" ON "directory"."user_profiles"("github");
CREATE INDEX IF NOT EXISTS "idx_user_profiles_deleted_at" ON "directory"."user_profiles"("deleted_at");
CREATE INDEX IF NOT EXISTS "idx_user_profiles_display_name" ON "directory"."user_profiles"("display_name");
CREATE INDEX IF NOT EXISTS "idx_user_profiles_first_and_last_name" ON "directory"."user_profiles"("first_name", "last_name");
CREATE INDEX IF NOT EXISTS "idx_user_profiles_birthday" ON "directory"."user_profiles"("birthday");
CREATE INDEX IF NOT EXISTS "idx_user_profiles_gender" ON "directory"."user_profiles"("gender");
CREATE INDEX IF NOT EXISTS "idx_user_profiles_location" ON "directory"."user_profiles"("location");
CREATE INDEX IF NOT EXISTS "idx_user_profiles_website" ON "directory"."user_profiles"("website");
CREATE INDEX IF NOT EXISTS "idx_user_profiles_role" ON "directory"."user_profiles"("role");


