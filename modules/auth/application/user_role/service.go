package user_role

import (
	userRoleQueries "nfxid/modules/auth/application/user_role/queries"
	userRoleDomain "nfxid/modules/auth/domain/user_role"
)

type Service struct {
	userRoleRepo  userRoleDomain.Repo
	userRoleQuery userRoleQueries.UserRoleQuery
}

func NewService(
	userRoleRepo userRoleDomain.Repo,
	userRoleQuery userRoleQueries.UserRoleQuery,
) *Service {
	return &Service{
		userRoleRepo:  userRoleRepo,
		userRoleQuery: userRoleQuery,
	}
}

