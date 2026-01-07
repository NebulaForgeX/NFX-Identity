# Tenants Schema 总览 / Tenants Schema Overview

## 概述 / Overview

`tenants` schema 是 NFX-Identity 平台的多租户管理模块，负责管理租户（公司/组织）、成员、邀请、组织架构等。

The `tenants` schema is the multi-tenant management module of the NFX-Identity platform, responsible for managing tenants (companies/organizations), members, invitations, organizational structure, etc.

## 核心概念 / Core Concepts

### 1. 租户（Tenants）
- **定义**: 租户主表，表示公司/组织
- **表**: `tenants.tenants`
- **特点**: 多租户隔离的核心，所有数据都关联到租户

### 2. 成员（Members）
- **定义**: 租户成员关系表
- **表**: `tenants.members`
- **特点**: 表示"谁属于哪个租户"

### 3. 邀请（Invitations）
- **定义**: 邀请机制
- **表**: `tenants.invitations`
- **特点**: 支持安全的成员邀请流程，包含预分配角色（role_ids 字段）

### 4. 成员角色（Member Roles）
- **定义**: 租户内的角色分配
- **表**: `tenants.member_roles`
- **特点**: 支持临时角色、过期角色

### 5. 组（Groups）
- **定义**: 组织架构管理
- **表**: `tenants.groups`、`tenants.member_groups`
- **特点**: 支持层级结构（部门、团队）

### 6. 租户应用（Tenant Apps）
- **定义**: 租户关联的应用
- **表**: `tenants.tenant_apps`
- **特点**: 管理租户可以使用的应用

### 7. 租户设置（Tenant Settings）
- **定义**: 租户级别的设置
- **表**: `tenants.tenant_settings`
- **特点**: 存储租户的配置和偏好

### 8. 域名验证（Domain Verifications）
- **定义**: 租户域名验证
- **表**: `tenants.domain_verifications`
- **特点**: 用于 SSO 和邮箱域名限制

## 表关系图 / Table Relationships

```
┌─────────────────┐
│    tenants      │
│  (租户主表)     │
└────────┬────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────┐
    │     members           │
    │  (租户成员)           │
    └────┬──────────────────┘
         │
         │ 1:N
         │
    ┌────▼──────────────────┐
    │   member_roles        │
    │  (成员角色)            │
    └───────────────────────┘

┌─────────────────┐
│  invitations    │
│  (邀请+角色)    │
└─────────────────┘

┌─────────────────┐
│     groups      │
│  (组织架构)     │
└────────┬────────┘
         │
         │ 层级结构
         │
    ┌────▼──────────────────┐
    │  member_groups        │
    │  (成员-组关系)        │
    └───────────────────────┘

┌─────────────────┐
│   tenant_apps   │
│  (租户应用)     │
└─────────────────┘

┌─────────────────┐
│ tenant_settings │
│  (租户设置)     │
└─────────────────┘

┌─────────────────┐
│member_app_roles │
│(成员应用角色)   │
└─────────────────┘

┌─────────────────┐
│domain_verifications│
│  (域名验证)     │
└─────────────────┘
```

## 表列表 / Table List

### 1. `tenants.tenants` - 租户表
- **用途**: 租户主表
- **关键字段**: `tenant_id`、`name`、`status`、`primary_domain`
- **详细文档**: [tenants.md](./tenants.md)

### 2. `tenants.members` - 成员表
- **用途**: 租户成员关系表
- **关键字段**: `member_id`、`tenant_id`、`user_id`、`status`、`source`
- **详细文档**: [members.md](./members.md)

### 3. `tenants.invitations` - 邀请表
- **用途**: 邀请机制（包含预分配角色）
- **关键字段**: `invite_id`、`tenant_id`、`email`、`token_hash`、`status`、`role_ids`
- **详细文档**: [invitations.md](./invitations.md)

### 4. `tenants.member_roles` - 成员角色表
- **用途**: 租户内的角色分配
- **关键字段**: `member_id`、`role_id`、`expires_at`
- **详细文档**: [member_roles.md](./member_roles.md)

### 6. `tenants.groups` - 组表
- **用途**: 组织架构管理
- **关键字段**: `tenant_id`、`name`、`parent_group_id`、`group_type`
- **详细文档**: [groups.md](./groups.md)

### 7. `tenants.member_groups` - 成员组表
- **用途**: 成员和组的多对多关系
- **关键字段**: `member_id`、`group_id`
- **详细文档**: [member_groups.md](./member_groups.md)

### 8. `tenants.tenant_apps` - 租户应用表
- **用途**: 租户关联的应用
- **关键字段**: `tenant_id`、`app_id`、`status`
- **详细文档**: [tenant_apps.md](./tenant_apps.md)

### 9. `tenants.tenant_settings` - 租户设置表
- **用途**: 租户级别的设置
- **关键字段**: `tenant_id`、`enforce_mfa`、`password_policy`、`login_policy`
- **详细文档**: [tenant_settings.md](./tenant_settings.md)

### 10. `tenants.member_app_roles` - 成员应用角色表
- **用途**: 成员在特定应用中的角色
- **关键字段**: `member_id`、`app_id`、`role_id`
- **详细文档**: [member_app_roles.md](./member_app_roles.md)

### 11. `tenants.domain_verifications` - 域名验证表
- **用途**: 租户域名验证
- **关键字段**: `tenant_id`、`domain`、`verification_status`
- **详细文档**: [domain_verifications.md](./domain_verifications.md)

## 成员生命周期 / Member Lifecycle

```
1. 管理员发送邀请（invitations，包含预分配角色 role_ids）
2. 用户接受邀请
3. 创建成员（members）
4. 从 invitations.role_ids 分配角色（member_roles）
5. 添加到组（member_groups）
6. 分配应用角色（member_app_roles）
7. 成员激活（members.status = 'ACTIVE'）
8. 成员移除（members.status = 'REMOVED'）
```

## 相关文档 / Related Documentation

- [tenants.md](./tenants.md) - 租户表详细文档
- [members.md](./members.md) - 成员表详细文档
- [invitations.md](./invitations.md) - 邀请表详细文档（包含预分配角色）
- [member_roles.md](./member_roles.md) - 成员角色表详细文档
- [groups.md](./groups.md) - 组表详细文档
- [member_groups.md](./member_groups.md) - 成员组表详细文档
- [tenant_apps.md](./tenant_apps.md) - 租户应用表详细文档
- [tenant_settings.md](./tenant_settings.md) - 租户设置表详细文档
- [member_app_roles.md](./member_app_roles.md) - 成员应用角色表详细文档
- [domain_verifications.md](./domain_verifications.md) - 域名验证表详细文档
