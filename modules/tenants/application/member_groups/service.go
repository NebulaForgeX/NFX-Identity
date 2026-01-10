package member_groups

import (
	memberGroupDomain "nfxid/modules/tenants/domain/member_groups"
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
