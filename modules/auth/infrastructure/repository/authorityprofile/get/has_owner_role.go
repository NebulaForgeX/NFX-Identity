package get

import (
	"context"
	"nfxidentity/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) HasOwnerRole(ctx context.Context, accountID uuid.UUID) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.AuthorityProfile{}).
		Where("account_id = ? AND deleted_at IS NULL AND authority_roles @> ARRAY['owner']::auth.authority_role[]", accountID).
		Count(&n).Error
	return n > 0, err
}
