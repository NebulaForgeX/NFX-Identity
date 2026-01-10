package refresh_tokens

import (
	refreshTokenDomain "nfxid/modules/auth/domain/refresh_tokens"
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
