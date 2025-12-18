package role

import (
	roleDomain "nfxid/modules/auth/domain/role"
)

type Service struct {
	roleRepo  *roleDomain.Repo
	roleQuery *roleDomain.Query
}

func NewService(
	roleRepo *roleDomain.Repo,
	roleQuery *roleDomain.Query,
) *Service {
	return &Service{
		roleRepo:  roleRepo,
		roleQuery: roleQuery,
	}
}
