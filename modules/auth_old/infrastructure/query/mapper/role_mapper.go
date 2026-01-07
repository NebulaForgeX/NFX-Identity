package mapper

import (
	roleDomain "nfxid/modules/auth/domain/role"
	roleDomainViews "nfxid/modules/auth/domain/role/views"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"
	"time"
)

func RoleModelToDomain(m *models.Role) roleDomainViews.RoleView {
	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	var desc *string
	if m.Description != nil && *m.Description != "" {
		desc = m.Description
	}

	return roleDomainViews.RoleView{
		ID:          m.ID,
		Name:        m.Name,
		Description: desc,
		Permissions: m.Permissions,
		IsSystem:    m.IsSystem,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}

func RoleListQueryToCommonQuery(q roleDomain.ListQuery) query.ListQueryParams {
	commonQuery := query.ListQueryParams{
		Offset: q.Offset,
		Limit:  q.Limit,
		All:    q.All,
		Search: ptr.Deref(q.Search),
	}

	// Convert sorts
	sortMapper := map[roleDomain.SortField]string{
		roleDomain.SortByCreatedTime: "created_at",
		roleDomain.SortByName:        "name",
	}

	for _, sort := range q.DomainSorts {
		if field, ok := sortMapper[sort.Field]; ok {
			commonQuery.Sorts = append(commonQuery.Sorts, query.Sort{
				Field: field,
				Order: sort.Order,
			})
		}
	}

	return commonQuery
}
