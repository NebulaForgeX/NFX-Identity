# Role Permissions Table

## 概述 / Overview

`access.role_permissions` 表定义了角色与权限之间的多对多关系，用于回答"某个角色（如 tenant.admin）实际上可以做什么？"这个问题。

企业需要"版本演进"能力：当权限被添加到管理员角色时，必须可追溯。

The `access.role_permissions` table defines the many-to-many relationship between roles and permissions, answering the question "What can a role (e.g., tenant.admin) actually do?"

Enterprise needs "version evolution" capability: when a permission is added to a role, it must be traceable.

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 关联关系的唯一标识符 |
| `role_id` | UUID | NOT NULL, REFERENCES `access.roles(id)` ON DELETE CASCADE | 角色 ID，外键引用 `access.roles.id` |
| `permission_id` | UUID | NOT NULL, REFERENCES `access.permissions(id)` ON DELETE CASCADE | 权限 ID，外键引用 `access.permissions.id` |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| `created_by` | UUID | NULL | 谁添加了这个权限到角色（用于审计），通常是管理员用户 ID |

## 唯一约束 / Unique Constraints

- **`(role_id, permission_id)`**: 确保同一个角色不能重复添加同一个权限

## 字段详解 / Field Details

### `id` (UUID)
- **用途**: 关联关系的主键标识符
- **生成方式**: 自动生成 UUID
- **唯一性**: 全局唯一

### `role_id` (UUID)
- **用途**: 关联的角色 ID
- **外键**: 引用 `access.roles.id`
- **级联删除**: 当角色被删除时，该角色的所有权限关联也会被自动删除（CASCADE）
- **不可为空**: 是

### `permission_id` (UUID)
- **用途**: 关联的权限 ID
- **外键**: 引用 `access.permissions.id`
- **级联删除**: 当权限被删除时，所有包含该权限的角色关联也会被自动删除（CASCADE）
- **不可为空**: 是

### `created_at` (TIMESTAMP)
- **用途**: 记录权限被添加到角色的时间
- **自动设置**: 插入时自动设置为当前时间
- **审计用途**: 可用于追踪权限变更历史

### `created_by` (UUID)
- **用途**: 记录是谁将权限添加到角色
- **引用**: 通常是管理员用户 ID（引用 `directory.users.id`，应用级一致性）
- **审计用途**: 用于审计追踪，了解权限变更的负责人
- **可为空**: 是（系统自动添加时可能为空）

## 索引 / Indexes

1. **`idx_role_permissions_role_id`**: 在 `role_id` 字段上创建索引，用于快速查找角色的所有权限
2. **`idx_role_permissions_permission_id`**: 在 `permission_id` 字段上创建索引，用于快速查找包含某个权限的所有角色
3. **`idx_role_permissions_role_permission`**: 在 `(role_id, permission_id)` 组合上创建唯一索引，用于快速检查权限是否已关联到角色

## 使用场景 / Use Cases

1. **角色权限配置**: 为角色分配权限，定义角色的能力范围
2. **权限查询**: 查询某个角色包含的所有权限
3. **角色查询**: 查询包含某个权限的所有角色
4. **权限变更追踪**: 通过 `created_at` 和 `created_by` 追踪权限变更历史
5. **权限验证**: 在授权检查时，通过角色查找其包含的权限

## 关系 / Relationships

- **多对一**: 多个权限关联可以属于同一个角色（通过 `role_id`）
- **多对一**: 多个角色关联可以包含同一个权限（通过 `permission_id`）
- **多对一**: 关联关系由某个用户创建（通过 `created_by`，应用级一致性）

## 查询示例 / Query Examples

```sql
-- 查询某个角色的所有权限
SELECT p.key, p.name, p.description
FROM access.role_permissions rp
JOIN access.permissions p ON rp.permission_id = p.id
WHERE rp.role_id = '...' AND p.deleted_at IS NULL;

-- 查询包含某个权限的所有角色
SELECT r.key, r.name, r.scope_type
FROM access.role_permissions rp
JOIN access.roles r ON rp.role_id = r.id
WHERE rp.permission_id = '...' AND r.deleted_at IS NULL;

-- 检查权限是否已关联到角色
SELECT EXISTS(
  SELECT 1 FROM access.role_permissions
  WHERE role_id = '...' AND permission_id = '...'
);
```

## 示例数据 / Example Data

```sql
-- 为 tenant.admin 角色添加权限
INSERT INTO access.role_permissions (role_id, permission_id, created_by) VALUES
  (
    (SELECT id FROM access.roles WHERE key = 'tenant.admin'),
    (SELECT id FROM access.permissions WHERE key = 'users.read'),
    'admin-user-id'
  ),
  (
    (SELECT id FROM access.roles WHERE key = 'tenant.admin'),
    (SELECT id FROM access.permissions WHERE key = 'users.write'),
    'admin-user-id'
  ),
  (
    (SELECT id FROM access.roles WHERE key = 'tenant.admin'),
    (SELECT id FROM access.permissions WHERE key = 'tenants.members.manage'),
    'admin-user-id'
  );
```

## 注意事项 / Notes

1. **唯一性约束**: 同一个角色不能重复添加同一个权限，数据库层面通过唯一约束保证
2. **级联删除**: 删除角色或权限时，关联关系会自动删除，无需手动清理
3. **审计追踪**: `created_by` 字段用于审计，建议在应用层始终设置该值
4. **权限继承**: 当用户被授予角色时，用户自动获得该角色包含的所有权限
5. **性能考虑**: 查询角色的权限时，建议使用索引优化查询性能

