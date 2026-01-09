package update

import (
	"context"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// UpdatePublic 更新 Image 的 IsPublic，实现 images.Update 接口
func (h *Handler) UpdatePublic(ctx context.Context, id uuid.UUID, isPublic bool) error {
	return h.db.WithContext(ctx).
		Model(&models.Image{}).
		Where("id = ?", id).
		Update(models.ImageCols.IsPublic, isPublic).Error
}
