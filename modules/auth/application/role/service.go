package role

import (
	roleQueries "nfxid/modules/auth/application/role/queries"
	roleDomain "nfxid/modules/auth/domain/role"
)

type Service struct {
	roleRepo  *roleDomain.Repo
	roleQuery roleQueries.RoleQuery
}

func NewService(
	roleRepo *roleDomain.Repo,
	roleQuery roleQueries.RoleQuery,
) *Service {
	return &Service{
		roleRepo:  roleRepo,
		roleQuery: roleQuery,
	}
}
