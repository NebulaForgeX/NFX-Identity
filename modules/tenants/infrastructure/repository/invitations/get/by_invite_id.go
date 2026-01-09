package get

import (
	"context"
	"errors"
	"nfxid/modules/tenants/domain/invitations"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/invitations/mapper"

	"gorm.io/gorm"
)

// ByInviteID 根据 InviteID 获取 Invitation，实现 invitations.Get 接口
func (h *Handler) ByInviteID(ctx context.Context, inviteID string) (*invitations.Invitation, error) {
	var m models.Invitation
	if err := h.db.WithContext(ctx).Where("invite_id = ?", inviteID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, invitations.ErrInvitationNotFound
		}
		return nil, err
	}
	return mapper.InvitationModelToDomain(&m), nil
}
