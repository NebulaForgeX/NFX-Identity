package update

import (
	"context"
	"errors"
	"time"
	"nfxid/enums"
	"nfxid/modules/clients/domain/client_credentials"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"gorm.io/gorm"
)

// Rotate 轮换 ClientCredential，实现 client_credentials.Update 接口
func (h *Handler) Rotate(ctx context.Context, clientID string, newSecretHash, newHashAlg string) error {
	// 先检查 ClientCredential 是否存在
	var m models.ClientCredential
	if err := h.db.WithContext(ctx).
		Where("client_id = ?", clientID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return client_credentials.ErrClientCredentialNotFound
		}
		return err
	}

	now := time.Now().UTC()
	updates := map[string]any{
		models.ClientCredentialCols.SecretHash: newSecretHash,
		models.ClientCredentialCols.HashAlg:    newHashAlg,
		models.ClientCredentialCols.Status:     enums.ClientsCredentialStatusActive,
		models.ClientCredentialCols.RotatedAt:  &now,
	}

	return h.db.WithContext(ctx).
		Model(&models.ClientCredential{}).
		Where("client_id = ?", clientID).
		Updates(updates).Error
}
