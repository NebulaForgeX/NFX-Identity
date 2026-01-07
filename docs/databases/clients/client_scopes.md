# Client Scopes Table

## 概述 / Overview

`clients.client_scopes` 表用于定义应用允许的权限范围（白名单），决定该应用可以请求哪些 scope。

在 OAuth 中：最终有效 scope = 请求的 scope ∩ 允许的 scope。

The `clients.client_scopes` table defines allowed permission scopes (Allow-list) for applications, determining which scopes this app can request.

In OAuth: final effective scope = requested scope ∩ allowed scope.

## 使用场景 / Use Cases

### 1. 授权应用 Scope
**场景**: 管理员为应用授权特定的 scope
**例子**:
- 为"数据分析服务"授权 `users.read` 和 `users.count` scope
- 操作：`INSERT INTO clients.client_scopes (app_id, scope, granted_by) VALUES ('app-uuid', 'users.read', 'admin-uuid'), ('app-uuid', 'users.count', 'admin-uuid')`
- 结果：应用只能请求这两个 scope，不能请求其他 scope

### 2. Scope 验证
**场景**: 应用请求 token 时，验证请求的 scope 是否被允许
**例子**:
- 应用请求 token，scope 为 `users.read,users.write`
- 查询：`SELECT scope FROM clients.client_scopes WHERE app_id = 'app-uuid' AND revoked_at IS NULL AND (expires_at IS NULL OR expires_at > NOW())`
- 允许的 scope：`['users.read', 'users.count']`
- 最终有效 scope：`users.read`（请求的 scope 与允许的 scope 的交集）

### 3. 临时 Scope 授权
**场景**: 临时授权 scope，过期后自动撤销
**例子**:
- 为应用临时授权 `users.export` scope，7 天后过期
- 操作：`INSERT INTO clients.client_scopes (app_id, scope, expires_at, granted_by) VALUES ('app-uuid', 'users.export', NOW() + INTERVAL '7 days', 'admin-uuid')`
- 结果：7 天后自动失效

### 4. 撤销 Scope
**场景**: 撤销应用的某个 scope 权限
**例子**:
- 管理员撤销应用的 `users.export` scope
- 操作：`UPDATE clients.client_scopes SET revoked_at = NOW(), revoked_by = 'admin-uuid', revoke_reason = 'no_longer_needed' WHERE app_id = 'app-uuid' AND scope = 'users.export'`
- 结果：应用无法再使用该 scope

