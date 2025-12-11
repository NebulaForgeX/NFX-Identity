package user

import (
	"context"
	userDomain "nfxid/modules/auth/domain/user"
	userDomainErrors "nfxid/modules/auth/domain/user/errors"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserCmd struct {
	Editable userDomain.UserEditable
	Status   string // pending, active, deactive
}

func (s *Service) CreateUser(ctx context.Context, cmd CreateUserCmd) (*userDomain.User, error) {
	// 检查用户名、邮箱、手机号是否已存在
	if exists, _ := s.userRepo.ExistsByUsername(ctx, cmd.Editable.Username); exists {
		return nil, userDomainErrors.ErrUsernameAlreadyExists
	}
	if exists, _ := s.userRepo.ExistsByEmail(ctx, cmd.Editable.Email); exists {
		return nil, userDomainErrors.ErrEmailAlreadyExists
	}
	if exists, _ := s.userRepo.ExistsByPhone(ctx, cmd.Editable.Phone); exists {
		return nil, userDomainErrors.ErrPhoneAlreadyExists
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

	if err := s.userRepo.Create(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}
