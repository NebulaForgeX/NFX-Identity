package tenants

import (
	tenantDomain "nfxidentity/modules/tenants/domain/tenants"
)

type Service struct {
	tenantRepo *tenantDomain.Repo
}

func NewService(
	tenantRepo *tenantDomain.Repo,
) *Service {
	return &Service{
		tenantRepo: tenantRepo,
	}
}
