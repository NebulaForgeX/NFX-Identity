package delete

import (
	"context"
	"nfxid/modules/clients/domain/client_credentials"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"github.com/google/uuid"
)

// ByID 根据 ID 删除 ClientCredential，实现 client_credentials.Delete 接口
func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.ClientCredential{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return client_credentials.ErrClientCredentialNotFound
	}
	return nil
}
