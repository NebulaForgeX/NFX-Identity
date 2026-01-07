# API Keys Table

## 概述 / Overview

`clients.api_keys` 表用于存储简化的 API 密钥凭证，用于脚本/内部系统/低复杂度集成商。

比 OAuth 风险更高：直接访问，需要更强的约束（IP/权限/过期时间）。

The `clients.api_keys` table stores simplified API key credentials for scripts/internal systems/low-complexity integrators.

Higher risk than OAuth: direct access, requires stronger constraints (IP/permissions/expiration).

## 使用场景 / Use Cases

### 1. 创建 API Key
**场景**: 为脚本或内部系统创建 API Key
**例子**:
- 为"批量数据导入脚本"创建 API Key
- 操作：`INSERT INTO clients.api_keys (key_id, app_id, key_hash, hash_alg, name, expires_at, created_by) VALUES ('key-prefix-123', 'app-uuid', 'hashed-key', 'argon2id', '批量数据导入脚本', NOW() + INTERVAL '90 days', 'admin-uuid')`
- 系统返回 API Key（明文，只返回一次）
- 结果：脚本可以使用该 API Key 访问 API

### 2. 验证 API Key
**场景**: 应用使用 API Key 进行认证
**例子**:
- 应用在请求头中发送 `Authorization: Bearer api-key-xxx`
- 查询：`SELECT key_hash, hash_alg, status, expires_at FROM clients.api_keys WHERE key_id = 'key-prefix-123' AND status = 'active'`
- 检查是否过期：`expires_at > NOW()`
- 对输入的 API Key 进行哈希，与存储的哈希值比较
- 如果匹配，更新 `last_used_at`

### 3. 设置过期时间
**场景**: 为 API Key 设置过期时间，提高安全性
**例子**:
- 创建临时 API Key，30 天后过期
- 操作：`INSERT INTO clients.api_keys (key_id, app_id, key_hash, name, expires_at) VALUES ('temp-key-123', 'app-uuid', 'hashed-key', '临时访问密钥', NOW() + INTERVAL '30 days')`
- 结果：30 天后自动失效

### 4. 撤销 API Key
**场景**: 发现 API Key 泄露，立即撤销
**例子**:
- 管理员发现 API Key 泄露
- 操作：`UPDATE clients.api_keys SET status = 'revoked', revoked_at = NOW(), revoked_by = 'admin-uuid', revoke_reason = 'key_leaked' WHERE key_id = 'key-prefix-123'`
- 结果：API Key 立即失效

