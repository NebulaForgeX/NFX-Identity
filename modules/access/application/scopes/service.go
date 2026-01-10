package scopes

import (
	scopeDomain "nfxid/modules/access/domain/scopes"
)

type Service struct {
	scopeRepo *scopeDomain.Repo
}

func NewService(
	scopeRepo *scopeDomain.Repo,
) *Service {
	return &Service{
		scopeRepo: scopeRepo,
	}
}
