// Image API - 基于 NFX-ID Backend

import type {
  BaseResponse,
  CreateImageRequest,
  CreateImageTagRequest,
  CreateImageTypeRequest,
  CreateImageVariantRequest,
  DataResponse,
  Image,
  ImageTag,
  ImageType,
  ImageVariant,
  UpdateImageRequest,
  UpdateImageTagRequest,
  UpdateImageTypeRequest,
  UpdateImageVariantRequest,
} from "@/types";

import { protectedClient } from "./clients";
import { URL_PATHS } from "./ip";

// ========== 图片相关 ==========

// 创建图片
export const CreateImage = async (params: CreateImageRequest): Promise<Image> => {
  const { data } = await protectedClient.post<DataResponse<Image>>(URL_PATHS.IMAGE.CREATE_IMAGE, params);
  return data.data;
};

// 根据 ID 获取图片
export const GetImage = async (id: string): Promise<Image> => {
  const url = URL_PATHS.IMAGE.GET_IMAGE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<Image>>(url);
  return data.data;
};

// 更新图片
export const UpdateImage = async (id: string, params: UpdateImageRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.IMAGE.UPDATE_IMAGE.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除图片
export const DeleteImage = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.IMAGE.DELETE_IMAGE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 图片类型相关 ==========

// 创建图片类型
export const CreateImageType = async (params: CreateImageTypeRequest): Promise<ImageType> => {
  const { data } = await protectedClient.post<DataResponse<ImageType>>(URL_PATHS.IMAGE.CREATE_IMAGE_TYPE, params);
  return data.data;
};

// 根据 ID 获取图片类型
export const GetImageType = async (id: string): Promise<ImageType> => {
  const url = URL_PATHS.IMAGE.GET_IMAGE_TYPE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<ImageType>>(url);
  return data.data;
};

// 更新图片类型
export const UpdateImageType = async (id: string, params: UpdateImageTypeRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.IMAGE.UPDATE_IMAGE_TYPE.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除图片类型
export const DeleteImageType = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.IMAGE.DELETE_IMAGE_TYPE.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 图片变体相关 ==========

// 创建图片变体
export const CreateImageVariant = async (params: CreateImageVariantRequest): Promise<ImageVariant> => {
  const { data } = await protectedClient.post<DataResponse<ImageVariant>>(URL_PATHS.IMAGE.CREATE_IMAGE_VARIANT, params);
  return data.data;
};

// 根据 ID 获取图片变体
export const GetImageVariant = async (id: string): Promise<ImageVariant> => {
  const url = URL_PATHS.IMAGE.GET_IMAGE_VARIANT.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<ImageVariant>>(url);
  return data.data;
};

// 更新图片变体
export const UpdateImageVariant = async (id: string, params: UpdateImageVariantRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.IMAGE.UPDATE_IMAGE_VARIANT.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除图片变体
export const DeleteImageVariant = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.IMAGE.DELETE_IMAGE_VARIANT.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};

// ========== 图片标签相关 ==========

// 创建图片标签
export const CreateImageTag = async (params: CreateImageTagRequest): Promise<ImageTag> => {
  const { data } = await protectedClient.post<DataResponse<ImageTag>>(URL_PATHS.IMAGE.CREATE_IMAGE_TAG, params);
  return data.data;
};

// 根据 ID 获取图片标签
export const GetImageTag = async (id: string): Promise<ImageTag> => {
  const url = URL_PATHS.IMAGE.GET_IMAGE_TAG.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<ImageTag>>(url);
  return data.data;
};

// 更新图片标签
export const UpdateImageTag = async (id: string, params: UpdateImageTagRequest): Promise<BaseResponse> => {
  const url = URL_PATHS.IMAGE.UPDATE_IMAGE_TAG.replace(":id", id);
  const { data } = await protectedClient.put<BaseResponse>(url, params);
  return data;
};

// 删除图片标签
export const DeleteImageTag = async (id: string): Promise<BaseResponse> => {
  const url = URL_PATHS.IMAGE.DELETE_IMAGE_TAG.replace(":id", id);
  const { data } = await protectedClient.delete<BaseResponse>(url);
  return data;
};
