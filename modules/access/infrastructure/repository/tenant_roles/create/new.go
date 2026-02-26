package create

import (
	"context"
	"nfxidentity/modules/access/domain/tenant_roles"
	"nfxidentity/modules/access/infrastructure/repository/tenant_roles/mapper"
)

func (h *Handler) New(ctx context.Context, r *tenant_roles.TenantRole) error {
	m := mapper.TenantRoleDomainToModel(r)
	return h.db.WithContext(ctx).Create(m).Error
}
