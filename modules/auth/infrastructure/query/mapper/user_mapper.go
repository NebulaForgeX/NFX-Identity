package mapper

import (
	userDomain "nfxid/modules/auth/domain/user"
	userDomainViews "nfxid/modules/auth/domain/user/views"
	"nfxid/modules/auth/infrastructure/rdb/models"
	"nfxid/modules/auth/infrastructure/rdb/views"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"
)

func UserViewToDomain(v *views.UserWithRoleView, u *models.User) userDomainViews.UserView {
	rolesJSON := []byte("[]")
	if v.Roles != nil {
		rolesJSON = *v.Roles
	}
	passwordHash := ""
	if u != nil {
		passwordHash = u.PasswordHash
	}
	return userDomainViews.UserView{
		ID:           v.UserID,
		Username:     v.Username,
		Email:        v.Email,
		Phone:        v.Phone,
		PasswordHash: passwordHash,
		Status:       v.Status,
		IsVerified:   v.IsVerified,
		LastLoginAt:  v.LastLoginAt,
		Roles:        rolesJSON,
		CreatedAt:    v.UserCreatedAt,
		UpdatedAt:    v.UserUpdatedAt,
	}
}

func UserBadgesViewToDomain(v *views.UserBadgesView) userDomainViews.UserBadgesView {
	return userDomainViews.UserBadgesView{
		UserID:      v.UserID,
		Username:    v.Username,
		Email:       v.Email,
		ProfileID:   v.ProfileID,
		DisplayName: v.DisplayName,
		Badges:      v.Badges,
	}
}

func UserListQueryToCommonQuery(q userDomain.ListQuery) query.ListQueryParams {
	commonQuery := query.ListQueryParams{
		Offset: q.Offset,
		Limit:  q.Limit,
		All:    q.All,
		Search: ptr.Deref(q.Search),
	}

	// Convert sorts
	sortMapper := map[userDomain.SortField]string{
		userDomain.SortByCreatedTime: "user_created_at",
		userDomain.SortByUsername:    "username",
		userDomain.SortByEmail:       "email",
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
