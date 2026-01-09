package update

import (
	"context"
	"nfxid/modules/directory/domain/user_preferences"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_preferences/mapper"
)

// Generic 通用更新 UserPreference，实现 user_preferences.Update 接口
func (h *Handler) Generic(ctx context.Context, up *user_preferences.UserPreference) error {
	m := mapper.UserPreferenceDomainToModel(up)
	updates := mapper.UserPreferenceModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.UserPreference{}).
		Where("id = ?", up.ID()).
		Updates(updates).Error
}
