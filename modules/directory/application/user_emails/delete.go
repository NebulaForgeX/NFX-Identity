package user_emails

import (
	"context"
	userEmailCommands "nfxid/modules/directory/application/user_emails/commands"
)

// DeleteUserEmail 删除用户邮箱（软删除）
func (s *Service) DeleteUserEmail(ctx context.Context, cmd userEmailCommands.DeleteUserEmailCmd) error {
	// Get domain entity
	userEmail, err := s.userEmailRepo.Get.ByID(ctx, cmd.UserEmailID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := userEmail.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.userEmailRepo.Update.Generic(ctx, userEmail)
}
