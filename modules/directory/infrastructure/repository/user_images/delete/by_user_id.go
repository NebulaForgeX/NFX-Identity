package delete

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserID 根据 UserID 删除 UserImage（软删除），实现 user_images.Delete 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Model(&models.UserImage{}).
		Where("user_id = ?", userID).
		Update("deleted_at", "NOW()").Error
}
