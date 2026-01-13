import type { QueryClient } from "@tanstack/react-query";
import { authEventEmitter, authEvents } from "@/events/auth";
import { AUTH_QUERY_KEY_PREFIXES } from "@/constants";

/**
 * Auth 相关的缓存失效事件处理
 */
export const useAuthCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateSessions = () => queryClient.invalidateQueries({ queryKey: AUTH_QUERY_KEY_PREFIXES.SESSIONS });
  const handleInvalidateSession = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUTH_QUERY_KEY_PREFIXES.SESSION, item] });
  const handleInvalidateUserCredentials = () => queryClient.invalidateQueries({ queryKey: AUTH_QUERY_KEY_PREFIXES.USER_CREDENTIALS });
  const handleInvalidateUserCredential = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUTH_QUERY_KEY_PREFIXES.USER_CREDENTIAL, item] });
  const handleInvalidateMFAFactors = () => queryClient.invalidateQueries({ queryKey: AUTH_QUERY_KEY_PREFIXES.MFA_FACTORS });
  const handleInvalidateMFAFactor = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUTH_QUERY_KEY_PREFIXES.MFA_FACTOR, item] });
  const handleInvalidateRefreshTokens = () => queryClient.invalidateQueries({ queryKey: AUTH_QUERY_KEY_PREFIXES.REFRESH_TOKENS });
  const handleInvalidateRefreshToken = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUTH_QUERY_KEY_PREFIXES.REFRESH_TOKEN, item] });
  const handleInvalidatePasswordResets = () => queryClient.invalidateQueries({ queryKey: AUTH_QUERY_KEY_PREFIXES.PASSWORD_RESETS });
  const handleInvalidatePasswordReset = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUTH_QUERY_KEY_PREFIXES.PASSWORD_RESET, item] });
  const handleInvalidatePasswordHistories = () => queryClient.invalidateQueries({ queryKey: AUTH_QUERY_KEY_PREFIXES.PASSWORD_HISTORIES });
  const handleInvalidatePasswordHistory = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUTH_QUERY_KEY_PREFIXES.PASSWORD_HISTORY, item] });
  const handleInvalidateLoginAttempts = () => queryClient.invalidateQueries({ queryKey: AUTH_QUERY_KEY_PREFIXES.LOGIN_ATTEMPTS });
  const handleInvalidateLoginAttempt = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUTH_QUERY_KEY_PREFIXES.LOGIN_ATTEMPT, item] });
  const handleInvalidateAccountLockouts = () => queryClient.invalidateQueries({ queryKey: AUTH_QUERY_KEY_PREFIXES.ACCOUNT_LOCKOUTS });
  const handleInvalidateAccountLockout = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUTH_QUERY_KEY_PREFIXES.ACCOUNT_LOCKOUT, item] });
  const handleInvalidateTrustedDevices = () => queryClient.invalidateQueries({ queryKey: AUTH_QUERY_KEY_PREFIXES.TRUSTED_DEVICES });
  const handleInvalidateTrustedDevice = (item: string) => queryClient.invalidateQueries({ queryKey: [...AUTH_QUERY_KEY_PREFIXES.TRUSTED_DEVICE, item] });

  // 注册监听器
  authEventEmitter.on(authEvents.INVALIDATE_SESSIONS, handleInvalidateSessions);
  authEventEmitter.on(authEvents.INVALIDATE_SESSION, handleInvalidateSession);
  authEventEmitter.on(authEvents.INVALIDATE_USER_CREDENTIALS, handleInvalidateUserCredentials);
  authEventEmitter.on(authEvents.INVALIDATE_USER_CREDENTIAL, handleInvalidateUserCredential);
  authEventEmitter.on(authEvents.INVALIDATE_MFA_FACTORS, handleInvalidateMFAFactors);
  authEventEmitter.on(authEvents.INVALIDATE_MFA_FACTOR, handleInvalidateMFAFactor);
  authEventEmitter.on(authEvents.INVALIDATE_REFRESH_TOKENS, handleInvalidateRefreshTokens);
  authEventEmitter.on(authEvents.INVALIDATE_REFRESH_TOKEN, handleInvalidateRefreshToken);
  authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_RESETS, handleInvalidatePasswordResets);
  authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_RESET, handleInvalidatePasswordReset);
  authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_HISTORIES, handleInvalidatePasswordHistories);
  authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_HISTORY, handleInvalidatePasswordHistory);
  authEventEmitter.on(authEvents.INVALIDATE_LOGIN_ATTEMPTS, handleInvalidateLoginAttempts);
  authEventEmitter.on(authEvents.INVALIDATE_LOGIN_ATTEMPT, handleInvalidateLoginAttempt);
  authEventEmitter.on(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS, handleInvalidateAccountLockouts);
  authEventEmitter.on(authEvents.INVALIDATE_ACCOUNT_LOCKOUT, handleInvalidateAccountLockout);
  authEventEmitter.on(authEvents.INVALIDATE_TRUSTED_DEVICES, handleInvalidateTrustedDevices);
  authEventEmitter.on(authEvents.INVALIDATE_TRUSTED_DEVICE, handleInvalidateTrustedDevice);

  // 清理监听器
  return () => {
    authEventEmitter.off(authEvents.INVALIDATE_SESSIONS, handleInvalidateSessions);
    authEventEmitter.off(authEvents.INVALIDATE_SESSION, handleInvalidateSession);
    authEventEmitter.off(authEvents.INVALIDATE_USER_CREDENTIALS, handleInvalidateUserCredentials);
    authEventEmitter.off(authEvents.INVALIDATE_USER_CREDENTIAL, handleInvalidateUserCredential);
    authEventEmitter.off(authEvents.INVALIDATE_MFA_FACTORS, handleInvalidateMFAFactors);
    authEventEmitter.off(authEvents.INVALIDATE_MFA_FACTOR, handleInvalidateMFAFactor);
    authEventEmitter.off(authEvents.INVALIDATE_REFRESH_TOKENS, handleInvalidateRefreshTokens);
    authEventEmitter.off(authEvents.INVALIDATE_REFRESH_TOKEN, handleInvalidateRefreshToken);
    authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_RESETS, handleInvalidatePasswordResets);
    authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_RESET, handleInvalidatePasswordReset);
    authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_HISTORIES, handleInvalidatePasswordHistories);
    authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_HISTORY, handleInvalidatePasswordHistory);
    authEventEmitter.off(authEvents.INVALIDATE_LOGIN_ATTEMPTS, handleInvalidateLoginAttempts);
    authEventEmitter.off(authEvents.INVALIDATE_LOGIN_ATTEMPT, handleInvalidateLoginAttempt);
    authEventEmitter.off(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS, handleInvalidateAccountLockouts);
    authEventEmitter.off(authEvents.INVALIDATE_ACCOUNT_LOCKOUT, handleInvalidateAccountLockout);
    authEventEmitter.off(authEvents.INVALIDATE_TRUSTED_DEVICES, handleInvalidateTrustedDevices);
    authEventEmitter.off(authEvents.INVALIDATE_TRUSTED_DEVICE, handleInvalidateTrustedDevice);
  };
};
