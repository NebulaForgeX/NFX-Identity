package password_history

import (
	passwordHistoryDomain "nfxid/modules/auth/domain/password_history"
)

type Service struct {
	passwordHistoryRepo *passwordHistoryDomain.Repo
}

func NewService(
	passwordHistoryRepo *passwordHistoryDomain.Repo,
) *Service {
	return &Service{
		passwordHistoryRepo: passwordHistoryRepo,
	}
}
