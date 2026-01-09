package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/user_badges"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_badges/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByUserIDAndBadgeID 根据 UserID 和 BadgeID 获取 UserBadge，实现 user_badges.Get 接口
func (h *Handler) ByUserIDAndBadgeID(ctx context.Context, userID, badgeID uuid.UUID) (*user_badges.UserBadge, error) {
	var m models.UserBadge
	if err := h.db.WithContext(ctx).
		Where("user_id = ? AND badge_id = ?", userID, badgeID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_badges.ErrUserBadgeNotFound
		}
		return nil, err
	}
	return mapper.UserBadgeModelToDomain(&m), nil
}
