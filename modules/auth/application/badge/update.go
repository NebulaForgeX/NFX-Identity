package badge

import (
	"context"
	badgeCommands "nfxid/modules/auth/application/badge/commands"
)

func (s *Service) UpdateBadge(ctx context.Context, cmd badgeCommands.UpdateBadgeCmd) error {
	b, err := s.badgeRepo.Get.ByID(ctx, cmd.BadgeID)
	if err != nil {
		return err
	}

	if err := b.Update(cmd.Editable); err != nil {
		return err
	}

	if err := s.badgeRepo.Update.Generic(ctx, b); err != nil {
		return err
	}

	return nil
}
