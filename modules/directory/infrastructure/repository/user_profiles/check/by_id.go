package check

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 检查 UserProfile 是否存在，实现 user_profiles.Check 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.UserProfile{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}
