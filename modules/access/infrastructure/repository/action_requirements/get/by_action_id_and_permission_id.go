package get

import (
	"context"
	"errors"

	"nfxid/modules/access/domain/action_requirements"
	"nfxid/modules/access/infrastructure/repository/action_requirements/mapper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) ByActionIDAndPermissionID(ctx context.Context, actionID, permissionID uuid.UUID) (*action_requirements.ActionRequirement, error) {
	var m mapper.ActionRequirementModel
	if err := h.db.WithContext(ctx).
		Where("action_id = ? AND permission_id = ?", actionID, permissionID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, action_requirements.ErrActionRequirementNotFound
		}
		return nil, err
	}
	return mapper.ActionRequirementModelToDomain(&m), nil
}
