package permission

import (
	permissionQueries "nfxid/modules/permission/application/permission/queries"
	permissionDomain "nfxid/modules/permission/domain/permission"
)

type Service struct {
	permissionRepo  permissionDomain.Repo
	permissionQuery permissionQueries.PermissionQuery
}

func NewService(
	permissionRepo permissionDomain.Repo,
	permissionQuery permissionQueries.PermissionQuery,
) *Service {
	return &Service{
		permissionRepo:  permissionRepo,
		permissionQuery: permissionQuery,
	}
}

