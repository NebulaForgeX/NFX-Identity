-- Active: 1765190245476@@127.0.0.1@10105@lyuauth_dev@permission
-- ========================================
-- 插入 SuperAdmin 权限
-- Insert SuperAdmin Permission
-- ========================================

INSERT INTO "permission"."permissions" (
  "tag",
  "name",
  "description",
  "category",
  "is_system"
) VALUES (
  'SuperAdmin',                    -- tag
  '超级管理员',                    -- name
  '拥有系统所有权限的超级管理员',  -- description
  'system'::"permission"."permission_category",  -- category (使用 ENUM)
  true                             -- is_system (系统权限，不可删除)
)
ON CONFLICT ("tag") DO NOTHING;    -- 如果 tag 已存在，则不做任何操作
