package update

import (
	"context"
	"errors"
	"time"
	"nfxid/modules/clients/domain/client_scopes"
	"nfxid/modules/clients/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Revoke 撤销 ClientScope，实现 client_scopes.Update 接口
func (h *Handler) Revoke(ctx context.Context, id uuid.UUID, revokedBy uuid.UUID, reason string) error {
	// 先检查 ClientScope 是否存在
	var m models.ClientScope
	if err := h.db.WithContext(ctx).
		Where("id = ?", id).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return client_scopes.ErrClientScopeNotFound
		}
		return err
	}

	// 检查是否已经撤销
	if m.RevokedAt != nil {
		return client_scopes.ErrClientScopeAlreadyRevoked
	}

	now := time.Now().UTC()
	updates := map[string]any{
		models.ClientScopeCols.RevokedAt:    &now,
		models.ClientScopeCols.RevokedBy:    &revokedBy,
		models.ClientScopeCols.RevokeReason: &reason,
	}

	return h.db.WithContext(ctx).
		Model(&models.ClientScope{}).
		Where("id = ?", id).
		Updates(updates).Error
}
