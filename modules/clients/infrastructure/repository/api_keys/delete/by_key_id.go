package delete

import (
	"context"
	"errors"
	"nfxid/modules/clients/domain/api_keys"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"gorm.io/gorm"
)

// ByKeyID 根据 KeyID 删除 APIKey，实现 api_keys.Delete 接口
func (h *Handler) ByKeyID(ctx context.Context, keyID string) error {
	result := h.db.WithContext(ctx).
		Where("key_id = ?", keyID).
		Delete(&models.ApiKey{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return api_keys.ErrAPIKeyNotFound
	}
	return nil
}
