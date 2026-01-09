package get

import (
	"context"
	"nfxid/modules/tenants/domain/groups"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/groups/mapper"

	"github.com/google/uuid"
)

// ByTenantID 根据 TenantID 获取 Group 列表，实现 groups.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*groups.Group, error) {
	var ms []models.Group
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*groups.Group, len(ms))
	for i, m := range ms {
		result[i] = mapper.GroupModelToDomain(&m)
	}
	return result, nil
}
