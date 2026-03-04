import { EventEmitter, defineEvents, type EventNamesOf } from "nfx-ui/events";
import { singleton } from "nfx-ui/utils";

export const imageEvents = defineEvents({
  INVALIDATE_IMAGE: "IMAGE:INVALIDATE_IMAGE",
  INVALIDATE_IMAGES: "IMAGE:INVALIDATE_IMAGES",
  INVALIDATE_IMAGE_TYPE: "IMAGE:INVALIDATE_IMAGE_TYPE",
  INVALIDATE_IMAGE_TYPES: "IMAGE:INVALIDATE_IMAGE_TYPES",
  INVALIDATE_IMAGE_VARIANT: "IMAGE:INVALIDATE_IMAGE_VARIANT",
  INVALIDATE_IMAGE_VARIANTS: "IMAGE:INVALIDATE_IMAGE_VARIANTS",
  INVALIDATE_IMAGE_TAG: "IMAGE:INVALIDATE_IMAGE_TAG",
  INVALIDATE_IMAGE_TAGS: "IMAGE:INVALIDATE_IMAGE_TAGS",
});

type ImageEvent = EventNamesOf<typeof imageEvents>;

class ImageEventEmitter extends EventEmitter<ImageEvent> {
  constructor() {
    super(imageEvents);
  }
}

export const imageEventEmitter = new (singleton(ImageEventEmitter))();
