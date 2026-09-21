package mapper

import (
	"nfxidentity/enums"
	rdbviews "nfxidentity/modules/auth/infrastructure/rdb/views"
	accountQuery "nfxidentity/modules/auth/query/account"
	"nfxidentity/pkgs/ptrx"
)

// FullAccountInformationWithCommunityProfile（当前登录态单 profile 读模型）专用 mapper：
// 把 FullAccountInformationWithCommunityProfile* view 行映射到通用 VO。

func FullAccountInformationWithForgerProfileAccountViewToVO(v *rdbviews.FullAccountInformationWithForgerProfileAccountView) accountQuery.AccountVO {
	if v == nil {
		return accountQuery.AccountVO{}
	}
	return accountQuery.AccountVO{
		ID:             ptrx.Deref(v.AccountID),
		AccountStatus:  ptrx.Deref(v.AccountStatus),
		SignupPlatform: ptrx.Deref(v.SignupPlatform),
		CreatedAt:      ptrx.Deref(v.CreatedAt),
		UpdatedAt:      ptrx.Deref(v.UpdatedAt),
	}
}

func FullAccountInformationWithForgerProfileEmailViewToVO(v *rdbviews.FullAccountInformationWithForgerProfileEmailView) accountQuery.EmailVO {
	if v == nil {
		return accountQuery.EmailVO{}
	}
	return accountQuery.EmailVO{
		ID:         ptrx.Deref(v.ID),
		AccountID:  ptrx.Deref(v.AccountID),
		Email:      ptrx.Deref(v.Email),
		IsPrimary:  ptrx.Deref(v.IsPrimary),
		VerifiedAt: v.VerifiedAt,
		CreatedAt:  ptrx.Deref(v.CreatedAt),
		UpdatedAt:  ptrx.Deref(v.UpdatedAt),
	}
}

func FullAccountInformationWithForgerProfilePhoneViewToVO(v *rdbviews.FullAccountInformationWithForgerProfilePhoneView) accountQuery.PhoneVO {
	if v == nil {
		return accountQuery.PhoneVO{}
	}
	return accountQuery.PhoneVO{
		ID:         ptrx.Deref(v.ID),
		AccountID:  ptrx.Deref(v.AccountID),
		Phone:      ptrx.Deref(v.Phone),
		IsPrimary:  ptrx.Deref(v.IsPrimary),
		VerifiedAt: v.VerifiedAt,
		CreatedAt:  ptrx.Deref(v.CreatedAt),
		UpdatedAt:  ptrx.Deref(v.UpdatedAt),
	}
}

func FullAccountInformationWithForgerProfileViewToVO(v *rdbviews.FullAccountInformationWithForgerProfileView) accountQuery.CommunityProfileVO {
	if v == nil {
		return accountQuery.CommunityProfileVO{}
	}
	return accountQuery.CommunityProfileVO{
		ProfileID:       ptrx.Deref(v.ID),
		AccountID:       ptrx.Deref(v.AccountID),
		ForgerRoles:  append([]enums.AuthForgerRole(nil), v.ForgerRoles...),
		ProfileLanguage: ptrx.Deref(v.ProfileLanguage),
		Preference:      v.Preference,
		DisplayName:     v.DisplayName,
		FirstName:       v.FirstName,
		LastName:        v.LastName,
		Country:         v.Country,
		City:            v.City,
		Gender:          v.Gender,
		Birthday:        v.Birthday,
		Website:         v.Website,
		Timezone:        v.Timezone,
		Bio:             v.Bio,
		CreatedAt:       ptrx.Deref(v.CreatedAt),
		UpdatedAt:       ptrx.Deref(v.UpdatedAt),
	}
}

func FullAccountInformationWithForgerProfileAvatarViewToVO(
	v *rdbviews.FullAccountInformationWithForgerProfileAvatarView,
) accountQuery.CommunityProfileAvatarVO {
	if v == nil {
		return accountQuery.CommunityProfileAvatarVO{}
	}
	return accountQuery.CommunityProfileAvatarVO{
		ID:        ptrx.Deref(v.ID),
		ProfileID: ptrx.Deref(v.ProfileID),
		ImageID:   ptrx.Deref(v.ImageID),
		IsActive:  ptrx.Deref(v.IsActive),
		CreatedAt: ptrx.Deref(v.CreatedAt),
		UpdatedAt: ptrx.Deref(v.UpdatedAt),
	}
}

func FullAccountInformationWithForgerProfileBackgroundViewToVO(
	v *rdbviews.FullAccountInformationWithForgerProfileBackgroundView,
) accountQuery.CommunityProfileBackgroundVO {
	if v == nil {
		return accountQuery.CommunityProfileBackgroundVO{}
	}
	return accountQuery.CommunityProfileBackgroundVO{
		ID:        ptrx.Deref(v.ID),
		ProfileID: ptrx.Deref(v.ProfileID),
		ImageID:   ptrx.Deref(v.ImageID),
		SortOrder: int32(ptrx.Deref(v.SortOrder)),
		CreatedAt: ptrx.Deref(v.CreatedAt),
		UpdatedAt: ptrx.Deref(v.UpdatedAt),
	}
}

func FullAccountInformationWithForgerProfileSettingsViewToVO(
	v *rdbviews.FullAccountInformationWithForgerProfileSettingsView,
) *accountQuery.CommunityProfileSettingsVO {
	if v == nil || v.ID == nil {
		return nil
	}
	return &accountQuery.CommunityProfileSettingsVO{
		LoginNotification: ptrx.Deref(v.LoginNotification),
		CreatedAt:         ptrx.Deref(v.CreatedAt),
		UpdatedAt:         ptrx.Deref(v.UpdatedAt),
	}
}
