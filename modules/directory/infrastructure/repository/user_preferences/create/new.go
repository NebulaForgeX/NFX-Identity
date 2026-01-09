package create

import (
	"context"
	"nfxid/modules/directory/domain/user_preferences"
	"nfxid/modules/directory/infrastructure/repository/user_preferences/mapper"
)

// New 创建新的 UserPreference，实现 user_preferences.Create 接口
func (h *Handler) New(ctx context.Context, up *user_preferences.UserPreference) error {
	m := mapper.UserPreferenceDomainToModel(up)
	return h.db.WithContext(ctx).Create(&m).Error
}
