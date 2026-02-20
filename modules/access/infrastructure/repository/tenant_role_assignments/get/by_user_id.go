package get

import (
	"context"

	dom "nfxid/modules/access/domain/tenant_role_assignments"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/tenant_role_assignments/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) ([]*dom.TenantRoleAssignment, error) {
	var list []*models.TenantRoleAssignment
	if err := h.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("assigned_at ASC").
		Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*dom.TenantRoleAssignment, len(list))
	for i := range list {
		out[i] = mapper.TenantRoleAssignmentModelToDomain(list[i])
	}
	return out, nil
}
