package update

import (
	"context"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// UpdateURL 更新 Image 的 URL，实现 images.Update 接口
func (h *Handler) UpdateURL(ctx context.Context, id uuid.UUID, url string) error {
	return h.db.WithContext(ctx).
		Model(&models.Image{}).
		Where("id = ?", id).
		Update(models.ImageCols.URL, url).Error
}
