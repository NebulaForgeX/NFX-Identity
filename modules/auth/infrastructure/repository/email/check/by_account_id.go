package check

import (
	"context"

	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"

	"github.com/google/uuid"
)

func (h *Handler) CountByAccountID(ctx context.Context, accountID uuid.UUID) (int64, error) {
	var n int64
	if err := h.db.WithContext(ctx).Model(&rdbmodels.Email{}).
		Where(rdbmodels.EmailCols.AccountID+" = ?", accountID.String()).
		Count(&n).Error; err != nil {
		return 0, err
	}
	return n, nil
}

func (h *Handler) CountVerifiedByAccountID(ctx context.Context, accountID uuid.UUID) (int64, error) {
	var n int64
	if err := h.db.WithContext(ctx).Model(&rdbmodels.Email{}).
		Where(rdbmodels.EmailCols.AccountID+" = ?", accountID.String()).
		Where(rdbmodels.EmailCols.VerifiedAt + " IS NOT NULL").
		Count(&n).Error; err != nil {
		return 0, err
	}
	return n, nil
}
