package delete

import (
	"context"
	"nfxid/modules/directory/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByUserIDAndBadgeID 根据 UserID 和 BadgeID 删除 UserBadge，实现 user_badges.Delete 接口
func (h *Handler) ByUserIDAndBadgeID(ctx context.Context, userID, badgeID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("user_id = ? AND badge_id = ?", userID, badgeID).
		Delete(&models.UserBadge{}).Error
}
