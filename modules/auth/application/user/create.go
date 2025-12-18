package user

import (
	"context"
	"nfxid/events"
	userDomain "nfxid/modules/auth/domain/user"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/safeexec"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserCmd struct {
	Editable userDomain.UserEditable
	Status   string // pending, active, deactive
}

func (s *Service) CreateUser(ctx context.Context, cmd CreateUserCmd) (*userDomain.User, error) {
	// 检查用户名、邮箱、手机号是否已存在
	if exists, _ := s.userRepo.Check.ByUsername(ctx, cmd.Editable.Username); exists {
		return nil, userDomainErrors.ErrUsernameAlreadyExists
	}
	if exists, _ := s.userRepo.Check.ByEmail(ctx, cmd.Editable.Email); exists {
		return nil, userDomainErrors.ErrEmailAlreadyExists
	}
	if cmd.Editable.Phone != nil {
		if exists, _ := s.userRepo.Check.ByPhone(ctx, *cmd.Editable.Phone); exists {
			return nil, userDomainErrors.ErrPhoneAlreadyExists
		}
	}

	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Editable.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 设置密码
	editable := cmd.Editable
	editable.Password = string(hashedPassword)

	// 使用 domain factory 创建实体
	status := cmd.Status
	if status == "" {
		status = "active"
	}

	u, err := userDomain.NewUser(userDomain.NewUserParams{
		Editable: editable,
		Status:   status,
	})
	if err != nil {
		return nil, err
	}

	if err := s.userRepo.Create.New(ctx, u); err != nil {
		return nil, err
	}

	// 发布用户创建事件（Auth -> Auth 内部事件，用于通知 profile 等服务创建关联数据）
	safeexec.SafeGo(func() error {
		phone := ""
		if u.Editable().Phone != nil {
			phone = *u.Editable().Phone
		}
		event := events.AuthToAuth_UserCreatedEvent{
			UserID:   u.ID().String(),
			Username: u.Editable().Username,
			Email:    u.Editable().Email,
			Phone:    phone,
			Status:   u.Status(),
			Details: map[string]interface{}{
				"is_verified": u.IsVerified(),
			},
		}
		return eventbus.PublishEvent(context.Background(), s.busPublisher, event)
	})

	return u, nil
}
