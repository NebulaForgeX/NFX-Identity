package user

import (
	"context"

	"nfxid/events"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/safeexec"

	"github.com/google/uuid"
)

// DeleteAccount 删除账户（特殊逻辑：删除用户及其关联数据）
func (s *Service) DeleteAccount(ctx context.Context, userID uuid.UUID) error {
	// 检查用户是否存在
	entity, err := s.userRepo.Get.ByID(ctx, userID)
	if err != nil {
		return err
	}

	// 删除用户
	if err := s.userRepo.Delete.ByID(ctx, userID); err != nil {
		return err
	}

	// 发布用户删除事件（Auth -> Auth 内部事件，用于通知 profile 等服务删除关联数据）
	safeexec.SafeGo(func() error {
		event := events.AuthToAuth_UserDeletedEvent{
			UserID:   userID.String(),
			Username: entity.Editable().Username,
			Email:    entity.Editable().Email,
			Details: map[string]interface{}{
				"phone": entity.Editable().Phone,
			},
		}
		return eventbus.PublishEvent(context.Background(), s.busPublisher, event)
	})

	// 发布删除账户事件（Auth -> Auth 内部事件）
	safeexec.SafeGo(func() error {
		event := events.AuthToAuth_SuccessEvent{
			Operation: "auth.account.deleted",
			EntityID:  userID.String(),
			UserID:    userID.String(),
			Details: map[string]interface{}{
				"username": entity.Editable().Username,
				"email":    entity.Editable().Email,
			},
		}
		return eventbus.PublishEvent(context.Background(), s.busPublisher, event)
	})

	return nil
}
