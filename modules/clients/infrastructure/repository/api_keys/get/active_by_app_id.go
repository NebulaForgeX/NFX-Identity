package get

import (
	"context"
	"nfxid/enums"
	"nfxid/modules/clients/domain/api_keys"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/api_keys/mapper"

	"github.com/google/uuid"
)

// ActiveByAppID 根据 AppID 获取活跃的 APIKey 列表，实现 api_keys.Get 接口
func (h *Handler) ActiveByAppID(ctx context.Context, appID uuid.UUID) ([]*api_keys.APIKey, error) {
	var ms []models.ApiKey
	if err := h.db.WithContext(ctx).
		Where("application_id = ? AND status = ?", appID, enums.ClientsApiKeyStatusActive).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*api_keys.APIKey, len(ms))
	for i, m := range ms {
		result[i] = mapper.APIKeyModelToDomain(&m)
	}
	return result, nil
}
