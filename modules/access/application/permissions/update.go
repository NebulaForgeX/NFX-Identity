package permissions

import (
	"context"
	permissionCommands "nfxid/modules/access/application/permissions/commands"
)

// UpdatePermission 更新权限
func (s *Service) UpdatePermission(ctx context.Context, cmd permissionCommands.UpdatePermissionCmd) error {
	// Get domain entity
	permission, err := s.permissionRepo.Get.ByID(ctx, cmd.PermissionID)
	if err != nil {
		return err
	}

	// Update domain entity
	key := permission.Key() // Keep existing key
	if err := permission.Update(key, cmd.Name, cmd.Description); err != nil {
		return err
	}

	// Save to repository
	return s.permissionRepo.Update.Generic(ctx, permission)
}
