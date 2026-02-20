package delete

import (
	"context"
	"nfxid/modules/clients/domain/client_scopes"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByAppIDAndScope 根据 AppID 和 Scope 删除 ClientScope，实现 client_scopes.Delete 接口
func (h *Handler) ByAppIDAndScope(ctx context.Context, appID uuid.UUID, scope string) error {
	result := h.db.WithContext(ctx).
		Where("application_id = ? AND scope = ?", appID, scope).
		Delete(&models.ClientScope{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return client_scopes.ErrClientScopeNotFound
	}
	return nil
}
