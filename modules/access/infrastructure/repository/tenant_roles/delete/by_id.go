package delete

import (
	"context"

	accessErr "nfxid/errors/src/access"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	res := h.db.WithContext(ctx).Where("id = ?", id).Delete(&models.TenantRole{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return accessErr.ErrTenantRoleNotFound
	}
	return nil
}
