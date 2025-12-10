package mapper

import (
	badgeAppQueries "nebulaid/modules/auth/application/badge/queries"
	badgeDomainViews "nebulaid/modules/auth/domain/badge/views"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/pkgs/query"
	"nebulaid/pkgs/utils/ptr"
	"nebulaid/pkgs/utils/slice"
)

func BadgeQueryToCommonQuery(q badgeAppQueries.BadgeListQuery) *query.ListQueryParams {
	return &query.ListQueryParams{
		All: q.All,
		Sorts: query.DomainSortToSort(q.DomainSorts, map[badgeAppQueries.SortField]string{
			badgeAppQueries.SortByCreatedTime: "created_at",
			badgeAppQueries.SortByName:        "name",
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
	}
}
