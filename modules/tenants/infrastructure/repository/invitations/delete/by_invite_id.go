package delete

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"
)

// ByInviteID 根据 InviteID 删除 Invitation，实现 invitations.Delete 接口
func (h *Handler) ByInviteID(ctx context.Context, inviteID string) error {
	return h.db.WithContext(ctx).
		Where("invite_id = ?", inviteID).
		Delete(&models.Invitation{}).Error
}
