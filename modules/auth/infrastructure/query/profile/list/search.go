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
	"nfxidentity/pkgs/query"
	"nfxidentity/pkgs/slicex"
	"strings"
)

//* =============================== Authority search.go =============================== !//
func (h *AuthorityHandler) Search(ctx context.Context, q profileQuery.ListQuery) (httpx.Page[profileQuery.AuthorityProfileItemVO], error) {
	q.Normalize()
	if q.Search != nil {
		if id, err := uuid.Parse(strings.TrimSpace(*q.Search)); err == nil {
			return h.searchByProfileID(ctx, q, id)
		}
	}

	params := mapper.ListQueryToParams(q)

	base := h.db.WithContext(ctx).Table(rdbviews.ListAuthorityProfileItem{}.TableName())
	if len(q.ExcludeProfileIDs) > 0 {
		base = base.Where(
			rdbviews.ListAuthorityProfileItemCols.ProfileID+" NOT IN ?",
			slicex.UuidSliceToStrSlice(q.ExcludeProfileIDs),
		)
	}

	rows, total, err := query.ExecuteQuery(
		ctx,
		base,
		params,
		authoritySearchQueryConfig,
		func(db *gorm.DB, data *[]rdbviews.ListAuthorityProfileItem) error {
			return db.Order(rdbviews.ListAuthorityProfileItemCols.CreatedAt + " ASC").Find(data).Error
		},
	)
	if err != nil {
		return httpx.Page[profileQuery.AuthorityProfileItemVO]{}, authErr.ErrAuthorityProfileSearchFailed.WithCause(err)
	}

	profiles := make([]profileQuery.AuthorityProfileItemVO, 0, len(rows))
	for i := range rows {
		item := mapper.ListAuthorityProfileItemViewToVO(&rows[i])
		if item.ProfileID == uuid.Nil {
			continue
		}
		profiles = append(profiles, item)
	}
	return httpx.NewPage(profiles, total), nil
}

//* =============================== Forger search.go =============================== !//
func (h *ForgerHandler) Search(ctx context.Context, q profileQuery.ListQuery) (httpx.Page[profileQuery.ForgerProfileItemVO], error) {
	q.Normalize()
	if q.Search != nil {
		if id, err := uuid.Parse(strings.TrimSpace(*q.Search)); err == nil {
			return h.searchByProfileID(ctx, q, id)
		}
	}

	params := mapper.ListQueryToParams(q)

	base := h.db.WithContext(ctx).Table(rdbviews.ListForgerProfileItem{}.TableName())
	if len(q.ExcludeProfileIDs) > 0 {
		base = base.Where(
			rdbviews.ListForgerProfileItemCols.ProfileID+" NOT IN ?",
			slicex.UuidSliceToStrSlice(q.ExcludeProfileIDs),
		)
	}

	rows, total, err := query.ExecuteQuery(
		ctx,
		base,
		params,
		forgerSearchQueryConfig,
		func(db *gorm.DB, data *[]rdbviews.ListForgerProfileItem) error {
			return db.Order(rdbviews.ListForgerProfileItemCols.CreatedAt + " ASC").Find(data).Error
		},
	)
	if err != nil {
		return httpx.Page[profileQuery.ForgerProfileItemVO]{}, authErr.ErrForgerProfileSearchFailed.WithCause(err)
	}

	profiles := make([]profileQuery.ForgerProfileItemVO, 0, len(rows))
	for i := range rows {
		item := mapper.ListForgerProfileItemViewToVO(&rows[i])
		if item.ProfileID == uuid.Nil {
			continue
		}
		profiles = append(profiles, item)
	}
	return httpx.NewPage(profiles, total), nil
}
