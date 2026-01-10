package groups

import (
	"context"
	groupCommands "nfxid/modules/tenants/application/groups/commands"
)

// UpdateGroup 更新组
func (s *Service) UpdateGroup(ctx context.Context, cmd groupCommands.UpdateGroupCmd) error {
	// Get domain entity
	group, err := s.groupRepo.Get.ByID(ctx, cmd.GroupID)
	if err != nil {
		return err
	}

	// Update domain entity
	if err := group.Update(cmd.Name, cmd.Type, cmd.ParentGroupID, cmd.Description, cmd.Metadata); err != nil {
		return err
	}

	// Save to repository
	return s.groupRepo.Update.Generic(ctx, group)
}
