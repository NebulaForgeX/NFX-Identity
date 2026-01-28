package check

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserIDAndImageID 根据 UserID 和 ImageID 检查 UserImage 是否存在，实现 user_images.Check 接口
func (h *Handler) ByUserIDAndImageID(ctx context.Context, userID, imageID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.UserImage{}).
		Where("user_id = ? AND image_id = ? AND deleted_at IS NULL", userID, imageID).
		Count(&count).Error
	return count > 0, err
}
