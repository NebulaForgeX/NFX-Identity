package get

import (
	"context"

	"nfxid/modules/access/domain/action_requirements"
	"nfxid/modules/access/infrastructure/repository/action_requirements/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByActionID(ctx context.Context, actionID uuid.UUID) ([]*action_requirements.ActionRequirement, error) {
	var ms []mapper.ActionRequirementModel
	if err := h.db.WithContext(ctx).
		Where("action_id = ?", actionID).
		Find(&ms).Error; err != nil {
		return nil, err
	}
	result := make([]*action_requirements.ActionRequirement, len(ms))
	for i := range ms {
		result[i] = mapper.ActionRequirementModelToDomain(&ms[i])
	}
	return result, nil
}
