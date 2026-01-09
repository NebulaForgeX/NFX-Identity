package check

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"
)

// ByName 根据 Name 检查 Badge 是否存在，实现 badges.Check 接口
func (h *Handler) ByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Badge{}).
		Where("name = ?", name).
		Count(&count).Error
	return count > 0, err
}
