package list

import (
	"context"
	"github.com/google/uuid"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/infrastructure/query/profile/mapper"
	rdbviews "nfxidentity/modules/auth/infrastructure/rdb/views"
	profileQuery "nfxidentity/modules/auth/query/profile"
)

//* =============================== Authority by_account_id.go =============================== !//
func (h *AuthorityHandler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]profileQuery.AuthorityProfileItemVO, error) {
	var rows []rdbviews.ListAuthorityProfileItem
	if err := h.db.WithContext(ctx).
		Table(rdbviews.ListAuthorityProfileItem{}.TableName()).
		Where(rdbviews.ListAuthorityProfileItemCols.AccountID+" = ?", accountID.String()).
		Order(rdbviews.ListAuthorityProfileItemCols.CreatedAt + " ASC").
		Find(&rows).Error; err != nil {
		return nil, authErr.ErrAccountProfileGetFailed
	}
	out := make([]profileQuery.AuthorityProfileItemVO, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.ListAuthorityProfileItemViewToVO(&rows[i]))
	}
	return out, nil
}

//* =============================== Forger by_account_id.go =============================== !//
func (h *ForgerHandler) ByAccountID(ctx context.Context, accountID uuid.UUID) ([]profileQuery.ForgerProfileItemVO, error) {
	var rows []rdbviews.ListForgerProfileItem
	if err := h.db.WithContext(ctx).
		Table(rdbviews.ListForgerProfileItem{}.TableName()).
		Where(rdbviews.ListForgerProfileItemCols.AccountID+" = ?", accountID.String()).
		Order(rdbviews.ListForgerProfileItemCols.CreatedAt + " ASC").
		Find(&rows).Error; err != nil {
		return nil, authErr.ErrAccountProfileGetFailed
	}
	out := make([]profileQuery.ForgerProfileItemVO, 0, len(rows))
	for i := range rows {
		out = append(out, mapper.ListForgerProfileItemViewToVO(&rows[i]))
	}
	return out, nil
}
