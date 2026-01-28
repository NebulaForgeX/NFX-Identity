package delete

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 删除 UserAvatar，实现 user_avatars.Delete 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&models.UserAvatar{}).Error
}
