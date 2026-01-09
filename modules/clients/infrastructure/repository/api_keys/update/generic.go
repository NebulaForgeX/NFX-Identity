package update

import (
	"context"
	"nfxid/modules/clients/domain/api_keys"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/api_keys/mapper"
)

// Generic 通用更新 APIKey，实现 api_keys.Update 接口
func (h *Handler) Generic(ctx context.Context, ak *api_keys.APIKey) error {
	m := mapper.APIKeyDomainToModel(ak)
	updates := mapper.APIKeyModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.ApiKey{}).
		Where("id = ?", ak.ID()).
		Updates(updates).Error
}
