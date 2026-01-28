package delete

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 UserImage（软删除），实现 user_images.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	return h.db.WithContext(ctx).
		Model(&models.UserImage{}).
		Where("id = ?", id).
		Update("deleted_at", "NOW()").Error
}
