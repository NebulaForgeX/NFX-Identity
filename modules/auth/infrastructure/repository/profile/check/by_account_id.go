package check

import (
	"context"
	"github.com/google/uuid"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
)

//* =============================== Authority by_account_id =============================== !//
func (h *Handler) AuthorityByAccountID(ctx context.Context, accountID uuid.UUID) (bool, error) {
	var n int64
	if err := h.db.WithContext(ctx).Model(&rdbmodels.AuthorityProfile{}).
		Where(rdbmodels.AuthorityProfileCols.AccountID+" = ?", accountID.String()).
		Count(&n).Error; err != nil {
		return false, err
	}
	return n > 0, nil
}

func (h *Handler) CountAuthorityByAccountID(ctx context.Context, accountID uuid.UUID) (int64, error) {
	var n int64
	if err := h.db.WithContext(ctx).Model(&rdbmodels.AuthorityProfile{}).
		Where(rdbmodels.AuthorityProfileCols.AccountID+" = ?", accountID.String()).
		Count(&n).Error; err != nil {
		return 0, err
	}
	return n, nil
}

//* =============================== Forger by_account_id =============================== !//
func (h *Handler) ForgerByAccountID(ctx context.Context, accountID uuid.UUID) (bool, error) {
	var n int64
	if err := h.db.WithContext(ctx).Model(&rdbmodels.ForgerProfile{}).
		Where(rdbmodels.ForgerProfileCols.AccountID+" = ?", accountID.String()).
		Count(&n).Error; err != nil {
		return false, err
	}
	return n > 0, nil
}

func (h *Handler) CountForgerByAccountID(ctx context.Context, accountID uuid.UUID) (int64, error) {
	var n int64
	if err := h.db.WithContext(ctx).Model(&rdbmodels.ForgerProfile{}).
		Where(rdbmodels.ForgerProfileCols.AccountID+" = ?", accountID.String()).
		Count(&n).Error; err != nil {
		return 0, err
	}
	return n, nil
}
