package get

import (
	"context"
	"nfxidentity/modules/tenants/domain/invitations"
	"nfxidentity/modules/tenants/infrastructure/rdb/models"
	"nfxidentity/modules/tenants/infrastructure/repository/invitations/mapper"

	"github.com/google/uuid"
)

// ByTenantIDAndStatus 根据 TenantID 和 Status 获取 Invitation 列表，实现 invitations.Get 接口
func (h *Handler) ByTenantIDAndStatus(ctx context.Context, tenantID uuid.UUID, status invitations.InvitationStatus) ([]*invitations.Invitation, error) {
	statusEnum := mapper.InvitationStatusDomainToEnum(status)
	var ms []models.Invitation
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ? AND status = ?", tenantID, statusEnum).
		Find(&ms).Error; err != nil {
		return nil, err
	}

	result := make([]*invitations.Invitation, len(ms))
	for i, m := range ms {
		result[i] = mapper.InvitationModelToDomain(&m)
	}
	return result, nil
}
