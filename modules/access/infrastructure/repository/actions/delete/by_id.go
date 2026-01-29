package delete

import (
	"context"

	"nfxid/modules/access/domain/actions"
	"nfxid/modules/access/infrastructure/repository/actions/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&mapper.ActionModel{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return actions.ErrActionNotFound
	}
	return nil
}
