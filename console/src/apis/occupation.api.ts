// Occupation API - 基于 NFX-ID

import type {
  CreateOccupationParams,
  Occupation,
  OccupationQueryParams,
  UpdateOccupationParams,
} from "@/types/domain";
import type { DataResponse } from "@/types/api";

import { protectedClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// 创建职业信息
export const CreateOccupation = async (params: CreateOccupationParams): Promise<Occupation> => {
  const { data } = await protectedClient.post<DataResponse<Occupation>>(
    URL_PATHS.AUTH.CREATE_OCCUPATION,
    params,
  );
  return data.data;
};

// 获取所有职业信息
export const GetOccupations = async (
  params?: OccupationQueryParams,
): Promise<{ occupations: Occupation[]; total: number }> => {
  const response = await protectedClient.get<DataResponse<Occupation[]>>(
    URL_PATHS.AUTH.GET_OCCUPATIONS,
    {
      params: params,
    },
  );
  const total = (response.data.meta?.total as number) || response.data.data.length;
  return {
    occupations: response.data.data,
    total,
  };
};

// 根据 ID 获取职业信息
export const GetOccupation = async (id: string): Promise<Occupation> => {
  const url = URL_PATHS.AUTH.GET_OCCUPATION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Occupation>>(url);
  return data.data;
};

// 根据资料 ID 获取职业信息
export const GetOccupationsByProfileId = async (profileId: string): Promise<Occupation[]> => {
  const url = URL_PATHS.AUTH.GET_OCCUPATIONS_BY_PROFILE_ID.replace(":profile_id", profileId);
  const { data } = await protectedClient.get<DataResponse<Occupation[]>>(url);
  return data.data;
};

// 更新职业信息
export const UpdateOccupation = async (
  id: string,
  params: Partial<UpdateOccupationParams>,
): Promise<void> => {
  const url = URL_PATHS.AUTH.UPDATE_OCCUPATION.replace(":id", id);
  await protectedClient.put<DataResponse<null>>(url, params);
};

// 删除职业信息
export const DeleteOccupation = async (id: string): Promise<void> => {
  const url = URL_PATHS.AUTH.DELETE_OCCUPATION.replace(":id", id);
  await protectedClient.delete(url);
};
