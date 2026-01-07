package check

import (
	"context"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByProfileAndBadge 根据 ProfileID 和 BadgeID 检查 ProfileBadge 是否存在，实现 profileBadge.Check 接口
func (h *Handler) ByProfileAndBadge(ctx context.Context, profileID, badgeID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.ProfileBadge{}).
		Where("profile_id = ? AND badge_id = ?", profileID, badgeID).
		Count(&count).Error
	return count > 0, err
}
