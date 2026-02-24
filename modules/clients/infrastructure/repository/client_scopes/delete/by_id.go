package delete

import (
	"context"

	clientsErr "nfxid/errors/src/clients"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 ClientScope，实现 client_scopes.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.ClientScope{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return clientsErr.ErrClientScopeNotFound
	}
	return nil
}
