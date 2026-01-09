package get

import (
	"context"
	"nfxid/modules/clients/domain/api_keys"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/api_keys/mapper"

	"github.com/google/uuid"
)

// ByAppID 根据 AppID 获取 APIKey 列表，实现 api_keys.Get 接口
func (h *Handler) ByAppID(ctx context.Context, appID uuid.UUID) ([]*api_keys.APIKey, error) {
	var ms []models.ApiKey
	if err := h.db.WithContext(ctx).
		Where("app_id = ?", appID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*api_keys.APIKey, len(ms))
	for i, m := range ms {
		result[i] = mapper.APIKeyModelToDomain(&m)
	}
	return result, nil
}
