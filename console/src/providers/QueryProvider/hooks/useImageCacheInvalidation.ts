import type { QueryClient } from "@tanstack/react-query";
import { imageEventEmitter, imageEvents } from "@/events/image";
import {
  IMAGE_IMAGE,
  IMAGE_IMAGE_LIST,
  IMAGE_IMAGE_TAG,
  IMAGE_IMAGE_TAG_LIST,
  IMAGE_IMAGE_TYPE,
  IMAGE_IMAGE_TYPE_LIST,
  IMAGE_IMAGE_VARIANT,
  IMAGE_IMAGE_VARIANT_LIST,
} from "@/constants";

type EventCb = (...args: unknown[]) => void;

/**
 * Image 相关的缓存失效事件处理
 */
export const useImageCacheInvalidation = (queryClient: QueryClient) => {
  const handleInvalidateImages = () => queryClient.invalidateQueries({ queryKey: IMAGE_IMAGE_LIST });
  const handleInvalidateImage = (item: string) => queryClient.invalidateQueries({ queryKey: IMAGE_IMAGE(item) });
  const handleInvalidateImageTypes = () => queryClient.invalidateQueries({ queryKey: IMAGE_IMAGE_TYPE_LIST });
  const handleInvalidateImageType = (item: string) => queryClient.invalidateQueries({ queryKey: IMAGE_IMAGE_TYPE(item) });
  const handleInvalidateImageVariants = () => queryClient.invalidateQueries({ queryKey: IMAGE_IMAGE_VARIANT_LIST });
  const handleInvalidateImageVariant = (item: string) => queryClient.invalidateQueries({ queryKey: IMAGE_IMAGE_VARIANT(item) });
  const handleInvalidateImageTags = () => queryClient.invalidateQueries({ queryKey: IMAGE_IMAGE_TAG_LIST });
  const handleInvalidateImageTag = (item: string) => queryClient.invalidateQueries({ queryKey: IMAGE_IMAGE_TAG(item) });

  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGES, handleInvalidateImages as EventCb);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE, handleInvalidateImage as EventCb);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TYPES, handleInvalidateImageTypes as EventCb);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TYPE, handleInvalidateImageType as EventCb);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_VARIANTS, handleInvalidateImageVariants as EventCb);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_VARIANT, handleInvalidateImageVariant as EventCb);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TAGS, handleInvalidateImageTags as EventCb);
  imageEventEmitter.on(imageEvents.INVALIDATE_IMAGE_TAG, handleInvalidateImageTag as EventCb);

  return () => {
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGES, handleInvalidateImages as EventCb);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE, handleInvalidateImage as EventCb);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TYPES, handleInvalidateImageTypes as EventCb);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TYPE, handleInvalidateImageType as EventCb);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_VARIANTS, handleInvalidateImageVariants as EventCb);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_VARIANT, handleInvalidateImageVariant as EventCb);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TAGS, handleInvalidateImageTags as EventCb);
    imageEventEmitter.off(imageEvents.INVALIDATE_IMAGE_TAG, handleInvalidateImageTag as EventCb);
  };
};
