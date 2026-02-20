package delete

import (
	"context"
	"nfxid/modules/access/domain/tenant_roles"
	"nfxid/modules/access/infrastructure/rdb/models"
	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	res := h.db.WithContext(ctx).Where("id = ?", id).Delete(&models.TenantRole{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return tenant_roles.ErrTenantRoleNotFound
	}
	return nil
}
