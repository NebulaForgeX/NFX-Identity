package get

import (
	"context"
	"nfxid/modules/clients/domain/apps"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/apps/mapper"

	"github.com/google/uuid"
)

// ByTenantIDAndEnvironment 根据 TenantID 和 Environment 获取 App 列表，实现 apps.Get 接口
func (h *Handler) ByTenantIDAndEnvironment(ctx context.Context, tenantID uuid.UUID, environment apps.Environment) ([]*apps.App, error) {
	envEnum := mapper.EnvironmentDomainToEnum(environment)
	var ms []models.Application
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ? AND environment = ?", tenantID, envEnum).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*apps.App, len(ms))
	for i, m := range ms {
		result[i] = mapper.AppModelToDomain(&m)
	}
	return result, nil
}
