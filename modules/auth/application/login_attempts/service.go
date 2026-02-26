package login_attempts

import (
	loginAttemptDomain "nfxidentity/modules/auth/domain/login_attempts"
)

type Service struct {
	loginAttemptRepo *loginAttemptDomain.Repo
}

func NewService(
	loginAttemptRepo *loginAttemptDomain.Repo,
) *Service {
	return &Service{
		loginAttemptRepo: loginAttemptRepo,
	}
}
