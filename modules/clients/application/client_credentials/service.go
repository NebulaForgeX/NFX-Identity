package client_credentials

import (
	clientCredentialDomain "nfxidentity/modules/clients/domain/client_credentials"
)

type Service struct {
	clientCredentialRepo *clientCredentialDomain.Repo
}

func NewService(
	clientCredentialRepo *clientCredentialDomain.Repo,
) *Service {
	return &Service{
		clientCredentialRepo: clientCredentialRepo,
	}
}
