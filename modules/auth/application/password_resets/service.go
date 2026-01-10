package password_resets

import (
	passwordResetDomain "nfxid/modules/auth/domain/password_resets"
)

type Service struct {
	passwordResetRepo *passwordResetDomain.Repo
}

func NewService(
	passwordResetRepo *passwordResetDomain.Repo,
) *Service {
	return &Service{
		passwordResetRepo: passwordResetRepo,
	}
}
