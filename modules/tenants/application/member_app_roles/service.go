package member_app_roles

import (
	memberAppRoleDomain "nfxid/modules/tenants/domain/member_app_roles"
)

type Service struct {
	memberAppRoleRepo *memberAppRoleDomain.Repo
}

func NewService(
	memberAppRoleRepo *memberAppRoleDomain.Repo,
) *Service {
	return &Service{
		memberAppRoleRepo: memberAppRoleRepo,
	}
}
