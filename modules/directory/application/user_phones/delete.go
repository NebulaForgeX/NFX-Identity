package user_phones

import (
	"context"
	userPhoneCommands "nfxid/modules/directory/application/user_phones/commands"
)

// DeleteUserPhone 删除用户手机号（软删除）
func (s *Service) DeleteUserPhone(ctx context.Context, cmd userPhoneCommands.DeleteUserPhoneCmd) error {
	// Get domain entity
	userPhone, err := s.userPhoneRepo.Get.ByID(ctx, cmd.UserPhoneID)
	if err != nil {
		return err
	}

	// Delete domain entity (soft delete)
	if err := userPhone.Delete(); err != nil {
		return err
	}

	// Save to repository
	return s.userPhoneRepo.Update.Generic(ctx, userPhone)
}
