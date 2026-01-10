package member_roles

import (
	memberRoleDomain "nfxid/modules/tenants/domain/member_roles"
)

type Service struct {
	memberRoleRepo *memberRoleDomain.Repo
}

func NewService(
	memberRoleRepo *memberRoleDomain.Repo,
) *Service {
	return &Service{
		memberRoleRepo: memberRoleRepo,
	}
}
