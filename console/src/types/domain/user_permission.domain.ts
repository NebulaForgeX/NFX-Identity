// User Permission Domain Types - 基于 NFX-ID Permission Service

export interface UserPermission {
  id: string;
  userId: string;
  permissionId: string;
  tag: string;
  name: string;
  category: string;
  createdAt: string;
}

export interface AssignPermissionParams {
  userId: string;
  permissionId: string;
}

export interface RevokePermissionParams {
  userId: string;
  permissionId: string;
}

export interface CheckPermissionParams {
  userId: string;
  tag: string;
}

export interface CheckPermissionResponse {
  hasPermission: boolean; // 后端返回 has_permission，axios-case-converter 会自动转换
}
