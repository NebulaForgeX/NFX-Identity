package mapper

import (
	userRoleDomain "nfxid/modules/auth/domain/user_role"
	userRoleDomainViews "nfxid/modules/auth/domain/user_role/views"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/query"
)

func UserRoleModelToDomain(m *models.UserRole) userRoleDomainViews.UserRoleView {
	return userRoleDomainViews.UserRoleView{
		ID:        m.ID,
		UserID:    m.UserID,
		RoleID:    m.RoleID,
		CreatedAt: m.CreatedAt,
	}
}

func UserRoleListQueryToCommonQuery(q userRoleDomain.ListQuery) query.ListQueryParams {
	commonQuery := query.ListQueryParams{
		Offset: q.Offset,
		Limit:  q.Limit,
		All:    q.All,
	}

	// Convert sorts
	sortMapper := map[userRoleDomain.SortField]string{
		userRoleDomain.SortByCreatedTime: "created_at",
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
