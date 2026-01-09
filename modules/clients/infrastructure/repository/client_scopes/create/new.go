package create

import (
	"context"
	"nfxid/modules/clients/domain/client_scopes"
	"nfxid/modules/clients/infrastructure/repository/client_scopes/mapper"
)

// New 创建新的 ClientScope，实现 client_scopes.Create 接口
func (h *Handler) New(ctx context.Context, cs *client_scopes.ClientScope) error {
	m := mapper.ClientScopeDomainToModel(cs)
	return h.db.WithContext(ctx).Create(&m).Error
}
