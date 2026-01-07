# Permissions Table

## 概述 / Overview

`access.permissions` 表定义了系统的能力点（权限点），用于回答"系统中可以执行哪些操作？"这个问题。

权限点是稳定的产品契约：系统升级和新功能本质上就是添加/废弃权限点。

The `access.permissions` table defines system capability points (permission points), answering the question "What actions can be performed in the system?"

Permission points are stable product contracts: system upgrades and new features are essentially adding/deprecating permission points.

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 权限的唯一标识符 |
| `key` | VARCHAR(255) | NOT NULL, UNIQUE | 稳定的权限键值，例如："users.read", "users.count", "users.export", "assets.read", "tenants.members.manage", "clients.credentials.rotate" |
| `name` | VARCHAR(255) | NOT NULL | 权限的显示名称 |
| `description` | TEXT | NULL | 权限的详细描述 |
| `is_system` | BOOLEAN | NOT NULL, DEFAULT false | 是否为系统内置权限，系统内置权限不能被删除 |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| `updated_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 更新时间 |
| `deleted_at` | TIMESTAMP | NULL | 软删除时间戳，NULL 表示未删除 |

## 字段详解 / Field Details

### `id` (UUID)
- **用途**: 权限的主键标识符
- **生成方式**: 自动生成 UUID
- **唯一性**: 全局唯一

### `key` (VARCHAR(255))
- **用途**: 权限的稳定标识符，用于代码中引用
- **命名规范**: 建议使用点分隔的层级结构，如 `resource.action` 格式
  - 示例: `users.read`, `users.write`, `users.export`, `tenants.members.manage`
- **唯一性**: 全局唯一，不允许重复
- **稳定性**: 一旦创建，不应随意修改，以保持向后兼容性

### `name` (VARCHAR(255))
- **用途**: 权限的显示名称，用于管理界面展示
- **示例**: "读取用户", "导出用户数据", "管理租户成员"

### `description` (TEXT)
- **用途**: 权限的详细描述，说明该权限的具体作用和使用场景
- **可为空**: 是

### `is_system` (BOOLEAN)
- **用途**: 标识是否为系统内置权限
- **系统权限**: `true` 表示系统内置，通常不允许删除
- **自定义权限**: `false` 表示用户或管理员创建的自定义权限
- **默认值**: `false`

### `created_at` (TIMESTAMP)
- **用途**: 记录权限创建的时间
- **自动设置**: 插入时自动设置为当前时间

### `updated_at` (TIMESTAMP)
- **用途**: 记录权限最后更新的时间
- **自动设置**: 更新时自动设置为当前时间

### `deleted_at` (TIMESTAMP)
- **用途**: 软删除标记
- **NULL**: 表示权限未被删除
- **非 NULL**: 表示权限已被删除（软删除）
- **查询**: 查询时应过滤 `deleted_at IS NULL` 的记录

## 索引 / Indexes

1. **`idx_permissions_key`**: 在 `key` 字段上创建索引，用于快速查找权限
2. **`idx_permissions_is_system`**: 在 `is_system` 字段上创建索引，用于筛选系统权限
3. **`idx_permissions_deleted_at`**: 在 `deleted_at` 字段上创建索引，用于软删除查询优化

## 使用场景 / Use Cases

1. **权限定义**: 系统初始化时创建基础权限点
2. **权限查询**: 通过 `key` 快速查找权限信息
3. **权限管理**: 在管理界面中展示所有可用权限
4. **权限关联**: 与 `role_permissions` 和 `scope_permissions` 表关联，定义角色和 OAuth scope 的权限

## 关系 / Relationships

- **一对多**: 一个权限可以被多个角色关联（通过 `role_permissions` 表）
- **一对多**: 一个权限可以被多个 OAuth scope 关联（通过 `scope_permissions` 表）
- **一对多**: 一个权限可以直接授予给用户或客户端（通过 `grants` 表，`grant_type='PERMISSION'`）

## 示例数据 / Example Data

```sql
-- 系统内置权限示例
INSERT INTO access.permissions (key, name, description, is_system) VALUES
  ('users.read', '读取用户', '查看用户基本信息', true),
  ('users.write', '写入用户', '创建和修改用户信息', true),
  ('users.export', '导出用户', '导出用户数据', true),
  ('tenants.members.manage', '管理租户成员', '添加、删除、修改租户成员', true),
  ('clients.credentials.rotate', '轮换客户端凭证', '轮换客户端密钥和 API Key', true);
```

## 注意事项 / Notes

1. **权限键的稳定性**: `key` 字段一旦创建，不应随意修改，以保持向后兼容性
2. **系统权限保护**: `is_system=true` 的权限通常不应被删除，应用层应进行保护
3. **软删除**: 使用软删除机制，保留历史数据用于审计
4. **权限粒度**: 权限点应设计得足够细粒度，以便灵活组合

