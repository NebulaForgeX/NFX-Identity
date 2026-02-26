package delete

import (
	"context"

	accessErr "nfxidentity/errors/src/access"
	"nfxidentity/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	res := h.db.WithContext(ctx).Where("id = ?", id).Delete(&models.ApplicationRole{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return accessErr.ErrApplicationRoleNotFound
	}
	return nil
}
