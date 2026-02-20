package get

import (
	"context"
	"nfxid/modules/clients/domain/client_scopes"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/client_scopes/mapper"

	"github.com/google/uuid"
)

// ActiveByAppID 根据 AppID 获取活跃的 ClientScope 列表，实现 client_scopes.Get 接口
func (h *Handler) ActiveByAppID(ctx context.Context, appID uuid.UUID) ([]*client_scopes.ClientScope, error) {
	var ms []models.ClientScope
	if err := h.db.WithContext(ctx).
		Where("application_id = ? AND revoked_at IS NULL", appID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*client_scopes.ClientScope, len(ms))
	for i, m := range ms {
		result[i] = mapper.ClientScopeModelToDomain(&m)
	}
	return result, nil
}
