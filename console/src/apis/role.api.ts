// Role API - 基于 NFX-ID

import type { CreateRoleParams, Role, RoleQueryParams, UpdateRoleParams } from "@/types/domain";
import type { DataResponse } from "@/types/api";

import { protectedClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// 创建角色
export const CreateRole = async (params: CreateRoleParams): Promise<Role> => {
  const { data } = await protectedClient.post<DataResponse<Role>>(URL_PATHS.AUTH.CREATE_ROLE, params);
  return data.data;
};

// 获取所有角色
export const GetRoles = async (params?: RoleQueryParams): Promise<{ roles: Role[]; total: number }> => {
  const response = await protectedClient.get<DataResponse<Role[]>>(URL_PATHS.AUTH.GET_ROLES, {
    params: params,
  });
  const total = (response.data.meta?.total as number) || response.data.data.length;
  return {
    roles: response.data.data,
    total,
  };
};

// 根据 ID 获取角色
export const GetRole = async (id: string): Promise<Role> => {
  const url = URL_PATHS.AUTH.GET_ROLE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Role>>(url);
  return data.data;
};

// 根据名称获取角色
export const GetRoleByName = async (name: string): Promise<Role> => {
  const url = URL_PATHS.AUTH.GET_ROLE_BY_NAME.replace(":name", name);
  const { data } = await protectedClient.get<DataResponse<Role>>(url);
  return data.data;
};

// 更新角色
export const UpdateRole = async (id: string, params: Partial<UpdateRoleParams>): Promise<void> => {
  const url = URL_PATHS.AUTH.UPDATE_ROLE.replace(":id", id);
  await protectedClient.put<DataResponse<null>>(url, params);
};

// 删除角色
export const DeleteRole = async (id: string): Promise<void> => {
  const url = URL_PATHS.AUTH.DELETE_ROLE.replace(":id", id);
  await protectedClient.delete(url);
};
