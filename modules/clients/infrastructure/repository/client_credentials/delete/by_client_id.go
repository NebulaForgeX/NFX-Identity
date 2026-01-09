package delete

import (
	"context"
	"errors"
	"nfxid/modules/clients/domain/client_credentials"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"gorm.io/gorm"
)

// ByClientID 根据 ClientID 删除 ClientCredential，实现 client_credentials.Delete 接口
func (h *Handler) ByClientID(ctx context.Context, clientID string) error {
	result := h.db.WithContext(ctx).
		Where("client_id = ?", clientID).
		Delete(&models.ClientCredential{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return client_credentials.ErrClientCredentialNotFound
	}
	return nil
}
