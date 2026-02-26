package ip_allowlist

import (
	ipAllowlistDomain "nfxidentity/modules/clients/domain/ip_allowlist"
)

type Service struct {
	ipAllowlistRepo *ipAllowlistDomain.Repo
}

func NewService(
	ipAllowlistRepo *ipAllowlistDomain.Repo,
) *Service {
	return &Service{
		ipAllowlistRepo: ipAllowlistRepo,
	}
}
