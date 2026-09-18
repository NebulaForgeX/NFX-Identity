package mapper

import (
	"nfxidentity/enums"
	rdbviews "nfxidentity/modules/auth/infrastructure/rdb/views"
	profileQuery "nfxidentity/modules/auth/query/profile"
	"nfxidentity/pkgs/ptrx"
	"nfxidentity/pkgs/query"
	"strings"
)

//* =============================== ListQuery params =============================== !//
func ListQueryToParams(q profileQuery.ListQuery) *query.ListQueryParams {
	search := ""
	if q.Search != nil {
		search = strings.TrimSpace(*q.Search)
	}

	return &query.ListQueryParams{
		Offset: q.DomainPagination.Offset,
		Limit:  q.DomainPagination.Limit,
		All:    q.DomainPagination.All,
		Search: search,
	}
}

//* =============================== Authority VO mapper =============================== !//
func ListAuthorityProfileItemViewToVO(v *rdbviews.ListAuthorityProfileItem) profileQuery.AuthorityProfileItemVO {
	if v == nil || v.ProfileID == nil {
		return profileQuery.AuthorityProfileItemVO{}
	}
	return profileQuery.AuthorityProfileItemVO{
		ProfileID:         *v.ProfileID,
		AccountID:         ptrx.Deref(v.AccountID),
		AuthorityRoles:    append([]enums.AuthAuthorityRole(nil), v.AuthorityRoles...),
		DisplayName:       v.DisplayName,
		ProfileLanguage:   ptrx.Deref(v.ProfileLanguage),
		City:              v.City,
		Country:           v.Country,
		Website:           v.Website,
		Timezone:          v.Timezone,
		Birthday:          v.Birthday,
		AvatarImageID:     v.AvatarImageID,
		BackgroundImageID: v.BackgroundImageID,
		CreatedAt:         ptrx.Deref(v.CreatedAt),
	}
}

//* =============================== Forger VO mapper =============================== !//
func ListForgerProfileItemViewToVO(v *rdbviews.ListForgerProfileItem) profileQuery.ForgerProfileItemVO {
	if v == nil || v.ProfileID == nil {
		return profileQuery.ForgerProfileItemVO{}
	}
	return profileQuery.ForgerProfileItemVO{
		ProfileID:         *v.ProfileID,
		AccountID:         ptrx.Deref(v.AccountID),
		ForgerRoles:    append([]enums.AuthForgerRole(nil), v.ForgerRoles...),
		DisplayName:       v.DisplayName,
		ProfileLanguage:   ptrx.Deref(v.ProfileLanguage),
		City:              v.City,
		Country:           v.Country,
		Website:           v.Website,
		Timezone:          v.Timezone,
		Birthday:          v.Birthday,
		AvatarImageID:     v.AvatarImageID,
		BackgroundImageID: v.BackgroundImageID,
		CreatedAt:         ptrx.Deref(v.CreatedAt),
	}
}
