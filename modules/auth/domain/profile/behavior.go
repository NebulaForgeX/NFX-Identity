package profile

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	"time"
)

// * =============================== Authority profile behavior =============================== !//
type AuthorityProfileEditable struct {
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

func (u *AuthorityProfile) EnsureNotDeleted() error {
	if u.DeletedAt() != nil {
		return authErr.ErrAuthorityProfileNotFound
	}
	return nil
}

func (u *AuthorityProfile) Update(ed AuthorityProfileEditable) error {
	if err := u.EnsureNotDeleted(); err != nil {
		return err
	}
	if err := ed.Validate(); err != nil {
		return err
	}
	u.state.ProfileLanguage = ed.ProfileLanguage
	u.state.Preference = ed.Preference
	u.state.DisplayName = ed.DisplayName
	u.state.FirstName = ed.FirstName
	u.state.LastName = ed.LastName
	u.state.Country = ed.Country
	u.state.City = ed.City
	u.state.Gender = ed.Gender
	u.state.Birthday = ed.Birthday
	u.state.Website = ed.Website
	u.state.Timezone = ed.Timezone
	u.state.Bio = ed.Bio
	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *AuthorityProfile) SetAuthorityRoles(roles []enums.AuthAuthorityRole) error {
	if err := u.EnsureNotDeleted(); err != nil {
		return err
	}
	normalized, err := normalizeAuthorityRoles(roles)
	if err != nil {
		return err
	}
	u.state.AuthorityRoles = normalized
	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *AuthorityProfile) Delete() error {
	if u.DeletedAt() != nil {
		return nil
	}
	now := time.Now().UTC()
	u.state.DeletedAt = &now
	u.state.UpdatedAt = now
	return nil
}

//! ====================================================================================== !//

// AuthorityProfileAvatarEditable 可编辑字段（nil 表示不修改）。
type AuthorityProfileAvatarEditable struct {
	ImageID  *uuid.UUID
	IsActive *bool
}

func (a *AuthorityProfileAvatar) Update(ed AuthorityProfileAvatarEditable) error {
	if ed.ImageID != nil {
		if *ed.ImageID == uuid.Nil {
			return authErr.ErrAuthorityProfileAvatarImageIDInvalid
		}
		a.state.ImageID = *ed.ImageID
	}
	if ed.IsActive != nil {
		a.state.IsActive = *ed.IsActive
	}
	a.state.UpdatedAt = time.Now().UTC()
	return nil
}

// Delete 软删头像关联行。
func (a *AuthorityProfileAvatar) Delete() error {
	if a.state.DeletedAt != nil {
		return nil
	}
	now := time.Now().UTC()
	a.state.DeletedAt = &now
	a.state.UpdatedAt = now
	return nil
}

//! ====================================================================================== !//

// AuthorityProfileBackgroundEditable 可编辑字段（nil 表示不修改）。
type AuthorityProfileBackgroundEditable struct {
	ImageID   *uuid.UUID
	SortOrder *int
}

func (b *AuthorityProfileBackground) Update(ed AuthorityProfileBackgroundEditable) error {
	if ed.ImageID != nil {
		if *ed.ImageID == uuid.Nil {
			return authErr.ErrAuthorityProfileBackgroundImageIDInvalid
		}
		b.state.ImageID = *ed.ImageID
	}
	if ed.SortOrder != nil {
		if *ed.SortOrder < 0 {
			return authErr.ErrAuthorityProfileBackgroundSortOrderInvalid
		}
		b.state.SortOrder = *ed.SortOrder
	}
	b.state.UpdatedAt = time.Now().UTC()
	return nil
}

// Delete 软删背景图关联行。
func (b *AuthorityProfileBackground) Delete() error {
	if b.state.DeletedAt != nil {
		return nil
	}
	now := time.Now().UTC()
	b.state.DeletedAt = &now
	b.state.UpdatedAt = now
	return nil
}

//! ====================================================================================== !//

func (s *AuthorityProfileSettings) EnsureNotDeleted() error {
	if s.DeletedAt() != nil {
		return authErr.ErrAuthorityProfileSettingsNotFound
	}
	return nil
}

// * =============================== Forger profile behavior =============================== !//
type ForgerProfileEditable struct {
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

func (u *ForgerProfile) EnsureNotDeleted() error {
	if u.DeletedAt() != nil {
		return authErr.ErrForgerProfileNotFound
	}
	return nil
}

func (u *ForgerProfile) Update(ed ForgerProfileEditable) error {
	if err := u.EnsureNotDeleted(); err != nil {
		return err
	}
	if err := ed.Validate(); err != nil {
		return err
	}
	u.state.ProfileLanguage = ed.ProfileLanguage
	u.state.Preference = ed.Preference
	u.state.DisplayName = ed.DisplayName
	u.state.FirstName = ed.FirstName
	u.state.LastName = ed.LastName
	u.state.Country = ed.Country
	u.state.City = ed.City
	u.state.Gender = ed.Gender
	u.state.Birthday = ed.Birthday
	u.state.Website = ed.Website
	u.state.Timezone = ed.Timezone
	u.state.Bio = ed.Bio
	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *ForgerProfile) SetForgerRoles(roles []enums.AuthForgerRole) error {
	if err := u.EnsureNotDeleted(); err != nil {
		return err
	}
	normalized, err := normalizeForgerRoles(roles)
	if err != nil {
		return err
	}
	u.state.ForgerRoles = normalized
	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *ForgerProfile) AddForgerRole(role enums.AuthForgerRole) error {
	if err := u.EnsureNotDeleted(); err != nil {
		return err
	}
	if err := validateForgerRole(role); err != nil {
		return err
	}
	for _, existing := range u.state.ForgerRoles {
		if existing == role {
			return nil
		}
	}
	u.state.ForgerRoles = append(append([]enums.AuthForgerRole(nil), u.state.ForgerRoles...), role)
	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *ForgerProfile) Delete() error {
	if u.DeletedAt() != nil {
		return nil
	}
	now := time.Now().UTC()
	u.state.DeletedAt = &now
	u.state.UpdatedAt = now
	return nil
}

//! ====================================================================================== !//

// ForgerProfileAvatarEditable 可编辑字段（nil 表示不修改）。
type ForgerProfileAvatarEditable struct {
	ImageID  *uuid.UUID
	IsActive *bool
}

func (a *ForgerProfileAvatar) Update(ed ForgerProfileAvatarEditable) error {
	if ed.ImageID != nil {
		if *ed.ImageID == uuid.Nil {
			return authErr.ErrForgerProfileAvatarImageIDInvalid
		}
		a.state.ImageID = *ed.ImageID
	}
	if ed.IsActive != nil {
		a.state.IsActive = *ed.IsActive
	}
	a.state.UpdatedAt = time.Now().UTC()
	return nil
}

// Delete 软删头像关联行。
func (a *ForgerProfileAvatar) Delete() error {
	if a.state.DeletedAt != nil {
		return nil
	}
	now := time.Now().UTC()
	a.state.DeletedAt = &now
	a.state.UpdatedAt = now
	return nil
}

//! ====================================================================================== !//

// ForgerProfileBackgroundEditable 可编辑字段（nil 表示不修改）。
type ForgerProfileBackgroundEditable struct {
	ImageID   *uuid.UUID
	SortOrder *int
}

func (b *ForgerProfileBackground) Update(ed ForgerProfileBackgroundEditable) error {
	if ed.ImageID != nil {
		if *ed.ImageID == uuid.Nil {
			return authErr.ErrForgerProfileBackgroundImageIDInvalid
		}
		b.state.ImageID = *ed.ImageID
	}
	if ed.SortOrder != nil {
		if *ed.SortOrder < 0 {
			return authErr.ErrForgerProfileBackgroundSortOrderInvalid
		}
		b.state.SortOrder = *ed.SortOrder
	}
	b.state.UpdatedAt = time.Now().UTC()
	return nil
}

// Delete 软删背景图关联行。
func (b *ForgerProfileBackground) Delete() error {
	if b.state.DeletedAt != nil {
		return nil
	}
	now := time.Now().UTC()
	b.state.DeletedAt = &now
	b.state.UpdatedAt = now
	return nil
}

//! ====================================================================================== !//

func (s *ForgerProfileSettings) EnsureNotDeleted() error {
	if s.DeletedAt() != nil {
		return authErr.ErrForgerProfileSettingsNotFound
	}
	return nil
}
