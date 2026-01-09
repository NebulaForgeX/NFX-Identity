package check

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 检查 UserPreference 是否存在，实现 user_preferences.Check 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.UserPreference{}).
		Where("user_id = ?", userID).
		Count(&count).Error
	return count > 0, err
}
