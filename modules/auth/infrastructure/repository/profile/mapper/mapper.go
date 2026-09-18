package mapper

import (
	"gorm.io/gorm"
	"nfxidentity/enums"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	rdbmodels "nfxidentity/modules/auth/infrastructure/rdb/models"
	"nfxidentity/pkgs/enumx"
	"nfxidentity/pkgs/patch"
	"nfxidentity/pkgs/ptrx"
)

//* =============================== Authority profile mapper =============================== !//
func AuthorityProfileDomainToModel(p *profileDomain.AuthorityProfile) *rdbmodels.AuthorityProfile {
	if p == nil {
		return nil
	}
	return &rdbmodels.AuthorityProfile{
		ID:              p.ID(),
		AccountID:       p.AccountID(),
		AuthorityRoles:  enumx.Array[enums.AuthAuthorityRole](p.AuthorityRoles()),
		ProfileLanguage: p.ProfileLanguage(),
		Preference:      p.Preference(),
		DisplayName:     p.DisplayName(),
		FirstName:       p.FirstName(),
		LastName:        p.LastName(),
		Country:         p.Country(),
		City:            p.City(),
		Gender:          p.Gender(),
		Birthday:        p.Birthday(),
		Website:         p.Website(),
		Timezone:        p.Timezone(),
		Bio:             p.Bio(),
		CreatedAt:       p.CreatedAt(),
		UpdatedAt:       p.UpdatedAt(),
		DeletedAt:       ptrx.TimePtrToDeletedAt(p.DeletedAt()),
	}
}

func AuthorityProfileModelToDomain(m *rdbmodels.AuthorityProfile) *profileDomain.AuthorityProfile {
	if m == nil {
		return nil
	}
	return profileDomain.NewAuthorityProfileFromState(profileDomain.AuthorityProfileState{
		ID:              m.ID,
		AccountID:       m.AccountID,
		AuthorityRoles:  append([]enums.AuthAuthorityRole(nil), m.AuthorityRoles...),
		ProfileLanguage: m.ProfileLanguage,
		Preference:      m.Preference,
		DisplayName:     m.DisplayName,
		FirstName:       m.FirstName,
		LastName:        m.LastName,
		Country:         m.Country,
		City:            m.City,
		Gender:          m.Gender,
		Birthday:        m.Birthday,
		Website:         m.Website,
		Timezone:        m.Timezone,
		Bio:             m.Bio,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		DeletedAt:       ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

var authorityProfilePatchColumnOpts = patch.Options{
	FieldToColumn: map[string]string{
		"ProfileLanguage": rdbmodels.AuthorityProfileCols.ProfileLanguage,
		"DisplayName":     rdbmodels.AuthorityProfileCols.DisplayName,
		"FirstName":       rdbmodels.AuthorityProfileCols.FirstName,
		"LastName":        rdbmodels.AuthorityProfileCols.LastName,
		"Country":         rdbmodels.AuthorityProfileCols.Country,
		"City":            rdbmodels.AuthorityProfileCols.City,
		"Gender":          rdbmodels.AuthorityProfileCols.Gender,
		"Birthday":        rdbmodels.AuthorityProfileCols.Birthday,
		"Website":         rdbmodels.AuthorityProfileCols.Website,
		"Timezone":        rdbmodels.AuthorityProfileCols.Timezone,
		"Bio":             rdbmodels.AuthorityProfileCols.Bio,
	},
	FallbackSnakeCase: false,
	NullValue:         func() any { return gorm.Expr("NULL") },
}

func AuthorityProfilePatchToUpdates(p profileDomain.AuthorityProfilePatch) map[string]any {
	return patch.PatchToColumns(p, authorityProfilePatchColumnOpts)
}

func AuthorityProfileDomainToUpdates(p *profileDomain.AuthorityProfile) map[string]any {
	m := AuthorityProfileDomainToModel(p)
	return map[string]any{
		rdbmodels.AuthorityProfileCols.AuthorityRoles:  m.AuthorityRoles,
		rdbmodels.AuthorityProfileCols.ProfileLanguage: m.ProfileLanguage,
		rdbmodels.AuthorityProfileCols.Preference:      m.Preference,
		rdbmodels.AuthorityProfileCols.DisplayName:     m.DisplayName,
		rdbmodels.AuthorityProfileCols.FirstName:       m.FirstName,
		rdbmodels.AuthorityProfileCols.LastName:        m.LastName,
		rdbmodels.AuthorityProfileCols.Country:         m.Country,
		rdbmodels.AuthorityProfileCols.City:            m.City,
		rdbmodels.AuthorityProfileCols.Gender:          m.Gender,
		rdbmodels.AuthorityProfileCols.Birthday:        m.Birthday,
		rdbmodels.AuthorityProfileCols.Website:         m.Website,
		rdbmodels.AuthorityProfileCols.Timezone:        m.Timezone,
		rdbmodels.AuthorityProfileCols.Bio:             m.Bio,
		rdbmodels.AuthorityProfileCols.UpdatedAt:       m.UpdatedAt,
		rdbmodels.AuthorityProfileCols.DeletedAt:       m.DeletedAt,
	}
}

func AuthorityProfileAvatarDomainToModel(a *profileDomain.AuthorityProfileAvatar) *rdbmodels.AuthorityProfileAvatar {
	if a == nil {
		return nil
	}
	st := a.State()
	return &rdbmodels.AuthorityProfileAvatar{
		ID:        st.ID,
		ProfileID: st.ProfileID,
		ImageID:   st.ImageID,
		IsActive:  st.IsActive,
		CreatedAt: st.CreatedAt,
		UpdatedAt: st.UpdatedAt,
		DeletedAt: ptrx.TimePtrToDeletedAt(st.DeletedAt),
	}
}

func AuthorityProfileAvatarModelToDomain(m *rdbmodels.AuthorityProfileAvatar) *profileDomain.AuthorityProfileAvatar {
	if m == nil {
		return nil
	}
	return profileDomain.NewAuthorityProfileAvatarFromState(profileDomain.AuthorityProfileAvatarState{
		ID:        m.ID,
		ProfileID: m.ProfileID,
		ImageID:   m.ImageID,
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

func AuthorityProfileBackgroundDomainToModel(b *profileDomain.AuthorityProfileBackground) *rdbmodels.AuthorityProfileBackground {
	if b == nil {
		return nil
	}
	st := b.State()
	return &rdbmodels.AuthorityProfileBackground{
		ID:        st.ID,
		ProfileID: st.ProfileID,
		ImageID:   st.ImageID,
		SortOrder: st.SortOrder,
		CreatedAt: st.CreatedAt,
		UpdatedAt: st.UpdatedAt,
		DeletedAt: ptrx.TimePtrToDeletedAt(st.DeletedAt),
	}
}

func AuthorityProfileBackgroundModelToDomain(m *rdbmodels.AuthorityProfileBackground) *profileDomain.AuthorityProfileBackground {
	if m == nil {
		return nil
	}
	return profileDomain.NewAuthorityProfileBackgroundFromState(profileDomain.AuthorityProfileBackgroundState{
		ID:        m.ID,
		ProfileID: m.ProfileID,
		ImageID:   m.ImageID,
		SortOrder: m.SortOrder,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

func AuthorityProfileSettingsDomainToModel(s *profileDomain.AuthorityProfileSettings) *rdbmodels.AuthorityProfileSetting {
	if s == nil {
		return nil
	}
	return &rdbmodels.AuthorityProfileSetting{
		ID:                s.ID(),
		LoginNotification: s.LoginNotification(),
		CreatedAt:         s.CreatedAt(),
		UpdatedAt:         s.UpdatedAt(),
		DeletedAt:         ptrx.TimePtrToDeletedAt(s.DeletedAt()),
	}
}

func AuthorityProfileSettingsModelToDomain(m *rdbmodels.AuthorityProfileSetting) *profileDomain.AuthorityProfileSettings {
	if m == nil {
		return nil
	}
	return profileDomain.NewAuthorityProfileSettingsFromState(profileDomain.AuthorityProfileSettingsState{
		ID:                m.ID,
		LoginNotification: m.LoginNotification,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
		DeletedAt:         ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

var authorityProfileSettingsPatchColumnOpts = patch.Options{
	FieldToColumn: map[string]string{
		"LoginNotification": rdbmodels.AuthorityProfileSettingCols.LoginNotification,
	},
	FallbackSnakeCase: false,
	NullValue:         func() any { return gorm.Expr("NULL") },
}

func AuthorityProfileSettingsPatchToUpdates(p profileDomain.AuthorityProfileSettingsPatch) map[string]any {
	return patch.PatchToColumns(p, authorityProfileSettingsPatchColumnOpts)
}

//* =============================== Forger profile mapper =============================== !//
func ForgerProfileDomainToModel(p *profileDomain.ForgerProfile) *rdbmodels.ForgerProfile {
	if p == nil {
		return nil
	}
	return &rdbmodels.ForgerProfile{
		ID:              p.ID(),
		AccountID:       p.AccountID(),
		ForgerRoles:  enumx.Array[enums.AuthForgerRole](p.ForgerRoles()),
		ProfileLanguage: p.ProfileLanguage(),
		Preference:      p.Preference(),
		DisplayName:     p.DisplayName(),
		FirstName:       p.FirstName(),
		LastName:        p.LastName(),
		Country:         p.Country(),
		City:            p.City(),
		Gender:          p.Gender(),
		Birthday:        p.Birthday(),
		Website:         p.Website(),
		Timezone:        p.Timezone(),
		Bio:             p.Bio(),
		CreatedAt:       p.CreatedAt(),
		UpdatedAt:       p.UpdatedAt(),
		DeletedAt:       ptrx.TimePtrToDeletedAt(p.DeletedAt()),
	}
}

func ForgerProfileModelToDomain(m *rdbmodels.ForgerProfile) *profileDomain.ForgerProfile {
	if m == nil {
		return nil
	}
	return profileDomain.NewForgerProfileFromState(profileDomain.ForgerProfileState{
		ID:              m.ID,
		AccountID:       m.AccountID,
		ForgerRoles:  append([]enums.AuthForgerRole(nil), m.ForgerRoles...),
		ProfileLanguage: m.ProfileLanguage,
		Preference:      m.Preference,
		DisplayName:     m.DisplayName,
		FirstName:       m.FirstName,
		LastName:        m.LastName,
		Country:         m.Country,
		City:            m.City,
		Gender:          m.Gender,
		Birthday:        m.Birthday,
		Website:         m.Website,
		Timezone:        m.Timezone,
		Bio:             m.Bio,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		DeletedAt:       ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

var forgerProfilePatchColumnOpts = patch.Options{
	FieldToColumn: map[string]string{
		"ProfileLanguage": rdbmodels.ForgerProfileCols.ProfileLanguage,
		"DisplayName":     rdbmodels.ForgerProfileCols.DisplayName,
		"FirstName":       rdbmodels.ForgerProfileCols.FirstName,
		"LastName":        rdbmodels.ForgerProfileCols.LastName,
		"Country":         rdbmodels.ForgerProfileCols.Country,
		"City":            rdbmodels.ForgerProfileCols.City,
		"Gender":          rdbmodels.ForgerProfileCols.Gender,
		"Birthday":        rdbmodels.ForgerProfileCols.Birthday,
		"Website":         rdbmodels.ForgerProfileCols.Website,
		"Timezone":        rdbmodels.ForgerProfileCols.Timezone,
		"Bio":             rdbmodels.ForgerProfileCols.Bio,
	},
	FallbackSnakeCase: false,
	NullValue:         func() any { return gorm.Expr("NULL") },
}

func ForgerProfilePatchToUpdates(p profileDomain.ForgerProfilePatch) map[string]any {
	return patch.PatchToColumns(p, forgerProfilePatchColumnOpts)
}

func ForgerProfileDomainToUpdates(p *profileDomain.ForgerProfile) map[string]any {
	m := ForgerProfileDomainToModel(p)
	return map[string]any{
		rdbmodels.ForgerProfileCols.ForgerRoles:  m.ForgerRoles,
		rdbmodels.ForgerProfileCols.ProfileLanguage: m.ProfileLanguage,
		rdbmodels.ForgerProfileCols.Preference:      m.Preference,
		rdbmodels.ForgerProfileCols.DisplayName:     m.DisplayName,
		rdbmodels.ForgerProfileCols.FirstName:       m.FirstName,
		rdbmodels.ForgerProfileCols.LastName:        m.LastName,
		rdbmodels.ForgerProfileCols.Country:         m.Country,
		rdbmodels.ForgerProfileCols.City:            m.City,
		rdbmodels.ForgerProfileCols.Gender:          m.Gender,
		rdbmodels.ForgerProfileCols.Birthday:        m.Birthday,
		rdbmodels.ForgerProfileCols.Website:         m.Website,
		rdbmodels.ForgerProfileCols.Timezone:        m.Timezone,
		rdbmodels.ForgerProfileCols.Bio:             m.Bio,
		rdbmodels.ForgerProfileCols.UpdatedAt:       m.UpdatedAt,
		rdbmodels.ForgerProfileCols.DeletedAt:       m.DeletedAt,
	}
}

func ForgerProfileAvatarDomainToModel(a *profileDomain.ForgerProfileAvatar) *rdbmodels.ForgerProfileAvatar {
	if a == nil {
		return nil
	}
	st := a.State()
	return &rdbmodels.ForgerProfileAvatar{
		ID:        st.ID,
		ProfileID: st.ProfileID,
		ImageID:   st.ImageID,
		IsActive:  st.IsActive,
		CreatedAt: st.CreatedAt,
		UpdatedAt: st.UpdatedAt,
		DeletedAt: ptrx.TimePtrToDeletedAt(st.DeletedAt),
	}
}

func ForgerProfileAvatarModelToDomain(m *rdbmodels.ForgerProfileAvatar) *profileDomain.ForgerProfileAvatar {
	if m == nil {
		return nil
	}
	return profileDomain.NewForgerProfileAvatarFromState(profileDomain.ForgerProfileAvatarState{
		ID:        m.ID,
		ProfileID: m.ProfileID,
		ImageID:   m.ImageID,
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

func ForgerProfileBackgroundDomainToModel(b *profileDomain.ForgerProfileBackground) *rdbmodels.ForgerProfileBackground {
	if b == nil {
		return nil
	}
	st := b.State()
	return &rdbmodels.ForgerProfileBackground{
		ID:        st.ID,
		ProfileID: st.ProfileID,
		ImageID:   st.ImageID,
		SortOrder: st.SortOrder,
		CreatedAt: st.CreatedAt,
		UpdatedAt: st.UpdatedAt,
		DeletedAt: ptrx.TimePtrToDeletedAt(st.DeletedAt),
	}
}

func ForgerProfileBackgroundModelToDomain(m *rdbmodels.ForgerProfileBackground) *profileDomain.ForgerProfileBackground {
	if m == nil {
		return nil
	}
	return profileDomain.NewForgerProfileBackgroundFromState(profileDomain.ForgerProfileBackgroundState{
		ID:        m.ID,
		ProfileID: m.ProfileID,
		ImageID:   m.ImageID,
		SortOrder: m.SortOrder,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

func ForgerProfileSettingsDomainToModel(s *profileDomain.ForgerProfileSettings) *rdbmodels.ForgerProfileSetting {
	if s == nil {
		return nil
	}
	return &rdbmodels.ForgerProfileSetting{
		ID:                s.ID(),
		LoginNotification: s.LoginNotification(),
		CreatedAt:         s.CreatedAt(),
		UpdatedAt:         s.UpdatedAt(),
		DeletedAt:         ptrx.TimePtrToDeletedAt(s.DeletedAt()),
	}
}

func ForgerProfileSettingsModelToDomain(m *rdbmodels.ForgerProfileSetting) *profileDomain.ForgerProfileSettings {
	if m == nil {
		return nil
	}
	return profileDomain.NewForgerProfileSettingsFromState(profileDomain.ForgerProfileSettingsState{
		ID:                m.ID,
		LoginNotification: m.LoginNotification,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
		DeletedAt:         ptrx.DeletedAtToTimePtr(m.DeletedAt),
	})
}

var forgerProfileSettingsPatchColumnOpts = patch.Options{
	FieldToColumn: map[string]string{
		"LoginNotification": rdbmodels.ForgerProfileSettingCols.LoginNotification,
	},
	FallbackSnakeCase: false,
	NullValue:         func() any { return gorm.Expr("NULL") },
}

func ForgerProfileSettingsPatchToUpdates(p profileDomain.ForgerProfileSettingsPatch) map[string]any {
	return patch.PatchToColumns(p, forgerProfileSettingsPatchColumnOpts)
}
