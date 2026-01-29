package delete

import (
	"context"

	"nfxid/modules/access/domain/action_requirements"
	"nfxid/modules/access/infrastructure/repository/action_requirements/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&mapper.ActionRequirementModel{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return action_requirements.ErrActionRequirementNotFound
	}
	return nil
}
