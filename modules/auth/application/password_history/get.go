package password_history

import (
	"context"
	passwordHistoryResult "nfxid/modules/auth/application/password_history/results"

	"github.com/google/uuid"
)

// GetPasswordHistory 根据ID获取密码历史
func (s *Service) GetPasswordHistory(ctx context.Context, passwordHistoryID uuid.UUID) (passwordHistoryResult.PasswordHistoryRO, error) {
	domainEntity, err := s.passwordHistoryRepo.Get.ByID(ctx, passwordHistoryID)
	if err != nil {
		return passwordHistoryResult.PasswordHistoryRO{}, err
	}
	return passwordHistoryResult.PasswordHistoryMapper(domainEntity), nil
}
