package mapper

import (
	profileBadgeAppQueries "nfxid/modules/auth/application/profile_badge/queries"
	profileBadgeDomainViews "nfxid/modules/auth/domain/profile_badge/views"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/pkgs/query"
)

func ProfileBadgeModelToDomain(m *models.ProfileBadge) profileBadgeDomainViews.ProfileBadgeView {
	return profileBadgeDomainViews.ProfileBadgeView{
		ID:          m.ID,
		ProfileID:   m.ProfileID,
		BadgeID:     m.BadgeID,
		Description: m.Description,
		Level:       m.Level,
		EarnedAt:    m.EarnedAt,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func ProfileBadgeListQueryToCommonQuery(q profileBadgeAppQueries.ProfileBadgeListQuery) query.ListQueryParams {
	commonQuery := query.ListQueryParams{
		Offset: q.Offset,
		Limit:  q.Limit,
		All:    q.All,
	}

	// Convert sorts
	sortMapper := map[profileBadgeAppQueries.SortField]string{
		profileBadgeAppQueries.SortByCreatedTime: "created_at",
		profileBadgeAppQueries.SortByEarnedAt:    "earned_at",
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
