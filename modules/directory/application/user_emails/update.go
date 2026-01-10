package user_emails

import (
	"context"
	userEmailCommands "nfxid/modules/directory/application/user_emails/commands"
)

// SetPrimaryEmail 设置主邮箱
func (s *Service) SetPrimaryEmail(ctx context.Context, cmd userEmailCommands.SetPrimaryEmailCmd) error {
	// Get domain entity
	userEmail, err := s.userEmailRepo.Get.ByID(ctx, cmd.UserEmailID)
	if err != nil {
		return err
	}

	// Set primary domain entity
	if err := userEmail.SetPrimary(); err != nil {
		return err
	}

	// Save to repository
	return s.userEmailRepo.Update.SetPrimary(ctx, cmd.UserEmailID)
}

// VerifyEmail 验证邮箱
func (s *Service) VerifyEmail(ctx context.Context, cmd userEmailCommands.VerifyEmailCmd) error {
	// Get domain entity
	userEmail, err := s.userEmailRepo.Get.ByID(ctx, cmd.UserEmailID)
	if err != nil {
		return err
	}

	// Verify domain entity
	if err := userEmail.Verify(); err != nil {
		return err
	}

	// Save to repository
	return s.userEmailRepo.Update.Verify(ctx, cmd.UserEmailID)
}
