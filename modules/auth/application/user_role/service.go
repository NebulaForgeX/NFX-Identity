package user_role

import (
	userRoleDomain "nfxid/modules/auth/domain/user_role"
)

type Service struct {
	userRoleRepo  *userRoleDomain.Repo
	userRoleQuery userRoleDomain.Query
}

func NewService(
	userRoleRepo *userRoleDomain.Repo,
	userRoleQuery userRoleDomain.Query,
) *Service {
	return &Service{
		userRoleRepo:  userRoleRepo,
		userRoleQuery: userRoleQuery,
	}
}
