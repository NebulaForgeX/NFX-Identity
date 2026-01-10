package account_lockouts

import (
	accountLockoutDomain "nfxid/modules/auth/domain/account_lockouts"
)

type Service struct {
	accountLockoutRepo *accountLockoutDomain.Repo
}

func NewService(
	accountLockoutRepo *accountLockoutDomain.Repo,
) *Service {
	return &Service{
		accountLockoutRepo: accountLockoutRepo,
	}
}
