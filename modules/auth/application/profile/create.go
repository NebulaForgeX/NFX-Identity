package profile

import (
	"context"
	profileCommands "nebulaid/modules/auth/application/profile/commands"
	profileDomain "nebulaid/modules/auth/domain/profile"
	profileDomainErrors "nebulaid/modules/auth/domain/profile/errors"
)

func (s *Service) CreateProfile(ctx context.Context, cmd profileCommands.CreateProfileCmd) (*profileDomain.Profile, error) {
	// 检查用户是否已有资料
	if exists, _ := s.profileRepo.ExistsByUserID(ctx, cmd.UserID); exists {
		return nil, profileDomainErrors.ErrProfileAlreadyExists
	}

	// 使用 domain factory 创建实体
	p, err := profileDomain.NewProfile(profileDomain.NewProfileParams{
		UserID:   cmd.UserID,
		Editable: cmd.Editable,
	})
	if err != nil {
		return nil, err
	}

	if err := s.profileRepo.Create(ctx, p); err != nil {
		return nil, err
	}

	return p, nil
}
