package delete

import (
	"context"

	"nfxid/modules/access/domain/application_role_assignments"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	res := h.db.WithContext(ctx).Where("id = ?", id).Delete(&models.ApplicationRoleAssignment{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return application_role_assignments.ErrApplicationRoleAssignmentNotFound
	}
	return nil
}
