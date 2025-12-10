package mapper

import (
	profileAppQueries "nebulaid/modules/auth/application/profile/queries"
	profileDomainViews "nebulaid/modules/auth/domain/profile/views"
	"nebulaid/modules/auth/infrastructure/rdb/views"
	"nebulaid/pkgs/query"
	"nebulaid/pkgs/utils/ptr"
	"strings"

	"github.com/google/uuid"
)

func ProfileViewToDomain(v *views.ProfileCompleteView) profileDomainViews.ProfileView {
	view := profileDomainViews.ProfileView{
		UserID:          v.UserID,
		Username:        v.Username,
		Email:           v.Email,
		UserPhone:       v.UserPhone,
		UserStatus:      v.UserStatus,
		IsVerified:      v.IsVerified,
		FirstName:       v.FirstName,
		LastName:        v.LastName,
		Nickname:        v.Nickname,
		DisplayName:     v.DisplayName,
		AvatarID:        v.AvatarID,
		BackgroundID:    v.BackgroundID,
		Bio:             v.Bio,
		Phone:           v.ProfilePhone,
		Birthday:        v.Birthday,
		Age:             v.Age,
		Gender:          v.Gender,
		Location:        v.Location,
		Website:         v.Website,
		Github:          v.Github,
		SocialLinks:     v.SocialLinks,
		Preferences:     v.Preferences,
		Skills:          v.Skills,
		PrivacySettings: v.PrivacySettings,
		Occupations:     v.Occupations,
		Educations:      v.Educations,
	}

	if v.ProfileID != nil {
		view.ID = *v.ProfileID
	}

	// Parse BackgroundIds from string to []uuid.UUID
	if v.BackgroundIds != nil && *v.BackgroundIds != "" {
		idsStr := strings.Trim(*v.BackgroundIds, "{}")
		if idsStr != "" {
			parts := strings.Split(idsStr, ",")
			view.BackgroundIds = make([]uuid.UUID, 0, len(parts))
			for _, part := range parts {
				part = strings.TrimSpace(part)
				if id, err := uuid.Parse(part); err == nil {
					view.BackgroundIds = append(view.BackgroundIds, id)
				}
			}
		}
	}

	if v.ProfileCreatedAt != nil {
		view.CreatedAt = *v.ProfileCreatedAt
	}
	if v.ProfileUpdatedAt != nil {
		view.UpdatedAt = *v.ProfileUpdatedAt
	}

	return view
}

func ProfileListQueryToCommonQuery(q profileAppQueries.ProfileListQuery) query.ListQueryParams {
	commonQuery := query.ListQueryParams{
		Offset: q.Offset,
		Limit:  q.Limit,
		All:    q.All,
		Search: ptr.Deref(q.Search),
	}

	// Convert sorts
	sortMapper := map[profileAppQueries.SortField]string{
		profileAppQueries.SortByCreatedTime: "profile_created_at",
		profileAppQueries.SortByDisplayName: "display_name",
		profileAppQueries.SortByNickname:    "nickname",
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
