package user

import (
	"context"
	userDomain "nebulaid/modules/auth/domain/user"
	userDomainErrors "nebulaid/modules/auth/domain/user/errors"

	"github.com/google/uuid"
)

type UpdateUserCmd struct {
	UserID   uuid.UUID
	Editable userDomain.UserEditable
}

func (s *Service) UpdateUser(ctx context.Context, cmd UpdateUserCmd) error {
	u, err := s.userRepo.GetByID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	// 检查用户名、邮箱、手机号是否已存在（排除当前用户）
	if cmd.Editable.Username != u.Editable().Username {
		if exists, _ := s.userRepo.ExistsByUsername(ctx, cmd.Editable.Username); exists {
			return userDomainErrors.ErrUsernameAlreadyExists
		}
	}
	if cmd.Editable.Email != u.Editable().Email {
		if exists, _ := s.userRepo.ExistsByEmail(ctx, cmd.Editable.Email); exists {
			return userDomainErrors.ErrEmailAlreadyExists
		}
	}
	if cmd.Editable.Phone != u.Editable().Phone {
		if exists, _ := s.userRepo.ExistsByPhone(ctx, cmd.Editable.Phone); exists {
			return userDomainErrors.ErrPhoneAlreadyExists
		}
	}

	if err := u.EnsureEditable(cmd.Editable); err != nil {
		return err
	}

	if err := u.Update(cmd.Editable); err != nil {
		return err
	}

	if err := s.userRepo.Update(ctx, u); err != nil {
		return err
	}

	return nil
}
