package check

import (
	"context"

	"nfxid/modules/access/infrastructure/repository/action_requirements/mapper"

	"github.com/google/uuid"
)

func (h *Handler) ByActionIDAndPermissionID(ctx context.Context, actionID, permissionID uuid.UUID) (bool, error) {
	var count int64
	err := h.db.WithContext(ctx).Model(&mapper.ActionRequirementModel{}).
		Where("action_id = ? AND permission_id = ?", actionID, permissionID).
		Count(&count).Error
	return count > 0, err
}
