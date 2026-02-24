package update

import (
	"context"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"time"
)

// UpdateLastUsed 更新最后使用时间，实现 api_keys.Update 接口
func (h *Handler) UpdateLastUsed(ctx context.Context, keyID string) error {
	now := time.Now().UTC()
	updates := map[string]any{
		models.ApiKeyCols.LastUsedAt: &now,
	}

	return h.db.WithContext(ctx).
		Model(&models.ApiKey{}).
		Where("key_id = ?", keyID).
		Updates(updates).Error
}
