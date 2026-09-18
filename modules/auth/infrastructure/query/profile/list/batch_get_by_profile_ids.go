package list

import (
	"context"
	"github.com/google/uuid"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/modules/auth/infrastructure/query/profile/mapper"
	rdbviews "nfxidentity/modules/auth/infrastructure/rdb/views"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/slicex"
)

//* =============================== Authority batch_get_by_profile_ids.go =============================== !//
func (h *AuthorityHandler) BatchGetByProfileIDs(ctx context.Context, profileIDs []uuid.UUID) ([]profileQuery.AuthorityProfileItemVO, error) {
	if len(profileIDs) == 0 {
		return nil, nil
	}

	var rows []rdbviews.ListAuthorityProfileItem
	if err := h.db.WithContext(ctx).
		Table(rdbviews.ListAuthorityProfileItem{}.TableName()).
		Where(rdbviews.ListAuthorityProfileItemCols.ProfileID+" IN ?", slicex.UuidSliceToStrSlice(profileIDs)).
		Find(&rows).Error; err != nil {
		return nil, authErr.ErrAuthorityProfileBatchGetFailed.WithCause(err)
	}

	byID := make(map[uuid.UUID]profileQuery.AuthorityProfileItemVO, len(rows))
	for i := range rows {
		item := mapper.ListAuthorityProfileItemViewToVO(&rows[i])
		if item.ProfileID == uuid.Nil {
			continue
		}
		byID[item.ProfileID] = item
	}

	out := make([]profileQuery.AuthorityProfileItemVO, 0, len(profileIDs))
	for _, id := range profileIDs {
		if item, ok := byID[id]; ok {
			out = append(out, item)
		}
	}
	return out, nil
}

//* =============================== Forger batch_get_by_profile_ids.go =============================== !//
func (h *ForgerHandler) BatchGetByProfileIDs(ctx context.Context, profileIDs []uuid.UUID) ([]profileQuery.ForgerProfileItemVO, error) {
	if len(profileIDs) == 0 {
		return nil, nil
	}

	var rows []rdbviews.ListForgerProfileItem
	if err := h.db.WithContext(ctx).
		Table(rdbviews.ListForgerProfileItem{}.TableName()).
		Where(rdbviews.ListForgerProfileItemCols.ProfileID+" IN ?", slicex.UuidSliceToStrSlice(profileIDs)).
		Find(&rows).Error; err != nil {
		return nil, authErr.ErrForgerProfileBatchGetFailed.WithCause(err)
	}

	byID := make(map[uuid.UUID]profileQuery.ForgerProfileItemVO, len(rows))
	for i := range rows {
		item := mapper.ListForgerProfileItemViewToVO(&rows[i])
		if item.ProfileID == uuid.Nil {
			continue
		}
		byID[item.ProfileID] = item
	}

	out := make([]profileQuery.ForgerProfileItemVO, 0, len(profileIDs))
	for _, id := range profileIDs {
		if item, ok := byID[id]; ok {
			out = append(out, item)
		}
	}
	return out, nil
}
