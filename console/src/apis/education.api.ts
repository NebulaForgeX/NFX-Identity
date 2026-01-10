// Education API - 基于 NFX-ID

import type {
  CreateEducationParams,
  Education,
  EducationQueryParams,
  UpdateEducationParams,
} from "@/types/domain";
import type { DataResponse } from "@/types/api";

import { protectedClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// 创建教育经历
export const CreateEducation = async (params: CreateEducationParams): Promise<Education> => {
  const { data } = await protectedClient.post<DataResponse<Education>>(
    URL_PATHS.AUTH.CREATE_EDUCATION,
    params,
  );
  return data.data;
};

// 获取所有教育经历
export const GetEducations = async (
  params?: EducationQueryParams,
): Promise<{ educations: Education[]; total: number }> => {
  const response = await protectedClient.get<DataResponse<Education[]>>(
    URL_PATHS.AUTH.GET_EDUCATIONS,
    {
      params: params,
    },
  );
  const total = (response.data.meta?.total as number) || response.data.data.length;
  return {
    educations: response.data.data,
    total,
  };
};

// 根据 ID 获取教育经历
export const GetEducation = async (id: string): Promise<Education> => {
  const url = URL_PATHS.AUTH.GET_EDUCATION.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Education>>(url);
  return data.data;
};

// 根据资料 ID 获取教育经历
export const GetEducationsByProfileId = async (profileId: string): Promise<Education[]> => {
  const url = URL_PATHS.AUTH.GET_EDUCATIONS_BY_PROFILE_ID.replace(":profile_id", profileId);
  const { data } = await protectedClient.get<DataResponse<Education[]>>(url);
  return data.data;
};

// 更新教育经历
export const UpdateEducation = async (
  id: string,
  params: Partial<UpdateEducationParams>,
): Promise<void> => {
  const url = URL_PATHS.AUTH.UPDATE_EDUCATION.replace(":id", id);
  await protectedClient.put<DataResponse<null>>(url, params);
};

// 删除教育经历
export const DeleteEducation = async (id: string): Promise<void> => {
  const url = URL_PATHS.AUTH.DELETE_EDUCATION.replace(":id", id);
  await protectedClient.delete(url);
};
