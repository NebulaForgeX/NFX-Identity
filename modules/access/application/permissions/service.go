package permissions

import (
	permissionDomain "nfxid/modules/access/domain/permissions"
)

type Service struct {
	permissionRepo *permissionDomain.Repo
}

func NewService(
	permissionRepo *permissionDomain.Repo,
) *Service {
	return &Service{
		permissionRepo: permissionRepo,
	}
}
