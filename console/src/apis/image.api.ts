// Image API - 基于 NFX-ID

import type {
  CreateImageParams,
  CreateImageTypeParams,
  Image,
  ImageQueryParams,
  ImageType,
  ImageTypeQueryParams,
  UpdateImageParams,
  UpdateImageTypeParams,
} from "@/types/domain";
import type { DataResponse } from "@/types/api";

import { protectedClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// ========== 图片相关 ==========

// 创建图片
export const CreateImage = async (params: CreateImageParams): Promise<Image> => {
  const { data } = await protectedClient.post<DataResponse<Image>>(
    URL_PATHS.IMAGE.CREATE_IMAGE,
    params,
  );
  return data.data;
};

// 获取所有图片
export const GetImages = async (
  params?: ImageQueryParams,
): Promise<{ images: Image[]; total: number }> => {
  const response = await protectedClient.get<DataResponse<Image[]>>(URL_PATHS.IMAGE.GET_IMAGES, {
    params: params,
  });
  const total = (response.data.meta?.total as number) || response.data.data.length;
  return {
    images: response.data.data,
    total,
  };
};

// 根据 ID 获取图片
export const GetImage = async (id: string): Promise<Image> => {
  const url = URL_PATHS.IMAGE.GET_IMAGE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Image>>(url);
  return data.data;
};

// 更新图片
export const UpdateImage = async (id: string, params: Partial<UpdateImageParams>): Promise<void> => {
  const url = URL_PATHS.IMAGE.UPDATE_IMAGE.replace(":id", id);
  await protectedClient.put<DataResponse<null>>(url, params);
};

// 删除图片
export const DeleteImage = async (id: string): Promise<void> => {
  const url = URL_PATHS.IMAGE.DELETE_IMAGE.replace(":id", id);
  await protectedClient.delete(url);
};

// ========== 图片类型相关 ==========

// 创建图片类型
export const CreateImageType = async (params: CreateImageTypeParams): Promise<ImageType> => {
  const { data } = await protectedClient.post<DataResponse<ImageType>>(
    URL_PATHS.IMAGE.CREATE_IMAGE_TYPE,
    params,
  );
  return data.data;
};

// 获取所有图片类型
export const GetImageTypes = async (
  params?: ImageTypeQueryParams,
): Promise<{ imageTypes: ImageType[]; total: number }> => {
  const response = await protectedClient.get<DataResponse<ImageType[]>>(
    URL_PATHS.IMAGE.GET_IMAGE_TYPES,
    {
      params: params,
    },
  );
  const total = (response.data.meta?.total as number) || response.data.data.length;
  return {
    imageTypes: response.data.data,
    total,
  };
};

// 根据 ID 获取图片类型
export const GetImageType = async (id: string): Promise<ImageType> => {
  const url = URL_PATHS.IMAGE.GET_IMAGE_TYPE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<ImageType>>(url);
  return data.data;
};

// 根据 key 获取图片类型
export const GetImageTypeByKey = async (key: string): Promise<ImageType> => {
  const url = URL_PATHS.IMAGE.GET_IMAGE_TYPE_BY_KEY.replace(":key", key);
  const { data } = await protectedClient.get<DataResponse<ImageType>>(url);
  return data.data;
};

// 更新图片类型
export const UpdateImageType = async (
  id: string,
  params: Partial<UpdateImageTypeParams>,
): Promise<void> => {
  const url = URL_PATHS.IMAGE.UPDATE_IMAGE_TYPE.replace(":id", id);
  await protectedClient.put<DataResponse<null>>(url, params);
};

// 删除图片类型
export const DeleteImageType = async (id: string): Promise<void> => {
  const url = URL_PATHS.IMAGE.DELETE_IMAGE_TYPE.replace(":id", id);
  await protectedClient.delete(url);
};
