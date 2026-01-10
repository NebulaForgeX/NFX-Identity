// Profile API - 基于 NFX-ID

import type {
  CreateProfileParams,
  Profile,
  ProfileQueryParams,
  UpdateProfileParams,
} from "@/types/domain";
import type { DataResponse } from "@/types/api";

import { protectedClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// 创建资料
export const CreateProfile = async (params: CreateProfileParams): Promise<Profile> => {
  const { data } = await protectedClient.post<DataResponse<Profile>>(
    URL_PATHS.AUTH.CREATE_PROFILE,
    params,
  );
  return data.data;
};

// 获取所有资料
export const GetProfiles = async (
  params?: ProfileQueryParams,
): Promise<{ profiles: Profile[]; total: number }> => {
  const response = await protectedClient.get<DataResponse<Profile[]>>(URL_PATHS.AUTH.GET_PROFILES, {
    params: params,
  });
  const total = (response.data.meta?.total as number) || response.data.data.length;
  return {
    profiles: response.data.data,
    total,
  };
};

// 根据 ID 获取资料
export const GetProfile = async (id: string): Promise<Profile> => {
  const url = URL_PATHS.AUTH.GET_PROFILE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Profile>>(url);
  return data.data;
};

// 根据用户 ID 获取资料
export const GetProfileByUserId = async (userId: string): Promise<Profile> => {
  const url = URL_PATHS.AUTH.GET_PROFILE_BY_USER_ID.replace(":user_id", userId);
  const { data } = await protectedClient.get<DataResponse<Profile>>(url);
  return data.data;
};

// 更新资料
export const UpdateProfile = async (id: string, params: Partial<UpdateProfileParams>): Promise<void> => {
  const url = URL_PATHS.AUTH.UPDATE_PROFILE.replace(":id", id);
  await protectedClient.put<DataResponse<null>>(url, params);
};

// 删除资料
export const DeleteProfile = async (id: string): Promise<void> => {
  const url = URL_PATHS.AUTH.DELETE_PROFILE.replace(":id", id);
  await protectedClient.delete(url);
};
