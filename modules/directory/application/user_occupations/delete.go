package user_occupations

import (
	"context"
	userOccupationCommands "nfxid/modules/directory/application/user_occupations/commands"
)

// DeleteUserOccupation 删除用户职业经历（软删除）
func (s *Service) DeleteUserOccupation(ctx context.Context, cmd userOccupationCommands.DeleteUserOccupationCmd) error {
	// Get domain entity
	userOccupation, err := s.userOccupationRepo.Get.ByID(ctx, cmd.UserOccupationID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := userOccupation.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.userOccupationRepo.Update.Generic(ctx, userOccupation)
}
