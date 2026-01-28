package update

import (
	"context"
	"nfxid/modules/directory/domain/user_avatars"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/user_avatars/mapper"
)

// Generic 通用更新 UserAvatar，实现 user_avatars.Update 接口
func (h *Handler) Generic(ctx context.Context, ua *user_avatars.UserAvatar) error {
	m := mapper.UserAvatarDomainToModel(ua)
	updates := mapper.UserAvatarModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.UserAvatar{}).
		Where("user_id = ?", ua.UserID()).
		Updates(updates).Error
}
