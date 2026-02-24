import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import { useTranslation } from "react-i18next";

import {
  CreateAccountLockout,
  CreateLoginAttempt,
  CreateMFAFactor,
  CreatePasswordHistory,
  CreatePasswordReset,
  CreateRefreshToken,
  CreateSession,
  CreateTrustedDevice,
  CreateUserCredential,
  DeleteAccountLockout,
  DeleteLoginAttempt,
  DeleteMFAFactor,
  DeletePasswordReset,
  DeleteRefreshToken,
  DeleteSession,
  DeleteTrustedDevice,
  DeleteUserCredential,
  GetAccountLockout,
  GetLoginAttempt,
  GetMFAFactor,
  GetPasswordHistory,
  GetPasswordReset,
  GetRefreshToken,
  GetSession,
  GetTrustedDevice,
  GetUserCredential,
  LoginByEmail,
  LoginByPhone,
  RevokeSession,
  SendVerificationCode,
  Signup,
  UpdateAccountLockout,
  UpdateMFAFactor,
  UpdatePasswordReset,
  UpdateRefreshToken,
  UpdateUserCredential,
} from "@/apis/auth.api";
import type {
  AccountLockout,
  CreateAccountLockoutRequest,
  CreateLoginAttemptRequest,
  CreateMFAFactorRequest,
  CreatePasswordHistoryRequest,
  CreatePasswordResetRequest,
  CreateRefreshTokenRequest,
  CreateSessionRequest,
  CreateTrustedDeviceRequest,
  CreateUserCredentialRequest,
  LoginAttempt,
  MFAFactor,
  PasswordHistory,
  PasswordReset,
  RefreshToken,
  RevokeSessionRequest,
  Session,
  TrustedDevice,
  UpdateAccountLockoutRequest,
  UpdateMFAFactorRequest,
  UpdatePasswordResetRequest,
  UpdateRefreshTokenRequest,
  UpdateUserCredentialRequest,
  UserCredential,
} from "@/types";
import { makeUnifiedQuery } from "@/hooks/core/makeUnifiedQuery";
import { authEventEmitter, authEvents } from "@/events/auth";
import { showError, showSuccess } from "@/stores/modalStore";
import AuthStore from "@/stores/authStore";
import { getApiError, getApiErrorMessage } from "@/utils/apiError";
import {
  AUTH_SESSION,
  AUTH_USER_CREDENTIAL,
  AUTH_MFA_FACTOR,
  AUTH_REFRESH_TOKEN,
  AUTH_PASSWORD_RESET,
  AUTH_PASSWORD_HISTORY,
  AUTH_LOGIN_ATTEMPT,
  AUTH_ACCOUNT_LOCKOUT,
  AUTH_TRUSTED_DEVICE,
} from "@/constants";
import type { UnifiedQueryParams } from "./core/type";

// ========== Session 相关 ==========

// 根据 ID 获取会话
export const useSession = (params: UnifiedQueryParams<Session> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetSession(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUTH_SESSION(id), { id }, options);
};

// 创建会话
export const useCreateSession = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: CreateSessionRequest) => {
      return await CreateSession(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_SESSIONS);
      showSuccess(t("session.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("session.createError")));
    },
  });
};

// 撤销会话
export const useRevokeSession = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: { sessionId: string; data: RevokeSessionRequest }) => {
      return await RevokeSession(params.sessionId, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_SESSIONS);
      authEventEmitter.emit(authEvents.INVALIDATE_SESSION, variables.sessionId);
      showSuccess(t("session.revokeSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("session.revokeError")));
    },
  });
};

// 删除会话
export const useDeleteSession = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteSession(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_SESSIONS);
      authEventEmitter.emit(authEvents.INVALIDATE_SESSION, id);
      showSuccess(t("session.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("session.deleteError")));
    },
  });
};

// ========== UserCredential 相关 ==========

// 根据 ID 获取用户凭证
export const useUserCredential = (params: UnifiedQueryParams<UserCredential> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetUserCredential(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUTH_USER_CREDENTIAL(id), { id }, options);
};

// 创建用户凭证
export const useCreateUserCredential = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: CreateUserCredentialRequest) => {
      return await CreateUserCredential(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_USER_CREDENTIALS);
      showSuccess(t("userCredential.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("userCredential.createError")));
    },
  });
};

// 更新用户凭证
export const useUpdateUserCredential = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserCredentialRequest }) => {
      return await UpdateUserCredential(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_USER_CREDENTIALS);
      authEventEmitter.emit(authEvents.INVALIDATE_USER_CREDENTIAL, variables.id);
      showSuccess(t("userCredential.updateSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("userCredential.updateError")));
    },
  });
};

// 删除用户凭证
export const useDeleteUserCredential = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteUserCredential(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_USER_CREDENTIALS);
      authEventEmitter.emit(authEvents.INVALIDATE_USER_CREDENTIAL, id);
      showSuccess(t("userCredential.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("userCredential.deleteError")));
    },
  });
};

// ========== MFAFactor 相关 ==========

// 根据 ID 获取 MFA 因子
export const useMFAFactor = (params: UnifiedQueryParams<MFAFactor> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetMFAFactor(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUTH_MFA_FACTOR(id), { id }, options);
};

// 创建 MFA 因子
export const useCreateMFAFactor = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: CreateMFAFactorRequest) => {
      return await CreateMFAFactor(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_MFA_FACTORS);
      showSuccess(t("mfaFactor.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("mfaFactor.createError")));
    },
  });
};

// 更新 MFA 因子
export const useUpdateMFAFactor = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateMFAFactorRequest }) => {
      return await UpdateMFAFactor(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_MFA_FACTORS);
      authEventEmitter.emit(authEvents.INVALIDATE_MFA_FACTOR, variables.id);
      showSuccess(t("mfaFactor.updateSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("mfaFactor.updateError")));
    },
  });
};

// 删除 MFA 因子
export const useDeleteMFAFactor = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteMFAFactor(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_MFA_FACTORS);
      authEventEmitter.emit(authEvents.INVALIDATE_MFA_FACTOR, id);
      showSuccess(t("mfaFactor.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("mfaFactor.deleteError")));
    },
  });
};

// ========== RefreshToken 相关 ==========

// 根据 ID 获取刷新令牌
export const useRefreshToken = (params: UnifiedQueryParams<RefreshToken> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetRefreshToken(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUTH_REFRESH_TOKEN(id), { id }, options);
};

// 创建刷新令牌
export const useCreateRefreshToken = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: CreateRefreshTokenRequest) => {
      return await CreateRefreshToken(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_REFRESH_TOKENS);
      showSuccess(t("refreshToken.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("refreshToken.createError")));
    },
  });
};

// 更新刷新令牌
export const useUpdateRefreshToken = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateRefreshTokenRequest }) => {
      return await UpdateRefreshToken(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_REFRESH_TOKENS);
      authEventEmitter.emit(authEvents.INVALIDATE_REFRESH_TOKEN, variables.id);
      showSuccess(t("refreshToken.updateSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("refreshToken.updateError")));
    },
  });
};

// 删除刷新令牌
export const useDeleteRefreshToken = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteRefreshToken(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_REFRESH_TOKENS);
      authEventEmitter.emit(authEvents.INVALIDATE_REFRESH_TOKEN, id);
      showSuccess(t("refreshToken.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("refreshToken.deleteError")));
    },
  });
};

// ========== PasswordReset 相关 ==========

// 根据 ID 获取密码重置
export const usePasswordReset = (params: UnifiedQueryParams<PasswordReset> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetPasswordReset(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUTH_PASSWORD_RESET(id), { id }, options);
};

// 创建密码重置
export const useCreatePasswordReset = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: CreatePasswordResetRequest) => {
      return await CreatePasswordReset(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_RESETS);
      showSuccess(t("passwordReset.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("passwordReset.createError")));
    },
  });
};

// 更新密码重置
export const useUpdatePasswordReset = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdatePasswordResetRequest }) => {
      return await UpdatePasswordReset(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_RESETS);
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_RESET, variables.id);
      showSuccess(t("passwordReset.updateSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("passwordReset.updateError")));
    },
  });
};

// 删除密码重置
export const useDeletePasswordReset = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeletePasswordReset(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_RESETS);
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_RESET, id);
      showSuccess(t("passwordReset.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("passwordReset.deleteError")));
    },
  });
};

// ========== PasswordHistory 相关 ==========

// 根据 ID 获取密码历史
export const usePasswordHistory = (params: UnifiedQueryParams<PasswordHistory> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetPasswordHistory(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUTH_PASSWORD_HISTORY(id), { id }, options);
};

// 创建密码历史
export const useCreatePasswordHistory = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: CreatePasswordHistoryRequest) => {
      return await CreatePasswordHistory(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_HISTORIES);
      showSuccess(t("passwordHistory.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("passwordHistory.createError")));
    },
  });
};

// ========== LoginAttempt 相关 ==========

// 根据 ID 获取登录尝试
export const useLoginAttempt = (params: UnifiedQueryParams<LoginAttempt> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetLoginAttempt(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUTH_LOGIN_ATTEMPT(id), { id }, options);
};

// 创建登录尝试
export const useCreateLoginAttempt = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: CreateLoginAttemptRequest) => {
      return await CreateLoginAttempt(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_LOGIN_ATTEMPTS);
      showSuccess(t("loginAttempt.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("loginAttempt.createError")));
    },
  });
};

// 删除登录尝试
export const useDeleteLoginAttempt = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteLoginAttempt(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_LOGIN_ATTEMPTS);
      authEventEmitter.emit(authEvents.INVALIDATE_LOGIN_ATTEMPT, id);
      showSuccess(t("loginAttempt.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("loginAttempt.deleteError")));
    },
  });
};

// ========== AccountLockout 相关 ==========

// 根据 ID 获取账户锁定
export const useAccountLockout = (params: UnifiedQueryParams<AccountLockout> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetAccountLockout(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUTH_ACCOUNT_LOCKOUT(id), { id }, options);
};

// 创建账户锁定
export const useCreateAccountLockout = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: CreateAccountLockoutRequest) => {
      return await CreateAccountLockout(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS);
      showSuccess(t("accountLockout.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("accountLockout.createError")));
    },
  });
};

// 更新账户锁定
export const useUpdateAccountLockout = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateAccountLockoutRequest }) => {
      return await UpdateAccountLockout(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS);
      authEventEmitter.emit(authEvents.INVALIDATE_ACCOUNT_LOCKOUT, variables.id);
      showSuccess(t("accountLockout.updateSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("accountLockout.updateError")));
    },
  });
};

// 删除账户锁定
export const useDeleteAccountLockout = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteAccountLockout(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS);
      authEventEmitter.emit(authEvents.INVALIDATE_ACCOUNT_LOCKOUT, id);
      showSuccess(t("accountLockout.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("accountLockout.deleteError")));
    },
  });
};

// ========== TrustedDevice 相关 ==========

// 根据 ID 获取受信任设备
export const useTrustedDevice = (params: UnifiedQueryParams<TrustedDevice> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetTrustedDevice(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(AUTH_TRUSTED_DEVICE(id), { id }, options);
};

// 创建受信任设备
export const useCreateTrustedDevice = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: CreateTrustedDeviceRequest) => {
      return await CreateTrustedDevice(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_TRUSTED_DEVICES);
      showSuccess(t("trustedDevice.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("trustedDevice.createError")));
    },
  });
};

// 删除受信任设备
export const useDeleteTrustedDevice = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteTrustedDevice(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_TRUSTED_DEVICES);
      authEventEmitter.emit(authEvents.INVALIDATE_TRUSTED_DEVICE, id);
      showSuccess(t("trustedDevice.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(getApiErrorMessage(error, t("trustedDevice.deleteError")));
    },
  });
};

// ========== 登录相关 ==========

// 通过邮箱登录
export const useLoginByEmail = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: { email: string; password: string }) => {
      const response = await LoginByEmail({
        email: params.email,
        password: params.password,
      });
      
      // 设置tokens
      AuthStore.getState().setTokens({
        accessToken: response.accessToken,
        refreshToken: response.refreshToken,
      });
      
      // 设置用户ID和认证状态
      AuthStore.getState().setCurrentUserId(response.userId);
      AuthStore.getState().setIsAuthValid(true);
      
      return response;
    },
    onSuccess: () => {
      showSuccess(t("login.success"));
      // 触发登录成功事件，由App.tsx监听并跳转
      authEventEmitter.emit(authEvents.LOGIN_SUCCESS);
    },
    onError: (error: AxiosError) => {
      const body = getApiError(error);
      if (body?.errCode === "ACCOUNT_LOCKED") {
        showError(t("login.accountLocked"));
      } else if (body?.errCode === "INVALID_CREDENTIALS") {
        showError(t("login.invalidCredentials"));
      } else {
        showError(body?.message ?? t("login.failed"));
      }
    },
  });
};

// 通过手机号登录
export const useLoginByPhone = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: { phone: string; password: string; countryCode?: string }) => {
      if (!params.countryCode) {
        throw new Error("countryCode is required");
      }
      const response = await LoginByPhone({
        countryCode: params.countryCode,
        phone: params.phone,
        password: params.password,
      });
      
      // 设置tokens
      AuthStore.getState().setTokens({
        accessToken: response.accessToken,
        refreshToken: response.refreshToken,
      });
      
      // 设置用户ID和认证状态
      AuthStore.getState().setCurrentUserId(response.userId);
      AuthStore.getState().setIsAuthValid(true);
      
      return response;
    },
    onSuccess: () => {
      showSuccess(t("login.success"));
      // 触发登录成功事件，由App.tsx监听并跳转
      authEventEmitter.emit(authEvents.LOGIN_SUCCESS);
    },
    onError: (error: AxiosError) => {
      const body = getApiError(error);
      if (body?.errCode === "ACCOUNT_LOCKED") {
        showError(t("login.accountLocked"));
      } else if (body?.errCode === "INVALID_CREDENTIALS") {
        showError(t("login.invalidPhoneCredentials"));
      } else {
        showError(body?.message ?? t("login.failed"));
      }
    },
  });
};

// ========== 注册相关 ==========

// 发送验证码
export const useSendVerificationCode = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: { email: string }) => {
      return await SendVerificationCode(params);
    },
    onSuccess: () => {
      showSuccess(t("verificationCode.sendSuccess"));
    },
    onError: (error: AxiosError) => {
      const body = getApiError(error);
      showError(body?.message ?? t("verificationCode.sendError"));
    },
  });
};

// 注册
export const useSignup = () => {
  const { t } = useTranslation("hooks.auth");
  return useMutation({
    mutationFn: async (params: {
      email: string;
      password: string;
      verificationCode: string;
    }) => {
      return await Signup(params);
    },
    onSuccess: (data) => {
      // 注册成功后，保存 token 到 store
      AuthStore.getState().setTokens({
        accessToken: data.accessToken,
        refreshToken: data.refreshToken,
      });
      // 设置用户ID和认证状态
      AuthStore.getState().setCurrentUserId(data.userId);
      AuthStore.getState().setIsAuthValid(true);
      showSuccess(t("signup.success"));
      // 触发登录成功事件，由App.tsx监听并跳转
      authEventEmitter.emit(authEvents.LOGIN_SUCCESS);
    },
    onError: (error: AxiosError) => {
      const body = getApiError(error);
      showError(body?.message ?? t("signup.error"));
    },
  });
};
