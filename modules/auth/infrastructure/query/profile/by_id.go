package profile

import (
	"context"
	"errors"
	profileDomainErrors "nfxid/modules/auth/domain/profile/errors"
	profileDomainViews "nfxid/modules/auth/domain/profile/views"
	"nfxid/modules/auth/infrastructure/query/mapper"
	"nfxid/modules/auth/infrastructure/rdb/views"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByID 根据 ID 获取 Profile，实现 profile.Query 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (profileDomainViews.ProfileView, error) {
	var v views.ProfileCompleteView
	if err := h.db.WithContext(ctx).Where("profile_id = ?", id).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return profileDomainViews.ProfileView{}, profileDomainErrors.ErrProfileViewNotFound
		}
		return profileDomainViews.ProfileView{}, err
	}
	return mapper.ProfileViewToDomain(&v), nil
}
