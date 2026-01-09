package get

import (
	"context"
	"nfxid/modules/tenants/domain/tenant_apps"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenant_apps/mapper"

	"github.com/google/uuid"
)

// ByTenantID 根据 TenantID 获取 TenantApp 列表，实现 tenant_apps.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*tenant_apps.TenantApp, error) {
	var ms []models.TenantApp
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*tenant_apps.TenantApp, len(ms))
	for i, m := range ms {
		result[i] = mapper.TenantAppModelToDomain(&m)
	}
	return result, nil
}
