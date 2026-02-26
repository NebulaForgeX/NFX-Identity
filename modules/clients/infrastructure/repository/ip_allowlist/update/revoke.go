package update

import (
	"context"
	"errors"
	"nfxidentity/enums"
	clientsErr "nfxidentity/errors/src/clients"
	"nfxidentity/modules/clients/infrastructure/rdb/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Revoke 撤销 IPAllowlist，实现 ip_allowlist.Update 接口
func (h *Handler) Revoke(ctx context.Context, ruleID string, revokedBy uuid.UUID, reason string) error {
	// 先检查 IPAllowlist 是否存在
	var m models.IpAllowlist
	if err := h.db.WithContext(ctx).
		Where("rule_id = ?", ruleID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return clientsErr.ErrIPAllowlistNotFound
		}
		return err
	}

	// 检查是否已经撤销
	if m.Status == enums.ClientsAllowlistStatusRevoked {
		return clientsErr.ErrIPAllowlistAlreadyRevoked
	}

	now := time.Now().UTC()
	updates := map[string]any{
		models.IpAllowlistCols.Status:       enums.ClientsAllowlistStatusRevoked,
		models.IpAllowlistCols.RevokedAt:    &now,
		models.IpAllowlistCols.RevokedBy:    &revokedBy,
		models.IpAllowlistCols.RevokeReason: &reason,
		models.IpAllowlistCols.UpdatedAt:    now,
	}

	return h.db.WithContext(ctx).
		Model(&models.IpAllowlist{}).
		Where("rule_id = ?", ruleID).
		Updates(updates).Error
}
