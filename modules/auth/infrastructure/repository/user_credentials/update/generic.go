package update

import (
	"context"
	"nfxid/modules/auth/domain/user_credentials"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/repository/user_credentials/mapper"
)

// Generic 通用更新 UserCredential，实现 user_credentials.Update 接口
func (h *Handler) Generic(ctx context.Context, uc *user_credentials.UserCredential) error {
	m := mapper.UserCredentialDomainToModel(uc)
	updates := mapper.UserCredentialModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.UserCredential{}).
		Where("id = ?", uc.ID()).
		Updates(updates).Error
}
