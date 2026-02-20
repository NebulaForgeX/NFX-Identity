package check

import (
	"context"
	"nfxid/modules/access/infrastructure/rdb/models"
	"github.com/google/uuid"
)

func (h *Handler) ByApplicationIDAndRoleKey(ctx context.Context, applicationID uuid.UUID, roleKey string) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.ApplicationRole{}).
		Where("application_id = ? AND role_key = ?", applicationID, roleKey).Count(&n).Error
	return n > 0, err
}
