# Sessions Table

## 概述 / Overview

`auth.sessions` 表用于设备/会话管理，支持：踢下线、查看活跃设备、限制并发会话、异常设备警报等功能。

The `auth.sessions` table manages device/session information, supporting: kick offline, view active devices, limit concurrent sessions, abnormal device alerts.

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY | 会话记录的唯一标识符 |
| `session_id` | VARCHAR(255) | NOT NULL, UNIQUE | 唯一的会话标识符 |
| `tenant_id` | UUID | NOT NULL | 多租户隔离 |
| `user_id` | UUID | NOT NULL | 用户 ID（引用 directory.users.id，应用级一致性） |
| `app_id` | UUID | NULL | 该会话所属的应用 |
| `client_id` | VARCHAR(255) | NULL | OAuth 客户端标识符 |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 会话创建时间 |
| `last_seen_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 最后活动时间 |
| `expires_at` | TIMESTAMP | NOT NULL | 会话过期时间 |
| `ip` | INET | NULL | IP 地址 |
| `ua_hash` | VARCHAR(64) | NULL | 用户代理哈希值 |
| `device_id` | VARCHAR(255) | NULL | 设备标识符 |
| `device_fingerprint` | VARCHAR(255) | NULL | 设备指纹（用于风险控制） |
| `device_name` | VARCHAR(255) | NULL | 用户友好的设备名称 |
| `revoked_at` | TIMESTAMP | NULL | 撤销时间 |
| `revoke_reason` | `auth.session_revoke_reason` ENUM | NULL | 撤销原因 |
| `revoked_by` | VARCHAR(255) | NULL | 撤销者：'user' 或管理员 user_id/username |
| `updated_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 更新时间 |

## 使用场景 / Use Cases

### 1. 用户登录创建会话
**场景**: 用户成功登录后创建会话
**例子**:
- 用户"张三"在 Chrome 浏览器登录
- 操作：`INSERT INTO auth.sessions (session_id, tenant_id, user_id, app_id, expires_at, ip, ua_hash, device_id, device_fingerprint, device_name) VALUES ('session-uuid', 'tenant-uuid', 'zhangsan-uuid', 'app-uuid', NOW() + INTERVAL '7 days', '203.0.113.1', 'ua-hash', 'device-uuid', 'fingerprint-hash', 'Chrome on Windows')`
- 结果：用户获得会话，可以在 7 天内使用该会话访问系统

### 2. 查看用户活跃会话
**场景**: 用户查看自己的所有活跃设备/会话
**例子**:
- 用户在"安全设置"中查看"活跃设备"
- 查询：`SELECT session_id, device_name, ip, created_at, last_seen_at, expires_at FROM auth.sessions WHERE user_id = 'zhangsan-uuid' AND revoked_at IS NULL AND expires_at > NOW() ORDER BY last_seen_at DESC`
- 结果：显示所有活跃会话，包括设备名称、IP、最后活动时间等

### 3. 踢下线（撤销会话）
**场景**: 用户或管理员撤销某个会话
**例子**:
- 用户发现异常设备，选择"踢下线"
- 操作：`UPDATE auth.sessions SET revoked_at = NOW(), revoke_reason = 'user_logout', revoked_by = 'user' WHERE session_id = 'suspicious-session-uuid'`
- 结果：该会话立即失效，用户需要重新登录

### 4. 限制并发会话
**场景**: 限制用户同时只能有 N 个活跃会话
**例子**:
- 系统限制每个用户最多 5 个并发会话
- 用户登录时，查询：`SELECT COUNT(*) FROM auth.sessions WHERE user_id = 'zhangsan-uuid' AND revoked_at IS NULL AND expires_at > NOW()`
- 如果已有 5 个会话，撤销最旧的会话：`UPDATE auth.sessions SET revoked_at = NOW(), revoke_reason = 'session_expired' WHERE user_id = 'zhangsan-uuid' AND revoked_at IS NULL ORDER BY last_seen_at ASC LIMIT 1`

### 5. 异常设备检测
**场景**: 检测异常设备登录
**例子**:
- 系统检测到用户从新设备登录
- 查询：`SELECT COUNT(*) FROM auth.sessions WHERE user_id = 'zhangsan-uuid' AND device_fingerprint = 'new-fingerprint'`
- 如果是新设备，发送通知或要求 MFA 验证

### 6. 会话自动过期
**场景**: 定期清理过期会话
**例子**:
- 定时任务每天清理过期会话
- 查询：`SELECT session_id FROM auth.sessions WHERE expires_at < NOW() AND revoked_at IS NULL`
- 操作：`UPDATE auth.sessions SET revoked_at = NOW(), revoke_reason = 'session_expired' WHERE expires_at < NOW() AND revoked_at IS NULL`

### 7. 更新最后活动时间
**场景**: 用户每次请求时更新会话活动时间
**例子**:
- 用户访问 API 时，系统更新会话活动时间
- 操作：`UPDATE auth.sessions SET last_seen_at = NOW(), updated_at = NOW() WHERE session_id = 'session-uuid' AND revoked_at IS NULL AND expires_at > NOW()`
- 结果：保持会话活跃，避免因长时间不活动而过期

### 8. 管理员撤销所有会话
**场景**: 管理员撤销用户的所有会话（如账户被锁定）
**例子**:
- 管理员锁定用户账户
- 操作：`UPDATE auth.sessions SET revoked_at = NOW(), revoke_reason = 'account_locked', revoked_by = 'admin-uuid' WHERE user_id = 'zhangsan-uuid' AND revoked_at IS NULL`
- 结果：用户的所有会话立即失效

