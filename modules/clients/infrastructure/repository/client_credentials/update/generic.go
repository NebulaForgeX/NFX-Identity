package update

import (
	"context"
	"nfxid/modules/clients/domain/client_credentials"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/client_credentials/mapper"
)

// Generic 通用更新 ClientCredential，实现 client_credentials.Update 接口
func (h *Handler) Generic(ctx context.Context, cc *client_credentials.ClientCredential) error {
	m := mapper.ClientCredentialDomainToModel(cc)
	updates := mapper.ClientCredentialModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.ClientCredential{}).
		Where("id = ?", cc.ID()).
		Updates(updates).Error
}
