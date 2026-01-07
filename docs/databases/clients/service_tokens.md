# Service Tokens Table

## 概述 / Overview

`clients.service_tokens` 表用于存储 M2M（机器对机器）访问令牌，存储已签发的服务令牌（如果是不透明令牌）或令牌元数据（如果是 JWT）。

支持令牌撤销、追踪和审计。

The `clients.service_tokens` table stores issued service tokens (if opaque) or token metadata (if JWT) for M2M access.

Supports token revocation, tracking, and audit.

## 使用场景 / Use Cases

### 1. 签发服务令牌
**场景**: 应用使用 client_credentials 获取服务令牌
**例子**:
- 应用使用 client_id 和 client_secret 获取服务令牌
- 系统签发 JWT 令牌
- 操作：`INSERT INTO clients.service_tokens (token_id, app_id, client_id, token_type, scopes, issued_at, expires_at, ip, ua_hash) VALUES ('jwt-jti-claim', 'app-uuid', 'client-id-123', 'jwt', ARRAY['users.read', 'users.count'], NOW(), NOW() + INTERVAL '1 hour', '203.0.113.1', 'ua-hash')`
- 结果：应用获得服务令牌，可以在 1 小时内使用

### 2. 验证服务令牌
**场景**: 应用使用服务令牌访问 API
**例子**:
- 应用在请求头中发送 `Authorization: Bearer jwt-token`
- 如果是 JWT，解析 JWT 获取 `jti` claim
- 查询：`SELECT status, expires_at, scopes FROM clients.service_tokens WHERE token_id = 'jwt-jti-claim' AND status = 'active'`
- 检查是否过期：`expires_at > NOW()`
- 如果有效，更新 `last_used_at`

### 3. 撤销服务令牌
**场景**: 发现令牌泄露，立即撤销
**例子**:
- 管理员发现令牌泄露
- 操作：`UPDATE clients.service_tokens SET status = 'revoked', revoked_at = NOW(), revoked_by = 'admin-uuid', revoke_reason = 'token_leaked' WHERE token_id = 'jwt-jti-claim'`
- 结果：令牌立即失效

### 4. 令牌使用追踪
**场景**: 追踪令牌的使用情况
**例子**:
- 查询令牌的使用统计
- 查询：`SELECT COUNT(*) as usage_count, MAX(last_used_at) as last_used FROM clients.service_tokens WHERE app_id = 'app-uuid' AND status = 'active'`
- 结果：显示令牌的使用次数和最后使用时间

