// Access API - 基于 NFX-ID Backend

import type {
  Action,
  ActionRequirement,
  BaseResponse,
  CreateActionRequest,
  CreateActionRequirementRequest,
  CreateGrantRequest,
  CreatePermissionRequest,
  CreateRolePermissionRequest,
  CreateRoleRequest,
  CreateScopePermissionRequest,
  CreateScopeRequest,
  DataResponse,
  Grant,
  Permission,
  Role,
  RolePermission,
  Scope,
  ScopePermission,
  UpdateGrantRequest,
  UpdatePermissionRequest,
  UpdateRoleRequest,
  UpdateScopeRequest,
} from "@/types";

import type { AxiosError } from "axios";

import { protectedClient } from "./clients";
import { URL_PATHS } from "./ip";

// ========== 角色相关 ==========

// 创建角色
export const CreateRole = async (params: CreateRoleRequest): Promise<Role> => {
  const { data } = await protectedClient.post<DataResponse<Role>>(URL_PATHS.ACCESS.CREATE_ROLE, params);
  return data.data;
};

// 根据 ID 获取角色
export const GetRole = async (id: string): Promise<Role> => {
  const url = URL_PATHS.ACCESS.GET_ROLE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Role>>(url);
  return data.data;
};

// 根据 Key 获取角色（404 视为未找到，返回 null，不抛错）
export const GetRoleByKey = async (key: string): Promise<Role | null> => {
  const url = URL_PATHS.ACCESS.GET_ROLE_BY_KEY.replace(":key", key);
  try {
    const { data } = await protectedClient.get<DataResponse<Role>>(url);
    return data.data;
  } catch (e) {
    const err = e as AxiosError & { response?: { status?: number } };
    if (err.response?.status === 404) return null;
    throw e;
  }
};

// 更新角色
export const UpdateRole = async (id: string, params: UpdateRoleRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.ACCESS.UPDATE_ROLE.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除角色
export const DeleteRole = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.ACCESS.DELETE_ROLE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 权限相关 ==========

// 创建权限
export const CreatePermission = async (params: CreatePermissionRequest): Promise<Permission> => {
  const { data } = await protectedClient.post<DataResponse<Permission>>(URL_PATHS.ACCESS.CREATE_PERMISSION, params);
  return data.data;
};

// 根据 ID 获取权限
export const GetPermission = async (id: string): Promise<Permission> => {
  const url = URL_PATHS.ACCESS.GET_PERMISSION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Permission>>(url);
  return data.data;
};

// 根据 Key 获取权限（404 视为未找到，返回 null，不抛错）
export const GetPermissionByKey = async (key: string): Promise<Permission | null> => {
  const url = URL_PATHS.ACCESS.GET_PERMISSION_BY_KEY.replace(":key", key);
  try {
    const { data } = await protectedClient.get<DataResponse<Permission>>(url);
    return data.data;
  } catch (e) {
    const err = e as AxiosError & { response?: { status?: number } };
    if (err.response?.status === 404) return null;
    throw e;
  }
};

// 更新权限
export const UpdatePermission = async (id: string, params: UpdatePermissionRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.ACCESS.UPDATE_PERMISSION.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除权限
export const DeletePermission = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.ACCESS.DELETE_PERMISSION.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 作用域相关 ==========

// 创建作用域
export const CreateScope = async (params: CreateScopeRequest): Promise<Scope> => {
  const { data } = await protectedClient.post<DataResponse<Scope>>(URL_PATHS.ACCESS.CREATE_SCOPE, params);
  return data.data;
};

// 根据 Scope 获取作用域
export const GetScope = async (scope: string): Promise<Scope> => {
  const url = URL_PATHS.ACCESS.GET_SCOPE.replace(":scope", scope);
  const { data } = await protectedClient.get<DataResponse<Scope>>(url);
  return data.data;
};

// 更新作用域
export const UpdateScope = async (scope: string, params: UpdateScopeRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.ACCESS.UPDATE_SCOPE.replace(":scope", scope);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除作用域
export const DeleteScope = async (scope: string): Promise<BaseResponse> => {
  const url = URL_PATHS.ACCESS.DELETE_SCOPE.replace(":scope", scope);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 授权相关 ==========

// 创建授权
export const CreateGrant = async (params: CreateGrantRequest): Promise<Grant> => {
  const { data } = await protectedClient.post<DataResponse<Grant>>(URL_PATHS.ACCESS.CREATE_GRANT, params);
  return data.data;
};

// 根据主体获取授权列表
export const GetGrantsBySubject = async (params: {
  subject_type: string;
  subject_id: string;
  tenant_id?: string;
}): Promise<Grant[]> => {
  const { data } = await protectedClient.get<DataResponse<Grant[]>>(URL_PATHS.ACCESS.GET_GRANTS_BY_SUBJECT, {
    params,
  });
  return data.data;
};

// 根据 ID 获取授权
export const GetGrant = async (id: string): Promise<Grant> => {
  const url = URL_PATHS.ACCESS.GET_GRANT.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Grant>>(url);
  return data.data;
};

// 更新授权
export const UpdateGrant = async (id: string, params: UpdateGrantRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.ACCESS.UPDATE_GRANT.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除授权
export const DeleteGrant = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.ACCESS.DELETE_GRANT.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 角色权限关联相关 ==========

// 创建角色权限关联
export const CreateRolePermission = async (params: CreateRolePermissionRequest): Promise<RolePermission> => {
  const { data } = await protectedClient.post<DataResponse<RolePermission>>(
    URL_PATHS.ACCESS.CREATE_ROLE_PERMISSION,
    params,
  );
  return data.data;
};

// 根据角色ID获取角色权限列表
export const GetRolePermissionsByRole = async (roleId: string): Promise<RolePermission[]> => {
  const url = URL_PATHS.ACCESS.GET_ROLE_PERMISSIONS_BY_ROLE.replace(":role_id", roleId);
  const { data } = await protectedClient.get<DataResponse<RolePermission[]>>(url);
  return data.data;
};

// 根据 ID 获取角色权限关联
export const GetRolePermission = async (id: string): Promise<RolePermission> => {
  const url = URL_PATHS.ACCESS.GET_ROLE_PERMISSION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<RolePermission>>(url);
  return data.data;
};

// 删除角色权限关联
export const DeleteRolePermission = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.ACCESS.DELETE_ROLE_PERMISSION.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 作用域权限关联相关 ==========

// 创建作用域权限关联
export const CreateScopePermission = async (params: CreateScopePermissionRequest): Promise<ScopePermission> => {
  const { data } = await protectedClient.post<DataResponse<ScopePermission>>(
    URL_PATHS.ACCESS.CREATE_SCOPE_PERMISSION,
    params,
  );
  return data.data;
};

// 根据 ID 获取作用域权限关联
export const GetScopePermission = async (id: string): Promise<ScopePermission> => {
  const url = URL_PATHS.ACCESS.GET_SCOPE_PERMISSION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<ScopePermission>>(url);
  return data.data;
};

// 删除作用域权限关联
export const DeleteScopePermission = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.ACCESS.DELETE_SCOPE_PERMISSION.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== Action 相关 ==========

export const CreateAction = async (params: CreateActionRequest): Promise<Action> => {
  const { data } = await protectedClient.post<DataResponse<Action>>(URL_PATHS.ACCESS.CREATE_ACTION, params);
  return data.data;
};

export const GetAction = async (id: string): Promise<Action> => {
  const url = URL_PATHS.ACCESS.GET_ACTION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Action>>(url);
  return data.data;
};

// 根据 Key 获取 Action（404 视为未找到，返回 null，不抛错）
export const GetActionByKey = async (key: string): Promise<Action | null> => {
  const url = URL_PATHS.ACCESS.GET_ACTION_BY_KEY.replace(":key", key);
  try {
    const { data } = await protectedClient.get<DataResponse<Action>>(url);
    return data.data;
  } catch (e) {
    const err = e as AxiosError & { response?: { status?: number } };
    if (err.response?.status === 404) return null;
    throw e;
  }
};

// ========== ActionRequirement 相关（Permission 关联的 Action） ==========

export const CreateActionRequirement = async (
  params: CreateActionRequirementRequest,
): Promise<ActionRequirement> => {
  const { data } = await protectedClient.post<DataResponse<ActionRequirement>>(
    URL_PATHS.ACCESS.CREATE_ACTION_REQUIREMENT,
    params,
  );
  return data.data;
};

export const GetActionRequirementsByPermission = async (
  permissionId: string,
): Promise<ActionRequirement[]> => {
  const url = URL_PATHS.ACCESS.GET_ACTION_REQUIREMENTS_BY_PERMISSION.replace(
    ":permission_id",
    permissionId,
  );
  const { data } = await protectedClient.get<DataResponse<ActionRequirement[]>>(url);
  return data.data;
};

export const GetActionRequirement = async (id: string): Promise<ActionRequirement> => {
  const url = URL_PATHS.ACCESS.GET_ACTION_REQUIREMENT.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<ActionRequirement>>(url);
  return data.data;
};

export const DeleteActionRequirement = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.ACCESS.DELETE_ACTION_REQUIREMENT.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};
