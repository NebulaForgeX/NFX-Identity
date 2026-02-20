package check

import (
	"context"

	"nfxid/modules/access/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) ByUserID(ctx context.Context, userID uuid.UUID) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.SuperAdmin{}).Where("user_id = ?", userID).Count(&n).Error
	return n > 0, err
}
