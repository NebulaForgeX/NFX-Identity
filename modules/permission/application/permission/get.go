package permission

import (
	"context"
	permissionCommands "nfxid/modules/permission/application/permission/commands"
	permissionViews "nfxid/modules/permission/application/permission/views"
	permissionDomain "nfxid/modules/permission/domain/permission"
)

func (s *Service) GetPermission(ctx context.Context, cmd permissionCommands.GetPermissionCmd) (*permissionViews.PermissionView, error) {
	domainView, err := s.permissionQuery.Single.ByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}
	view := permissionViews.PermissionViewMapper(*domainView)
	return &view, nil
}

func (s *Service) GetPermissionByTag(ctx context.Context, cmd permissionCommands.GetPermissionByTagCmd) (*permissionViews.PermissionView, error) {
	domainView, err := s.permissionQuery.Single.ByTag(ctx, cmd.Tag)
	if err != nil {
		return nil, err
	}
	view := permissionViews.PermissionViewMapper(*domainView)
	return &view, nil
}

func (s *Service) ListPermissions(ctx context.Context, cmd permissionCommands.ListPermissionsCmd) ([]*permissionViews.PermissionView, error) {
	if cmd.Category != "" {
		domainViews, err := s.permissionQuery.List.ByCategory(ctx, cmd.Category)
		if err != nil {
			return nil, err
		}
		result := make([]*permissionViews.PermissionView, len(domainViews))
		for i, v := range domainViews {
			view := permissionViews.PermissionViewMapper(*v)
			result[i] = &view
		}
		return result, nil
	}
	// Use Generic list query
	listQuery := permissionDomain.ListQuery{}
	listQuery.Normalize()
	domainViews, _, err := s.permissionQuery.List.Generic(ctx, listQuery)
	if err != nil {
		return nil, err
	}
	result := make([]*permissionViews.PermissionView, len(domainViews))
	for i, v := range domainViews {
		view := permissionViews.PermissionViewMapper(*v)
		result[i] = &view
	}
	return result, nil
}
