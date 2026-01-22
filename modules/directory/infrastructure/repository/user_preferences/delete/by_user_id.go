package delete

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 删除 UserPreference，实现 user_preferences.Delete 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("id = ?", userID).
		Delete(&models.UserPreference{}).Error
}
