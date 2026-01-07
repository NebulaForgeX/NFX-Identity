-- User Occupations table for user work/employment history
-- One user can have multiple occupation records
CREATE TABLE IF NOT EXISTS "directory"."user_occupations" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" UUID NOT NULL REFERENCES "directory"."users"("id") ON DELETE CASCADE,
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

CREATE INDEX IF NOT EXISTS "idx_user_occupations_user_id" ON "directory"."user_occupations"("user_id");
CREATE INDEX IF NOT EXISTS "idx_user_occupations_company" ON "directory"."user_occupations"("company");
CREATE INDEX IF NOT EXISTS "idx_user_occupations_position" ON "directory"."user_occupations"("position");
CREATE INDEX IF NOT EXISTS "idx_user_occupations_industry" ON "directory"."user_occupations"("industry");
CREATE INDEX IF NOT EXISTS "idx_user_occupations_is_current" ON "directory"."user_occupations"("is_current");
CREATE INDEX IF NOT EXISTS "idx_user_occupations_deleted_at" ON "directory"."user_occupations"("deleted_at");
