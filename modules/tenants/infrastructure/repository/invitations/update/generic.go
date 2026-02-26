package update

import (
	"context"
	"nfxidentity/modules/tenants/domain/invitations"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/invitations/mapper"
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
