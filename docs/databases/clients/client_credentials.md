# Client Credentials Table

## 概述 / Overview

`clients.client_credentials` 表用于存储 OAuth 客户端凭证（client_id 和 secret hash），支持多版本轮换。

只存储哈希值，永远不存储明文（明文只在创建时返回一次）。

The `clients.client_credentials` table stores OAuth client_id and secret hash, supports multi-version rotation.

Only stores hash, never plaintext (plaintext only returned once during creation).

## 使用场景 / Use Cases

### 1. 创建客户端凭证
**场景**: 为新应用创建 OAuth 客户端凭证
**例子**:
- 为"数据分析服务"创建 OAuth 凭证
- 操作：`INSERT INTO clients.client_credentials (app_id, client_id, secret_hash, hash_alg, status, created_by) VALUES ('app-uuid', 'client-id-123', 'hashed-secret', 'argon2id', 'active', 'admin-uuid')`
- 系统返回 `client_id` 和 `client_secret`（明文，只返回一次）
- 结果：应用可以使用这些凭证进行 OAuth 认证

### 2. 凭证轮换
**场景**: 定期轮换客户端凭证，提高安全性
**例子**:
- 管理员决定轮换凭证
- 创建新凭证：`INSERT INTO clients.client_credentials (app_id, client_id, secret_hash, hash_alg, status, rotated_at) VALUES ('app-uuid', 'client-id-456', 'new-hashed-secret', 'argon2id', 'active', NOW())`
- 标记旧凭证为已轮换：`UPDATE clients.client_credentials SET status = 'rotating', rotated_at = NOW() WHERE client_id = 'client-id-123'`
- 结果：新凭证生效，旧凭证在过渡期后失效

### 3. 验证客户端凭证
**场景**: 应用使用 client_id 和 client_secret 进行认证
**例子**:
- 应用发送 `client_id` 和 `client_secret` 进行认证
- 查询：`SELECT secret_hash, hash_alg, status FROM clients.client_credentials WHERE client_id = 'client-id-123' AND status = 'active'`
- 对输入的 `client_secret` 进行哈希，与存储的哈希值比较
- 如果匹配，更新 `last_used_at`：`UPDATE clients.client_credentials SET last_used_at = NOW() WHERE client_id = 'client-id-123'`

### 4. 撤销凭证
**场景**: 发现凭证泄露，立即撤销
**例子**:
- 管理员发现凭证泄露
- 操作：`UPDATE clients.client_credentials SET status = 'revoked', revoked_at = NOW(), revoked_by = 'admin-uuid', revoke_reason = 'credential_leaked' WHERE client_id = 'client-id-123'`
- 结果：凭证立即失效，应用无法使用

