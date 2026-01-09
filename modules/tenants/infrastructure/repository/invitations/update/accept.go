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

// Accept 接受邀请，实现 invitations.Update 接口
func (h *Handler) Accept(ctx context.Context, inviteID string, userID uuid.UUID) error {
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

	// 检查是否已经接受
	if m.Status == enums.TenantsInvitationStatusAccepted {
		return invitations.ErrInvitationAlreadyAccepted
	}

	// 检查是否已过期
	if m.ExpiresAt.Before(time.Now().UTC()) {
		return invitations.ErrInvitationExpired
	}

	now := time.Now().UTC()
	status := enums.TenantsInvitationStatusAccepted
	updates := map[string]any{
		models.InvitationCols.Status:           status,
		models.InvitationCols.AcceptedByUserID: &userID,
		models.InvitationCols.AcceptedAt:      &now,
	}

	return h.db.WithContext(ctx).
		Model(&models.Invitation{}).
		Where("invite_id = ?", inviteID).
		Updates(updates).Error
}
