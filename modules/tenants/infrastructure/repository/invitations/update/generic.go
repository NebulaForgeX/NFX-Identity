package update

import (
	"context"
	"nfxid/modules/tenants/domain/invitations"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/invitations/mapper"
)

// Generic 通用更新 Invitation，实现 invitations.Update 接口
func (h *Handler) Generic(ctx context.Context, i *invitations.Invitation) error {
	m := mapper.InvitationDomainToModel(i)
	updates := mapper.InvitationModelToUpdates(m)
	return h.db.WithContext(ctx).
		Model(&models.Invitation{}).
		Where("id = ?", i.ID()).
		Updates(updates).Error
}
