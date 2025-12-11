package badge

import (
	"context"
	badgeCommands "nfxid/modules/auth/application/badge/commands"
	"nfxid/modules/auth/domain/badge"
)

func (s *Service) CreateBadge(ctx context.Context, cmd badgeCommands.CreateBadgeCmd) (*badge.Badge, error) {
	b, err := badge.NewBadge(badge.NewBadgeParams(cmd))
	if err != nil {
		return nil, err
	}

	if err := s.badgeRepo.Create(ctx, b); err != nil {
		return nil, err
	}

	return b, nil
}
