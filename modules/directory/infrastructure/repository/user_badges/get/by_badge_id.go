package get

import (
	"context"
	"nfxid/modules/directory/domain/user_badges"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_badges/mapper"

	"github.com/google/uuid"
)

// ByBadgeID 根据 BadgeID 获取 UserBadge 列表，实现 user_badges.Get 接口
func (h *Handler) ByBadgeID(ctx context.Context, badgeID uuid.UUID) ([]*user_badges.UserBadge, error) {
	var ms []models.UserBadge
	if err := h.db.WithContext(ctx).
		Where("badge_id = ?", badgeID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*user_badges.UserBadge, len(ms))
	for i, m := range ms {
		result[i] = mapper.UserBadgeModelToDomain(&m)
	}
	return result, nil
}
