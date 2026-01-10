package groups

import (
	"context"
	groupCommands "nfxid/modules/tenants/application/groups/commands"
)

// DeleteGroup 删除组（软删除）
func (s *Service) DeleteGroup(ctx context.Context, cmd groupCommands.DeleteGroupCmd) error {
	// Get domain entity
	group, err := s.groupRepo.Get.ByID(ctx, cmd.GroupID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := group.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.groupRepo.Update.Generic(ctx, group)
}
