package update

import (
	"context"
	"errors"
	"time"
	"nfxid/enums"
	"nfxid/modules/clients/domain/api_keys"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Revoke 撤销 APIKey，实现 api_keys.Update 接口
func (h *Handler) Revoke(ctx context.Context, keyID string, revokedBy uuid.UUID, reason string) error {
	// 先检查 APIKey 是否存在
	var m models.ApiKey
	if err := h.db.WithContext(ctx).
		Where("key_id = ?", keyID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return api_keys.ErrAPIKeyNotFound
		}
		return err
	}

	// 检查是否已经撤销
	if m.Status == enums.ClientsApiKeyStatusRevoked {
		return api_keys.ErrAPIKeyAlreadyRevoked
	}

	now := time.Now().UTC()
	updates := map[string]any{
		models.ApiKeyCols.Status:      enums.ClientsApiKeyStatusRevoked,
		models.ApiKeyCols.RevokedAt:   &now,
		models.ApiKeyCols.RevokedBy:   &revokedBy,
		models.ApiKeyCols.RevokeReason: &reason,
	}

	return h.db.WithContext(ctx).
		Model(&models.ApiKey{}).
		Where("key_id = ?", keyID).
		Updates(updates).Error
}
