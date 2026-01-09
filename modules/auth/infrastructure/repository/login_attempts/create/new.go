package create

import (
	"context"
	"nfxid/modules/auth/domain/login_attempts"
	"nfxid/modules/auth/infrastructure/repository/login_attempts/mapper"
)

// New 创建新的 LoginAttempt，实现 login_attempts.Create 接口
func (h *Handler) New(ctx context.Context, la *login_attempts.LoginAttempt) error {
	m := mapper.LoginAttemptDomainToModel(la)
	return h.db.WithContext(ctx).Create(&m).Error
}
