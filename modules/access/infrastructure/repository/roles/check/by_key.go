package check

import (
	"context"
	"nfxid/modules/access/infrastructure/rdb/models"
)

// ByKey 根据 Key 检查 Role 是否存在，实现 roles.Check 接口
func (h *Handler) ByKey(ctx context.Context, key string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Role{}).
		Where("key = ?", key).
		Count(&count).Error
	return count > 0, err
}
