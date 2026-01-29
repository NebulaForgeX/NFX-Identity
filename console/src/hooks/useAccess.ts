import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";

import {
  CreateAction,
  CreateActionRequirement,
  CreateGrant,
  CreatePermission,
  CreateRole,
  CreateRolePermission,
  CreateScope,
  CreateScopePermission,
  DeleteActionRequirement,
  DeleteGrant,
  DeletePermission,
  DeleteRole,
  DeleteRolePermission,
  DeleteScope,
  DeleteScopePermission,
  GetAction,
  GetActionByKey,
  GetActionRequirementsByPermission,
  GetGrant,
  GetGrantsBySubject,
  GetPermission,
  GetPermissionByKey,
  GetRole,
  GetRoleByKey,
  GetRolePermission,
  GetRolePermissionsByRole,
  GetScope,
  GetScopePermission,
  UpdateGrant,
  UpdatePermission,
  UpdateRole,
  UpdateScope,
} from "@/apis";
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
import { makeUnifiedQuery } from "@/hooks/core/makeUnifiedQuery";
import { accessEventEmitter, accessEvents } from "@/events/access";
import { showError, showSuccess } from "@/stores/modalStore";

import {
  ACCESS_ACTION,
  ACCESS_ACTION_REQUIREMENTS_BY_PERMISSION,
  ACCESS_ROLE,
  ACCESS_PERMISSION,
  ACCESS_SCOPE,
  ACCESS_GRANT,
  ACCESS_GRANTS_BY_SUBJECT,
  ACCESS_ROLE_PERMISSION,
  ACCESS_ROLE_PERMISSIONS_BY_ROLE,
  ACCESS_SCOPE_PERMISSION,
} from "@/constants";
import type { UnifiedQueryParams } from "./core/type";

// ========== Role 相关 ==========
export const useRoleById = (params: UnifiedQueryParams<Role> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => await GetRole(params.id),
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_ROLE(id), {id}, options);
};

export const useRoleByKey = (params: UnifiedQueryParams<Role> & { key: string }) => {
  const { key, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { key: string }) => await GetRoleByKey(params.key),
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_ROLE(key), {key}, options);
};


// 创建角色
export const useCreateRole = () => {
  return useMutation({
    mutationFn: async (params: CreateRoleRequest) => {
      return await CreateRole(params);
    },
    onSuccess: () => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_ROLES);
      showSuccess("角色创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建角色失败，请稍后重试。" + error.message);
    },
  });
};

// 更新角色
export const useUpdateRole = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateRoleRequest }) => {
      return await UpdateRole(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_ROLES);
      accessEventEmitter.emit(accessEvents.INVALIDATE_ROLE, variables.id);
      showSuccess("角色更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新角色失败，请稍后重试。" + error.message);
    },
  });
};

// 删除角色
export const useDeleteRole = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteRole(id);
    },
    onSuccess: (_, id) => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_ROLES);
      accessEventEmitter.emit(accessEvents.INVALIDATE_ROLE, id);
      showSuccess("角色删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除角色失败，请稍后重试。" + error.message);
    },
  });
};

// ========== Permission 相关 ==========

// 根据 ID 获取权限
export const usePermission = (params: UnifiedQueryParams<Permission> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetPermission(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_PERMISSION(id), { id }, options);
};

// 根据 Key 获取权限
export const usePermissionByKey = (params: UnifiedQueryParams<Permission> & { key: string }) => {
  const { key, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { key: string }) => {
      return await GetPermissionByKey(params.key);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_PERMISSION(key), { key }, options);
};

// 创建权限
export const useCreatePermission = () => {
  

  return useMutation({
    mutationFn: async (params: CreatePermissionRequest) => {
      return await CreatePermission(params);
    },
    onSuccess: () => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSIONS);
      showSuccess("权限创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建权限失败，请稍后重试。" + error.message);
    },
  });
};

// 更新权限
export const useUpdatePermission = () => {
  

  return useMutation({
    mutationFn: async (params: { id: string; data: UpdatePermissionRequest }) => {
      return await UpdatePermission(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSIONS);
      accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSION, variables.id);
      showSuccess("权限更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新权限失败，请稍后重试。" + error.message);
    },
  });
};

// 删除权限
export const useDeletePermission = () => {
  

  return useMutation({
    mutationFn: async (id: string) => {
      return await DeletePermission(id);
    },
    onSuccess: (_, id) => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSIONS);
      accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSION, id);
      showSuccess("权限删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除权限失败，请稍后重试。" + error.message);
    },
  });
};

// ========== Scope 相关 ==========

// 根据 Scope 获取作用域
export const useScope = (params: UnifiedQueryParams<Scope> & { scope: string }) => {
  const { scope, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { scope: string }) => {
      return await GetScope(params.scope);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_SCOPE(scope), { scope }, options);
};

// 创建作用域
export const useCreateScope = () => {
  

  return useMutation({
    mutationFn: async (params: CreateScopeRequest) => {
      return await CreateScope(params);
    },
    onSuccess: () => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPES);
      showSuccess("作用域创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建作用域失败，请稍后重试。" + error.message);
    },
  });
};

// 更新作用域
export const useUpdateScope = () => {
  

  return useMutation({
    mutationFn: async (params: { scope: string; data: UpdateScopeRequest }) => {
      return await UpdateScope(params.scope, params.data);
    },
    onSuccess: (_, variables) => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPES);
      accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPE, variables.scope);
      showSuccess("作用域更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新作用域失败，请稍后重试。" + error.message);
    },
  });
};

// 删除作用域
export const useDeleteScope = () => {
  

  return useMutation({
    mutationFn: async (scope: string) => {
      return await DeleteScope(scope);
    },
    onSuccess: (_, scope) => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPES);
      accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPE, scope);
      showSuccess("作用域删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除作用域失败，请稍后重试。" + error.message);
    },
  });
};

// ========== Grant 相关 ==========

// 根据 ID 获取授权
export const useGrant = (params: UnifiedQueryParams<Grant> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetGrant(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_GRANT(id), { id }, options);
};

// 根据主体获取授权列表
export const useGrantsBySubject = (params: UnifiedQueryParams<Grant[]> & {
  subject_type: string;
  subject_id: string;
  tenant_id?: string;
}) => {
  const { subject_type, subject_id, tenant_id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { subject_type: string; subject_id: string; tenant_id?: string }) => {
      return await GetGrantsBySubject(params);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(
    ACCESS_GRANTS_BY_SUBJECT(subject_type, subject_id, tenant_id),
    { subject_type, subject_id, tenant_id },
    options,
  );
};

// 创建授权
export const useCreateGrant = () => {
  

  return useMutation({
    mutationFn: async (params: CreateGrantRequest) => {
      return await CreateGrant(params);
    },
    onSuccess: () => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_GRANTS);
      showSuccess("授权创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建授权失败，请稍后重试。" + error.message);
    },
  });
};

// 更新授权
export const useUpdateGrant = () => {
  

  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateGrantRequest }) => {
      return await UpdateGrant(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_GRANTS);
      accessEventEmitter.emit(accessEvents.INVALIDATE_GRANT, variables.id);
      showSuccess("授权更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新授权失败，请稍后重试。" + error.message);
    },
  });
};

// 删除授权
export const useDeleteGrant = () => {
  

  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteGrant(id);
    },
    onSuccess: (_, id) => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_GRANTS);
      accessEventEmitter.emit(accessEvents.INVALIDATE_GRANT, id);
      showSuccess("授权删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除授权失败，请稍后重试。" + error.message);
    },
  });
};

// ========== RolePermission 相关 ==========

// 根据 ID 获取角色权限关联
export const useRolePermission = (params: UnifiedQueryParams<RolePermission> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetRolePermission(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_ROLE_PERMISSION(id), { id }, options);
};

// 根据角色ID获取角色权限列表
export const useRolePermissionsByRole = (params: UnifiedQueryParams<RolePermission[]> & { roleId: string }) => {
  const { roleId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { roleId: string }) => {
      return await GetRolePermissionsByRole(params.roleId);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_ROLE_PERMISSIONS_BY_ROLE(roleId), { roleId }, options);
};

// 创建角色权限关联
export const useCreateRolePermission = () => {
  

  return useMutation({
    mutationFn: async (params: CreateRolePermissionRequest) => {
      return await CreateRolePermission(params);
    },
    onSuccess: () => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_ROLE_PERMISSIONS);
      accessEventEmitter.emit(accessEvents.INVALIDATE_ROLES);
      accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSIONS);
      showSuccess("角色权限关联创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建角色权限关联失败，请稍后重试。" + error.message);
    },
  });
};

// 删除角色权限关联
export const useDeleteRolePermission = () => {
  

  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteRolePermission(id);
    },
    onSuccess: (_, id) => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_ROLE_PERMISSIONS);
      accessEventEmitter.emit(accessEvents.INVALIDATE_ROLE_PERMISSION, id);
      accessEventEmitter.emit(accessEvents.INVALIDATE_ROLES);
      accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSIONS);
      showSuccess("角色权限关联删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除角色权限关联失败，请稍后重试。" + error.message);
    },
  });
};

// ========== ScopePermission 相关 ==========

// 根据 ID 获取作用域权限关联
export const useScopePermission = (params: UnifiedQueryParams<ScopePermission> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetScopePermission(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_SCOPE_PERMISSION(id), { id }, options);
};

// 创建作用域权限关联
export const useCreateScopePermission = () => {
  return useMutation({
    mutationFn: async (params: CreateScopePermissionRequest) => {
      return await CreateScopePermission(params);
    },
    onSuccess: () => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPE_PERMISSIONS);
      accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPES);
      accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSIONS);
      showSuccess("作用域权限关联创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建作用域权限关联失败，请稍后重试。" + error.message);
    },
  });
};

// 删除作用域权限关联
export const useDeleteScopePermission = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteScopePermission(id);
    },
    onSuccess: (_, id) => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPE_PERMISSIONS);
      accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPE_PERMISSION, id);
      accessEventEmitter.emit(accessEvents.INVALIDATE_SCOPES);
      accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSIONS);
      showSuccess("作用域权限关联删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除作用域权限关联失败，请稍后重试。" + error.message);
    },
  });
};

// ========== Action 相关 ==========

export const useAction = (params: UnifiedQueryParams<Action> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => await GetAction(params.id),
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_ACTION(id), { id }, options);
};

export const useActionByKey = (params: UnifiedQueryParams<Action> & { key: string }) => {
  const { key, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { key: string }) => await GetActionByKey(params.key),
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_ACTION(key), { key }, options);
};

export const useCreateAction = () => {
  return useMutation({
    mutationFn: async (params: CreateActionRequest) => {
      return await CreateAction(params);
    },
    onSuccess: () => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_ACTIONS);
      showSuccess("Action 创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 Action 失败，请稍后重试。" + error.message);
    },
  });
};

// ========== ActionRequirement 相关（Permission 关联的 Action） ==========

export const useActionRequirementsByPermission = (
  params: UnifiedQueryParams<ActionRequirement[]> & { permissionId: string },
) => {
  const { permissionId, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { permissionId: string }) =>
      await GetActionRequirementsByPermission(params.permissionId),
    "suspense",
    postProcess,
  );
  return makeQuery(ACCESS_ACTION_REQUIREMENTS_BY_PERMISSION(permissionId), { permissionId }, options);
};

export const useCreateActionRequirement = () => {
  return useMutation({
    mutationFn: async (params: CreateActionRequirementRequest) => {
      return await CreateActionRequirement(params);
    },
    onSuccess: () => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_ACTION_REQUIREMENTS);
      accessEventEmitter.emit(accessEvents.INVALIDATE_ACTIONS);
      accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSIONS);
      showSuccess("Action 关联权限创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建 Action 关联权限失败，请稍后重试。" + error.message);
    },
  });
};

export const useDeleteActionRequirement = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteActionRequirement(id);
    },
    onSuccess: (_, id) => {
      accessEventEmitter.emit(accessEvents.INVALIDATE_ACTION_REQUIREMENTS);
      accessEventEmitter.emit(accessEvents.INVALIDATE_ACTION_REQUIREMENT, id);
      accessEventEmitter.emit(accessEvents.INVALIDATE_PERMISSIONS);
      showSuccess("Action 关联权限删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除 Action 关联权限失败，请稍后重试。" + error.message);
    },
  });
};
