package delete

import (
	"context"

	accessErr "nfxid/errors/src/access"
	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) error {
	res := h.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&models.SuperAdmin{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return accessErr.ErrSuperAdminNotFound
	}
	return nil
}
