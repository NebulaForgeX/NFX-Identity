package scope_permissions

import (
	"context"
	scopePermissionCommands "nfxid/modules/access/application/scope_permissions/commands"
)

// DeleteScopePermission 删除作用域权限
func (s *Service) DeleteScopePermission(ctx context.Context, cmd scopePermissionCommands.DeleteScopePermissionCmd) error {
	// Delete from repository (hard delete)
	return s.scopePermissionRepo.Delete.ByID(ctx, cmd.ScopePermissionID)
}

// DeleteScopePermissionByScopeAndPermission 根据作用域和权限删除作用域权限
func (s *Service) DeleteScopePermissionByScopeAndPermission(ctx context.Context, cmd scopePermissionCommands.DeleteScopePermissionByScopeAndPermissionCmd) error {
	// Delete from repository (hard delete)
	return s.scopePermissionRepo.Delete.ByScopeAndPermissionID(ctx, cmd.Scope, cmd.PermissionID)
}

// DeleteScopePermissionsByScope 根据作用域删除所有作用域权限
func (s *Service) DeleteScopePermissionsByScope(ctx context.Context, cmd scopePermissionCommands.DeleteScopePermissionsByScopeCmd) error {
	// Delete from repository (hard delete)
	return s.scopePermissionRepo.Delete.ByScope(ctx, cmd.Scope)
}
