# NFX-Identity 数据库表清单 / Database Tables Overview

## Directory Schema - 用户目录模块

| 表名 / Table | 作用 / Purpose | 文档 / Documentation |
|-------------|--------------|---------------------|
| `users` | 用户主表，存储用户名和账户状态（认证凭据在 auth.user_credentials） | [users.md](./directory/users.md) |
| `user_profiles` | 用户详细资料信息 | [user_profiles.md](./directory/user_profiles.md) |
| `user_emails` | 用户的多个邮箱地址 | [user_emails.md](./directory/user_emails.md) |
| `user_phones` | 用户的多个手机号码 | [user_phones.md](./directory/user_phones.md) |
| `user_preferences` | 用户设置和偏好 | [user_preferences.md](./directory/user_preferences.md) |
| `badges` | 徽章定义 | [badges.md](./directory/badges.md) |
| `user_badges` | 用户和徽章的多对多关系 | [user_badges.md](./directory/user_badges.md) |
| `user_educations` | 用户教育经历 | [user_educations.md](./directory/user_educations.md) |
| `user_occupations` | 用户工作经历 | [user_occupations.md](./directory/user_occupations.md) |

**总览文档**: [directory.md](./directory/directory.md)

## Auth Schema - 身份认证模块

| 表名 / Table | 作用 / Purpose | 文档 / Documentation |
|-------------|--------------|---------------------|
| `user_credentials` | 用户凭证（密码、OAuth等） | [user_credentials.md](./auth/user_credentials.md) |
| `sessions` | 用户会话管理 | [sessions.md](./auth/sessions.md) |
| `refresh_tokens` | OAuth刷新令牌 | [refresh_tokens.md](./auth/refresh_tokens.md) |
| `login_attempts` | 登录尝试记录 | [login_attempts.md](./auth/login_attempts.md) |
| `account_lockouts` | 账户锁定记录 | [account_lockouts.md](./auth/account_lockouts.md) |
| `password_resets` | 密码重置请求 | [password_resets.md](./auth/password_resets.md) |
| `mfa_factors` | 多因素认证因子 | [mfa_factors.md](./auth/mfa_factors.md) |
| `password_history` | 密码历史记录 | [password_history.md](./auth/password_history.md) |
| `trusted_devices` | 受信任设备 | [trusted_devices.md](./auth/trusted_devices.md) |

**总览文档**: [auth.md](./auth/auth.md)

## Clients Schema - 客户端/应用管理模块

| 表名 / Table | 作用 / Purpose | 文档 / Documentation |
|-------------|--------------|---------------------|
| `apps` | 应用主表 | [apps.md](./clients/apps.md) |
| `client_credentials` | OAuth客户端凭证 | [client_credentials.md](./clients/client_credentials.md) |
| `api_keys` | API密钥 | [api_keys.md](./clients/api_keys.md) |
| `client_scopes` | 客户端允许的权限范围 | [client_scopes.md](./clients/client_scopes.md) |
| `ip_allowlist` | IP白名单 | [ip_allowlist.md](./clients/ip_allowlist.md) |
| `rate_limits` | 客户端级别限流配置 | [rate_limits.md](./clients/rate_limits.md) |

**总览文档**: [clients.md](./clients/clients.md)

## Access Schema - 访问控制/授权模块

| 表名 / Table | 作用 / Purpose | 文档 / Documentation |
|-------------|--------------|---------------------|
| `permissions` | 权限定义 | [permissions.md](./access/permissions.md) |
| `roles` | 角色定义 | [roles.md](./access/roles.md) |
| `role_permissions` | 角色和权限的多对多关系 | [role_permissions.md](./access/role_permissions.md) |
| `scopes` | 权限范围定义 | [scopes.md](./access/scopes.md) |
| `scope_permissions` | 权限范围和权限的多对多关系 | [scope_permissions.md](./access/scope_permissions.md) |
| `grants` | 授权记录（主体-资源-权限） | [grants.md](./access/grants.md) |

**总览文档**: [access.md](./access/access.md)

## Tenants Schema - 多租户管理模块

| 表名 / Table | 作用 / Purpose | 文档 / Documentation |
|-------------|--------------|---------------------|
| `tenants` | 租户主表 | [tenants.md](./tenants/tenants.md) |
| `members` | 租户成员关系表 | [members.md](./tenants/members.md) |
| `member_roles` | 租户内的角色分配 | [member_roles.md](./tenants/member_roles.md) |
| `invitations` | 邀请机制（包含预分配角色） | [invitations.md](./tenants/invitations.md) |
| `groups` | 组织架构管理 | [groups.md](./tenants/groups.md) |
| `member_groups` | 成员和组的多对多关系 | [member_groups.md](./tenants/member_groups.md) |
| `tenant_apps` | 租户关联的应用 | [tenant_apps.md](./tenants/tenant_apps.md) |
| `tenant_settings` | 租户级别的设置 | [tenant_settings.md](./tenants/tenant_settings.md) |
| `member_app_roles` | 成员在特定应用中的角色 | [member_app_roles.md](./tenants/member_app_roles.md) |
| `domain_verifications` | 租户域名验证 | [domain_verifications.md](./tenants/domain_verifications.md) |

**总览文档**: [tenants.md](./tenants/tenants.md)

## Image Schema - 图片管理模块

| 表名 / Table | 作用 / Purpose | 文档 / Documentation |
|-------------|--------------|---------------------|
| `image_types` | 图片类型定义 | [image_types.md](./image/image_types.md) |
| `images` | 图片主表 | [images.md](./image/images.md) |
| `image_variants` | 图片变体（缩略图、不同尺寸等） | [image_variants.md](./image/image_variants.md) |
| `image_tags` | 图片标签 | [image_tags.md](./image/image_tags.md) |

**总览文档**: [image.md](./image/image.md)

## Audit Schema - 审计模块

| 表名 / Table | 作用 / Purpose | 文档 / Documentation |
|-------------|--------------|---------------------|
| `events` | 审计事件事实表 | [events.md](./audit/events.md) |
| `event_search_index` | 事件搜索索引表 | [event_search_index.md](./audit/event_search_index.md) |
| `actor_snapshots` | 操作者快照 | [actor_snapshots.md](./audit/actor_snapshots.md) |
| `hash_chain_checkpoints` | 哈希链检查点 | [hash_chain_checkpoints.md](./audit/hash_chain_checkpoints.md) |
| `event_retention_policies` | 事件保留策略 | [event_retention_policies.md](./audit/event_retention_policies.md) |

**总览文档**: [audit.md](./audit/audit.md)

---

## 统计 / Statistics

- **Schema 数量**: 7
- **表总数**: 52
- **Directory Schema**: 9 张表
- **Auth Schema**: 9 张表
- **Clients Schema**: 6 张表
- **Access Schema**: 6 张表
- **Tenants Schema**: 10 张表
- **Image Schema**: 4 张表
- **Audit Schema**: 5 张表

## 已移除的表 / Removed Tables

以下表已从默认架构中移除（建议使用替代方案）：

- **auth.credential_events** - 建议使用 `audit.events` 统一审计
- **clients.service_tokens** - 仅 Opaque token 需要；JWT 模式默认不启用
- **clients.token_usage_logs** - 建议写入 `audit.events`，避免重复且爆量
- **access.policies** - ABAC 复杂度高，先做 RBAC+grants 足够
- **tenants.invitation_roles** - 已并入 `invitations.role_ids` 字段
