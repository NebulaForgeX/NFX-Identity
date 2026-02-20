package get

import (
	"context"

	"nfxid/modules/access/domain/tenant_roles"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/tenant_roles/mapper"

	"github.com/google/uuid"
)

// ByTenantID 按租户ID列表
func (h *Handler) ByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*tenant_roles.TenantRole, error) {
	var list []*models.TenantRole
	if err := h.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Order("created_at ASC").
		Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*tenant_roles.TenantRole, len(list))
	for i := range list {
		out[i] = mapper.TenantRoleModelToDomain(list[i])
	}
	return out, nil
}
