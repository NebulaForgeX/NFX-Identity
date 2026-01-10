package members

import (
	memberDomain "nfxid/modules/tenants/domain/members"
)

type Service struct {
	memberRepo *memberDomain.Repo
}

func NewService(
	memberRepo *memberDomain.Repo,
) *Service {
	return &Service{
		memberRepo: memberRepo,
	}
}
