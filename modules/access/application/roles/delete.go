package roles

import (
	"context"
	roleCommands "nfxid/modules/access/application/roles/commands"
)

// DeleteRole 删除角色（软删除）
func (s *Service) DeleteRole(ctx context.Context, cmd roleCommands.DeleteRoleCmd) error {
	// Get domain entity
	role, err := s.roleRepo.Get.ByID(ctx, cmd.RoleID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := role.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.roleRepo.Update.Generic(ctx, role)
}
