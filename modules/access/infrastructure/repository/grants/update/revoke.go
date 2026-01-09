package update

import (
	"context"
	"errors"
	"time"
	"nfxid/modules/access/domain/grants"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Revoke 撤销 Grant，实现 grants.Update 接口
func (h *Handler) Revoke(ctx context.Context, id uuid.UUID, revokedBy uuid.UUID, reason string) error {
	// 先检查 Grant 是否存在且未被撤销
	var m models.Grant
	if err := h.db.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return grants.ErrGrantNotFound
		}
		return err
	}

	if m.RevokedAt != nil {
		return grants.ErrGrantAlreadyRevoked
	}

	now := time.Now().UTC()
	updates := map[string]any{
		models.GrantCols.RevokedAt:    &now,
		models.GrantCols.RevokedBy:    &revokedBy,
		models.GrantCols.RevokeReason: &reason,
	}

	return h.db.WithContext(ctx).
		Model(&models.Grant{}).
		Where("id = ?", id).
		Updates(updates).Error
}
