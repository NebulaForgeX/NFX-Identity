package user_phones

import (
	"context"
	"time"
	userPhoneCommands "nfxid/modules/directory/application/user_phones/commands"
)

// SetPrimaryPhone 设置主手机号
func (s *Service) SetPrimaryPhone(ctx context.Context, cmd userPhoneCommands.SetPrimaryPhoneCmd) error {
	// Get domain entity
	userPhone, err := s.userPhoneRepo.Get.ByID(ctx, cmd.UserPhoneID)
	if err != nil {
		return err
	}

	// Set primary domain entity
	if err := userPhone.SetPrimary(); err != nil {
		return err
	}

	// Save to repository
	return s.userPhoneRepo.Update.SetPrimary(ctx, cmd.UserPhoneID)
}

// VerifyPhone 验证手机号
func (s *Service) VerifyPhone(ctx context.Context, cmd userPhoneCommands.VerifyPhoneCmd) error {
	// Get domain entity
	userPhone, err := s.userPhoneRepo.Get.ByID(ctx, cmd.UserPhoneID)
	if err != nil {
		return err
	}

	// Verify domain entity
	if err := userPhone.Verify(); err != nil {
		return err
	}

	// Save to repository
	return s.userPhoneRepo.Update.Verify(ctx, cmd.UserPhoneID)
}

// UpdateVerificationCode 更新验证码
func (s *Service) UpdateVerificationCode(ctx context.Context, cmd userPhoneCommands.UpdateVerificationCodeCmd) error {
	// Get domain entity
	userPhone, err := s.userPhoneRepo.Get.ByID(ctx, cmd.UserPhoneID)
	if err != nil {
		return err
	}

	// Parse expires at
	expiresAt, err := time.Parse(time.RFC3339, cmd.VerificationExpiresAt)
	if err != nil {
		return err
	}

	// Update verification code domain entity
	if err := userPhone.UpdateVerificationCode(cmd.VerificationCode, expiresAt); err != nil {
		return err
	}

	// Save to repository
	return s.userPhoneRepo.Update.UpdateVerificationCode(ctx, cmd.UserPhoneID, cmd.VerificationCode, expiresAt)
}
