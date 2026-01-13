import type { QueryClient } from "@tanstack/react-query";
import { imageEventEmitter, imageEvents } from "@/events/image";
import { IMAGE_QUERY_KEY_PREFIXES } from "@/constants";

/**
 * Image 相关的缓存失效事件处理
 */
export const useImageCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateImages = () => queryClient.invalidateQueries({ queryKey: IMAGE_QUERY_KEY_PREFIXES.IMAGES });
  const handleInvalidateImage = (item: string) => queryClient.invalidateQueries({ queryKey: [...IMAGE_QUERY_KEY_PREFIXES.IMAGE, item] });
  const handleInvalidateImageTypes = () => queryClient.invalidateQueries({ queryKey: IMAGE_QUERY_KEY_PREFIXES.IMAGE_TYPES });
  const handleInvalidateImageType = (item: string) => queryClient.invalidateQueries({ queryKey: [...IMAGE_QUERY_KEY_PREFIXES.IMAGE_TYPE, item] });
  const handleInvalidateImageVariants = () => queryClient.invalidateQueries({ queryKey: IMAGE_QUERY_KEY_PREFIXES.IMAGE_VARIANTS });
  const handleInvalidateImageVariant = (item: string) => queryClient.invalidateQueries({ queryKey: [...IMAGE_QUERY_KEY_PREFIXES.IMAGE_VARIANT, item] });
  const handleInvalidateImageTags = () => queryClient.invalidateQueries({ queryKey: IMAGE_QUERY_KEY_PREFIXES.IMAGE_TAGS });
  const handleInvalidateImageTag = (item: string) => queryClient.invalidateQueries({ queryKey: [...IMAGE_QUERY_KEY_PREFIXES.IMAGE_TAG, item] });

  // 注册监听器
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGES, handleInvalidateImages);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE, handleInvalidateImage);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TYPES, handleInvalidateImageTypes);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TYPE, handleInvalidateImageType);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_VARIANTS, handleInvalidateImageVariants);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_VARIANT, handleInvalidateImageVariant);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TAGS, handleInvalidateImageTags);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TAG, handleInvalidateImageTag);

  // 清理监听器
  return () => {
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGES, handleInvalidateImages);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE, handleInvalidateImage);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TYPES, handleInvalidateImageTypes);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TYPE, handleInvalidateImageType);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_VARIANTS, handleInvalidateImageVariants);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_VARIANT, handleInvalidateImageVariant);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TAGS, handleInvalidateImageTags);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TAG, handleInvalidateImageTag);
  };
};
