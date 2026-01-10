// Access Request Types - 基于 NFX-ID Backend
// 注意：由于使用了 axios-case-converter，前端使用 camelCase，会自动转换为后端的 snake_case

// ========== 角色相关 ==========

export interface CreateRoleRequest {
  key: string;
  name: string;
  description?: string;
  scopeType?: string;
  isSystem?: boolean;
}

export interface UpdateRoleRequest {
  name: string;
  description?: string;
  scopeType?: string;
}

// ========== 权限相关 ==========

export interface CreatePermissionRequest {
  key: string;
  name: string;
  description?: string;
  isSystem?: boolean;
}

export interface UpdatePermissionRequest {
  name: string;
  description?: string;
}

// ========== 作用域相关 ==========

export interface CreateScopeRequest {
  scope: string;
  description?: string;
  isSystem?: boolean;
}

export interface UpdateScopeRequest {
  description?: string;
}

// ========== 授权相关 ==========

export interface CreateGrantRequest {
  subjectType: string;
  subjectId: string;
  grantType: string;
  grantRefId: string;
  tenantId?: string;
  appId?: string;
  resourceType?: string;
  resourceId?: string;
  effect: string;
  expiresAt?: string;
}

export interface UpdateGrantRequest {
  expiresAt?: string;
}

export interface RevokeGrantRequest {
  revokedBy: string;
  revokeReason?: string;
}

// ========== 角色权限关联相关 ==========

export interface CreateRolePermissionRequest {
  roleId: string;
  permissionId: string;
}

// ========== 作用域权限关联相关 ==========

export interface CreateScopePermissionRequest {
  scope: string;
  permissionId: string;
}
