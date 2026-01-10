package badges

import (
	"context"
	badgeCommands "nfxid/modules/directory/application/badges/commands"
)

// DeleteBadge 删除徽章（软删除）
func (s *Service) DeleteBadge(ctx context.Context, cmd badgeCommands.DeleteBadgeCmd) error {
	// Get domain entity
	badge, err := s.badgeRepo.Get.ByID(ctx, cmd.BadgeID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := badge.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.badgeRepo.Update.Generic(ctx, badge)
}
