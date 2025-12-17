package delete

import (
	"context"
	profileBadgeDomainErrors "nfxid/modules/auth/domain/profile_badge/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByProfileAndBadge 根据 ProfileID 和 BadgeID 删除 ProfileBadge，实现 profileBadge.Delete 接口
func (h *Handler) ByProfileAndBadge(ctx context.Context, profileID, badgeID uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("profile_id = ? AND badge_id = ?", profileID, badgeID).
		Delete(&models.ProfileBadge{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return profileBadgeDomainErrors.ErrProfileBadgeNotFound
	}
	return nil
}
