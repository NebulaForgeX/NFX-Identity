-- Occupations table for user work/employment history
-- One profile can have multiple occupation records
CREATE TABLE IF NOT EXISTS "auth"."occupations" (
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

CREATE INDEX IF NOT EXISTS "idx_occupations_profile_id" ON "auth"."occupations"("profile_id");
CREATE INDEX IF NOT EXISTS "idx_occupations_company" ON "auth"."occupations"("company");
CREATE INDEX IF NOT EXISTS "idx_occupations_position" ON "auth"."occupations"("position");
CREATE INDEX IF NOT EXISTS "idx_occupations_industry" ON "auth"."occupations"("industry");
CREATE INDEX IF NOT EXISTS "idx_occupations_is_current" ON "auth"."occupations"("is_current");
CREATE INDEX IF NOT EXISTS "idx_occupations_deleted_at" ON "auth"."occupations"("deleted_at");

