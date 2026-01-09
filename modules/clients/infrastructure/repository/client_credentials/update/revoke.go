package update

import (
	"context"
	"errors"
	"nfxid/enums"
	"nfxid/modules/clients/domain/client_credentials"
	"nfxid/modules/clients/infrastructure/rdb/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Revoke 撤销 ClientCredential，实现 client_credentials.Update 接口
func (h *Handler) Revoke(ctx context.Context, clientID string, revokedBy uuid.UUID, reason string) error {
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

	// 检查是否已经撤销
	if m.Status == enums.ClientsCredentialStatusRevoked {
		return client_credentials.ErrCredentialAlreadyRevoked
	}

	now := time.Now().UTC()
	updates := map[string]any{
		models.ClientCredentialCols.Status:       enums.ClientsCredentialStatusRevoked,
		models.ClientCredentialCols.RevokedAt:    &now,
		models.ClientCredentialCols.RevokedBy:    &revokedBy,
		models.ClientCredentialCols.RevokeReason: &reason,
	}

	return h.db.WithContext(ctx).
		Model(&models.ClientCredential{}).
		Where("client_id = ?", clientID).
		Updates(updates).Error
}
