package update

import (
	"context"
	"nfxidentity/modules/directory/domain/users"
	"nfxidentity/modules/directory/infrastructure/rdb/models"
	"nfxidentity/modules/directory/infrastructure/repository/users/mapper"
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
