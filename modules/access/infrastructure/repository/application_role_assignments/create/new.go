package create

import (
	"context"

	"nfxid/modules/access/domain/application_role_assignments"
	"nfxid/modules/access/infrastructure/repository/application_role_assignments/mapper"
)

func (h *Handler) New(ctx context.Context, a *application_role_assignments.ApplicationRoleAssignment) error {
	m := mapper.ApplicationRoleAssignmentDomainToModel(a)
	return h.db.WithContext(ctx).Create(m).Error
}
