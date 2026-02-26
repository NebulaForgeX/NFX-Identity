package get

import (
	"context"
	"errors"

	dirErr "nfxidentity/errors/src/directory"
	"nfxidentity/modules/directory/domain/user_profiles"
	"nfxidentity/modules/directory/infrastructure/rdb/models"
	"nfxidentity/modules/directory/infrastructure/repository/user_profiles/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByUserID 根据 UserID 获取 UserProfile，实现 user_profiles.Get 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) (*user_profiles.UserProfile, error) {
	var m models.UserProfile
	if err := h.db.WithContext(ctx).Where("id = ?", userID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dirErr.ErrUserProfileNotFound
		}
		return nil, err
	}
	return mapper.UserProfileModelToDomain(&m), nil
}
