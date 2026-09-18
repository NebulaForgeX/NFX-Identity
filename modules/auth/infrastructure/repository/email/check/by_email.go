package check

import (
	"context"

	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
)

// ByEmail reports whether a non-deleted auth email row exists for normalizedEmail (must match DB-stored normalization).
func (h *Handler) ByEmail(ctx context.Context, normalizedEmail string) (bool, error) {
	var n int64
	if err := h.db.WithContext(ctx).Model(&rdbmodels.Email{}).
		Where(rdbmodels.EmailCols.Email+" = ?", normalizedEmail).
		Count(&n).Error; err != nil {
		return false, err
	}
	return n > 0, nil
}
