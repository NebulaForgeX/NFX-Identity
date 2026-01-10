package account_lockouts

import (
	"context"
	accountLockoutResult "nfxid/modules/auth/application/account_lockouts/results"

	"github.com/google/uuid"
)

// GetAccountLockout 根据UserID和TenantID获取账户锁定
func (s *Service) GetAccountLockout(ctx context.Context, userID, tenantID uuid.UUID) (accountLockoutResult.AccountLockoutRO, error) {
	domainEntity, err := s.accountLockoutRepo.Get.ByUserIDAndTenantID(ctx, userID, tenantID)
	if err != nil {
		return accountLockoutResult.AccountLockoutRO{}, err
	}
	return accountLockoutResult.AccountLockoutMapper(domainEntity), nil
}
