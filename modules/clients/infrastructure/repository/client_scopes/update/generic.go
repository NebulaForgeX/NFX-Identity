package update

import (
	"context"
	"nfxid/modules/clients/domain/client_scopes"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"nfxid/modules/clients/infrastructure/repository/client_scopes/mapper"
)

// Generic 通用更新 ClientScope，实现 client_scopes.Update 接口
func (h *Handler) Generic(ctx context.Context, cs *client_scopes.ClientScope) error {
	m := mapper.ClientScopeDomainToModel(cs)
	updates := mapper.ClientScopeModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.ClientScope{}).
		Where("id = ?", cs.ID()).
		Updates(updates).Error
}
