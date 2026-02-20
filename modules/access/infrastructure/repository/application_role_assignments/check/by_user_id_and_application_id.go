package check

import (
	"context"

	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) ByUserIDAndApplicationID(ctx context.Context, userID, applicationID uuid.UUID) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.ApplicationRoleAssignment{}).
		Where("user_id = ? AND application_id = ?", userID, applicationID).Count(&n).Error
	return n > 0, err
}
