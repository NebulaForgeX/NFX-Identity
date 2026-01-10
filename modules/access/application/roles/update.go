package roles

import (
	"context"
	roleCommands "nfxid/modules/access/application/roles/commands"
)

// UpdateRole 更新角色
func (s *Service) UpdateRole(ctx context.Context, cmd roleCommands.UpdateRoleCmd) error {
	// Get domain entity
	role, err := s.roleRepo.Get.ByID(ctx, cmd.RoleID)
	if err != nil {
		return err
	}

	// Update domain entity
	key := role.Key() // Keep existing key
	if err := role.Update(key, cmd.Name, cmd.Description, cmd.ScopeType); err != nil {
		return err
	}

	// Save to repository
	return s.roleRepo.Update.Generic(ctx, role)
}
