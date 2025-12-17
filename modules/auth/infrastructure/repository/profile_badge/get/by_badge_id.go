package get

import (
	"context"
	profileBadge "nfxid/modules/auth/domain/profile_badge"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
)

// ByBadgeID 根据 BadgeID 获取 ProfileBadge 列表，实现 profileBadge.Get 接口
func (h *Handler) ByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]*profileBadge.ProfileBadge, error) {
	var ms []models.ProfileBadge
	if err := h.db.WithContext(ctx).
		Where("badge_id = ?", badgeID).
		Order("earned_at DESC").
		Find(&ms).Error; err != nil {
		return nil, err
	}

	entities := make([]*profileBadge.ProfileBadge, len(ms))
	for i := range ms {
		entities[i] = mapper.ProfileBadgeModelToDomain(&ms[i])
	}
	return entities, nil
}
