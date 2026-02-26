package sessions

import (
	sessionDomain "nfxidentity/modules/auth/domain/sessions"
)

type Service struct {
	sessionRepo *sessionDomain.Repo
}

func NewService(
	sessionRepo *sessionDomain.Repo,
) *Service {
	return &Service{
		sessionRepo: sessionRepo,
	}
}
