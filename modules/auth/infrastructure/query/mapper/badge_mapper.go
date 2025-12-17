package mapper

import (
	badgeDomain "nfxid/modules/auth/domain/badge"
	badgeDomainViews "nfxid/modules/auth/domain/badge/views"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"
	"nfxid/pkgs/utils/slice"
	"time"
)

func BadgeQueryToCommonQuery(q badgeDomain.ListQuery) *query.ListQueryParams {
	return &query.ListQueryParams{
		All: q.All,
		Sorts: query.DomainSortToSort(q.DomainSorts, map[badgeDomain.SortField]string{
			badgeDomain.SortByCreatedTime: "created_at",
			badgeDomain.SortByName:        "name",
		}),
		Search: ptr.Deref(q.Search),
		Filters: map[string][]any{
			"id":        slice.ToAnySlice(q.BadgeIDs),
			"category":  {q.Category},
			"is_system": {q.IsSystem},
		},
	}
}

func BadgeDBToDomain(m *models.Badge) badgeDomainViews.BadgeView {
	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}
	return badgeDomainViews.BadgeView{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		IconURL:     m.IconURL,
		Color:       m.Color,
		Category:    m.Category,
		IsSystem:    m.IsSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}
