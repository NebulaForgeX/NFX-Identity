# Credential Events Table

## 概述 / Overview

`auth.credential_events` 表用于记录凭据变更的审计轨迹，包括密码更改、MFA 绑定、凭据禁用等，用于安全审计。

The `auth.credential_events` table records credential change audit trail, including password changes, MFA bindings, credential disables, etc. for security audit.

## 使用场景 / Use Cases

### 1. 记录密码更改
**场景**: 用户更改密码时记录事件
**例子**:
- 用户更改密码
- 操作：`INSERT INTO auth.credential_events (event_id, tenant_id, user_id, event_type, actor_id, actor_type, metadata) VALUES ('event-uuid', 'tenant-uuid', 'zhangsan-uuid', 'password_changed', 'zhangsan-uuid', 'user', '{"ip": "203.0.113.1", "device": "Chrome"}')`
- 结果：记录密码更改事件，用于审计

### 2. 记录 MFA 启用
**场景**: 用户启用 MFA 时记录事件
**例子**:
- 用户启用 TOTP
- 操作：`INSERT INTO auth.credential_events (event_id, tenant_id, user_id, event_type, actor_id, actor_type, metadata) VALUES ('event-uuid', 'tenant-uuid', 'zhangsan-uuid', 'mfa_enabled', 'zhangsan-uuid', 'user', '{"factor_type": "totp"}')`
- 结果：记录 MFA 启用事件

### 3. 管理员操作审计
**场景**: 管理员重置用户密码时记录事件
**例子**:
- 管理员重置用户密码
- 操作：`INSERT INTO auth.credential_events (event_id, tenant_id, user_id, event_type, actor_id, actor_type, metadata) VALUES ('event-uuid', 'tenant-uuid', 'zhangsan-uuid', 'password_reset', 'admin-uuid', 'admin', '{"ip": "203.0.113.2", "reason": "user_request"}')`
- 结果：记录管理员操作，用于审计追踪

### 4. 查询凭据变更历史
**场景**: 安全团队查询用户的凭据变更历史
**例子**:
- 安全团队调查用户账户异常
- 查询：`SELECT event_type, actor_type, created_at, metadata FROM auth.credential_events WHERE user_id = 'zhangsan-uuid' ORDER BY created_at DESC LIMIT 20`
- 结果：显示用户的所有凭据变更历史，包括时间、操作者、IP 等

