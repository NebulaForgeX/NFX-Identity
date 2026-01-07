package update

import (
	"context"
	profileBadge "nfxid/modules/auth/domain/profile_badge"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// Generic 通用更新 ProfileBadge，实现 profileBadge.Update 接口
func (h *Handler) Generic(ctx context.Context, pb *profileBadge.ProfileBadge) error {
	m := mapper.ProfileBadgeDomainToModel(pb)
	updates := mapper.ProfileBadgeModelsToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.ProfileBadge{}).
		Where("id = ?", pb.ID()).
		Updates(updates).Error
}
