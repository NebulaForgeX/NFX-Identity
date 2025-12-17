package user

import (
	"context"
	"nfxid/events"
	userDomain "nfxid/modules/auth/domain/user"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/safeexec"

	"github.com/google/uuid"
)

type UpdateUserCmd struct {
	UserID   uuid.UUID
	Editable userDomain.UserEditable
}

func (s *Service) UpdateUser(ctx context.Context, cmd UpdateUserCmd) error {
	u, err := s.userRepo.Get.ByID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	// 检查用户名、邮箱、手机号是否已存在（排除当前用户）
	if cmd.Editable.Username != u.Editable().Username {
		if exists, _ := s.userRepo.Check.ByUsername(ctx, cmd.Editable.Username); exists {
			return userDomainErrors.ErrUsernameAlreadyExists
		}
	}
	if cmd.Editable.Email != u.Editable().Email {
		if exists, _ := s.userRepo.Check.ByEmail(ctx, cmd.Editable.Email); exists {
			return userDomainErrors.ErrEmailAlreadyExists
		}
	}
	if cmd.Editable.Phone != u.Editable().Phone {
		if exists, _ := s.userRepo.Check.ByPhone(ctx, cmd.Editable.Phone); exists {
			return userDomainErrors.ErrPhoneAlreadyExists
		}
	}

	if err := u.EnsureEditable(cmd.Editable); err != nil {
		return err
	}

	if err := u.Update(cmd.Editable); err != nil {
		return err
	}

	if err := s.userRepo.Update.Generic(ctx, u); err != nil {
		return err
	}

	// 发布用户更新事件（Auth -> Auth 内部事件，用于通知其他服务）
	safeexec.SafeGo(func() error {
		event := events.AuthToAuth_UserUpdatedEvent{
			UserID:   u.ID().String(),
			Username: u.Editable().Username,
			Email:    u.Editable().Email,
			Phone:    u.Editable().Phone,
			Details: map[string]interface{}{
				"status":      u.Status(),
				"is_verified": u.IsVerified(),
			},
		}
		return eventbus.PublishEvent(context.Background(), s.busPublisher, event)
	})

	return nil
}
