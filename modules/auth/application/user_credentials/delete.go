package user_credentials

import (
	"context"
	userCredentialCommands "nfxid/modules/auth/application/user_credentials/commands"
)

// DeleteUserCredential 删除用户凭证（软删除）
func (s *Service) DeleteUserCredential(ctx context.Context, cmd userCredentialCommands.DeleteUserCredentialCmd) error {
	// Get domain entity
	userCredential, err := s.userCredentialRepo.Get.ByID(ctx, cmd.UserCredentialID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := userCredential.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.userCredentialRepo.Update.Generic(ctx, userCredential)
}
