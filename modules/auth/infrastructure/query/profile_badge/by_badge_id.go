package profile_badge

import (
	"context"
	profileBadgeDomainViews "nfxid/modules/auth/domain/profile_badge/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/utils/slice"

	"github.com/google/uuid"
)

// ByBadgeID 根据 BadgeID 获取 ProfileBadge 列表，实现 profile_badge.Query 接口
func (h *Handler) ByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]profileBadgeDomainViews.ProfileBadgeView, error) {
	var items []models.ProfileBadge
	if err := h.db.WithContext(ctx).
		Where("badge_id = ?", badgeID).
		Order("earned_at DESC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return slice.MapP(items, mapper.ProfileBadgeModelToDomain), nil
}
