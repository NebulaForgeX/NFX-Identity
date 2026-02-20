package check

import (
	"context"
	"nfxid/modules/access/infrastructure/rdb/models"
	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.ApplicationRole{}).Where("id = ?", id).Count(&n).Error
	return n > 0, err
}
