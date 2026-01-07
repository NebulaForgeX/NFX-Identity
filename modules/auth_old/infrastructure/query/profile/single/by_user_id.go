package single

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

// ByUserID 根据 UserID 获取 Profile，实现 profileDomain.Single 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) (*profileDomainViews.ProfileView, error) {
	var v views.ProfileCompleteView
	if err := h.db.WithContext(ctx).Where("user_id = ?", userID).First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, profileDomainErrors.ErrProfileViewNotFound
		}
		return nil, err
	}
	view := mapper.ProfileViewToDomain(&v)
	return &view, nil
}
