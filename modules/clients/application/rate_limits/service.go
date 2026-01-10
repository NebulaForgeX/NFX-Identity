package rate_limits

import (
	rateLimitDomain "nfxid/modules/clients/domain/rate_limits"
)

type Service struct {
	rateLimitRepo *rateLimitDomain.Repo
}

func NewService(
	rateLimitRepo *rateLimitDomain.Repo,
) *Service {
	return &Service{
		rateLimitRepo: rateLimitRepo,
	}
}
