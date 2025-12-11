package permission

import (
	"context"
	permissionCommands "nfxid/modules/permission/application/permission/commands"
	permissionViews "nfxid/modules/permission/application/permission/views"
)

func (s *Service) GetPermission(ctx context.Context, cmd permissionCommands.GetPermissionCmd) (*permissionViews.PermissionView, error) {
	return s.permissionQuery.GetByID(ctx, cmd.ID)
}

func (s *Service) GetPermissionByTag(ctx context.Context, cmd permissionCommands.GetPermissionByTagCmd) (*permissionViews.PermissionView, error) {
	return s.permissionQuery.GetByTag(ctx, cmd.Tag)
}

func (s *Service) ListPermissions(ctx context.Context, cmd permissionCommands.ListPermissionsCmd) ([]*permissionViews.PermissionView, error) {
	if cmd.Category != "" {
		return s.permissionQuery.GetByCategory(ctx, cmd.Category)
	}
	return s.permissionQuery.List(ctx)
}

