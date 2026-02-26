package update

import (
	"context"
	"nfxidentity/modules/image/domain/images"
	"nfxidentity/modules/image/infrastructure/rdb/models"
	"nfxidentity/modules/image/infrastructure/repository/images/mapper"
)

// Generic 通用更新 Image，实现 images.Update 接口
func (h *Handler) Generic(ctx context.Context, i *images.Image) error {
	m := mapper.ImageDomainToModel(i)
	updates := mapper.ImageModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Image{}).
		Where("id = ?", i.ID()).
		Updates(updates).Error
}
