package apps

import (
	appDomain "nfxid/modules/clients/domain/apps"
)

type Service struct {
	appRepo *appDomain.Repo
}

func NewService(
	appRepo *appDomain.Repo,
) *Service {
	return &Service{
		appRepo: appRepo,
	}
}
