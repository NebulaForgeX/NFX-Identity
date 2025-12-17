package profile_occupation

import (
	"context"
	occupationCommands "nfxid/modules/auth/application/profile_occupation/commands"
	occupationDomain "nfxid/modules/auth/domain/profile_occupation"
)

func (s *Service) CreateOccupation(ctx context.Context, cmd occupationCommands.CreateOccupationCmd) (*occupationDomain.Occupation, error) {
	// 使用 domain factory 创建实体
	o, err := occupationDomain.NewOccupation(occupationDomain.NewOccupationParams{
		ProfileID: cmd.ProfileID,
		Editable:  cmd.Editable,
	})
	if err != nil {
		return nil, err
	}

	if err := s.occupationRepo.Create.New(ctx, o); err != nil {
		return nil, err
	}

	return o, nil
}
