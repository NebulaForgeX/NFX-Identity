package get

import (
	"context"
	"errors"
	"nfxid/modules/directory/domain/user_preferences"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_preferences/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ByUserID 根据 UserID 获取 UserPreference，实现 user_preferences.Get 接口
func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) (*user_preferences.UserPreference, error) {
	var m models.UserPreference
	if err := h.db.WithContext(ctx).Where("id = ?", userID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user_preferences.ErrUserPreferenceNotFound
		}
		return nil, err
	}
	return mapper.UserPreferenceModelToDomain(&m), nil
}
