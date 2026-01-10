package tenant_apps

import (
	"context"
	tenantAppResult "nfxid/modules/tenants/application/tenant_apps/results"

	"github.com/google/uuid"
)

// GetTenantApp 根据ID获取租户应用
func (s *Service) GetTenantApp(ctx context.Context, tenantAppID uuid.UUID) (tenantAppResult.TenantAppRO, error) {
	domainEntity, err := s.tenantAppRepo.Get.ByID(ctx, tenantAppID)
	if err != nil {
		return tenantAppResult.TenantAppRO{}, err
	}
	return tenantAppResult.TenantAppMapper(domainEntity), nil
}
