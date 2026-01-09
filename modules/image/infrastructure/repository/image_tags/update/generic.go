package update

import (
	"context"
	"nfxid/modules/image/domain/image_tags"
	"nfxid/modules/image/infrastructure/rdb/models"
	"nfxid/modules/image/infrastructure/repository/image_tags/mapper"
)

// Generic 通用更新 ImageTag，实现 image_tags.Update 接口
func (h *Handler) Generic(ctx context.Context, it *image_tags.ImageTag) error {
	m := mapper.ImageTagDomainToModel(it)
	updates := mapper.ImageTagModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.ImageTag{}).
		Where("id = ?", it.ID()).
		Updates(updates).Error
}
