// Auth API - 基于 NFX-ID

import type {
  CreateUserParams,
  LoginParams,
  LoginResponse,
  RefreshResponse,
  RefreshTokenParams,
  UpdateUserParams,
  User,
  UserQueryParams,
} from "@/types/domain";
import type { DataResponse } from "@/types/api";

import { protectedClient, publicClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// ========== 公开路由 ==========

// 登录
export const Login = async (params: LoginParams): Promise<LoginResponse> => {
  const { data } = await publicClient.post<DataResponse<LoginResponse>>(URL_PATHS.AUTH.LOGIN, params);
  return data.data;
};

// 通过邮箱登录（包装函数）
export const LoginByEmail = async (params: { email: string; password: string }): Promise<LoginResponse & { token?: string }> => {
  const response = await Login({ identifier: params.email, password: params.password });
  // 兼容旧代码，添加 token 字段（指向 accessToken）
  return { ...response, token: response.accessToken };
};

// 通过手机号登录（包装函数）
export const LoginByPhone = async (params: { phone: string; password: string }): Promise<LoginResponse & { token?: string }> => {
  const response = await Login({ identifier: params.phone, password: params.password });
  // 兼容旧代码，添加 token 字段（指向 accessToken）
  return { ...response, token: response.accessToken };
};

// 刷新 Token
export const RefreshToken = async (params: RefreshTokenParams): Promise<RefreshResponse> => {
  const { data } = await publicClient.post<DataResponse<RefreshResponse>>(
    URL_PATHS.AUTH.REFRESH_TOKEN,
    params,
  );
  return data.data;
};

// ========== 需要认证的路由 ==========

// 创建用户
export const CreateUser = async (params: CreateUserParams): Promise<User> => {
  const { data } = await protectedClient.post<DataResponse<User>>(URL_PATHS.AUTH.CREATE_USER, params);
  return data.data;
};

// 获取所有用户
export const GetUsers = async (params?: UserQueryParams): Promise<{ users: User[]; total: number }> => {
  const response = await protectedClient.get<DataResponse<User[]>>(URL_PATHS.AUTH.GET_USERS, {
    params: params,
  });
  // 后端返回 data 是 User[] 数组，total 在 meta 中
  const total = (response.data.meta?.total as number) || response.data.data.length;
  return {
    users: response.data.data,
    total,
  };
};

// 根据 ID 获取用户
export const GetUser = async (id: string): Promise<User> => {
  const url = URL_PATHS.AUTH.GET_USER.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<User>>(url);
  return data.data;
};

// 验证登录状态（检查 token 是否有效）
export const CheckLogin = async (): Promise<void> => {
  // 调用一个需要认证的 API 来验证 token 是否有效
  // 使用 GetUsers 并设置 limit=1 来最小化数据传输
  await protectedClient.get<DataResponse<User[]>>(URL_PATHS.AUTH.GET_USERS, {
    params: { limit: 1 },
  });
  // 如果请求成功（没有抛出 401 错误），说明 token 有效
};

// 更新用户
export const UpdateUser = async (id: string, params: Partial<UpdateUserParams>): Promise<void> => {
  const url = URL_PATHS.AUTH.UPDATE_USER.replace(":id", id);
  await protectedClient.put<DataResponse<null>>(url, params);
};

// 删除用户
export const DeleteUser = async (id: string): Promise<void> => {
  const url = URL_PATHS.AUTH.DELETE_USER.replace(":id", id);
  await protectedClient.delete(url);
};

// 删除用户账户
export const DeleteUserAccount = async (id: string): Promise<void> => {
  const url = URL_PATHS.AUTH.DELETE_USER_ACCOUNT.replace(":id", id);
  await protectedClient.delete(url);
};

// TODO: 以下函数在后端尚未实现，暂时提供 stub 实现

// 注册（Signup - 用户自主注册）
export const Signup = async (params: { email: string; password: string; username?: string }): Promise<LoginResponse & { token?: string }> => {
  // TODO: 实现用户注册功能
  throw new Error("Signup not implemented yet");
};

// 管理员创建用户（Register - 需要认证）
export const Register = async (params: CreateUserParams, avatarFile?: File): Promise<User> => {
  // TODO: 实现管理员创建用户功能（包含头像上传）
  throw new Error("Register not implemented yet");
};

// 发送验证码（发送邮件验证码，存储到 Redis）
export const SendVerificationCode = async (params: { email: string }): Promise<void> => {
  await publicClient.post<DataResponse<null>>(URL_PATHS.AUTH.SEND_VERIFICATION_CODE, params);
};

// 发送验证码到当前用户的邮箱
export const SendVerificationCodeToCurrentEmail = async (): Promise<void> => {
  // TODO: 实现发送验证码到当前用户邮箱功能
  throw new Error("SendVerificationCodeToCurrentEmail not implemented yet");
};

// 更新密码
export const UpdatePassword = async (params: { verificationCode: string; newPassword: string }): Promise<void> => {
  // TODO: 实现更新密码功能
  throw new Error("UpdatePassword not implemented yet");
};

// 更新邮箱
export const UpdateEmail = async (params: { oldEmailCode: string; newEmail: string; newEmailCode: string }): Promise<void> => {
  // TODO: 实现更新邮箱功能
  throw new Error("UpdateEmail not implemented yet");
};
