import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import { useTranslation } from "react-i18next";

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
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (params: CreateImageRequest) => {
      return await CreateImage(params);
    },
    onSuccess: () => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGES);
      showSuccess(t("image.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("image.createError") + error.message);
    },
  });
};

// 更新图片
export const useUpdateImage = () => {
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateImageRequest }) => {
      return await UpdateImage(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGES);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE, variables.id);
      showSuccess(t("image.updateSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("image.updateError") + error.message);
    },
  });
};

// 删除图片
export const useDeleteImage = () => {
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteImage(id);
    },
    onSuccess: (_, id) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGES);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE, id);
      showSuccess(t("image.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("image.deleteError") + error.message);
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
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (params: CreateImageTypeRequest) => {
      return await CreateImageType(params);
    },
    onSuccess: () => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TYPES);
      showSuccess(t("imageType.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("imageType.createError") + error.message);
    },
  });
};

// 更新图片类型
export const useUpdateImageType = () => {
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateImageTypeRequest }) => {
      return await UpdateImageType(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TYPES);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TYPE, variables.id);
      showSuccess(t("imageType.updateSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("imageType.updateError") + error.message);
    },
  });
};

// 删除图片类型
export const useDeleteImageType = () => {
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteImageType(id);
    },
    onSuccess: (_, id) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TYPES);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TYPE, id);
      showSuccess(t("imageType.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("imageType.deleteError") + error.message);
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
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (params: CreateImageVariantRequest) => {
      return await CreateImageVariant(params);
    },
    onSuccess: () => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_VARIANTS);
      showSuccess(t("imageVariant.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("imageVariant.createError") + error.message);
    },
  });
};

// 更新图片变体
export const useUpdateImageVariant = () => {
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateImageVariantRequest }) => {
      return await UpdateImageVariant(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_VARIANTS);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_VARIANT, variables.id);
      showSuccess(t("imageVariant.updateSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("imageVariant.updateError") + error.message);
    },
  });
};

// 删除图片变体
export const useDeleteImageVariant = () => {
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteImageVariant(id);
    },
    onSuccess: (_, id) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_VARIANTS);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_VARIANT, id);
      showSuccess(t("imageVariant.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("imageVariant.deleteError") + error.message);
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
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (params: CreateImageTagRequest) => {
      return await CreateImageTag(params);
    },
    onSuccess: () => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TAGS);
      showSuccess(t("imageTag.createSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("imageTag.createError") + error.message);
    },
  });
};

// 更新图片标签
export const useUpdateImageTag = () => {
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (params: { id: string; data: UpdateImageTagRequest }) => {
      return await UpdateImageTag(params.id, params.data);
    },
    onSuccess: (_, variables) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TAGS);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TAG, variables.id);
      showSuccess(t("imageTag.updateSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("imageTag.updateError") + error.message);
    },
  });
};

// 删除图片标签
export const useDeleteImageTag = () => {
  const { t } = useTranslation("hooks.image");
  return useMutation({
    mutationFn: async (id: string) => {
      return await DeleteImageTag(id);
    },
    onSuccess: (_, id) => {
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TAGS);
      imageEventEmitter.emit(imageEvents.INVALIDATE_IMAGE_TAG, id);
      showSuccess(t("imageTag.deleteSuccess"));
    },
    onError: (error: AxiosError) => {
      showError(t("imageTag.deleteError") + error.message);
    },
  });
};
