package update

import (
	"context"
	"nfxid/modules/auth/domain/badge"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// Generic 通用更新 Badge，实现 badge.Update 接口
func (h *Handler) Generic(ctx context.Context, b *badge.Badge) error {
	m := mapper.BadgeDomainToModel(b)
	updates := mapper.BadgeModelsToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Badge{}).
		Where("id = ?", b.ID()).
		Updates(updates).Error
}
