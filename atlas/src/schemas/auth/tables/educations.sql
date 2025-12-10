-- Educations table for user education history
-- One profile can have multiple education records
CREATE TABLE IF NOT EXISTS "auth"."educations" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "profile_id" UUID NOT NULL REFERENCES "auth"."profiles"("id") ON DELETE CASCADE,
  "school" VARCHAR(255) NOT NULL, -- 学校名称
  "degree" VARCHAR(100), -- 学位：Bachelor, Master, PhD, etc.
  "major" VARCHAR(255), -- 专业
  "field_of_study" VARCHAR(255), -- 研究领域
  "start_date" DATE, -- 开始日期
  "end_date" DATE, -- 结束日期（NULL 表示在读）
  "is_current" BOOLEAN NOT NULL DEFAULT false, -- 是否在读
  "description" TEXT, -- 描述
  "grade" VARCHAR(50), -- 成绩/GPA
  "activities" TEXT, -- 活动/社团
  "achievements" TEXT, -- 成就
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_educations_profile_id" ON "auth"."educations"("profile_id");
CREATE INDEX IF NOT EXISTS "idx_educations_school" ON "auth"."educations"("school");
CREATE INDEX IF NOT EXISTS "idx_educations_degree" ON "auth"."educations"("degree");
CREATE INDEX IF NOT EXISTS "idx_educations_deleted_at" ON "auth"."educations"("deleted_at");

