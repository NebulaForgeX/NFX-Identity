package get

import (
	"context"

	"nfxid/modules/access/domain/application_role_assignments"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/application_role_assignments/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByApplicationID(ctx context.Context, applicationID uuid.UUID) ([]*application_role_assignments.ApplicationRoleAssignment, error) {
	var list []*models.ApplicationRoleAssignment
	if err := h.db.WithContext(ctx).
		Where("application_id = ?", applicationID).
		Order("assigned_at ASC").
		Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*application_role_assignments.ApplicationRoleAssignment, len(list))
	for i := range list {
		out[i] = mapper.ApplicationRoleAssignmentModelToDomain(list[i])
	}
	return out, nil
}
