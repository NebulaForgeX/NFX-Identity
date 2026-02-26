package create

import (
	"context"
	"nfxidentity/modules/auth/domain/password_resets"
	"nfxidentity/modules/auth/infrastructure/repository/password_resets/mapper"
)

// New 创建新的 PasswordReset，实现 password_resets.Create 接口
func (h *Handler) New(ctx context.Context, pr *password_resets.PasswordReset) error {
	m := mapper.PasswordResetDomainToModel(pr)
	return h.db.WithContext(ctx).Create(&m).Error
}
