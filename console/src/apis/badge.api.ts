// Badge API - 基于 NFX-ID

import type { Badge, BadgeQueryParams, CreateBadgeParams, UpdateBadgeParams } from "@/types/domain";
import type { DataResponse } from "@/types/api";

import { protectedClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// 创建徽章
export const CreateBadge = async (params: CreateBadgeParams): Promise<Badge> => {
  const { data } = await protectedClient.post<DataResponse<Badge>>(URL_PATHS.AUTH.CREATE_BADGE, params);
  return data.data;
};

// 获取所有徽章
export const GetBadges = async (params?: BadgeQueryParams): Promise<{ badges: Badge[]; total: number }> => {
  const response = await protectedClient.get<DataResponse<Badge[]>>(URL_PATHS.AUTH.GET_BADGES, {
    params: params,
  });
  const total = (response.data.meta?.total as number) || response.data.data.length;
  return {
    badges: response.data.data,
    total,
  };
};

// 根据 ID 获取徽章
export const GetBadge = async (id: string): Promise<Badge> => {
  const url = URL_PATHS.AUTH.GET_BADGE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Badge>>(url);
  return data.data;
};

// 根据名称获取徽章
export const GetBadgeByName = async (name: string): Promise<Badge> => {
  const url = URL_PATHS.AUTH.GET_BADGE_BY_NAME.replace(":name", name);
  const { data } = await protectedClient.get<DataResponse<Badge>>(url);
  return data.data;
};

// 更新徽章
export const UpdateBadge = async (id: string, params: Partial<UpdateBadgeParams>): Promise<void> => {
  const url = URL_PATHS.AUTH.UPDATE_BADGE.replace(":id", id);
  await protectedClient.put<DataResponse<null>>(url, params);
};

// 删除徽章
export const DeleteBadge = async (id: string): Promise<void> => {
  const url = URL_PATHS.AUTH.DELETE_BADGE.replace(":id", id);
  await protectedClient.delete(url);
};
