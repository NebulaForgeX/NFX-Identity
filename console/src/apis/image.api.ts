// Image API - 基于 NFX-ID Backend

import type { BaseResponse, DataResponse } from "@/types/api";

import { protectedClient } from "@/apis/clients";
import { URL_PATHS } from "@/apis/ip";

// ========== 图片相关 ==========

// 创建图片
export const CreateImage = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.IMAGE.CREATE_IMAGE, params);
  return data.data;
};

// 根据 ID 获取图片
export const GetImage = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.IMAGE.GET_IMAGE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新图片
export const UpdateImage = async (id: string, params: unknown): Promise<BaseResponse> => {
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
export const CreateImageType = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.IMAGE.CREATE_IMAGE_TYPE, params);
  return data.data;
};

// 根据 ID 获取图片类型
export const GetImageType = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.IMAGE.GET_IMAGE_TYPE.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新图片类型
export const UpdateImageType = async (id: string, params: unknown): Promise<BaseResponse> => {
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
export const CreateImageVariant = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.IMAGE.CREATE_IMAGE_VARIANT, params);
  return data.data;
};

// 根据 ID 获取图片变体
export const GetImageVariant = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.IMAGE.GET_IMAGE_VARIANT.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新图片变体
export const UpdateImageVariant = async (id: string, params: unknown): Promise<BaseResponse> => {
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
export const CreateImageTag = async (params: unknown): Promise<unknown> => {
  const { data } = await protectedClient.post<DataResponse<unknown>>(URL_PATHS.IMAGE.CREATE_IMAGE_TAG, params);
  return data.data;
};

// 根据 ID 获取图片标签
export const GetImageTag = async (id: string): Promise<unknown> => {
  const url = URL_PATHS.IMAGE.GET_IMAGE_TAG.replace(":id", id);
  const { data } = await protectedClient.get<DataResponse<unknown>>(url);
  return data.data;
};

// 更新图片标签
export const UpdateImageTag = async (id: string, params: unknown): Promise<BaseResponse> => {
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
