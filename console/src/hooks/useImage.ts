import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";

import {
  CreateImage,
  CreateImageTag,
  CreateImageType,
  CreateImageVariant,
  DeleteImage,
  DeleteImageTag,
  DeleteImageType,
  DeleteImageVariant,
  GetImage,
  GetImageTag,
  GetImageType,
  GetImageVariant,
  UpdateImage,
  UpdateImageTag,
  UpdateImageType,
  UpdateImageVariant,
} from "@/apis/image.api";
import type {
  CreateImageRequest,
  CreateImageTagRequest,
  CreateImageTypeRequest,
  CreateImageVariantRequest,
  Image,
  ImageTag,
  ImageType,
  ImageVariant,
  UpdateImageRequest,
  UpdateImageTagRequest,
  UpdateImageTypeRequest,
  UpdateImageVariantRequest,
} from "@/types";
import { makeUnifiedQuery } from "@/hooks/core/makeUnifiedQuery";
import { imageEventEmitter, imageEvents } from "@/events/image";
import { showError, showSuccess } from "@/stores/modalStore";
import { IMAGE_IMAGE, IMAGE_IMAGE_TYPE, IMAGE_IMAGE_VARIANT, IMAGE_IMAGE_TAG } from "@/constants";
import type { UnifiedQueryParams } from "./core/type";

// ========== Image 相关 ==========

// 根据 ID 获取图片
export const useImage = (params: UnifiedQueryParams<Image> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetImage(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(IMAGE_IMAGE(id), { id }, options);
};

// 创建图片
export const useCreateImage = () => {
  return useMutation({
    mutationFn: async (params: CreateImageRequest) => {
      return await CreateImage(params);
    },
    onSuccess: () => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGES);
      showSuccess("图片创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建图片失败，请稍后重试。" + error.message);
    },
  });
};

// 更新图片
export const useUpdateImage = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateImageRequest }) => {
      return await UpdateImage(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGES);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE, variables.id);
      showSuccess("图片更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新图片失败，请稍后重试。" + error.message);
    },
  });
};

// 删除图片
export const useDeleteImage = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteImage(id);
    },
    onSuccess: (_, id) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGES);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE, id);
      showSuccess("图片删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除图片失败，请稍后重试。" + error.message);
    },
  });
};

// ========== ImageType 相关 ==========

// 根据 ID 获取图片类型
export const useImageType = (params: UnifiedQueryParams<ImageType> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetImageType(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(IMAGE_IMAGE_TYPE(id), { id }, options);
};

// 创建图片类型
export const useCreateImageType = () => {
  return useMutation({
    mutationFn: async (params: CreateImageTypeRequest) => {
      return await CreateImageType(params);
    },
    onSuccess: () => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TYPES);
      showSuccess("图片类型创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建图片类型失败，请稍后重试。" + error.message);
    },
  });
};

// 更新图片类型
export const useUpdateImageType = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateImageTypeRequest }) => {
      return await UpdateImageType(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TYPES);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TYPE, variables.id);
      showSuccess("图片类型更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新图片类型失败，请稍后重试。" + error.message);
    },
  });
};

// 删除图片类型
export const useDeleteImageType = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteImageType(id);
    },
    onSuccess: (_, id) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TYPES);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TYPE, id);
      showSuccess("图片类型删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除图片类型失败，请稍后重试。" + error.message);
    },
  });
};

// ========== ImageVariant 相关 ==========

// 根据 ID 获取图片变体
export const useImageVariant = (params: UnifiedQueryParams<ImageVariant> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetImageVariant(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(IMAGE_IMAGE_VARIANT(id), { id }, options);
};

// 创建图片变体
export const useCreateImageVariant = () => {
  return useMutation({
    mutationFn: async (params: CreateImageVariantRequest) => {
      return await CreateImageVariant(params);
    },
    onSuccess: () => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_VARIANTS);
      showSuccess("图片变体创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建图片变体失败，请稍后重试。" + error.message);
    },
  });
};

// 更新图片变体
export const useUpdateImageVariant = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateImageVariantRequest }) => {
      return await UpdateImageVariant(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_VARIANTS);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_VARIANT, variables.id);
      showSuccess("图片变体更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新图片变体失败，请稍后重试。" + error.message);
    },
  });
};

// 删除图片变体
export const useDeleteImageVariant = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteImageVariant(id);
    },
    onSuccess: (_, id) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_VARIANTS);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_VARIANT, id);
      showSuccess("图片变体删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除图片变体失败，请稍后重试。" + error.message);
    },
  });
};

// ========== ImageTag 相关 ==========

// 根据 ID 获取图片标签
export const useImageTag = (params: UnifiedQueryParams<ImageTag> & { id: string }) => {
  const { id, options, postProcess } = params;
  const makeQuery = makeUnifiedQuery(
    async (params: { id: string }) => {
      return await GetImageTag(params.id);
    },
    "suspense",
    postProcess,
  );
  return makeQuery(IMAGE_IMAGE_TAG(id), { id }, options);
};

// 创建图片标签
export const useCreateImageTag = () => {
  return useMutation({
    mutationFn: async (params: CreateImageTagRequest) => {
      return await CreateImageTag(params);
    },
    onSuccess: () => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TAGS);
      showSuccess("图片标签创建成功！");
    },
    onError: (error: AxiosError) => {
      showError("创建图片标签失败，请稍后重试。" + error.message);
    },
  });
};

// 更新图片标签
export const useUpdateImageTag = () => {
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateImageTagRequest }) => {
      return await UpdateImageTag(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TAGS);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TAG, variables.id);
      showSuccess("图片标签更新成功！");
    },
    onError: (error: AxiosError) => {
      showError("更新图片标签失败，请稍后重试。" + error.message);
    },
  });
};

// 删除图片标签
export const useDeleteImageTag = () => {
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteImageTag(id);
    },
    onSuccess: (_, id) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TAGS);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TAG, id);
      showSuccess("图片标签删除成功！");
    },
    onError: (error: AxiosError) => {
      showError("删除图片标签失败，请稍后重试。" + error.message);
    },
  });
};
