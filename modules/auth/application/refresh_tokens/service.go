package refresh_tokens

import (
	refreshTokenDomain "nfxidentity/modules/auth/domain/refresh_tokens"
)

type Service struct {
	refreshTokenRepo *refreshTokenDomain.Repo
}

func NewService(
	refreshTokenRepo *refreshTokenDomain.Repo,
) *Service {
	return &Service{
		refreshTokenRepo: refreshTokenRepo,
	}
}
