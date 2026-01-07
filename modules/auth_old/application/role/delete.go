package role

import (
	"context"

	"github.com/google/uuid"
)

type DeleteRoleCmd struct {
	RoleID uuid.UUID
}

func (s *Service) DeleteRole(ctx context.Context, cmd DeleteRoleCmd) error {
	r, err := s.roleRepo.Get.ByID(ctx, cmd.RoleID)
	if err != nil {
		return err
	}

	if err := r.Delete(); err != nil {
		return err
	}

	if err := s.roleRepo.Update.Generic(ctx, r); err != nil {
		return err
	}

	return nil
}
