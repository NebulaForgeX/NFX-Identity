package update

import (
	"context"
	"nfxidentity/modules/clients/infrastructure/rdb/models"
	"time"
)

// UpdateLastUsed 更新最后使用时间，实现 client_credentials.Update 接口
func (h *Handler) UpdateLastUsed(ctx context.Context, clientID string) error {
	now := time.Now().UTC()
	updates := map[string]any{
		models.ClientCredentialCols.LastUsedAt: &now,
	}

	return h.db.WithContext(ctx).
		Model(&models.ClientCredential{}).
		Where("client_id = ?", clientID).
		Updates(updates).Error
}
