package badges

import (
	"context"
	badgeCommands "nfxid/modules/directory/application/badges/commands"
)

// UpdateBadge 更新徽章
func (s *Service) UpdateBadge(ctx context.Context, cmd badgeCommands.UpdateBadgeCmd) error {
	// Get domain entity
	badge, err := s.badgeRepo.Get.ByID(ctx, cmd.BadgeID)
	if err != nil {
		return err
	}

	// Update domain entity
	if err := badge.Update(cmd.Name, cmd.Description, cmd.IconURL, cmd.Color, cmd.Category); err != nil {
		return err
	}

	// Save to repository
	return s.badgeRepo.Update.Generic(ctx, badge)
}
