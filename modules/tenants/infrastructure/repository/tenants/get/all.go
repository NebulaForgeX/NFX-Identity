package get

import (
	"context"
	"nfxid/modules/tenants/domain/tenants"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenants/mapper"
)

// All 获取所有 Tenant，实现 tenants.Get 接口
func (h *Handler) All(ctx context.Context) ([]*tenants.Tenant, error) {
	var ms []models.Tenant
	if err := h.db.WithContext(ctx).Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*tenants.Tenant, len(ms))
	for i, m := range ms {
		result[i] = mapper.TenantModelToDomain(&m)
	}
	return result, nil
}
