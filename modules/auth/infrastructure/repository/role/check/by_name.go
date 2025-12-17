package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// ByName 根据 Name 检查 Role 是否存在，实现 role.Check 接口
func (h *Handler) ByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Role{}).
		Where("name = ?", name).
		Count(&count).Error
	return count > 0, err
}
