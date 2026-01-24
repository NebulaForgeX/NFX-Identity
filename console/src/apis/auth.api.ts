// Auth API - 基于 NFX-ID Backend

import type {
  AccountLockout,
  BaseResponse,
  CreateAccountLockoutRequest,
  CreateLoginAttemptRequest,
  CreateMFAFactorRequest,
  CreatePasswordHistoryRequest,
  CreatePasswordResetRequest,
  CreateRefreshTokenRequest,
  CreateSessionRequest,
  CreateTrustedDeviceRequest,
  CreateUserCredentialRequest,
  DataResponse,
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

import { protectedClient, publicClient } from "./clients";
import { URL_PATHS } from "./ip";

// ========== 会话相关 ==========

// 创建会话
export const CreateSession = async (params: CreateSessionRequest): Promise<Session> => {
  const { data } = await protectedClient.post<DataResponse<Session>>(URL_PATHS.AUTH.CREATE_SESSION, params);
  return data.data;
};

// 根据 ID 获取会话
export const GetSession = async (id: string): Promise<Session> => {
  const url = URL_PATHS.AUTH.GET_SESSION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Session>>(url);
  return data.data;
};

// 撤销会话
export const RevokeSession = async (sessionId: string, params: RevokeSessionRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.REVOKE_SESSION.replace(":session_id", sessionId);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除会话
export const DeleteSession = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.DELETE_SESSION.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 用户凭证相关 ==========

// 创建用户凭证
export const CreateUserCredential = async (params: CreateUserCredentialRequest): Promise<UserCredential> => {
  const { data } = await protectedClient.post<DataResponse<UserCredential>>(
    URL_PATHS.AUTH.CREATE_USER_CREDENTIAL,
    params,
  );
  return data.data;
};

// 根据 ID 获取用户凭证
export const GetUserCredential = async (id: string): Promise<UserCredential> => {
  const url = URL_PATHS.AUTH.GET_USER_CREDENTIAL.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<UserCredential>>(url);
  return data.data;
};

// 更新用户凭证
export const UpdateUserCredential = async (id: string, params: UpdateUserCredentialRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.UPDATE_USER_CREDENTIAL.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除用户凭证
export const DeleteUserCredential = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.DELETE_USER_CREDENTIAL.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== MFA 因子相关 ==========

// 创建 MFA 因子
export const CreateMFAFactor = async (params: CreateMFAFactorRequest): Promise<MFAFactor> => {
  const { data } = await protectedClient.post<DataResponse<MFAFactor>>(URL_PATHS.AUTH.CREATE_MFA_FACTOR, params);
  return data.data;
};

// 根据 ID 获取 MFA 因子
export const GetMFAFactor = async (id: string): Promise<MFAFactor> => {
  const url = URL_PATHS.AUTH.GET_MFA_FACTOR.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<MFAFactor>>(url);
  return data.data;
};

// 更新 MFA 因子
export const UpdateMFAFactor = async (id: string, params: UpdateMFAFactorRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.UPDATE_MFA_FACTOR.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除 MFA 因子
export const DeleteMFAFactor = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.DELETE_MFA_FACTOR.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 刷新令牌相关 ==========

// 创建刷新令牌
export const CreateRefreshToken = async (params: CreateRefreshTokenRequest): Promise<RefreshToken> => {
  const { data } = await protectedClient.post<DataResponse<RefreshToken>>(URL_PATHS.AUTH.CREATE_REFRESH_TOKEN, params);
  return data.data;
};

// 根据 ID 获取刷新令牌
export const GetRefreshToken = async (id: string): Promise<RefreshToken> => {
  const url = URL_PATHS.AUTH.GET_REFRESH_TOKEN.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<RefreshToken>>(url);
  return data.data;
};

// 更新刷新令牌
export const UpdateRefreshToken = async (id: string, params: UpdateRefreshTokenRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.UPDATE_REFRESH_TOKEN.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除刷新令牌
export const DeleteRefreshToken = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.DELETE_REFRESH_TOKEN.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 密码重置相关 ==========

// 创建密码重置
export const CreatePasswordReset = async (params: CreatePasswordResetRequest): Promise<PasswordReset> => {
  const { data } = await protectedClient.post<DataResponse<PasswordReset>>(
    URL_PATHS.AUTH.CREATE_PASSWORD_RESET,
    params,
  );
  return data.data;
};

// 根据 ID 获取密码重置
export const GetPasswordReset = async (id: string): Promise<PasswordReset> => {
  const url = URL_PATHS.AUTH.GET_PASSWORD_RESET.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<PasswordReset>>(url);
  return data.data;
};

// 更新密码重置
export const UpdatePasswordReset = async (id: string, params: UpdatePasswordResetRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.UPDATE_PASSWORD_RESET.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除密码重置
export const DeletePasswordReset = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.DELETE_PASSWORD_RESET.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 密码历史相关 ==========

// 创建密码历史
export const CreatePasswordHistory = async (params: CreatePasswordHistoryRequest): Promise<PasswordHistory> => {
  const { data } = await protectedClient.post<DataResponse<PasswordHistory>>(
    URL_PATHS.AUTH.CREATE_PASSWORD_HISTORY,
    params,
  );
  return data.data;
};

// 根据 ID 获取密码历史
export const GetPasswordHistory = async (id: string): Promise<PasswordHistory> => {
  const url = URL_PATHS.AUTH.GET_PASSWORD_HISTORY.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<PasswordHistory>>(url);
  return data.data;
};

// ========== 登录尝试相关 ==========

// 创建登录尝试
export const CreateLoginAttempt = async (params: CreateLoginAttemptRequest): Promise<LoginAttempt> => {
  const { data } = await protectedClient.post<DataResponse<LoginAttempt>>(URL_PATHS.AUTH.CREATE_LOGIN_ATTEMPT, params);
  return data.data;
};

// 根据 ID 获取登录尝试
export const GetLoginAttempt = async (id: string): Promise<LoginAttempt> => {
  const url = URL_PATHS.AUTH.GET_LOGIN_ATTEMPT.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<LoginAttempt>>(url);
  return data.data;
};

// 删除登录尝试
export const DeleteLoginAttempt = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.DELETE_LOGIN_ATTEMPT.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 账户锁定相关 ==========

// 创建账户锁定
export const CreateAccountLockout = async (params: CreateAccountLockoutRequest): Promise<AccountLockout> => {
  const { data } = await protectedClient.post<DataResponse<AccountLockout>>(
    URL_PATHS.AUTH.CREATE_ACCOUNT_LOCKOUT,
    params,
  );
  return data.data;
};

// 根据 ID 获取账户锁定
export const GetAccountLockout = async (id: string): Promise<AccountLockout> => {
  const url = URL_PATHS.AUTH.GET_ACCOUNT_LOCKOUT.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<AccountLockout>>(url);
  return data.data;
};

// 更新账户锁定
export const UpdateAccountLockout = async (id: string, params: UpdateAccountLockoutRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.UPDATE_ACCOUNT_LOCKOUT.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除账户锁定
export const DeleteAccountLockout = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.DELETE_ACCOUNT_LOCKOUT.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 受信任设备相关 ==========

// 创建受信任设备
export const CreateTrustedDevice = async (params: CreateTrustedDeviceRequest): Promise<TrustedDevice> => {
  const { data } = await protectedClient.post<DataResponse<TrustedDevice>>(
    URL_PATHS.AUTH.CREATE_TRUSTED_DEVICE,
    params,
  );
  return data.data;
};

// 根据 ID 获取受信任设备
export const GetTrustedDevice = async (id: string): Promise<TrustedDevice> => {
  const url = URL_PATHS.AUTH.GET_TRUSTED_DEVICE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<TrustedDevice>>(url);
  return data.data;
};

// 删除受信任设备
export const DeleteTrustedDevice = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.AUTH.DELETE_TRUSTED_DEVICE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 登录和刷新 Token 相关 ==========

// 登录请求类型
export interface LoginRequest {
  loginType: "email" | "phone";
  email?: string;
  phone?: string;
  countryCode?: string;
  password: string;
}

// 登录响应类型
export interface LoginResponse {
  accessToken: string;
  refreshToken: string;
  expiresIn: number;
  userId: string;
}

// 刷新 Token 请求类型
export interface RefreshTokenRequest {
  refreshToken: string;
}

// 刷新 Token 响应类型
export interface RefreshTokenResponse {
  accessToken: string;
  refreshToken: string;
  expiresIn: number;
}

// 登录（使用 publicClient，不需要认证）
export const Login = async (params: LoginRequest): Promise<LoginResponse> => {
  const { data } = await publicClient.post<DataResponse<LoginResponse>>("/auth/login", params);
  return data.data;
};

// 刷新 Token（使用 publicClient，不需要认证）
export const RefreshAccessToken = async (params: RefreshTokenRequest): Promise<RefreshTokenResponse> => {
  const { data } = await publicClient.post<DataResponse<RefreshTokenResponse>>("/auth/refresh", params);
  return data.data;
};
