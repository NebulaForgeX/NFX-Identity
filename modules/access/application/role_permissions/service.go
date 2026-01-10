package role_permissions

import (
	rolePermissionDomain "nfxid/modules/access/domain/role_permissions"
)

type Service struct {
	rolePermissionRepo *rolePermissionDomain.Repo
}

func NewService(
	rolePermissionRepo *rolePermissionDomain.Repo,
) *Service {
	return &Service{
		rolePermissionRepo: rolePermissionRepo,
	}
}
