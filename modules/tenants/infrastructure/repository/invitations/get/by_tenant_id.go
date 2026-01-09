package get

import (
	"context"
	"nfxid/modules/tenants/domain/invitations"
	"nfxid/modules/tenants/infrastructure/rdb/models"
	"nfxid/modules/tenants/infrastructure/repository/invitations/mapper"

	"github.com/google/uuid"
)

// ByTenantID 根据 TenantID 获取 Invitation 列表，实现 invitations.Get 接口
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*invitations.Invitation, error) {
	var ms []models.Invitation
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*invitations.Invitation, len(ms))
	for i, m := range ms {
		result[i] = mapper.InvitationModelToDomain(&m)
	}
	return result, nil
}
