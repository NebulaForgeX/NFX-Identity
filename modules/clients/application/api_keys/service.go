package api_keys

import (
	apiKeyDomain "nfxidentity/modules/clients/domain/api_keys"
)

type Service struct {
	apiKeyRepo *apiKeyDomain.Repo
}

func NewService(
	apiKeyRepo *apiKeyDomain.Repo,
) *Service {
	return &Service{
		apiKeyRepo: apiKeyRepo,
	}
}
