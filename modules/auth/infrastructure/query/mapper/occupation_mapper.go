package mapper

import (
	occupationAppQueries "nebulaid/modules/auth/application/occupation/queries"
	occupationDomainViews "nebulaid/modules/auth/domain/occupation/views"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/pkgs/query"
	"nebulaid/pkgs/utils/ptr"
	"time"
)

func OccupationModelToDomain(m *models.Occupation) occupationDomainViews.OccupationView {
	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	return occupationDomainViews.OccupationView{
		ID:               m.ID,
		ProfileID:        m.ProfileID,
		Company:          m.Company,
		Position:         m.Position,
		Department:       m.Department,
		Industry:         m.Industry,
		Location:         m.Location,
		EmploymentType:   m.EmploymentType,
		StartDate:        m.StartDate,
		EndDate:          m.EndDate,
		IsCurrent:        m.IsCurrent,
		Description:      m.Description,
		Responsibilities: m.Responsibilities,
		Achievements:     m.Achievements,
		SkillsUsed:       m.SkillsUsed,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
		DeletedAt:        deletedAt,
	}
}

func OccupationListQueryToCommonQuery(q occupationAppQueries.OccupationListQuery) query.ListQueryParams {
	commonQuery := query.ListQueryParams{
		Offset: q.Offset,
		Limit:  q.Limit,
		All:    q.All,
		Search: ptr.Deref(q.Search),
	}

	// Convert sorts
	sortMapper := map[occupationAppQueries.SortField]string{
		occupationAppQueries.SortByCreatedTime: "created_at",
		occupationAppQueries.SortByStartDate:   "start_date",
		occupationAppQueries.SortByCompany:     "company",
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
