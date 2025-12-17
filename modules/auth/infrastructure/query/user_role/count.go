package user_role

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"
)

// Count 获取 UserRole 总数，实现 user_role.Query 接口
func (h *Handler) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := h.db.WithContext(ctx).Model(&models.UserRole{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
