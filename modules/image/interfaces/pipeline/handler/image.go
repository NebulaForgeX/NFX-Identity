package handler

import (
	"context"

	"nfxid/events/image"
	"nfxid/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
)

type ImageHandler struct {
	// 可以注入 application services 或其他依赖
}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{}
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
