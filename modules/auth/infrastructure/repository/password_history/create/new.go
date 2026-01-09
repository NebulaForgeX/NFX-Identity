package create

import (
	"context"
	"nfxid/modules/auth/domain/password_history"
	"nfxid/modules/auth/infrastructure/repository/password_history/mapper"
)

// New 创建新的 PasswordHistory，实现 password_history.Create 接口
func (h *Handler) New(ctx context.Context, ph *password_history.PasswordHistory) error {
	m := mapper.PasswordHistoryDomainToModel(ph)
	return h.db.WithContext(ctx).Create(&m).Error
}
