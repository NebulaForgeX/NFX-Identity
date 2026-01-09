package update

import (
	"context"
	"nfxid/modules/directory/domain/user_badges"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_badges/mapper"
)

// Generic 通用更新 UserBadge，实现 user_badges.Update 接口
func (h *Handler) Generic(ctx context.Context, ub *user_badges.UserBadge) error {
	m := mapper.UserBadgeDomainToModel(ub)
	updates := mapper.UserBadgeModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.UserBadge{}).
		Where("id = ?", ub.ID()).
		Updates(updates).Error
}
