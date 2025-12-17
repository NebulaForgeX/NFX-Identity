package profile_badge

import (
	"context"
	"errors"
	profileBadgeDomainErrors "nfxid/modules/auth/domain/profile_badge/errors"
	profileBadgeDomainViews "nfxid/modules/auth/domain/profile_badge/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 ProfileBadge，实现 profile_badge.Query 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (profileBadgeDomainViews.ProfileBadgeView, error) {
	var m models.ProfileBadge
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return profileBadgeDomainViews.ProfileBadgeView{}, profileBadgeDomainErrors.ErrProfileBadgeViewNotFound
		}
		return profileBadgeDomainViews.ProfileBadgeView{}, err
	}
	return mapper.ProfileBadgeModelToDomain(&m), nil
}
