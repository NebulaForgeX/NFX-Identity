# Scope Permissions Table

## 概述 / Overview

`access.scope_permissions` 表定义了 OAuth scope 与内部权限之间的映射关系，用于回答"一个 scope 包含哪些内部权限？"这个问题。

例如："users:read" scope 可能包含权限：users.read, users.count, users.detail。

The `access.scope_permissions` table maps OAuth scopes to internal permissions, answering the question "Which internal permissions does a scope include?"

Example: "users:read" scope may include permissions: users.read, users.count, users.detail.

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 映射关系的唯一标识符 |
| `scope` | VARCHAR(255) | NOT NULL, REFERENCES `access.scopes(scope)` ON DELETE CASCADE | Scope 标识符，外键引用 `access.scopes.scope` |
| `permission_id` | UUID | NOT NULL, REFERENCES `access.permissions(id)` ON DELETE CASCADE | 权限 ID，外键引用 `access.permissions.id` |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |

## 唯一约束 / Unique Constraints

- **`(scope, permission_id)`**: 确保同一个 scope 不能重复映射同一个权限

## 字段详解 / Field Details

### `id` (UUID)
- **用途**: 映射关系的主键标识符
- **生成方式**: 自动生成 UUID
- **唯一性**: 全局唯一

### `scope` (VARCHAR(255))
- **用途**: 关联的 scope 标识符
- **外键**: 引用 `access.scopes.scope`（注意：引用的是主键 `scope`，不是 `id`）
- **级联删除**: 当 scope 被删除时，该 scope 的所有权限映射也会被自动删除（CASCADE）
- **不可为空**: 是

### `permission_id` (UUID)
- **用途**: 关联的权限 ID
- **外键**: 引用 `access.permissions.id`
- **级联删除**: 当权限被删除时，所有包含该权限的 scope 映射也会被自动删除（CASCADE）
- **不可为空**: 是

### `created_at` (TIMESTAMP)
- **用途**: 记录权限被映射到 scope 的时间
- **自动设置**: 插入时自动设置为当前时间
- **审计用途**: 可用于追踪权限映射变更历史

## 索引 / Indexes

1. **`idx_scope_permissions_scope`**: 在 `scope` 字段上创建索引，用于快速查找 scope 的所有权限
2. **`idx_scope_permissions_permission_id`**: 在 `permission_id` 字段上创建索引，用于快速查找包含某个权限的所有 scope
3. **`idx_scope_permissions_scope_permission`**: 在 `(scope, permission_id)` 组合上创建唯一索引，用于快速检查权限是否已映射到 scope

## 使用场景 / Use Cases

1. **Scope 权限配置**: 为 scope 分配权限，定义 scope 的能力范围
2. **权限查询**: 查询某个 scope 包含的所有权限
3. **Scope 查询**: 查询包含某个权限的所有 scope
4. **Token 权限验证**: 在验证 token 权限时，通过 scope 查找其包含的权限
5. **权限映射管理**: 在管理界面中管理 scope 与权限的映射关系

## 关系 / Relationships

- **多对一**: 多个权限映射可以属于同一个 scope（通过 `scope`）
- **多对一**: 多个 scope 映射可以包含同一个权限（通过 `permission_id`）

## 工作流程 / Workflow

### OAuth Token 授权流程
1. 客户端请求 token 时指定 scope（如 `users:read`）
2. 系统通过 `scope_permissions` 表查找该 scope 包含的所有权限
3. Token 中包含这些权限信息（或通过 scope 间接引用）
4. 在 API 调用时，验证 token 的 scope 是否包含执行操作所需的权限

### 权限验证流程
1. 客户端使用 token 调用 API
2. API 检查操作所需的权限（如 `users.read`）
3. 系统查找 token 的 scope（如 `users:read`）
4. 通过 `scope_permissions` 表验证 scope 是否包含所需权限

## 查询示例 / Query Examples

```sql
-- 查询某个 scope 的所有权限
SELECT p.key, p.name, p.description
FROM access.scope_permissions sp
JOIN access.permissions p ON sp.permission_id = p.id
WHERE sp.scope = 'users:read' AND p.deleted_at IS NULL;

-- 查询包含某个权限的所有 scope
SELECT s.scope, s.description
FROM access.scope_permissions sp
JOIN access.scopes s ON sp.scope = s.scope
WHERE sp.permission_id = '...' AND s.deleted_at IS NULL;

-- 检查权限是否已映射到 scope
SELECT EXISTS(
  SELECT 1 FROM access.scope_permissions
  WHERE scope = 'users:read' AND permission_id = '...'
);

-- 验证 token scope 是否包含所需权限
SELECT EXISTS(
  SELECT 1 FROM access.scope_permissions sp
  WHERE sp.scope = ANY(ARRAY['users:read', 'users:write']) -- token 的 scopes
    AND sp.permission_id = (SELECT id FROM access.permissions WHERE key = 'users.read')
);
```

## 示例数据 / Example Data

```sql
-- 为 users:read scope 映射权限
INSERT INTO access.scope_permissions (scope, permission_id) VALUES
  (
    'users:read',
    (SELECT id FROM access.permissions WHERE key = 'users.read')
  ),
  (
    'users:read',
    (SELECT id FROM access.permissions WHERE key = 'users.count')
  ),
  (
    'users:read',
    (SELECT id FROM access.permissions WHERE key = 'users.detail')
  );

-- 为 users:write scope 映射权限
INSERT INTO access.scope_permissions (scope, permission_id) VALUES
  (
    'users:write',
    (SELECT id FROM access.permissions WHERE key = 'users.write')
  ),
  (
    'users:write',
    (SELECT id FROM access.permissions WHERE key = 'users.create')
  );
```

## 注意事项 / Notes

1. **唯一性约束**: 同一个 scope 不能重复映射同一个权限，数据库层面通过唯一约束保证
2. **级联删除**: 删除 scope 或权限时，映射关系会自动删除，无需手动清理
3. **一对多关系**: 一个 scope 可以映射到多个权限，这是设计上的灵活性
4. **性能考虑**: 查询 scope 的权限时，建议使用索引优化查询性能
5. **权限继承**: 当 token 包含某个 scope 时，token 自动获得该 scope 映射的所有权限
6. **映射变更**: 修改 scope 的权限映射时，需要考虑对现有 token 的影响

