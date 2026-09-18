package profile

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	"time"
)

// * =============================== Authority profile factory =============================== !//
type NewAuthorityProfileParams struct {
	AccountID       uuid.UUID
	AuthorityRoles  []enums.AuthAuthorityRole
	ProfileLanguage enums.AuthProfileLanguage
	Preference      *datatypes.JSON
	DisplayName     *string
	FirstName       *string
	LastName        *string
	Country         *string
	City            *string
	Gender          *string
	Birthday        *time.Time
	Website         *string
	Timezone        *string
	Bio             *string
}

func NewAuthorityProfile(p NewAuthorityProfileParams) (*AuthorityProfile, error) {
	if p.AccountID == uuid.Nil {
		return nil, authErr.ErrAuthorityProfileAccountIDInvalid
	}
	lang := p.ProfileLanguage
	if lang == "" {
		lang = enums.AuthProfileLanguageEn
	}
	if err := validateAuthorityProfileLanguage(lang); err != nil {
		return nil, err
	}
	roles := p.AuthorityRoles
	if len(roles) == 0 {
		roles = []enums.AuthAuthorityRole{enums.AuthAuthorityRoleAdministrator}
	}
	normalized, err := normalizeAuthorityRoles(roles)
	if err != nil {
		return nil, err
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	return NewAuthorityProfileFromState(AuthorityProfileState{
		ID:              id,
		AccountID:       p.AccountID,
		AuthorityRoles:  normalized,
		ProfileLanguage: lang,
		Preference:      p.Preference,
		DisplayName:     p.DisplayName,
		FirstName:       p.FirstName,
		LastName:        p.LastName,
		Country:         p.Country,
		City:            p.City,
		Gender:          p.Gender,
		Birthday:        p.Birthday,
		Website:         p.Website,
		Timezone:        p.Timezone,
		Bio:             p.Bio,
		CreatedAt:       now,
		UpdatedAt:       now,
	}), nil
}

//! ====================================================================================== !//

type NewAuthorityProfileAvatarParams struct {
	ProfileID uuid.UUID
	ImageID   uuid.UUID
	IsActive  bool
}

func NewAuthorityProfileAvatar(p NewAuthorityProfileAvatarParams) (*AuthorityProfileAvatar, error) {
	if p.ProfileID == uuid.Nil {
		return nil, authErr.ErrAuthorityProfileAvatarAccountIDInvalid
	}
	if p.ImageID == uuid.Nil {
		return nil, authErr.ErrAuthorityProfileAvatarImageIDInvalid
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	return NewAuthorityProfileAvatarFromState(AuthorityProfileAvatarState{
		ID:        id,
		ProfileID: p.ProfileID,
		ImageID:   p.ImageID,
		IsActive:  p.IsActive,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

//! ====================================================================================== !//

type NewAuthorityProfileBackgroundParams struct {
	ProfileID uuid.UUID
	ImageID   uuid.UUID
	SortOrder int
}

func NewAuthorityProfileBackground(p NewAuthorityProfileBackgroundParams) (*AuthorityProfileBackground, error) {
	if p.ProfileID == uuid.Nil {
		return nil, authErr.ErrAuthorityProfileBackgroundAccountIDInvalid
	}
	if p.ImageID == uuid.Nil {
		return nil, authErr.ErrAuthorityProfileBackgroundImageIDInvalid
	}
	if p.SortOrder < 0 {
		return nil, authErr.ErrAuthorityProfileBackgroundSortOrderInvalid
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	return NewAuthorityProfileBackgroundFromState(AuthorityProfileBackgroundState{
		ID:        id,
		ProfileID: p.ProfileID,
		ImageID:   p.ImageID,
		SortOrder: p.SortOrder,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

//! ====================================================================================== !//

type NewAuthorityProfileSettingsParams struct {
	ProfileID         uuid.UUID
	LoginNotification bool
}

func NewAuthorityProfileSettings(p NewAuthorityProfileSettingsParams) (*AuthorityProfileSettings, error) {
	if p.ProfileID == uuid.Nil {
		return nil, authErr.ErrAuthorityProfileSettingsProfileIDInvalid
	}
	now := time.Now().UTC()
	return NewAuthorityProfileSettingsFromState(AuthorityProfileSettingsState{
		ID:                p.ProfileID,
		LoginNotification: p.LoginNotification,
		CreatedAt:         now,
		UpdatedAt:         now,
	}), nil
}

// * =============================== Forger profile factory =============================== !//
type NewForgerProfileParams struct {
	AccountID       uuid.UUID
	ForgerRoles     []enums.AuthForgerRole
	ProfileLanguage enums.AuthProfileLanguage
	Preference      *datatypes.JSON
	DisplayName     *string
	FirstName       *string
	LastName        *string
	Country         *string
	City            *string
	Gender          *string
	Birthday        *time.Time
	Website         *string
	Timezone        *string
	Bio             *string
}

func NewForgerProfile(p NewForgerProfileParams) (*ForgerProfile, error) {
	if p.AccountID == uuid.Nil {
		return nil, authErr.ErrForgerProfileAccountIDInvalid
	}
	lang := p.ProfileLanguage
	if lang == "" {
		lang = enums.AuthProfileLanguageEn
	}
	if err := validateForgerProfileLanguage(lang); err != nil {
		return nil, err
	}
	roles := p.ForgerRoles
	if len(roles) == 0 {
		roles = []enums.AuthForgerRole{enums.AuthForgerRoleForger}
	}
	normalized, err := normalizeForgerRoles(roles)
	if err != nil {
		return nil, err
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	return NewForgerProfileFromState(ForgerProfileState{
		ID:              id,
		AccountID:       p.AccountID,
		ForgerRoles:     normalized,
		ProfileLanguage: lang,
		Preference:      p.Preference,
		DisplayName:     p.DisplayName,
		FirstName:       p.FirstName,
		LastName:        p.LastName,
		Country:         p.Country,
		City:            p.City,
		Gender:          p.Gender,
		Birthday:        p.Birthday,
		Website:         p.Website,
		Timezone:        p.Timezone,
		Bio:             p.Bio,
		CreatedAt:       now,
		UpdatedAt:       now,
	}), nil
}

//! ====================================================================================== !//

type NewForgerProfileAvatarParams struct {
	ProfileID uuid.UUID
	ImageID   uuid.UUID
	IsActive  bool
}

func NewForgerProfileAvatar(p NewForgerProfileAvatarParams) (*ForgerProfileAvatar, error) {
	if p.ProfileID == uuid.Nil {
		return nil, authErr.ErrForgerProfileAvatarAccountIDInvalid
	}
	if p.ImageID == uuid.Nil {
		return nil, authErr.ErrForgerProfileAvatarImageIDInvalid
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	return NewForgerProfileAvatarFromState(ForgerProfileAvatarState{
		ID:        id,
		ProfileID: p.ProfileID,
		ImageID:   p.ImageID,
		IsActive:  p.IsActive,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

//! ====================================================================================== !//

type NewForgerProfileBackgroundParams struct {
	ProfileID uuid.UUID
	ImageID   uuid.UUID
	SortOrder int
}

func NewForgerProfileBackground(p NewForgerProfileBackgroundParams) (*ForgerProfileBackground, error) {
	if p.ProfileID == uuid.Nil {
		return nil, authErr.ErrForgerProfileBackgroundAccountIDInvalid
	}
	if p.ImageID == uuid.Nil {
		return nil, authErr.ErrForgerProfileBackgroundImageIDInvalid
	}
	if p.SortOrder < 0 {
		return nil, authErr.ErrForgerProfileBackgroundSortOrderInvalid
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	return NewForgerProfileBackgroundFromState(ForgerProfileBackgroundState{
		ID:        id,
		ProfileID: p.ProfileID,
		ImageID:   p.ImageID,
		SortOrder: p.SortOrder,
		CreatedAt: now,
		UpdatedAt: now,
	}), nil
}

//! ====================================================================================== !//

type NewForgerProfileSettingsParams struct {
	ProfileID         uuid.UUID
	LoginNotification bool
}

func NewForgerProfileSettings(p NewForgerProfileSettingsParams) (*ForgerProfileSettings, error) {
	if p.ProfileID == uuid.Nil {
		return nil, authErr.ErrForgerProfileSettingsProfileIDInvalid
	}
	now := time.Now().UTC()
	return NewForgerProfileSettingsFromState(ForgerProfileSettingsState{
		ID:                p.ProfileID,
		LoginNotification: p.LoginNotification,
		CreatedAt:         now,
		UpdatedAt:         now,
	}), nil
}
