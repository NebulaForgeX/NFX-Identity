package tenants

import (
	"context"
	tenantResult "nfxid/modules/tenants/application/tenants/results"

	"github.com/google/uuid"
)

// GetTenant 根据ID获取租户
func (s *Service) GetTenant(ctx context.Context, tenantID uuid.UUID) (tenantResult.TenantRO, error) {
	domainEntity, err := s.tenantRepo.Get.ByID(ctx, tenantID)
	if err != nil {
		return tenantResult.TenantRO{}, err
	}
	return tenantResult.TenantMapper(domainEntity), nil
}

// GetTenantByTenantID 根据TenantID获取租户
func (s *Service) GetTenantByTenantID(ctx context.Context, tenantID string) (tenantResult.TenantRO, error) {
	domainEntity, err := s.tenantRepo.Get.ByTenantID(ctx, tenantID)
	if err != nil {
		return tenantResult.TenantRO{}, err
	}
	return tenantResult.TenantMapper(domainEntity), nil
}
