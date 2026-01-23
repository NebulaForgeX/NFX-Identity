/*
 * prettier-ignore
 * This file contains constants for the backend API endpoints and WebSocket URLs.
 * API endpoints are based on NFX-ID routes.
 */

// 从环境变量获取配置
// 通过 Traefik 反向代理访问：10166 是 Traefik 的 HTTP 端口
const HTTP_BASE_URL = import.meta.env.VITE_API_URL || "http://localhost:10166";
const WS_BASE_URL = import.meta.env.VITE_WS_URL || "ws://localhost:10166";

// 图片服务 URL（独立配置）
// 通过 Traefik 反向代理访问
// API路由定义 - 通过 Traefik 反向代理，路径前缀为 /access, /audit, /auth, /clients, /directory, /image, /system, /tenants
export const URL_PATHS = {
  // Access 模块 - /access/auth
  ACCESS: {
    // 角色相关 - 需要认证
    CREATE_ROLE: "/access/auth/roles",
    GET_ROLE: "/access/auth/roles/:id",
    GET_ROLE_BY_KEY: "/access/auth/roles/key/:key",
    UPDATE_ROLE: "/access/auth/roles/:id",
    DELETE_ROLE: "/access/auth/roles/:id",
    // 权限相关 - 需要认证
    CREATE_PERMISSION: "/access/auth/permissions",
    GET_PERMISSION: "/access/auth/permissions/:id",
    GET_PERMISSION_BY_KEY: "/access/auth/permissions/key/:key",
    UPDATE_PERMISSION: "/access/auth/permissions/:id",
    DELETE_PERMISSION: "/access/auth/permissions/:id",
    // 作用域相关 - 需要认证
    CREATE_SCOPE: "/access/auth/scopes",
    GET_SCOPE: "/access/auth/scopes/:scope",
    UPDATE_SCOPE: "/access/auth/scopes/:scope",
    DELETE_SCOPE: "/access/auth/scopes/:scope",
    // 授权相关 - 需要认证
    CREATE_GRANT: "/access/auth/grants",
    GET_GRANT: "/access/auth/grants/:id",
    UPDATE_GRANT: "/access/auth/grants/:id",
    DELETE_GRANT: "/access/auth/grants/:id",
    // 角色权限关联相关 - 需要认证
    CREATE_ROLE_PERMISSION: "/access/auth/role-permissions",
    GET_ROLE_PERMISSION: "/access/auth/role-permissions/:id",
    DELETE_ROLE_PERMISSION: "/access/auth/role-permissions/:id",
    // 作用域权限关联相关 - 需要认证
    CREATE_SCOPE_PERMISSION: "/access/auth/scope-permissions",
    GET_SCOPE_PERMISSION: "/access/auth/scope-permissions/:id",
    DELETE_SCOPE_PERMISSION: "/access/auth/scope-permissions/:id",
  },
  // Audit 模块 - /audit/auth
  AUDIT: {
    // 事件相关 - 需要认证
    CREATE_EVENT: "/audit/auth/events",
    GET_EVENT: "/audit/auth/events/:id",
    DELETE_EVENT: "/audit/auth/events/:id",
    // Actor Snapshot 相关 - 需要认证
    CREATE_ACTOR_SNAPSHOT: "/audit/auth/actor-snapshots",
    GET_ACTOR_SNAPSHOT: "/audit/auth/actor-snapshots/:id",
    DELETE_ACTOR_SNAPSHOT: "/audit/auth/actor-snapshots/:id",
    // Event Retention Policy 相关 - 需要认证
    CREATE_EVENT_RETENTION_POLICY: "/audit/auth/event-retention-policies",
    GET_EVENT_RETENTION_POLICY: "/audit/auth/event-retention-policies/:id",
    UPDATE_EVENT_RETENTION_POLICY: "/audit/auth/event-retention-policies/:id",
    DELETE_EVENT_RETENTION_POLICY: "/audit/auth/event-retention-policies/:id",
    // Event Search Index 相关 - 需要认证
    CREATE_EVENT_SEARCH_INDEX: "/audit/auth/event-search-index",
    GET_EVENT_SEARCH_INDEX: "/audit/auth/event-search-index/:id",
    DELETE_EVENT_SEARCH_INDEX: "/audit/auth/event-search-index/:id",
    // Hash Chain Checkpoint 相关 - 需要认证
    CREATE_HASH_CHAIN_CHECKPOINT: "/audit/auth/hash-chain-checkpoints",
    GET_HASH_CHAIN_CHECKPOINT: "/audit/auth/hash-chain-checkpoints/:id",
    DELETE_HASH_CHAIN_CHECKPOINT: "/audit/auth/hash-chain-checkpoints/:id",
  },
  // Auth 模块 - /auth/auth
  AUTH: {
    // 会话相关 - 需要认证
    CREATE_SESSION: "/auth/auth/sessions",
    GET_SESSION: "/auth/auth/sessions/:id",
    REVOKE_SESSION: "/auth/auth/sessions/:session_id/revoke",
    DELETE_SESSION: "/auth/auth/sessions/:id",
    // 用户凭证相关 - 需要认证
    CREATE_USER_CREDENTIAL: "/auth/auth/user-credentials",
    GET_USER_CREDENTIAL: "/auth/auth/user-credentials/:id",
    UPDATE_USER_CREDENTIAL: "/auth/auth/user-credentials/:id",
    DELETE_USER_CREDENTIAL: "/auth/auth/user-credentials/:id",
    // MFA 因子相关 - 需要认证
    CREATE_MFA_FACTOR: "/auth/auth/mfa-factors",
    GET_MFA_FACTOR: "/auth/auth/mfa-factors/:id",
    UPDATE_MFA_FACTOR: "/auth/auth/mfa-factors/:id",
    DELETE_MFA_FACTOR: "/auth/auth/mfa-factors/:id",
    // 刷新令牌相关 - 需要认证
    CREATE_REFRESH_TOKEN: "/auth/auth/refresh-tokens",
    GET_REFRESH_TOKEN: "/auth/auth/refresh-tokens/:id",
    UPDATE_REFRESH_TOKEN: "/auth/auth/refresh-tokens/:id",
    DELETE_REFRESH_TOKEN: "/auth/auth/refresh-tokens/:id",
    // 密码重置相关 - 需要认证
    CREATE_PASSWORD_RESET: "/auth/auth/password-resets",
    GET_PASSWORD_RESET: "/auth/auth/password-resets/:id",
    UPDATE_PASSWORD_RESET: "/auth/auth/password-resets/:id",
    DELETE_PASSWORD_RESET: "/auth/auth/password-resets/:id",
    // 密码历史相关 - 需要认证
    CREATE_PASSWORD_HISTORY: "/auth/auth/password-history",
    GET_PASSWORD_HISTORY: "/auth/auth/password-history/:id",
    // 登录尝试相关 - 需要认证
    CREATE_LOGIN_ATTEMPT: "/auth/auth/login-attempts",
    GET_LOGIN_ATTEMPT: "/auth/auth/login-attempts/:id",
    DELETE_LOGIN_ATTEMPT: "/auth/auth/login-attempts/:id",
    // 账户锁定相关 - 需要认证
    CREATE_ACCOUNT_LOCKOUT: "/auth/auth/account-lockouts",
    GET_ACCOUNT_LOCKOUT: "/auth/auth/account-lockouts/:id",
    UPDATE_ACCOUNT_LOCKOUT: "/auth/auth/account-lockouts/:id",
    DELETE_ACCOUNT_LOCKOUT: "/auth/auth/account-lockouts/:id",
    // 受信任设备相关 - 需要认证
    CREATE_TRUSTED_DEVICE: "/auth/auth/trusted-devices",
    GET_TRUSTED_DEVICE: "/auth/auth/trusted-devices/:id",
    DELETE_TRUSTED_DEVICE: "/auth/auth/trusted-devices/:id",
  },
  // Clients 模块 - /clients/auth
  CLIENTS: {
    // 应用相关 - 需要认证
    CREATE_APP: "/clients/auth/apps",
    GET_APP: "/clients/auth/apps/:id",
    GET_APP_BY_APP_ID: "/clients/auth/apps/app-id/:app_id",
    UPDATE_APP: "/clients/auth/apps/:id",
    DELETE_APP: "/clients/auth/apps/:id",
    // API Key 相关 - 需要认证
    CREATE_API_KEY: "/clients/auth/api-keys",
    GET_API_KEY: "/clients/auth/api-keys/:id",
    DELETE_API_KEY_BY_KEY_ID: "/clients/auth/api-keys/key-id/:key_id",
    // Client Credential 相关 - 需要认证
    CREATE_CLIENT_CREDENTIAL: "/clients/auth/client-credentials",
    GET_CLIENT_CREDENTIAL: "/clients/auth/client-credentials/:id",
    DELETE_CLIENT_CREDENTIAL_BY_CLIENT_ID: "/clients/auth/client-credentials/client-id/:client_id",
    // Client Scope 相关 - 需要认证
    CREATE_CLIENT_SCOPE: "/clients/auth/client-scopes",
    GET_CLIENT_SCOPE: "/clients/auth/client-scopes/:id",
    DELETE_CLIENT_SCOPE: "/clients/auth/client-scopes/:id",
    // IP Allowlist 相关 - 需要认证
    CREATE_IP_ALLOWLIST: "/clients/auth/ip-allowlist",
    GET_IP_ALLOWLIST: "/clients/auth/ip-allowlist/:id",
    DELETE_IP_ALLOWLIST_BY_RULE_ID: "/clients/auth/ip-allowlist/rule-id/:rule_id",
    // Rate Limit 相关 - 需要认证
    CREATE_RATE_LIMIT: "/clients/auth/rate-limits",
    GET_RATE_LIMIT: "/clients/auth/rate-limits/:id",
    DELETE_RATE_LIMIT: "/clients/auth/rate-limits/:id",
  },
  // Directory 模块 - /directory/auth
  DIRECTORY: {
    // 用户相关 - 需要认证
    CREATE_USER: "/directory/auth/users",
    GET_USER: "/directory/auth/users/:id",
    GET_USER_BY_USERNAME: "/directory/auth/users/username/:username",
    UPDATE_USER_STATUS: "/directory/auth/users/:id/status",
    UPDATE_USER_USERNAME: "/directory/auth/users/:id/username",
    VERIFY_USER: "/directory/auth/users/:id/verify",
    DELETE_USER: "/directory/auth/users/:id",
    // 徽章相关 - 需要认证
    CREATE_BADGE: "/directory/auth/badges",
    GET_BADGE: "/directory/auth/badges/:id",
    GET_BADGE_BY_NAME: "/directory/auth/badges/name/:name",
    UPDATE_BADGE: "/directory/auth/badges/:id",
    DELETE_BADGE: "/directory/auth/badges/:id",
    // 用户徽章相关 - 需要认证
    CREATE_USER_BADGE: "/directory/auth/user-badges",
    GET_USER_BADGE: "/directory/auth/user-badges/:id",
    DELETE_USER_BADGE: "/directory/auth/user-badges/:id",
    // 用户教育相关 - 需要认证
    CREATE_USER_EDUCATION: "/directory/auth/user-educations",
    GET_USER_EDUCATION: "/directory/auth/user-educations/:id",
    UPDATE_USER_EDUCATION: "/directory/auth/user-educations/:id",
    DELETE_USER_EDUCATION: "/directory/auth/user-educations/:id",
    // 用户邮箱相关 - 需要认证
    CREATE_USER_EMAIL: "/directory/auth/user-emails",
    GET_USER_EMAIL: "/directory/auth/user-emails/:id",
    UPDATE_USER_EMAIL: "/directory/auth/user-emails/:id",
    SET_PRIMARY_USER_EMAIL: "/directory/auth/user-emails/:id/set-primary",
    VERIFY_USER_EMAIL: "/directory/auth/user-emails/:id/verify",
    DELETE_USER_EMAIL: "/directory/auth/user-emails/:id",
    // 用户职业相关 - 需要认证
    CREATE_USER_OCCUPATION: "/directory/auth/user-occupations",
    GET_USER_OCCUPATION: "/directory/auth/user-occupations/:id",
    UPDATE_USER_OCCUPATION: "/directory/auth/user-occupations/:id",
    DELETE_USER_OCCUPATION: "/directory/auth/user-occupations/:id",
    // 用户电话相关 - 需要认证
    CREATE_USER_PHONE: "/directory/auth/user-phones",
    GET_USER_PHONE: "/directory/auth/user-phones/:id",
    UPDATE_USER_PHONE: "/directory/auth/user-phones/:id",
    SET_PRIMARY_USER_PHONE: "/directory/auth/user-phones/:id/set-primary",
    VERIFY_USER_PHONE: "/directory/auth/user-phones/:id/verify",
    DELETE_USER_PHONE: "/directory/auth/user-phones/:id",
    // 用户偏好相关 - 需要认证
    CREATE_USER_PREFERENCE: "/directory/auth/user-preferences",
    GET_USER_PREFERENCE: "/directory/auth/user-preferences/:id",
    UPDATE_USER_PREFERENCE: "/directory/auth/user-preferences/:id",
    DELETE_USER_PREFERENCE: "/directory/auth/user-preferences/:id",
    // 用户资料相关 - 需要认证
    CREATE_USER_PROFILE: "/directory/auth/user-profiles",
    GET_USER_PROFILE: "/directory/auth/user-profiles/:id",
    UPDATE_USER_PROFILE: "/directory/auth/user-profiles/:id",
    DELETE_USER_PROFILE: "/directory/auth/user-profiles/:id",
  },
  // Image 模块 - /image/auth
  IMAGE: {
    // 图片相关 - 需要认证
    CREATE_IMAGE: "/image/auth/images",
    GET_IMAGE: "/image/auth/images/:id",
    UPDATE_IMAGE: "/image/auth/images/:id",
    DELETE_IMAGE: "/image/auth/images/:id",
    // 图片类型相关 - 需要认证
    CREATE_IMAGE_TYPE: "/image/auth/image-types",
    GET_IMAGE_TYPE: "/image/auth/image-types/:id",
    UPDATE_IMAGE_TYPE: "/image/auth/image-types/:id",
    DELETE_IMAGE_TYPE: "/image/auth/image-types/:id",
    // 图片变体相关 - 需要认证
    CREATE_IMAGE_VARIANT: "/image/auth/image-variants",
    GET_IMAGE_VARIANT: "/image/auth/image-variants/:id",
    UPDATE_IMAGE_VARIANT: "/image/auth/image-variants/:id",
    DELETE_IMAGE_VARIANT: "/image/auth/image-variants/:id",
    // 图片标签相关 - 需要认证
    CREATE_IMAGE_TAG: "/image/auth/image-tags",
    GET_IMAGE_TAG: "/image/auth/image-tags/:id",
    UPDATE_IMAGE_TAG: "/image/auth/image-tags/:id",
    DELETE_IMAGE_TAG: "/image/auth/image-tags/:id",
  },
  // System 模块 - /system
  SYSTEM: {
    // 系统状态相关 - 公开接口（不需要认证）
    GET_SYSTEM_STATE_LATEST: "/system/system-state/latest",
    INITIALIZE_SYSTEM_STATE: "/system/system-state/initialize",
    // 系统状态相关 - 需要认证
    GET_SYSTEM_STATE: "/system/auth/system-state/:id",
    RE_INITIALIZE_SYSTEM_STATE: "/system/auth/system-state/initialize",
    RESET_SYSTEM_STATE: "/system/auth/system-state/reset",
    DELETE_SYSTEM_STATE: "/system/auth/system-state/:id",
  },
  // Tenants 模块 - /tenants/auth
  TENANTS: {
    // 租户相关 - 需要认证
    CREATE_TENANT: "/tenants/auth/",
    GET_TENANT: "/tenants/auth/:id",
    GET_TENANT_BY_TENANT_ID: "/tenants/auth/tenant-id/:tenant_id",
    UPDATE_TENANT: "/tenants/auth/:id",
    UPDATE_TENANT_STATUS: "/tenants/auth/:id/status",
    DELETE_TENANT: "/tenants/auth/:id",
    // 组相关 - 需要认证
    CREATE_GROUP: "/tenants/auth/groups",
    GET_GROUP: "/tenants/auth/groups/:id",
    UPDATE_GROUP: "/tenants/auth/groups/:id",
    DELETE_GROUP: "/tenants/auth/groups/:id",
    // 成员相关 - 需要认证
    CREATE_MEMBER: "/tenants/auth/members",
    GET_MEMBER: "/tenants/auth/members/:id",
    UPDATE_MEMBER: "/tenants/auth/members/:id",
    DELETE_MEMBER: "/tenants/auth/members/:id",
    // 邀请相关 - 需要认证
    CREATE_INVITATION: "/tenants/auth/invitations",
    GET_INVITATION: "/tenants/auth/invitations/:id",
    GET_INVITATION_BY_INVITE_ID: "/tenants/auth/invitations/invite-id/:invite_id",
    ACCEPT_INVITATION: "/tenants/auth/invitations/invite-id/:invite_id/accept",
    REVOKE_INVITATION: "/tenants/auth/invitations/invite-id/:invite_id/revoke",
    DELETE_INVITATION: "/tenants/auth/invitations/:id",
    // 租户应用相关 - 需要认证
    CREATE_TENANT_APP: "/tenants/auth/tenant-apps",
    GET_TENANT_APP: "/tenants/auth/tenant-apps/:id",
    UPDATE_TENANT_APP: "/tenants/auth/tenant-apps/:id",
    DELETE_TENANT_APP: "/tenants/auth/tenant-apps/:id",
    // 租户设置相关 - 需要认证
    CREATE_TENANT_SETTING: "/tenants/auth/tenant-settings",
    GET_TENANT_SETTING: "/tenants/auth/tenant-settings/:id",
    UPDATE_TENANT_SETTING: "/tenants/auth/tenant-settings/:id",
    DELETE_TENANT_SETTING: "/tenants/auth/tenant-settings/:id",
    // 域名验证相关 - 需要认证
    CREATE_DOMAIN_VERIFICATION: "/tenants/auth/domain-verifications",
    GET_DOMAIN_VERIFICATION: "/tenants/auth/domain-verifications/:id",
    UPDATE_DOMAIN_VERIFICATION: "/tenants/auth/domain-verifications/:id",
    DELETE_DOMAIN_VERIFICATION: "/tenants/auth/domain-verifications/:id",
    // 成员角色相关 - 需要认证
    CREATE_MEMBER_ROLE: "/tenants/auth/member-roles",
    GET_MEMBER_ROLE: "/tenants/auth/member-roles/:id",
    REVOKE_MEMBER_ROLE: "/tenants/auth/member-roles/:id/revoke",
    DELETE_MEMBER_ROLE: "/tenants/auth/member-roles/:id",
    // 成员组相关 - 需要认证
    CREATE_MEMBER_GROUP: "/tenants/auth/member-groups",
    GET_MEMBER_GROUP: "/tenants/auth/member-groups/:id",
    REVOKE_MEMBER_GROUP: "/tenants/auth/member-groups/:id/revoke",
    DELETE_MEMBER_GROUP: "/tenants/auth/member-groups/:id",
    // 成员应用角色相关 - 需要认证
    CREATE_MEMBER_APP_ROLE: "/tenants/auth/member-app-roles",
    GET_MEMBER_APP_ROLE: "/tenants/auth/member-app-roles/:id",
    REVOKE_MEMBER_APP_ROLE: "/tenants/auth/member-app-roles/:id/revoke",
    DELETE_MEMBER_APP_ROLE: "/tenants/auth/member-app-roles/:id",
  },
} as const;

export const API_ENDPOINTS = {
  PURE: HTTP_BASE_URL,
  WS: WS_BASE_URL,
} as const;

// 类型定义
export type URL_PATHS_TYPE = typeof URL_PATHS;
export type API_ENDPOINTS_TYPE = typeof API_ENDPOINTS;
