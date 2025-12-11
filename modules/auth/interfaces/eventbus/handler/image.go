package handler

import (
	"context"

	"nfxid/events"
	"nfxid/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
)

type ImageHandler struct {
	// 可以注入 application services
}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{}
}

// OnImageToAuth_ImageDelete 监听 Image 服务通知的图片删除事件
func (h *ImageHandler) OnImageToAuth_ImageDelete(ctx context.Context, evt events.ImageToAuth_ImageDeleteEvent, msg *message.Message) error {
	logx.S().Infof("🗑️ [Auth Worker] Image 服务通知图片已删除: image_path=%s, entity_id=%s, entity_type=%s, user_id=%s",
		evt.ImagePath, evt.EntityID, evt.EntityType, evt.UserID)
	// TODO: 处理图片删除逻辑，比如清理用户头像/背景引用
	return nil
}

// OnImageToAuth_ImageSuccess 监听 Image 服务通知的成功事件
func (h *ImageHandler) OnImageToAuth_ImageSuccess(ctx context.Context, evt events.ImageToAuth_ImageSuccessEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Auth Worker] Image 服务通知操作成功: operation=%s, entity_id=%s, user_id=%s",
		evt.Operation, evt.EntityID, evt.UserID)
	return nil
}

// OnImageToAuth_ImageTest 监听 Image 服务发送的测试消息
func (h *ImageHandler) OnImageToAuth_ImageTest(ctx context.Context, evt events.ImageToAuth_ImageTestEvent, msg *message.Message) error {
	logx.S().Infof("📨 [Auth Worker] Image 服务发送测试消息: %s", evt.Message)
	return nil
}
