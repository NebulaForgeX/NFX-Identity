package mapper

import (
	"time"

	profileQuery "nfxidentity/modules/auth/query/profile"
	authorityprofilepb "nfxidentity/protos/gen/auth/authority_profile"
	forgerprofilepb "nfxidentity/protos/gen/auth/forger_profile"
)

func ForgerProfileItemToPublicProfileCardProto(vo profileQuery.ForgerProfileItemVO) *forgerprofilepb.PublicProfileCard {
	out := &forgerprofilepb.PublicProfileCard{
		ProfileId: vo.ProfileID.String(),
		City:      vo.City,
		Country:   vo.Country,
		Website:   vo.Website,
		Timezone:  vo.Timezone,
	}
	if vo.DisplayName != nil {
		out.DisplayName = vo.DisplayName
	}
	if vo.AvatarImageID != nil {
		s := vo.AvatarImageID.String()
		out.AvatarImageId = &s
	}
	if vo.BackgroundImageID != nil {
		s := vo.BackgroundImageID.String()
		out.BackgroundImageId = &s
	}
	if vo.Birthday != nil {
		s := vo.Birthday.UTC().Format(time.RFC3339)
		out.Birthday = &s
	}
	return out
}

func ForgerProfileItemsToPublicProfileCardProto(vos []profileQuery.ForgerProfileItemVO) []*forgerprofilepb.PublicProfileCard {
	out := make([]*forgerprofilepb.PublicProfileCard, 0, len(vos))
	for _, vo := range vos {
		out = append(out, ForgerProfileItemToPublicProfileCardProto(vo))
	}
	return out
}

func AuthorityProfileItemToPublicProfileCardProto(vo profileQuery.AuthorityProfileItemVO) *authorityprofilepb.PublicProfileCard {
	out := &authorityprofilepb.PublicProfileCard{
		ProfileId: vo.ProfileID.String(),
		City:      vo.City,
		Country:   vo.Country,
		Website:   vo.Website,
		Timezone:  vo.Timezone,
	}
	if vo.DisplayName != nil {
		out.DisplayName = vo.DisplayName
	}
	if vo.AvatarImageID != nil {
		s := vo.AvatarImageID.String()
		out.AvatarImageId = &s
	}
	if vo.BackgroundImageID != nil {
		s := vo.BackgroundImageID.String()
		out.BackgroundImageId = &s
	}
	if vo.Birthday != nil {
		s := vo.Birthday.UTC().Format(time.RFC3339)
		out.Birthday = &s
	}
	return out
}

func AuthorityProfileItemsToPublicProfileCardProto(vos []profileQuery.AuthorityProfileItemVO) []*authorityprofilepb.PublicProfileCard {
	out := make([]*authorityprofilepb.PublicProfileCard, 0, len(vos))
	for _, vo := range vos {
		out = append(out, AuthorityProfileItemToPublicProfileCardProto(vo))
	}
	return out
}
