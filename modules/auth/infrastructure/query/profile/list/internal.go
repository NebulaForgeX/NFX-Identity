package list

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/infrastructure/query/profile/mapper"
	rdbviews "nfxidentity/modules/auth/infrastructure/rdb/views"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/httpx"
)

//* =============================== Authority internal.go =============================== !//
func (h *AuthorityHandler) searchByProfileID(
	ctx context.Context,
	q profileQuery.ListQuery,
	profileID uuid.UUID,
) (httpx.Page[profileQuery.AuthorityProfileItemVO], error) {
	for _, excluded := range q.ExcludeProfileIDs {
		if excluded == profileID {
			return httpx.NewPage([]profileQuery.AuthorityProfileItemVO{}, 0), nil
		}
	}

	var row rdbviews.ListAuthorityProfileItem
	err := h.db.WithContext(ctx).
		Table(rdbviews.ListAuthorityProfileItem{}.TableName()).
		Where(rdbviews.ListAuthorityProfileItemCols.ProfileID+" = ?", profileID).
		First(&row).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return httpx.NewPage([]profileQuery.AuthorityProfileItemVO{}, 0), nil
		}
		return httpx.Page[profileQuery.AuthorityProfileItemVO]{}, authErr.ErrAuthorityProfileSearchFailed.WithCause(err)
	}
	item := mapper.ListAuthorityProfileItemViewToVO(&row)
	if item.ProfileID == uuid.Nil {
		return httpx.NewPage([]profileQuery.AuthorityProfileItemVO{}, 0), nil
	}
	return httpx.NewPage([]profileQuery.AuthorityProfileItemVO{item}, 1), nil
}

//* =============================== Forger internal.go =============================== !//
// searchByProfileID handles exact UUID search for user profile list queries.
func (h *ForgerHandler) searchByProfileID(
	ctx context.Context,
	q profileQuery.ListQuery,
	profileID uuid.UUID,
) (httpx.Page[profileQuery.ForgerProfileItemVO], error) {
	for _, excluded := range q.ExcludeProfileIDs {
		if excluded == profileID {
			return httpx.NewPage([]profileQuery.ForgerProfileItemVO{}, 0), nil
		}
	}

	var row rdbviews.ListForgerProfileItem
	err := h.db.WithContext(ctx).
		Table(rdbviews.ListForgerProfileItem{}.TableName()).
		Where(rdbviews.ListForgerProfileItemCols.ProfileID+" = ?", profileID).
		First(&row).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return httpx.NewPage([]profileQuery.ForgerProfileItemVO{}, 0), nil
		}
		return httpx.Page[profileQuery.ForgerProfileItemVO]{}, authErr.ErrForgerProfileSearchFailed.WithCause(err)
	}
	item := mapper.ListForgerProfileItemViewToVO(&row)
	if item.ProfileID == uuid.Nil {
		return httpx.NewPage([]profileQuery.ForgerProfileItemVO{}, 0), nil
	}
	return httpx.NewPage([]profileQuery.ForgerProfileItemVO{item}, 1), nil
}
