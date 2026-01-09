package create

import (
	"context"
	"nfxid/modules/clients/domain/client_credentials"
	"nfxid/modules/clients/infrastructure/repository/client_credentials/mapper"
)

// New 创建新的 ClientCredential，实现 client_credentials.Create 接口
func (h *Handler) New(ctx context.Context, cc *client_credentials.ClientCredential) error {
	m := mapper.ClientCredentialDomainToModel(cc)
	return h.db.WithContext(ctx).Create(&m).Error
}
