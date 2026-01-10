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

// GetTenantAppsByTenantID 根据租户ID获取租户应用列表
func (s *Service) GetTenantAppsByTenantID(ctx context.Context, tenantID uuid.UUID) ([]tenantAppResult.TenantAppRO, error) {
	domainEntities, err := s.tenantAppRepo.Get.ByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}
	
	results := make([]tenantAppResult.TenantAppRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = tenantAppResult.TenantAppMapper(entity)
	}
	return results, nil
}

// GetTenantAppsByAppID 根据应用ID获取租户应用列表
func (s *Service) GetTenantAppsByAppID(ctx context.Context, appID uuid.UUID) ([]tenantAppResult.TenantAppRO, error) {
	domainEntities, err := s.tenantAppRepo.Get.ByAppID(ctx, appID)
	if err != nil {
		return nil, err
	}
	
	results := make([]tenantAppResult.TenantAppRO, len(domainEntities))
	for i, entity := range domainEntities {
		results[i] = tenantAppResult.TenantAppMapper(entity)
	}
	return results, nil
}
