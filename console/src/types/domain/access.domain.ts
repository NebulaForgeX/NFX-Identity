// Access Domain Types - 基于 NFX-ID Backend
// 注意：由于使用了 axios-case-converter，后端返回的 snake_case 会自动转换为 camelCase

import { ScopeType, SubjectType, GrantType, GrantEffect } from "./enums";

export interface Role {
  id: string;
  key: string;
  name: string;
  description?: string;
  scopeType: ScopeType;
  isSystem: boolean;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface Permission {
  id: string;
  key: string;
  name: string;
  description?: string;
  isSystem: boolean;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface Scope {
  scope: string;
  description?: string;
  isSystem: boolean;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface Grant {
  id: string;
  subjectType: SubjectType;
  subjectId: string;
  grantType: GrantType;
  grantRefId: string;
  tenantId?: string;
  appId?: string;
  resourceType?: string;
  resourceId?: string;
  effect: GrantEffect;
  expiresAt?: string;
  createdAt: string;
  createdBy?: string;
  revokedAt?: string;
  revokedBy?: string;
  revokeReason?: string;
}

export interface RolePermission {
  id: string;
  roleId: string;
  permissionId: string;
  createdAt: string;
}

export interface ScopePermission {
  id: string;
  scope: string;
  permissionId: string;
  createdAt: string;
}
