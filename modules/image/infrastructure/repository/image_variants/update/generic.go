package update

import (
	"context"
	"nfxid/modules/image/domain/image_variants"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/image_variants/mapper"
)

// Generic 通用更新 ImageVariant，实现 image_variants.Update 接口
func (h *Handler) Generic(ctx context.Context, iv *image_variants.ImageVariant) error {
	m := mapper.ImageVariantDomainToModel(iv)
	updates := mapper.ImageVariantModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.ImageVariant{}).
		Where("id = ?", iv.ID()).
		Updates(updates).Error
}
