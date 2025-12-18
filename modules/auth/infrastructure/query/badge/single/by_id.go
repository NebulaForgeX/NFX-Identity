package single

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

// ByID 根据 ID 获取 Badge，实现 badgeDomain.Single 接口
func (h *Handler) ByID(ctx context.Context, badgeID uuid.UUID) (*badgeDomainViews.BadgeView, error) {
	var m models.Badge
	if err := h.db.WithContext(ctx).Where("id = ?", badgeID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, badgeDomainErrors.ErrBadgeViewNotFound
		}
		return nil, err
	}
	view := mapper.BadgeDBToDomain(&m)
	return &view, nil
}

