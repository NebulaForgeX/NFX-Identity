package user

import (
	"context"

	"nebulaid/events"
	"nebulaid/pkgs/eventbus"
	"nebulaid/pkgs/safeexec"

	"github.com/google/uuid"
)

// DeleteAccount 删除账户（特殊逻辑：删除用户及其关联数据）
func (s *Service) DeleteAccount(ctx context.Context, userID uuid.UUID) error {
	// 检查用户是否存在
	entity, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	// 删除用户资料（如果存在）
	profile, err := s.profileRepo.GetByUserID(ctx, userID)
	if err == nil && profile != nil {
		if err := profile.Delete(); err != nil {
			// 记录错误但不阻止删除用户
		} else {
			if err := s.profileRepo.Update(ctx, profile); err != nil {
				// 记录错误但不阻止删除用户
			}
		}
	}

	// 删除用户
	if err := s.userRepo.Delete(ctx, userID); err != nil {
		return err
	}

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
