package check

import (
	"context"

	"nfxid/modules/access/infrastructure/repository/action_requirements/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).Model(&mapper.ActionRequirementModel{}).
		Where("id = ?", id).Count(&count).Error
	return count > 0, err
}
