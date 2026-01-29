package delete

import (
	"context"

	"nfxid/modules/access/infrastructure/repository/action_requirements/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByActionID(ctx context.Context, actionID uuid.UUID) error {
	return h.db.WithContext(ctx).
		Where("action_id = ?", actionID).
		Delete(&mapper.ActionRequirementModel{}).Error
}
