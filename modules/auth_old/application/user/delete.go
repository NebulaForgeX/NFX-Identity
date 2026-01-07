package user

import (
	"context"

	"github.com/google/uuid"
)

type DeleteUserCmd struct {
	UserID uuid.UUID
}

func (s *Service) DeleteUser(ctx context.Context, cmd DeleteUserCmd) error {
	u, err := s.userRepo.Get.ByID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	if err := u.Delete(); err != nil {
		return err
	}

	if err := s.userRepo.Update.Generic(ctx, u); err != nil {
		return err
	}

	return nil
}
