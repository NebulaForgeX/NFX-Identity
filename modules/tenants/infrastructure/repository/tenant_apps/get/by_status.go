package get

import (
	"context"
	"nfxid/modules/tenants/domain/tenant_apps"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenant_apps/mapper"
)

// ByStatus 根据 Status 获取 TenantApp 列表，实现 tenant_apps.Get 接口
func (h *Handler) ByStatus(ctx context.Context, status tenant_apps.TenantAppStatus) ([]*tenant_apps.TenantApp, error) {
	statusEnum := mapper.TenantAppStatusDomainToEnum(status)
	var ms []models.TenantApplication
	if err := h.db.WithContext(ctx).
		Where("status = ?", statusEnum).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*tenant_apps.TenantApp, len(ms))
	for i, m := range ms {
		result[i] = mapper.TenantAppModelToDomain(&m)
	}
	return result, nil
}
