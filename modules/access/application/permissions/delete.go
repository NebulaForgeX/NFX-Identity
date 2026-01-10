package permissions

import (
	"context"
	permissionCommands "nfxid/modules/access/application/permissions/commands"
)

// DeletePermission 删除权限（软删除）
func (s *Service) DeletePermission(ctx context.Context, cmd permissionCommands.DeletePermissionCmd) error {
	// Get domain entity
	permission, err := s.permissionRepo.Get.ByID(ctx, cmd.PermissionID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := permission.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.permissionRepo.Update.Generic(ctx, permission)
}
