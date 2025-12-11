-- Profile Occupations table for user work/employment history
-- One profile can have multiple occupation records
CREATE TABLE IF NOT EXISTS "auth"."profile_occupations" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "profile_id" UUID NOT NULL REFERENCES "auth"."profiles"("id") ON DELETE CASCADE,
  "company" VARCHAR(255) NOT NULL, -- 公司名称
  "position" VARCHAR(255) NOT NULL, -- 职位
  "department" VARCHAR(255), -- 部门
  "industry" VARCHAR(100), -- 行业
  "location" VARCHAR(255), -- 工作地点
  "employment_type" VARCHAR(50), -- 雇佣类型：full-time, part-time, contract, internship, etc.
  "start_date" DATE, -- 开始日期
  "end_date" DATE, -- 结束日期（NULL 表示当前工作）
  "is_current" BOOLEAN NOT NULL DEFAULT false, -- 是否当前工作
  "description" TEXT, -- 工作描述
  "responsibilities" TEXT, -- 职责
  "achievements" TEXT, -- 成就
  "skills_used" TEXT[], -- 使用的技能列表
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_profile_occupations_profile_id" ON "auth"."profile_occupations"("profile_id");
CREATE INDEX IF NOT EXISTS "idx_profile_occupations_company" ON "auth"."profile_occupations"("company");
CREATE INDEX IF NOT EXISTS "idx_profile_occupations_position" ON "auth"."profile_occupations"("position");
CREATE INDEX IF NOT EXISTS "idx_profile_occupations_industry" ON "auth"."profile_occupations"("industry");
CREATE INDEX IF NOT EXISTS "idx_profile_occupations_is_current" ON "auth"."profile_occupations"("is_current");
CREATE INDEX IF NOT EXISTS "idx_profile_occupations_deleted_at" ON "auth"."profile_occupations"("deleted_at");
