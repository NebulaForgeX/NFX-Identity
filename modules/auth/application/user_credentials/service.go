package user_credentials

import (
	userCredentialDomain "nfxidentity/modules/auth/domain/user_credentials"
)

type Service struct {
	userCredentialRepo *userCredentialDomain.Repo
}

func NewService(
	userCredentialRepo *userCredentialDomain.Repo,
) *Service {
	return &Service{
		userCredentialRepo: userCredentialRepo,
	}
}
