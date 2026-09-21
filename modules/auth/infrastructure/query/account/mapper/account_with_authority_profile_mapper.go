package mapper

import (
	"nfxidentity/enums"
	rdbviews "nfxidentity/modules/auth/infrastructure/rdb/views"
	accountQuery "nfxidentity/modules/auth/query/account"
	"nfxidentity/pkgs/ptrx"
)

// FullAccountInformationWithAuthorityProfile（当前登录态单 profile 读模型）专用 mapper：
// 把 FullAccountInformationWithAuthorityProfile* view 行映射到通用 VO。

func FullAccountInformationWithAuthorityProfileAccountViewToVO(v *rdbviews.FullAccountInformationWithAuthorityProfileAccountView) accountQuery.AccountVO {
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

func FullAccountInformationWithAuthorityProfileEmailViewToVO(v *rdbviews.FullAccountInformationWithAuthorityProfileEmailView) accountQuery.EmailVO {
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

func FullAccountInformationWithAuthorityProfilePhoneViewToVO(v *rdbviews.FullAccountInformationWithAuthorityProfilePhoneView) accountQuery.PhoneVO {
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

func FullAccountInformationWithAuthorityProfileViewToVO(v *rdbviews.FullAccountInformationWithAuthorityProfileView) accountQuery.AuthorityProfileVO {
	if v == nil {
		return accountQuery.AuthorityProfileVO{}
	}
	return accountQuery.AuthorityProfileVO{
		ProfileID:       ptrx.Deref(v.ID),
		AccountID:       ptrx.Deref(v.AccountID),
		AuthorityRoles:  append([]enums.AuthAuthorityRole(nil), v.AuthorityRoles...),
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

func FullAccountInformationWithAuthorityProfileAvatarViewToVO(
	v *rdbviews.FullAccountInformationWithAuthorityProfileAvatarView,
) accountQuery.AuthorityProfileAvatarVO {
	if v == nil {
		return accountQuery.AuthorityProfileAvatarVO{}
	}
	return accountQuery.AuthorityProfileAvatarVO{
		ID:        ptrx.Deref(v.ID),
		ProfileID: ptrx.Deref(v.ProfileID),
		ImageID:   ptrx.Deref(v.ImageID),
		IsActive:  ptrx.Deref(v.IsActive),
		CreatedAt: ptrx.Deref(v.CreatedAt),
		UpdatedAt: ptrx.Deref(v.UpdatedAt),
	}
}

func FullAccountInformationWithAuthorityProfileBackgroundViewToVO(
	v *rdbviews.FullAccountInformationWithAuthorityProfileBackgroundView,
) accountQuery.AuthorityProfileBackgroundVO {
	if v == nil {
		return accountQuery.AuthorityProfileBackgroundVO{}
	}
	return accountQuery.AuthorityProfileBackgroundVO{
		ID:        ptrx.Deref(v.ID),
		ProfileID: ptrx.Deref(v.ProfileID),
		ImageID:   ptrx.Deref(v.ImageID),
		SortOrder: int32(ptrx.Deref(v.SortOrder)),
		CreatedAt: ptrx.Deref(v.CreatedAt),
		UpdatedAt: ptrx.Deref(v.UpdatedAt),
	}
}

func FullAccountInformationWithAuthorityProfileSettingsViewToVO(
	v *rdbviews.FullAccountInformationWithAuthorityProfileSettingsView,
) *accountQuery.AuthorityProfileSettingsVO {
	if v == nil || v.ID == nil {
		return nil
	}
	return &accountQuery.AuthorityProfileSettingsVO{
		LoginNotification: ptrx.Deref(v.LoginNotification),
		CreatedAt:         ptrx.Deref(v.CreatedAt),
		UpdatedAt:         ptrx.Deref(v.UpdatedAt),
	}
}
