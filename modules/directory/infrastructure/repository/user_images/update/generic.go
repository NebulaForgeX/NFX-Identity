package update

import (
	"context"
	"nfxid/modules/directory/domain/user_images"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_images/mapper"
)

// Generic 通用更新 UserImage，实现 user_images.Update 接口
func (h *Handler) Generic(ctx context.Context, ui *user_images.UserImage) error {
	m := mapper.UserImageDomainToModel(ui)
	updates := mapper.UserImageModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.UserImage{}).
		Where("id = ?", ui.ID()).
		Updates(updates).Error
}
