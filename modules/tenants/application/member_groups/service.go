package member_groups

import (
	memberGroupDomain "nfxidentity/modules/tenants/domain/member_groups"
)

type Service struct {
	memberGroupRepo *memberGroupDomain.Repo
}

func NewService(
	memberGroupRepo *memberGroupDomain.Repo,
) *Service {
	return &Service{
		memberGroupRepo: memberGroupRepo,
	}
}
