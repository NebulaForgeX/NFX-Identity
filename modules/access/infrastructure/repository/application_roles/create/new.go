package create

import (
	"context"
	"nfxidentity/modules/access/domain/application_roles"
	"nfxidentity/modules/access/infrastructure/repository/application_roles/mapper"
)

func (h *Handler) New(ctx context.Context, r *application_roles.ApplicationRole) error {
	m := mapper.ApplicationRoleDomainToModel(r)
	return h.db.WithContext(ctx).Create(m).Error
}
