-- Active: 1765190245476@@127.0.0.1@10105@lyuauth_dev@permission
-- ========================================
-- 插入授权码示例 (Insert Authorization Code)
-- ========================================
-- 使用方法：根据实际需求修改参数后执行
-- ========================================

-- 示例 1：插入一个单次使用的授权码（不过期，立即激活）
INSERT INTO "permission"."authorization_codes" (
  "code",
  "max_uses",
  "used_count",
  "created_by",
  "expires_at",
  "is_active"
) VALUES (
  'ADMIN-2025-001',           -- 授权码（必须唯一）
  1,                          -- 最大使用次数
  0,                          -- 当前使用次数（初始为0）
  NULL,                       -- 创建者用户ID（可选，NULL表示系统创建）
  NULL,                       -- 过期时间（可选，NULL表示永不过期）
  true                        -- 是否激活
);

-- 示例 2：插入一个多次使用的授权码（30天后过期）
INSERT INTO "permission"."authorization_codes" (
  "code",
  "max_uses",
  "used_count",
  "created_by",
  "expires_at",
  "is_active"
) VALUES (
  'COMPANY-2024-001',         -- 授权码
  10,                         -- 最多可以使用10次
  0,                          -- 初始使用次数为0
  NULL,                       -- 创建者用户ID
  CURRENT_TIMESTAMP + INTERVAL '30 days',  -- 30天后过期
  true                        -- 激活状态
);

-- 示例 3：插入一个指定创建者的授权码
-- 注意：created_by 必须是 auth.users 表中存在的用户ID
INSERT INTO "permission"."authorization_codes" (
  "code",
  "max_uses",
  "used_count",
  "created_by",
  "expires_at",
  "is_active"
) VALUES (
  'USER-2024-001',
  5,
  0,
  NULL,                       -- 创建者用户ID
  NULL,
  true
);

-- ========================================
-- 批量插入示例
-- ========================================
INSERT INTO "permission"."authorization_codes" (
  "code",
  "max_uses",
  "used_count",
  "created_by",
  "expires_at",
  "is_active"
) VALUES
  ('BATCH-001', 1, 0, NULL, NULL, true),
  ('BATCH-002', 1, 0, NULL, NULL, true),
  ('BATCH-003', 1, 0, NULL, NULL, true);
