package user_permission

import (
	"context"
	userPermissionCommands "nfxid/modules/permission/application/user_permission/commands"
	userPermissionDomain "nfxid/modules/permission/domain/user_permission"
	userPermissionDomainErrors "nfxid/modules/permission/domain/user_permission/errors"
)

func (s *Service) AssignPermission(ctx context.Context, cmd userPermissionCommands.AssignPermissionCmd) error {
	// Check if already exists
	exists, err := s.userPermissionRepo.Check.ByUserIDAndPermissionID(ctx, cmd.UserID, cmd.PermissionID)
	if err != nil {
		return err
	}
	if exists {
		return userPermissionDomainErrors.ErrUserPermissionAlreadyExists
	}

	up, err := userPermissionDomain.NewUserPermission(userPermissionDomain.NewUserPermissionParams{
		UserID:       cmd.UserID,
		PermissionID: cmd.PermissionID,
	})
	if err != nil {
		return err
	}

	return s.userPermissionRepo.Create.New(ctx, up)
}
