import type { QueryClient } from "@tanstack/react-query";
import { authEventEmitter, authEvents } from "@/events/auth";
import {
  AUTH_ACCOUNT_LOCKOUT,
  AUTH_ACCOUNT_LOCKOUT_LIST,
  AUTH_LOGIN_ATTEMPT,
  AUTH_LOGIN_ATTEMPT_LIST,
  AUTH_MFA_FACTOR,
  AUTH_MFA_FACTOR_LIST,
  AUTH_PASSWORD_HISTORY,
  AUTH_PASSWORD_HISTORY_LIST,
  AUTH_PASSWORD_RESET,
  AUTH_PASSWORD_RESET_LIST,
  AUTH_REFRESH_TOKEN,
  AUTH_REFRESH_TOKEN_LIST,
  AUTH_SESSION,
  AUTH_SESSION_LIST,
  AUTH_TRUSTED_DEVICE,
  AUTH_TRUSTED_DEVICE_LIST,
  AUTH_USER_CREDENTIAL,
  AUTH_USER_CREDENTIAL_LIST,
} from "@/constants";
import { AuthStore } from "@/stores/authStore";
import { ChatStore } from "@/stores/chatStore";

type EventCb = (...args: unknown[]) => void;

/**
 * Auth 相关的缓存失效事件处理
 */
export const useAuthCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateSessions = () => queryClient.invalidateQueries({ queryKey: AUTH_SESSION_LIST });
  const handleInvalidateSession = (item: string) => queryClient.invalidateQueries({ queryKey: AUTH_SESSION(item) });
  const handleInvalidateUserCredentials = () => queryClient.invalidateQueries({ queryKey: AUTH_USER_CREDENTIAL_LIST });
  const handleInvalidateUserCredential = (item: string) => queryClient.invalidateQueries({ queryKey: AUTH_USER_CREDENTIAL(item) });
  const handleInvalidateMFAFactors = () => queryClient.invalidateQueries({ queryKey: AUTH_MFA_FACTOR_LIST });
  const handleInvalidateMFAFactor = (item: string) => queryClient.invalidateQueries({ queryKey: AUTH_MFA_FACTOR(item) });
  const handleInvalidateRefreshTokens = () => queryClient.invalidateQueries({ queryKey: AUTH_REFRESH_TOKEN_LIST });
  const handleInvalidateRefreshToken = (item: string) => queryClient.invalidateQueries({ queryKey: AUTH_REFRESH_TOKEN(item) });
  const handleInvalidatePasswordResets = () => queryClient.invalidateQueries({ queryKey: AUTH_PASSWORD_RESET_LIST });
  const handleInvalidatePasswordReset = (item: string) => queryClient.invalidateQueries({ queryKey: AUTH_PASSWORD_RESET(item) });
  const handleInvalidatePasswordHistories = () => queryClient.invalidateQueries({ queryKey: AUTH_PASSWORD_HISTORY_LIST });
  const handleInvalidatePasswordHistory = (item: string) => queryClient.invalidateQueries({ queryKey: AUTH_PASSWORD_HISTORY(item) });
  const handleInvalidateLoginAttempts = () => queryClient.invalidateQueries({ queryKey: AUTH_LOGIN_ATTEMPT_LIST });
  const handleInvalidateLoginAttempt = (item: string) => queryClient.invalidateQueries({ queryKey: AUTH_LOGIN_ATTEMPT(item) });
  const handleInvalidateAccountLockouts = () => queryClient.invalidateQueries({ queryKey: AUTH_ACCOUNT_LOCKOUT_LIST });
  const handleInvalidateAccountLockout = (item: string) => queryClient.invalidateQueries({ queryKey: AUTH_ACCOUNT_LOCKOUT(item) });
  const handleInvalidateTrustedDevices = () => queryClient.invalidateQueries({ queryKey: AUTH_TRUSTED_DEVICE_LIST });
  const handleInvalidateTrustedDevice = (item: string) => queryClient.invalidateQueries({ queryKey: AUTH_TRUSTED_DEVICE(item) });
  
  // 处理退出登录事件 - 清理所有缓存和 stores
  const handleLogout = () => {
    // 清理认证状态
    AuthStore.getState().clearAuth();
    // 清理 React Query 缓存
    queryClient.clear();
    // 清理 ChatStore 状态
    ChatStore.getState().clearUnreadCount();
    ChatStore.getState().setTotalMessages(0);
  };

  authEventEmitter.on(authEvents.INVALIDATE_SESSIONS, handleInvalidateSessions as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_SESSION, handleInvalidateSession as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_USER_CREDENTIALS, handleInvalidateUserCredentials as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_USER_CREDENTIAL, handleInvalidateUserCredential as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_MFA_FACTORS, handleInvalidateMFAFactors as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_MFA_FACTOR, handleInvalidateMFAFactor as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_REFRESH_TOKENS, handleInvalidateRefreshTokens as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_REFRESH_TOKEN, handleInvalidateRefreshToken as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_RESETS, handleInvalidatePasswordResets as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_RESET, handleInvalidatePasswordReset as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_HISTORIES, handleInvalidatePasswordHistories as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_PASSWORD_HISTORY, handleInvalidatePasswordHistory as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_LOGIN_ATTEMPTS, handleInvalidateLoginAttempts as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_LOGIN_ATTEMPT, handleInvalidateLoginAttempt as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS, handleInvalidateAccountLockouts as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_ACCOUNT_LOCKOUT, handleInvalidateAccountLockout as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_TRUSTED_DEVICES, handleInvalidateTrustedDevices as EventCb);
  authEventEmitter.on(authEvents.INVALIDATE_TRUSTED_DEVICE, handleInvalidateTrustedDevice as EventCb);
  authEventEmitter.on(authEvents.LOGOUT, handleLogout as EventCb);

  return () => {
    authEventEmitter.off(authEvents.INVALIDATE_SESSIONS, handleInvalidateSessions as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_SESSION, handleInvalidateSession as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_USER_CREDENTIALS, handleInvalidateUserCredentials as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_USER_CREDENTIAL, handleInvalidateUserCredential as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_MFA_FACTORS, handleInvalidateMFAFactors as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_MFA_FACTOR, handleInvalidateMFAFactor as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_REFRESH_TOKENS, handleInvalidateRefreshTokens as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_REFRESH_TOKEN, handleInvalidateRefreshToken as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_RESETS, handleInvalidatePasswordResets as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_RESET, handleInvalidatePasswordReset as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_HISTORIES, handleInvalidatePasswordHistories as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_PASSWORD_HISTORY, handleInvalidatePasswordHistory as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_LOGIN_ATTEMPTS, handleInvalidateLoginAttempts as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_LOGIN_ATTEMPT, handleInvalidateLoginAttempt as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS, handleInvalidateAccountLockouts as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_ACCOUNT_LOCKOUT, handleInvalidateAccountLockout as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_TRUSTED_DEVICES, handleInvalidateTrustedDevices as EventCb);
    authEventEmitter.off(authEvents.INVALIDATE_TRUSTED_DEVICE, handleInvalidateTrustedDevice as EventCb);
    authEventEmitter.off(authEvents.LOGOUT, handleLogout as EventCb);
  };
};
