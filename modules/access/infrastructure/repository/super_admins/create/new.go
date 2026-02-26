package create

import (
	"context"

	"nfxidentity/modules/access/domain/super_admins"
	"nfxidentity/modules/access/infrastructure/repository/super_admins/mapper"
)

func (h *Handler) New(ctx context.Context, s *super_admins.SuperAdmin) error {
	m := mapper.SuperAdminDomainToModel(s)
	return h.db.WithContext(ctx).Create(m).Error
}
