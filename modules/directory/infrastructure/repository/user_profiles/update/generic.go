package update

import (
	"context"
	"nfxidentity/modules/directory/domain/user_profiles"
	"nfxidentity/modules/directory/infrastructure/rdb/models"
	"nfxidentity/modules/directory/infrastructure/repository/user_profiles/mapper"
)

// Generic 通用更新 UserProfile，实现 user_profiles.Update 接口
func (h *Handler) Generic(ctx context.Context, up *user_profiles.UserProfile) error {
	m := mapper.UserProfileDomainToModel(up)
	updates := mapper.UserProfileModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.UserProfile{}).
		Where("id = ?", up.ID()).
		Updates(updates).Error
}
