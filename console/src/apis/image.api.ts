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

// ========== 图片上传相关 ==========

// 上传图片（上传到 tmp 目录，后续通过 Directory 服务创建时自动移动）
export const UploadImage = async (file: File, isPublic: boolean = true): Promise<Image> => {
  const formData = new FormData();
  formData.append("file", file);

  const url = `${URL_PATHS.IMAGE.upload}?is_public=${isPublic}`;
  const { data } = await protectedClient.post<DataResponse<Image>>(url, formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return data.data;
};

// ========== 图片相关 ==========

// 创建图片
export const CreateImage = async (params: CreateImageRequest): Promise<Image> => {
  const { data } = await protectedClient.post<DataResponse<Image>>(
    URL_PATHS.IMAGE.images,
    params,
  );
  return data.data;
};

// 根据 ID 获取图片
export const GetImage = async (id: string): Promise<Image> => {
  const { data } = await protectedClient.get<DataResponse<Image>>(
    URL_PATHS.IMAGE.images.byId(id),
  );
  return data.data;
};

// 更新图片
export const UpdateImage = async (id: string, params: UpdateImageRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.IMAGE.images.byId(id),
    params,
  );
  return data;
};

// 删除图片
export const DeleteImage = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.IMAGE.images.byId(id),
  );
  return data;
};

// ========== 图片类型相关 ==========

// 创建图片类型
export const CreateImageType = async (params: CreateImageTypeRequest): Promise<ImageType> => {
  const { data } = await protectedClient.post<DataResponse<ImageType>>(
    URL_PATHS.IMAGE.imageTypes,
    params,
  );
  return data.data;
};

// 根据 ID 获取图片类型
export const GetImageType = async (id: string): Promise<ImageType> => {
  const { data } = await protectedClient.get<DataResponse<ImageType>>(
    URL_PATHS.IMAGE.imageTypes.byId(id),
  );
  return data.data;
};

// 更新图片类型
export const UpdateImageType = async (id: string, params: UpdateImageTypeRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.IMAGE.imageTypes.byId(id),
    params,
  );
  return data;
};

// 删除图片类型
export const DeleteImageType = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.IMAGE.imageTypes.byId(id),
  );
  return data;
};

// ========== 图片变体相关 ==========

// 创建图片变体
export const CreateImageVariant = async (params: CreateImageVariantRequest): Promise<ImageVariant> => {
  const { data } = await protectedClient.post<DataResponse<ImageVariant>>(
    URL_PATHS.IMAGE.imageVariants,
    params,
  );
  return data.data;
};

// 根据 ID 获取图片变体
export const GetImageVariant = async (id: string): Promise<ImageVariant> => {
  const { data } = await protectedClient.get<DataResponse<ImageVariant>>(
    URL_PATHS.IMAGE.imageVariants.byId(id),
  );
  return data.data;
};

// 更新图片变体
export const UpdateImageVariant = async (id: string, params: UpdateImageVariantRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.IMAGE.imageVariants.byId(id),
    params,
  );
  return data;
};

// 删除图片变体
export const DeleteImageVariant = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.IMAGE.imageVariants.byId(id),
  );
  return data;
};

// ========== 图片标签相关 ==========

// 创建图片标签
export const CreateImageTag = async (params: CreateImageTagRequest): Promise<ImageTag> => {
  const { data } = await protectedClient.post<DataResponse<ImageTag>>(
    URL_PATHS.IMAGE.imageTags,
    params,
  );
  return data.data;
};

// 根据 ID 获取图片标签
export const GetImageTag = async (id: string): Promise<ImageTag> => {
  const { data } = await protectedClient.get<DataResponse<ImageTag>>(
    URL_PATHS.IMAGE.imageTags.byId(id),
  );
  return data.data;
};

// 更新图片标签
export const UpdateImageTag = async (id: string, params: UpdateImageTagRequest): Promise<BaseResponse> => {
  const { data } = await protectedClient.put<BaseResponse>(
    URL_PATHS.IMAGE.imageTags.byId(id),
    params,
  );
  return data;
};

// 删除图片标签
export const DeleteImageTag = async (id: string): Promise<BaseResponse> => {
  const { data } = await protectedClient.delete<BaseResponse>(
    URL_PATHS.IMAGE.imageTags.byId(id),
  );
  return data;
};
