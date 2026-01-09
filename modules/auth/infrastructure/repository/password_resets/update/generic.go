package update

import (
	"context"
	"nfxid/modules/auth/domain/password_resets"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/password_resets/mapper"
)

// Generic 通用更新 PasswordReset，实现 password_resets.Update 接口
func (h *Handler) Generic(ctx context.Context, pr *password_resets.PasswordReset) error {
	m := mapper.PasswordResetDomainToModel(pr)
	updates := mapper.PasswordResetModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.PasswordReset{}).
		Where("id = ?", pr.ID()).
		Updates(updates).Error
}
