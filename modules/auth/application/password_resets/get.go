package password_resets

import (
	"context"
	passwordResetResult "nfxid/modules/auth/application/password_resets/results"

	"github.com/google/uuid"
)

// GetPasswordReset 根据ID获取密码重置
func (s *Service) GetPasswordReset(ctx context.Context, passwordResetID uuid.UUID) (passwordResetResult.PasswordResetRO, error) {
	domainEntity, err := s.passwordResetRepo.Get.ByID(ctx, passwordResetID)
	if err != nil {
		return passwordResetResult.PasswordResetRO{}, err
	}
	return passwordResetResult.PasswordResetMapper(domainEntity), nil
}

// GetPasswordResetByResetID 根据ResetID获取密码重置
func (s *Service) GetPasswordResetByResetID(ctx context.Context, resetID string) (passwordResetResult.PasswordResetRO, error) {
	domainEntity, err := s.passwordResetRepo.Get.ByResetID(ctx, resetID)
	if err != nil {
		return passwordResetResult.PasswordResetRO{}, err
	}
	return passwordResetResult.PasswordResetMapper(domainEntity), nil
}
