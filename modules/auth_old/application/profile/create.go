package profile

import (
	"context"
	profileCommands "nfxid/modules/auth/application/profile/commands"
	profileDomain "nfxid/modules/auth/domain/profile"
	profileDomainErrors "nfxid/modules/auth/domain/profile/errors"
)

func (s *Service) CreateProfile(ctx context.Context, cmd profileCommands.CreateProfileCmd) (*profileDomain.Profile, error) {
	// 检查用户是否已有资料
	if exists, _ := s.profileRepo.Check.ByUserID(ctx, cmd.UserID); exists {
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

	if err := s.profileRepo.Create.New(ctx, p); err != nil {
		return nil, err
	}

	return p, nil
}
