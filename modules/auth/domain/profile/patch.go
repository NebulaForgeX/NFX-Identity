package profile

import (
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/patch"
	"strings"
	"time"
)

// * =============================== Shared patch helpers =============================== !//
func validatePatchStringField(f patch.PatchField[string], maxLen int, fieldErr *errx.Error) error {
	if !f.IsSet() || f.IsNull() {
		return nil
	}
	if v, ok := f.Value(); ok && len(v) > maxLen {
		return fieldErr
	}
	return nil
}

func diffStringPatchField(in patch.PatchField[string], old *string) patch.PatchField[string] {
	return in.DiffPtrBy(old, func(a, b string) bool {
		return strings.TrimSpace(a) == strings.TrimSpace(b)
	})
}

// * =============================== Authority profile patch =============================== !//
type AuthorityProfilePatch struct {
	ProfileLanguage patch.PatchField[enums.AuthProfileLanguage]
	DisplayName     patch.PatchField[string]
	FirstName       patch.PatchField[string]
	LastName        patch.PatchField[string]
	Country         patch.PatchField[string]
	City            patch.PatchField[string]
	Gender          patch.PatchField[string]
	Birthday        patch.PatchField[time.Time]
	Website         patch.PatchField[string]
	Timezone        patch.PatchField[string]
	Bio             patch.PatchField[string]
}

func (p AuthorityProfilePatch) Validate() error {
	if patch.IsPatchEmpty(p) {
		return authErr.ErrAuthorityProfilePatchEmpty
	}
	if p.ProfileLanguage.IsSet() {
		if p.ProfileLanguage.IsNull() {
			return authErr.ErrAuthorityProfileLanguageInvalid
		}
		if v, ok := p.ProfileLanguage.Value(); ok {
			if err := validateAuthorityProfileLanguage(v); err != nil {
				return err
			}
		}
	}
	if err := validatePatchStringField(p.DisplayName, 150, authErr.ErrAuthorityProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.FirstName, 100, authErr.ErrAuthorityProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.LastName, 100, authErr.ErrAuthorityProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.Country, 100, authErr.ErrAuthorityProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.City, 100, authErr.ErrAuthorityProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.Gender, 100, authErr.ErrAuthorityProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.Website, 100, authErr.ErrAuthorityProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.Timezone, 100, authErr.ErrAuthorityProfileFieldInvalid); err != nil {
		return err
	}
	if p.Bio.IsSet() && !p.Bio.IsNull() {
		if v, ok := p.Bio.Value(); ok && len(v) > 5000 {
			return authErr.ErrAuthorityProfileFieldInvalid
		}
	}
	return nil
}

// DiffAgainst keeps only incoming fields that differ from u; unchanged values become Unset.
func (p AuthorityProfilePatch) DiffAgainst(u *AuthorityProfile) AuthorityProfilePatch {
	return AuthorityProfilePatch{
		ProfileLanguage: p.ProfileLanguage.DiffVal(u.ProfileLanguage()),
		DisplayName:     diffStringPatchField(p.DisplayName, u.DisplayName()),
		FirstName:       diffStringPatchField(p.FirstName, u.FirstName()),
		LastName:        diffStringPatchField(p.LastName, u.LastName()),
		Country:         diffStringPatchField(p.Country, u.Country()),
		City:            diffStringPatchField(p.City, u.City()),
		Gender:          diffStringPatchField(p.Gender, u.Gender()),
		Birthday:        p.Birthday.DiffPtr(u.Birthday()),
		Website:         diffStringPatchField(p.Website, u.Website()),
		Timezone:        diffStringPatchField(p.Timezone, u.Timezone()),
		Bio:             diffStringPatchField(p.Bio, u.Bio()),
	}
}

func (u *AuthorityProfile) ApplyPatch(p AuthorityProfilePatch) error {
	if err := u.EnsureNotDeleted(); err != nil {
		return err
	}

	patch.ApplyVal(&u.state.ProfileLanguage, p.ProfileLanguage)
	patch.ApplyPtr(&u.state.DisplayName, p.DisplayName)
	patch.ApplyPtr(&u.state.FirstName, p.FirstName)
	patch.ApplyPtr(&u.state.LastName, p.LastName)
	patch.ApplyPtr(&u.state.Country, p.Country)
	patch.ApplyPtr(&u.state.City, p.City)
	patch.ApplyPtr(&u.state.Gender, p.Gender)
	patch.ApplyPtr(&u.state.Website, p.Website)
	patch.ApplyPtr(&u.state.Timezone, p.Timezone)
	patch.ApplyPtrWith(&u.state.Bio, p.Bio, strings.TrimSpace)
	patch.ApplyPtr(&u.state.Birthday, p.Birthday)

	u.state.UpdatedAt = time.Now().UTC()
	return nil
}

// * =============================== Forger profile patch =============================== !//
type ForgerProfilePatch struct {
	ProfileLanguage patch.PatchField[enums.AuthProfileLanguage]
	DisplayName     patch.PatchField[string]
	FirstName       patch.PatchField[string]
	LastName        patch.PatchField[string]
	Country         patch.PatchField[string]
	City            patch.PatchField[string]
	Gender          patch.PatchField[string]
	Birthday        patch.PatchField[time.Time]
	Website         patch.PatchField[string]
	Timezone        patch.PatchField[string]
	Bio             patch.PatchField[string]
}

func (p ForgerProfilePatch) Validate() error {
	if patch.IsPatchEmpty(p) {
		return authErr.ErrForgerProfilePatchEmpty
	}
	if p.ProfileLanguage.IsSet() {
		if p.ProfileLanguage.IsNull() {
			return authErr.ErrForgerProfileLanguageInvalid
		}
		if v, ok := p.ProfileLanguage.Value(); ok {
			if err := validateForgerProfileLanguage(v); err != nil {
				return err
			}
		}
	}
	if err := validatePatchStringField(p.DisplayName, 150, authErr.ErrForgerProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.FirstName, 100, authErr.ErrForgerProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.LastName, 100, authErr.ErrForgerProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.Country, 100, authErr.ErrForgerProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.City, 100, authErr.ErrForgerProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.Gender, 100, authErr.ErrForgerProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.Website, 100, authErr.ErrForgerProfileFieldInvalid); err != nil {
		return err
	}
	if err := validatePatchStringField(p.Timezone, 100, authErr.ErrForgerProfileFieldInvalid); err != nil {
		return err
	}
	if p.Bio.IsSet() && !p.Bio.IsNull() {
		if v, ok := p.Bio.Value(); ok && len(v) > 5000 {
			return authErr.ErrForgerProfileFieldInvalid
		}
	}
	return nil
}

// DiffAgainst keeps only incoming fields that differ from u; unchanged values become Unset.
func (p ForgerProfilePatch) DiffAgainst(u *ForgerProfile) ForgerProfilePatch {
	return ForgerProfilePatch{
		ProfileLanguage: p.ProfileLanguage.DiffVal(u.ProfileLanguage()),
		DisplayName:     diffStringPatchField(p.DisplayName, u.DisplayName()),
		FirstName:       diffStringPatchField(p.FirstName, u.FirstName()),
		LastName:        diffStringPatchField(p.LastName, u.LastName()),
		Country:         diffStringPatchField(p.Country, u.Country()),
		City:            diffStringPatchField(p.City, u.City()),
		Gender:          diffStringPatchField(p.Gender, u.Gender()),
		Birthday:        p.Birthday.DiffPtr(u.Birthday()),
		Website:         diffStringPatchField(p.Website, u.Website()),
		Timezone:        diffStringPatchField(p.Timezone, u.Timezone()),
		Bio:             diffStringPatchField(p.Bio, u.Bio()),
	}
}

func (u *ForgerProfile) ApplyPatch(p ForgerProfilePatch) error {
	if err := u.EnsureNotDeleted(); err != nil {
		return err
	}

	patch.ApplyVal(&u.state.ProfileLanguage, p.ProfileLanguage)
	patch.ApplyPtr(&u.state.DisplayName, p.DisplayName)
	patch.ApplyPtr(&u.state.FirstName, p.FirstName)
	patch.ApplyPtr(&u.state.LastName, p.LastName)
	patch.ApplyPtr(&u.state.Country, p.Country)
	patch.ApplyPtr(&u.state.City, p.City)
	patch.ApplyPtr(&u.state.Gender, p.Gender)
	patch.ApplyPtr(&u.state.Website, p.Website)
	patch.ApplyPtr(&u.state.Timezone, p.Timezone)
	patch.ApplyPtrWith(&u.state.Bio, p.Bio, strings.TrimSpace)
	patch.ApplyPtr(&u.state.Birthday, p.Birthday)

	u.state.UpdatedAt = time.Now().UTC()
	return nil
}
