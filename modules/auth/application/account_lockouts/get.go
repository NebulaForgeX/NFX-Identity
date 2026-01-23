package account_lockouts

import (
	"context"
	accountLockoutResult "nfxid/modules/auth/application/account_lockouts/results"

	"github.com/google/uuid"
)

// GetAccountLockout 根据UserID获取账户锁定
func (s *Service) GetAccountLockout(ctx context.Context, userID uuid.UUID) (accountLockoutResult.AccountLockoutRO, error) {
	domainEntity, err := s.accountLockoutRepo.Get.ByUserID(ctx, userID)
	if err != nil {
		return accountLockoutResult.AccountLockoutRO{}, err
	}
	return accountLockoutResult.AccountLockoutMapper(domainEntity), nil
}
