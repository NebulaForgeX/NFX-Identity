package update

import (
	"context"
	"nfxid/modules/image/domain/image_types"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/image_types/mapper"
)

// Generic 通用更新 ImageType，实现 image_types.Update 接口
func (h *Handler) Generic(ctx context.Context, it *image_types.ImageType) error {
	m := mapper.ImageTypeDomainToModel(it)
	updates := mapper.ImageTypeModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.ImageType{}).
		Where("id = ?", it.ID()).
		Updates(updates).Error
}
