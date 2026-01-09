package check

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserIDAndBadgeID 根据 UserID 和 BadgeID 检查 UserBadge 是否存在，实现 user_badges.Check 接口
func (h *Handler) ByUserIDAndBadgeID(ctx context.Context, userID, badgeID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.UserBadge{}).
		Where("user_id = ? AND badge_id = ?", userID, badgeID).
		Count(&count).Error
	return count > 0, err
}
