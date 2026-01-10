// Role Domain Types - 基于 NFX-ID

export interface Role {
  id: string;
  name: string;
  description?: string;
  permissions: string[];
  isSystem: boolean;
  createdAt: string;
  updatedAt: string;
  deletedAt?: string;
}

export interface CreateRoleParams {
  name: string;
  description?: string;
  permissions?: string[];
  isSystem?: boolean;
}

export interface UpdateRoleParams {
  name?: string;
  description?: string;
  permissions?: string[];
  isSystem?: boolean;
}

export interface RoleQueryParams {
  offset?: number;
  limit?: number;
  search?: string;
  isSystem?: boolean;
  sort?: string[];
}
