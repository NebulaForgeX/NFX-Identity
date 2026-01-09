package get

import (
	"context"
	"nfxid/modules/clients/domain/client_credentials"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/client_credentials/mapper"

	"github.com/google/uuid"
)

// ByAppID 根据 AppID 获取 ClientCredential 列表，实现 client_credentials.Get 接口
func (h *Handler) ByAppID(ctx context.Context, appID uuid.UUID) ([]*client_credentials.ClientCredential, error) {
	var ms []models.ClientCredential
	if err := h.db.WithContext(ctx).
		Where("app_id = ?", appID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*client_credentials.ClientCredential, len(ms))
	for i, m := range ms {
		result[i] = mapper.ClientCredentialModelToDomain(&m)
	}
	return result, nil
}
