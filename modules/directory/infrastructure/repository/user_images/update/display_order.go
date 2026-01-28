package update

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// DisplayOrder 更新 UserImage 的 DisplayOrder，实现 user_images.Update 接口
func (h *Handler) DisplayOrder(ctx context.Context, id uuid.UUID, displayOrder int) error {
	return h.db.WithContext(ctx).
		Model(&models.UserImage{}).
		Where("id = ?", id).
		Update("display_order", displayOrder).Error
}
