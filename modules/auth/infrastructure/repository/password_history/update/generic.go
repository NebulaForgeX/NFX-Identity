package update

import (
	"context"
	"nfxid/modules/auth/domain/password_history"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/password_history/mapper"
)

// Generic 通用更新 PasswordHistory，实现 password_history.Update 接口
func (h *Handler) Generic(ctx context.Context, ph *password_history.PasswordHistory) error {
	m := mapper.PasswordHistoryDomainToModel(ph)
	updates := mapper.PasswordHistoryModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.PasswordHistory{}).
		Where("id = ?", ph.ID()).
		Updates(updates).Error
}
