package update

import (
	"context"
	"nfxid/modules/auth/domain/user"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/mapper"
)

// Generic 通用更新 User，实现 user.Update 接口
func (h *Handler) Generic(ctx context.Context, u *user.User) error {
	m := mapper.UserDomainToModel(u)
	updates := mapper.UserModelsToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", u.ID()).
		Updates(updates).Error
}
