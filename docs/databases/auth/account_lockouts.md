# Account Lockouts Table

## 概述 / Overview

`auth.account_lockouts` 表用于管理用户账户的锁定/解锁状态，提供明确的锁定状态，而不是从 login_attempts 聚合。

The `auth.account_lockouts` table manages user lock/unlock state, providing explicit lock state instead of aggregating from login_attempts.

## 使用场景 / Use Cases

### 1. 自动锁定账户
**场景**: 检测到多次失败登录后自动锁定
**例子**:
- 用户连续 5 次输入错误密码
- 操作：`INSERT INTO auth.account_lockouts (user_id, tenant_id, locked_until, lock_reason, locked_by) VALUES ('zhangsan-uuid', 'tenant-uuid', NOW() + INTERVAL '30 minutes', 'too_many_attempts', 'system')`
- 结果：账户被锁定 30 分钟，用户无法登录

### 2. 管理员手动锁定
**场景**: 管理员因安全原因手动锁定账户
**例子**:
- 管理员发现账户异常，手动锁定
- 操作：`INSERT INTO auth.account_lockouts (user_id, tenant_id, locked_until, lock_reason, locked_by, actor_id) VALUES ('zhangsan-uuid', 'tenant-uuid', NULL, 'suspicious_activity', 'admin-uuid', 'admin-uuid')`
- 结果：账户永久锁定（`locked_until = NULL`），直到管理员手动解锁

### 3. 检查账户是否锁定
**场景**: 用户登录时检查账户锁定状态
**例子**:
- 用户尝试登录
- 查询：`SELECT locked_until, lock_reason FROM auth.account_lockouts WHERE user_id = 'zhangsan-uuid' AND (locked_until IS NULL OR locked_until > NOW())`
- 如果存在记录，拒绝登录，返回锁定原因

### 4. 自动解锁
**场景**: 临时锁定期满后自动解锁
**例子**:
- 定时任务检查过期的锁定
- 查询：`SELECT user_id FROM auth.account_lockouts WHERE locked_until IS NOT NULL AND locked_until <= NOW()`
- 操作：`UPDATE auth.account_lockouts SET unlocked_at = NOW(), unlocked_by = 'system' WHERE locked_until <= NOW() AND unlocked_at IS NULL`
- 结果：账户自动解锁

### 5. 管理员手动解锁
**场景**: 管理员手动解锁账户
**例子**:
- 管理员确认账户安全后解锁
- 操作：`UPDATE auth.account_lockouts SET unlocked_at = NOW(), unlocked_by = 'admin-uuid', unlock_actor_id = 'admin-uuid' WHERE user_id = 'zhangsan-uuid' AND unlocked_at IS NULL`
- 结果：账户立即解锁

