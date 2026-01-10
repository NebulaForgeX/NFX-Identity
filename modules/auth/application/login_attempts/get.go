package login_attempts

import (
	"context"
	loginAttemptResult "nfxid/modules/auth/application/login_attempts/results"

	"github.com/google/uuid"
)

// GetLoginAttempt 根据ID获取登录尝试
func (s *Service) GetLoginAttempt(ctx context.Context, loginAttemptID uuid.UUID) (loginAttemptResult.LoginAttemptRO, error) {
	domainEntity, err := s.loginAttemptRepo.Get.ByID(ctx, loginAttemptID)
	if err != nil {
		return loginAttemptResult.LoginAttemptRO{}, err
	}
	return loginAttemptResult.LoginAttemptMapper(domainEntity), nil
}
