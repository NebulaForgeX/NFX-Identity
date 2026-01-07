package get

import (
	"context"
	"errors"
	"nfxid/modules/auth/domain/profile"
	profileDomainErrors "nfxid/modules/auth/domain/profile/errors"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByUserID 根据 UserID 获取 Profile，实现 profile.Get 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) (*profile.Profile, error) {
	var m models.Profile
	if err := h.db.WithContext(ctx).Where("user_id = ?", userID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, profileDomainErrors.ErrProfileNotFound
		}
		return nil, err
	}
	return mapper.ProfileModelToDomain(&m), nil
}
