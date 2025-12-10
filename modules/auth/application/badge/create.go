package badge

import (
	"context"
	badgeCommands "nebulaid/modules/auth/application/badge/commands"
	"nebulaid/modules/auth/domain/badge"
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
