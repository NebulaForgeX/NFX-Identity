package reqdto

import (
	"strings"
	"time"

	"nfxidentity/constants"
	"nfxidentity/enums"
	authErr "nfxidentity/errors/src/auth"
	profileDomain "nfxidentity/modules/auth/domain/profile"
	"nfxidentity/pkgs/patch"

	"github.com/google/uuid"
)

// CreateCommunityProfile is the request body for POST /auth/me/profiles.
type CreateCommunityProfile struct {
	DisplayName     string                    `json:"display_name"`
	ProfileLanguage enums.AuthProfileLanguage `json:"profile_language"`
}

func (r CreateCommunityProfile) Normalize() CreateCommunityProfile {
	return CreateCommunityProfile{
		DisplayName:     strings.TrimSpace(r.DisplayName),
		ProfileLanguage: r.ProfileLanguage,
	}
}

func (r CreateCommunityProfile) Validate() error {
	if r.DisplayName == "" {
		return authErr.ErrForgerProfileFieldInvalid
	}
	if len(r.DisplayName) > 150 {
		return authErr.ErrForgerProfileFieldInvalid
	}
	if r.ProfileLanguage != "" && !constants.AuthLanguage.Valid(r.ProfileLanguage) {
		return authErr.ErrForgerProfileLanguageInvalid
	}
	return nil
}

// DeleteCommunityProfileURI binds the profile id path param for DELETE /auth/me/profiles/:profileId.
type DeleteCommunityProfileURI struct {
	ProfileID uuid.UUID `uri:"profileId"`
}

// PublicProfileCardURI binds the profile id path param for GET /auth/me/profiles/:profileId/public-card.
type PublicProfileCardURI struct {
	ProfileID uuid.UUID `uri:"profileId"`
}

// UpdatePreference is the request body for PUT /auth/me/forger-profile/preference.
type UpdatePreference struct {
	Preference string `json:"preference"`
}

// ConfirmCommunityProfileAvatar is the request body for PUT /auth/me/forger-profile/avatars.
type ConfirmCommunityProfileAvatar struct {
	ImageID string `json:"image_id"`
}

// PatchCommunityProfile is the request body for PATCH /auth/me/forger-profile.
type PatchCommunityProfile struct {
	ProfileLanguage patch.Field[enums.AuthProfileLanguage] `json:"profile_language"`
	DisplayName     patch.Field[string]                    `json:"display_name"`
	FirstName       patch.Field[string]                    `json:"first_name"`
	LastName        patch.Field[string]                    `json:"last_name"`
	Country         patch.Field[string]                    `json:"country"`
	City            patch.Field[string]                    `json:"city"`
	Gender          patch.Field[string]                    `json:"gender"`
	Birthday        patch.Field[time.Time]                 `json:"birthday"`
	Website         patch.Field[string]                    `json:"website"`
	Timezone        patch.Field[string]                    `json:"timezone"`
	Bio             patch.Field[string]                    `json:"bio"`
}

func (r PatchCommunityProfile) ToPatch() profileDomain.ForgerProfilePatch {
	return profileDomain.ForgerProfilePatch{
		ProfileLanguage: r.ProfileLanguage.PatchField,
		DisplayName:     normalizeOptionalStringPatch(r.DisplayName.PatchField),
		FirstName:       normalizeOptionalStringPatch(r.FirstName.PatchField),
		LastName:        normalizeOptionalStringPatch(r.LastName.PatchField),
		Country:         normalizeOptionalStringPatch(r.Country.PatchField),
		City:            normalizeOptionalStringPatch(r.City.PatchField),
		Gender:          normalizeOptionalStringPatch(r.Gender.PatchField),
		Birthday:        r.Birthday.PatchField,
		Website:         normalizeOptionalStringPatch(r.Website.PatchField),
		Timezone:        normalizeOptionalStringPatch(r.Timezone.PatchField),
		Bio:             normalizeOptionalStringPatch(r.Bio.PatchField),
	}
}

func normalizeOptionalStringPatch(f patch.PatchField[string]) patch.PatchField[string] {
	if f.IsUnset() {
		return patch.Unset[string]()
	}
	if f.IsNull() {
		return patch.SetNull[string]()
	}
	if v, ok := f.Value(); ok {
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			return patch.SetNull[string]()
		}
		return patch.Set(trimmed)
	}
	return patch.Unset[string]()
}

// PatchCommunityProfileSettings is the request body for PATCH /auth/me/forger-profile-settings.
type PatchCommunityProfileSettings struct {
	LoginNotification patch.Field[bool] `json:"login_notification"`
}

func (r PatchCommunityProfileSettings) ToPatch() profileDomain.ForgerProfileSettingsPatch {
	return profileDomain.ForgerProfileSettingsPatch{
		LoginNotification: r.LoginNotification.PatchField,
	}
}
