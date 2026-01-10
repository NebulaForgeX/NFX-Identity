package role_permissions

import (
	"context"
	rolePermissionCommands "nfxid/modules/access/application/role_permissions/commands"
)

// DeleteRolePermission 删除角色权限
func (s *Service) DeleteRolePermission(ctx context.Context, cmd rolePermissionCommands.DeleteRolePermissionCmd) error {
	// Delete from repository (hard delete)
	return s.rolePermissionRepo.Delete.ByID(ctx, cmd.RolePermissionID)
}

// DeleteRolePermissionByRoleAndPermission 根据角色和权限删除角色权限
func (s *Service) DeleteRolePermissionByRoleAndPermission(ctx context.Context, cmd rolePermissionCommands.DeleteRolePermissionByRoleAndPermissionCmd) error {
	// Delete from repository (hard delete)
	return s.rolePermissionRepo.Delete.ByRoleIDAndPermissionID(ctx, cmd.RoleID, cmd.PermissionID)
}

// DeleteRolePermissionsByRole 根据角色删除所有角色权限
func (s *Service) DeleteRolePermissionsByRole(ctx context.Context, cmd rolePermissionCommands.DeleteRolePermissionsByRoleCmd) error {
	// Delete from repository (hard delete)
	return s.rolePermissionRepo.Delete.ByRoleID(ctx, cmd.RoleID)
}
