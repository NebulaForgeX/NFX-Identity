package permission

import (
	permissionDomain "nfxid/modules/permission/domain/permission"
)

type Service struct {
	permissionRepo  *permissionDomain.Repo
	permissionQuery *permissionDomain.Query
}

func NewService(
	permissionRepo *permissionDomain.Repo,
	permissionQuery *permissionDomain.Query,
) *Service {
	return &Service{
		permissionRepo:  permissionRepo,
		permissionQuery: permissionQuery,
	}
}
