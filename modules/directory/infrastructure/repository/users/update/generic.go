package update

import (
	"context"
	"nfxid/modules/directory/domain/users"
	"nfxid/modules/directory/infrastructure/rdb/models"
	"nfxid/modules/directory/infrastructure/repository/users/mapper"
)

// Generic 通用更新 User，实现 users.Update 接口
func (h *Handler) Generic(ctx context.Context, u *users.User) error {
	m := mapper.UserDomainToModel(u)
	updates := mapper.UserModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", u.ID()).
		Updates(updates).Error
}
