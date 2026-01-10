// Auth Request Types - 基于 NFX-ID Backend

// ========== 会话相关 ==========

export interface CreateSessionRequest {
  sessionId: string;
  tenantId: string;
  userId: string;
  appId?: string;
  clientId?: string;
  expiresAt: string;
  ip?: string;
  uaHash?: string;
  deviceId?: string;
  deviceFingerprint?: string;
  deviceName?: string;
}

export interface RevokeSessionRequest {
  revokeReason: string;
  revokedBy: string;
}

// ========== 用户凭证相关 ==========

export interface CreateUserCredentialRequest {
  userId: string;
  tenantId: string;
  credentialType: string;
  passwordHash?: string;
  hashAlg?: string;
  hashParams?: Record<string, unknown>;
  status?: string;
  mustChangePassword?: boolean;
}

export interface UpdateUserCredentialRequest {
  status?: string;
  mustChangePassword?: boolean;
}

// ========== MFA 因子相关 ==========

export interface CreateMFAFactorRequest {
  userId: string;
  factorType: string;
  factorId: string;
  status?: string;
}

export interface UpdateMFAFactorRequest {
  status?: string;
}

// ========== 刷新令牌相关 ==========

export interface CreateRefreshTokenRequest {
  token: string;
  sessionId: string;
  userId: string;
  clientId?: string;
  expiresAt?: string;
}

export interface UpdateRefreshTokenRequest {
  status?: string;
  expiresAt?: string;
}

// ========== 密码重置相关 ==========

export interface CreatePasswordResetRequest {
  resetToken: string;
  userId: string;
  expiresAt?: string;
}

export interface UpdatePasswordResetRequest {
  status?: string;
  usedAt?: string;
}

// ========== 密码历史相关 ==========

export interface CreatePasswordHistoryRequest {
  userId: string;
  passwordHash: string;
}

// ========== 登录尝试相关 ==========

export interface CreateLoginAttemptRequest {
  userId?: string;
  identifier: string;
  ipAddress?: string;
  userAgent?: string;
  success: boolean;
  attemptedAt: string;
}

// ========== 账户锁定相关 ==========

export interface CreateAccountLockoutRequest {
  userId: string;
  reason: string;
  lockedUntil?: string;
  status?: string;
}

export interface UpdateAccountLockoutRequest {
  reason?: string;
  lockedUntil?: string;
  status?: string;
}

// ========== 受信任设备相关 ==========

export interface CreateTrustedDeviceRequest {
  userId: string;
  deviceId: string;
  deviceName?: string;
  deviceFingerprint?: string;
  ipAddress?: string;
}
