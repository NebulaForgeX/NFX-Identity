package get

import (
	"context"
	"errors"
	"nfxid/modules/clients/domain/api_keys"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/api_keys/mapper"

	"gorm.io/gorm"
)

// ByKeyID 根据 KeyID 获取 APIKey，实现 api_keys.Get 接口
func (h *Handler) ByKeyID(ctx context.Context, keyID string) (*api_keys.APIKey, error) {
	var m models.ApiKey
	if err := h.db.WithContext(ctx).Where("key_id = ?", keyID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, api_keys.ErrAPIKeyNotFound
		}
		return nil, err
	}
	return mapper.APIKeyModelToDomain(&m), nil
}
