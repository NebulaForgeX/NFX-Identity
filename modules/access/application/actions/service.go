package actions

import (
	actionDomain "nfxid/modules/access/domain/actions"
)

type Service struct {
	actionRepo *actionDomain.Repo
}

func NewService(actionRepo *actionDomain.Repo) *Service {
	return &Service{
		actionRepo: actionRepo,
	}
}
