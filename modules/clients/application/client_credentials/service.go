package client_credentials

import (
	clientCredentialDomain "nfxid/modules/clients/domain/client_credentials"
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
