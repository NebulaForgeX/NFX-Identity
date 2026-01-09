package get

import (
	"context"
	"nfxid/modules/tenants/domain/invitations"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/invitations/mapper"
)

// ByEmail 根据 Email 获取 Invitation 列表，实现 invitations.Get 接口
func (h *Handler) ByEmail(ctx context.Context, email string) ([]*invitations.Invitation, error) {
	var ms []models.Invitation
	if err := h.db.WithContext(ctx).
		Where("email = ?", email).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*invitations.Invitation, len(ms))
	for i, m := range ms {
		result[i] = mapper.InvitationModelToDomain(&m)
	}
	return result, nil
}
