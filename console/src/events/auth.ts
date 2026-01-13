export const authEvents = {
  // Session 相关
  INVALIDATE_SESSION: "AUTH:INVALIDATE_SESSION",
  INVALIDATE_SESSIONS: "AUTH:INVALIDATE_SESSIONS",

  // UserCredential 相关
  INVALIDATE_USER_CREDENTIAL: "AUTH:INVALIDATE_USER_CREDENTIAL",
  INVALIDATE_USER_CREDENTIALS: "AUTH:INVALIDATE_USER_CREDENTIALS",

  // MFAFactor 相关
  INVALIDATE_MFA_FACTOR: "AUTH:INVALIDATE_MFA_FACTOR",
  INVALIDATE_MFA_FACTORS: "AUTH:INVALIDATE_MFA_FACTORS",

  // RefreshToken 相关
  INVALIDATE_REFRESH_TOKEN: "AUTH:INVALIDATE_REFRESH_TOKEN",
  INVALIDATE_REFRESH_TOKENS: "AUTH:INVALIDATE_REFRESH_TOKENS",

  // PasswordReset 相关
  INVALIDATE_PASSWORD_RESET: "AUTH:INVALIDATE_PASSWORD_RESET",
  INVALIDATE_PASSWORD_RESETS: "AUTH:INVALIDATE_PASSWORD_RESETS",

  // PasswordHistory 相关
  INVALIDATE_PASSWORD_HISTORY: "AUTH:INVALIDATE_PASSWORD_HISTORY",
  INVALIDATE_PASSWORD_HISTORIES: "AUTH:INVALIDATE_PASSWORD_HISTORIES",

  // LoginAttempt 相关
  INVALIDATE_LOGIN_ATTEMPT: "AUTH:INVALIDATE_LOGIN_ATTEMPT",
  INVALIDATE_LOGIN_ATTEMPTS: "AUTH:INVALIDATE_LOGIN_ATTEMPTS",

  // AccountLockout 相关
  INVALIDATE_ACCOUNT_LOCKOUT: "AUTH:INVALIDATE_ACCOUNT_LOCKOUT",
  INVALIDATE_ACCOUNT_LOCKOUTS: "AUTH:INVALIDATE_ACCOUNT_LOCKOUTS",

  // TrustedDevice 相关
  INVALIDATE_TRUSTED_DEVICE: "AUTH:INVALIDATE_TRUSTED_DEVICE",
  INVALIDATE_TRUSTED_DEVICES: "AUTH:INVALIDATE_TRUSTED_DEVICES",
} as const;

type AuthEvent = (typeof authEvents)[keyof typeof authEvents];

class AuthEventEmitter {
  private listeners: Record<AuthEvent, Set<Function>> = {
    [authEvents.INVALIDATE_SESSION]: new Set<Function>(),
    [authEvents.INVALIDATE_SESSIONS]: new Set<Function>(),
    [authEvents.INVALIDATE_USER_CREDENTIAL]: new Set<Function>(),
    [authEvents.INVALIDATE_USER_CREDENTIALS]: new Set<Function>(),
    [authEvents.INVALIDATE_MFA_FACTOR]: new Set<Function>(),
    [authEvents.INVALIDATE_MFA_FACTORS]: new Set<Function>(),
    [authEvents.INVALIDATE_REFRESH_TOKEN]: new Set<Function>(),
    [authEvents.INVALIDATE_REFRESH_TOKENS]: new Set<Function>(),
    [authEvents.INVALIDATE_PASSWORD_RESET]: new Set<Function>(),
    [authEvents.INVALIDATE_PASSWORD_RESETS]: new Set<Function>(),
    [authEvents.INVALIDATE_PASSWORD_HISTORY]: new Set<Function>(),
    [authEvents.INVALIDATE_PASSWORD_HISTORIES]: new Set<Function>(),
    [authEvents.INVALIDATE_LOGIN_ATTEMPT]: new Set<Function>(),
    [authEvents.INVALIDATE_LOGIN_ATTEMPTS]: new Set<Function>(),
    [authEvents.INVALIDATE_ACCOUNT_LOCKOUT]: new Set<Function>(),
    [authEvents.INVALIDATE_ACCOUNT_LOCKOUTS]: new Set<Function>(),
    [authEvents.INVALIDATE_TRUSTED_DEVICE]: new Set<Function>(),
    [authEvents.INVALIDATE_TRUSTED_DEVICES]: new Set<Function>(),
  };

  on(event: AuthEvent, callback: Function) {
    this.listeners[event].add(callback);
  }

  off(event: AuthEvent, callback: Function) {
    this.listeners[event].delete(callback);
  }

  emit(event: AuthEvent, ...args: unknown[]) {
    this.listeners[event].forEach((callback) => {
      callback(...args);
    });
  }
}

export const authEventEmitter = new AuthEventEmitter();
