package check

import (
	"context"
	"nfxid/modules/clients/infrastructure/rdb/models"
)

// ByKeyID 根据 KeyID 检查 APIKey 是否存在，实现 api_keys.Check 接口
func (h *Handler) ByKeyID(ctx context.Context, keyID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.ApiKey{}).
		Where("key_id = ?", keyID).
		Count(&count).Error
	return count > 0, err
}
