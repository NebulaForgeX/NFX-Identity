package update

import (
	"context"
	"nfxid/modules/access/domain/tenant_roles"
	"nfxid/modules/access/infrastructure/rdb/models"
	"nfxid/modules/access/infrastructure/repository/tenant_roles/mapper"
)

func (h *Handler) Generic(ctx context.Context, r *tenant_roles.TenantRole) error {
	m := mapper.TenantRoleDomainToModel(r)
	return h.db.WithContext(ctx).Model(&models.TenantRole{}).
		Where("id = ?", r.ID()).
		Updates(map[string]interface{}{
			models.TenantRoleCols.RoleKey: m.RoleKey,
			models.TenantRoleCols.Name:    m.Name,
		}).Error
}
