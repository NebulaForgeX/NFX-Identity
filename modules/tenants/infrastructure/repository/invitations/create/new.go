package create

import (
	"context"
	"nfxidentity/modules/tenants/domain/invitations"
	"nfxidentity/modules/tenants/infrastructure/repository/invitations/mapper"
)

// New 创建新的 Invitation，实现 invitations.Create 接口
func (h *Handler) New(ctx context.Context, i *invitations.Invitation) error {
	m := mapper.InvitationDomainToModel(i)
	return h.db.WithContext(ctx).Create(&m).Error
}
