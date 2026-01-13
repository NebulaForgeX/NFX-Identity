export const imageEvents = {
  // Image 相关
  INVALIDATE_IMAGE: "IMAGE:INVALIDATE_IMAGE",
  INVALIDATE_IMAGES: "IMAGE:INVALIDATE_IMAGES",

  // ImageType 相关
  INVALIDATE_IMAGE_TYPE: "IMAGE:INVALIDATE_IMAGE_TYPE",
  INVALIDATE_IMAGE_TYPES: "IMAGE:INVALIDATE_IMAGE_TYPES",

  // ImageVariant 相关
  INVALIDATE_IMAGE_VARIANT: "IMAGE:INVALIDATE_IMAGE_VARIANT",
  INVALIDATE_IMAGE_VARIANTS: "IMAGE:INVALIDATE_IMAGE_VARIANTS",

  // ImageTag 相关
  INVALIDATE_IMAGE_TAG: "IMAGE:INVALIDATE_IMAGE_TAG",
  INVALIDATE_IMAGE_TAGS: "IMAGE:INVALIDATE_IMAGE_TAGS",
} as const;

type ImageEvent = (typeof imageEvents)[keyof typeof imageEvents];

class ImageEventEmitter {
  private listeners: Record<ImageEvent, Set<Function>> = {
    [imageEvents.INVALIDATE_IMAGE]: new Set<Function>(),
    [imageEvents.INVALIDATE_IMAGES]: new Set<Function>(),
    [imageEvents.INVALIDATE_IMAGE_TYPE]: new Set<Function>(),
    [imageEvents.INVALIDATE_IMAGE_TYPES]: new Set<Function>(),
    [imageEvents.INVALIDATE_IMAGE_VARIANT]: new Set<Function>(),
    [imageEvents.INVALIDATE_IMAGE_VARIANTS]: new Set<Function>(),
    [imageEvents.INVALIDATE_IMAGE_TAG]: new Set<Function>(),
    [imageEvents.INVALIDATE_IMAGE_TAGS]: new Set<Function>(),
  };

  on(event: ImageEvent, callback: Function) {
    this.listeners[event].add(callback);
  }

  off(event: ImageEvent, callback: Function) {
    this.listeners[event].delete(callback);
  }

  emit(event: ImageEvent, ...args: unknown[]) {
    this.listeners[event].forEach((callback) => {
      callback(...args);
    });
  }
}

export const imageEventEmitter = new ImageEventEmitter();
