package update

import (
	"context"
	imageDomain "nfxid/modules/image/domain/image"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/mapper"
)

// Generic 更新 Image，实现 imageDomain.Update 接口
func (h *Handler) Generic(ctx context.Context, img *imageDomain.Image) error {
	m := mapper.ImageDomainToModel(img)
	updates := mapper.ImageModelsToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Image{}).
		Where("id = ?", img.ID()).
		Updates(updates).Error
}
