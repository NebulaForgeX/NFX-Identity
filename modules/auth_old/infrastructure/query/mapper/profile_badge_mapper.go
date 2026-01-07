package mapper

import (
	profileBadgeDomain "nfxid/modules/auth/domain/profile_badge"
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

func ProfileBadgeListQueryToCommonQuery(q profileBadgeDomain.ListQuery) query.ListQueryParams {
	commonQuery := query.ListQueryParams{
		Offset: q.Offset,
		Limit:  q.Limit,
		All:    q.All,
	}

	// Convert sorts
	sortMapper := map[profileBadgeDomain.SortField]string{
		profileBadgeDomain.SortByCreatedTime: "created_at",
		profileBadgeDomain.SortByEarnedAt:    "earned_at",
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
