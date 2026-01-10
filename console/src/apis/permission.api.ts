// Permission API - 基于 NFX-ID Permission Service

import type {
  AuthorizationCode,
  CreateAuthorizationCodeParams,
  UseAuthorizationCodeParams,
  Permission,
  PermissionList,
  CreatePermissionParams,
  UpdatePermissionParams,
  PermissionQueryParams,
  UserPermission,
  AssignPermissionParams,
  RevokePermissionParams,
  CheckPermissionParams,
  CheckPermissionResponse,
} from "@/types/domain";
import type { DataResponse } from "@/types/api";

import { protectedClient, publicClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// ========== 公开路由 ==========

// 登录（Permission Service 的登录，返回权限信息）
export interface PermissionLoginParams {
  type: "password" | "code";
  identifier: string; // username, email 或 phone
  password?: string; // 当 type=password 时
  code?: string; // 当 type=code 时
}

export interface PermissionLoginResponse {
  accessToken: string;
  refreshToken: string;
  userId: string;
  username: string;
  email: string;
  phone: string;
  permissions: UserPermission[];
  permissionTags: string[];
}

export const PermissionLogin = async (params: PermissionLoginParams): Promise<PermissionLoginResponse> => {
  const { data } = await publicClient.post<DataResponse<PermissionLoginResponse>>(
    URL_PATHS.PERMISSION.LOGIN,
    params,
  );
  return data.data;
};

// 注册（用于 Identity-Admin 平台）
export interface PermissionRegisterParams {
  email: string;
  verification_code: string; // 邮箱验证码（6位）
  authorization_code: string; // 授权码/邀请码
  password: string;
}

export const PermissionRegister = async (params: PermissionRegisterParams): Promise<PermissionLoginResponse> => {
  const { data } = await publicClient.post<DataResponse<PermissionLoginResponse>>(
    URL_PATHS.PERMISSION.REGISTER,
    params,
  );
  return data.data;
};

// ========== 需要认证的路由 ==========

// ========== Permission 管理 ==========

// 创建权限
export const CreatePermission = async (params: CreatePermissionParams): Promise<Permission> => {
  const { data } = await protectedClient.post<DataResponse<Permission>>(
    URL_PATHS.PERMISSION.CREATE_PERMISSION,
    params,
  );
  return data.data;
};

// 获取权限列表
export const GetPermissions = async (
  params?: PermissionQueryParams,
): Promise<{ permissions: Permission[]; total: number }> => {
  const response = await protectedClient.get<DataResponse<Permission[]>>(URL_PATHS.PERMISSION.GET_PERMISSIONS, {
    params: params,
  });
  const total = (response.data.meta?.total as number) || response.data.data.length;
  return {
    permissions: response.data.data,
    total,
  };
};

// 根据 ID 获取权限
export const GetPermission = async (id: string): Promise<Permission> => {
  const url = URL_PATHS.PERMISSION.GET_PERMISSION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Permission>>(url);
  return data.data;
};

// 根据 Tag 获取权限
export const GetPermissionByTag = async (tag: string): Promise<Permission> => {
  const url = URL_PATHS.PERMISSION.GET_PERMISSION_BY_TAG.replace(":tag", tag);
  const { data } = await protectedClient.get<DataResponse<Permission>>(url);
  return data.data;
};

// 更新权限
export const UpdatePermission = async (id: string, params: UpdatePermissionParams): Promise<void> => {
  const url = URL_PATHS.PERMISSION.UPDATE_PERMISSION.replace(":id", id);
  await protectedClient.put<DataResponse<null>>(url, params);
};

// 删除权限
export const DeletePermission = async (id: string): Promise<void> => {
  const url = URL_PATHS.PERMISSION.DELETE_PERMISSION.replace(":id", id);
  await protectedClient.delete(url);
};

// ========== User Permission 管理 ==========

// 分配权限给用户
export const AssignUserPermission = async (params: AssignPermissionParams): Promise<void> => {
  await protectedClient.post<DataResponse<null>>(URL_PATHS.PERMISSION.ASSIGN_USER_PERMISSION, params);
};

// 撤销用户权限
export const RevokeUserPermission = async (params: RevokePermissionParams): Promise<void> => {
  // DELETE 请求，后端使用 BodyParser，所以传递 data
  await protectedClient.delete<DataResponse<null>>(URL_PATHS.PERMISSION.REVOKE_USER_PERMISSION, {
    data: params,
  });
};

// 获取用户的所有权限
export const GetUserPermissions = async (userId: string): Promise<UserPermission[]> => {
  const url = URL_PATHS.PERMISSION.GET_USER_PERMISSIONS.replace(":user_id", userId);
  const { data } = await protectedClient.get<DataResponse<UserPermission[]>>(url);
  return data.data;
};

// 获取用户的所有权限标签
export const GetUserPermissionTags = async (userId: string): Promise<string[]> => {
  const url = URL_PATHS.PERMISSION.GET_USER_PERMISSION_TAGS.replace(":user_id", userId);
  const { data } = await protectedClient.get<DataResponse<string[]>>(url);
  return data.data;
};

// 检查用户是否有某个权限
export const CheckUserPermission = async (params: CheckPermissionParams): Promise<CheckPermissionResponse> => {
  const { data } = await protectedClient.post<DataResponse<CheckPermissionResponse>>(
    URL_PATHS.PERMISSION.CHECK_USER_PERMISSION,
    params,
  );
  return data.data;
};

// ========== Authorization Code 管理 ==========

// 创建授权码
export const CreateAuthorizationCode = async (params: CreateAuthorizationCodeParams): Promise<AuthorizationCode> => {
  const { data } = await protectedClient.post<DataResponse<AuthorizationCode>>(
    URL_PATHS.PERMISSION.CREATE_AUTHORIZATION_CODE,
    params,
  );
  return data.data;
};

// 根据 ID 获取授权码
export const GetAuthorizationCode = async (id: string): Promise<AuthorizationCode> => {
  const url = URL_PATHS.PERMISSION.GET_AUTHORIZATION_CODE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<AuthorizationCode>>(url);
  return data.data;
};

// 根据 Code 获取授权码
export const GetAuthorizationCodeByCode = async (code: string): Promise<AuthorizationCode> => {
  const url = URL_PATHS.PERMISSION.GET_AUTHORIZATION_CODE_BY_CODE.replace(":code", code);
  const { data } = await protectedClient.get<DataResponse<AuthorizationCode>>(url);
  return data.data;
};

// 使用授权码
export const UseAuthorizationCode = async (params: UseAuthorizationCodeParams): Promise<void> => {
  await protectedClient.post<DataResponse<null>>(URL_PATHS.PERMISSION.USE_AUTHORIZATION_CODE, params);
};

// 删除授权码
export const DeleteAuthorizationCode = async (id: string): Promise<void> => {
  const url = URL_PATHS.PERMISSION.DELETE_AUTHORIZATION_CODE.replace(":id", id);
  await protectedClient.delete(url);
};

// 激活授权码
export const ActivateAuthorizationCode = async (id: string): Promise<void> => {
  const url = URL_PATHS.PERMISSION.ACTIVATE_AUTHORIZATION_CODE.replace(":id", id);
  await protectedClient.post<DataResponse<null>>(url);
};

// 停用授权码
export const DeactivateAuthorizationCode = async (id: string): Promise<void> => {
  const url = URL_PATHS.PERMISSION.DEACTIVATE_AUTHORIZATION_CODE.replace(":id", id);
  await protectedClient.post<DataResponse<null>>(url);
};
