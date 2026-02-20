package create

import (
	"context"

	dom "nfxid/modules/access/domain/tenant_role_assignments"
	"nfxid/modules/access/infrastructure/repository/tenant_role_assignments/mapper"
)

func (h *Handler) New(ctx context.Context, a *dom.TenantRoleAssignment) error {
	m := mapper.TenantRoleAssignmentDomainToModel(a)
	return h.db.WithContext(ctx).Create(m).Error
}
