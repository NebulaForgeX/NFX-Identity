# Refresh Tokens Table

## 概述 / Overview

`auth.refresh_tokens` 表用于存储长期会话的刷新令牌，支持令牌轮换和撤销，满足企业级安全要求。

The `auth.refresh_tokens` table stores refresh tokens for long-lived sessions with rotation support, supporting token rotation and revocation for enterprise security.

## 使用场景 / Use Cases

### 1. 用户登录获取刷新令牌
**场景**: 用户登录后获得访问令牌和刷新令牌
**例子**:
- 用户"张三"登录，系统生成访问令牌（1 小时过期）和刷新令牌（7 天过期）
- 操作：`INSERT INTO auth.refresh_tokens (token_id, user_id, tenant_id, app_id, expires_at, ip, ua_hash, device_id) VALUES ('refresh-token-uuid', 'zhangsan-uuid', 'tenant-uuid', 'app-uuid', NOW() + INTERVAL '7 days', '203.0.113.1', 'ua-hash', 'device-uuid')`
- 结果：用户可以使用刷新令牌获取新的访问令牌

### 2. 令牌轮换（Token Rotation）
**场景**: 使用刷新令牌获取新访问令牌时，轮换刷新令牌
**例子**:
- 用户使用刷新令牌获取新访问令牌
- 系统创建新刷新令牌：`INSERT INTO auth.refresh_tokens (token_id, user_id, tenant_id, app_id, expires_at, rotated_from, ip) VALUES ('new-refresh-token-uuid', 'zhangsan-uuid', 'tenant-uuid', 'app-uuid', NOW() + INTERVAL '7 days', 'old-refresh-token-uuid', '203.0.113.1')`
- 撤销旧令牌：`UPDATE auth.refresh_tokens SET revoked_at = NOW(), revoke_reason = 'rotation' WHERE token_id = 'old-refresh-token-uuid'`
- 结果：旧令牌失效，新令牌生效，提高安全性

### 3. 用户登出撤销令牌
**场景**: 用户主动登出，撤销所有刷新令牌
**例子**:
- 用户点击"登出"
- 操作：`UPDATE auth.refresh_tokens SET revoked_at = NOW(), revoke_reason = 'user_logout' WHERE user_id = 'zhangsan-uuid' AND revoked_at IS NULL`
- 结果：用户的所有刷新令牌失效，需要重新登录

### 4. 检测可疑活动撤销令牌
**场景**: 检测到可疑活动，自动撤销相关令牌
**例子**:
- 系统检测到用户从异常 IP 登录
- 操作：`UPDATE auth.refresh_tokens SET revoked_at = NOW(), revoke_reason = 'suspicious_activity' WHERE user_id = 'zhangsan-uuid' AND ip != '203.0.113.1' AND revoked_at IS NULL`
- 结果：异常 IP 的令牌被撤销，用户需要重新验证身份

### 5. 密码更改撤销令牌
**场景**: 用户更改密码后，撤销所有刷新令牌
**例子**:
- 用户更改密码
- 操作：`UPDATE auth.refresh_tokens SET revoked_at = NOW(), revoke_reason = 'password_changed' WHERE user_id = 'zhangsan-uuid' AND revoked_at IS NULL`
- 结果：所有设备需要重新登录，提高安全性

