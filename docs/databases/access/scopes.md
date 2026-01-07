# Scopes Table

## 概述 / Overview

`access.scopes` 表定义了 OAuth/OIDC 语义层的权限，用于回答"这个 token 可以做什么？（用于 client_credentials / user token）"这个问题。

Scopes 是外部 API 契约，permissions 是内部能力点。例如："users:read" scope 可能映射到多个内部权限：users.read + users.count。或者你可以让 scope 和 permission 一对一（但长期扩展会很痛苦）。

The `access.scopes` table defines OAuth/OIDC semantic layer permissions, answering the question "What can this token do? (for client_credentials / user token)".

Scopes are external API contracts, permissions are internal capability points. Example: "users:read" scope may map to multiple internal permissions: users.read + users.count. Or you can make scope and permission 1:1 (but long-term expansion will be painful).

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `scope` | VARCHAR(255) | PRIMARY KEY | Scope 标识符，例如："users:read", "users:write", "assets:read" 等 |
| `description` | TEXT | NULL | Scope 的详细描述 |
| `is_system` | BOOLEAN | NOT NULL, DEFAULT false | 是否为系统内置 scope |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| `updated_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 更新时间 |
| `deleted_at` | TIMESTAMP | NULL | 软删除时间戳，NULL 表示未删除 |

## 字段详解 / Field Details

### `scope` (VARCHAR(255))
- **用途**: Scope 的唯一标识符，作为主键
- **命名规范**: 建议使用冒号分隔的格式，如 `resource:action` 或 `resource:action:subaction`
  - 示例: `users:read`, `users:write`, `users:export`, `assets:read`, `tenants:members:manage`
- **唯一性**: 全局唯一（主键）
- **稳定性**: 一旦创建，不应随意修改，以保持向后兼容性
- **OAuth 标准**: 遵循 OAuth 2.0 和 OIDC 的 scope 命名约定

### `description` (TEXT)
- **用途**: Scope 的详细描述，说明该 scope 的具体作用和使用场景
- **可为空**: 是
- **用途**: 帮助开发者理解 scope 的权限范围

### `is_system` (BOOLEAN)
- **用途**: 标识是否为系统内置 scope
- **系统 scope**: `true` 表示系统内置，通常不允许删除
- **自定义 scope**: `false` 表示用户或管理员创建的自定义 scope
- **默认值**: `false`

### `created_at` (TIMESTAMP)
- **用途**: 记录 scope 创建的时间
- **自动设置**: 插入时自动设置为当前时间

### `updated_at` (TIMESTAMP)
- **用途**: 记录 scope 最后更新的时间
- **自动设置**: 更新时自动设置为当前时间

### `deleted_at` (TIMESTAMP)
- **用途**: 软删除标记
- **NULL**: 表示 scope 未被删除
- **非 NULL**: 表示 scope 已被删除（软删除）
- **查询**: 查询时应过滤 `deleted_at IS NULL` 的记录

## 索引 / Indexes

1. **`idx_scopes_is_system`**: 在 `is_system` 字段上创建索引，用于筛选系统 scope
2. **`idx_scopes_deleted_at`**: 在 `deleted_at` 字段上创建索引，用于软删除查询优化

## 使用场景 / Use Cases

1. **OAuth Token 授权**: 在 OAuth 2.0 授权流程中，客户端请求特定的 scope
2. **Token 权限验证**: 验证 token 是否包含执行某个操作所需的 scope
3. **Scope 管理**: 在管理界面中展示所有可用 scope
4. **权限映射**: 与 `scope_permissions` 表关联，将 scope 映射到内部权限

## 关系 / Relationships

- **一对多**: 一个 scope 可以映射到多个内部权限（通过 `scope_permissions` 表）
- **多对一**: 多个 scope 可以映射到同一个权限（通过 `scope_permissions` 表）

## Scope vs Permission 的区别 / Scope vs Permission

### Scope（外部 API 契约）
- **用途**: 面向外部 API 的权限声明
- **使用场景**: OAuth token 中声明权限
- **粒度**: 通常较粗，一个 scope 可能包含多个操作
- **示例**: `users:read` 可能包含 `users.read`、`users.count`、`users.detail` 等多个权限

### Permission（内部能力点）
- **用途**: 系统内部的能力点定义
- **使用场景**: 内部授权检查
- **粒度**: 通常较细，一个权限对应一个具体操作
- **示例**: `users.read`、`users.count`、`users.detail` 是独立的权限

### 映射关系
- **一对多**: 一个 scope 可以映射到多个 permissions
- **灵活性**: 允许在不改变外部 API 契约的情况下，调整内部权限结构

## 查询示例 / Query Examples

```sql
-- 查询所有系统 scope
SELECT scope, description
FROM access.scopes
WHERE is_system = true AND deleted_at IS NULL;

-- 查询某个 scope 映射的所有权限
SELECT p.key, p.name, p.description
FROM access.scope_permissions sp
JOIN access.permissions p ON sp.permission_id = p.id
WHERE sp.scope = 'users:read' AND p.deleted_at IS NULL;

-- 检查 scope 是否存在
SELECT EXISTS(
  SELECT 1 FROM access.scopes
  WHERE scope = 'users:read' AND deleted_at IS NULL
);
```

## 示例数据 / Example Data

```sql
-- 系统内置 scope 示例
INSERT INTO access.scopes (scope, description, is_system) VALUES
  ('users:read', '读取用户信息', true),
  ('users:write', '创建和修改用户信息', true),
  ('users:export', '导出用户数据', true),
  ('assets:read', '读取资源信息', true),
  ('tenants:members:manage', '管理租户成员', true),
  ('clients:credentials:rotate', '轮换客户端凭证', true);
```

## 注意事项 / Notes

1. **Scope 的稳定性**: `scope` 字段一旦创建，不应随意修改，以保持向后兼容性
2. **系统 Scope 保护**: `is_system=true` 的 scope 通常不应被删除，应用层应进行保护
3. **软删除**: 使用软删除机制，保留历史数据用于审计
4. **OAuth 兼容性**: Scope 命名应遵循 OAuth 2.0 和 OIDC 标准
5. **粒度设计**: Scope 的粒度应适中，既不过细（增加管理复杂度），也不过粗（降低灵活性）
6. **权限映射**: 建议通过 `scope_permissions` 表建立 scope 与权限的映射关系，而不是硬编码

