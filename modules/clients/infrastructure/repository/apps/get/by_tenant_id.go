package get

import (
	"context"
	"nfxid/modules/clients/domain/apps"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/apps/mapper"

	"github.com/google/uuid"
)

// ByTenantID 根据 TenantID 获取 App 列表，实现 apps.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*apps.App, error) {
	var ms []models.App
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*apps.App, len(ms))
	for i, m := range ms {
		result[i] = mapper.AppModelToDomain(&m)
	}
	return result, nil
}
