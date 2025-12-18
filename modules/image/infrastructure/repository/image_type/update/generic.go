package update

import (
	"context"
	imageTypeDomain "nfxid/modules/image/domain/image_type"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/mapper"
)

// Generic 更新 ImageType，实现 imageTypeDomain.Update 接口
func (h *Handler) Generic(ctx context.Context, it *imageTypeDomain.ImageType) error {
	m := mapper.ImageTypeDomainToModel(it)
	updates := mapper.ImageTypeModelsToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.ImageType{}).
		Where("id = ?", it.ID()).
		Updates(updates).Error
}
