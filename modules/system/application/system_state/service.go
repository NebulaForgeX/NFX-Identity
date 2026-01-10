package system_state

import (
	systemStateDomain "nfxid/modules/system/domain/system_state"
)

type Service struct {
	systemStateRepo *systemStateDomain.Repo
}

func NewService(
	systemStateRepo *systemStateDomain.Repo,
) *Service {
	return &Service{
		systemStateRepo: systemStateRepo,
	}
}
