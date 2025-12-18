package user_permission

import (
	"context"
	userPermissionCommands "nfxid/modules/permission/application/user_permission/commands"
	userPermissionViews "nfxid/modules/permission/application/user_permission/views"
)

func (s *Service) GetUserPermissions(ctx context.Context, cmd userPermissionCommands.GetUserPermissionsCmd) ([]*userPermissionViews.UserPermissionView, error) {
	domainViews, err := s.userPermissionQuery.List.ByUserID(ctx, cmd.UserID)
	if err != nil {
		return nil, err
	}
	result := make([]*userPermissionViews.UserPermissionView, len(domainViews))
	for i, v := range domainViews {
		view := userPermissionViews.UserPermissionViewMapper(*v)
		result[i] = &view
	}
	return result, nil
}

func (s *Service) GetUserPermissionTags(ctx context.Context, cmd userPermissionCommands.GetUserPermissionsCmd) ([]string, error) {
	return s.userPermissionQuery.List.PermissionTagsByUserID(ctx, cmd.UserID)
}

func (s *Service) CheckPermission(ctx context.Context, cmd userPermissionCommands.CheckPermissionCmd) (bool, error) {
	tags, err := s.userPermissionQuery.List.PermissionTagsByUserID(ctx, cmd.UserID)
	if err != nil {
		return false, err
	}

	for _, tag := range tags {
		if tag == cmd.Tag {
			return true, nil
		}
	}
	return false, nil
}
