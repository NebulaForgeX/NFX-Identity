# Access Schema 总览 / Access Schema Overview

## 概述 / Overview

`access` schema 是 NFX-Identity 平台的授权（Authorization）核心模块，负责管理权限、角色、OAuth scope、策略和授权关系。

The `access` schema is the authorization core module of the NFX-Identity platform, responsible for managing permissions, roles, OAuth scopes, policies, and authorization relationships.

## 核心概念 / Core Concepts

### 1. 权限（Permissions）
- **定义**: 系统能力点，定义"可以做什么"
- **表**: `access.permissions`
- **特点**: 稳定的产品契约，系统升级和新功能本质上就是添加/废弃权限点

### 2. 角色（Roles）
- **定义**: 权限集合的命名包装器，定义"存在哪些职位/身份"
- **表**: `access.roles`
- **特点**: 为了人类可读性和操作便利性，企业客户通常需要自定义角色

### 3. OAuth Scope
- **定义**: OAuth/OIDC 语义层权限，定义"token 可以做什么"
- **表**: `access.scopes`
- **特点**: 外部 API 契约，与内部权限（permissions）分离

### 4. 策略（Policies）
- **定义**: ABAC 条件授权策略，定义"在什么条件下允许/拒绝访问"
- **表**: `access.policies`
- **特点**: 用于处理复杂的条件授权场景

### 5. 授权（Grants）
- **定义**: 核心授权表，定义"谁在什么范围内被授予了什么"
- **表**: `access.grants`
- **特点**: 支持角色授权、权限授权、租户级授权、应用级授权、资源级授权

## 表关系图 / Table Relationships

```
┌─────────────────┐
│   permissions   │
│  (权限定义)     │
└────────┬────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────────────────┐
    │      role_permissions             │
    │  (角色-权限关联)                   │
    └────┬──────────────────────────────┘
         │
         │ N:1
         │
┌────────▼────────┐              ┌─────────────────┐
│     roles       │              │    scopes       │
│   (角色定义)    │              │  (OAuth Scope)   │
└────────┬────────┘              └────────┬─────────┘
         │                               │
         │ 1:N                           │ 1:N
         │                               │
    ┌────▼──────────────────────────┐   │
    │          grants               │   │
    │    (授权关系-角色)            │   │
    └────┬──────────────────────────┘   │
         │                               │
         │ N:1                           │
         │                               │
    ┌────▼──────────────────────────┐   │
    │          grants               │   │
    │  (授权关系-权限)              │   │
    └───────────────────────────────┘   │
                                        │
                                    ┌───▼──────────────────┐
                                    │  scope_permissions   │
                                    │ (Scope-权限映射)      │
                                    └──────────────────────┘

┌─────────────────┐
│    policies     │
│  (ABAC 策略)    │
│  (独立表)       │
└─────────────────┘
```

## 表列表 / Table List

### 1. `access.permissions` - 权限表
- **用途**: 定义系统能力点
- **关键字段**: `key`（权限键值，如 "users.read"）、`is_system`（系统内置标识）
- **详细文档**: [permissions.md](./permissions.md)

### 2. `access.roles` - 角色表
- **用途**: 定义角色集合
- **关键字段**: `key`（角色键值，如 "tenant.admin"）、`scope_type`（作用范围：TENANT/APP/GLOBAL）
- **详细文档**: [roles.md](./roles.md)

### 3. `access.role_permissions` - 角色权限关联表
- **用途**: 定义角色与权限的多对多关系
- **关键字段**: `role_id`、`permission_id`、`created_by`（审计追踪）
- **详细文档**: [role_permissions.md](./role_permissions.md)

### 4. `access.scopes` - OAuth Scope 表
- **用途**: 定义 OAuth/OIDC 语义层权限
- **关键字段**: `scope`（主键，如 "users:read"）、`is_system`（系统内置标识）
- **详细文档**: [scopes.md](./scopes.md)

### 5. `access.scope_permissions` - Scope 权限映射表
- **用途**: 映射 OAuth scope 到内部权限
- **关键字段**: `scope`、`permission_id`
- **详细文档**: [scope_permissions.md](./scope_permissions.md)

### 6. `access.policies` - 策略表
- **用途**: 定义 ABAC 条件授权策略
- **关键字段**: `effect`（ALLOW/DENY）、`priority`（优先级）、`condition`（JSONB 条件表达式）
- **详细文档**: [policies.md](./policies.md)

### 7. `access.grants` - 授权表
- **用途**: 核心授权表，定义谁被授予了什么
- **关键字段**: `subject_type`（USER/CLIENT）、`grant_type`（ROLE/PERMISSION）、`tenant_id`、`app_id`、`resource_type`、`resource_id`
- **详细文档**: [grants.md](./grants.md)

## 数据流 / Data Flow

### 1. 权限定义流程
```
1. 系统初始化时创建基础权限（permissions）
2. 创建角色（roles）并关联权限（role_permissions）
3. 创建 OAuth scope（scopes）并映射到权限（scope_permissions）
```

### 2. 授权流程
```
1. 管理员创建授权（grants）
   - 选择主体（用户或客户端）
   - 选择授权类型（角色或权限）
   - 设置作用范围（租户、应用、资源）
2. 系统查询授权
   - 根据主体查找所有授权
   - 如果是角色授权，查找角色包含的权限
   - 应用作用范围过滤
3. 权限验证
   - 检查操作所需的权限
   - 验证主体是否拥有该权限
   - 应用策略（policies）进行条件判断
```

### 3. OAuth Token 授权流程
```
1. 客户端请求 token 时指定 scope
2. 系统通过 scope_permissions 查找 scope 包含的权限
3. Token 中包含权限信息
4. API 调用时验证 token 的权限
```

## 授权模型 / Authorization Models

### 1. RBAC（基于角色的访问控制）
- **实现**: 通过 `roles`、`role_permissions`、`grants` 表实现
- **流程**: 用户被授予角色 → 角色包含权限 → 用户获得权限

### 2. ABAC（基于属性的访问控制）
- **实现**: 通过 `policies` 表实现
- **流程**: 定义策略条件 → 评估请求上下文 → 应用策略结果

### 3. 混合模型
- **实现**: RBAC + ABAC 结合
- **流程**: 先通过 RBAC 确定基础权限 → 再通过 ABAC 策略进行条件判断

## 作用范围 / Scopes

### 1. 全局范围（Global）
- **定义**: `grants.tenant_id IS NULL AND grants.app_id IS NULL`
- **用途**: 跨租户和应用的系统级权限

### 2. 租户范围（Tenant）
- **定义**: `grants.tenant_id IS NOT NULL`
- **用途**: 在特定租户内生效的权限

### 3. 应用范围（App）
- **定义**: `grants.app_id IS NOT NULL`
- **用途**: 在特定应用内生效的权限

### 4. 资源范围（Resource）
- **定义**: `grants.resource_type IS NOT NULL AND grants.resource_id IS NOT NULL`
- **用途**: 在特定资源上生效的权限

## 查询示例 / Query Examples

### 查询用户的所有权限
```sql
-- 1. 查找用户的所有授权
WITH user_grants AS (
  SELECT grant_type, grant_ref_id, tenant_id, app_id, resource_type, resource_id
  FROM access.grants
  WHERE subject_type = 'USER'
    AND subject_id = 'user-uuid'
    AND revoked_at IS NULL
    AND (expires_at IS NULL OR expires_at > NOW())
),
-- 2. 如果是角色授权，查找角色包含的权限
role_perms AS (
  SELECT DISTINCT rp.permission_id, ug.tenant_id, ug.app_id, ug.resource_type, ug.resource_id
  FROM user_grants ug
  JOIN access.role_permissions rp ON ug.grant_type = 'ROLE' AND ug.grant_ref_id = rp.role_id
),
-- 3. 如果是权限授权，直接使用
perm_grants AS (
  SELECT grant_ref_id AS permission_id, tenant_id, app_id, resource_type, resource_id
  FROM user_grants
  WHERE grant_type = 'PERMISSION'
)
-- 4. 合并结果
SELECT p.key, p.name, COALESCE(rp.tenant_id, pg.tenant_id) AS tenant_id
FROM access.permissions p
LEFT JOIN role_perms rp ON p.id = rp.permission_id
LEFT JOIN perm_grants pg ON p.id = pg.permission_id
WHERE (rp.permission_id IS NOT NULL OR pg.permission_id IS NOT NULL)
  AND p.deleted_at IS NULL;
```

### 验证用户是否有权限
```sql
-- 检查用户是否有特定权限（在特定租户内）
SELECT EXISTS(
  SELECT 1
  FROM access.grants g
  WHERE g.subject_type = 'USER'
    AND g.subject_id = 'user-uuid'
    AND g.tenant_id = 'tenant-uuid'
    AND (
      (g.grant_type = 'PERMISSION' AND g.grant_ref_id = (SELECT id FROM access.permissions WHERE key = 'users.read'))
      OR
      (g.grant_type = 'ROLE' AND g.grant_ref_id IN (
        SELECT role_id FROM access.role_permissions
        WHERE permission_id = (SELECT id FROM access.permissions WHERE key = 'users.read')
      ))
    )
    AND g.revoked_at IS NULL
    AND (g.expires_at IS NULL OR g.expires_at > NOW())
);
```

## 最佳实践 / Best Practices

1. **权限粒度**: 权限应设计得足够细粒度，以便灵活组合
2. **角色设计**: 角色应基于实际业务场景设计，避免过度复杂
3. **Scope 映射**: OAuth scope 与内部权限分离，通过映射表关联
4. **策略优先级**: 合理设置策略优先级，避免冲突
5. **授权撤销**: 使用软删除机制，保留历史记录用于审计
6. **性能优化**: 授权查询可能涉及多表关联，建议使用索引和缓存优化

## 相关文档 / Related Documentation

- [permissions.md](./permissions.md) - 权限表详细文档
- [roles.md](./roles.md) - 角色表详细文档
- [role_permissions.md](./role_permissions.md) - 角色权限关联表详细文档
- [scopes.md](./scopes.md) - OAuth Scope 表详细文档
- [scope_permissions.md](./scope_permissions.md) - Scope 权限映射表详细文档
- [policies.md](./policies.md) - 策略表详细文档
- [grants.md](./grants.md) - 授权表详细文档

