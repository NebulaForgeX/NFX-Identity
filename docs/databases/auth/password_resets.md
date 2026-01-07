# Password Resets Table

## 概述 / Overview

`auth.password_resets` 表用于安全的密码重置流程，支持审计、滥用防护和安全令牌管理。

The `auth.password_resets` table manages secure password recovery flow, supporting audit, abuse prevention, and secure token management.

## 使用场景 / Use Cases

### 1. 用户请求密码重置
**场景**: 用户忘记密码，请求重置
**例子**:
- 用户点击"忘记密码"，输入邮箱
- 系统生成重置令牌：`INSERT INTO auth.password_resets (reset_id, tenant_id, user_id, delivery, code_hash, expires_at, requested_ip, ua_hash) VALUES ('reset-uuid', 'tenant-uuid', 'zhangsan-uuid', 'email', 'hashed-code', NOW() + INTERVAL '1 hour', '203.0.113.1', 'ua-hash')`
- 发送邮件包含重置链接（包含 reset_id）
- 结果：用户收到重置邮件，1 小时内有效

### 2. 验证重置令牌
**场景**: 用户点击重置链接，验证令牌
**例子**:
- 用户点击邮件中的重置链接
- 查询：`SELECT * FROM auth.password_resets WHERE reset_id = 'reset-uuid' AND status = 'issued' AND expires_at > NOW()`
- 如果存在且未使用，允许用户重置密码
- 如果不存在或已过期，拒绝请求

### 3. 防止滥用
**场景**: 限制同一用户的密码重置请求频率
**例子**:
- 系统检测用户在过去 1 小时内已请求 3 次重置
- 查询：`SELECT COUNT(*) FROM auth.password_resets WHERE user_id = 'zhangsan-uuid' AND created_at >= NOW() - INTERVAL '1 hour' AND status = 'issued'`
- 如果 >= 3，拒绝新请求，防止滥用

### 4. 标记重置令牌为已使用
**场景**: 用户成功重置密码后，标记令牌为已使用
**例子**:
- 用户成功重置密码
- 操作：`UPDATE auth.password_resets SET status = 'used', used_at = NOW() WHERE reset_id = 'reset-uuid'`
- 结果：令牌失效，不能重复使用

### 5. 过期令牌清理
**场景**: 定期清理过期的重置令牌
**例子**:
- 定时任务清理过期令牌
- 查询：`SELECT reset_id FROM auth.password_resets WHERE expires_at < NOW() AND status = 'issued'`
- 操作：`UPDATE auth.password_resets SET status = 'expired' WHERE expires_at < NOW() AND status = 'issued'`
- 结果：过期令牌被标记为已过期

