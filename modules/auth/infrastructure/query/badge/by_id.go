package badge

import (
	"context"
	"errors"
	badgeDomainErrors "nfxid/modules/auth/domain/badge/errors"
	badgeDomainViews "nfxid/modules/auth/domain/badge/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Badge，实现 badge.Query 接口
func (h *Handler) ByID(ctx context.Context, badgeID uuid.UUID) (badgeDomainViews.BadgeView, error) {
	var m models.Badge
	if err := h.db.WithContext(ctx).Where("id = ?", badgeID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return badgeDomainViews.BadgeView{}, badgeDomainErrors.ErrBadgeViewNotFound
		}
		return badgeDomainViews.BadgeView{}, err
	}
	return mapper.BadgeDBToDomain(&m), nil
}
