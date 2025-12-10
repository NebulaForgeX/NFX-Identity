package badge

import (
	"context"
	badgeCommands "nebulaid/modules/auth/application/badge/commands"
)

func (s *Service) UpdateBadge(ctx context.Context, cmd badgeCommands.UpdateBadgeCmd) error {
	b, err := s.badgeRepo.GetByID(ctx, cmd.BadgeID)
	if err != nil {
		return err
	}

	if err := b.Update(cmd.Editable); err != nil {
		return err
	}

	if err := s.badgeRepo.Update(ctx, b); err != nil {
		return err
	}

	return nil
}
