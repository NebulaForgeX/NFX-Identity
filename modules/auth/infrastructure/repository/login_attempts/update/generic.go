package update

import (
	"context"
	"nfxid/modules/auth/domain/login_attempts"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/login_attempts/mapper"
)

// Generic 通用更新 LoginAttempt，实现 login_attempts.Update 接口
func (h *Handler) Generic(ctx context.Context, la *login_attempts.LoginAttempt) error {
	m := mapper.LoginAttemptDomainToModel(la)
	updates := mapper.LoginAttemptModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.LoginAttempt{}).
		Where("id = ?", la.ID()).
		Updates(updates).Error
}
