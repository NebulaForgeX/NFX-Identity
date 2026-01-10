// Auth Domain Types - 基于 NFX-ID Backend

export interface Session {
  id: string;
  sessionId: string;
  tenantId: string;
  userId: string;
  appId?: string;
  clientId?: string;
  createdAt: string;
  lastSeenAt: string;
  expiresAt: string;
  ip?: string;
  uaHash?: string;
  deviceId?: string;
  deviceFingerprint?: string;
  deviceName?: string;
  revokedAt?: string;
  revokeReason?: string;
  revokedBy?: string;
  updatedAt: string;
}

export interface UserCredential {
  id: string;
  userId: string;
  credentialType: string;
  credentialValue: string;
  status: string;
  verifiedAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface MFAFactor {
  id: string;
  userId: string;
  factorType: string;
  factorId: string;
  status: string;
  verifiedAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface RefreshToken {
  id: string;
  token: string;
  sessionId: string;
  userId: string;
  clientId?: string;
  status: string;
  expiresAt?: string;
  revokedAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface PasswordReset {
  id: string;
  resetToken: string;
  userId: string;
  status: string;
  expiresAt?: string;
  usedAt?: string;
  createdAt: string;
  updatedAt: string;
}

export interface PasswordHistory {
  id: string;
  userId: string;
  passwordHash: string;
  createdAt: string;
}

export interface LoginAttempt {
  id: string;
  userId?: string;
  identifier: string;
  ipAddress?: string;
  userAgent?: string;
  success: boolean;
  attemptedAt: string;
  createdAt: string;
}

export interface AccountLockout {
  id: string;
  userId: string;
  reason: string;
  lockedUntil?: string;
  status: string;
  createdAt: string;
  updatedAt: string;
}

export interface TrustedDevice {
  id: string;
  userId: string;
  deviceId: string;
  deviceName?: string;
  deviceFingerprint?: string;
  ipAddress?: string;
  lastUsedAt?: string;
  createdAt: string;
}
