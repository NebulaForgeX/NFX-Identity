package mapper

import (
	educationAppQueries "nebulaid/modules/auth/application/education/queries"
	educationDomainViews "nebulaid/modules/auth/domain/education/views"
	"nebulaid/modules/auth/infrastructure/rdb/models"
	"nebulaid/pkgs/query"
	"nebulaid/pkgs/utils/ptr"
	"time"
)

func EducationModelToDomain(m *models.Education) educationDomainViews.EducationView {
	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}

	return educationDomainViews.EducationView{
		ID:           m.ID,
		ProfileID:    m.ProfileID,
		School:       m.School,
		Degree:       m.Degree,
		Major:        m.Major,
		FieldOfStudy: m.FieldOfStudy,
		StartDate:    m.StartDate,
		EndDate:      m.EndDate,
		IsCurrent:    m.IsCurrent,
		Description:  m.Description,
		Grade:        m.Grade,
		Activities:   m.Activities,
		Achievements: m.Achievements,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		DeletedAt:    deletedAt,
	}
}

func EducationListQueryToCommonQuery(q educationAppQueries.EducationListQuery) query.ListQueryParams {
	commonQuery := query.ListQueryParams{
		Offset: q.Offset,
		Limit:  q.Limit,
		All:    q.All,
		Search: ptr.Deref(q.Search),
	}

	// Convert sorts
	sortMapper := map[educationAppQueries.SortField]string{
		educationAppQueries.SortByCreatedTime: "created_at",
		educationAppQueries.SortByStartDate:   "start_date",
		educationAppQueries.SortBySchool:      "school",
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
