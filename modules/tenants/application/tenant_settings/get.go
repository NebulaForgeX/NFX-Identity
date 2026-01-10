package tenant_settings

import (
	"context"
	tenantSettingResult "nfxid/modules/tenants/application/tenant_settings/results"

	"github.com/google/uuid"
)

// GetTenantSetting 根据ID获取租户设置
func (s *Service) GetTenantSetting(ctx context.Context, tenantSettingID uuid.UUID) (tenantSettingResult.TenantSettingRO, error) {
	domainEntity, err := s.tenantSettingRepo.Get.ByID(ctx, tenantSettingID)
	if err != nil {
		return tenantSettingResult.TenantSettingRO{}, err
	}
	return tenantSettingResult.TenantSettingMapper(domainEntity), nil
}

// GetTenantSettingByTenantID 根据TenantID获取租户设置
func (s *Service) GetTenantSettingByTenantID(ctx context.Context, tenantID uuid.UUID) (tenantSettingResult.TenantSettingRO, error) {
	domainEntity, err := s.tenantSettingRepo.Get.ByTenantID(ctx, tenantID)
	if err != nil {
		return tenantSettingResult.TenantSettingRO{}, err
	}
	return tenantSettingResult.TenantSettingMapper(domainEntity), nil
}
