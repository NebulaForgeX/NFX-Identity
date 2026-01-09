package update

import (
	"context"
	"nfxid/modules/image/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// UpdateURL 更新 ImageVariant 的 URL，实现 image_variants.Update 接口
func (h *Handler) UpdateURL(ctx context.Context, id uuid.UUID, url string) error {
	return h.db.WithContext(ctx).
		Model(&models.ImageVariant{}).
		Where("id = ?", id).
		Update(models.ImageVariantCols.URL, url).Error
}
