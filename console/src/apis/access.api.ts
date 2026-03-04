// Access API - 基于 NFX-ID Backend

import type { BaseResponse, DataResponse } from "nfx-ui/types";
import type {
  Action,
  ActionRequirement,
  CreateActionRequest,
  CreateActionRequirementRequest,
  CreateGrantRequest,
  CreatePermissionRequest,
  CreateRolePermissionRequest,
  CreateRoleRequest,
  CreateScopePermissionRequest,
  CreateScopeRequest,
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

import { protectedClient } from "./clients";
import { URL_PATHS } from "./ip";

// ========== 角色相关 ==========

// 创建角色
export const CreateRole = async (params: CreateRoleRequest): Promise<Role> => {
  const { data } = await protectedClient.post<DataResponse<Role>>(URL_PATHS.ACCESS.roles, params);
  return data.data;
};

// 根据 ID 获取角色
export const GetRole = async (id: string): Promise<Role> => {
  const { data } = await protectedClient.get<DataResponse<Role>>(URL_PATHS.ACCESS.roles.byId(id));
  return data.data;
};

// 根据 Key 获取角色（后端 404 时返回 err_code + message，axios 抛错，不返回 null）
export const GetRoleByKey = async (key: string): Promise<Role> => {
  const { data } = await protectedClient.get<DataResponse<Role>>(URL_PATHS.ACCESS.roles.byKey(key));
  return data.data;
};

// 更新角色
export const UpdateRole = async (id: string, params: UpdateRoleRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.ACCESS.roles.byId(id), params);
  return data;
};

// 删除角色
export const DeleteRole = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.ACCESS.roles.byId(id));
  return data;
};

// ========== 权限相关 ==========

// 创建权限
export const CreatePermission = async (params: CreatePermissionRequest): Promise<Permission> => {
  const { data } = await protectedClient.post<DataResponse<Permission>>(URL_PATHS.ACCESS.permissions, params);
  return data.data;
};

// 根据 ID 获取权限
export const GetPermission = async (id: string): Promise<Permission> => {
  const { data } = await protectedClient.get<DataResponse<Permission>>(URL_PATHS.ACCESS.permissions.byId(id));
  return data.data;
};

// 根据 Key 获取权限（后端 404 时返回 err_code + message，axios 抛错，不返回 null）
export const GetPermissionByKey = async (key: string): Promise<Permission> => {
  const { data } = await protectedClient.get<DataResponse<Permission>>(URL_PATHS.ACCESS.permissions.byKey(key));
  return data.data;
};

// 更新权限
export const UpdatePermission = async (id: string, params: UpdatePermissionRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.ACCESS.permissions.byId(id), params);
  return data;
};

// 删除权限
export const DeletePermission = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.ACCESS.permissions.byId(id));
  return data;
};

// ========== 作用域相关 ==========

// 创建作用域
export const CreateScope = async (params: CreateScopeRequest): Promise<Scope> => {
  const { data } = await protectedClient.post<DataResponse<Scope>>(URL_PATHS.ACCESS.scopes, params);
  return data.data;
};

// 根据 Scope 获取作用域
export const GetScope = async (scope: string): Promise<Scope> => {
  const { data } = await protectedClient.get<DataResponse<Scope>>(URL_PATHS.ACCESS.scopes.byScope(scope));
  return data.data;
};

// 更新作用域
export const UpdateScope = async (scope: string, params: UpdateScopeRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.ACCESS.scopes.byScope(scope), params);
  return data;
};

// 删除作用域
export const DeleteScope = async (scope: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.ACCESS.scopes.byScope(scope));
  return data;
};

// ========== 授权相关 ==========

// 创建授权
export const CreateGrant = async (params: CreateGrantRequest): Promise<Grant> => {
  const { data } = await protectedClient.post<DataResponse<Grant>>(URL_PATHS.ACCESS.grants, params);
  return data.data;
};

// 根据主体获取授权列表
export const GetGrantsBySubject = async (params: {
  subject_type: string;
  subject_id: string;
  tenant_id?: string;
}): Promise<Grant[]> => {
  const { data } = await protectedClient.get<DataResponse<Grant[]>>(URL_PATHS.ACCESS.grants, {
    params,
  });
  return data.data;
};

// 根据 ID 获取授权
export const GetGrant = async (id: string): Promise<Grant> => {
  const { data } = await protectedClient.get<DataResponse<Grant>>(URL_PATHS.ACCESS.grants.byId(id));
  return data.data;
};

// 更新授权
export const UpdateGrant = async (id: string, params: UpdateGrantRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(URL_PATHS.ACCESS.grants.byId(id), params);
  return data;
};

// 删除授权
export const DeleteGrant = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.ACCESS.grants.byId(id));
  return data;
};

// ========== 角色权限关联相关 ==========

// 创建角色权限关联
export const CreateRolePermission = async (params: CreateRolePermissionRequest): Promise<RolePermission> => {
  const { data } = await protectedClient.post<DataResponse<RolePermission>>(
    URL_PATHS.ACCESS.rolePermissions,
    params,
  );
  return data.data;
};

// 根据角色ID获取角色权限列表
export const GetRolePermissionsByRole = async (roleId: string): Promise<RolePermission[]> => {
  const { data } = await protectedClient.get<DataResponse<RolePermission[]>>(
    URL_PATHS.ACCESS.rolePermissions.byRole(roleId),
  );
  return data.data;
};

// 根据 ID 获取角色权限关联
export const GetRolePermission = async (id: string): Promise<RolePermission> => {
  const { data } = await protectedClient.get<DataResponse<RolePermission>>(
    URL_PATHS.ACCESS.rolePermissions.byId(id),
  );
  return data.data;
};

// 删除角色权限关联
export const DeleteRolePermission = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.ACCESS.rolePermissions.byId(id));
  return data;
};

// ========== 作用域权限关联相关 ==========

// 创建作用域权限关联
export const CreateScopePermission = async (params: CreateScopePermissionRequest): Promise<ScopePermission> => {
  const { data } = await protectedClient.post<DataResponse<ScopePermission>>(
    URL_PATHS.ACCESS.scopePermissions,
    params,
  );
  return data.data;
};

// 根据 ID 获取作用域权限关联
export const GetScopePermission = async (id: string): Promise<ScopePermission> => {
  const { data } = await protectedClient.get<DataResponse<ScopePermission>>(
    URL_PATHS.ACCESS.scopePermissions.byId(id),
  );
  return data.data;
};

// 删除作用域权限关联
export const DeleteScopePermission = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(URL_PATHS.ACCESS.scopePermissions.byId(id));
  return data;
};

// ========== Action 相关 ==========

export const CreateAction = async (params: CreateActionRequest): Promise<Action> => {
  const { data } = await protectedClient.post<DataResponse<Action>>(URL_PATHS.ACCESS.actions, params);
  return data.data;
};

export const GetAction = async (id: string): Promise<Action> => {
  const { data } = await protectedClient.get<DataResponse<Action>>(URL_PATHS.ACCESS.actions.byId(id));
  return data.data;
};

// 根据 Key 获取 Action（后端 404 时返回 err_code + message，axios 抛错，不返回 null）
export const GetActionByKey = async (key: string): Promise<Action> => {
  const { data } = await protectedClient.get<DataResponse<Action>>(URL_PATHS.ACCESS.actions.byKey(key));
  return data.data;
};

// ========== ActionRequirement 相关（Permission 关联的 Action） ==========

export const CreateActionRequirement = async (
  params: CreateActionRequirementRequest,
): Promise<ActionRequirement> => {
  const { data } = await protectedClient.post<DataResponse<ActionRequirement>>(
    URL_PATHS.ACCESS.actionRequirements,
    params,
  );
  return data.data;
};

export const GetActionRequirementsByPermission = async (
  permissionId: string,
): Promise<ActionRequirement[]> => {
  const { data } = await protectedClient.get<DataResponse<ActionRequirement[]>>(
    URL_PATHS.ACCESS.actionRequirements.byPermission(permissionId),
  );
  return data.data;
};

export const GetActionRequirement = async (id: string): Promise<ActionRequirement> => {
  const { data } = await protectedClient.get<DataResponse<ActionRequirement>>(
    URL_PATHS.ACCESS.actionRequirements.byId(id),
  );
  return data.data;
};

export const DeleteActionRequirement = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.ACCESS.actionRequirements.byId(id),
  );
  return data;
};
