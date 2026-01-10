package groups

import (
	groupDomain "nfxid/modules/tenants/domain/groups"
)

type Service struct {
	groupRepo *groupDomain.Repo
}

func NewService(
	groupRepo *groupDomain.Repo,
) *Service {
	return &Service{
		groupRepo: groupRepo,
	}
}
