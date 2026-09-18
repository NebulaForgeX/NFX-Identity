// Domain Enums used by auth.domain.ts (Identity login center).

/**
 * Credential type
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
 * Credential status
 * auth.credential_status
 */
export enum CredentialStatus {
  ACTIVE = "active",
  DISABLED = "disabled",
  EXPIRED = "expired",
}

/**
 * MFA type
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
 * Revoke reason
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
 * Session revoke reason
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
 * Login failure code
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
 * Lock reason
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
 * Password reset delivery
 * auth.reset_delivery
 */
export enum ResetDelivery {
  EMAIL = "email",
  SMS = "sms",
}

/**
 * Password reset status
 * auth.reset_status
 */
export enum ResetStatus {
  ISSUED = "issued",
  USED = "used",
  EXPIRED = "expired",
  REVOKED = "revoked",
}
