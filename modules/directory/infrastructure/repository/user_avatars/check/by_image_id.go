package check

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByImageID 根据 ImageID 检查 UserAvatar 是否存在，实现 user_avatars.Check 接口
func (h *Handler) ByImageID(ctx context.Context, imageID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.UserAvatar{}).
		Where("image_id = ?", imageID).
		Count(&count).Error
	return count > 0, err
}
