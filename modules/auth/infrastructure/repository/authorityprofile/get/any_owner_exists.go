package get

import (
	"context"

	"nfxidentity/modules/auth/infrastructure/rdb/models"
)

func (h *Handler) AnyOwnerExists(ctx context.Context) (bool, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(&models.Authorityprofile{}).
		Where("deleted_at IS NULL AND authority_roles @> ARRAY['owner']::auth.authority_role[]").
		Count(&n).Error
	return n > 0, err
}
