// Profile Badge API - 基于 NFX-ID

import type {
  CreateProfileBadgeParams,
  ProfileBadge,
  UpdateProfileBadgeParams,
} from "@/types/domain";
import type { DataResponse } from "@/types/api";

import { protectedClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// 创建用户徽章关联
export const CreateProfileBadge = async (params: CreateProfileBadgeParams): Promise<ProfileBadge> => {
  const { data } = await protectedClient.post<DataResponse<ProfileBadge>>(
    URL_PATHS.AUTH.CREATE_PROFILE_BADGE,
    params,
  );
  return data.data;
};

// 根据 ID 获取用户徽章关联
export const GetProfileBadge = async (id: string): Promise<ProfileBadge> => {
  const url = URL_PATHS.AUTH.GET_PROFILE_BADGE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<ProfileBadge>>(url);
  return data.data;
};

// 根据资料 ID 获取用户徽章关联
export const GetProfileBadgesByProfileId = async (profileId: string): Promise<ProfileBadge[]> => {
  const url = URL_PATHS.AUTH.GET_PROFILE_BADGES_BY_PROFILE_ID.replace(":profile_id", profileId);
  const { data } = await protectedClient.get<DataResponse<ProfileBadge[]>>(url);
  return data.data;
};

// 根据徽章 ID 获取用户徽章关联
export const GetProfileBadgesByBadgeId = async (badgeId: string): Promise<ProfileBadge[]> => {
  const url = URL_PATHS.AUTH.GET_PROFILE_BADGES_BY_BADGE_ID.replace(":badge_id", badgeId);
  const { data } = await protectedClient.get<DataResponse<ProfileBadge[]>>(url);
  return data.data;
};

// 更新用户徽章关联
export const UpdateProfileBadge = async (
  id: string,
  params: Partial<UpdateProfileBadgeParams>,
): Promise<void> => {
  const url = URL_PATHS.AUTH.UPDATE_PROFILE_BADGE.replace(":id", id);
  await protectedClient.put<DataResponse<null>>(url, params);
};

// 删除用户徽章关联
export const DeleteProfileBadge = async (id: string): Promise<void> => {
  const url = URL_PATHS.AUTH.DELETE_PROFILE_BADGE.replace(":id", id);
  await protectedClient.delete(url);
};

// 根据资料 ID 和徽章 ID 删除用户徽章关联
export const DeleteProfileBadgeByProfileAndBadge = async (
  profileId: string,
  badgeId: string,
): Promise<void> => {
  const url = URL_PATHS.AUTH.DELETE_PROFILE_BADGE_BY_PROFILE_AND_BADGE.replace(
    ":profile_id",
    profileId,
  ).replace(":badge_id", badgeId);
  await protectedClient.delete(url);
};
