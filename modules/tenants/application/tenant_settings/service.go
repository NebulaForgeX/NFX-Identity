package tenant_settings

import (
	tenantSettingDomain "nfxidentity/modules/tenants/domain/tenant_settings"
)

type Service struct {
	tenantSettingRepo *tenantSettingDomain.Repo
}

func NewService(
	tenantSettingRepo *tenantSettingDomain.Repo,
) *Service {
	return &Service{
		tenantSettingRepo: tenantSettingRepo,
	}
}
