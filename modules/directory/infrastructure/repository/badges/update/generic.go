package update

import (
	"context"
	"nfxid/modules/directory/domain/badges"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/badges/mapper"
)

// Generic 通用更新 Badge，实现 badges.Update 接口
func (h *Handler) Generic(ctx context.Context, b *badges.Badge) error {
	m := mapper.BadgeDomainToModel(b)
	updates := mapper.BadgeModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Badge{}).
		Where("id = ?", b.ID()).
		Updates(updates).Error
}
