package action_requirements

import (
	arDomain "nfxid/modules/access/domain/action_requirements"
)

type Service struct {
	arRepo *arDomain.Repo
}

func NewService(arRepo *arDomain.Repo) *Service {
	return &Service{arRepo: arRepo}
}
