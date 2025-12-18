package single

import (
	"context"
	"errors"
	badgeDomainErrors "nfxid/modules/auth/domain/badge/errors"
	badgeDomainViews "nfxid/modules/auth/domain/badge/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"gorm.io/gorm"
)

// ByName 根据名称获取 Badge，实现 badgeDomain.Single 接口
func (h *Handler) ByName(ctx context.Context, name string) (*badgeDomainViews.BadgeView, error) {
	var m models.Badge
	if err := h.db.WithContext(ctx).Where("name = ?", name).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, badgeDomainErrors.ErrBadgeViewNotFound
		}
		return nil, err
	}
	view := mapper.BadgeDBToDomain(&m)
	return &view, nil
}

