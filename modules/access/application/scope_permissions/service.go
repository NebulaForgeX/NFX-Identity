package scope_permissions

import (
	scopePermissionDomain "nfxid/modules/access/domain/scope_permissions"
)

type Service struct {
	scopePermissionRepo *scopePermissionDomain.Repo
}

func NewService(
	scopePermissionRepo *scopePermissionDomain.Repo,
) *Service {
	return &Service{
		scopePermissionRepo: scopePermissionRepo,
	}
}
