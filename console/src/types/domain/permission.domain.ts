// Permission Domain Types - 基于 NFX-ID Permission Service

export interface Permission {
  id: string;
  tag: string;
  name: string;
  description?: string;
  category?: string;
  isSystem: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface PermissionList {
  permissions: Permission[];
  total: number;
}

export interface CreatePermissionParams {
  tag: string;
  name: string;
  description?: string;
  category?: string;
  isSystem?: boolean;
}

export interface UpdatePermissionParams {
  tag: string;
  name: string;
  description?: string;
  category?: string;
}

export interface PermissionQueryParams {
  category?: string;
}
