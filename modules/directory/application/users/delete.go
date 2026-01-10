package users

import (
	"context"
	userCommands "nfxid/modules/directory/application/users/commands"
)

// DeleteUser 删除用户（软删除）
func (s *Service) DeleteUser(ctx context.Context, cmd userCommands.DeleteUserCmd) error {
	// Get domain entity
	user, err := s.userRepo.Get.ByID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := user.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.userRepo.Update.Generic(ctx, user)
}
