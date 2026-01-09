package check

import (
	"context"
	"nfxid/modules/tenants/infrastructure/rdb/models"
)

// ByInviteID 根据 InviteID 检查 Invitation 是否存在，实现 invitations.Check 接口
func (h *Handler) ByInviteID(ctx context.Context, inviteID string) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).
		Model(&models.Invitation{}).
		Where("invite_id = ?", inviteID).
		Count(&count).Error
	return count > 0, err
}
