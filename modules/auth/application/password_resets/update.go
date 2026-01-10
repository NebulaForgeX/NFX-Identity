package password_resets

import (
	"context"
	passwordResetCommands "nfxid/modules/auth/application/password_resets/commands"
)

// MarkAsUsed 标记密码重置为已使用
func (s *Service) MarkAsUsed(ctx context.Context, cmd passwordResetCommands.MarkAsUsedCmd) error {
	// Get domain entity
	passwordReset, err := s.passwordResetRepo.Get.ByResetID(ctx, cmd.ResetID)
	if err != nil {
		return err
	}

	// Mark as used domain entity
	if err := passwordReset.MarkAsUsed(); err != nil {
		return err
	}

	// Save to repository
	return s.passwordResetRepo.Update.MarkAsUsed(ctx, cmd.ResetID)
}

// IncrementAttemptCount 增加尝试次数
func (s *Service) IncrementAttemptCount(ctx context.Context, cmd passwordResetCommands.IncrementAttemptCountCmd) error {
	// Get domain entity
	passwordReset, err := s.passwordResetRepo.Get.ByResetID(ctx, cmd.ResetID)
	if err != nil {
		return err
	}

	// Increment attempt count domain entity
	if err := passwordReset.IncrementAttemptCount(); err != nil {
		return err
	}

	// Save to repository
	return s.passwordResetRepo.Update.IncrementAttemptCount(ctx, cmd.ResetID)
}

// UpdateStatus 更新状态
func (s *Service) UpdateStatus(ctx context.Context, cmd passwordResetCommands.UpdateStatusCmd) error {
	// Get domain entity
	passwordReset, err := s.passwordResetRepo.Get.ByResetID(ctx, cmd.ResetID)
	if err != nil {
		return err
	}

	// Update status domain entity
	if err := passwordReset.UpdateStatus(cmd.Status); err != nil {
		return err
	}

	// Save to repository
	return s.passwordResetRepo.Update.UpdateStatus(ctx, cmd.ResetID, cmd.Status)
}
