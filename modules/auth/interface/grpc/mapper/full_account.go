package mapper

import (
	"time"

	"nfxidentity/enums"
	accountQuery "nfxidentity/modules/auth/query/account"
	accountpb "nfxidentity/protos/gen/auth/account"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CommunityFullToProto(in *accountQuery.FullAccountInformationWithCommunityProfileVO) *accountpb.FullAccountInformation {
	if in == nil {
		return &accountpb.FullAccountInformation{}
	}
	full := &accountpb.FullAccountInformation{
		Account: accountToProto(in.Account),
		Emails:  emailsToProto(in.Emails),
		Phones:  phonesToProto(in.Phones),
	}
	if in.CommunityProfile != nil {
		full.ForgerProfile = communityProfileToProto(in.CommunityProfile)
	}
	return full
}

func AuthorityFullToProto(in *accountQuery.FullAccountInformationWithAuthorityProfileVO) *accountpb.FullAccountInformation {
	if in == nil {
		return &accountpb.FullAccountInformation{}
	}
	full := &accountpb.FullAccountInformation{
		Account: accountToProto(in.Account),
		Emails:  emailsToProto(in.Emails),
		Phones:  phonesToProto(in.Phones),
	}
	if in.AuthorityProfile != nil {
		full.AuthorityProfile = authorityProfileToProto(in.AuthorityProfile)
	}
	return full
}

func accountToProto(acc accountQuery.AccountVO) *accountpb.Account {
	return &accountpb.Account{
		Id:             acc.ID.String(),
		SignupPlatform: string(acc.SignupPlatform),
		AccountStatus:  accountStatusToProto(acc.AccountStatus),
		CreatedAt:      ts(acc.CreatedAt),
		UpdatedAt:      ts(acc.UpdatedAt),
	}
}

func emailsToProto(rows []accountQuery.EmailVO) []*accountpb.Email {
	out := make([]*accountpb.Email, 0, len(rows))
	for _, e := range rows {
		item := &accountpb.Email{
			Id:        e.ID.String(),
			AccountId: e.AccountID.String(),
			Email:     e.Email,
			IsPrimary: e.IsPrimary,
			CreatedAt: ts(e.CreatedAt),
			UpdatedAt: ts(e.UpdatedAt),
		}
		if e.VerifiedAt != nil {
			item.VerifiedAt = ts(*e.VerifiedAt)
		}
		out = append(out, item)
	}
	return out
}

func phonesToProto(rows []accountQuery.PhoneVO) []*accountpb.Phone {
	out := make([]*accountpb.Phone, 0, len(rows))
	for _, p := range rows {
		item := &accountpb.Phone{
			Id:        p.ID.String(),
			AccountId: p.AccountID.String(),
			Phone:     p.Phone,
			IsPrimary: p.IsPrimary,
			CreatedAt: ts(p.CreatedAt),
			UpdatedAt: ts(p.UpdatedAt),
		}
		if p.VerifiedAt != nil {
			item.VerifiedAt = ts(*p.VerifiedAt)
		}
		out = append(out, item)
	}
	return out
}

func communityProfileToProto(p *accountQuery.CommunityProfileVO) *accountpb.ForgerProfile {
	out := &accountpb.ForgerProfile{
		AccountId:       p.AccountID.String(),
		ProfileId:       p.ProfileID.String(),
		ProfileLanguage: profileLangToProto(p.ProfileLanguage),
		DisplayName:     p.DisplayName,
		FirstName:       p.FirstName,
		LastName:        p.LastName,
		Country:         p.Country,
		City:            p.City,
		Gender:          p.Gender,
		Website:         p.Website,
		Timezone:        p.Timezone,
		Bio:             p.Bio,
		CreatedAt:       ts(p.CreatedAt),
		UpdatedAt:       ts(p.UpdatedAt),
	}
	if p.Preference != nil {
		out.PreferenceJson = string(*p.Preference)
	}
	if p.Birthday != nil {
		s := p.Birthday.Format(time.RFC3339)
		out.Birthday = &s
	}
	for _, role := range p.ForgerRoles {
		if role == enums.AuthForgerRoleForger {
			out.ForgerRoles = append(out.ForgerRoles, accountpb.ForgerRole_FORGER_ROLE_FORGER)
		}
	}
	for _, a := range p.Avatars {
		out.Avatars = append(out.Avatars, &accountpb.ForgerProfileAvatar{
			Id:        a.ID.String(),
			ProfileId: a.ProfileID.String(),
			ImageId:   a.ImageID.String(),
			IsActive:  a.IsActive,
			CreatedAt: ts(a.CreatedAt),
			UpdatedAt: ts(a.UpdatedAt),
		})
	}
	for _, b := range p.Backgrounds {
		out.Backgrounds = append(out.Backgrounds, &accountpb.ForgerProfileBackground{
			Id:        b.ID.String(),
			ProfileId: b.ProfileID.String(),
			ImageId:   b.ImageID.String(),
			SortOrder: b.SortOrder,
			CreatedAt: ts(b.CreatedAt),
			UpdatedAt: ts(b.UpdatedAt),
		})
	}
	return out
}

func authorityProfileToProto(p *accountQuery.AuthorityProfileVO) *accountpb.AuthorityProfile {
	out := &accountpb.AuthorityProfile{
		AccountId:       p.AccountID.String(),
		ProfileId:       p.ProfileID.String(),
		ProfileLanguage: profileLangToProto(p.ProfileLanguage),
		DisplayName:     p.DisplayName,
		FirstName:       p.FirstName,
		LastName:        p.LastName,
		Country:         p.Country,
		City:            p.City,
		Gender:          p.Gender,
		Website:         p.Website,
		Timezone:        p.Timezone,
		Bio:             p.Bio,
		CreatedAt:       ts(p.CreatedAt),
		UpdatedAt:       ts(p.UpdatedAt),
	}
	if p.Preference != nil {
		out.PreferenceJson = string(*p.Preference)
	}
	if p.Birthday != nil {
		s := p.Birthday.Format(time.RFC3339)
		out.Birthday = &s
	}
	for _, role := range p.AuthorityRoles {
		switch role {
		case enums.AuthAuthorityRoleAuditor:
			out.AuthorityRoles = append(out.AuthorityRoles, accountpb.AuthorityRole_AUTHORITY_ROLE_AUDITOR)
		case enums.AuthAuthorityRoleAdministrator:
			out.AuthorityRoles = append(out.AuthorityRoles, accountpb.AuthorityRole_AUTHORITY_ROLE_ADMINISTRATOR)
		case enums.AuthAuthorityRoleOwner:
			out.AuthorityRoles = append(out.AuthorityRoles, accountpb.AuthorityRole_AUTHORITY_ROLE_OWNER)
		}
	}
	for _, a := range p.Avatars {
		out.Avatars = append(out.Avatars, &accountpb.AuthorityProfileAvatar{
			Id:        a.ID.String(),
			ProfileId: a.ProfileID.String(),
			ImageId:   a.ImageID.String(),
			IsActive:  a.IsActive,
			CreatedAt: ts(a.CreatedAt),
			UpdatedAt: ts(a.UpdatedAt),
		})
	}
	for _, b := range p.Backgrounds {
		out.Backgrounds = append(out.Backgrounds, &accountpb.AuthorityProfileBackground{
			Id:        b.ID.String(),
			ProfileId: b.ProfileID.String(),
			ImageId:   b.ImageID.String(),
			SortOrder: b.SortOrder,
			CreatedAt: ts(b.CreatedAt),
			UpdatedAt: ts(b.UpdatedAt),
		})
	}
	return out
}

func ts(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

func accountStatusToProto(s enums.AuthAccountStatus) accountpb.AccountStatus {
	switch s {
	case enums.AuthAccountStatusActive:
		return accountpb.AccountStatus_ACCOUNT_STATUS_ACTIVE
	case enums.AuthAccountStatusSuspended:
		return accountpb.AccountStatus_ACCOUNT_STATUS_SUSPENDED
	case enums.AuthAccountStatusDeleted:
		return accountpb.AccountStatus_ACCOUNT_STATUS_DELETED
	default:
		return accountpb.AccountStatus_ACCOUNT_STATUS_UNSPECIFIED
	}
}

func profileLangToProto(s enums.AuthProfileLanguage) accountpb.ProfileLanguage {
	switch s {
	case enums.AuthProfileLanguageEn:
		return accountpb.ProfileLanguage_PROFILE_LANGUAGE_EN
	case enums.AuthProfileLanguageZh:
		return accountpb.ProfileLanguage_PROFILE_LANGUAGE_ZH
	case enums.AuthProfileLanguageFr:
		return accountpb.ProfileLanguage_PROFILE_LANGUAGE_FR
	default:
		return accountpb.ProfileLanguage_PROFILE_LANGUAGE_UNSPECIFIED
	}
}
