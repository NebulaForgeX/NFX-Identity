package update

import (
	"context"

	"nfxid/modules/access/domain/application_roles"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/application_roles/mapper"
)

func (h *Handler) Generic(ctx context.Context, r *application_roles.ApplicationRole) error {
	m := mapper.ApplicationRoleDomainToModel(r)
	return h.db.WithContext(ctx).Model(&models.ApplicationRole{}).
		Where("id = ?", r.ID()).
		Updates(map[string]interface{}{
			models.ApplicationRoleCols.RoleKey: m.RoleKey,
			models.ApplicationRoleCols.Name:    m.Name,
		}).Error
}
