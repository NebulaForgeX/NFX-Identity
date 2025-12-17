package user_role

import (
	"context"

	"github.com/google/uuid"
)

type DeleteUserRoleCmd struct {
	UserRoleID uuid.UUID
}

func (s *Service) DeleteUserRole(ctx context.Context, cmd DeleteUserRoleCmd) error {
	return s.userRoleRepo.Delete.ByID(ctx, cmd.UserRoleID)
}

type DeleteUserRoleByUserAndRoleCmd struct {
	UserID uuid.UUID
	RoleID uuid.UUID
}

func (s *Service) DeleteUserRoleByUserAndRole(ctx context.Context, cmd DeleteUserRoleByUserAndRoleCmd) error {
	return s.userRoleRepo.Delete.ByUserAndRole(ctx, cmd.UserID, cmd.RoleID)
}
