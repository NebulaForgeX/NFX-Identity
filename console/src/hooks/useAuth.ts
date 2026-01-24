import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";

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
  RevokeSession,
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
  return useMutation({
    mutationFn: async (params: CreateSessionRequest) => {
      return await CreateSession(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_SESSIONS);
      showSuccess("会话创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建会话失败，请稍后重试。" + error.message);
    },
  });
};

// 撤销会话
export const useRevokeSession = () => {
  return useMutation({
    mutationFn: async (params: { sessionId: string; data: RevokeSessionRequest }) => {
      return await RevokeSession(params.sessionId, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_SESSIONS);
      authEventEmitter.emit(authEvents.INVALIDATE_SESSION, variables.sessionId);
      showSuccess("会话撤销成功！");
    },
    onError: (error: AxiosError) => {
      showError("撤销会话失败，请稍后重试。" + error.message);
    },
  });
};

// 删除会话
export const useDeleteSession = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteSession(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_SESSIONS);
      authEventEmitter.emit(authEvents.INVALIDATE_SESSION, id);
      showSuccess("会话删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除会话失败，请稍后重试。" + error.message);
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
  return useMutation({
    mutationFn: async (params: CreateUserCredentialRequest) => {
      return await CreateUserCredential(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_USER_CREDENTIALS);
      showSuccess("用户凭证创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建用户凭证失败，请稍后重试。" + error.message);
    },
  });
};

// 更新用户凭证
export const useUpdateUserCredential = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateUserCredentialRequest }) => {
      return await UpdateUserCredential(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_USER_CREDENTIALS);
      authEventEmitter.emit(authEvents.INVALIDATE_USER_CREDENTIAL, variables.id);
      showSuccess("用户凭证更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新用户凭证失败，请稍后重试。" + error.message);
    },
  });
};

// 删除用户凭证
export const useDeleteUserCredential = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteUserCredential(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_USER_CREDENTIALS);
      authEventEmitter.emit(authEvents.INVALIDATE_USER_CREDENTIAL, id);
      showSuccess("用户凭证删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除用户凭证失败，请稍后重试。" + error.message);
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
  return useMutation({
    mutationFn: async (params: CreateMFAFactorRequest) => {
      return await CreateMFAFactor(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_MFA_FACTORS);
      showSuccess("MFA 因子创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 MFA 因子失败，请稍后重试。" + error.message);
    },
  });
};

// 更新 MFA 因子
export const useUpdateMFAFactor = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateMFAFactorRequest }) => {
      return await UpdateMFAFactor(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_MFA_FACTORS);
      authEventEmitter.emit(authEvents.INVALIDATE_MFA_FACTOR, variables.id);
      showSuccess("MFA 因子更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新 MFA 因子失败，请稍后重试。" + error.message);
    },
  });
};

// 删除 MFA 因子
export const useDeleteMFAFactor = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteMFAFactor(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_MFA_FACTORS);
      authEventEmitter.emit(authEvents.INVALIDATE_MFA_FACTOR, id);
      showSuccess("MFA 因子删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除 MFA 因子失败，请稍后重试。" + error.message);
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
  return useMutation({
    mutationFn: async (params: CreateRefreshTokenRequest) => {
      return await CreateRefreshToken(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_REFRESH_TOKENS);
      showSuccess("刷新令牌创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建刷新令牌失败，请稍后重试。" + error.message);
    },
  });
};

// 更新刷新令牌
export const useUpdateRefreshToken = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateRefreshTokenRequest }) => {
      return await UpdateRefreshToken(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_REFRESH_TOKENS);
      authEventEmitter.emit(authEvents.INVALIDATE_REFRESH_TOKEN, variables.id);
      showSuccess("刷新令牌更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新刷新令牌失败，请稍后重试。" + error.message);
    },
  });
};

// 删除刷新令牌
export const useDeleteRefreshToken = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteRefreshToken(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_REFRESH_TOKENS);
      authEventEmitter.emit(authEvents.INVALIDATE_REFRESH_TOKEN, id);
      showSuccess("刷新令牌删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除刷新令牌失败，请稍后重试。" + error.message);
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
  return useMutation({
    mutationFn: async (params: CreatePasswordResetRequest) => {
      return await CreatePasswordReset(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_RESETS);
      showSuccess("密码重置创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建密码重置失败，请稍后重试。" + error.message);
    },
  });
};

// 更新密码重置
export const useUpdatePasswordReset = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdatePasswordResetRequest }) => {
      return await UpdatePasswordReset(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_RESETS);
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_RESET, variables.id);
      showSuccess("密码重置更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新密码重置失败，请稍后重试。" + error.message);
    },
  });
};

// 删除密码重置
export const useDeletePasswordReset = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeletePasswordReset(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_RESETS);
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_RESET, id);
      showSuccess("密码重置删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除密码重置失败，请稍后重试。" + error.message);
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
  return useMutation({
    mutationFn: async (params: CreatePasswordHistoryRequest) => {
      return await CreatePasswordHistory(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_PASSWORD_HISTORIES);
      showSuccess("密码历史创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建密码历史失败，请稍后重试。" + error.message);
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
  return useMutation({
    mutationFn: async (params: CreateLoginAttemptRequest) => {
      return await CreateLoginAttempt(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_LOGIN_ATTEMPTS);
      showSuccess("登录尝试创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建登录尝试失败，请稍后重试。" + error.message);
    },
  });
};

// 删除登录尝试
export const useDeleteLoginAttempt = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteLoginAttempt(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_LOGIN_ATTEMPTS);
      authEventEmitter.emit(authEvents.INVALIDATE_LOGIN_ATTEMPT, id);
      showSuccess("登录尝试删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除登录尝试失败，请稍后重试。" + error.message);
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
  return useMutation({
    mutationFn: async (params: CreateAccountLockoutRequest) => {
      return await CreateAccountLockout(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS);
      showSuccess("账户锁定创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建账户锁定失败，请稍后重试。" + error.message);
    },
  });
};

// 更新账户锁定
export const useUpdateAccountLockout = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateAccountLockoutRequest }) => {
      return await UpdateAccountLockout(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      authEventEmitter.emit(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS);
      authEventEmitter.emit(authEvents.INVALIDATE_ACCOUNT_LOCKOUT, variables.id);
      showSuccess("账户锁定更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新账户锁定失败，请稍后重试。" + error.message);
    },
  });
};

// 删除账户锁定
export const useDeleteAccountLockout = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteAccountLockout(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_ACCOUNT_LOCKOUTS);
      authEventEmitter.emit(authEvents.INVALIDATE_ACCOUNT_LOCKOUT, id);
      showSuccess("账户锁定删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除账户锁定失败，请稍后重试。" + error.message);
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
  return useMutation({
    mutationFn: async (params: CreateTrustedDeviceRequest) => {
      return await CreateTrustedDevice(params);
    },
    onSuccess: () => {
      authEventEmitter.emit(authEvents.INVALIDATE_TRUSTED_DEVICES);
      showSuccess("受信任设备创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建受信任设备失败，请稍后重试。" + error.message);
    },
  });
};

// 删除受信任设备
export const useDeleteTrustedDevice = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteTrustedDevice(id);
    },
    onSuccess: (_, id) => {
      authEventEmitter.emit(authEvents.INVALIDATE_TRUSTED_DEVICES);
      authEventEmitter.emit(authEvents.INVALIDATE_TRUSTED_DEVICE, id);
      showSuccess("受信任设备删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除受信任设备失败，请稍后重试。" + error.message);
    },
  });
};

// ========== 登录相关 ==========

// 通过邮箱登录
export const useLoginByEmail = () => {
  return useMutation({
    mutationFn: async (params: { email: string; password: string }) => {
      const { Login } = await import("@/apis/auth.api");
      const response = await Login({
        loginType: "email",
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
      showSuccess("登录成功！");
      // 触发登录成功事件，由App.tsx监听并跳转
      authEventEmitter.emit(authEvents.LOGIN_SUCCESS);
    },
    onError: (error: AxiosError) => {
      const errorMessage = (error.response?.data as { message?: string })?.message;
      if (errorMessage?.includes("locked")) {
        showError("账户已被锁定，请联系管理员");
      } else if (errorMessage?.includes("invalid")) {
        showError("邮箱或密码错误");
      } else {
        showError("登录失败，请稍后重试");
      }
    },
  });
};

// 通过手机号登录
export const useLoginByPhone = () => {
  return useMutation({
    mutationFn: async (params: { phone: string; password: string; countryCode?: string }) => {
      const { Login } = await import("@/apis/auth.api");
      const response = await Login({
        loginType: "phone",
        phone: params.phone,
        countryCode: params.countryCode,
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
      showSuccess("登录成功！");
      // 触发登录成功事件，由App.tsx监听并跳转
      authEventEmitter.emit(authEvents.LOGIN_SUCCESS);
    },
    onError: (error: AxiosError) => {
      const errorMessage = (error.response?.data as { message?: string })?.message;
      if (errorMessage?.includes("locked")) {
        showError("账户已被锁定，请联系管理员");
      } else if (errorMessage?.includes("invalid")) {
        showError("手机号或密码错误");
      } else {
        showError("登录失败，请稍后重试");
      }
    },
  });
};

// ========== 注册相关 ==========

// 发送验证码
export const useSendVerificationCode = () => {
  return useMutation({
    mutationFn: async (params: { email: string }) => {
      // TODO: 实现发送验证码逻辑
      throw new Error("发送验证码功能尚未实现");
    },
    onError: (error: AxiosError) => {
      showError("发送验证码失败，请稍后重试。" + error.message);
    },
  });
};

// 注册
export const useSignup = () => {
  return useMutation({
    mutationFn: async (params: {
      email: string;
      password: string;
      verificationCode: string;
    }) => {
      // TODO: 实现注册逻辑
      // 1. 验证验证码
      // 2. 创建用户
      // 3. 创建用户凭证
      // 4. 创建会话
      throw new Error("注册功能尚未实现");
    },
    onError: (error: AxiosError) => {
      showError("注册失败，请稍后重试。" + error.message);
    },
  });
};
