# Roles Table

## 概述 / Overview

`access.roles` 表定义了角色的集合，用于回答"系统中存在哪些职位/身份？"这个问题。

角色是为了人类可读性和操作便利性而存在的命名包装器。企业客户通常需要"自定义角色"（或基于模板修改）。

The `access.roles` table defines role collections, answering the question "What positions/identities exist in the system?"

Roles are named wrappers for human readability and operational convenience. Enterprise customers often need "custom roles" (or modify based on templates).

## 表结构 / Table Structure

| 字段名 / Field | 类型 / Type | 约束 / Constraints | 说明 / Description |
|---------------|------------|-------------------|-------------------|
| `id` | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() | 角色的唯一标识符 |
| `key` | VARCHAR(255) | NOT NULL, UNIQUE | 唯一的角色键值，例如："tenant.owner", "tenant.admin", "tenant.viewer", "app.operator", "app.support", "service.reader", "service.writer" |
| `name` | VARCHAR(255) | NOT NULL | 角色的显示名称 |
| `description` | TEXT | NULL | 角色的详细描述 |
| `scope_type` | `access.scope_type` ENUM | NOT NULL, DEFAULT 'TENANT' | 角色生效范围：TENANT（租户级）、APP（应用级）、GLOBAL（全局） |
| `is_system` | BOOLEAN | NOT NULL, DEFAULT false | 是否为系统内置角色 |
| `created_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 创建时间 |
| `updated_at` | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP | 更新时间 |
| `deleted_at` | TIMESTAMP | NULL | 软删除时间戳，NULL 表示未删除 |

## 枚举类型 / Enum Types

### `access.scope_type`
- **`TENANT`**: 租户级角色，仅在特定租户内生效
- **`APP`**: 应用级角色，仅在特定应用内生效
- **`GLOBAL`**: 全局角色，跨租户和应用生效

## 字段详解 / Field Details

### `id` (UUID)
- **用途**: 角色的主键标识符
- **生成方式**: 自动生成 UUID
- **唯一性**: 全局唯一

### `key` (VARCHAR(255))
- **用途**: 角色的稳定标识符，用于代码中引用
- **命名规范**: 建议使用点分隔的层级结构，如 `scope.role_name` 格式
  - 租户级: `tenant.owner`, `tenant.admin`, `tenant.viewer`, `tenant.member`
  - 应用级: `app.operator`, `app.support`, `app.developer`
  - 全局: `service.reader`, `service.writer`, `system.admin`
- **唯一性**: 全局唯一，不允许重复
- **稳定性**: 一旦创建，不应随意修改，以保持向后兼容性

### `name` (VARCHAR(255))
- **用途**: 角色的显示名称，用于管理界面展示
- **示例**: "租户所有者", "租户管理员", "应用操作员"

### `description` (TEXT)
- **用途**: 角色的详细描述，说明该角色的职责和使用场景
- **可为空**: 是

### `scope_type` (`access.scope_type` ENUM)
- **用途**: 定义角色的作用范围
- **TENANT**: 租户级角色，通常与 `tenants.member_roles` 表关联使用
- **APP**: 应用级角色，通常与 `tenants.member_app_roles` 表关联使用
- **GLOBAL**: 全局角色，通常用于系统级或跨租户的权限管理
- **默认值**: `'TENANT'`

### `is_system` (BOOLEAN)
- **用途**: 标识是否为系统内置角色
- **系统角色**: `true` 表示系统内置，通常不允许删除
- **自定义角色**: `false` 表示用户或管理员创建的自定义角色
- **默认值**: `false`

### `created_at` (TIMESTAMP)
- **用途**: 记录角色创建的时间
- **自动设置**: 插入时自动设置为当前时间

### `updated_at` (TIMESTAMP)
- **用途**: 记录角色最后更新的时间
- **自动设置**: 更新时自动设置为当前时间

### `deleted_at` (TIMESTAMP)
- **用途**: 软删除标记
- **NULL**: 表示角色未被删除
- **非 NULL**: 表示角色已被删除（软删除）
- **查询**: 查询时应过滤 `deleted_at IS NULL` 的记录

## 索引 / Indexes

1. **`idx_roles_key`**: 在 `key` 字段上创建索引，用于快速查找角色
2. **`idx_roles_scope_type`**: 在 `scope_type` 字段上创建索引，用于按作用范围筛选角色
3. **`idx_roles_is_system`**: 在 `is_system` 字段上创建索引，用于筛选系统角色
4. **`idx_roles_deleted_at`**: 在 `deleted_at` 字段上创建索引，用于软删除查询优化

## 使用场景 / Use Cases

1. **角色定义**: 系统初始化时创建基础角色
2. **角色查询**: 通过 `key` 或 `scope_type` 快速查找角色
3. **角色管理**: 在管理界面中展示所有可用角色
4. **权限关联**: 与 `role_permissions` 表关联，定义角色包含的权限
5. **授权关联**: 与 `grants` 表关联，将角色授予给用户或客户端

## 关系 / Relationships

- **一对多**: 一个角色可以包含多个权限（通过 `role_permissions` 表）
- **一对多**: 一个角色可以授予给多个用户或客户端（通过 `grants` 表，`grant_type='ROLE'`）
- **一对多**: 一个角色可以分配给多个租户成员（通过 `tenants.member_roles` 表，当 `scope_type='TENANT'` 时）
- **一对多**: 一个角色可以分配给多个应用成员（通过 `tenants.member_app_roles` 表，当 `scope_type='APP'` 时）

## 示例数据 / Example Data

```sql
-- 系统内置角色示例
INSERT INTO access.roles (key, name, description, scope_type, is_system) VALUES
  ('tenant.owner', '租户所有者', '拥有租户的完全控制权', 'TENANT', true),
  ('tenant.admin', '租户管理员', '管理租户成员和设置', 'TENANT', true),
  ('tenant.viewer', '租户查看者', '只能查看租户信息', 'TENANT', true),
  ('app.operator', '应用操作员', '管理应用的日常操作', 'APP', true),
  ('app.support', '应用支持', '提供应用技术支持', 'APP', true),
  ('service.reader', '服务读取者', '全局只读权限', 'GLOBAL', true),
  ('service.writer', '服务写入者', '全局读写权限', 'GLOBAL', true);
```

## 注意事项 / Notes

1. **角色键的稳定性**: `key` 字段一旦创建，不应随意修改，以保持向后兼容性
2. **系统角色保护**: `is_system=true` 的角色通常不应被删除，应用层应进行保护
3. **作用范围**: `scope_type` 决定了角色的使用场景，选择时需谨慎
4. **软删除**: 使用软删除机制，保留历史数据用于审计
5. **角色与权限**: 角色本身不包含权限，需要通过 `role_permissions` 表关联权限

