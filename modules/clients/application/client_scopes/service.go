package client_scopes

import (
	clientScopeDomain "nfxid/modules/clients/domain/client_scopes"
)

type Service struct {
	clientScopeRepo *clientScopeDomain.Repo
}

func NewService(
	clientScopeRepo *clientScopeDomain.Repo,
) *Service {
	return &Service{
		clientScopeRepo: clientScopeRepo,
	}
}
