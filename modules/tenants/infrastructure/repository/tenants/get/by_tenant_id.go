package get

import (
	"context"
	"errors"

	tenantsErr "nfxidentity/errors/src/tenants"
	"nfxidentity/modules/tenants/domain/tenants"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/tenants/mapper"

	"gorm.io/gorm"
)

// ByTenantID 根据 TenantID 获取 Tenant，实现 tenants.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID string) (*tenants.Tenant, error) {
	var m models.Tenant
	if err := h.db.WithContext(ctx).Where("tenant_id = ?", tenantID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, tenantsErr.ErrTenantNotFound
		}
		return nil, err
	}
	return mapper.TenantModelToDomain(&m), nil
}
