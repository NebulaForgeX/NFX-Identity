package get

import (
	"context"
	"nfxidentity/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) Owned(ctx context.Context, accountID, id uuid.UUID) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.Forgerprofile{}).Where("id = ? AND account_id = ? AND deleted_at IS NULL", id, accountID).Count(&n).Error
	return n > 0, err
}
