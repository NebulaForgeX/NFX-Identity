package handler

import (
	"context"

	"nfxid/events/directory"
	"nfxid/events/image"
	imageCommands "nfxid/modules/image/application/images/commands"
	imageApp "nfxid/modules/image/application/images"
	"nfxid/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

type ImageHandler struct {
	imageAppSvc *imageApp.Service
}

func NewImageHandler(imageAppSvc *imageApp.Service) *ImageHandler {
	return &ImageHandler{imageAppSvc: imageAppSvc}
}

// OnUserAvatarReplaced 监听用户头像替换事件，删除被替换的旧图片
func (h *ImageHandler) OnUserAvatarReplaced(ctx context.Context, evt directory.UserAvatarReplacedEvent, msg *message.Message) error {
	if evt.OldImageID == "" {
		return nil
	}
	imageID, err := uuid.Parse(evt.OldImageID)
	if err != nil {
		logx.S().Warnf("[Image Pipeline] invalid old_image_id in UserAvatarReplaced: %s", evt.OldImageID)
		return nil
	}
	err = h.imageAppSvc.DeleteImage(ctx, imageCommands.DeleteImageCmd{ImageID: imageID})
	if err != nil {
		logx.S().Warnf("[Image Pipeline] delete old avatar image failed: image_id=%s, err=%v", evt.OldImageID, err)
		return err
	}
	logx.S().Infof("[Image Pipeline] deleted old avatar image: image_id=%s", evt.OldImageID)
	return nil
}

// OnImagesInvalidateCache 监听 Images 缓存清除事件
func (h *ImageHandler) OnImagesInvalidateCache(ctx context.Context, evt image.ImagesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Image Pipeline] 收到 Images 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnImageTypesInvalidateCache 监听 ImageTypes 缓存清除事件
func (h *ImageHandler) OnImageTypesInvalidateCache(ctx context.Context, evt image.ImageTypesInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Image Pipeline] 收到 ImageTypes 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnImageVariantsInvalidateCache 监听 ImageVariants 缓存清除事件
func (h *ImageHandler) OnImageVariantsInvalidateCache(ctx context.Context, evt image.ImageVariantsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Image Pipeline] 收到 ImageVariants 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}

// OnImageTagsInvalidateCache 监听 ImageTags 缓存清除事件
func (h *ImageHandler) OnImageTagsInvalidateCache(ctx context.Context, evt image.ImageTagsInvalidateCacheEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Image Pipeline] 收到 ImageTags 缓存清除事件: id=%s, prefix=%s, namespace=%s",
		evt.ID, evt.Prefix, evt.Namespace)
	// 缓存清除事件处理完成
	return nil
}
