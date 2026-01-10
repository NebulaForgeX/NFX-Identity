package tenant_apps

import (
	tenantAppDomain "nfxid/modules/tenants/domain/tenant_apps"
)

type Service struct {
	tenantAppRepo *tenantAppDomain.Repo
}

func NewService(
	tenantAppRepo *tenantAppDomain.Repo,
) *Service {
	return &Service{
		tenantAppRepo: tenantAppRepo,
	}
}
