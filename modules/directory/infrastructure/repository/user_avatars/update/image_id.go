package update

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ImageID 更新 UserAvatar 的 ImageID，实现 user_avatars.Update 接口
func (h *Handler) ImageID(ctx context.Context, userID, imageID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Model(&models.UserAvatar{}).
		Where("user_id = ?", userID).
		Update("image_id", imageID).Error
}
