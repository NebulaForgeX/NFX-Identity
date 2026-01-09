package update

import (
	"context"
	"errors"
	"time"
	"nfxid/enums"
	"nfxid/modules/tenants/domain/invitations"
	"nfxid/modules/tenants/infrastructure/rdb/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Revoke 撤销邀请，实现 invitations.Update 接口
func (h *Handler) Revoke(ctx context.Context, inviteID string, revokedBy uuid.UUID, reason string) error {
	// 先检查 Invitation 是否存在
	var m models.Invitation
	if err := h.db.WithContext(ctx).
		Where("invite_id = ?", inviteID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return invitations.ErrInvitationNotFound
		}
		return err
	}

	// 检查是否已经撤销
	if m.Status == enums.TenantsInvitationStatusRevoked {
		return invitations.ErrInvitationAlreadyRevoked
	}

	now := time.Now().UTC()
	status := enums.TenantsInvitationStatusRevoked
	updates := map[string]any{
		models.InvitationCols.Status:       status,
		models.InvitationCols.RevokedBy:    &revokedBy,
		models.InvitationCols.RevokedAt:    &now,
		models.InvitationCols.RevokeReason: &reason,
	}

	return h.db.WithContext(ctx).
		Model(&models.Invitation{}).
		Where("invite_id = ?", inviteID).
		Updates(updates).Error
}
