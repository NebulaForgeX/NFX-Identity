package check

import (
	"context"
	"nfxid/modules/image/infrastructure/rdb/models"
)

// ByKey 根据 Key 检查 ImageType 是否存在，实现 image_types.Check 接口
func (h *Handler) ByKey(ctx context.Context, key string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.ImageType{}).
		Where("key = ?", key).
		Count(&count).Error
	return count > 0, err
}
