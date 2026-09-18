package get

import (
	"context"
	"nfxidentity/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) HasOwnerRole(ctx context.Context, accountID uuid.UUID) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.Forgerprofile{}).
		Where("account_id = ? AND deleted_at IS NULL AND forger_roles @> ARRAY['owner']::auth.forger_role[]", accountID).
		Count(&n).Error
	return n > 0, err
}
