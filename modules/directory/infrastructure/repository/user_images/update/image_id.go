package update

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ImageID 更新 UserImage 的 ImageID，实现 user_images.Update 接口
func (h *Handler) ImageID(ctx context.Context, id, imageID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Model(&models.UserImage{}).
		Where("id = ?", id).
		Update("image_id", imageID).Error
}
