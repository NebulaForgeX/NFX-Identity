package create

import (
	"context"
	"nfxidentity/modules/auth/domain/user_credentials"
	"nfxidentity/modules/auth/infrastructure/repository/user_credentials/mapper"
)

// New 创建新的 UserCredential，实现 user_credentials.Create 接口
func (h *Handler) New(ctx context.Context, uc *user_credentials.UserCredential) error {
	m := mapper.UserCredentialDomainToModel(uc)
	return h.db.WithContext(ctx).Create(&m).Error
}
