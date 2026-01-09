package get

import (
	"context"
	"nfxid/modules/tenants/domain/tenants"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/tenants/mapper"
)

// ByStatus 根据 Status 获取 Tenant 列表，实现 tenants.Get 接口
func (h *Handler) ByStatus(ctx context.Context, status tenants.TenantStatus) ([]*tenants.Tenant, error) {
	statusEnum := mapper.TenantStatusDomainToEnum(status)
	var ms []models.Tenant
	if err := h.db.WithContext(ctx).
		Where("status = ?", statusEnum).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*tenants.Tenant, len(ms))
	for i, m := range ms {
		result[i] = mapper.TenantModelToDomain(&m)
	}
	return result, nil
}
