package roles

import (
	roleDomain "nfxid/modules/access/domain/roles"
)

type Service struct {
	roleRepo *roleDomain.Repo
}

func NewService(
	roleRepo *roleDomain.Repo,
) *Service {
	return &Service{
		roleRepo: roleRepo,
	}
}
