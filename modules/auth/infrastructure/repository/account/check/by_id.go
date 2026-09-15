package check

import (
	"context"
	"nfxidentity/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) ByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.Account{}).Where("id = ? AND deleted_at IS NULL", id).Count(&n).Error
	return n > 0, err
}
