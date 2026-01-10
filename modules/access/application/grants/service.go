package grants

import (
	grantDomain "nfxid/modules/access/domain/grants"
)

type Service struct {
	grantRepo *grantDomain.Repo
}

func NewService(
	grantRepo *grantDomain.Repo,
) *Service {
	return &Service{
		grantRepo: grantRepo,
	}
}
