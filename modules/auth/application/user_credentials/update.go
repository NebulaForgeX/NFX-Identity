package user_credentials

import (
	"context"
	userCredentialCommands "nfxid/modules/auth/application/user_credentials/commands"
)

// UpdateUserCredential 更新用户凭证
func (s *Service) UpdateUserCredential(ctx context.Context, cmd userCredentialCommands.UpdateUserCredentialCmd) error {
	// Get domain entity
	userCredential, err := s.userCredentialRepo.Get.ByID(ctx, cmd.UserCredentialID)
	if err != nil {
		return err
	}

	// Update status if provided
	if err := userCredential.UpdateStatus(cmd.Status); err != nil {
		return err
	}

	// Update must change password if needed
	if err := userCredential.SetMustChangePassword(cmd.MustChangePassword); err != nil {
		return err
	}

	// Save to repository
	return s.userCredentialRepo.Update.Generic(ctx, userCredential)
}

// UpdatePassword 更新密码
func (s *Service) UpdatePassword(ctx context.Context, cmd userCredentialCommands.UpdatePasswordCmd) error {
	// Get domain entity
	userCredential, err := s.userCredentialRepo.Get.ByUserID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	// Update password domain entity
	if err := userCredential.UpdatePassword(cmd.PasswordHash, cmd.HashAlg, cmd.HashParams); err != nil {
		return err
	}

	// Save to repository
	return s.userCredentialRepo.Update.UpdatePassword(ctx, cmd.UserID, cmd.PasswordHash, cmd.HashAlg, cmd.HashParams)
}

// UpdateStatus 更新状态
func (s *Service) UpdateStatus(ctx context.Context, cmd userCredentialCommands.UpdateStatusCmd) error {
	// Get domain entity
	userCredential, err := s.userCredentialRepo.Get.ByUserID(ctx, cmd.UserID)
	if err != nil {
		return err
	}

	// Update status domain entity
	if err := userCredential.UpdateStatus(cmd.Status); err != nil {
		return err
	}

	// Save to repository
	return s.userCredentialRepo.Update.UpdateStatus(ctx, cmd.UserID, cmd.Status)
}
