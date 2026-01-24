// Domain Enums - 基于后端数据库定义
// 所有枚举类型都从后端数据库 schema 中提取

// ========== Directory Schema ==========

/**
 * 用户状态
 * directory.user_status
 */
export enum UserStatus {
  PENDING = "pending",
  ACTIVE = "active",
  DEACTIVE = "deactive",
}

// ========== Auth Schema ==========

/**
 * 凭证类型
 * auth.credential_type
 */
export enum CredentialType {
  PASSWORD = "password",
  PASSKEY = "passkey",
  OAUTH_LINK = "oauth_link",
  SAML = "saml",
  LDAP = "ldap",
}

/**
 * 凭证状态
 * auth.credential_status
 */
export enum CredentialStatus {
  ACTIVE = "active",
  DISABLED = "disabled",
  EXPIRED = "expired",
}

/**
 * MFA 类型
 * auth.mfa_type
 */
export enum MFAType {
  TOTP = "totp",
  SMS = "sms",
  EMAIL = "email",
  WEBAUTHN = "webauthn",
  BACKUP_CODE = "backup_code",
}

/**
 * 撤销原因
 * auth.revoke_reason
 */
export enum RevokeReason {
  USER_LOGOUT = "user_logout",
  ADMIN_REVOKE = "admin_revoke",
  PASSWORD_CHANGED = "password_changed",
  ROTATION = "rotation",
  ACCOUNT_LOCKED = "account_locked",
  DEVICE_CHANGED = "device_changed",
  SUSPICIOUS_ACTIVITY = "suspicious_activity",
  OTHER = "other",
}

/**
 * 会话撤销原因
 * auth.session_revoke_reason
 */
export enum SessionRevokeReason {
  USER_LOGOUT = "user_logout",
  ADMIN_REVOKE = "admin_revoke",
  PASSWORD_CHANGED = "password_changed",
  DEVICE_CHANGED = "device_changed",
  ACCOUNT_LOCKED = "account_locked",
  SUSPICIOUS_ACTIVITY = "suspicious_activity",
  SESSION_EXPIRED = "session_expired",
  OTHER = "other",
}

/**
 * 登录失败代码
 * auth.failure_code
 */
export enum FailureCode {
  BAD_PASSWORD = "bad_password",
  USER_NOT_FOUND = "user_not_found",
  LOCKED = "locked",
  MFA_REQUIRED = "mfa_required",
  MFA_FAILED = "mfa_failed",
  ACCOUNT_DISABLED = "account_disabled",
  CREDENTIAL_EXPIRED = "credential_expired",
  RATE_LIMITED = "rate_limited",
  IP_BLOCKED = "ip_blocked",
  DEVICE_NOT_TRUSTED = "device_not_trusted",
  OTHER = "other",
}

/**
 * 锁定原因
 * auth.lock_reason
 */
export enum LockReason {
  TOO_MANY_ATTEMPTS = "too_many_attempts",
  ADMIN_LOCK = "admin_lock",
  RISK_DETECTED = "risk_detected",
  SUSPICIOUS_ACTIVITY = "suspicious_activity",
  COMPLIANCE = "compliance",
  OTHER = "other",
}

/**
 * 密码重置交付方式
 * auth.reset_delivery
 */
export enum ResetDelivery {
  EMAIL = "email",
  SMS = "sms",
}

/**
 * 密码重置状态
 * auth.reset_status
 */
export enum ResetStatus {
  ISSUED = "issued",
  USED = "used",
  EXPIRED = "expired",
  REVOKED = "revoked",
}

// ========== Tenants Schema ==========

/**
 * 租户状态
 * tenants.tenant_status
 */
export enum TenantStatus {
  ACTIVE = "ACTIVE",
  SUSPENDED = "SUSPENDED",
  CLOSED = "CLOSED",
  PENDING = "PENDING",
}

/**
 * 成员状态
 * tenants.member_status
 */
export enum MemberStatus {
  INVITED = "INVITED",
  ACTIVE = "ACTIVE",
  SUSPENDED = "SUSPENDED",
  REMOVED = "REMOVED",
}

/**
 * 成员来源
 * tenants.member_source
 */
export enum MemberSource {
  MANUAL = "MANUAL",
  INVITE = "INVITE",
  SCIM = "SCIM",
  SSO = "SSO",
  HR_SYNC = "HR_SYNC",
  IMPORT = "IMPORT",
}

/**
 * 租户应用状态
 * tenants.tenant_app_status
 */
export enum TenantAppStatus {
  ACTIVE = "ACTIVE",
  DISABLED = "DISABLED",
  SUSPENDED = "SUSPENDED",
}

/**
 * 验证状态
 * tenants.verification_status
 */
export enum VerificationStatus {
  PENDING = "PENDING",
  VERIFIED = "VERIFIED",
  FAILED = "FAILED",
  EXPIRED = "EXPIRED",
}

/**
 * 验证方法
 * tenants.verification_method
 */
export enum VerificationMethod {
  DNS = "DNS",
  TXT = "TXT",
  HTML = "HTML",
  FILE = "FILE",
}

/**
 * 邀请状态
 * tenants.invitation_status
 */
export enum InvitationStatus {
  PENDING = "PENDING",
  ACCEPTED = "ACCEPTED",
  EXPIRED = "EXPIRED",
  REVOKED = "REVOKED",
}

/**
 * 组类型
 * tenants.group_type
 */
export enum GroupType {
  DEPARTMENT = "department",
  TEAM = "team",
  GROUP = "group",
  OTHER = "other",
}

// ========== Clients Schema ==========

/**
 * 应用类型
 * clients.app_type
 */
export enum AppType {
  SERVER = "server",
  SERVICE = "service",
  INTERNAL = "internal",
  PARTNER = "partner",
  THIRD_PARTY = "third_party",
}

/**
 * 应用状态
 * clients.app_status
 */
export enum AppStatus {
  ACTIVE = "active",
  DISABLED = "disabled",
  SUSPENDED = "suspended",
  PENDING = "pending",
}

/**
 * 环境
 * clients.environment
 */
export enum Environment {
  PRODUCTION = "production",
  STAGING = "staging",
  DEVELOPMENT = "development",
  TEST = "test",
}

/**
 * API 密钥状态
 * clients.api_key_status
 */
export enum ApiKeyStatus {
  ACTIVE = "active",
  REVOKED = "revoked",
  EXPIRED = "expired",
}

/**
 * 客户端凭证状态
 * clients.credential_status
 */
export enum ClientCredentialStatus {
  ACTIVE = "active",
  EXPIRED = "expired",
  REVOKED = "revoked",
  ROTATING = "rotating",
}

/**
 * 白名单状态
 * clients.allowlist_status
 */
export enum AllowlistStatus {
  ACTIVE = "active",
  DISABLED = "disabled",
  REVOKED = "revoked",
}

/**
 * 速率限制类型
 * clients.rate_limit_type
 */
export enum RateLimitType {
  REQUESTS_PER_SECOND = "requests_per_second",
  REQUESTS_PER_MINUTE = "requests_per_minute",
  REQUESTS_PER_HOUR = "requests_per_hour",
  REQUESTS_PER_DAY = "requests_per_day",
}

// ========== Access Schema ==========

/**
 * 作用域类型
 * access.scope_type
 */
export enum ScopeType {
  TENANT = "TENANT",
  APP = "APP",
  GLOBAL = "GLOBAL",
}

/**
 * 主体类型
 * access.subject_type
 */
export enum SubjectType {
  USER = "USER",
  CLIENT = "CLIENT",
}

/**
 * 授权类型
 * access.grant_type
 */
export enum GrantType {
  ROLE = "ROLE",
  PERMISSION = "PERMISSION",
}

/**
 * 授权效果
 * access.grant_effect
 */
export enum GrantEffect {
  ALLOW = "ALLOW",
  DENY = "DENY",
}

// ========== Audit Schema ==========

/**
 * 操作者类型
 * audit.actor_type
 */
export enum ActorType {
  USER = "user",
  SERVICE = "service",
  SYSTEM = "system",
  ADMIN = "admin",
}

/**
 * 结果类型
 * audit.result_type
 */
export enum ResultType {
  SUCCESS = "success",
  FAILURE = "failure",
  DENY = "deny",
  ERROR = "error",
}

/**
 * 风险级别
 * audit.risk_level
 */
export enum RiskLevel {
  LOW = "low",
  MEDIUM = "medium",
  HIGH = "high",
  CRITICAL = "critical",
}

/**
 * 数据分类
 * audit.data_classification
 */
export enum DataClassification {
  PUBLIC = "public",
  INTERNAL = "internal",
  CONFIDENTIAL = "confidential",
  RESTRICTED = "restricted",
}

/**
 * 保留操作
 * audit.retention_action
 */
export enum RetentionAction {
  ARCHIVE = "archive",
  DELETE = "delete",
  EXPORT = "export",
}
