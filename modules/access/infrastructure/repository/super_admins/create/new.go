package create

import (
	"context"

	"nfxid/modules/access/domain/super_admins"
	"nfxid/modules/access/infrastructure/repository/super_admins/mapper"
)

func (h *Handler) New(ctx context.Context, s *super_admins.SuperAdmin) error {
	m := mapper.SuperAdminDomainToModel(s)
	return h.db.WithContext(ctx).Create(m).Error
}
