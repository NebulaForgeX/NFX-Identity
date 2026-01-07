package user_permission

import (
	"context"
	userPermissionCommands "nfxid/modules/permission/application/user_permission/commands"
)

func (s *Service) RevokePermission(ctx context.Context, cmd userPermissionCommands.RevokePermissionCmd) error {
	return s.userPermissionRepo.Delete.ByUserIDAndPermissionID(ctx, cmd.UserID, cmd.PermissionID)
}
