package handler

import (
	"context"
	"nfxid/events"
	imageApp "nfxid/modules/image/application/image"
	"nfxid/pkgs/logx"

	"github.com/ThreeDotsLabs/watermill/message"
)

type AuthHandler struct {
	imageAppSvc *imageApp.Service
}

func NewAuthHandler(imageAppSvc *imageApp.Service) *AuthHandler {
	return &AuthHandler{
		imageAppSvc: imageAppSvc,
	}
}

// OnAuthToImage_ImageDelete 监听 Auth 服务请求删除图片
func (h *AuthHandler) OnAuthToImage_ImageDelete(ctx context.Context, evt events.AuthToImage_ImageDeleteEvent, msg *message.Message) error {
	logx.S().Infof("🗑️ [Image Worker] Auth 服务请求删除图片: image_path=%s, entity_id=%s, entity_type=%s, user_id=%s",
		evt.ImagePath, evt.EntityID, evt.EntityType, evt.UserID)

	// 删除图片
	cmd := imageApp.DeleteImageByStoragePathCmd{
		StoragePath: evt.ImagePath,
	}
	if err := h.imageAppSvc.DeleteImageByStoragePath(ctx, cmd); err != nil {
		logx.S().Errorf("删除图片失败: image_path=%s, error: %v", evt.ImagePath, err)
		return err
	}

	return nil
}

// OnAuthToImage_ImageSuccess 监听 Auth 服务通知的成功事件
func (h *AuthHandler) OnAuthToImage_ImageSuccess(ctx context.Context, evt events.AuthToImage_ImageSuccessEvent, msg *message.Message) error {
	logx.S().Infof("✅ [Image Worker] Auth 服务通知操作成功: operation=%s, entity_id=%s, user_id=%s",
		evt.Operation, evt.EntityID, evt.UserID)
	return nil
}

// OnAuthToImage_ImageTest 监听 Auth 服务发送的测试消息
func (h *AuthHandler) OnAuthToImage_ImageTest(ctx context.Context, evt events.AuthToImage_ImageTestEvent, msg *message.Message) error {
	logx.S().Infof("📨 [Image Worker] Auth 服务发送测试消息: %s", evt.Message)
	return nil
}
